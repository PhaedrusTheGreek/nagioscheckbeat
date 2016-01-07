package beat

import (
	"github.com/PhaedrusTheGreek/nagioscheckbeat/nagiosperf"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/mattn/go-shellwords"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type NagiosCheckBeat struct {
	// Config
	interval string
	checks   []NagiosCheckConfig

	// State
	isAlive  bool
	duration time.Duration

	// Handles
	events publisher.Client
}

func New() *NagiosCheckBeat {
	return &NagiosCheckBeat{}
}

func (nagiosCheck *NagiosCheckBeat) Config(b *beat.Beat) error {
	var config ConfigSettings
	err := cfgfile.Read(&config, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	nagiosCheck.checks = config.Input.Checks

	if config.Input.Interval != nil {
		nagiosCheck.interval = *config.Input.Interval
	} else {
		nagiosCheck.interval = "5s"
	}

	duration, err := time.ParseDuration(nagiosCheck.interval)
	if err != nil {
		logp.Err("Invalid Interval: %v", err)
		return err
	} else {
		nagiosCheck.duration = duration
	}

	return nil
}

func (nagiosCheck *NagiosCheckBeat) Setup(b *beat.Beat) error {

	nagiosCheck.events = b.Events

	return nil

}

func (nagiosCheck *NagiosCheckBeat) Run(b *beat.Beat) error {

	nagiosCheck.isAlive = true

	for nagiosCheck.isAlive {

		for _, check := range nagiosCheck.checks {

			for _, event := range Check(&check) {
				nagiosCheck.events.PublishEvent(event)
			}

		}

		time.Sleep(nagiosCheck.duration)

	}

	return nil
}

func Check(check *NagiosCheckConfig) (events []common.MapStr) {

	startTime := time.Now()
	startMs := startTime.UnixNano() / int64(time.Millisecond)

	if check == nil {
		logp.Err("Invalid/Missing Nagios Check Configuration")
		return
	}

	if check.Name == nil {
		logp.Err("Must Specify a Nagios Check Name")
		return
	}

	if check.Cmd == nil {
		logp.Err("Must Specify a Nagios Check Command")
		return
	}

	args := ""
	if check.Args != nil {
		args = *check.Args
	}

	check_event := common.MapStr{
		"@timestamp": common.Time(startTime),
		"type":       "nagioscheck",
		"cmd":        *check.Cmd,
		"args":       args,
	}

	logp.Debug("nagioscheck", "Running Command: %q", *check.Cmd)

	//arg_fields := strings.Fields(args)
	arg_fields, err := shellwords.Parse(args) // Smarter

	if err != nil {
		logp.Err("Could not parse args %q", args)
	}

	cmd := exec.Command(*check.Cmd, arg_fields...)
	var waitStatus syscall.WaitStatus

	/* Go will return 'err' if the command exits abnormally (non-zero return code).
	Nagios commands always will exit abnormally when a check fails, so from
	a funcational perspective, this doesn't help us.  Instead, if the ProcessState is nil,
	that tells us that the command coulnd't run for some reason, which does help.
	*/

	output, err := cmd.CombinedOutput()
	if cmd.ProcessState == nil {
		logp.Err("Command Error: %v", err)
		return
	}
	waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)

	logp.Debug("nagioscheck", "Command Returned: %q, exit code %d", output, waitStatus.ExitStatus())

	parts := strings.Split(string(output), "|")
	check_event["message"] = parts[0]
	check_event["status"] = nagiosperf.NiceStatus(waitStatus.ExitStatus())
	check_event["took_ms"] = time.Now().UnixNano()/int64(time.Millisecond) - startMs

	// publish the check result, even if there is no perf data
	events = append(events, check_event)

	if len(parts) > 1 {
		logp.Debug("nagioscheck", "Parsing: %q", parts[1])
		perfs, errors := nagiosperf.ParsePerfString(parts[1])
		if len(errors) > 0 {
			for _, err := range errors {
				logp.Err("Command Error: %v", err)
			}
		} else {

			logp.Debug("nagioscheck", "Command Returned '%d' Perf Metrics: %v", len(perfs), perfs)

			for _, perf := range perfs {

				metric_event := common.MapStr{
					"@timestamp": common.Time(startTime),
					"type":       "nagiosmetric",
					"name":       *check.Name,
					"label":      perf.Label,
					"uom":        perf.Uom,
					"value":      perf.Value,
					"min":        perf.Min,
					"max":        perf.Max,
					"warning":    perf.Warning,
					"critical":   perf.Critical,
				}

				events = append(events, metric_event)

			}
		}
	}

	return
}

func (nagiosCheck *NagiosCheckBeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (nagiosCheck *NagiosCheckBeat) Stop() {
	nagiosCheck.isAlive = false
}

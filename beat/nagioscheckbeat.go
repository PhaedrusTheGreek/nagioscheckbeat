package beat

import (
	"github.com/PhaedrusTheGreek/nagioscheckbeat/nagiosperf"
	"github.com/elastic/libbeat/beat"
	"github.com/elastic/libbeat/cfgfile"
	"github.com/elastic/libbeat/common"
	"github.com/elastic/libbeat/logp"
	"github.com/mattn/go-shellwords"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type NagiosCheckBeat struct {
	// from configuration
	interval string
	checks   []NagiosCheckConfig

	// state
	isAlive  bool
	duration time.Duration
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

	//TODO: Check if target script exists / test run

	return nil

}

func (nagiosCheck *NagiosCheckBeat) Run(b *beat.Beat) error {

	nagiosCheck.isAlive = true

	for nagiosCheck.isAlive {

		time.Sleep(nagiosCheck.duration)

		for _, check := range nagiosCheck.checks {

			startTime := time.Now()
			startMs := startTime.UnixNano() / int64(time.Millisecond)
			event := common.MapStr{
				"@timestamp": common.Time(startTime),
				"type":       "nagioscheck",
				"cmd":        *check.Cmd,
				"args":       *check.Args,
			}

			logp.Debug("nagioscheck", "Running Command: %q", *check.Cmd)

			//arg_fields := strings.Fields(*check.Args)
			arg_fields, err := shellwords.Parse(*check.Args)

			if err != nil {
				logp.Err("Could not parse args %q", *check.Args)
			}

			cmd := exec.Command(*check.Cmd, arg_fields...)
			var waitStatus syscall.WaitStatus

			/* Go will return 'err' if the command exits abnormally (non-zero return code).
			Nagios commands always will exit abnormally when a check fails, so from
			a funcational perspective, we don't care about that.
			*/

			output, _ := cmd.CombinedOutput()
			waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)

			logp.Debug("nagioscheck", "Command Returned: %q, exit code %d", output, waitStatus.ExitStatus())

			parts := strings.Split(string(output), "|")
			event["message"] = parts[0]
			if len(parts) > 1 {
				logp.Debug("nagioscheck", "Parsing: %q", parts[1])
				perfs, errors := nagiosperf.ParsePerfString(parts[1])
				if len(errors) > 0 {
					for _, err := range errors {
						logp.Err("Command Error: %v", err)
					}
				} else {
					logp.Debug("nagioscheck", "Command Returned '%d' Perf Metrics: %v", len(perfs), perfs)
					addPerfsToEvent(*check.Name, perfs, &event)
				}
			}

			event["status"] = nagiosperf.NiceStatus(waitStatus.ExitStatus())
			event["took_ms"] = time.Now().UnixNano()/int64(time.Millisecond) - startMs

			b.Events.PublishEvent(event)

		}
	}

	return nil
}

func addPerfsToEvent(name string, perfs []nagiosperf.Perf, event *common.MapStr) {

	var data = common.MapStr{}

	for _, perf := range perfs {
		data[string(perf.Label)] = perf
	}

	(*event)[name] = data

}

func (nagiosCheck *NagiosCheckBeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (nagiosCheck *NagiosCheckBeat) Stop() {
	nagiosCheck.isAlive = false
}

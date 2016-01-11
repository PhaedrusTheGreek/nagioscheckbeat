package check

import (
	"github.com/PhaedrusTheGreek/nagioscheckbeat/config"
	"github.com/elastic/beats/libbeat/common"
	"strings"
	"testing"
)

func TestCheck(t *testing.T) {

	/*
		1st Check, makes sure command runs and returns OK
	*/
	cmd := "/usr/local/sbin/check_dummy"
	arg := "0 hello"
	period := "1s"
	name := "dummy"

	checkConfig := config.NagiosCheckConfig{Cmd: &cmd, Args: &arg, Period: &period, Name: &name}

	check := NagiosCheck{}
	check.Setup(&checkConfig)
	got, err := check.Check()

	if err != nil {
		t.Errorf("%v", err)
	}

	// [{"@timestamp":"2016-01-10T05:40:19.738Z","args":"0 hello","cmd":"/usr/local/sbin/check_dummy","message":"OK: hello\n","status":"OK","took_ms":6,"type":"nagioscheck"}]

	var event common.MapStr = got[0]
	var message = event["message"].(string)
	if strings.Compare(message, "OK: hello\n") != 0 {
		t.Errorf("Expected 'OK: hello\\n', and got %q", message)
	}

	/*
		2nd Check - makes sure perf data comes along
	*/
	cmd = "/usr/local/sbin/check_load"
	arg = "-w 5 -c 10"
	period = "1s"
	name = "load"

	checkConfig = config.NagiosCheckConfig{Cmd: &cmd, Args: &arg, Period: &period, Name: &name}

	check = NagiosCheck{}
	check.Setup(&checkConfig)
	got, err = check.Check()

	if err != nil {
		t.Errorf("%v", err)
	}

	if len(got) != 4 {
		t.Errorf("Expected 4 events, but got %d", len(got))
	}

}

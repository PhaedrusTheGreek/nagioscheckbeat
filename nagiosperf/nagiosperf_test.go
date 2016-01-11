package nagiosperf

import "testing"
import "reflect"

func TestParse(t *testing.T) {

	cases := []struct {
		in   string
		want []Perf
	}{
		{"load1=0.010;5.000;9.000;0; load5=0.060;5.000;9.000;0; load15=2.010;5.000;9.000;0;", []Perf{
			{Label: "load1", Value: 0.01, Warning: 5, Critical: 9, Min: 0},
			{Label: "load5", Value: 0.06, Warning: 5, Critical: 9, Min: 0},
			{Label: "load15", Value: 2.01, Warning: 5, Critical: 9, Min: 0},
		}},
		{"users=4;20;50;0", []Perf{
			{Label: "users", Value: 4, Warning: 20, Critical: 50, Min: 0},
		}},
		{"/home/a-m=0;0;0 shared-folder:big=20 12345678901234567890=20", []Perf{
			{Label: "/home/a-m", Value: 0, Warning: 0, Critical: 0},
			{Label: "shared-folder:big", Value: 20},
			{Label: "12345678901234567890", Value: 20},
		}},
		{"time=0.002722s;0.000000;0.000000;0.000000;10.000000", []Perf{
			{Label: "time", Value: 0.002722, Uom: "s", Max: 10},
		}},
		{"'Waiting for Connection'=22 'Starting Up'=1\n", []Perf{
			{Label: "Waiting for Connection", Value: 22 },
			{Label: "Starting Up", Value: 1 },
		}},
		{"", []Perf{}},
	}
	for _, c := range cases {
		got, errors := ParsePerfString(c.in)
		if len(errors) > 0 {
			for _, err := range errors {
				t.Errorf("ParsePerfString(%q) returned error: '%s'", c.in, err)
			}
			continue
		}

		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("ParsePerfString(%q):\n WANT: %v\n GOT: %v", c.in, c.want, got)
			continue
		}
	}
}

package nagiosperf

/*
Partial Go Implementation of:
http://cpansearch.perl.org/src/NIERLEIN/Monitoring-Plugin-0.39/lib/Monitoring/Plugin/Performance.pm
*/

import (
	"errors"
	"regexp"
	"strconv"
)

type Perf struct {
	Label string
	Uom   string

	Value    float64
	Warning  float64
	Critical float64
	Min      float64
	Max      float64
}

func NiceStatus(status int) string {
	switch status {
	case 0:
		return "OK"
	case 1:
		return "WARNING"
	case 2:
		return "CRITICAL"
	case 3:
		return "UNKNOWN"
	}
	return "INVALID STATUS"
}

func parse(perfString string) (Perf, error) {

	const (
		Value                        = `[-+]?[\d\.,]+`
		Value_re                     = Value + `(?:e` + Value + `)?`
		Value_with_negative_infinity = Value_re + `|~`
	)

	var perf Perf

	re := regexp.MustCompile(`^'?([^'=]+)'?=(` + Value_re + `)([\w%]*);?(` + Value_with_negative_infinity + `\:?` + Value_re + `?)?;?(` + Value_with_negative_infinity + `\:?` + Value_re + `?)?;?(` + Value_re + `)?;?(` + Value_re + `)?`)

	results := re.FindStringSubmatch(perfString)
	if results == nil || len(results) < 3 {
		return perf, errors.New("Cannot Parse: " + perfString)
	}

	perf = Perf{}

	perf.Label = results[1]
	perf.Uom = results[3]

	if v, err := strconv.ParseFloat(results[2], 64); err == nil {
		perf.Value = v
	} else if results[2] != "" {
		return perf, err
	}

	if v, err := strconv.ParseFloat(results[4], 64); err == nil {
		perf.Warning = v
	} else if results[4] != "" {
		return perf, err
	}

	if v, err := strconv.ParseFloat(results[5], 64); err == nil {
		perf.Critical = v
	} else if results[5] != "" {
		return perf, err
	}

	if v, err := strconv.ParseFloat(results[6], 64); err == nil {
		perf.Min = v
	} else if results[6] != "" {
		return perf, err
	}

	if v, err := strconv.ParseFloat(results[7], 64); err == nil {
		perf.Max = v
	} else if results[7] != "" {
		return perf, err
	}

	return perf, nil

}

func ParsePerfString(perfString string) ([]Perf, []error) {

	var errors []error

	perfs := []Perf{}

	for _, element := range eachPerf(perfString) {

		perf, err := parse(element)

		perfs = append(perfs, perf)

		if err != nil {
			errors = append(errors, err)
		}

	}
	return perfs, errors

}

/*
Splits string by spaces, considering quotes
*/
func eachPerf(perfString string) []string {
	ir := regexp.MustCompile("'.+'|\".+\"|\\S+")
	return ir.FindAllString(perfString, -1)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func initializeGama(size int) []map[string]int {

	gama := make([]map[string]int, 0)

	for i := 0; i < 12; i++ {
		gama = append(gama,
			map[string]int{
				"0": int(0),
				"1": int(0),
			})
	}

	return gama
}

type DiagnosticReport struct {
	Gama  []map[string]int
	input []string
}

func (d *DiagnosticReport) process(input []string) {
	for _, line := range input {
		fmt.Println(line)
		diag := strings.Split(line, "")

		for i, bit := range diag {
			d.Gama[i][bit] = d.Gama[i][bit] + 1
		}
	}
}

func (d *DiagnosticReport) oxygenGeneratorRating() int {
	return d.part2BitCriteria(d.input, 0, oxygenCriteria)
}

func (d *DiagnosticReport) co2Rating() int {
	return d.part2BitCriteria(d.input, 0, co2Criteria)
}

func oxygenCriteria(counter map[string]int) string {
	if counter["1"] > counter["0"] || counter["1"] == counter["0"] {
		return "1"
	} else {
		return "0"
	}
}

func co2Criteria(counter map[string]int) string {
	if counter["1"] == counter["0"] {
		return "0"
	}

	if counter["1"] > counter["0"] {
		return "0"
	} else {
		return "1"
	}
}

func filterCommon(input []string, position int, filterBit string) []string {

	filteredInput := []string{}

	for _, line := range input {
		if string(line[position]) == filterBit {
			filteredInput = append(filteredInput, line)
		}
	}

	return filteredInput
}

func (d *DiagnosticReport) part2BitCriteria(input []string, position int, criteria func(map[string]int) string) int {
	if len(input) < 1 {
		return 0 // we shouldn't get here
	}

	if len(input) == 1 {
		value, _ := strconv.ParseInt(input[0], 2, 32)
		return int(value)
	}

	counter := map[string]int{
		"0": int(0),
		"1": int(0),
	}

	for _, element := range input {
		bits := strings.Split(element, "")
		counter[bits[position]] = counter[bits[position]] + 1
	}

	filterBit := criteria(counter)
	common := filterCommon(input, position, filterBit)

	return d.part2BitCriteria(common, position+1, criteria)
}

func NewDiagnosticReport(input []string) DiagnosticReport {
	gama := initializeGama(12)
	report := DiagnosticReport{
		Gama:  gama,
		input: input,
	}
	report.process(input)
	return report
}

func (d *DiagnosticReport) EpsilonRate() int {
	rate := ""

	for _, element := range d.Gama {
		if element["0"] < element["1"] {
			rate = rate + "0"
		} else {
			rate = rate + "1"
		}
	}

	value, _ := strconv.ParseInt(rate, 2, 32)

	return int(value)
}

func (d *DiagnosticReport) GamaRate() int {
	rate := ""

	for _, element := range d.Gama {
		if element["0"] > element["1"] {
			rate = rate + "0"
		} else {
			rate = rate + "1"
		}
	}

	value, _ := strconv.ParseInt(rate, 2, 32)

	return int(value)
}

func (d *DiagnosticReport) PowerConsumption() int {
	return d.GamaRate() * d.EpsilonRate()
}

func (d *DiagnosticReport) LifeSupportRating() int {
	return d.co2Rating() * d.oxygenGeneratorRating()
}

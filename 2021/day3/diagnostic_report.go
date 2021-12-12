package main

import (
	"fmt"
	"strings"
	"strconv"
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
	Gama []map[string]int
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

func NewDiagnosticReport(input []string) DiagnosticReport {
	gama := initializeGama(12)
	report := DiagnosticReport{
		Gama: gama,
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

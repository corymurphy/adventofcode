package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	lines := readLines("input")

	streamStart := true

	startIndex := 0
	endIndex := 0

	passports := Passports{}

	valid1 := 0
	valid2 := 0

	for i, line := range lines {

		if streamStart {
			startIndex = i
		}

		if line == "" {
			endIndex = i
			passport := NewPassport(lines[startIndex:endIndex])
			passports = append(passports, passport)

			if passport.Valid1() {
				valid1++
			}

			if passport.Valid2() {
				valid2++
			}
			streamStart = true
			continue
		} else {
			streamStart = false
		}
	}

	log.Printf("valid1 passport count: %s\n", fmt.Sprint(valid1))
	log.Printf("valid2 passport count: %s\n", fmt.Sprint(valid2))
}

type Passports []Passport

func contains(char string) bool {

	validHex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

	for _, item := range validHex {
		if char == item {
			return true
		}
	}

	return false
}

func main() {
	part1()

	lines := []string{"eyr:2033", "hgt:177cm pid:173cm", "ecl:utc byr:2029 hcl:#efcc98 iyr:2023", "cid:derp"}

	passport := NewPassport(lines)

	log.Printf("valid passport: %s\n", fmt.Sprint(passport.Valid1()))

	log.Printf("byr: %s\n", fmt.Sprint(int64(2000)))

	//ValidIntRange(p.BirthYear, 1920, 2002)

	log.Printf("byr valid range: %s\n", ValidIntRange("2000", 1920, 2002))

	log.Printf("is valid: %s\n", contains("x"))

	passport.Print()
}

type Passport struct {
	BirthYear      string // byr
	IssueYear      string // iyr
	ExpirationYear string // eyr
	Height         string // hgt
	HairColor      string // hcl
	EyeColor       string // ecl
	PassportId     string // pid
	CountryId      string // cid
}

func (p Passport) Valid1() bool {
	if len(p.BirthYear) != 0 && len(p.IssueYear) != 0 && len(p.ExpirationYear) != 0 && len(p.Height) != 0 && len(p.HairColor) != 0 && len(p.EyeColor) != 0 && len(p.PassportId) != 0 {
		return true
	} else {
		return false
	}
}

func (p Passport) Valid2() bool {
	if p.Byr() && p.Iyr() && p.Eyr() && p.Hgt() && p.Hcl() && p.Ecl() && p.Pid() {
		return true
	} else {
		return false
	}
}

func (p Passport) Print() {
	fmt.Printf("byr: %s\n", p.BirthYear)
}

func (p Passport) Byr() bool {

	if len(p.BirthYear) <= 0 {
		return false
	}

	if ValidIntRange(p.BirthYear, 1920, 2002) {
		return true
	} else {
		return false
	}
}

func ValidIntRange(value string, start int, end int) bool {
	if val, err := strconv.ParseInt(value, 10, 64); err != nil {
		return false
	} else {
		if val < int64(start) || val > int64(end) {
			return false
		} else {
			return true
		}
	}
}

func (p Passport) Iyr() bool {

	if len(p.IssueYear) <= 0 {
		return false
	}

	if ValidIntRange(p.IssueYear, 2010, 2020) {
		return true
	} else {
		return false
	}
}

func (p Passport) Eyr() bool {

	if len(p.ExpirationYear) <= 0 {
		return false
	}

	if ValidIntRange(p.ExpirationYear, 2020, 2030) {
		return true
	} else {
		return false
	}
}

func (p Passport) Hgt() bool {

	if len(p.Height) <= 0 {
		return false
	}

	if strings.HasSuffix(p.Height, "cm") {
		if ValidIntRange(strings.Replace(p.Height, "cm", "", -1), 150, 193) {
			return true
		} else {
			return false
		}
	}

	if strings.HasSuffix(p.Height, "in") {
		if ValidIntRange(strings.Replace(p.Height, "in", "", -1), 59, 76) {
			return true
		} else {
			return false
		}
	}

	return false
}

func (p Passport) Pid() bool {
	if len(p.PassportId) != 9 {
		return false
	}

	if val, err := strconv.ParseInt(p.PassportId, 10, 64); err != nil {

		return false
	} else {
		val++
	}

	return true
}

func (p Passport) Hcl() bool {
	if len(p.HairColor) != 7 {
		return false
	}

	hairColor := strings.Replace(p.HairColor, "#", "", -1)

	for _, char := range strings.Split(hairColor, "") {
		if !contains(char) {
			return false
		}
	}

	return true
}

func (p Passport) Ecl() bool {
	valids := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, valid := range valids {
		if p.EyeColor == valid {
			return true
		}
	}
	return false
}

func NewPassport(slice []string) Passport {
	passport := Passport{}

	for _, line := range slice {

		splitLine := strings.Split(line, " ")

		for _, item := range splitLine {

			if strings.HasPrefix(item, "byr") {
				passport.BirthYear = strings.Split(item, ":")[1]
			}

			if strings.HasPrefix(item, "pid") {
				passport.PassportId = strings.Split(item, ":")[1]
			}

			if strings.HasPrefix(item, "iyr") {
				passport.IssueYear = strings.Split(item, ":")[1]
			}

			if strings.HasPrefix(item, "eyr") {
				passport.ExpirationYear = strings.Split(item, ":")[1]
			}

			if strings.HasPrefix(item, "hgt") {
				passport.Height = strings.Split(item, ":")[1]
			}

			if strings.HasPrefix(item, "hcl") {
				passport.HairColor = strings.Split(item, ":")[1]
			}

			if strings.HasPrefix(item, "ecl") {
				passport.EyeColor = strings.Split(item, ":")[1]
			}

			if strings.HasPrefix(item, "cid") {
				passport.CountryId = strings.Split(item, ":")[1]
			}
		}
	}

	return passport
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

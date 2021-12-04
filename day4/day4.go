package day4

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	BirthYear      = "byr"
	IssueYear      = "iyr"
	ExpirationYear = "eyr"
	Height         = "hgt"
	HairColor      = "hcl"
	EyeColor       = "ecl"
	PassportID     = "pid"
	CountryID      = "cid"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func ParsePassport(str string) Passport {
	p := strings.SplitN(str, " ", 8)

	var pp Passport
	for _, v := range p {
		kv := strings.SplitN(v, ":", 2)

		if kv[0] == BirthYear {
			pp.BirthYear = kv[1]
		}
		if kv[0] == IssueYear {
			pp.IssueYear = kv[1]
		}
		if kv[0] == ExpirationYear {
			pp.ExpirationYear = kv[1]
		}
		if kv[0] == Height {
			pp.Height = kv[1]
		}
		if kv[0] == HairColor {
			pp.HairColor = kv[1]
		}
		if kv[0] == EyeColor {
			pp.EyeColor = kv[1]
		}
		if kv[0] == PassportID {
			pp.PassportID = kv[1]
		}
		if kv[0] == CountryID {
			pp.CountryID = kv[1]
		}
	}

	return pp
}

func CountValidPassports(lines []string, validate func(Passport) bool) int {
	var r int
	var s []string

	var ml string
	for i, v := range lines {
		if v == "" {
			s = append(s, ml)
			ml = ""
			continue
		}
		if ml != "" {
			ml += " "
		}
		ml += v

		if i == len(lines)-1 {
			s = append(s, ml)
			ml = ""
		}
	}

	for _, v := range s {
		p := ParsePassport(v)
		if validate(p) {
			r++
		}
	}

	return r
}

func ValidatePassport(p Passport) bool {

	if p.BirthYear == "" {
		return false
	}
	if p.IssueYear == "" {
		return false
	}
	if p.ExpirationYear == "" {
		return false
	}
	if p.Height == "" {
		return false
	}
	if p.HairColor == "" {
		return false
	}
	if p.EyeColor == "" {
		return false
	}
	if p.PassportID == "" {
		return false
	}
	if p.CountryID == "" {
		// noop
		// missing CountryID is ok
	}

	return true
}

var (
	hairRegexp = regexp.MustCompile(`^[#][0-9a-f]{6}$`)
	pidRegexp  = regexp.MustCompile(`^[0-9]{9}$`)
	eyeRegexp  = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
)

func ValidatePassportStrict(p Passport) bool {

	if !ValidatePassport(p) {
		return false
	}

	{
		v := parseNumber(p.BirthYear)
		if v < 1920 || v > 2002 {
			return false
		}
	}
	{
		v := parseNumber(p.IssueYear)
		if v < 2010 || v > 2020 {
			return false
		}
	}
	{
		v := parseNumber(p.ExpirationYear)
		if v < 2020 || v > 2030 {
			return false
		}
	}
	{
		if strings.HasSuffix(p.Height, "cm") {
			v := parseNumber(strings.TrimSuffix(p.Height, "cm"))
			if v < 150 || v > 193 {
				return false
			}
		} else if strings.HasSuffix(p.Height, "in") {
			v := parseNumber(strings.TrimSuffix(p.Height, "in"))
			if v < 59 || v > 76 {
				return false
			}
		} else {
			return false
		}
	}
	{
		if !hairRegexp.MatchString(p.HairColor) {
			return false
		}
	}
	{
		if !eyeRegexp.MatchString(p.EyeColor) {
			return false
		}
	}
	{
		if !pidRegexp.MatchString(p.PassportID) {
			return false
		}

		v := parseNumber(p.PassportID)
		if v <= 0 {
			return false
		}
	}

	// CountryID is skipped

	return true
}

// parseNumber .
//
// Returns -1 on parse error.
func parseNumber(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1
	}
	return v
}

package day2

import (
	"strconv"
	"strings"
)

type IPolicy interface {
	Assert(pwd string) bool
}

// AssertPolicy .
//
// Returns count of valid passwords.
//
func AssertPolicy(lines []string, policyBuilder func(string) IPolicy) int {
	var i int

	for _, v := range lines {
		pwd := v[strings.Index(v, ": ")+2:]

		p := policyBuilder(v)
		if p.Assert(pwd) {
			i++
		}
	}

	return i
}

type Policy1 struct {
	char     uint8
	min, max int64 // char must be repeated...
}

func NewParsePolicy1(s string) *Policy1 {
	p := strings.SplitN(s[0:strings.Index(s, ":")], " ", 2)
	n := strings.SplitN(p[0], "-", 2)

	var policy Policy1

	policy.char = p[1][0]

	// lets skip errors
	policy.min, _ = strconv.ParseInt(n[0], 10, 64)
	policy.max, _ = strconv.ParseInt(n[1], 10, 64)

	return &policy
}

func (s *Policy1) Assert(password string) bool {
	c := strings.Count(password, string(s.char))

	return int64(c) >= s.min && int64(c) <= s.max
}

type Policy2 struct {
	char       uint8
	idx1, idx2 int64
}

func NewParsePolicy2(s string) *Policy2 {
	p := strings.SplitN(s[0:strings.Index(s, ":")], " ", 2)
	n := strings.SplitN(p[0], "-", 2)

	var policy Policy2

	policy.char = p[1][0]

	// lets skip errors
	policy.idx1, _ = strconv.ParseInt(n[0], 10, 64)
	policy.idx2, _ = strconv.ParseInt(n[1], 10, 64)

	policy.idx1--
	policy.idx2--

	return &policy
}

func (s *Policy2) Assert(password string) bool {
	var a, b bool

	if int64(len(password)) >= s.idx1 {
		a = password[s.idx1] == s.char
	}

	if int64(len(password)) >= s.idx2 {
		b = password[s.idx2] == s.char
	}

	return a != b
}

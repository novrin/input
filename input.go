// Package input provides utilities for checking and
// validating input data fields.
package input

import (
	"net/url"
	"strconv"
	"time"
	"unicode/utf8"
)

// Rule represents a validation rule for a field.
// If OK is false, Message is recorded as the error reason.
type Rule struct {
	OK      bool   // whether the rule passed
	Message string // message to record if the rule fails
}

// Log collects validation errors for input fields.
//
// Typically, Log is embedded into a struct to track
// validation errors. After performing checks, use OK()
// to see if all rules passed.
type Log struct {
	// Errors maps field names to a list of error messages.
	Errors map[string][]string
}

// OK returns true if no errors have been recorded in Log.
func (log *Log) OK() bool { return len(log.Errors) == 0 }

// Check applies a set of rules to a field.
// If a rule fails (OK is false), its Message is added to
// the field's error list.
func (log *Log) Check(field string, rules ...Rule) {
	if log.Errors == nil {
		log.Errors = make(map[string][]string)
	}
	for _, rule := range rules {
		if !rule.OK {
			log.Errors[field] = append(log.Errors[field], rule.Message)
		}
	}
}

// IsMember returns true if s is found in ss.
func IsMember(s string, ss []string) bool {
	for _, x := range ss {
		if s == x {
			return true
		}
	}
	return false
}

// IsInCharLimit returns true if the rune count of s is
// between min and max, inclusive.
func IsInCharLimit(s string, min int, max int) bool {
	x := utf8.RuneCountInString(s)
	return x >= min && x <= max
}

// IsTime returns true if s can be parsed as a time value using layout.
func IsTime(s string, layout string) bool {
	_, err := time.Parse(layout, s)
	return err == nil
}

// IsTimePast returns true if s can be parsed as a time
// value using layout and is before the current time.
func IsTimePast(s string, layout string) bool {
	t, err := time.Parse(layout, s)
	if err != nil {
		return false
	}
	return t.UTC().Before(time.Now().UTC())
}

// IsTimeFuture returns true if s can be parsed as a time
// value using layout and is after the current time.
func IsTimeFuture(s string, layout string) bool {
	t, err := time.Parse(layout, s)
	if err != nil {
		return false
	}
	return t.UTC().After(time.Now().UTC())
}

// IsURL returns true if s can be parsed into a URL structure.
func IsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

// IsBool returns true if s can be parsed as a boolean value.
func IsBool(s string) bool {
	_, err := strconv.ParseBool(s)
	return err == nil
}

// IsInt returns true if s can be parsed as an int of the
// given base and bit size.
func IsInt(s string, base int, bitSize int) bool {
	_, err := strconv.ParseInt(s, base, bitSize)
	return err == nil
}

// IsUint returns true if s can be parsed as a uint of the
// given base and bit size. A sign prefix is not permitted.
func IsUint(s string, base int, bitSize int) bool {
	_, err := strconv.ParseUint(s, base, bitSize)
	return err == nil
}

// IsFloat returns true if s can be parsed as a float of
// the given bit size.
func IsFloat(s string, bitSize int) bool {
	_, err := strconv.ParseFloat(s, bitSize)
	return err == nil
}

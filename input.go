// Package input provides utilities for checking and
// validating input data fields.
package input

import (
	"net/url"
	"strconv"
	"time"
	"unicode/utf8"
)

// Check evaluates a field.
//
// If 'ok' is false, message is appended to errs[field].
// If errs is nil and a message needs to be added, a new map
// is allocated and returned.
//
// Check returns the updated map. It is therefore necessary
// to store the result of Check, often in the variable
// holding the map itself.
//
// Example:
//
//	var errors map[string][]string // nil map
//	errors = Check(errors, "name", name != "", "name cannot be blank")
//	errors now contains: map[string][]string{"name": {"name cannot be blank"}}
func Check(errs map[string][]string, field string, ok bool, message string) map[string][]string {
	if errs == nil {
		errs = make(map[string][]string)
	}
	if !ok {
		errs[field] = append(errs[field], message)
	}
	return errs
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

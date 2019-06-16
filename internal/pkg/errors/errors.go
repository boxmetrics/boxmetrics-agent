// Package errors implements functions to manipulate errors.
package errors

import (
	"encoding/json"
)

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// Convert error to a new error with marshall support
func Convert(e error) error {
	if e != nil {
		return &errorString{e.Error()}
	}

	return e
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func (e *errorString) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	e.s = s

	return nil
}

// MarshalJSON return JSON
func (e errorString) MarshalJSON() ([]byte, error) {
	var s string

	s = e.s

	return json.Marshal(s)
}

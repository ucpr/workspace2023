// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: example.proto

package gen

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on TestMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TestMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TestMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TestMessageMultiError, or
// nil if none found.
func (m *TestMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *TestMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetExample()))
		i := 0
		for key := range m.GetExample() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetExample()[key]
			_ = val

			if _, ok := _TestMessage_Example_InLookup[key]; !ok {
				err := TestMessageValidationError{
					field:  fmt.Sprintf("Example[%v]", key),
					reason: "value must be in list [missing missing2]",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			// no validation rules for Example[key]
		}
	}

	if _, ok := _TestMessage_Example2_InLookup[m.GetExample2()]; !ok {
		err := TestMessageValidationError{
			field:  "Example2",
			reason: "value must be in list [foo bar baz]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	{
		sorted_keys := make([]string, len(m.GetExample3()))
		i := 0
		for key := range m.GetExample3() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetExample3()[key]
			_ = val

			// no validation rules for Example3[key]

			if utf8.RuneCountInString(val) < 3 {
				err := TestMessageValidationError{
					field:  fmt.Sprintf("Example3[%v]", key),
					reason: "value length must be at least 3 runes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	{
		sorted_keys := make([]string, len(m.GetExample4()))
		i := 0
		for key := range m.GetExample4() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetExample4()[key]
			_ = val

			if !_TestMessage_Example4_Pattern.MatchString(key) {
				err := TestMessageValidationError{
					field:  fmt.Sprintf("Example4[%v]", key),
					reason: "value does not match regex pattern \"^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$\"",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			// no validation rules for Example4[key]
		}
	}

	{
		sorted_keys := make([]string, len(m.GetExample5()))
		i := 0
		for key := range m.GetExample5() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetExample5()[key]
			_ = val

			if _, ok := _TestMessage_Example5_NotInLookup[key]; ok {
				err := TestMessageValidationError{
					field:  fmt.Sprintf("Example5[%v]", key),
					reason: "value must not be in list [hoge]",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			// no validation rules for Example5[key]
		}
	}

	{
		sorted_keys := make([]string, len(m.GetExample6()))
		i := 0
		for key := range m.GetExample6() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetExample6()[key]
			_ = val

			if _, ok := _TestMessage_Example6_InLookup[key]; !ok {
				err := TestMessageValidationError{
					field:  fmt.Sprintf("Example6[%v]", key),
					reason: "value must be in list [foo bar]",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if _, ok := _TestMessage_Example6_InLookup[val]; !ok {
				err := TestMessageValidationError{
					field:  fmt.Sprintf("Example6[%v]", key),
					reason: "value must be in list [baz qux]",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return TestMessageMultiError(errors)
	}

	return nil
}

// TestMessageMultiError is an error wrapping multiple validation errors
// returned by TestMessage.ValidateAll() if the designated constraints aren't met.
type TestMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestMessageMultiError) AllErrors() []error { return m }

// TestMessageValidationError is the validation error returned by
// TestMessage.Validate if the designated constraints aren't met.
type TestMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestMessageValidationError) ErrorName() string { return "TestMessageValidationError" }

// Error satisfies the builtin error interface
func (e TestMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestMessageValidationError{}

var _TestMessage_Example2_InLookup = map[string]struct{}{
	"foo": {},
	"bar": {},
	"baz": {},
}

var _TestMessage_Example4_Pattern = regexp.MustCompile("^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$")

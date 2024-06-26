// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/system/v1/sys_post.proto

package v1

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

// Validate checks the field values on SysPostRep with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysPostRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysPostRep with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysPostRepMultiError, or
// nil if none found.
func (m *SysPostRep) ValidateAll() error {
	return m.validate(true)
}

func (m *SysPostRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PostId

	// no validation rules for PostCode

	// no validation rules for PostName

	// no validation rules for PostSort

	// no validation rules for Status

	// no validation rules for Remark

	if len(errors) > 0 {
		return SysPostRepMultiError(errors)
	}

	return nil
}

// SysPostRepMultiError is an error wrapping multiple validation errors
// returned by SysPostRep.ValidateAll() if the designated constraints aren't met.
type SysPostRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysPostRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysPostRepMultiError) AllErrors() []error { return m }

// SysPostRepValidationError is the validation error returned by
// SysPostRep.Validate if the designated constraints aren't met.
type SysPostRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysPostRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysPostRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysPostRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysPostRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysPostRepValidationError) ErrorName() string { return "SysPostRepValidationError" }

// Error satisfies the builtin error interface
func (e SysPostRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysPostRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysPostRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysPostRepValidationError{}

// Validate checks the field values on DeleteSysPostRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteSysPostRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSysPostRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSysPostRepMultiError, or nil if none found.
func (m *DeleteSysPostRep) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSysPostRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := DeleteSysPostRepValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteSysPostRepMultiError(errors)
	}

	return nil
}

// DeleteSysPostRepMultiError is an error wrapping multiple validation errors
// returned by DeleteSysPostRep.ValidateAll() if the designated constraints
// aren't met.
type DeleteSysPostRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSysPostRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSysPostRepMultiError) AllErrors() []error { return m }

// DeleteSysPostRepValidationError is the validation error returned by
// DeleteSysPostRep.Validate if the designated constraints aren't met.
type DeleteSysPostRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSysPostRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSysPostRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSysPostRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSysPostRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSysPostRepValidationError) ErrorName() string { return "DeleteSysPostRepValidationError" }

// Error satisfies the builtin error interface
func (e DeleteSysPostRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSysPostRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSysPostRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSysPostRepValidationError{}

// Validate checks the field values on GetSysPostRep with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetSysPostRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSysPostRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetSysPostRepMultiError, or
// nil if none found.
func (m *GetSysPostRep) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSysPostRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := GetSysPostRepValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetSysPostRepMultiError(errors)
	}

	return nil
}

// GetSysPostRepMultiError is an error wrapping multiple validation errors
// returned by GetSysPostRep.ValidateAll() if the designated constraints
// aren't met.
type GetSysPostRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSysPostRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSysPostRepMultiError) AllErrors() []error { return m }

// GetSysPostRepValidationError is the validation error returned by
// GetSysPostRep.Validate if the designated constraints aren't met.
type GetSysPostRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSysPostRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSysPostRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSysPostRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSysPostRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSysPostRepValidationError) ErrorName() string { return "GetSysPostRepValidationError" }

// Error satisfies the builtin error interface
func (e GetSysPostRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSysPostRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSysPostRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSysPostRepValidationError{}

// Validate checks the field values on GetSysPostReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetSysPostReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSysPostReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSysPostReplyMultiError, or nil if none found.
func (m *GetSysPostReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSysPostReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetSysPostReplyValidationError{
					field:  "Post",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetSysPostReplyValidationError{
					field:  "Post",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetSysPostReplyValidationError{
				field:  "Post",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetSysPostReplyMultiError(errors)
	}

	return nil
}

// GetSysPostReplyMultiError is an error wrapping multiple validation errors
// returned by GetSysPostReply.ValidateAll() if the designated constraints
// aren't met.
type GetSysPostReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSysPostReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSysPostReplyMultiError) AllErrors() []error { return m }

// GetSysPostReplyValidationError is the validation error returned by
// GetSysPostReply.Validate if the designated constraints aren't met.
type GetSysPostReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSysPostReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSysPostReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSysPostReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSysPostReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSysPostReplyValidationError) ErrorName() string { return "GetSysPostReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetSysPostReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSysPostReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSysPostReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSysPostReplyValidationError{}

// Validate checks the field values on ListSysPostRep with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListSysPostRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSysPostRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListSysPostRepMultiError,
// or nil if none found.
func (m *ListSysPostRep) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSysPostRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPageNum() <= 0 {
		err := ListSysPostRepValidationError{
			field:  "PageNum",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPageSize() <= 0 {
		err := ListSysPostRepValidationError{
			field:  "PageSize",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListSysPostRepMultiError(errors)
	}

	return nil
}

// ListSysPostRepMultiError is an error wrapping multiple validation errors
// returned by ListSysPostRep.ValidateAll() if the designated constraints
// aren't met.
type ListSysPostRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSysPostRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSysPostRepMultiError) AllErrors() []error { return m }

// ListSysPostRepValidationError is the validation error returned by
// ListSysPostRep.Validate if the designated constraints aren't met.
type ListSysPostRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSysPostRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSysPostRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSysPostRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSysPostRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSysPostRepValidationError) ErrorName() string { return "ListSysPostRepValidationError" }

// Error satisfies the builtin error interface
func (e ListSysPostRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSysPostRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSysPostRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSysPostRepValidationError{}

// Validate checks the field values on ListSysPostReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListSysPostReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSysPostReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSysPostReplyMultiError, or nil if none found.
func (m *ListSysPostReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSysPostReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetRows() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListSysPostReplyValidationError{
						field:  fmt.Sprintf("Rows[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListSysPostReplyValidationError{
						field:  fmt.Sprintf("Rows[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSysPostReplyValidationError{
					field:  fmt.Sprintf("Rows[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListSysPostReplyMultiError(errors)
	}

	return nil
}

// ListSysPostReplyMultiError is an error wrapping multiple validation errors
// returned by ListSysPostReply.ValidateAll() if the designated constraints
// aren't met.
type ListSysPostReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSysPostReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSysPostReplyMultiError) AllErrors() []error { return m }

// ListSysPostReplyValidationError is the validation error returned by
// ListSysPostReply.Validate if the designated constraints aren't met.
type ListSysPostReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSysPostReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSysPostReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSysPostReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSysPostReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSysPostReplyValidationError) ErrorName() string { return "ListSysPostReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListSysPostReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSysPostReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSysPostReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSysPostReplyValidationError{}

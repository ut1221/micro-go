// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/system/v1/sys_config.proto

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

// Validate checks the field values on CreateSysConfigRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSysConfigRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSysConfigRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSysConfigRepMultiError, or nil if none found.
func (m *CreateSysConfigRep) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSysConfigRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateSysConfigRepMultiError(errors)
	}

	return nil
}

// CreateSysConfigRepMultiError is an error wrapping multiple validation errors
// returned by CreateSysConfigRep.ValidateAll() if the designated constraints
// aren't met.
type CreateSysConfigRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSysConfigRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSysConfigRepMultiError) AllErrors() []error { return m }

// CreateSysConfigRepValidationError is the validation error returned by
// CreateSysConfigRep.Validate if the designated constraints aren't met.
type CreateSysConfigRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSysConfigRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSysConfigRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSysConfigRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSysConfigRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSysConfigRepValidationError) ErrorName() string {
	return "CreateSysConfigRepValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSysConfigRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSysConfigRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSysConfigRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSysConfigRepValidationError{}

// Validate checks the field values on EmptySysConfigReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EmptySysConfigReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmptySysConfigReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EmptySysConfigReplyMultiError, or nil if none found.
func (m *EmptySysConfigReply) ValidateAll() error {
	return m.validate(true)
}

func (m *EmptySysConfigReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptySysConfigReplyMultiError(errors)
	}

	return nil
}

// EmptySysConfigReplyMultiError is an error wrapping multiple validation
// errors returned by EmptySysConfigReply.ValidateAll() if the designated
// constraints aren't met.
type EmptySysConfigReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptySysConfigReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptySysConfigReplyMultiError) AllErrors() []error { return m }

// EmptySysConfigReplyValidationError is the validation error returned by
// EmptySysConfigReply.Validate if the designated constraints aren't met.
type EmptySysConfigReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptySysConfigReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptySysConfigReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptySysConfigReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptySysConfigReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptySysConfigReplyValidationError) ErrorName() string {
	return "EmptySysConfigReplyValidationError"
}

// Error satisfies the builtin error interface
func (e EmptySysConfigReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptySysConfigReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptySysConfigReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptySysConfigReplyValidationError{}

// Validate checks the field values on UpdateSysConfigRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSysConfigRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSysConfigRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSysConfigRepMultiError, or nil if none found.
func (m *UpdateSysConfigRep) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSysConfigRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateSysConfigRepMultiError(errors)
	}

	return nil
}

// UpdateSysConfigRepMultiError is an error wrapping multiple validation errors
// returned by UpdateSysConfigRep.ValidateAll() if the designated constraints
// aren't met.
type UpdateSysConfigRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSysConfigRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSysConfigRepMultiError) AllErrors() []error { return m }

// UpdateSysConfigRepValidationError is the validation error returned by
// UpdateSysConfigRep.Validate if the designated constraints aren't met.
type UpdateSysConfigRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSysConfigRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSysConfigRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSysConfigRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSysConfigRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSysConfigRepValidationError) ErrorName() string {
	return "UpdateSysConfigRepValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSysConfigRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSysConfigRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSysConfigRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSysConfigRepValidationError{}

// Validate checks the field values on DeleteSysConfigRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSysConfigRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSysConfigRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSysConfigRepMultiError, or nil if none found.
func (m *DeleteSysConfigRep) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSysConfigRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteSysConfigRepMultiError(errors)
	}

	return nil
}

// DeleteSysConfigRepMultiError is an error wrapping multiple validation errors
// returned by DeleteSysConfigRep.ValidateAll() if the designated constraints
// aren't met.
type DeleteSysConfigRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSysConfigRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSysConfigRepMultiError) AllErrors() []error { return m }

// DeleteSysConfigRepValidationError is the validation error returned by
// DeleteSysConfigRep.Validate if the designated constraints aren't met.
type DeleteSysConfigRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSysConfigRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSysConfigRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSysConfigRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSysConfigRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSysConfigRepValidationError) ErrorName() string {
	return "DeleteSysConfigRepValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSysConfigRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSysConfigRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSysConfigRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSysConfigRepValidationError{}

// Validate checks the field values on GetSysConfigRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetSysConfigRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSysConfigRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSysConfigRepMultiError, or nil if none found.
func (m *GetSysConfigRep) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSysConfigRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := GetSysConfigRepValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetSysConfigRepMultiError(errors)
	}

	return nil
}

// GetSysConfigRepMultiError is an error wrapping multiple validation errors
// returned by GetSysConfigRep.ValidateAll() if the designated constraints
// aren't met.
type GetSysConfigRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSysConfigRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSysConfigRepMultiError) AllErrors() []error { return m }

// GetSysConfigRepValidationError is the validation error returned by
// GetSysConfigRep.Validate if the designated constraints aren't met.
type GetSysConfigRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSysConfigRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSysConfigRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSysConfigRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSysConfigRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSysConfigRepValidationError) ErrorName() string { return "GetSysConfigRepValidationError" }

// Error satisfies the builtin error interface
func (e GetSysConfigRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSysConfigRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSysConfigRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSysConfigRepValidationError{}

// Validate checks the field values on GetSysConfigReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetSysConfigReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSysConfigReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSysConfigReplyMultiError, or nil if none found.
func (m *GetSysConfigReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSysConfigReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetSysConfigReplyValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetSysConfigReplyValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetSysConfigReplyValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetSysConfigReplyMultiError(errors)
	}

	return nil
}

// GetSysConfigReplyMultiError is an error wrapping multiple validation errors
// returned by GetSysConfigReply.ValidateAll() if the designated constraints
// aren't met.
type GetSysConfigReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSysConfigReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSysConfigReplyMultiError) AllErrors() []error { return m }

// GetSysConfigReplyValidationError is the validation error returned by
// GetSysConfigReply.Validate if the designated constraints aren't met.
type GetSysConfigReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSysConfigReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSysConfigReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSysConfigReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSysConfigReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSysConfigReplyValidationError) ErrorName() string {
	return "GetSysConfigReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetSysConfigReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSysConfigReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSysConfigReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSysConfigReplyValidationError{}

// Validate checks the field values on ListSysConfigRep with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListSysConfigRep) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSysConfigRep with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSysConfigRepMultiError, or nil if none found.
func (m *ListSysConfigRep) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSysConfigRep) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPageNum() <= 0 {
		err := ListSysConfigRepValidationError{
			field:  "PageNum",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPageSize() <= 0 {
		err := ListSysConfigRepValidationError{
			field:  "PageSize",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListSysConfigRepMultiError(errors)
	}

	return nil
}

// ListSysConfigRepMultiError is an error wrapping multiple validation errors
// returned by ListSysConfigRep.ValidateAll() if the designated constraints
// aren't met.
type ListSysConfigRepMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSysConfigRepMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSysConfigRepMultiError) AllErrors() []error { return m }

// ListSysConfigRepValidationError is the validation error returned by
// ListSysConfigRep.Validate if the designated constraints aren't met.
type ListSysConfigRepValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSysConfigRepValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSysConfigRepValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSysConfigRepValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSysConfigRepValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSysConfigRepValidationError) ErrorName() string { return "ListSysConfigRepValidationError" }

// Error satisfies the builtin error interface
func (e ListSysConfigRepValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSysConfigRep.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSysConfigRepValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSysConfigRepValidationError{}

// Validate checks the field values on ListSysConfigReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListSysConfigReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSysConfigReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSysConfigReplyMultiError, or nil if none found.
func (m *ListSysConfigReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSysConfigReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRows() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListSysConfigReplyValidationError{
						field:  fmt.Sprintf("Rows[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListSysConfigReplyValidationError{
						field:  fmt.Sprintf("Rows[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSysConfigReplyValidationError{
					field:  fmt.Sprintf("Rows[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListSysConfigReplyMultiError(errors)
	}

	return nil
}

// ListSysConfigReplyMultiError is an error wrapping multiple validation errors
// returned by ListSysConfigReply.ValidateAll() if the designated constraints
// aren't met.
type ListSysConfigReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSysConfigReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSysConfigReplyMultiError) AllErrors() []error { return m }

// ListSysConfigReplyValidationError is the validation error returned by
// ListSysConfigReply.Validate if the designated constraints aren't met.
type ListSysConfigReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSysConfigReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSysConfigReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSysConfigReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSysConfigReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSysConfigReplyValidationError) ErrorName() string {
	return "ListSysConfigReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListSysConfigReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSysConfigReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSysConfigReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSysConfigReplyValidationError{}

// Validate checks the field values on ConfigByKeyReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ConfigByKeyReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigByKeyReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConfigByKeyReqMultiError,
// or nil if none found.
func (m *ConfigByKeyReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigByKeyReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if len(errors) > 0 {
		return ConfigByKeyReqMultiError(errors)
	}

	return nil
}

// ConfigByKeyReqMultiError is an error wrapping multiple validation errors
// returned by ConfigByKeyReq.ValidateAll() if the designated constraints
// aren't met.
type ConfigByKeyReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigByKeyReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigByKeyReqMultiError) AllErrors() []error { return m }

// ConfigByKeyReqValidationError is the validation error returned by
// ConfigByKeyReq.Validate if the designated constraints aren't met.
type ConfigByKeyReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigByKeyReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigByKeyReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigByKeyReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigByKeyReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigByKeyReqValidationError) ErrorName() string { return "ConfigByKeyReqValidationError" }

// Error satisfies the builtin error interface
func (e ConfigByKeyReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigByKeyReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigByKeyReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigByKeyReqValidationError{}

// Validate checks the field values on ConfigByKeyReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConfigByKeyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigByKeyReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfigByKeyReplyMultiError, or nil if none found.
func (m *ConfigByKeyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigByKeyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return ConfigByKeyReplyMultiError(errors)
	}

	return nil
}

// ConfigByKeyReplyMultiError is an error wrapping multiple validation errors
// returned by ConfigByKeyReply.ValidateAll() if the designated constraints
// aren't met.
type ConfigByKeyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigByKeyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigByKeyReplyMultiError) AllErrors() []error { return m }

// ConfigByKeyReplyValidationError is the validation error returned by
// ConfigByKeyReply.Validate if the designated constraints aren't met.
type ConfigByKeyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigByKeyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigByKeyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigByKeyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigByKeyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigByKeyReplyValidationError) ErrorName() string { return "ConfigByKeyReplyValidationError" }

// Error satisfies the builtin error interface
func (e ConfigByKeyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigByKeyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigByKeyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigByKeyReplyValidationError{}

// Validate checks the field values on ConfigReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ConfigReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConfigReplyMultiError, or
// nil if none found.
func (m *ConfigReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ConfigId

	// no validation rules for ConfigName

	// no validation rules for ConfigKey

	// no validation rules for ConfigValue

	// no validation rules for ConfigType

	// no validation rules for Remark

	// no validation rules for CreatedAt

	if len(errors) > 0 {
		return ConfigReplyMultiError(errors)
	}

	return nil
}

// ConfigReplyMultiError is an error wrapping multiple validation errors
// returned by ConfigReply.ValidateAll() if the designated constraints aren't met.
type ConfigReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigReplyMultiError) AllErrors() []error { return m }

// ConfigReplyValidationError is the validation error returned by
// ConfigReply.Validate if the designated constraints aren't met.
type ConfigReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigReplyValidationError) ErrorName() string { return "ConfigReplyValidationError" }

// Error satisfies the builtin error interface
func (e ConfigReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigReplyValidationError{}

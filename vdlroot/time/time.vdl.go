// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: time

// Package time defines standard representations of absolute and relative times.
//
// The representations described below are required to provide wire
// compatibility between different programming environments.  Generated code for
// different environments typically provide automatic conversions into native
// representations, for simpler idiomatic usage.
package time

import (
	"fmt"
	"reflect"
	"time"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Duration represents the elapsed duration between two points in time, with
// up to nanosecond precision.
type Duration struct {
	// Seconds represents the seconds in the duration.  The range is roughly
	// +/-290 billion years, larger than the estimated age of the universe.
	Seconds int64
	// Nanos represents the fractions of a second at nanosecond resolution.  Must
	// be in the inclusive range between +/-999,999,999.
	//
	// In normalized form, durations less than one second are represented with 0
	// Seconds and +/-Nanos.  For durations one second or more, the sign of Nanos
	// must match Seconds, or be 0.
	Nanos int32
}

func (Duration) __VDLReflect(struct {
	Name string `vdl:"time.Duration"`
}) {
}

func (m *Duration) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Seconds")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromInt(int64(m.Seconds), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Nanos")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.Nanos), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Duration) MakeVDLTarget() vdl.Target {
	return nil
}

type DurationTarget struct {
	Value         *time.Duration
	wireValue     Duration
	secondsTarget vdl.Int64Target
	nanosTarget   vdl.Int32Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *DurationTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	t.wireValue = reflect.Zero(reflect.TypeOf(t.wireValue)).Interface().(Duration)
	if ttWant := vdl.TypeOf((*Duration)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *DurationTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Seconds":
		t.secondsTarget.Value = &t.wireValue.Seconds
		target, err := &t.secondsTarget, error(nil)
		return nil, target, err
	case "Nanos":
		t.nanosTarget.Value = &t.wireValue.Nanos
		target, err := &t.nanosTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct time.Duration", name)
	}
}
func (t *DurationTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *DurationTarget) FinishFields(_ vdl.FieldsTarget) error {

	if err := DurationToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

// Time represents an absolute point in time with up to nanosecond precision.
//
// Time is represented as the duration before or after a fixed epoch.  The zero
// Time represents the epoch 0001-01-01T00:00:00.000000000Z.  This uses the
// proleptic Gregorian calendar; the calendar runs on an exact 400 year cycle.
// Leap seconds are "smeared", ensuring that no leap second table is necessary
// for interpretation.
//
// This is similar to Go time.Time, but always in the UTC location.
// http://golang.org/pkg/time/#Time
//
// This is similar to conventional "unix time", but with the epoch defined at
// year 1 rather than year 1970.  This allows the zero Time to be used as a
// natural sentry, since it isn't a valid time for many practical applications.
// http://en.wikipedia.org/wiki/Unix_time
type Time struct {
	Seconds int64
	Nanos   int32
}

func (Time) __VDLReflect(struct {
	Name string `vdl:"time.Time"`
}) {
}

func (m *Time) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Seconds")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromInt(int64(m.Seconds), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Nanos")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.Nanos), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Time) MakeVDLTarget() vdl.Target {
	return nil
}

type TimeTarget struct {
	Value         *time.Time
	wireValue     Time
	secondsTarget vdl.Int64Target
	nanosTarget   vdl.Int32Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *TimeTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	t.wireValue = reflect.Zero(reflect.TypeOf(t.wireValue)).Interface().(Time)
	if ttWant := vdl.TypeOf((*Time)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *TimeTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Seconds":
		t.secondsTarget.Value = &t.wireValue.Seconds
		target, err := &t.secondsTarget, error(nil)
		return nil, target, err
	case "Nanos":
		t.nanosTarget.Value = &t.wireValue.Nanos
		target, err := &t.nanosTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct time.Time", name)
	}
}
func (t *TimeTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *TimeTarget) FinishFields(_ vdl.FieldsTarget) error {

	if err := TimeToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

// WireDeadline represents the deadline for an operation, where the operation is
// expected to finish before the deadline.  The intended usage is for a client
// to set a deadline on an operation, say one minute from "now", and send the
// deadline to a server.  The server is expected to finish the operation before
// the deadline.
//
// On a single device, it is simplest to represent a deadline as an absolute
// time; when the time now reaches the deadline, the deadline has expired.
// However when sending a deadline between devices with potential clock skew, it
// is often more robust to represent the deadline as a duration from "now".  The
// sender computes the duration from its notion of "now", while the receiver
// computes the absolute deadline from its own notion of "now".
//
// This representation doesn't account for propagation delay, but does ensure
// that the deadline used by the receiver is no earlier than the deadline
// intended by the client.  In many common scenarios the propagation delay is
// small compared to the potential clock skew, making this a simple but
// effective approach.
//
// WireDeadline typically has a native representation called Deadline that is an
// absolute Time, which automatically performs the sender and receiver
// conversions from "now".
type WireDeadline struct {
	// FromNow represents the deadline as a duration from "now".
	FromNow time.Duration
	// NoDeadline indicates there is no deadline; the analogous sentry for the
	// native Deadline is the zero Time.
	NoDeadline bool
}

func (WireDeadline) __VDLReflect(struct {
	Name string `vdl:"time.WireDeadline"`
}) {
}

func (m *WireDeadline) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var wireValue2 Duration
	if err := DurationFromNative(&wireValue2, m.FromNow); err != nil {
		return err
	}

	keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("FromNow")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue2.FillVDLTarget(fieldTarget4, tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
			return err
		}
	}
	keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("NoDeadline")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget6.FromBool(bool(m.NoDeadline), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WireDeadline) MakeVDLTarget() vdl.Target {
	return nil
}

type WireDeadlineTarget struct {
	Value            *Deadline
	wireValue        WireDeadline
	fromNowTarget    DurationTarget
	noDeadlineTarget vdl.BoolTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WireDeadlineTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	t.wireValue = reflect.Zero(reflect.TypeOf(t.wireValue)).Interface().(WireDeadline)
	if ttWant := vdl.TypeOf((*WireDeadline)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *WireDeadlineTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "FromNow":
		t.fromNowTarget.Value = &t.wireValue.FromNow
		target, err := &t.fromNowTarget, error(nil)
		return nil, target, err
	case "NoDeadline":
		t.noDeadlineTarget.Value = &t.wireValue.NoDeadline
		target, err := &t.noDeadlineTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct time.WireDeadline", name)
	}
}
func (t *WireDeadlineTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WireDeadlineTarget) FinishFields(_ vdl.FieldsTarget) error {

	if err := WireDeadlineToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

// Type-check Duration conversion functions.
var _ func(Duration, *time.Duration) error = DurationToNative
var _ func(*Duration, time.Duration) error = DurationFromNative

// Type-check Time conversion functions.
var _ func(Time, *time.Time) error = TimeToNative
var _ func(*Time, time.Time) error = TimeFromNative

// Type-check WireDeadline conversion functions.
var _ func(WireDeadline, *Deadline) error = WireDeadlineToNative
var _ func(*WireDeadline, Deadline) error = WireDeadlineFromNative

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}

	// Register native type conversions first, so that vdl.TypeOf works.
	vdl.RegisterNative(DurationToNative, DurationFromNative)
	vdl.RegisterNative(TimeToNative, TimeFromNative)
	vdl.RegisterNative(WireDeadlineToNative, WireDeadlineFromNative)

	// Register types.
	vdl.Register((*Duration)(nil))
	vdl.Register((*Time)(nil))
	vdl.Register((*WireDeadline)(nil))

	return struct{}{}
}

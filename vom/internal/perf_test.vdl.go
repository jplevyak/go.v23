// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: perf_test.vdl

package internal

import (
	// VDL system imports
	"fmt"
	"v.io/v23/vdl"
)

type AddressInfo struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (AddressInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom/internal.AddressInfo"`
}) {
}

type CreditAgency int

const (
	CreditAgencyEquifax CreditAgency = iota
	CreditAgencyExperian
	CreditAgencyTransUnion
)

// CreditAgencyAll holds all labels for CreditAgency.
var CreditAgencyAll = [...]CreditAgency{CreditAgencyEquifax, CreditAgencyExperian, CreditAgencyTransUnion}

// CreditAgencyFromString creates a CreditAgency from a string label.
func CreditAgencyFromString(label string) (x CreditAgency, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *CreditAgency) Set(label string) error {
	switch label {
	case "Equifax", "equifax":
		*x = CreditAgencyEquifax
		return nil
	case "Experian", "experian":
		*x = CreditAgencyExperian
		return nil
	case "TransUnion", "transunion":
		*x = CreditAgencyTransUnion
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in internal.CreditAgency", label)
}

// String returns the string label of x.
func (x CreditAgency) String() string {
	switch x {
	case CreditAgencyEquifax:
		return "Equifax"
	case CreditAgencyExperian:
		return "Experian"
	case CreditAgencyTransUnion:
		return "TransUnion"
	}
	return ""
}

func (CreditAgency) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom/internal.CreditAgency"`
	Enum struct{ Equifax, Experian, TransUnion string }
}) {
}

type ExperianRating int

const (
	ExperianRatingGood ExperianRating = iota
	ExperianRatingBad
)

// ExperianRatingAll holds all labels for ExperianRating.
var ExperianRatingAll = [...]ExperianRating{ExperianRatingGood, ExperianRatingBad}

// ExperianRatingFromString creates a ExperianRating from a string label.
func ExperianRatingFromString(label string) (x ExperianRating, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *ExperianRating) Set(label string) error {
	switch label {
	case "Good", "good":
		*x = ExperianRatingGood
		return nil
	case "Bad", "bad":
		*x = ExperianRatingBad
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in internal.ExperianRating", label)
}

// String returns the string label of x.
func (x ExperianRating) String() string {
	switch x {
	case ExperianRatingGood:
		return "Good"
	case ExperianRatingBad:
		return "Bad"
	}
	return ""
}

func (ExperianRating) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom/internal.ExperianRating"`
	Enum struct{ Good, Bad string }
}) {
}

type EquifaxCreditReport struct {
	Rating byte
}

func (EquifaxCreditReport) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom/internal.EquifaxCreditReport"`
}) {
}

type ExperianCreditReport struct {
	Rating ExperianRating
}

func (ExperianCreditReport) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom/internal.ExperianCreditReport"`
}) {
}

type TransUnionCreditReport struct {
	Rating int16
}

func (TransUnionCreditReport) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom/internal.TransUnionCreditReport"`
}) {
}

type (
	// AgencyReport represents any single field of the AgencyReport union type.
	AgencyReport interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the AgencyReport union type.
		__VDLReflect(__AgencyReportReflect)
	}
	// AgencyReportEquifaxReport represents field EquifaxReport of the AgencyReport union type.
	AgencyReportEquifaxReport struct{ Value EquifaxCreditReport }
	// AgencyReportExperianReport represents field ExperianReport of the AgencyReport union type.
	AgencyReportExperianReport struct{ Value ExperianCreditReport }
	// AgencyReportTransUnionReport represents field TransUnionReport of the AgencyReport union type.
	AgencyReportTransUnionReport struct{ Value TransUnionCreditReport }
	// __AgencyReportReflect describes the AgencyReport union type.
	__AgencyReportReflect struct {
		Name  string `vdl:"v.io/v23/vom/internal.AgencyReport"`
		Type  AgencyReport
		Union struct {
			EquifaxReport    AgencyReportEquifaxReport
			ExperianReport   AgencyReportExperianReport
			TransUnionReport AgencyReportTransUnionReport
		}
	}
)

func (x AgencyReportEquifaxReport) Index() int                         { return 0 }
func (x AgencyReportEquifaxReport) Interface() interface{}             { return x.Value }
func (x AgencyReportEquifaxReport) Name() string                       { return "EquifaxReport" }
func (x AgencyReportEquifaxReport) __VDLReflect(__AgencyReportReflect) {}

func (x AgencyReportExperianReport) Index() int                         { return 1 }
func (x AgencyReportExperianReport) Interface() interface{}             { return x.Value }
func (x AgencyReportExperianReport) Name() string                       { return "ExperianReport" }
func (x AgencyReportExperianReport) __VDLReflect(__AgencyReportReflect) {}

func (x AgencyReportTransUnionReport) Index() int                         { return 2 }
func (x AgencyReportTransUnionReport) Interface() interface{}             { return x.Value }
func (x AgencyReportTransUnionReport) Name() string                       { return "TransUnionReport" }
func (x AgencyReportTransUnionReport) __VDLReflect(__AgencyReportReflect) {}

type CreditReport struct {
	Agency CreditAgency
	Report AgencyReport
}

func (CreditReport) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom/internal.CreditReport"`
}) {
}

type Customer struct {
	Name    string
	Id      int64
	Active  bool
	Address AddressInfo
	Credit  CreditReport
}

func (Customer) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom/internal.Customer"`
}) {
}

func init() {
	vdl.Register((*AddressInfo)(nil))
	vdl.Register((*CreditAgency)(nil))
	vdl.Register((*ExperianRating)(nil))
	vdl.Register((*EquifaxCreditReport)(nil))
	vdl.Register((*ExperianCreditReport)(nil))
	vdl.Register((*TransUnionCreditReport)(nil))
	vdl.Register((*AgencyReport)(nil))
	vdl.Register((*CreditReport)(nil))
	vdl.Register((*Customer)(nil))
}
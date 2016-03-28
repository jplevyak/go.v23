// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syncbase

import (
	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/security/access"
	wire "v.io/v23/services/syncbase"
	"v.io/v23/syncbase/util"
)

func newApp(parentFullName, relativeName string) App {
	fullName := naming.Join(parentFullName, util.Escape(relativeName))
	return &app{
		c:        wire.AppClient(fullName),
		fullName: fullName,
		name:     relativeName,
	}
}

type app struct {
	c        wire.AppClientMethods
	fullName string
	name     string
}

var _ App = (*app)(nil)

// Name implements App.Name.
func (a *app) Name() string {
	return a.name
}

// FullName implements App.FullName.
func (a *app) FullName() string {
	return a.fullName
}

// Exists implements App.Exists.
func (a *app) Exists(ctx *context.T) (bool, error) {
	return a.c.Exists(ctx)
}

// Database implements App.Database.
func (a *app) Database(relativeName string, schema *Schema) Database {
	return NewDatabase(a.fullName, relativeName, schema)
}

// ListDatabases implements App.ListDatabases.
func (a *app) ListDatabases(ctx *context.T) ([]string, error) {
	return util.ListChildren(ctx, a.fullName)
}

// Create implements App.Create.
func (a *app) Create(ctx *context.T, perms access.Permissions) error {
	return a.c.Create(ctx, perms)
}

// Destroy implements App.Destroy.
func (a *app) Destroy(ctx *context.T) error {
	return a.c.Destroy(ctx)
}

// SetPermissions implements App.SetPermissions.
func (a *app) SetPermissions(ctx *context.T, perms access.Permissions, version string) error {
	return a.c.SetPermissions(ctx, perms, version)
}

// GetPermissions implements App.GetPermissions.
func (a *app) GetPermissions(ctx *context.T) (perms access.Permissions, version string, err error) {
	return a.c.GetPermissions(ctx)
}

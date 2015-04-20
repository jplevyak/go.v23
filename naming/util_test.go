// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package naming_test

import (
	"testing"

	"v.io/v23/naming"
	"v.io/v23/rpc/version"
)

func TestFormat(t *testing.T) {
	testcases := []struct {
		network, address string
		opts             []naming.EndpointOpt
		output           string
	}{
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{}, "@5@tcp@127.0.0.1:21@@@@@"},
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{naming.NullRoutingID}, "@5@tcp@127.0.0.1:21@00000000000000000000000000000000@@@@"},
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{version.RPCVersionRange{3, 5}}, "@5@tcp@127.0.0.1:21@@@@@"},
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{naming.ServesMountTable(true)}, "@5@tcp@127.0.0.1:21@@m@@@"},
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{naming.ServesMountTable(false)}, "@5@tcp@127.0.0.1:21@@s@@@"},
		{"tcp", "127.0.0.1:22", []naming.EndpointOpt{naming.BlessingOpt("batman@dccomics.com")}, "@5@tcp@127.0.0.1:22@@@batman@dccomics.com@@"},
		{"tcp", "127.0.0.1:22", []naming.EndpointOpt{naming.BlessingOpt("batman@dccomics.com"), naming.BlessingOpt("bugs@bunny.com"), naming.ServesMountTable(true)}, "@5@tcp@127.0.0.1:22@@m@batman@dccomics.com,bugs@bunny.com@@"},
		{"tcp", "127.0.0.1:22", []naming.EndpointOpt{naming.BlessingOpt("batman@dccomics.com"), naming.BlessingOpt("bugs@bunny.com"), version.RPCVersionRange{5, 7}}, "@5@tcp@127.0.0.1:22@@@batman@dccomics.com,bugs@bunny.com@@"},
		{"tcp", "127.0.0.1:22", []naming.EndpointOpt{naming.BlessingOpt("@s@@")}, "@5@tcp@127.0.0.1:22@@@@s@@@@"},
		{"tcp", "127.0.0.1:22", []naming.EndpointOpt{naming.BlessingOpt("dev.v.io/services/mounttabled")}, "@5@tcp@127.0.0.1:22@@@dev.v.io/services/mounttabled@@"},
	}
	for i, test := range testcases {
		str := naming.FormatEndpoint(test.network, test.address, test.opts...)
		if str != test.output {
			t.Errorf("%d: unexpected endpoint string: got %q != %q", i, str, test.output)
		}
	}
}

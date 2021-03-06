// SPDX-License-Identifier: ISC
// Copyright (c) 2019-2021 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ACLTestSuite struct {
	suite.Suite
}

func (suite *ACLTestSuite) TestHasCommandAccess() {
	access := map[string]map[AccessMode]bool{
		"bind":                {AccessModeFull: true, AccessModeLimited: true, AccessModeMinimal: true},
		"bind_ack":            {AccessModeFull: true, AccessModeLimited: true, AccessModeMinimal: true},
		"bitcoind":            {AccessModeFull: true, AccessModeLimited: true, AccessModeMinimal: true},
		"get_bitcoind_status": {AccessModeFull: true, AccessModeLimited: true, AccessModeMinimal: true},
		"create_wallet":       {AccessModeFull: true},
		"finish_psbt":         {AccessModeFull: true},
		"set_member":          {AccessModeFull: true},
		"remove_member":       {AccessModeFull: true},
		"start_bitcoind":      {AccessModeFull: true},
		"stop_bitcoind":       {AccessModeFull: true},
	}
	for command, access := range access {
		for _, mode := range []AccessMode{AccessModeNotApplicant, AccessModeFull, AccessModeLimited, AccessModeMinimal} {
			suite.Equal(access[mode], HasCommandAccess(command, mode))
		}
	}
}

func (suite *ACLTestSuite) TestHasBitcoinRPCAccess() {
	access := map[AccessMode]map[string]bool{
		AccessModeFull: {
			"getbalances":            true,
			"getblockchaininfo":      true,
			"getmininginfo":          true,
			"getnettotals":           true,
			"getnetworkinfo":         true,
			"getnewaddress":          true,
			"getreceivedbyaddress":   true,
			"gettransaction":         true,
			"getwalletinfo":          true,
			"listtransactions":       true,
			"walletcreatefundedpsbt": true,
		},
		AccessModeLimited: {
			"getbalances":            false,
			"getblockchaininfo":      false,
			"getmininginfo":          false,
			"getnettotals":           false,
			"getnetworkinfo":         false,
			"getnewaddress":          false,
			"getreceivedbyaddress":   false,
			"gettransaction":         false,
			"getwalletinfo":          false,
			"listtransactions":       false,
			"walletcreatefundedpsbt": false,
		},
		AccessModeMinimal: {
			"getbalances":            false,
			"getblockchaininfo":      false,
			"getmininginfo":          false,
			"getnettotals":           false,
			"getnetworkinfo":         false,
			"getnewaddress":          false,
			"getreceivedbyaddress":   false,
			"gettransaction":         false,
			"getwalletinfo":          false,
			"listtransactions":       false,
			"walletcreatefundedpsbt": false,
		},
		AccessModeNotApplicant: {
			"getbalances":            false,
			"getblockchaininfo":      false,
			"getmininginfo":          false,
			"getnettotals":           false,
			"getnetworkinfo":         false,
			"getnewaddress":          false,
			"getreceivedbyaddress":   false,
			"gettransaction":         false,
			"getwalletinfo":          false,
			"listtransactions":       false,
			"walletcreatefundedpsbt": false,
		},
	}
	for mode, access := range access {
		for rpc, allowed := range access {
			suite.Equal(allowed, HasBitcoinRPCAccess(rpc, mode))
		}
	}
}

func TestACLTestSuite(t *testing.T) {
	suite.Run(t, &ACLTestSuite{})
}

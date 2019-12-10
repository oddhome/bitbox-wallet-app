// Copyright 2018 Shift Devices AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package accounts

import (
	"github.com/btcsuite/btcd/wire"
	"github.com/digitalbitbox/bitbox-wallet-app/backend/coins/coin"
	"github.com/digitalbitbox/bitbox-wallet-app/backend/keystore"
	"github.com/digitalbitbox/bitbox-wallet-app/backend/rates"
	"github.com/digitalbitbox/bitbox-wallet-app/backend/signing"
)

// Interface is the API of a Account.
type Interface interface {
	Info() *Info
	// Code is an identifier for the account *type* (part of account database filenames, apis, etc.).
	// Type as in btc-p2wpkh, eth-erc20-usdt, etc.
	Code() string
	// FilesFolder is path to a directory for account files, like databases, etc. Only available
	// after Initialize(). It must be unique not only up to the type, but also the exact
	// keystores/signing configuration (e.g. a btc-p2wpkh account for one xpub/xprv should have a
	// different ID).
	FilesFolder() string
	Coin() coin.Coin
	// Name returns a human readable long name.
	Name() string
	// Initialize only starts the initialization, the account is not initialized right afterwards.
	Initialize() error
	Initialized() bool
	Offline() bool
	FatalError() bool
	Close()
	Notifier() Notifier
	Transactions() ([]Transaction, error)
	Balance() (*Balance, error)
	// Creates, signs and broadcasts a transaction. Returns keystore.ErrSigningAborted on user
	// abort.
	SendTx(string, coin.SendAmount, FeeTargetCode, map[wire.OutPoint]struct{}, []byte) error
	FeeTargets() ([]FeeTarget, FeeTargetCode)
	TxProposal(string, coin.SendAmount, FeeTargetCode, map[wire.OutPoint]struct{}, []byte) (
		coin.Amount, coin.Amount, coin.Amount, error)
	GetUnusedReceiveAddresses() []Address
	CanVerifyAddresses() (bool, bool, error)
	VerifyAddress(addressID string) (bool, error)
	Keystores() *keystore.Keystores
	RateUpdater() *rates.RateUpdater
}

// Info holds account information.
type Info struct {
	SigningConfiguration *signing.Configuration `json:"signingConfiguration"`
}

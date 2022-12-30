// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate mockery --name System --filename system.go

package system

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type System interface {
	Properties() (types.ChainProperties, error)
	Health() (types.Health, error)
	Peers() ([]types.PeerInfo, error)
	Name() (types.Text, error)
	Chain() (types.Text, error)
	Version() (types.Text, error)
	NetworkState() (types.NetworkState, error)
}

// system exposes methods for retrieval of system data
type system struct {
	client client.Client
}

// NewSystem creates a new system struct
func NewSystem(cl client.Client) System {
	return &system{cl}
}

// Chain retrieves the chain
func (c *system) Chain() (types.Text, error) {
	var t types.Text
	err := c.client.Call(&t, "system_chain")
	return t, err
}

// Health retrieves the health status of the connected node
func (c *system) Health() (types.Health, error) {
	var h types.Health
	err := c.client.Call(&h, "system_health")
	return h, err
}

// Name retrieves the node name
func (c *system) Name() (types.Text, error) {
	var t types.Text
	err := c.client.Call(&t, "system_name")
	return t, err
}

// NetworkState retrieves the current state of the network
func (c *system) NetworkState() (types.NetworkState, error) {
	var n types.NetworkState
	err := c.client.Call(&n, "system_networkState")
	return n, err
}

// Peers retrieves the currently connected peers
func (c *system) Peers() ([]types.PeerInfo, error) {
	var p []types.PeerInfo
	err := c.client.Call(&p, "system_peers")
	return p, err
}

// Properties retrieves a custom set of properties as a JSON object, defined in the chain spec
func (c *system) Properties() (types.ChainProperties, error) {
	var p types.ChainProperties
	err := c.client.Call(&p, "system_properties")
	return p, err
}

// Version retrieves the version of the node
func (c *system) Version() (types.Text, error) {
	var t types.Text
	err := c.client.Call(&t, "system_version")
	return t, err
}

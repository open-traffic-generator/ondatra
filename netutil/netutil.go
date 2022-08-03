// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package netutil provides network-related helper functions for testing.
package netutil

import (
	"math/big"
)

// DefaultNetworkInstance is the constant name for referring to the default network
// instance on the device.  It has been standardized in OpenConfig.
const DefaultNetworkInstance = "DEFAULT"

var one = big.NewInt(1)

func maxVal(numBytes uint) *big.Int {
	limit := new(big.Int).Lsh(one, 8*numBytes)
	return limit.Sub(limit, one)
}

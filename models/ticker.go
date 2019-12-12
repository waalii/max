// Copyright 2018 MaiCoin Technologies
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

package models

import (
	"time"

	"github.com/waalii/max/types"
)

type Ticker struct {
	// timestamp in seconds since Unix epoch
	At     time.Time    `json:"at,omitempty"`
	Buy    types.Price  `json:"buy,omitempty"`
	Sell   types.Price  `json:"sell,omitempty"`
	Open   types.Price  `json:"open,omitempty"`
	Last   types.Price  `json:"last,omitempty"`
	High   types.Price  `json:"high,omitempty"`
	Low    types.Price  `json:"low,omitempty"`
	Volume types.Volume `json:"vol,omitempty"`
}

type Tickers map[string]*Ticker

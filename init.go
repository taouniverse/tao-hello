// Copyright 2021
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tao_hello

import (
	"encoding/json"
	"github.com/taouniverse/tao"
)

/**
import _ "github.com/taouniverse/tao_hello"
*/

var h = new(HelloConfig)

func init() {
	err := tao.Register(ConfigKey, func() error {
		// 1. transfer config bytes to object
		bytes, err := tao.GetConfigBytes(ConfigKey)
		if err != nil {
			h = h.Default().(*HelloConfig)
		} else {
			err = json.Unmarshal(bytes, &h)
			if err != nil {
				return err
			}
		}
		// 2. set object to tao
		return tao.SetConfig(ConfigKey, h)
	})
	if err != nil {
		panic(err)
	}
}

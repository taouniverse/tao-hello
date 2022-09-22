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

package hello

import (
	"context"
	"fmt"
	"github.com/taouniverse/tao"
)

// ConfigKey for this repo
const ConfigKey = "hello"

// Config implements tao.Config
type Config struct {
	Print     string   `json:"print"`
	Times     int      `json:"times"`
	RunAfters []string `json:"run_after,omitempty"`
}

var defaultHello = &Config{
	Print: `
  _   _  U _____ u  _       _       U  ___ u      _____      _      U  ___ u 
 |'| |'| \| ___"|/ |"|     |"|       \/"_ \/     |_ " _| U  /"\  u   \/"_ \/ 
/| |_| |\ |  _|" U | | u U | | u     | | | |       | |    \/ _ \/    | | | | 
U|  _  |u | |___  \| |/__ \| |/__.-,_| |_| |      /| |\   / ___ \.-,_| |_| | 
 |_| |_|  |_____|  |_____| |_____|\_)-\___/      u |_|U  /_/   \_\\_)-\___/  
 //   \\  <<   >>  //  \\  //  \\      \\        _// \\_  \\    >>     \\    
(_") ("_)(__) (__)(_")("_)(_")("_)    (__)      (__) (__)(__)  (__)   (__)   
`,
	Times:     1,
	RunAfters: []string{},
}

// Name of Config
func (h *Config) Name() string {
	return ConfigKey
}

// ValidSelf with some default values
func (h *Config) ValidSelf() {
	if h.Times == 0 {
		h.Times = defaultHello.Times
	}
	if h.Print == "" {
		h.Print = defaultHello.Print
	}
	if h.RunAfters == nil {
		h.RunAfters = defaultHello.RunAfters
	}
}

// ToTask transform itself to Task
func (h *Config) ToTask() tao.Task {
	return tao.NewTask(
		ConfigKey,
		func(ctx context.Context, param tao.Parameter) (tao.Parameter, error) {
			// non-block check
			select {
			case <-ctx.Done():
				return param, tao.NewError(tao.ContextCanceled, "%s: context has been canceled", ConfigKey)
			default:
			}
			// print times
			for i := 0; i < h.Times; i++ {
				fmt.Println(h.Print)
			}
			return param, nil
		})
}

// RunAfter defines pre task names
func (h *Config) RunAfter() []string {
	return h.RunAfters
}

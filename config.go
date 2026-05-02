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

// InstanceConfig 单实例配置
type InstanceConfig struct {
	Print string `json:"print"`
	Times int    `json:"times"`
}

// Config 总配置，实现 tao.MultiConfig 接口
type Config struct {
	tao.BaseMultiConfig[InstanceConfig]
	RunAfters []string `json:"run_after,omitempty" yaml:"run_after,omitempty"`
}

var defaultInstance = &InstanceConfig{
	Print: `
  _   _  U _____ u  _       _       U  ___ u      _____      _      U  ___ u 
 |'| |'| \| ___"|/ |"|     |"|       \/"_ \/     |_ " _| U  /"\  u   \/"_ \/ 
/| |_| |\ |  _|" U | | u U | | u     | | | |       | |    \/ _ \/    | | | | 
U|  _  |u | |___  \| |/__ \| |/__.-,_| |_| |      /| |\   / ___ \.-,_| |_| | 
 |_| |_|  |_____|  |_____| |_____|\_)-\___/      u |_|U  /_/   \_\\_)-\___/  
 //   \\  <<   >>  //  \\  //  \\      \\        _// \\_  \\    >>     \\    
(_") ("_)(__) (__)(_")("_)(_")("_)    (__)      (__) (__)(__)  (__)   (__)   
`,
	Times: 1,
}

// Name of Config
func (h *Config) Name() string {
	return ConfigKey
}

// ValidSelf with some default values
func (h *Config) ValidSelf() {
	for i := range h.Instances {
		inst := &h.Instances[i].Cfg
		if inst.Times == 0 {
			inst.Times = defaultInstance.Times
		}
		if inst.Print == "" {
			inst.Print = defaultInstance.Print
		}
	}
	if h.RunAfters == nil {
		h.RunAfters = []string{}
	}
}

// ToTask transform itself to Task
func (h *Config) ToTask() tao.Task {
	return tao.NewTask(
		ConfigKey,
		func(ctx context.Context, param tao.Parameter) (tao.Parameter, error) {
			select {
			case <-ctx.Done():
				return param, tao.NewError(tao.ContextCanceled, "%s: context has been canceled", ConfigKey)
			default:
			}
			for _, inst := range h.Instances {
				for i := 0; i < inst.Cfg.Times; i++ {
					fmt.Println(inst.Cfg.Print)
				}
			}
			return param, nil
		})
}

// RunAfter defines pre task names
func (h *Config) RunAfter() []string {
	return h.RunAfters
}

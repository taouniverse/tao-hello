package tao_hello

import (
	"context"
	"fmt"
	"github.com/taouniverse/tao"
)

// ConfigKey for this repo
const ConfigKey = "hello"

// HelloConfig implements tao.Config
type HelloConfig struct {
	Print     string   `json:"print"`
	Times     int      `json:"times"`
	RunAfter_ []string `json:"run_after,omitempty"`
}

var defaultHello = &HelloConfig{
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
	RunAfter_: []string{},
}

// Default config
func (h *HelloConfig) Default() tao.Config {
	return defaultHello
}

// ValidSelf with some default values
func (h *HelloConfig) ValidSelf() {
	if h.Times == 0 {
		h.Times = defaultHello.Times
	}
	if h.Print == "" {
		h.Print = defaultHello.Print
	}
	if h.RunAfter_ == nil {
		h.RunAfter_ = defaultHello.RunAfter_
	}
}

// ToTask transform itself to Task
func (h *HelloConfig) ToTask() tao.Task {
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
func (h *HelloConfig) RunAfter() []string {
	return h.RunAfter_
}

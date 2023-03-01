//go:build !windows

package win_eventlog

import (
	_ "embed"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

//go:embed sample.conf
var sampleConfig string

type WinEventLog struct {
	Log telegraf.Logger `toml:"-"`
}

func (w *WinEventLog) Init() error {
	w.Log.Warn("current platform is not supported")
	return nil
}
func (w *WinEventLog) SampleConfig() string                { return sampleConfig }
func (w *WinEventLog) Gather(_ telegraf.Accumulator) error { return nil }

func init() {
	inputs.Add("win_eventlog", func() telegraf.Input {
		return &WinEventLog{}
	})
}

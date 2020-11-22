package tracer

import (
	"github.com/k0kubun/pp"
	"github.com/c3sr/config"
	"github.com/c3sr/vipertags"
)

type tracerConfig struct {
	Enabled     bool          `json:"enabled" config:"tracer.enabled" default:"true"`
	Provider    string        `json:"provider" config:"tracer.provider" default:"jaeger"`
	LevelString string        `json:"level" config:"tracer.level"`
	Level       Level         `json:"-" config:"-"`
	done        chan struct{} `json:"-" config:"-"`
}

var (
	// Config holds the data read by c3sr/config
	Config = &tracerConfig{
		done:  make(chan struct{}),
		Level: NO_TRACE,
	}
)

func (tracerConfig) ConfigName() string {
	return "Tracer"
}

func (c *tracerConfig) SetDefaults() {
	vipertags.SetDefaults(c)
}

func (c *tracerConfig) Read() {
	defer close(c.done)
	vipertags.Fill(c)
	c.Level = LevelFromName(c.LevelString)
}

func (c tracerConfig) Wait() {
	<-c.done
}

func (c tracerConfig) String() string {
	return pp.Sprintln(c)
}

func (c tracerConfig) Debug() {
	log.Debug("Tracer Config = ", c)
}

func init() {
	config.Register(Config)
}

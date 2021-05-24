package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

type EdgeDNSOptions struct {
	ListenIfName string
	ListenPort   int32
}

func NewEdgeDNSOptions() *EdgeDNSOptions {
	o := &EdgeDNSOptions{
		ListenIfName: "docker0",
		ListenPort:   53,
	}

	return o
}

// Validate validates NewEdgeDNSOptions
func Validate(options *EdgeDNSOptions) error {
	if len(options.ListenIfName) == 0 {
		return fmt.Errorf("listen-if-name is empty")
	}
	if 0 > options.ListenPort || options.ListenPort > 65535 {
		return fmt.Errorf("listen-port is invalid")
	}
	return nil
}

func (o *EdgeDNSOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.ListenIfName, "listen-if-name", o.ListenIfName, "The name of interface that edgedns listens to")
	fs.Int32Var(&o.ListenPort, "listen-port", o.ListenPort, "The port that edgedns listens to")
}

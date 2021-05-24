package config

import (
	"net"

	"github.com/kubeedge/edgemesh/cmd/edgedns/app/options"
	"github.com/kubeedge/edgemesh/pkg/networking/util"
)

type EdgeDNSConfiguration struct {
	ListenIfIP net.IP
	ListenPort int32
}

// Complete converts *options.EdgeDNSOptions to *EdgeDNSConfiguration
func Complete(options *options.EdgeDNSOptions) (*EdgeDNSConfiguration, error) {
	ip, err := util.GetInterfaceIP(options.ListenIfName)
	if err != nil {
		return nil, err
	}

	cfg := &EdgeDNSConfiguration{
		ListenIfIP: ip,
		ListenPort: options.ListenPort,
	}

	return cfg, nil
}

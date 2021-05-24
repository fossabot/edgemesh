package app

import (
	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/version/verflag"
	"k8s.io/klog/v2"

	"github.com/kubeedge/edgemesh/cmd/edgedns/app/config"
	"github.com/kubeedge/edgemesh/cmd/edgedns/app/options"
)

func NewEdgeDNSCommand() *cobra.Command {
	opts := options.NewEdgeDNSOptions()

	cmd := &cobra.Command{
		Use: "edgedns",
		Long: `EdgeDNS is the dns server module of EdgeMesh.`,
		Run: func(cmd *cobra.Command, args []string) {
			verflag.PrintAndExitIfRequested()
			cliflag.PrintFlags(cmd.Flags())

			if err := options.Validate(opts); err != nil {
				klog.Fatalf("failed validate: %v", err)
			}

			cfg, err := config.Complete(opts)
			if err != nil {
				klog.Fatalf("failed complete: %v", err)
			}

			if err := Run(cfg); err != nil {
				klog.Exit(err)
			}
		},
	}

	opts.AddFlags(cmd.Flags())
	return cmd
}

func Run(cfg *config.EdgeDNSConfiguration) error {
	klog.Infof("edge dns config: %v", cfg)
	return nil
}

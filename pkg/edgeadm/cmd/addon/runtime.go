package addon

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/superedge/superedge/pkg/edgeadm/constant"
	"github.com/superedge/superedge/pkg/edgeadm/steps"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
)

func NewInstallContainerRuntime() *cobra.Command {
	initRunner := workflow.NewRunner()

	cmd := &cobra.Command{
		Use:   "install-runtime",
		Short: "Install container runtime",
		Long:  "Run this on any machine you wish to install an container runtime",
		RunE: func(cmd *cobra.Command, args []string) error {

			var runtimeType string
			flags := cmd.Flags()
			flags.StringVar(&runtimeType, constant.RuntimeType, "docker", "container runtime type")
			if runtimeType == "" {
				return errors.New("you must specify the container runtime type")
			}

			if runtimeType == "docker" {
				initRunner.AppendPhase(steps.NewDockerPhase())
			} else if runtimeType == "containerd" {
				initRunner.AppendPhase(steps.NewContainerDPhase())
			} else {
				return errors.New("container runtime type error")
			}

			if err := initRunner.Run(args); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

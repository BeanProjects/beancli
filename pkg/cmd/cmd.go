package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	Version      string
	GitCommit    string
	GitTreeState string
	BuildDate    string
)

type BeanOptions struct {
	Arguments   []string
	ConfigFlags *genericclioptions.ConfigFlags

	genericiooptions.IOStreams
}

func NewDefaultBeanCommand() *cobra.Command {

	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "bean",
		Short: i18n.T("kubectl controls the Kubernetes cluster manager"),
		Long: templates.LongDesc(`
		  The bean command controls the whole ecosystem of BeanProjects.
		`),
	}

	// version cmd
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Bean",
		Long:  "All software has versions. This is Bean's",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("YourApp Version: %s\n", Version)
			fmt.Printf("Git Commit: %s\n", GitCommit)
			fmt.Printf("Git Tree State: %s\n", GitTreeState)
			fmt.Printf("Build Date: %s\n", BuildDate)
			fmt.Printf("Go OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		},
	}

	cmds.AddCommand(versionCmd)

	return cmds
}

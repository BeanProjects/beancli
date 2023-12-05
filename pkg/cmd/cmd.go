package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
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

	return cmds
}

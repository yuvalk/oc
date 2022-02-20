package catalog

import (
	"fmt"

	kcmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	listLong = templates.LongDesc(`
		Lists the content
	`)
	listExample = templates.Examples(`
		# list catalog
		oc adm catalog list quay.io/smthg
	`)
)

func init() {
	subCommands = append(subCommands, NewListCatalog)
}

func NewListCatalog(f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := NewListCatalogOptions(streams)

	cmd := &cobra.Command{
		Use:     "list CATALOG",
		Short:   "List catalog content",
		Long:    listLong,
		Example: listExample,
		Run: func(cmd *cobra.Command, args []string) {
			kcmdutil.CheckErr(o.Complete(cmd, args))
			kcmdutil.CheckErr(o.Validate())
			kcmdutil.CheckErr(o.Run())
		},
	}
	//flags := cmd.Flags()

	//o.SecurityOptions.Bind(flags)
	//o.ParallelOptions.Bind(flags)

	return cmd
}

type ListCatalogOptions struct {
}

func NewListCatalogOptions(streams genericclioptions.IOStreams) *ListCatalogOptions {
	return &ListCatalogOptions{}
}

func (o *ListCatalogOptions) Complete(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("must specify catalog")
	}

	catalog := args[0]

	fmt.Println("vim-go" + catalog)

	return nil
}

func (o *ListCatalogOptions) Validate() error {
	return nil
}

func (o *ListCatalogOptions) Run() error {
	return nil
}

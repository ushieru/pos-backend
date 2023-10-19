package cobra_app

import (
	"github.com/spf13/cobra"
	"github.com/ushieru/pos/app/cobra/cmd"
)

func NewCobraApp() *CobraApp {
	cobraApp := &CobraApp{
		root: cmd.RootCmd,
	}
	return cobraApp
}

func (c *CobraApp) Run() error {
	c.root.AddCommand(cmd.ServeCmd)
	if err := c.root.Execute(); err != nil {
		return err
	}
	return nil
}

type CobraApp struct {
	root *cobra.Command
}

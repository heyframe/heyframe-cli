package account

import (
	"github.com/spf13/cobra"
)

var accountCompanyProducerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Manage your HeyFrame manufacturer",
}

func init() {
	accountRootCmd.AddCommand(accountCompanyProducerCmd)
}

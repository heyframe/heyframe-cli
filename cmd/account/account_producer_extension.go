package account

import (
	"github.com/spf13/cobra"
)

var accountCompanyProducerExtensionCmd = &cobra.Command{
	Use:   "extension",
	Short: "Manage your HeyFrame extensions",
}

func init() {
	accountCompanyProducerCmd.AddCommand(accountCompanyProducerExtensionCmd)
}

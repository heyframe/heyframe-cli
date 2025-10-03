package account

import (
	"github.com/spf13/cobra"

	account_api "github.com/heyframe/heyframe-cli/internal/account-api"
	"github.com/heyframe/heyframe-cli/internal/config"
)

var accountRootCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage your HeyFrame Account",
}

type ServiceContainer struct {
	Conf          config.Config
	AccountClient *account_api.Client
}

var services *ServiceContainer

func Register(rootCmd *cobra.Command, onInit func(commandName string) (*ServiceContainer, error)) {
	accountRootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		ser, err := onInit(cmd.Name())
		services = ser
		return err
	}
	rootCmd.AddCommand(accountRootCmd)
}

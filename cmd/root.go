package cmd

import (
	"context"
	"os"
	"slices"

	"github.com/spf13/cobra"

	"github.com/heyframe/heyframe-cli/cmd/account"
	"github.com/heyframe/heyframe-cli/cmd/extension"
	"github.com/heyframe/heyframe-cli/cmd/project"
	accountApi "github.com/heyframe/heyframe-cli/internal/account-api"
	"github.com/heyframe/heyframe-cli/internal/config"
	"github.com/heyframe/heyframe-cli/internal/system"
	"github.com/heyframe/heyframe-cli/logging"
)

var (
	cfgFile string
	version = "dev"
)

var rootCmd = &cobra.Command{
	Use:     "heyFrame-cli",
	Short:   "A cli for common HeyFrame tasks",
	Long:    `This application contains some utilities like extension management`,
	Version: version,
}

func Execute(ctx context.Context) {
	ctx = logging.WithLogger(ctx, logging.NewLogger(slices.Contains(os.Args, "--verbose")))
	accountApi.SetUserAgent("heyFrame-cli/" + version)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		logging.FromContext(ctx).Fatalln(err)
	}
}

func init() {
	rootCmd.SilenceErrors = true

	cobra.OnInitialize(func() {
		_ = config.InitConfig(cfgFile)
	})

	cobra.OnFinalize(func() {
		_ = system.CloseCaches()
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.heyFrame-cli.yaml)")
	rootCmd.PersistentFlags().Bool("verbose", false, "show debug output")

	project.Register(rootCmd)
	extension.Register(rootCmd)
	account.Register(rootCmd, func(commandName string) (*account.ServiceContainer, error) {
		err := config.InitConfig(cfgFile)
		if err != nil {
			return nil, err
		}
		conf := config.Config{}
		if commandName == "login" || commandName == "logout" {
			return &account.ServiceContainer{
				Conf:          conf,
				AccountClient: nil,
			}, nil
		}
		client, err := accountApi.NewApi(rootCmd.Context(), conf)
		if err != nil {
			return nil, err
		}
		return &account.ServiceContainer{
			Conf:          conf,
			AccountClient: client,
		}, nil
	})
}

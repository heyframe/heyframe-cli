package project

import (
	"os"

	adminSdk "github.com/heyframe/go-heyframe-admin-api-sdk"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/heyframe/heyframe-cli/logging"
	"github.com/heyframe/heyframe-cli/platform"
)

var projectConfigPullCmd = &cobra.Command{
	Use:   "pull",
	Short: "[Deprecated] Synchronizes your shop config to local",
	RunE: func(cmd *cobra.Command, _ []string) error {
		logging.FromContext(cmd.Context()).Warnf("This command is deprecated and will be removed in the future. Please use Fixture Bundle instead https://developer.heyFrame.com/docs/resources/tooling/fixture-bundle/")

		var cfg *platform.Config
		var err error

		if cfg, err = platform.ReadConfig(projectConfigPath, false); err != nil {
			return err
		}

		client, err := platform.NewPlatformClient(cmd.Context(), cfg)
		if err != nil {
			return err
		}

		for _, applyer := range NewSyncApplyers(cfg) {
			if err := applyer.Pull(adminSdk.NewApiContext(cmd.Context()), client, cfg); err != nil {
				return err
			}
		}

		content, err := yaml.Marshal(cfg)
		if err != nil {
			return err
		}

		if err := os.WriteFile(projectConfigPath, content, os.ModePerm); err != nil {
			return err
		}

		logging.FromContext(cmd.Context()).Infof("%s has been updated", projectConfigPath)

		return nil
	},
}

func init() {
	projectConfigCmd.AddCommand(projectConfigPullCmd)
}

package project

import (
	"fmt"
	"os"

	adminSdk "github.com/heyframe/go-heyframe-admin-api-sdk"
	"github.com/spf13/cobra"

	"github.com/heyframe/heyframe-cli/logging"
	"github.com/heyframe/heyframe-cli/platform"
)

var projectClearCacheCmd = &cobra.Command{
	Use:   "clear-cache",
	Short: "Clears the Shop cache",
	RunE: func(cmd *cobra.Command, _ []string) error {
		var cfg *platform.Config
		var err error

		if cfg, err = platform.ReadConfig(projectConfigPath, false); err != nil {
			return err
		}

		if cfg.AdminApi == nil {
			logging.FromContext(cmd.Context()).Infof("Clearing cache localy")

			projectRoot, err := findClosestHeyFrameProject()
			if err != nil {
				return err
			}

			return os.RemoveAll(fmt.Sprintf("%s/var/cache", projectRoot))
		}

		logging.FromContext(cmd.Context()).Infof("Clearing cache using admin-api")

		client, err := platform.NewPlatformClient(cmd.Context(), cfg)
		if err != nil {
			return err
		}

		_, err = client.CacheManager.Clear(adminSdk.NewApiContext(cmd.Context()))

		return err
	},
}

func init() {
	projectRootCmd.AddCommand(projectClearCacheCmd)
}

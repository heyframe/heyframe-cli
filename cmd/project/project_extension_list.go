package project

import (
	"encoding/json"
	"fmt"
	"os"

	adminSdk "github.com/heyframe/go-heyframe-admin-api-sdk"
	"github.com/spf13/cobra"

	"github.com/heyframe/heyframe-cli/internal/table"
	"github.com/heyframe/heyframe-cli/platform"
)

var projectExtensionListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all installed extensions",
	RunE: func(cmd *cobra.Command, _ []string) error {
		var cfg *platform.Config
		var err error

		outputAsJson, _ := cmd.PersistentFlags().GetBool("json")

		if cfg, err = platform.ReadConfig(projectConfigPath, true); err != nil {
			return err
		}

		client, err := platform.NewPlatformClient(cmd.Context(), cfg)
		if err != nil {
			return err
		}

		if _, err := client.ExtensionManager.Refresh(adminSdk.NewApiContext(cmd.Context())); err != nil {
			return err
		}

		extensions, _, err := client.ExtensionManager.ListAvailableExtensions(adminSdk.NewApiContext(cmd.Context()))
		if err != nil {
			return err
		}

		if outputAsJson {
			content, err := json.Marshal(extensions)
			if err != nil {
				return err
			}

			fmt.Println(string(content))

			return nil
		}

		table := table.NewWriter(os.Stdout)
		table.Header([]string{"Name", "Version", "Status"})

		for _, extension := range extensions {
			_ = table.Append([]string{extension.Name, extension.Version, extension.Status()})
		}

		_ = table.Render()

		return nil
	},
}

func init() {
	projectExtensionCmd.AddCommand(projectExtensionListCmd)
	projectExtensionListCmd.PersistentFlags().Bool("json", false, "Output as json")
}

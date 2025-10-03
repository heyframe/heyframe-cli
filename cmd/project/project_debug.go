package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/heyframe/heyframe-cli/extension"
	"github.com/heyframe/heyframe-cli/internal/table"
	"github.com/heyframe/heyframe-cli/logging"
	"github.com/heyframe/heyframe-cli/platform"
)

var projectDebug = &cobra.Command{
	Use:   "debug",
	Short: "Shows detected HeyFrame version and detected extensions for further debugging",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		args[0], err = filepath.Abs(args[0])
		if err != nil {
			return err
		}

		shopCfg, err := platform.ReadConfig(projectConfigPath, true)
		if err != nil {
			return err
		}

		heyframeConstraint, err := extension.GetHeyFrameProjectConstraint(args[0])
		if err != nil {
			return err
		}

		if shopCfg.IsFallback() {
			fmt.Printf("Could not find a %s, using fallback config\n", projectConfigPath)
		} else {
			fmt.Printf("Found config: Yes\n")
		}
		fmt.Printf("Detected following HeyFrame version: %s\n", heyframeConstraint.String())

		sources := extension.FindAssetSourcesOfProject(logging.DisableLogger(cmd.Context()), args[0], shopCfg)

		fmt.Println("Following extensions/bundles has been detected")
		table := table.NewWriter(os.Stdout)
		table.Header([]string{"Name", "Path"})

		for _, source := range sources {
			_ = table.Append([]string{source.Name, source.Path})
		}

		_ = table.Render()

		return nil
	},
}

func init() {
	projectRootCmd.AddCommand(projectDebug)
}

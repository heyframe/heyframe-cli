package extension

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/heyframe/heyframe-cli/extension"
)

var extensionAssetBundleCmd = &cobra.Command{
	Use:   "build [path]",
	Short: "Builds assets for extensions",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		assetCfg := extension.AssetBuildConfig{
			HeyFrameRoot: os.Getenv("HEYFRAME_PROJECT_ROOT"),
		}
		validatedExtensions := make([]extension.Extension, 0)

		for _, arg := range args {
			path, err := filepath.Abs(arg)
			if err != nil {
				return fmt.Errorf("cannot open file: %w", err)
			}

			ext, err := extension.GetExtensionByFolder(path)
			if err != nil {
				return fmt.Errorf("cannot open extension: %w", err)
			}

			validatedExtensions = append(validatedExtensions, ext)
		}

		if assetCfg.HeyFrameRoot != "" {
			constraint, err := extension.GetHeyFrameProjectConstraint(assetCfg.HeyFrameRoot)
			if err != nil {
				return fmt.Errorf("cannot get heyframe version constraint from project %s: %w", assetCfg.HeyFrameRoot, err)
			}
			assetCfg.HeyFrameVersion = constraint
		} else {
			constraint, err := validatedExtensions[0].GetHeyFrameVersionConstraint()
			if err != nil {
				return fmt.Errorf("cannot get heyframe version constraint: %w", err)
			}

			assetCfg.HeyFrameVersion = constraint
		}

		if err := extension.BuildAssetsForExtensions(cmd.Context(), extension.ConvertExtensionsToSources(cmd.Context(), validatedExtensions), assetCfg); err != nil {
			return fmt.Errorf("cannot build assets: %w", err)
		}

		return nil
	},
}

func init() {
	extensionRootCmd.AddCommand(extensionAssetBundleCmd)
}

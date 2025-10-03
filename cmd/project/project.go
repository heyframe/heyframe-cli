package project

import (
	"github.com/spf13/cobra"

	"github.com/heyframe/heyframe-cli/platform"
)

var projectConfigPath string

var projectRootCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage your HeyFrame Project",
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(projectRootCmd)
	projectRootCmd.PersistentFlags().StringVar(&projectConfigPath, "project-config", platform.DefaultConfigFileName(), "Path to config")
}

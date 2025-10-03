package project

import (
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/heyframe/heyframe-cli/extension"
	"github.com/heyframe/heyframe-cli/internal/phpexec"
	"github.com/heyframe/heyframe-cli/platform"
)

var projectFrontendWatchCmd = &cobra.Command{
	Use:     "frontend-watch [path]",
	Short:   "Starts the HeyFrame Frontend Watcher",
	Aliases: []string{"watch-frontend"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var projectRoot string
		var err error

		if len(args) == 1 {
			projectRoot = args[0]
		} else if projectRoot, err = findClosestHeyFrameProject(); err != nil {
			return err
		}

		if err := extension.LoadSymfonyEnvFile(projectRoot); err != nil {
			return err
		}

		shopCfg, err := platform.ReadConfig(projectConfigPath, true)
		if err != nil {
			return err
		}

		if err := filterAndWritePluginJson(cmd, projectRoot, shopCfg); err != nil {
			return err
		}

		if err := runTransparentCommand(commandWithRoot(phpexec.ConsoleCommand(cmd.Context(), "feature:dump"), projectRoot)); err != nil {
			return err
		}

		activeOnly := "--active-only"

		if !themeCompileSupportsActiveOnly(projectRoot) {
			activeOnly = "-v"
		}

		if err := runTransparentCommand(commandWithRoot(phpexec.ConsoleCommand(cmd.Context(), "theme:compile", activeOnly), projectRoot)); err != nil {
			return err
		}

		if err := runTransparentCommand(commandWithRoot(phpexec.ConsoleCommand(cmd.Context(), "theme:dump"), projectRoot)); err != nil {
			return err
		}

		if err := os.Setenv("PROJECT_ROOT", projectRoot); err != nil {
			return err
		}

		if err := os.Setenv("STOREFRONT_ROOT", extension.PlatformPath(projectRoot, "Frontend", "")); err != nil {
			return err
		}

		if _, err := os.Stat(extension.PlatformPath(projectRoot, "Frontend", "Resources/app/frontend/node_modules/webpack-dev-server")); os.IsNotExist(err) {
			if err := extension.InstallNPMDependencies(cmd.Context(), extension.PlatformPath(projectRoot, "Frontend", "Resources/app/frontend"), extension.NpmPackage{Dependencies: map[string]string{"not-empty": "not-empty"}}); err != nil {
				return err
			}
		}

		return runTransparentCommand(commandWithRoot(exec.CommandContext(cmd.Context(), "npm", "run-script", "hot-proxy"), extension.PlatformPath(projectRoot, "Frontend", "Resources/app/frontend")))
	},
}

func themeCompileSupportsActiveOnly(projectRoot string) bool {
	themeFile := extension.PlatformPath(projectRoot, "Frontend", "Theme/Command/ThemeCompileCommand.php")

	bytes, err := os.ReadFile(themeFile)
	if err != nil {
		return false
	}

	return strings.Contains(string(bytes), "active-only")
}

func init() {
	projectRootCmd.AddCommand(projectFrontendWatchCmd)
	projectFrontendWatchCmd.PersistentFlags().String("only-extensions", "", "Only watch the given extensions (comma separated)")
	projectFrontendWatchCmd.PersistentFlags().String("skip-extensions", "", "Skips the given extensions (comma separated)")
	projectFrontendWatchCmd.PersistentFlags().Bool("only-custom-static-extensions", false, "Only build extensions from custom/static-plugins directory")
}

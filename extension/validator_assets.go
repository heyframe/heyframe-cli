package extension

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/heyframe/heyframe-cli/internal/validation"
)

func validateAssets(ext Extension, check validation.Check) {
	if !ext.GetExtensionConfig().Validation.StoreCompliance {
		return
	}

	for _, resourceDir := range ext.GetResourcesDirs() {
		validateAssetByResourceDir(check, resourceDir)
	}

	for _, extraBundle := range ext.GetExtensionConfig().Build.ExtraBundles {
		bundlePath := ext.GetRootDir()

		if extraBundle.Path != "" {
			bundlePath = fmt.Sprintf("%s/%s", bundlePath, extraBundle.Path)
		} else {
			bundlePath = fmt.Sprintf("%s/%s", bundlePath, extraBundle.Name)
		}

		validateAssetByResourceDir(check, filepath.Join(bundlePath, "Resources"))
	}
}

func validateAssetByResourceDir(check validation.Check, resourceDir string) {
	_, foundAdminBuildFiles := os.Stat(filepath.Join(resourceDir, "public", "administration"))
	foundAdminEntrypoint := hasJavascriptEntrypoint(filepath.Join(resourceDir, "app", "administration", "src"))
	foundFrontendEntrypoint := hasJavascriptEntrypoint(filepath.Join(resourceDir, "app", "frontend", "src"))
	_, foundFrontendDistFiles := os.Stat(filepath.Join(resourceDir, "app", "frontend", "dist"))

	if foundAdminBuildFiles == nil && !foundAdminEntrypoint {
		check.AddResult(validation.CheckResult{
			Path:       resourceDir,
			Identifier: "assets.administration.sources_missing",
			Message:    fmt.Sprintf("Found administration build files in %s but no source files to rebuild the assets.", resourceDir),
			Severity:   validation.SeverityError,
		})
	}

	if foundAdminBuildFiles != nil && foundAdminEntrypoint {
		check.AddResult(validation.CheckResult{
			Path:       resourceDir,
			Identifier: "assets.administration.build_missing",
			Message:    fmt.Sprintf("Found administration source files in %s but no build files. Please run the build command to generate the assets.", resourceDir),
			Severity:   validation.SeverityError,
		})
	}

	if foundFrontendDistFiles != nil && foundFrontendEntrypoint {
		check.AddResult(validation.CheckResult{
			Path:       resourceDir,
			Identifier: "assets.frontend.sources_missing",
			Message:    fmt.Sprintf("Found frontend build files in %s but no source files to rebuild the assets.", resourceDir),
			Severity:   validation.SeverityError,
		})
	}

	if foundFrontendDistFiles == nil && !foundFrontendEntrypoint {
		check.AddResult(validation.CheckResult{
			Path:       resourceDir,
			Identifier: "assets.frontend.build_missing",
			Message:    fmt.Sprintf("Found frontend source files in %s but no build files. Please run the build command to generate the assets.", resourceDir),
			Severity:   validation.SeverityError,
		})
	}
}

func hasJavascriptEntrypoint(jsRoot string) bool {
	entrypointFiles := []string{"main.js", "main.ts"}
	for _, file := range entrypointFiles {
		if _, err := os.Stat(filepath.Join(jsRoot, file)); err == nil {
			return true
		}
	}
	return false
}

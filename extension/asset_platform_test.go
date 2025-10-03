package extension

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/heyframe/heyframe-cli/internal/asset"
	"github.com/heyframe/heyframe-cli/logging"
)

func getTestContext() context.Context {
	logger := logging.NewLogger(false)

	return logging.WithLogger(context.TODO(), logger)
}

func TestGenerateConfigWithAdminAndFrontendFiles(t *testing.T) {
	dir := t.TempDir()

	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "administration", "src"), os.ModePerm))
	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "administration", "src", "main.js"), []byte("test"), os.ModePerm))
	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "frontend", "src"), os.ModePerm))
	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "frontend", "src", "main.js"), []byte("test"), os.ModePerm))

	config := BuildAssetConfigFromExtensions(getTestContext(), []asset.Source{{Name: "FroshTools", Path: dir}}, AssetBuildConfig{})

	assert.True(t, config.Has("FroshTools"))
	assert.True(t, config.RequiresAdminBuild())
	assert.True(t, config.RequiresFrontendBuild())
	assert.Equal(t, "frosh-tools", config["FroshTools"].TechnicalName)
	assert.Equal(t, "Resources/app/administration/src/main.js", *config["FroshTools"].Administration.EntryFilePath)
	assert.Equal(t, "Resources/app/frontend/src/main.js", *config["FroshTools"].Frontend.EntryFilePath)
	assert.Nil(t, config["FroshTools"].Administration.Webpack)
	assert.Nil(t, config["FroshTools"].Frontend.Webpack)
	assert.Equal(t, "Resources/app/administration/src", config["FroshTools"].Administration.Path)
	assert.Equal(t, "Resources/app/frontend/src", config["FroshTools"].Frontend.Path)
}

func TestGenerateConfigWithTypeScript(t *testing.T) {
	dir := t.TempDir()

	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "administration", "src"), os.ModePerm))
	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "administration", "build"), os.ModePerm))

	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "frontend", "src"), os.ModePerm))
	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "frontend", "build"), os.ModePerm))

	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "administration", "src", "main.ts"), []byte("test"), os.ModePerm))

	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "administration", "build", "webpack.config.js"), []byte("test"), os.ModePerm))

	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "frontend", "src", "main.ts"), []byte("test"), os.ModePerm))
	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "frontend", "build", "webpack.config.js"), []byte("test"), os.ModePerm))

	config := BuildAssetConfigFromExtensions(getTestContext(), []asset.Source{{Name: "FroshTools", Path: dir}}, AssetBuildConfig{})

	assert.True(t, config.Has("FroshTools"))
	assert.True(t, config.RequiresAdminBuild())
	assert.True(t, config.RequiresFrontendBuild())
	assert.Equal(t, "frosh-tools", config["FroshTools"].TechnicalName)
	assert.Equal(t, "Resources/app/administration/src/main.ts", *config["FroshTools"].Administration.EntryFilePath)
	assert.Equal(t, "Resources/app/frontend/src/main.ts", *config["FroshTools"].Frontend.EntryFilePath)
	assert.Equal(t, "Resources/app/administration/build/webpack.config.js", *config["FroshTools"].Administration.Webpack)
	assert.Equal(t, "Resources/app/frontend/build/webpack.config.js", *config["FroshTools"].Frontend.Webpack)
	assert.Equal(t, "Resources/app/administration/src", config["FroshTools"].Administration.Path)
	assert.Equal(t, "Resources/app/frontend/src", config["FroshTools"].Frontend.Path)
}

func TestGenerateWebpackCJS(t *testing.T) {
	dir := t.TempDir()

	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "administration", "src"), os.ModePerm))
	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "administration", "build"), os.ModePerm))

	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "frontend", "src"), os.ModePerm))
	assert.NoError(t, os.MkdirAll(filepath.Join(dir, "Resources", "app", "frontend", "build"), os.ModePerm))

	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "administration", "src", "main.ts"), []byte("test"), os.ModePerm))

	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "administration", "build", "webpack.config.cjs"), []byte("test"), os.ModePerm))

	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "frontend", "src", "main.ts"), []byte("test"), os.ModePerm))
	assert.NoError(t, os.WriteFile(filepath.Join(dir, "Resources", "app", "frontend", "build", "webpack.config.cjs"), []byte("test"), os.ModePerm))

	config := BuildAssetConfigFromExtensions(getTestContext(), []asset.Source{{Name: "FroshTools", Path: dir}}, AssetBuildConfig{})

	assert.True(t, config.Has("FroshTools"))
	assert.True(t, config.RequiresAdminBuild())
	assert.True(t, config.RequiresFrontendBuild())
	assert.Equal(t, "frosh-tools", config["FroshTools"].TechnicalName)
	assert.Equal(t, "Resources/app/administration/src/main.ts", *config["FroshTools"].Administration.EntryFilePath)
	assert.Equal(t, "Resources/app/frontend/src/main.ts", *config["FroshTools"].Frontend.EntryFilePath)
	assert.Equal(t, "Resources/app/administration/build/webpack.config.cjs", *config["FroshTools"].Administration.Webpack)
	assert.Equal(t, "Resources/app/frontend/build/webpack.config.cjs", *config["FroshTools"].Frontend.Webpack)
	assert.Equal(t, "Resources/app/administration/src", config["FroshTools"].Administration.Path)
	assert.Equal(t, "Resources/app/frontend/src", config["FroshTools"].Frontend.Path)
}

func TestGenerateConfigAddsFrontendAlwaysAsEntrypoint(t *testing.T) {
	config := BuildAssetConfigFromExtensions(getTestContext(), []asset.Source{}, AssetBuildConfig{})

	assert.False(t, config.RequiresFrontendBuild())
	assert.False(t, config.RequiresAdminBuild())
}

func TestGenerateConfigDoesNotAddExtensionWithoutConfig(t *testing.T) {
	dir := t.TempDir()

	config := BuildAssetConfigFromExtensions(getTestContext(), []asset.Source{{Name: "FroshApp", Path: dir}}, AssetBuildConfig{})

	assert.False(t, config.Has("FroshApp"))
}

func TestGenerateConfigDoesNotAddExtensionWithoutName(t *testing.T) {
	dir := t.TempDir()

	config := BuildAssetConfigFromExtensions(getTestContext(), []asset.Source{{Name: "", Path: dir}}, AssetBuildConfig{})

	assert.Len(t, config, 0)
}

func TestOnlyFilterOnAssetConfig(t *testing.T) {
	cfg := make(ExtensionAssetConfig)

	cfg["FroshTools"] = &ExtensionAssetConfigEntry{}
	cfg["FroshTest"] = &ExtensionAssetConfigEntry{}

	filtered := cfg.Only([]string{"FroshTools"})

	assert.Len(t, filtered, 1)
	assert.Contains(t, filtered, "FroshTools")
}

func TestSkipFilterOnAssetConfig(t *testing.T) {
	cfg := make(ExtensionAssetConfig)

	cfg["FroshTools"] = &ExtensionAssetConfigEntry{}
	cfg["FroshTest"] = &ExtensionAssetConfigEntry{}

	filtered := cfg.Not([]string{"FroshTools"})

	assert.Len(t, filtered, 1)
	assert.Contains(t, filtered, "FroshTest")
}

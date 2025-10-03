package extension

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnippetValidateNoExistingFolderAdmin(t *testing.T) {
	check := &testCheck{}
	plugin := PlatformPlugin{
		path:   "test",
		config: &Config{},
	}

	validateAdministrationSnippets(plugin, check)
}

func TestSnippetValidateNoExistingFolderFrontend(t *testing.T) {
	check := &testCheck{}
	plugin := PlatformPlugin{
		path:   "test",
		config: &Config{},
	}

	validateAdministrationSnippets(plugin, check)
}

func TestSnippetValidateFrontendByPathOneFileIsIgnored(t *testing.T) {
	tmpDir := t.TempDir()

	check := &testCheck{}
	_ = os.MkdirAll(path.Join(tmpDir, "Resources", "snippet"), os.ModePerm)
	_ = os.WriteFile(path.Join(tmpDir, "Resources", "snippet", "frontend.en-GB.json"), []byte(`{}`), os.ModePerm)

	assert.NoError(t, validateFrontendSnippetsByPath(tmpDir, tmpDir, check))
	assert.Len(t, check.Results, 0)
	assert.Len(t, check.Results, 0)
}

func TestSnippetValidateFrontendByPathSameFile(t *testing.T) {
	tmpDir := t.TempDir()

	check := &testCheck{}

	_ = os.MkdirAll(path.Join(tmpDir, "Resources", "snippet"), os.ModePerm)
	_ = os.WriteFile(path.Join(tmpDir, "Resources", "snippet", "frontend.en-GB.json"), []byte(`{"test": "1"}`), os.ModePerm)
	_ = os.WriteFile(path.Join(tmpDir, "Resources", "snippet", "frontend.de-DE.json"), []byte(`{"test": "2"}`), os.ModePerm)

	assert.NoError(t, validateFrontendSnippetsByPath(tmpDir, tmpDir, check))
	assert.Len(t, check.Results, 0)
	assert.Len(t, check.Results, 0)
}

func TestSnippetValidateFrontendByPathTestDifferent(t *testing.T) {
	tmpDir := t.TempDir()

	check := &testCheck{}

	_ = os.MkdirAll(path.Join(tmpDir, "Resources", "snippet"), os.ModePerm)
	_ = os.WriteFile(path.Join(tmpDir, "Resources", "snippet", "frontend.en-GB.json"), []byte(`{"a": "1"}`), os.ModePerm)
	_ = os.WriteFile(path.Join(tmpDir, "Resources", "snippet", "frontend.de-DE.json"), []byte(`{"b": "2"}`), os.ModePerm)

	assert.NoError(t, validateFrontendSnippetsByPath(tmpDir, tmpDir, check))
	assert.Len(t, check.Results, 2)
	assert.Contains(t, check.Results[0].Message, "key /a is missing, but defined in the main language file")
	assert.Contains(t, check.Results[1].Message, "missing key \"/b\" in this snippet file, but defined in the main language")
}

func TestSnippetValidateFindsInvalidJsonInMainFile(t *testing.T) {
	tmpDir := t.TempDir()

	check := &testCheck{}

	_ = os.MkdirAll(path.Join(tmpDir, "Resources", "snippet"), os.ModePerm)
	_ = os.WriteFile(path.Join(tmpDir, "Resources", "snippet", "frontend.en-GB.json"), []byte(`{"a": "1",}`), os.ModePerm)
	_ = os.WriteFile(path.Join(tmpDir, "Resources", "snippet", "frontend.de-DE.json"), []byte(`{"a": "2"}`), os.ModePerm)

	assert.NoError(t, validateFrontendSnippetsByPath(tmpDir, tmpDir, check))
	assert.Len(t, check.Results, 1)
	assert.Contains(t, check.Results[0].Message, "contains invalid JSON")
}

func TestSnippetValidateFindsInvalidJsonInGermanFile(t *testing.T) {
	tmpDir := t.TempDir()

	check := &testCheck{}

	_ = os.MkdirAll(path.Join(tmpDir, "Resources", "snippet"), os.ModePerm)
	_ = os.WriteFile(path.Join(tmpDir, "Resources", "snippet", "frontend.en-GB.json"), []byte(`{"a": "1"}`), os.ModePerm)
	_ = os.WriteFile(path.Join(tmpDir, "Resources", "snippet", "frontend.de-DE.json"), []byte(`{"a": "2",}`), os.ModePerm)

	assert.NoError(t, validateFrontendSnippetsByPath(tmpDir, tmpDir, check))
	assert.Len(t, check.Results, 1)
	assert.Contains(t, check.Results[0].Message, "contains invalid JSON")
}

package system

import (
	"os"
	"path"
)

// GetHeyFrameCliCacheDir returns the base cache directory for heyframe-c
func GetHeyFrameCliCacheDir() string {
	if dir := os.Getenv("HEYFRAME_CLI_CACHE_DIR"); dir != "" {
		return dir
	}

	cacheDir, _ := os.UserCacheDir()

	return path.Join(cacheDir, "heyframe-cli")
}

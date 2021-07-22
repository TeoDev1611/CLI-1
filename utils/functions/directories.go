package functions

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetCacheDir() string {
	home, err := os.UserHomeDir()
    CheckErrors(err,"Code 2","Error in get the home dir","Report the error on github")
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(home, "AppData", "Local", "moldy", "cache")
	case "linux":
		return filepath.Join(home, ".moldy", "cache")
	case "darwin":
		return filepath.Join(home, ".moldy", "cache")
	default:
		return "Error"
	}
}

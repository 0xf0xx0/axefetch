package icons

import (
	"os"
	"strings"
)

// tries to load an icon by path, and if that fails searches for it in the icon dirs
func SearchAndLoadIcon(name string) []string {
	name, ok := Icons[name]
	if !ok {
		return loadIcon(name)
	}
	return loadIcon(name)
}

// loads an icon from a path
func loadIcon(path string) []string {
	contents, err := os.ReadFile(path)
	if err != nil {
		//println(err.Error()) /// TODO: what to do with this error...
		return nil
	}
	/// trim trailing newlines, used for padding the info string when the icon is shorter
	/// Trim over TrimSpace cause TrimSpace doesnt get newlines
	return strings.Split(strings.Trim(string(contents), "\n"), "\n")
}

package icons

import (
	"os"
	"strings"
)

// tries to load an icon by path, and if that fails searches for it in the icon dirs
func SearchAndLoadIcon(uri string) []string {
	path, ok := Icons[uri]
	if !ok {
		return loadIconFromPath(uri)
	}
	return loadIconFromPath(path)
}

func loadIconFromPath(path string) []string {
	contents, err := os.ReadFile(path)
	if err != nil {
		//println(err.Error()) /// TODO: what to do with this error...
		return nil
	}
	/// trim trailing newlines, used for padding the info string when the icon is shorter
	/// Trim over TrimSpace cause TrimSpace doesnt get newlines
	return strings.Split(strings.Trim(string(contents), "\n"), "\n")
}

package icons

import (
	"os"
	"strings"
)

// searches and loads an icon
func LoadIcon(nameOrPath string) []string {
	contents := loadIcon(nameOrPath)
	if contents == nil {
		/// look in icon dirs
		contents = SearchAndLoadIcon(nameOrPath)
		if contents == nil {
			/// doesnt exist
			return nil
		}
	}
	return contents
}

// loads an icon from a path
func loadIcon(path string) []string {
	contents, err := os.ReadFile(path)
	if err != nil {
		//println(err.Error()) /// TODO: what to do with this error...
		return nil
	}
	/// trim trailing newlines, used for padding the info string when the icon is shorter
	return strings.Split(strings.Trim(string(contents), "\n"), "\n")
}

// searches for an icon in the icon dirs and returns it
func SearchAndLoadIcon(name string) []string {
	for _, m := range append([]map[string][]string{}, Asics, Models, Vendors) {
		icon, ok := m[name]
		if !ok {
			continue
		}
		return icon
	}
	return nil
}

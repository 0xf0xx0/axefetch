package icons

import (
	"os"
	"strings"
)

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

// tries to load an icon by path, and if that fails searches for it in the icon dirs
func SearchAndLoadIcon(name string) []string {
	icon := loadIcon(name)
	if icon == nil {
		/// search
		/// MAYBE: does merging the maps result in better perf? doubtful
		/// ordered by most to least used
		for _, m := range append([]map[string][]string{}, Models, ModelFamilies, Asics, Vendors, Misc) {
			var ok bool
			icon, ok = m[name]
			if !ok {
				continue
			}
			return icon
		}
	}
	return icon
}

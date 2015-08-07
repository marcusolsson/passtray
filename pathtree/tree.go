package pathtree

// Items ...
type Items map[string]Items

// Add recursively adds a slice of strings to an Items map and makes parent
// keys as needed.
func (i Items) Add(path []string) {
	if len(path) > 0 {
		if v, ok := i[path[0]]; ok {
			v.Add(path[1:])
		} else {
			result := make(Items)
			result.Add(path[1:])
			i[path[0]] = result
		}
	}
}

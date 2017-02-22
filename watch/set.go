package watch

// Set describes a set of files and directories to watch for changes
type Set struct {
	Root  string // Directory to watch. If it's not absolute, it's relative to Job.Root
	Depth int
}

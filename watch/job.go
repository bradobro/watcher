package watch

type Job struct {
	Verbosity int     // how much we show: -1, 0, 1, or 2
	Sets      []*Set  // what to watch
	Delay     int     // how long to wait before changes
	Debounce  int     // how long o ignore changes after a change
	Tasks     []*task // things to run
}

type task struct {
	pid         int
	commandline string
}

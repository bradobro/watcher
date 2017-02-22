package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	verbose = flag.Bool("v", false, "verbose")
	depth   = flag.Int("depth", 1, "recursion depth")
	dir     = flag.String("dir", ".", "directory root to use for watching")
	quiet   = flag.Duration("quiet", 800*time.Millisecond, "quiet period after command execution")
	ignore  = flag.String("ignore", "", "path ignore pattern")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [flags] [command to execute and args]\n", os.Args[0])
	flag.PrintDefaults()
}

func cli(testInjects ...string) (testResult int) {
	if len(testInjects) == 1 {
		fmt.Println("File watcher thing.")
		testResult = 1
	}
	return
}

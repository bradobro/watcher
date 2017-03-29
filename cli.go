package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

/*
For repeated flags See https://lawlessguy.wordpress.com/2013/07/23/filling-a-slice-using-command-line-flags-in-go-golang/

we could have a per-path depth, ignore, include, but the syntax will be ugly:

-watch './**/* depth=5 exclude=**/*_test.go include=**/*.go include=**/\d\d\_**/*.md exclude=./vendor/**/*'

\d = digit
\_ = literal space (or could use \x0020)

Use flag setter/string for a custom type watchslice = []watch

*/

var (
	verbose = flag.Bool("v", false, "verbose")
	depth   = flag.Int("depth", 1, "recursion depth")
	dir     = flag.String("dir", ".", "directory root to use for watching")
	quiet   = flag.Duration("quiet", 800*time.Millisecond, "quiet period after command execution")
	ignore  = flag.String("ignore", "", "path ignore pattern")
	// command
	// process
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

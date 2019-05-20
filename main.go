package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Individual path prefixes we are searching for.
var components []string

// This is like filepath.Glob, but it's a bit quirky. It tries to find
// project root directories using prefixes stored in `components`.
func quirkyGlob(dir string, index int) (m []string, e error) {
	m = []string{}
	e = nil

	// Do not iterate through hidden directories.
	if strings.HasPrefix(path.Base(dir), ".") {
		return
	}

	fi, err := os.Stat(dir)
	// Ignore the directory if we couldn't stat it.
	if err != nil {
		return
	}

	// Only directories are considered, we can't CD into files.
	if !fi.IsDir() {
		return
	}

	// If the current directory matches the component, increment the
	// component index and keep searching. Do not continue searching
	// sub-directories for the same component index since they will
	// all result in longer paths than this one, and we have already
	// satisfied the search for this component.
	if strings.HasPrefix(path.Base(dir), components[index]) {
		index = index + 1
	}

	// If we hit the end of the component list, then we have a match
	// and are done searching.
	if index == len(components) {
		return []string{dir}, nil
	}

	directories, err := filepath.Glob(fmt.Sprint(dir, "/*"))
	if err != nil {
		return m, err
	}

	// If any directory is .git, then we are at a project root and
	// we do not want to continue searching subdirectories.
	for _, fl := range directories {
		matched, err := filepath.Match(".git", path.Base(fl))
		if err != nil {
			continue
		}
		if matched {
			return
		}
	}

	// For each subdirectory, check for matches.
	for _, path := range directories {
		matches, err := quirkyGlob(path, index)

		if err != nil {
			continue
		}

		if len(matches) == 0 {
			continue
		}

		m = append(m, matches...)
	}

	return
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [SEARCH]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  SEARCH string\n")
		fmt.Fprintf(os.Stderr, "         a fuzzy path to search for (example \"s/foo\")\n")
		flag.PrintDefaults()
	}

	var root string
	flag.StringVar(&root, "root", os.Getenv("HOME"), "root directory to search from")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println(root)
		return
	}

	arg := flag.Args()[0]

	// Split the search expression into components on the "/" character.
	for _, c := range strings.Split(arg, "/") {
		components = append(components, c)
	}

	matches, err := quirkyGlob(root, 0)

	// If we didn't get any matches, or we got an error, just return root.
	if err != nil || len(matches) == 0 {
		fmt.Println(root)
		return
	}

	// Find the shortest path that matches.
	shortest := matches[0]
	jumps := strings.Count(shortest, "/")

	for _, path := range matches {
		if j := strings.Count(path, "/"); j < jumps {
			shortest = path
			jumps = j
		}
	}

	fmt.Println(shortest)
}

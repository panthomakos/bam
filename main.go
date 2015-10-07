package main

import (
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

	fi, err := os.Stat(dir)
	if err != nil {
		return
	}

	// Only directories are considered, we can't CD into files.
	if !fi.IsDir() {
		return
	}

	// Skip hidden directories.
	if strings.HasPrefix(path.Base(dir), ".") {
		return
	}

	// If we have reached the end of the components, we have matched.
	if index == len(components) {
		return []string{dir}, nil
	}

	// If the current directory matches the component, increment the
	// component and keep searching. Do not continue searching
	// sub-directories since they will all be longer than this path.
	if strings.HasPrefix(path.Base(dir), components[index]) {
		matches, _ := quirkyGlob(dir, index+1)
		if err != nil {
			return m, err
		}
		return append(m, matches...), nil
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
	root := filepath.Join(os.Getenv("HOME"), "Projects")

	if len(os.Args) == 1 {
		fmt.Println(root)
		return
	}

	arg := os.Args[1]

	for _, c := range strings.Split(arg, "/") {
		components = append(components, c)
	}

	matches, err := quirkyGlob(root, 0)

	if err != nil || len(matches) == 0 {
		fmt.Println(root)
		return
	}

	shortest := matches[0]

	for _, path := range matches {
		if len(path) < len(shortest) {
			shortest = path
		}
	}

	fmt.Println(shortest)
}

//go:build (linux || aix || zos) && !appengine
// +build linux aix zos
// +build !appengine

package isatty

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// IsTerminal return true if the file descriptor is terminal.
func IsTerminal(fd uintptr) bool {
	_, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
	fmt.Println("Test change")
	return err == nil
}

// IsCygwinTerminal return true if the file descriptor is a cygwin or msys2
// terminal. This is also always false on this environment.
func IsCygwinTerminal(fd uintptr) bool {
	return false
}

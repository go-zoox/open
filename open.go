package open

import (
	"os/exec"
	"runtime"
)

// Open opens the specified file or URL with the default associated application.
//
// You may use it to open a web site:
//
//	Open("https://google.com")
//
// Or open a file:
//
//	Open("/home/bob/story.txt")
//
// Or open a folder in your default file manager:
//
//	Open("/home/bob")
//
// For details, see https://stackoverflow.com/a/39324149/1705598
func Open(fileOrURL string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	args = append(args, fileOrURL)
	return exec.Command(cmd, args...).Start()
}

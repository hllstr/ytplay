package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Audio Player
func Play(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("afplay", path)
	// IOS support using mpv (via ISH SHELL)
	case "linux", "android", "ios":
		cmd = exec.Command("mpv", "--no-video", path)
	case "windows":
		ps := fmt.Sprintf("(New-Object Media.SoundPlayer '%s').PlaySync()", path)
		cmd = exec.Command("powershell", "-c", ps)
	default:
		return fmt.Errorf("OS tidak didukung: %s", runtime.GOOS)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

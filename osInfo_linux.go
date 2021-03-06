package osinfo

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func GetInfo() (*OSInfo, error) {
	out, err := getInfo()
	if err != nil {
		return nil, err
	}

	for strings.Index(out, "broken pipe") != -1 {
		out, _ = getInfo()
		time.Sleep(500 * time.Millisecond)
	}
	osStr := strings.Replace(out, "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	osInfo := strings.Split(osStr, " ")

	gio := &OSInfo{GoOSBig: osInfo[0], Kernel: osInfo[1], Platform: osInfo[2], OS: osInfo[3], GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
	gio.Hostname, _ = os.Hostname()

	return gio, nil
}

func getInfo() (string, error) {
	cmd := exec.Command("uname", "-srio")

	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return out.String(), nil
}

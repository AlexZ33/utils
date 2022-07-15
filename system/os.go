// Package system contain some functions about os, runtime, shell command
package system

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
)

func IsWindow() bool {
	return runtime.GOOS == "windows"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}

func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// GetOsEnv Getenv retrieves the value of the environment variable named by the key.
//It returns the value, which will be empty if the variable is not present.
//To distinguish between an empty value and an unset value, use os.LookupEnv.
func GetOsEnv(key string) string {
	return os.Getenv(key)
}

// SetOsEnv sets the value of the environment variable named by the key.
//It returns an error, if any.
func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

// CompareOsEnv gets env named by the key and compare it with comparedEnv
func CompareOsEnv(key, comparedEnv string) bool {
	env := GetOsEnv(key)
	if env == "" {
		return false
	}
	return env == comparedEnv
}

// ExecCommand use shell /bin/bash -c to execute command
func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errout bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindow() {
		cmd = exec.Command("cmd")
	}
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err = cmd.Run()

	if err != nil {
		stderr = string(errout.Bytes())
	}
	stdout = string(out.Bytes())
	return
}

package utils

import "testing"

func Test_Run(t *testing.T) {
	RunPrint("go", "version")
	RunPrint("mkdir", "-p", "/usr/sbin/")
}

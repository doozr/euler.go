package main

import (
	"os"
	"os/exec"
	"path"
)

func main() {
	gopath := os.Getenv("GOPATH")
	eulerDir := path.Join(gopath, "src", "github.com", "doozr", "euler.go")
	os.Chdir(eulerDir)

	cmd := exec.Command("go", "test")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

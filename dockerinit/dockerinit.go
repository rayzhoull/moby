package main

import (
	_ "github.com/rayzhoull/moby/daemon/execdriver/lxc"
	_ "github.com/rayzhoull/moby/daemon/execdriver/native"
	"github.com/rayzhoull/moby/reexec"
)

func main() {
	// Running in init mode
	reexec.Init()
}

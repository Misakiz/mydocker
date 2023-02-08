//go:build linux
// +build linux

package main

import (
	mydocker "mydocker/run"
)

func main() {
	mydocker.MyDockerRun()

}


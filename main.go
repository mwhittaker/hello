package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		panic(fmt.Errorf("cannot read build info"))
	}
	fmt.Println("hello")
	fmt.Printf("info.Main.Version = %s\n", info.Main.Version)
	for _, setting := range info.Settings {
		fmt.Printf("%s = %s\n", setting.Key, setting.Value)
	}
}

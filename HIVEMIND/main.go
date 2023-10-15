package main

import "HIVEMIND/core/hivemind"

func main() {
	hb := hivemind.Init()
	err := hb.Start()
	if err != nil {
		panic("FATAL HIVEMIND ERROR")
	}
}

package main

import (
        "log"
        ui "github.com/gizak/termui"
)

const (
	VERSION = "0.0.0"
	//USAGE   = 
)

func init() {
}

func main() {
    if err := ui.Init(); err != nil {
        log.Fatalln(err)
    }

    defer ui.Close()
}


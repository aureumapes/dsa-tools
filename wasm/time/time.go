package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("submitEntry", SubmitEntry())
	js.Global().Set("hourName", HourName())
	js.Global().Set("addMinute", AddMinute())
	js.Global().Set("addDay", AddDay())
	js.Global().Set("pad", Pad())
	js.Global().Set("saveStatus", SaveStatus())
	select {}
}

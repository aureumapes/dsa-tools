package main

import (
	"github.com/aureumapes/dsa-tools/wasm/time"
	"syscall/js"
)

func main() {
	js.Global().Set("submitEntry", time.SubmitEntry())
	js.Global().Set("hourName", time.HourName())
	js.Global().Set("addMinute", time.AddMinute())
	select {}
}

package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"syscall/js"
)

func AddMinute() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		timeStr := js.Global().Get("document").Call("getElementById", "time").Get("innerHTML").String()
		timStrs := strings.Split(timeStr, ":")
		hour, _ := strconv.ParseInt(timStrs[0], 10, 64)
		minute, _ := strconv.ParseInt(timStrs[1], 10, 64)
		colors := []string{"#303050", "#354264", "#385676", "#3a6a85", "#3e7f91", "#48949a", "#59a8a0", "#72bca3", "#90cea4", "#b3dfa5", "#d8eea7", "#fffcad"}
		addMinutes := int64(args[0].Int())
		minute += addMinutes
		if minute >= 60 {
			minute -= 60
			hour++
		}
		if hour >= 24 {
			hour = 0
			minute = 0
			js.Global().Call("addDay")
		}
		color := ""
		if hour < 12 {
			color = colors[hour]
		} else {
			slices.Reverse(colors)
			color = colors[hour-12]
		}
		js.Global().Get("document").Call("getElementById", "time").Get("style").Set("background", color)
		timeStr = fmt.Sprintf("%s:%s", js.Global().Call("pad", hour), js.Global().Call("pad", minute))
		js.Global().Get("document").Call("getElementById", "time").Set("innerHTML", timeStr)
		js.Global().Call("hourName")
		js.Global().Call("saveStatus")
		return nil
	})
}

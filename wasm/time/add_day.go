package time

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"syscall/js"
)

func AddDay() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		months := []string{"PRA", "RON", "EFF", "TRA", "BOR", "HES", "FIR", "TSA", "PHE", "PER", "ING", "RAH", "NL"}
		document := js.Global().Get("document")
		document.Call("getElementById", "time").Set("innerText", "00:00")
		date := strings.Split(document.Call("getElementById", "date").Get("innerText").String(), " ")
		year, _ := strconv.ParseInt(strings.Split(date[2], "BF")[0], 10, 64)
		month := slices.Index(months, date[1]) + 1
		day, _ := strconv.ParseInt(strings.Split(date[0], ".")[0], 10, 64)
		day++
		if month != 13 && day > 30 {
			month++
			day = 1
		} else if month == 13 && day > 5 {
			month = 1
			day = 1
			year++
		}
		document.Call("getElementById", "date").Set("innerText", fmt.Sprintf("%d. %s %dBF", day, months[month-1], year))
		js.Global().Call("saveStatus")
		return nil
	})
}

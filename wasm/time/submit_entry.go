package time

import (
	"strconv"
	"syscall/js"
)

func SubmitEntry() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		text := js.Global().Get("document").Call("getElementById", "entry-text").Get("value").String()
		jubilee := js.Global().Get("document").Call("getElementById", "jubilee").Get("checked").Bool()
		url := "/entry?jubilee=" + strconv.FormatBool(jubilee)
		data := js.ValueOf(map[string]any{
			"method": "PUT",
			"body":   text,
			"headers": map[string]any{
				"Content-Type": "text/plain",
			},
		})
		js.Global().Call("fetch", url, data)
		return nil
	})
}

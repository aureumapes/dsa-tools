package time

import (
	"syscall/js"
)

func SaveStatus() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		text := js.Global().Get("document").Call("getElementById", "date").Get("innerText").String()
		url := "/save"
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

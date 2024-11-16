package main

import (
	"encoding/json"
	"syscall/js"
)

type Char struct {
	Name        string `gorm:"primarykey"`
	TDay        string
	BDay        string
	Type        string
	Description string
	Member      string
}

func submit() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		document := js.Global().Get("document")
		name := document.Call("getElementById", "name").Get("value").String()
		description := document.Call("getElementById", "description").Get("value").String()
		tday := document.Call("getElementById", "tday").Get("value").String()
		bday := document.Call("getElementById", "bday").Get("value").String()
		member := document.Call("getElementById", "member").Get("value").String()
		char := Char{
			Name:        name,
			TDay:        tday,
			BDay:        bday,
			Type:        "SC",
			Description: description,
			Member:      member,
		}
		charJson, _ := json.Marshal(&char)
		url := "/newchar"
		data := js.ValueOf(map[string]any{
			"method": "PUT",
			"body":   string(charJson),
			"headers": map[string]any{
				"Content-Type": "text/plain",
			},
		})
		js.Global().Call("fetch", url, data)
		return nil
	})
}

func main() {
	js.Global().Set("send", submit())
	select {}
}

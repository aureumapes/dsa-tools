package time

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"
	"syscall/js"
)

func determineTextColor(colour color.RGBA) string {
	if 0.2126*float64(colour.R)+0.7152*float64(colour.G)+0.0722*float64(colour.B) > math.Sqrt(1.05*0.05)-0.05 {
		return "black"
	} else {
		return "white"
	}
}

func HourName() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		hourString := strings.Split(js.Global().Get("document").Call("getElementById", "time").Get("innerText").String(), ":")[0]
		b, _ := hex.DecodeString("ffd700")
		hour, _ := strconv.ParseInt(hourString, 10, 64)
		names := []string{
			"Praiosstunde",
			"Rondrastunde",
			"Efferdsstunde",
			"Traviastunde",
			"Boronsstunde",
			"Hesindestunde",
			"Firunsstunde",
			"Tsastunde",
			"Phexensstunde",
			"Perainestunde",
			"Ingerimmstunde",
			"Rahjasstunde",
		}
		colours := []string{
			"ffd700",
			"8b0000",
			"7fffd4",
			"ff8c00",
			"000000",
			"663399",
			"00ffff",
			"ff69b4",
			"a9a9a9",
			"006400",
			"b22222",
			"9690AB",
		}
		if hour >= 12 {
			b, _ = hex.DecodeString(colours[hour-12])
		} else {
			b, _ = hex.DecodeString(colours[hour])
		}
		background := color.RGBA{R: b[0], G: b[1], B: b[2], A: 0xff}
		c := determineTextColor(background)
		bStr := fmt.Sprintf("#%X", []byte{background.R, background.G, background.B})
		name := names[0]
		if hour >= 12 {
			name = names[hour-12]
		} else {
			name = names[hour]
		}
		js.Global().Get("document").Call("getElementById", "name").Get("style").Set("color", c)
		js.Global().Get("document").Call("getElementById", "name").Get("style").Set("background", bStr)
		js.Global().Get("document").Call("getElementById", "name").Set("innerText", name)
		return nil
	})
}

func Pad() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		d := args[0].Int()
		if d < 10 {
			return fmt.Sprintf("0%d", d)
		} else {
			return fmt.Sprintf("%d", d)
		}
	})
}

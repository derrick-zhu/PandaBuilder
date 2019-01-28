package logger

import (
	"fmt"

	"github.com/bclicn/color"
)

func formatStr(tmpl string, content ...interface{}) string {
	finalTmpl := color.BLightBlue("*** ") + tmpl
	finalStr := fmt.Sprintf(finalTmpl, content...)
	return finalStr
}

func formatRawStr(tmpl string, content ...interface{}) string {
	return fmt.Sprintf(tmpl, content...)
}

func Println(tmpl string, content ...interface{}) {
	Print(formatStr(tmpl, content...) + "\n")
}

func Printf(tmpl string, content ...interface{}) {
	Print(formatStr(tmpl, content...))
}

func PrintlnRaw(tmpl string, content ...interface{}) {
	Print(formatRawStr(tmpl, content...) + "\n")
}

func PrintfRaw(tmpl string, content ...interface{}) {
	Print(formatRawStr(tmpl, content...))
}

package util

import "github.com/russross/blackfriday/v2"

func Md2html(in string) string {
	input := []byte(in)
	unsafe := blackfriday.Run(input, blackfriday.WithExtensions(blackfriday.CommonExtensions))
	html := string(unsafe)
	return html
}
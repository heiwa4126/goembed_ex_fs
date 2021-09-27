package main

import (
	"embed"
	"log"
)

//go:embed static other_static
var f embed.FS

// 複数のディレクトリをまとめられる。他特定のextentionのファイルだけ、とかも出来る。
// https://pkg.go.dev/embed の例参照。

func readFromEmbedFS(file string) string {
	data, err := f.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	print(readFromEmbedFS("static/hello.txt"))
	print(readFromEmbedFS("other_static/goodbye.txt"))
}

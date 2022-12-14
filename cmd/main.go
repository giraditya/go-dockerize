package main

import (
	"fmt"

	"github.com/giriaditya/hot-reload/pkg/cmd/http"
)

func main() {
	http.RunHttp()
	echoHello()
}

func echoHello() {
	fmt.Println("Hello")
}

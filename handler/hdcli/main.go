package main

import (
	"fmt"
	"os"

	"github.com/sihiver/BrowserBridge/handler"
)

func main() {

	fmt.Println("V3")
	handler.Handle(&handler.HandleSettings{
		ListenAddr: os.Args[2],
		RemoteAddr: os.Args[1],
	})
}

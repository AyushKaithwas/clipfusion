package main

import (
	"fmt"

	"github.com/AyushKaithwas/clipfusion/server"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	closeChan := make(chan bool)
	sigChan := make(chan os.Signal, 1)

	// TODO: Bind various API endpoints to serve stuff for our perverse needs

	go func() {
		fmt.Println("chronicler is listening on :8080")
		fmt.Println("jon sucks")
		http.ListenAndServe(":8080", http.FileServer(http.Dir("./www")))
		closeChan <- true
	}()

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigChan:
	case <-closeChan:
	}
}

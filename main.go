package main

import (
	"context"
	"fmt"
	"time"

	"github.com/percipia/eslgo"
	"github.com/percipia/eslgo/command"
)

func main() {
	eslgo.ListenAndServe(":8084", handleConnection)
}

func handleConnection(ctx context.Context, conn *eslgo.Conn, response *eslgo.RawResponse) {
	fmt.Printf("Got connection! %#v\n", response)

	response, err := conn.SendCommand(ctx, command.API{Command: "echo", Arguments: "", Background: false})
	if err != nil {
		panic(err)
	}

	if response.IsOk() {
		fmt.Println("Echoing!")
	}

	fmt.Println("Sleeping for 5s")
	time.Sleep(5000)
	fmt.Println("Waking up")

	conn.HangupCall(ctx, response.ChannelUUID(), "NORMAL_CLEARING")
}

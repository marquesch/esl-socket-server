package main

import (
	"context"
	"fmt"
	"net/textproto"

	"github.com/percipia/eslgo"
	"github.com/percipia/eslgo/command"
)


func main() {
	eslgo.ListenAndServe(":8084", handleConnection)
}

func handleConnection(ctx context.Context, conn *eslgo.Conn, response *eslgo.RawResponse) {
	defer conn.Close()
	uuid := response.GetHeader("Unique-ID")
	fmt.Println("UUID:", uuid)

	callCh := make(chan bool)
	conn.RegisterEventListener(uuid, func(event *eslgo.Event) {
		fmt.Printf("Got event %s\n", event.GetHeader("Event-Name"))
		if event.GetHeader("Event-Name") == "CHANNEL_HANGUP_COMPLETE" {
			close(callCh)
		}
	})

	conn.EnableEvents(ctx)
	conn.SendCommand(ctx, command.Linger{Enabled: true})

	_, err := conn.SendCommand(ctx, &command.SendMessage{
		UUID: uuid,
		Headers: textproto.MIMEHeader{
			"call-command":     []string{"execute"},
			"execute-app-name": []string{"bridge"},
			"execute-app-arg": []string{
				"{hangup_after_bridge=true,ignore_early_media=true}user/1000",
			},
		},
		Sync:    false,
		SyncPri: false,
	})
	if err != nil {
		fmt.Println("Bridge error:", err)
		return
	}
	fmt.Println("Bridge command sent!")

	<-callCh
	fmt.Println("Call hungup! Cleaning up resources.")

	fmt.Println("Call completed:", uuid)
}

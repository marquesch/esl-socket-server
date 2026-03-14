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
	uuid := response.GetHeader("Unique-ID")
	fmt.Println("UUID:", uuid)

	conn.EnableEvents(ctx)
	conn.SendCommand(ctx, command.Linger{Enabled: true})

	_, err := conn.SendCommand(ctx, &command.SendMessage{
		UUID: uuid,
		Headers: textproto.MIMEHeader{
			"call-command":     []string{"execute"},
			"execute-app-name": []string{"bridge"},
			"execute-app-arg":  []string{"{ignore_early_media=true}user/1000"},
		},
		Sync: false,
		SyncPri: false,
	})
	if err != nil {
		fmt.Println("Bridge error:", err)
		return
	}
	fmt.Println("Bridge command sent, waiting...")

	conn.SendCommand(ctx, &command.SendMessage{
		UUID: uuid,
		Headers: textproto.MIMEHeader{
			"call-command": []string{"hangup"},
			"hangup-cause": []string{"NORMAL_CLEARING"},
		},
	})

	fmt.Println("Call completed:", uuid)
}

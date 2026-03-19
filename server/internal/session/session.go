package session

import (
	"context"
	"net/textproto"

	"github.com/marquesch/esl-socket-server/internal/call"
	"github.com/percipia/eslgo"
	"github.com/percipia/eslgo/command"
)

type ESLSession struct {
	conn        *eslgo.Conn
	ChannelUUID string
}

func (e *ESLSession) Bridge(
	ctx context.Context,
	dialParams call.DialParams,
) (*eslgo.RawResponse, error) {
	return e.conn.SendCommand(ctx, &command.SendMessage{
		UUID: e.ChannelUUID,
		Headers: textproto.MIMEHeader{
			"call-command":     []string{"execute"},
			"execute-app-name": []string{"bridge"},
			"execute-app-arg":  []string{dialParams.String()},
		},
		Sync:    false,
		SyncPri: false,
	})
}

func New(conn *eslgo.Conn, uuid string) *ESLSession {
	return &ESLSession{conn: conn, ChannelUUID: uuid}
}

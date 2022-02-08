package socketio

import (
	socketio "github.com/googollee/go-socket.io"
)

const MAX_MESSAGE_LENGTH = 500

type IngressMessage struct {
	ToRoom  string `json:"toRoom"`
	Content string `json:"content"`
}

type EgressMessage struct {
	ToRoom   string `json:"toRoom"`
	FromId   string `json:"fromId"`
	FromName string `json:"fromName"`
	Content  string `json:"content"`
}

func handleChat(s socketio.Conn, m IngressMessage) {
	if len(m.ToRoom) == 0 || len(m.Content) == 0 || len(m.Content) > MAX_MESSAGE_LENGTH {
		s.Emit("chat:error", "message format error")
		return
	}
	if s.Context() == nil {
		s.Emit("chat:error", "unauthorized")
		return
	}
	context := s.Context().(SocketContext)
	resp := EgressMessage{
		ToRoom:   m.ToRoom,
		FromId:   context.Id,
		FromName: context.Username,
		Content:  m.Content,
	}
	IO.BroadcastToRoom("/", resp.ToRoom, "chat", resp)
}

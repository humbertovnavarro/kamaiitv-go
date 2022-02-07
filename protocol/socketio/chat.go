package socketio

import socketio "github.com/googollee/go-socket.io"

func handleChat(s socketio.Conn, msg string) {
	s.Emit("chat", "Hello, world!")
}

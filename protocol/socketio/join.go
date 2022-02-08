package socketio

import socketio "github.com/googollee/go-socket.io"

func handleJoin(s socketio.Conn, room string) {
	s.Join(room + ":public")
	s.Emit("join:success", room+":public")
}

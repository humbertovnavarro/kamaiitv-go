package socketio

import socketio "github.com/googollee/go-socket.io"

func handleJoin(s socketio.Conn, room string) {
	if len(s.Rooms()) > 300 {
		s.Emit("join:error", "joined too many rooms")
		return
	}
	s.Join(room + ":public")
	s.Emit("join:success", room+":public")
}

func handleLeave(s socketio.Conn, room string) {
	s.Leave(room)
}

func handleLeaveAll(s socketio.Conn, room string) {
	s.LeaveAll()
}

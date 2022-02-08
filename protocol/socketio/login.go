package socketio

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/gwuhaolin/livego/lib"
)

func handleLogin(s socketio.Conn, token string) {
	claims, err := lib.DecodeToken(token)
	if err != nil {
		s.Emit("login:error", "bad token")
		return
	}
	privRoom := claims.Id + ":private"
	s.Join(privRoom)
	s.Emit("login:success", claims.Subject+":"+claims.Id)
	context := SocketContext{
		Id:       claims.Id,
		Username: claims.Subject,
	}
	// Addresses multiple username edge case
	IO.ForEach("/", privRoom, func(c socketio.Conn) {
		c.SetContext(context)
	})
}

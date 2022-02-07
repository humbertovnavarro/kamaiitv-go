package socketio

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

var IO = &socketio.Server{}

func Start(redisAddress string) {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	IO = server
	_, err := server.Adapter(&socketio.RedisAdapterOptions{
		Addr:    redisAddress,
		Network: "tcp",
	})
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		authenticateSocket(s)
		return nil
	})

	server.OnEvent("/", "chat", func(s socketio.Conn, msg string) {
		handleChat(s, msg)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()
	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func authenticateSocket(socketio.Conn) string {
	return "1234"
}

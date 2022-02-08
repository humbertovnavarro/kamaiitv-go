package socketio

import (
	"net/http"

	log "github.com/sirupsen/logrus"

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

type SocketContext struct {
	Id       string
	Username string
}

func Start(redisAddress string) {
	// Cors
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
	// Redis Cluster
	_, err := server.Adapter(&socketio.RedisAdapterOptions{
		Addr:    redisAddress,
		Network: "tcp",
	})
	if err != nil {
		log.Fatal(err)
	}
	IO = server

	server.OnConnect("/", handleConnect)
	server.OnEvent("/", "chat", handleChat)
	server.OnEvent("/", "join", handleJoin)
	server.OnEvent("/", "login", handleLogin)
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

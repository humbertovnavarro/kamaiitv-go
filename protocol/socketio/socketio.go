package socketio

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

_, err := server.Adapter(&socketio.RedisAdapterOptions{
    Host:   "127.0.0.1",
    Port:   "6379",
    Prefix: "socket.io",
})
if err != nil {
    log.Fatal("error:", err)
}

func Start() {
	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	go server.Serve()
	defer server.Close()
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

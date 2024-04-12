package chat

import (
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

func BenchmarkSendMessage(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/room"}
		conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			b.Fatalf("could not open websocket connection: %v", err)
		}
		defer conn.Close()

		message := []byte("Hello, Parallel World!")
		for pb.Next() {
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				b.Fatalf("failed to write message: %v", err)
			}
		}
	})
}

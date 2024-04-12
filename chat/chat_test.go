package chat

import (
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

func BenchmarkMessageThroughput(b *testing.B) {
	// Setup WebSocket connection
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/room"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		b.Fatalf("could not open websocket connection: %v", err)
	}
	defer conn.Close()

	message := []byte("test message")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			b.Fatalf("failed to write message: %v", err)
		}
	}
}

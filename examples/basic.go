package main

import (
	"github.com/recws-org/recws"
	"log"
	"time"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	ws := recws.RecConn{
		//KeepAliveTimeout: 10 * time.Second,
	}
	ws.Dial("ws://127.0.0.1:3563", nil)

	//go func() {
	//	time.Sleep(2 * time.Second)
	//	cancel()
	//}()

	for {
		select {
		//case <-ctx.Done():
		//	go ws.Close()
		//	log.Printf("Websocket closed %s", ws.GetURL())
		//	return
		default:
			if !ws.IsConnected() {
				log.Printf("Websocket disconnected %s", ws.GetURL())
				<-time.After(3 * time.Second)
			}

			if err := ws.WriteMessage(1, []byte("Incoming")); err != nil {
				log.Printf("Error: WriteMessage %s", ws.GetURL())
				<-time.After(3 * time.Second)
			}

			_, _, err := ws.ReadMessage()
			if err != nil {
				log.Printf("Error: ReadMessage %s", ws.GetURL())
				<-time.After(3 * time.Second)
			}
		}
	}
}

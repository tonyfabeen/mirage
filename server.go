package server 

import(
	"fmt"
	"config"
	"http"
	"websocket"
	"io"
)


type Server struct{ }

func ChannelHandler(ws *websocket.Conn){
	io.Copy(ws,ws)
}

func (server *Server) Start(){
	var config = new(config.Config)
	ok := config.Parse()

	if ok {
		for key, channel := range config.Channels {
		    fmt.Printf("K %s",key)
			http.Handle("/"+ channel, websocket.Handler(ChannelHandler))
		}
		err := http.ListenAndServe(":"+config.Port, nil)
		if err != nil {
			panic("ListenAndServe: ", err.String())
		}
	}
}

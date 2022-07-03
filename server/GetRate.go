package server

import (
	"fmt"
	"net/http"
)

func (srv *Server) GetRate(resp http.ResponseWriter, req *http.Request) {
	_, err := http.Post("http://127.0.0.1:8081/api/publisher", "application/json;content=UTF-8", req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (srv *Server) Health(resp http.ResponseWriter, req *http.Request) {
	_, err := resp.Write([]byte("Success"))
	if err != nil {
		return
	}
}

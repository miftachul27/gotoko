package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	fmt.Println("Selamat datang di GoToko") // Memperbaiki pernyataan print
	server.Router = mux.NewRouter()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Mendengarkan port %s\n", addr)          // Memperbaiki pernyataan print
	log.Fatal(http.ListenAndServe(addr, server.Router)) // Memperbaiki nama fungsi ListenAndServe
}

func Run() {
	var server = Server{}

	server.Initialize() // Memperbaiki pemanggilan metode initialize
	server.Run(":9000") // Menghapus addr: dan memperbaiki port
}

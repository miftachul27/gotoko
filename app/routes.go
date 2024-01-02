package app

import (
	"github.com/miftcahul27/gotoko/controller" // Ganti dengan path yang sesuai
)

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET") // Memperbaiki format pemanggilan HandleFunc dan Methods
}

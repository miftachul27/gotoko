package app

import (
	"github.com/miftcahul27/gotoko/controllers" // Ganti dengan path yang sesuai
)

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET") // Memperbaiki format pemanggilan HandleFunc dan Methods
}

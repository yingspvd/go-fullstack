package controllers

import (
	"github.com/yingspvd/go-fullstack/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcom To This Awesome API")
}

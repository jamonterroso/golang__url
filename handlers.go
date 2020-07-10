package main

import (
	"encoding/json"
	"net/http"
)

//Usersurl redirecciona a la pagina de usuarios
func Usersurl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a, _ := obtenerResultados()
	json.NewEncoder(w).Encode(a)
}

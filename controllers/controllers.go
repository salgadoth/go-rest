package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/salgadoth/go-rest-api/database"
	"github.com/salgadoth/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var p []models.Personalidade

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var p models.Personalidade

	json.NewDecoder(r.Body).Decode(&p)

	database.DB.Create(&p)

	json.NewEncoder(w).Encode(p)
}

func RemovePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var p models.Personalidade

	database.DB.Delete(&p, id)

	json.NewEncoder(w).Encode(p)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)

	json.NewDecoder(r.Body).Decode(&p)

	database.DB.Save(&p)

	json.NewEncoder(w).Encode(p)
}

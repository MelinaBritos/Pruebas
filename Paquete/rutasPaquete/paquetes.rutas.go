package rutasPaquete

import (
	"encoding/json"
	"net/http"
	"strconv"

	dataPaquete "github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/modelosPaquete"
	"github.com/MelinaBritos/TP-Principal-AMAZONA/Paquete/validaciones"

	"github.com/gorilla/mux"
)

func GetPaquetesHandler(w http.ResponseWriter, r *http.Request) {

	paquetes := dataPaquete.ObtenerPaquetes()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(paquetes); err != nil {
		http.Error(w, "Error al codificar paquetes en JSON", http.StatusInternalServerError)
		return
	}
}

func GetPaqueteHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id_paquete_uint, err := strconv.ParseUint(params["id"], 10, 32)

	paquete, err := dataPaquete.ObtenerPaquete(uint(id_paquete_uint))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("El paquete no existe: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&paquete)
}

func PostPaqueteHandler(w http.ResponseWriter, r *http.Request) {

	var paquetes []modelosPaquete.Paquete

	if err := json.NewDecoder(r.Body).Decode(&paquetes); err != nil {
		http.Error(w, "Error al decodificar los paquetes: "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, paquete := range paquetes {
		if err := validaciones.ValidarPaquete(paquete); err != nil {
			http.Error(w, "Datos del paquete inválidos: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err := dataPaquete.CrearPaquetes(paquetes); err != nil {
		http.Error(w, "Error al crear los paquetes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&paquetes)
}

func PutPaqueteHandler(w http.ResponseWriter, r *http.Request) {
	var paquetesInput []*modelosPaquete.Paquete

	if err := json.NewDecoder(r.Body).Decode(&paquetesInput); err != nil {
		http.Error(w, "Error al decodificar los paquetes: "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, paquete := range paquetesInput {
		if err := validaciones.ValidarPaquete(*paquete); err != nil {
			http.Error(w, "Datos del paquete inválidos: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err := dataPaquete.ActualizarPaquetes(paquetesInput); err != nil {
		http.Error(w, "Error al actualizar los paquetes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Paquetes actualizados"))
}

func DeletePaqueteHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id_paquete_str := params["id"]

	id_paquete, err := strconv.ParseUint(id_paquete_str, 10, 64)
	if err != nil {
		http.Error(w, "ID del paquete inválido", http.StatusBadRequest)
		return
	}

	if err := dataPaquete.BorrarPaquete(uint(id_paquete)); err != nil {
		http.Error(w, "Error al borrar el paquete: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Paquete borrado"))
	w.WriteHeader(http.StatusOK)
}

func GetPaquetesSinAsignar(w http.ResponseWriter, r *http.Request) {
	paquetes := dataPaquete.ObtenerPaquetesSinAsignar()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(paquetes); err != nil {
		http.Error(w, "Error al codificar paquetes en JSON", http.StatusInternalServerError)
		return
	}

}

func PutEntregarPaquete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id_paquete_str := params["id"]

	id_paquete, err := strconv.ParseUint(id_paquete_str, 10, 64)
	if err != nil {
		http.Error(w, "ID del paquete inválido", http.StatusBadRequest)
		return
	}

	if err := dataPaquete.EntregarPaquete(uint(id_paquete)); err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error al entregar el paquete: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(&paquete)

}

package infraestructure

import (
	"actividad/src/perfumes/application"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type EditPerfumeController struct {
	useCase application.EditPerfume
}

func NewEditPerfumeController(useCase application.EditPerfume) *EditPerfumeController {
	return &EditPerfumeController{useCase: useCase}
}

func (ep_c *EditPerfumeController) Execute(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de perfume inválido", http.StatusBadRequest)
		return
	}

	var p struct {
		Marca  string  `json:"marca"`
		Modelo string  `json:"modelo"`
		Precio float32 `json:"precio"`
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	err = ep_c.useCase.Execute(int32(id), p.Marca, p.Modelo, p.Precio)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar el perfume: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Perfume actualizado correctamente"))
}
package infraestructure

import (
	"actividad/src/tiendas/application"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type EditTiendaController struct {
	useCase application.EditTienda
}

func NewEditTiendaController(useCase application.EditTienda) *EditTiendaController {
	return &EditTiendaController{useCase: useCase}
}

func (ep_c *EditTiendaController) Execute(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de tienda inválido", http.StatusBadRequest)
		return
	}

	var p struct {
		Nombre  string  `json:"nombre"`
		Ubicacion string  `json:"ubicacion"`
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	err = ep_c.useCase.Execute(int32(id), p.Nombre, p.Ubicacion)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar la tienda: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tienda actualizado correctamente"))
}
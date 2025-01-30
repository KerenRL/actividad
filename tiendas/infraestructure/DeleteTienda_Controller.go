package infraestructure

import (
	"actividad/src/tiendas/application"
	"fmt"
	"net/http"
	"strconv"
)

type DeleteTiendaController struct {
	useCase application.DeleteTienda
}

func NewDeleteTiendaController(useCase application.DeleteTienda) *DeleteTiendaController {
	return &DeleteTiendaController{useCase: useCase}
}

func (dp_c *DeleteTiendaController) Execute(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de perfume inválido", http.StatusBadRequest)
		return
	}

	err = dp_c.useCase.Execute(int32(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar la tienda: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tienda eliminada correctamente"))
}

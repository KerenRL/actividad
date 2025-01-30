package infraestructure

import (
	"actividad/src/perfumes/application"
	"fmt"
	"net/http"
	"strconv"
)

type DeletePerfumeController struct {
	useCase application.DeletePerfume
}

func NewDeletePerfumeController(useCase application.DeletePerfume) *DeletePerfumeController {
	return &DeletePerfumeController{useCase: useCase}
}

func (dp_c *DeletePerfumeController) Execute(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de perfume inválido", http.StatusBadRequest)
		return
	}

	err = dp_c.useCase.Execute(int32(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar el perfume: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Perfume eliminado correctamente"))
}

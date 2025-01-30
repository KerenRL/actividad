package infraestructure

import (
	"actividad/src/tiendas/application"
	"net/http"
)

type ViewTiendaController struct {
	useCase application.ViewTienda
}

func NewViewTiendaController(useCase application.ViewTienda) *ViewTiendaController {
	return &ViewTiendaController{useCase: useCase}
}

func (vp_c *ViewTiendaController) Execute(w http.ResponseWriter, r *http.Request) {
	vp_c.useCase.Execute()
	w.Write([]byte("Lista de tiendas"))
}

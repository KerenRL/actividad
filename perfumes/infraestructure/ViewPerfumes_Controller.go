package infraestructure

import (
	"actividad/src/perfumes/application"
	"net/http"
)

type ViewPerfumesController struct {
	useCase application.ViewPerfumes
}

func NewViewPerfumesController(useCase application.ViewPerfumes) *ViewPerfumesController {
	return &ViewPerfumesController{useCase: useCase}
}

func (vp_c *ViewPerfumesController) Execute(w http.ResponseWriter, r *http.Request) {
	vp_c.useCase.Execute()
	w.Write([]byte("Lista de perfumes"))
}

package infraestructure

import (
	"actividad/src/perfumes/application"
	"actividad/src/perfumes/domain"
	"net/http"
)

func SetupRouter(repo domain.IPerfume) {
	createPerfume := application.NewCreatePerfume(repo)
	createPerfumeController := NewCreatePerfumeController(*createPerfume)

	viewPerfumes := application.NewViewPerfumes(repo)
	viewPerfumesController := NewViewPerfumesController(*viewPerfumes)

	editPerfumeUseCase := application.NewEditPerfume(repo)
	editPerfumeController := NewEditPerfumeController(*editPerfumeUseCase)

	deletePerfumeUseCase := application.NewDeletePerfume(repo)
	deletePerfumeController := NewDeletePerfumeController(*deletePerfumeUseCase)

	http.HandleFunc("/perfumes", createPerfumeController.Execute)
	http.HandleFunc("/perfume", viewPerfumesController.Execute)
	http.HandleFunc("/editPerfume", editPerfumeController.Execute)
	http.HandleFunc("/deletePerfume", deletePerfumeController.Execute)

	http.ListenAndServe(":8080", nil)
}

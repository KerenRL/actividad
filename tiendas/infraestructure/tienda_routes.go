package infraestructure

import (
	"actividad/src/tiendas/application"
	"actividad/src/tiendas/domain"
	"net/http"
)

func SetupRouter(repo domain.ITienda) {
	createTienda := application.NewCreatePerfume(repo)
	createTiendaController := NewCreateTiendaController(*createTienda)

	viewTienda := application.NewViewTienda(repo)
	viewTiendaController := NewViewTiendaController(*viewTienda)

	editTiendaUseCase := application.NewEditTienda(repo)
	editTiendaController := NewEditTiendaController(*editTiendaUseCase)

	deleteTiendaUseCase := application.NewDeleteTienda(repo)
	deleteTiendaController := NewDeleteTiendaController(*deleteTiendaUseCase)

	http.HandleFunc("/tiendas", createTiendaController.Execute)
	http.HandleFunc("/tienda", viewTiendaController.Execute)
	http.HandleFunc("/editTienda", editTiendaController.Execute)
	http.HandleFunc("/deleteTienda", deleteTiendaController.Execute)

	http.ListenAndServe(":8080", nil)
}

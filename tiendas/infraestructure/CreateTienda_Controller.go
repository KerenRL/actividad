package infraestructure

import (
	"actividad/src/tiendas/application"
	"encoding/json"
	"net/http"
)

type CreateTiendaController struct {
	useCase application.CreateTienda
}

func NewCreateTiendaController(useCase application.CreateTienda) *CreateTiendaController {
	return &CreateTiendaController{useCase: useCase}
}

type RequestBody struct {
	Nombre  string  `json:"nombre"`
	Ubicacion string  `json:"ubicacion"`
}

func (cp_c *CreateTiendaController) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Error al leer el JSON", http.StatusBadRequest)
		return
	}

	cp_c.useCase.Execute(body.Nombre, body.Ubicacion)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Perfume agregado correctamente"})
}

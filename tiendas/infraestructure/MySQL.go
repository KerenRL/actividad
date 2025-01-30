package infraestructure

import (
	"actividad/src/config"
	"actividad/src/tiendas/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.ITienda = (*MySQL)(nil)

func NewMySQL() domain.ITienda {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}
func (mysql *MySQL) SaveTienda(nombre string, ubicacion string) {
	query := "INSERT INTO tienda (nombre, ubicacion) VALUES ( ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, nombre, ubicacion)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Tienda guardada correctamente: Nombre: %s - Ubicacion: %.2f", nombre, ubicacion)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
}

func (mysql *MySQL) GetAll() {
	query := "SELECT * FROM tienda"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombre, ubicacion string
		if err := rows.Scan(&id, &nombre, &ubicacion); err != nil {
			fmt.Printf("Error al escanear la fila: %v\n", err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Ubicacion: %.2f\n", id, nombre, ubicacion)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterando sobre las filas: %v\n", err)
	}
}

func (mysql *MySQL) UpdateTienda(id int32, nombre string, ubicacion string) error {
	query := "UPDATE tienda SET nombre = ?, ubicacion = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, nombre, ubicacion, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de actualización: %v", err)
	}
	return nil
}

func (mysql *MySQL) DeleteTienda(id int32) error {
	query := "DELETE FROM tienda WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de eliminación: %v", err)
	}
	return nil
}
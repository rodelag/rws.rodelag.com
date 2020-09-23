package reportes_elconix

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
)

type InventarioSucursal struct {
	Bodega,
	Inventario string
}

func conexionInventarioSucursal() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.elconix.user"),
		viper.GetString("basedatos.mysql.elconix.password"),
		viper.GetString("basedatos.mysql.elconix.server"),
		viper.GetString("basedatos.mysql.elconix.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: ", errMySQL)
	}
	return connMySQL
}

func ListarInventarioSucursal() []InventarioSucursal {
	connMySQL := conexionInventarioSucursal()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consulta())
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	pi, psi := InventarioSucursal{}, []InventarioSucursal{}

	for rows.Next() {
		err := rows.Scan(&pi.Bodega, &pi.Inventario)
		utils.LogError("Problemas leer los datos: ", err)
		psi = append(psi, InventarioSucursal{
			Bodega:     pi.Bodega,
			Inventario: pi.Inventario,
		})
	}
	return psi
}

//Esta función más adelante puede cambiarse por un llamado a una API
func consulta() string {
	consulta := `
		SELECT
			Warehouse AS bodega,
			SUM(a.instock * Item_AvgCost_Calc(p.id, 'NACIONALIZADA', SYSDATE())) AS inventario
		FROM
			products_mview_instock_actualizado AS a
				INNER JOIN products AS p
					ON a.Item = p.id
			GROUP BY Warehouse
			ORDER BY Sucursal
	`
	return consulta
}

package conteo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
)

type Conteo struct {
	RegistroID,
	RegistroNumero,
	RegistroEmpresa,
	RegistroSucursal,
	RegistroSucursalNombre,
	RegistroEntradaAnt,
	RegistroEntrada,
	RegistroSalidaAnt,
	RegistroSalida,
	RegistroFacturasAnt,
	RegistroFacturas,
	RegistroTiquetePromedioAnt,
	RegistroTiquetePromedio,
	RegistroArticulosAnt,
	RegistroArticulos,
	RegistroVentaAnt,
	RegistroVenta,
	RegistroFecha,
	RegistroIP string
}

func conexionConteo() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		viper.GetString("basedatos.mysql.rodelag.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: ", errMySQL)
	}
	return connMySQL
}

func ListarConteo(inicio, fin string) []Conteo {
	connMySQL := conexionConteo()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(consulta(inicio, fin))
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	con, cons := Conteo{}, []Conteo{}

	for rows.Next() {
		err := rows.Scan(
			&con.RegistroID,
			&con.RegistroNumero,
			&con.RegistroEmpresa,
			&con.RegistroSucursal,
			&con.RegistroSucursalNombre,
			&con.RegistroEntradaAnt,
			&con.RegistroEntrada,
			&con.RegistroSalidaAnt,
			&con.RegistroSalida,
			&con.RegistroFacturasAnt,
			&con.RegistroFacturas,
			&con.RegistroTiquetePromedioAnt,
			&con.RegistroTiquetePromedio,
			&con.RegistroArticulosAnt,
			&con.RegistroArticulos,
			&con.RegistroVentaAnt,
			&con.RegistroVenta,
			&con.RegistroFecha,
			&con.RegistroIP,
		)
		utils.LogError("Problemas leer los datos: ", err)
		cons = append(cons, Conteo{
			RegistroID:                 con.RegistroID,
			RegistroNumero:             con.RegistroNumero,
			RegistroEmpresa:            con.RegistroEmpresa,
			RegistroSucursal:           con.RegistroSucursal,
			RegistroSucursalNombre:     con.RegistroSucursalNombre,
			RegistroEntradaAnt:         con.RegistroEntradaAnt,
			RegistroEntrada:            con.RegistroEntrada,
			RegistroSalidaAnt:          con.RegistroSalidaAnt,
			RegistroSalida:             con.RegistroSalida,
			RegistroFacturasAnt:        con.RegistroFacturasAnt,
			RegistroFacturas:           con.RegistroFacturas,
			RegistroTiquetePromedioAnt: con.RegistroTiquetePromedioAnt,
			RegistroTiquetePromedio:    con.RegistroTiquetePromedio,
			RegistroArticulosAnt:       con.RegistroArticulosAnt,
			RegistroArticulos:          con.RegistroArticulos,
			RegistroVentaAnt:           con.RegistroVentaAnt,
			RegistroVenta:              con.RegistroVenta,
			RegistroFecha:              con.RegistroFecha,
			RegistroIP:                 con.RegistroIP,
		})
	}
	return cons
}

//Esta función más adelante puede cambiarse por un llamado a una API
func consulta(inicio, fin string) string {
	consulta := `
		SELECT
			IFNULL(a.registroID, '') AS registroID,
			IFNULL(a.registroNumero, '') AS registroNumero,
			IFNULL(a.registroEmpresa, '') AS registroEmpresa,
			IFNULL(a.registroSucursal, '') AS registroSucursal,
			IFNULL(a.registroSucursalNombre, '') AS registroSucursalNombre,
			SUM(IFNULL(b.registroEntrada, 0)) AS registroEntradaAnt,
			SUM(IFNULL(a.registroEntrada, 0)) AS registroEntrada,
			IFNULL(b.registroSalida, '') AS registroSalidaAnt,
			IFNULL(a.registroSalida, '') AS registroSalida,
			IFNULL(b.registroFacturas, '') AS registroFacturasAnt,
			IFNULL(a.registroFacturas, '') AS registroFacturas,
			IFNULL(b.registroTiquetePromedio, '') AS registroTiquetePromedioAnt,
			IFNULL(a.registroTiquetePromedio, '') AS registroTiquetePromedio,
			IFNULL(b.registroArticulos, '') AS registroArticulosAnt,
			IFNULL(a.registroArticulos, '') AS registroArticulos,
			SUM(IFNULL(b.registroVenta, 0)) AS registroVentaAnt,
			SUM(IFNULL(a.registroVenta, 0)) AS registroVenta,
			IFNULL(a.registroFecha, '') AS registroFecha,
			IFNULL(a.registroIP, '') AS registroIP
		FROM
			rodelag_conteo.trafico AS a
			INNER JOIN
			(SELECT
				 IFNULL(b.registroID, '') AS registroID,
				 IFNULL(b.registroNumero, '') AS registroNumero,
				 IFNULL(b.registroEmpresa, '') AS registroEmpresa,
				 IFNULL(b.registroSucursal, '') AS registroSucursal,
				 IFNULL(b.registroSucursalNombre, '') AS registroSucursalNombre,
				 SUM(IFNULL(b.registroEntrada, 0)) AS registroEntrada,
				 SUM(IFNULL(b.registroSalida, 0)) AS registroSalida,
				 IFNULL(b.registroFacturas, '') AS registroFacturas,
				 IFNULL(b.registroTiquetePromedio, '') AS registroTiquetePromedio,
				 IFNULL(b.registroArticulos, '') AS registroArticulos,
				 SUM(IFNULL(b.registroVenta, 0)) AS registroVenta,
				 IFNULL(b.registroFecha, '') AS registroFecha,
				 IFNULL(b.registroIP, '') AS registroIP
			 FROM
				 rodelag_conteo.trafico AS b
			 WHERE
				DATE(b.registroFecha) BETWEEN DATE_ADD(DATE_ADD('%s', INTERVAL -1 YEAR), INTERVAL +1 DAY) AND DATE_ADD(DATE_ADD('%s', INTERVAL -1 YEAR), INTERVAL +1 DAY) GROUP BY b.registroSucursal) AS b
				ON a.registroSucursal = b.registroSucursal
		WHERE
			DATE(a.registroFecha) BETWEEN '%s' AND '%s'
			GROUP BY a.registroSucursal
	`
	return fmt.Sprintf(consulta, inicio, fin, inicio, fin)
}

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
	RegistroEntrada,
	RegistroSalida,
	RegistroFacturas,
	RegistroTiquetePromedio,
	RegistroArticulos,
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
			&con.RegistroEntrada,
			&con.RegistroSalida,
			&con.RegistroFacturas,
			&con.RegistroTiquetePromedio,
			&con.RegistroArticulos,
			&con.RegistroVenta,
			&con.RegistroFecha,
			&con.RegistroIP,
		)
		utils.LogError("Problemas leer los datos: ", err)
		cons = append(cons, Conteo{
			RegistroID:              con.RegistroID,
			RegistroNumero:          con.RegistroNumero,
			RegistroEmpresa:         con.RegistroEmpresa,
			RegistroSucursal:        con.RegistroSucursal,
			RegistroSucursalNombre:  con.RegistroSucursalNombre,
			RegistroEntrada:         con.RegistroEntrada,
			RegistroSalida:          con.RegistroSalida,
			RegistroFacturas:        con.RegistroFacturas,
			RegistroTiquetePromedio: con.RegistroTiquetePromedio,
			RegistroArticulos:       con.RegistroArticulos,
			RegistroVenta:           con.RegistroVenta,
			RegistroFecha:           con.RegistroFecha,
			RegistroIP:              con.RegistroIP,
		})
	}
	return cons
}

//Esta función más adelante puede cambiarse por un llamado a una API
func consulta(inicio, fin string) string {
	consulta := `
		SELECT
			IFNULL(registroID, '') AS registroID,
			IFNULL(registroNumero, '') AS registroNumero,
			IFNULL(registroEmpresa, '') AS registroEmpresa,
			IFNULL(registroSucursal, '') AS registroSucursal,
			IFNULL(registroSucursalNombre, '') AS registroSucursalNombre,
			IFNULL(registroEntrada, '') AS registroEntrada,
			IFNULL(registroSalida, '') AS registroSalida,
			IFNULL(registroFacturas, '') AS registroFacturas,
			IFNULL(registroTiquetePromedio, '') AS registroTiquetePromedio,
			IFNULL(registroArticulos, '') AS registroArticulos,
			IFNULL(registroVenta, '') AS registroVenta,
			IFNULL(registroFecha, '') AS registroFecha,
			IFNULL(registroIP, '') AS registroIP
		FROM
			rodelag_conteo.trafico
		WHERE
			DATE(registroFecha) BETWEEN '%s' AND '%s'
	`
	return fmt.Sprintf(consulta, inicio, fin)
}

package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type ContactenosVC struct {
	ID int
	Nombre,
	Apellido,
	Cedula,
	Correo,
	Telefono,
	ActividadEconomica,
	DetalleSolicitud,
	FechaRegistro string
}

func conexionContactenosVC() *sql.DB {
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

func ListarContactenosVC() []ContactenosVC {
	connMySQL := conexionContactenosVC()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_contactenosvc;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	ct := ContactenosVC{}
	cts := []ContactenosVC{}

	for rows.Next() {
		err := rows.Scan(
			&ct.ID,
			&ct.Nombre,
			&ct.Apellido,
			&ct.Cedula,
			&ct.Correo,
			&ct.Telefono,
			&ct.ActividadEconomica,
			&ct.DetalleSolicitud,
			&ct.FechaRegistro,
		)
		utils.LogError("Problemas leer los datos: ", err)
		cts = append(cts, ContactenosVC{
			ID:                 ct.ID,
			Nombre:             ct.Nombre,
			Apellido:           ct.Apellido,
			Cedula:             ct.Cedula,
			Correo:             ct.Correo,
			Telefono:           ct.Telefono,
			ActividadEconomica: ct.ActividadEconomica,
			DetalleSolicitud:   ct.DetalleSolicitud,
			FechaRegistro:      ct.FechaRegistro,
		})
	}
	return cts
}

func CrearContactenosVC(nombre, apellido, cedula, correo, telefono, actividadEconomica, detalleSolicitud string) ContactenosVC {
	ct := ContactenosVC{
		Nombre:             nombre,
		Apellido:           apellido,
		Cedula:             cedula,
		Correo:             correo,
		Telefono:           telefono,
		ActividadEconomica: actividadEconomica,
		DetalleSolicitud:   detalleSolicitud,
		FechaRegistro:      time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionContactenosVC()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_contactenosvc (nombre, apellido, cedula, correo, telefono, actividadEconomica, detalleSolicitud, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		ct.Nombre,
		ct.Apellido,
		ct.Cedula,
		ct.Correo,
		ct.Telefono,
		ct.ActividadEconomica,
		ct.DetalleSolicitud,
		ct.FechaRegistro,
	)

	return ct
}

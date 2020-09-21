package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type Contactenos struct {
	ID int
	Nombre,
	Apellido,
	Correo,
	Telefono,
	Mensaje,
	FechaRegistro string
}

func conexionContactenos() *sql.DB {
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

func ListarContactenos() []Contactenos {
	connMySQL := conexionContactenos()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_contactenos;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	ct := Contactenos{}
	cts := []Contactenos{}

	for rows.Next() {
		err := rows.Scan(
			&ct.ID,
			&ct.Nombre,
			&ct.Apellido,
			&ct.Correo,
			&ct.Telefono,
			&ct.Mensaje,
			&ct.FechaRegistro,
		)
		utils.LogError("Problemas leer los datos: ", err)
		cts = append(cts, Contactenos{
			ID:            ct.ID,
			Nombre:        ct.Nombre,
			Apellido:      ct.Apellido,
			Correo:        ct.Correo,
			Telefono:      ct.Telefono,
			Mensaje:       ct.Mensaje,
			FechaRegistro: ct.FechaRegistro,
		})
	}
	return cts
}

func CrearContactenos(nombre, apellido, correo, telefono, mensaje string) Contactenos {
	ct := Contactenos{
		Nombre:        nombre,
		Apellido:      apellido,
		Correo:        correo,
		Telefono:      telefono,
		Mensaje:       mensaje,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionContactenos()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_contactenos (nombre, apellido, correo, telefono, mensaje, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		ct.Nombre,
		ct.Apellido,
		ct.Correo,
		ct.Telefono,
		ct.Mensaje,
		ct.FechaRegistro,
	)

	return ct
}

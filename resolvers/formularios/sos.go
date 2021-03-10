package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type Sos struct {
	ID int
	Nombre,
	Apellido,
	Cedula,
	Correo,
	Telefono,
	FechaRegistro,
	Estado string
}

func conexionSos() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		"rodelag_ovnicom",
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (Sos) ", true, errMySQL)
	}
	return connMySQL
}

func VerSos(id int) Sos {
	connMySQL := conexionSos()
	defer connMySQL.Close()

	sos := Sos{}

	err := connMySQL.QueryRow("SELECT * FROM formulario_sos WHERE id = ?;", id).Scan(
		&sos.ID,
		&sos.Nombre,
		&sos.Apellido,
		&sos.Cedula,
		&sos.Correo,
		&sos.Telefono,
		&sos.FechaRegistro,
		&sos.Estado,
	)
	utils.LogError("Problemas al leer registro: (Sos) ", false, err)

	return sos
}

func ListarSos() []Sos {
	connMySQL := conexionSos()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_sos;")
	utils.LogError("Problemas al listar los registros de la base de datos: (Sos) ", true, err)
	defer rows.Close()

	sos := Sos{}
	soss := []Sos{}

	for rows.Next() {
		err := rows.Scan(&sos.ID, &sos.Nombre, &sos.Apellido, &sos.Cedula, &sos.Correo, &sos.Telefono, &sos.FechaRegistro, &sos.Estado)
		utils.LogError("Problemas leer los datos: (Sos) ", true, err)
		soss = append(soss, Sos{
			ID:            sos.ID,
			Nombre:        sos.Nombre,
			Apellido:      sos.Apellido,
			Cedula:        sos.Cedula,
			Correo:        sos.Correo,
			Telefono:      sos.Telefono,
			FechaRegistro: sos.FechaRegistro,
			Estado:        sos.Estado,
		})
	}
	return soss
}

func BusquedaSos(busqueda string) []Sos {
	connMySQL := conexionSos()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(
		fmt.Sprintf("SELECT * FROM formulario_sos WHERE (nombre LIKE '%%%s%%') OR (apellido LIKE '%%%s%%') OR (cedula LIKE '%%%s%%') OR (estado LIKE '%%%s%%');",
			busqueda,
			busqueda,
			busqueda,
			busqueda))

	utils.LogError("Problemas al listar los registros de la base de datos: (Sos) ", true, err)
	defer rows.Close()

	sos := Sos{}
	soss := []Sos{}

	for rows.Next() {
		err := rows.Scan(&sos.ID, &sos.Nombre, &sos.Apellido, &sos.Cedula, &sos.Correo, &sos.Telefono, &sos.FechaRegistro, &sos.Estado)
		utils.LogError("Problemas leer los datos: (Sos) ", true, err)
		soss = append(soss, Sos{
			ID:            sos.ID,
			Nombre:        sos.Nombre,
			Apellido:      sos.Apellido,
			Cedula:        sos.Cedula,
			Correo:        sos.Correo,
			Telefono:      sos.Telefono,
			FechaRegistro: sos.FechaRegistro,
			Estado:        sos.Estado,
		})
	}
	return soss
}

func EditarSos(id int, estado string) Sos {
	sos := Sos{
		ID:     id,
		Estado: estado,
	}

	connMySQL := conexionSos()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("UPDATE formulario_sos SET estado = ? WHERE id = ?;")
	utils.LogError("Problemas al crear el registro en la base de datos: (Sos) ", true, err)
	defer conn.Close()

	conn.Exec(sos.Estado, sos.ID)

	return sos
}

func CrearSos(nombre, apellido, cedula, correo, telefono string) Sos {
	sos := Sos{
		Nombre:        nombre,
		Apellido:      apellido,
		Cedula:        cedula,
		Correo:        correo,
		Telefono:      telefono,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionSos()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_sos (nombre, apellido, cedula, correo, telefono, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: (Sos) ", false, err)
	defer conn.Close()

	conn.Exec(sos.Nombre, sos.Apellido, sos.Cedula, sos.Correo, sos.Telefono, sos.FechaRegistro)

	return sos
}

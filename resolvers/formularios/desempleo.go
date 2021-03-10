package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type Desempleo struct {
	ID int
	Nombre,
	Apellido,
	Cedula,
	Edad,
	Correo,
	Telefono,
	DireccionDomicilio,
	NombreEmpresa,
	TiempoLaboral,
	FechaRegistro,
	Estado string
}

func conexionDesempleo() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		"rodelag_ovnicom",
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (Desempleo) ", true, errMySQL)
	}
	return connMySQL
}

func VerDesempleo(id int) Desempleo {
	connMySQL := conexionDesempleo()
	defer connMySQL.Close()

	des := Desempleo{}

	err := connMySQL.QueryRow("SELECT * FROM formulario_desempleo WHERE id = ?;", id).Scan(
		&des.ID,
		&des.Nombre,
		&des.Apellido,
		&des.Cedula,
		&des.Edad,
		&des.Correo,
		&des.Telefono,
		&des.DireccionDomicilio,
		&des.NombreEmpresa,
		&des.TiempoLaboral,
		&des.FechaRegistro,
		&des.Estado,
	)
	utils.LogError("Problemas al leer registro: (Desempleo) ", false, err)

	return des
}

func ListarDesempleo() []Desempleo {
	connMySQL := conexionDesempleo()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_desempleo;")
	utils.LogError("Problemas al listar los registros de la base de datos: (Desempleo) ", true, err)
	defer rows.Close()

	des := Desempleo{}
	dess := []Desempleo{}

	for rows.Next() {
		err := rows.Scan(&des.ID, &des.Nombre, &des.Apellido, &des.Cedula, &des.Edad, &des.Correo, &des.Telefono, &des.DireccionDomicilio, &des.NombreEmpresa, &des.TiempoLaboral, &des.FechaRegistro, &des.Estado)
		utils.LogError("Problemas leer los datos: (Desempleo) ", true, err)
		dess = append(dess, Desempleo{
			ID:                 des.ID,
			Nombre:             des.Nombre,
			Apellido:           des.Apellido,
			Cedula:             des.Cedula,
			Edad:               des.Edad,
			Correo:             des.Correo,
			Telefono:           des.Telefono,
			DireccionDomicilio: des.DireccionDomicilio,
			NombreEmpresa:      des.NombreEmpresa,
			TiempoLaboral:      des.TiempoLaboral,
			FechaRegistro:      des.FechaRegistro,
			Estado:             des.Estado,
		})
	}
	return dess
}

func BusquedaDesempleo(busqueda string) []Desempleo {
	connMySQL := conexionDesempleo()
	defer connMySQL.Close()

	rows, err := connMySQL.Query(
		fmt.Sprintf("SELECT * FROM formulario_desempleo WHERE (nombre LIKE '%%%s%%') OR (apellido LIKE '%%%s%%') OR (cedula LIKE '%%%s%%') OR (estado LIKE '%%%s%%');",
			busqueda,
			busqueda,
			busqueda,
			busqueda))

	utils.LogError("Problemas al listar los registros de la base de datos: (Desempleo) ", true, err)
	defer rows.Close()

	des := Desempleo{}
	dess := []Desempleo{}

	for rows.Next() {
		err := rows.Scan(&des.ID, &des.Nombre, &des.Apellido, &des.Cedula, &des.Edad, &des.Correo, &des.Telefono, &des.DireccionDomicilio, &des.NombreEmpresa, &des.TiempoLaboral, &des.FechaRegistro, &des.Estado)
		utils.LogError("Problemas leer los datos: (Desempleo) ", true, err)
		dess = append(dess, Desempleo{
			ID:                 des.ID,
			Nombre:             des.Nombre,
			Apellido:           des.Apellido,
			Cedula:             des.Cedula,
			Edad:               des.Edad,
			Correo:             des.Correo,
			Telefono:           des.Telefono,
			DireccionDomicilio: des.DireccionDomicilio,
			NombreEmpresa:      des.NombreEmpresa,
			TiempoLaboral:      des.TiempoLaboral,
			FechaRegistro:      des.FechaRegistro,
			Estado:             des.Estado,
		})
	}
	return dess
}

func EditarDesempleo(id int, estado string) Desempleo {
	cp := Desempleo{
		ID:     id,
		Estado: estado,
	}

	connMySQL := conexionDesempleo()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("UPDATE formulario_desempleo SET estado = ? WHERE id = ?;")
	utils.LogError("Problemas al crear el registro en la base de datos: (Desempleo) ", true, err)
	defer conn.Close()

	conn.Exec(cp.Estado, cp.ID)

	return cp
}

func CrearDesempleo(nombre, apellido, cedula, edad, correo, telefono, direccionDomicilio, nombreEmpresa, tiempoLaboral string) Desempleo {
	des := Desempleo{
		Nombre:             nombre,
		Apellido:           apellido,
		Cedula:             cedula,
		Edad:               edad,
		Correo:             correo,
		Telefono:           telefono,
		DireccionDomicilio: direccionDomicilio,
		NombreEmpresa:      nombreEmpresa,
		TiempoLaboral:      tiempoLaboral,
		FechaRegistro:      time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionDesempleo()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_desempleo (nombre, apellido, cedula, edad, correo, telefono, direccionDomicilio, nombreEmpresa, tiempoLaboral, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: (Desempleo) ", true, err)
	defer conn.Close()

	conn.Exec(des.Nombre, des.Apellido, des.Cedula, des.Edad, des.Correo, des.Telefono, des.DireccionDomicilio, des.NombreEmpresa, des.TiempoLaboral, des.FechaRegistro)

	return des
}

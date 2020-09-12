package resolvers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type PagoACH struct {
	ID              int
	Nombre          string
	Apellido        string
	TitularCuenta   string
	Cedula          string
	Correo          string
	Telefono        string
	CompraOrigen    string
	NumeroOrden     string
	FotoComprobante string
	Estado          string
	FechaRegistro   string
	Estados         []EstadoPagoACH
}

type EstadoPagoACH struct {
	ID,
	Estado,
	Comentario,
	FechaRegistro,
	Formulario,
	Usuario,
	CorreoUsuario string
	IDFormulario int
}

func configuracion() {
	viper.SetConfigName("configuracion")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func logError(m string, e error) {
	configuracion()
	f, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	if e != nil {
		log.Printf(m+" %v", e)
		notificacion(m, e)
	}
}

func conexion() *sql.DB {
	configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.user"),
		viper.GetString("basedatos.mysql.password"),
		viper.GetString("basedatos.mysql.server"),
		viper.GetString("basedatos.mysql.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		logError("Problemas con la conexion a mysql: ", errMySQL)
	}
	return connMySQL
}

func VerPagoACH(id int) PagoACH {
	connMySQL := conexion()
	defer connMySQL.Close()

	pagoACH := PagoACH{
		Estados: func() []EstadoPagoACH {
			consulta := fmt.Sprintf("SELECT * FROM formulario_estado WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_pagosach", id)

			rows, err := connMySQL.Query(consulta)
			logError("Problemas al listar los estados de los registros de la base de datos: ", err)
			defer rows.Close()

			estado, estados := EstadoPagoACH{}, []EstadoPagoACH{}

			for rows.Next() {
				err := rows.Scan(&estado.ID, &estado.Estado, &estado.Comentario, &estado.FechaRegistro, &estado.Formulario, &estado.Usuario, &estado.CorreoUsuario, &estado.IDFormulario)
				logError("Problemas leer los estados: ", err)
				estados = append(estados, EstadoPagoACH{
					ID:            estado.ID,
					Estado:        estado.Estado,
					Comentario:    estado.Comentario,
					FechaRegistro: estado.FechaRegistro,
					Formulario:    estado.Formulario,
					Usuario:       estado.Usuario,
					CorreoUsuario: estado.CorreoUsuario,
					IDFormulario:  estado.IDFormulario,
				})
			}
			return estados
		}(),
	}

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_estado WHERE formulario = 'formulario_pagosach' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_pagosach AS a WHERE a.id = ?;", id).Scan(
		&pagoACH.ID,
		&pagoACH.Nombre,
		&pagoACH.Apellido,
		&pagoACH.TitularCuenta,
		&pagoACH.Cedula,
		&pagoACH.Correo,
		&pagoACH.Telefono,
		&pagoACH.CompraOrigen,
		&pagoACH.NumeroOrden,
		&pagoACH.FotoComprobante,
		&pagoACH.FechaRegistro,
		&pagoACH.Estado,
	)
	logError("Problemas al leer registro: ", err)

	return pagoACH
}

func ListarPagoACH() []PagoACH {
	connMySQL := conexion()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_estado WHERE formulario = 'formulario_pagosach' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_pagosach AS a;")
	logError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	pagoACH := PagoACH{}
	pagosACH := []PagoACH{}

	for rows.Next() {
		err := rows.Scan(&pagoACH.ID, &pagoACH.Nombre, &pagoACH.Apellido, &pagoACH.TitularCuenta, &pagoACH.Cedula, &pagoACH.Correo, &pagoACH.Telefono, &pagoACH.CompraOrigen, &pagoACH.NumeroOrden, &pagoACH.FotoComprobante, &pagoACH.FechaRegistro, &pagoACH.Estado)
		logError("Problemas leer los datos: ", err)
		pagosACH = append(pagosACH, PagoACH{
			ID:              pagoACH.ID,
			Nombre:          pagoACH.Nombre,
			Apellido:        pagoACH.Apellido,
			TitularCuenta:   pagoACH.TitularCuenta,
			Cedula:          pagoACH.Cedula,
			Correo:          pagoACH.Correo,
			Telefono:        pagoACH.Telefono,
			CompraOrigen:    pagoACH.CompraOrigen,
			NumeroOrden:     pagoACH.NumeroOrden,
			FotoComprobante: pagoACH.FotoComprobante,
			Estado:          pagoACH.Estado,
			FechaRegistro:   pagoACH.FechaRegistro,
			Estados: func() []EstadoPagoACH {
				consulta := fmt.Sprintf("SELECT * FROM formulario_estado WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_pagosach", pagoACH.ID)

				rows, err := connMySQL.Query(consulta)
				logError("Problemas al listar los estados de los registros de la base de datos: ", err)
				defer rows.Close()

				estado, estados := EstadoPagoACH{}, []EstadoPagoACH{}

				for rows.Next() {
					err := rows.Scan(&estado.ID, &estado.Estado, &estado.Comentario, &estado.FechaRegistro, &estado.Formulario, &estado.Usuario, &estado.CorreoUsuario, &estado.IDFormulario)
					logError("Problemas leer los estados: ", err)
					estados = append(estados, EstadoPagoACH{
						ID:            estado.ID,
						Estado:        estado.Estado,
						Comentario:    estado.Comentario,
						FechaRegistro: estado.FechaRegistro,
						Formulario:    estado.Formulario,
						Usuario:       estado.Usuario,
						CorreoUsuario: estado.CorreoUsuario,
						IDFormulario:  estado.IDFormulario,
					})
				}
				return estados
			}(),
		})
	}
	return pagosACH
}

func CrearPagoACH(nombre string, apellido string, titularCuenta string, cedula string, correo string, telefono string, compraOrigen string, numeroOrden string, fotoComprobante string) PagoACH {
	pagosACH := PagoACH{
		Nombre:          nombre,
		Apellido:        apellido,
		TitularCuenta:   titularCuenta,
		Cedula:          cedula,
		Correo:          correo,
		Telefono:        telefono,
		CompraOrigen:    compraOrigen,
		NumeroOrden:     numeroOrden,
		FotoComprobante: fotoComprobante,
		FechaRegistro:   time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexion()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_pagosach (nombre, apellido, titularCuenta, cedula, correo, telefono, compraOrigen, numeroOrden, fotoComprobante, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(pagosACH.Nombre, pagosACH.Apellido, pagosACH.TitularCuenta, pagosACH.Cedula, pagosACH.Correo, pagosACH.Telefono, pagosACH.CompraOrigen, pagosACH.NumeroOrden, pagosACH.FotoComprobante, pagosACH.FechaRegistro)

	return pagosACH
}

func CrearEstadoPagoACH(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) EstadoPagoACH {
	estadoPagoACH := EstadoPagoACH{
		Estado:        estado,
		Comentario:    comentario,
		Formulario:    formulario,
		Usuario:       usuario,
		CorreoUsuario: correoUsuario,
		IDFormulario:  idFormulario,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexion()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_estado (estado, comentario, formulario, usuario, correoUsuario, idFormulario, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el estado en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(estadoPagoACH.Estado, estadoPagoACH.Comentario, estadoPagoACH.Formulario, estadoPagoACH.Usuario, estadoPagoACH.CorreoUsuario, estadoPagoACH.IDFormulario, estadoPagoACH.FechaRegistro)

	return estadoPagoACH
}

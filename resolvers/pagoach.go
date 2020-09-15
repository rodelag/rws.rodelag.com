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
	Comentarios     []ComentarioPagoACH
}

type ComentarioPagoACH struct {
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
		Comentarios: func() []ComentarioPagoACH {
			consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_pagosach", id)

			rows, err := connMySQL.Query(consulta)
			logError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
			defer rows.Close()

			comentario, comentarios := ComentarioPagoACH{}, []ComentarioPagoACH{}

			for rows.Next() {
				err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
				logError("Problemas leer los estados: ", err)
				comentarios = append(comentarios, ComentarioPagoACH{
					ID:            comentario.ID,
					Estado:        comentario.Estado,
					Comentario:    comentario.Comentario,
					FechaRegistro: comentario.FechaRegistro,
					Formulario:    comentario.Formulario,
					Usuario:       comentario.Usuario,
					CorreoUsuario: comentario.CorreoUsuario,
					IDFormulario:  comentario.IDFormulario,
				})
			}
			return comentarios
		}(),
	}

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_pagosach' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_pagosach AS a WHERE a.id = ?;", id).Scan(
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

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_pagosach' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_pagosach AS a;")
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
			Comentarios: func() []ComentarioPagoACH {
				consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_pagosach", pagoACH.ID)

				rows, err := connMySQL.Query(consulta)
				logError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
				defer rows.Close()

				comentario, comentarios := ComentarioPagoACH{}, []ComentarioPagoACH{}

				for rows.Next() {
					err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
					logError("Problemas leer los estados: ", err)
					comentarios = append(comentarios, ComentarioPagoACH{
						ID:            comentario.ID,
						Estado:        comentario.Estado,
						Comentario:    comentario.Comentario,
						FechaRegistro: comentario.FechaRegistro,
						Formulario:    comentario.Formulario,
						Usuario:       comentario.Usuario,
						CorreoUsuario: comentario.CorreoUsuario,
						IDFormulario:  comentario.IDFormulario,
					})
				}
				return comentarios
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

func CrearComentarioPagoACH(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) ComentarioPagoACH {
	comentarioPagoACH := ComentarioPagoACH{
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

	conn, err := connMySQL.Prepare("INSERT INTO formulario_comentarios (estado, comentario, formulario, usuario, correoUsuario, idFormulario, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el estado en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(comentarioPagoACH.Estado, comentarioPagoACH.Comentario, comentarioPagoACH.Formulario, comentarioPagoACH.Usuario, comentarioPagoACH.CorreoUsuario, comentarioPagoACH.IDFormulario, comentarioPagoACH.FechaRegistro)

	return comentarioPagoACH
}

package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type NotificacionDescuento struct {
	ID int
	Nombre,
	Apellido,
	Sucursal,
	Fecha,
	NombreProducto,
	FotoPrecio,
	CodigoParte,
	CantidadVendidas,
	NombreCompetidor,
	PrecioCompetidor,
	PrecioRodelag,
	FotoCotizacion,
	Estado,
	FechaRegistro string
	Comentarios []ComentarioNotificacionDescuento
}

type ComentarioNotificacionDescuento struct {
	ID,
	Estado,
	Comentario,
	FechaRegistro,
	Formulario,
	Usuario,
	CorreoUsuario string
	IDFormulario int
}

func conexionNotificacionDescuento() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		viper.GetString("basedatos.mysql.rodelag.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (Notificacion Descuento) ", true, errMySQL)
	}
	return connMySQL
}

func VerNotificacionDescuento(id int) NotificacionDescuento {
	connMySQL := conexionNotificacionDescuento()
	defer connMySQL.Close()

	reg := NotificacionDescuento{
		Comentarios: func() []ComentarioNotificacionDescuento {
			consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_notificaciondescuento", id)

			rows, err := connMySQL.Query(consulta)
			utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: (Notificacion Descuento) ", false, err)
			defer rows.Close()

			comentario, comentarios := ComentarioNotificacionDescuento{}, []ComentarioNotificacionDescuento{}

			for rows.Next() {
				err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
				utils.LogError("Problemas leer los estados: (Notificacion Descuento) ", false, err)
				comentarios = append(comentarios, ComentarioNotificacionDescuento{
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

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_notificaciondescuento' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_notificaciondescuento AS a WHERE a.id = ?;", id).Scan(
		&reg.ID,
		&reg.Nombre,
		&reg.Apellido,
		&reg.Sucursal,
		&reg.Fecha,
		&reg.NombreProducto,
		&reg.FotoPrecio,
		&reg.CodigoParte,
		&reg.CantidadVendidas,
		&reg.NombreCompetidor,
		&reg.PrecioCompetidor,
		&reg.PrecioRodelag,
		&reg.FotoCotizacion,
		&reg.FechaRegistro,
		&reg.Estado,
	)
	utils.LogError("Problemas al leer registro: (Notificacion Descuento) ", false, err)
	return reg
}

func ListarNotificacionDescuento() []NotificacionDescuento {
	connMySQL := conexionNotificacionDescuento()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_notificaciondescuento' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_notificaciondescuento AS a;")
	utils.LogError("Problemas al listar los registros de la base de datos: (Notificacion Descuento) ", true, err)
	defer rows.Close()

	nd := NotificacionDescuento{}
	nds := []NotificacionDescuento{}

	for rows.Next() {
		err := rows.Scan(
			&nd.ID,
			&nd.Nombre,
			&nd.Apellido,
			&nd.Sucursal,
			&nd.Fecha,
			&nd.NombreProducto,
			&nd.FotoPrecio,
			&nd.CodigoParte,
			&nd.CantidadVendidas,
			&nd.NombreCompetidor,
			&nd.PrecioCompetidor,
			&nd.PrecioRodelag,
			&nd.FotoCotizacion,
			&nd.FechaRegistro,
			&nd.Estado,
		)
		utils.LogError("Problemas leer los datos: (Notificacion Descuento) ", true, err)
		nds = append(nds, NotificacionDescuento{
			ID:               nd.ID,
			Nombre:           nd.Nombre,
			Apellido:         nd.Apellido,
			Sucursal:         nd.Sucursal,
			Fecha:            nd.Fecha,
			NombreProducto:   nd.NombreProducto,
			FotoPrecio:       nd.FotoPrecio,
			CodigoParte:      nd.CodigoParte,
			CantidadVendidas: nd.CantidadVendidas,
			NombreCompetidor: nd.NombreCompetidor,
			PrecioCompetidor: nd.PrecioCompetidor,
			PrecioRodelag:    nd.PrecioRodelag,
			FotoCotizacion:   nd.FotoCotizacion,
			FechaRegistro:    nd.FechaRegistro,
			Estado:           nd.Estado,
			Comentarios: func() []ComentarioNotificacionDescuento {
				consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_notificaciondescuento", nd.ID)

				rows, err := connMySQL.Query(consulta)
				utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: (Notificacion Descuento) ", false, err)
				defer rows.Close()

				comentario, comentarios := ComentarioNotificacionDescuento{}, []ComentarioNotificacionDescuento{}

				for rows.Next() {
					err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
					utils.LogError("Problemas leer los estados: (Notificacion Descuento) ", false, err)
					comentarios = append(comentarios, ComentarioNotificacionDescuento{
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
	return nds
}

func CrearNotificacionDescuento(nombre, apellido, sucursal, fecha, nombreProducto, fotoPrecio, codigoParte, cantidadVendidas, nombreCompetidor, precioCompetidor, precioRodelag, fotoCotizacion string) NotificacionDescuento {
	nd := NotificacionDescuento{
		Nombre:           nombre,
		Apellido:         apellido,
		Sucursal:         sucursal,
		Fecha:            fecha,
		NombreProducto:   nombreProducto,
		FotoPrecio:       fotoPrecio,
		CodigoParte:      codigoParte,
		CantidadVendidas: cantidadVendidas,
		NombreCompetidor: nombreCompetidor,
		PrecioCompetidor: precioCompetidor,
		PrecioRodelag:    precioRodelag,
		FotoCotizacion:   fotoCotizacion,
		FechaRegistro:    time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionNotificacionDescuento()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_notificaciondescuento (nombre, apellido, sucursal, fecha, nombreProducto, fotoPrecio, codigoParte, cantidadVendidas, nombreCompetidor, precioCompetidor, precioRodelag, fotoCotizacion, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: (Notificacion Descuento) ", true, err)
	defer conn.Close()

	conn.Exec(
		nd.Nombre,
		nd.Apellido,
		nd.Sucursal,
		nd.Fecha,
		nd.NombreProducto,
		nd.FotoPrecio,
		nd.CodigoParte,
		nd.CantidadVendidas,
		nd.NombreCompetidor,
		nd.PrecioCompetidor,
		nd.PrecioRodelag,
		nd.FotoCotizacion,
		nd.FechaRegistro,
	)

	return nd
}

func CrearComentarioNotificacionDescuento(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) ComentarioNotificacionDescuento {
	c := ComentarioNotificacionDescuento{
		Estado:        estado,
		Comentario:    comentario,
		Formulario:    formulario,
		Usuario:       usuario,
		CorreoUsuario: correoUsuario,
		IDFormulario:  idFormulario,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionNotificacionDescuento()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_comentarios (estado, comentario, formulario, usuario, correoUsuario, idFormulario, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el estado en la base de datos: (Notificacion Descuento) ", true, err)
	defer conn.Close()

	conn.Exec(c.Estado, c.Comentario, c.Formulario, c.Usuario, c.CorreoUsuario, c.IDFormulario, c.FechaRegistro)

	return c
}

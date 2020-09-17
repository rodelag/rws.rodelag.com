package resolvers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type SolicitudAplazamientoPago struct {
	ID int
	Nombre,
	Apellido,
	Correo,
	TelefonoCasa,
	Celular,
	TipoProducto,
	TipoCliente,
	TipoActividadEconomica,
	LugarTrabajo,
	MotivoSolicitud,
	DetalleMotivo,
	Cedula,
	Talonario,
	CartaMotivo,
	Gestion,
	EstadoCuenta,
	APC,
	Propuesta,
	Estado,
	FechaRegistro string
	Comentarios []ComentarioSolicitudAplazamientoPago
}

type ComentarioSolicitudAplazamientoPago struct {
	ID,
	Estado,
	Comentario,
	FechaRegistro,
	Formulario,
	Usuario,
	CorreoUsuario string
	IDFormulario int
}

func conexionSolicitudAplazamientoPago() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.user"),
		viper.GetString("basedatos.mysql.password"),
		viper.GetString("basedatos.mysql.server"),
		viper.GetString("basedatos.mysql.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: ", errMySQL)
	}
	return connMySQL
}

func VerSolicitudAplazamientoPago(id int) SolicitudAplazamientoPago {
	connMySQL := conexionSolicitudAplazamientoPago()
	defer connMySQL.Close()

	sap := SolicitudAplazamientoPago{
		Comentarios: func() []ComentarioSolicitudAplazamientoPago {
			consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_aplazamientopago", id)

			rows, err := connMySQL.Query(consulta)
			utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
			defer rows.Close()

			comentario, comentarios := ComentarioSolicitudAplazamientoPago{}, []ComentarioSolicitudAplazamientoPago{}

			for rows.Next() {
				err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
				utils.LogError("Problemas leer los estados: ", err)
				comentarios = append(comentarios, ComentarioSolicitudAplazamientoPago{
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

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_aplazamientopago' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_aplazamientopago AS a WHERE a.id = ?;", id).Scan(
		&sap.ID,
		&sap.Nombre,
		&sap.Apellido,
		&sap.Correo,
		&sap.TelefonoCasa,
		&sap.Celular,
		&sap.TipoProducto,
		&sap.TipoCliente,
		&sap.TipoActividadEconomica,
		&sap.LugarTrabajo,
		&sap.MotivoSolicitud,
		&sap.DetalleMotivo,
		&sap.Cedula,
		&sap.Talonario,
		&sap.CartaMotivo,
		&sap.Gestion,
		&sap.EstadoCuenta,
		&sap.APC,
		&sap.Propuesta,
		&sap.FechaRegistro,
		&sap.Estado,
	)
	utils.LogError("Problemas al leer registro: ", err)

	return sap
}

func ListarSolicitudAplazamientoPago() []SolicitudAplazamientoPago {
	connMySQL := conexionSolicitudAplazamientoPago()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_aplazamientopago' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_aplazamientopago AS a;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	sap := SolicitudAplazamientoPago{}
	saps := []SolicitudAplazamientoPago{}

	for rows.Next() {
		err := rows.Scan(
			&sap.ID,
			&sap.Nombre,
			&sap.Apellido,
			&sap.Correo,
			&sap.TelefonoCasa,
			&sap.Celular,
			&sap.TipoProducto,
			&sap.TipoCliente,
			&sap.TipoActividadEconomica,
			&sap.LugarTrabajo,
			&sap.MotivoSolicitud,
			&sap.DetalleMotivo,
			&sap.Cedula,
			&sap.Talonario,
			&sap.CartaMotivo,
			&sap.Gestion,
			&sap.EstadoCuenta,
			&sap.APC,
			&sap.Propuesta,
			&sap.FechaRegistro,
			&sap.Estado,
		)
		utils.LogError("Problemas leer los datos: ", err)
		saps = append(saps, SolicitudAplazamientoPago{
			ID:                     sap.ID,
			Nombre:                 sap.Nombre,
			Apellido:               sap.Apellido,
			Correo:                 sap.Correo,
			TelefonoCasa:           sap.TelefonoCasa,
			Celular:                sap.Celular,
			TipoProducto:           sap.TipoProducto,
			TipoCliente:            sap.TipoCliente,
			TipoActividadEconomica: sap.TipoActividadEconomica,
			LugarTrabajo:           sap.LugarTrabajo,
			MotivoSolicitud:        sap.MotivoSolicitud,
			DetalleMotivo:          sap.DetalleMotivo,
			Cedula:                 sap.Cedula,
			Talonario:              sap.Talonario,
			CartaMotivo:            sap.CartaMotivo,
			Estado:                 sap.Estado,
			Gestion:                sap.Gestion,
			EstadoCuenta:           sap.EstadoCuenta,
			APC:                    sap.APC,
			Propuesta:              sap.Propuesta,
			FechaRegistro:          sap.FechaRegistro,
			Comentarios: func() []ComentarioSolicitudAplazamientoPago {
				consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_aplazamientopago", sap.ID)

				rows, err := connMySQL.Query(consulta)
				utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
				defer rows.Close()

				comentario, comentarios := ComentarioSolicitudAplazamientoPago{}, []ComentarioSolicitudAplazamientoPago{}

				for rows.Next() {
					err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
					utils.LogError("Problemas leer los estados: ", err)
					comentarios = append(comentarios, ComentarioSolicitudAplazamientoPago{
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
	return saps
}

func CrearSolicitudAplazamientoPago(nombre, apellido, correo, telefonoCasa, celular, tipoProducto, tipoCliente, tipoActividadEconomica, lugarTrabajo, motivoSolicitud, detalleMotivo, cedula, talonario, cartaMotivo string) SolicitudAplazamientoPago {
	sap := SolicitudAplazamientoPago{
		Nombre:                 nombre,
		Apellido:               apellido,
		Correo:                 correo,
		TelefonoCasa:           telefonoCasa,
		Celular:                celular,
		TipoProducto:           tipoProducto,
		TipoCliente:            tipoCliente,
		TipoActividadEconomica: tipoActividadEconomica,
		LugarTrabajo:           lugarTrabajo,
		MotivoSolicitud:        motivoSolicitud,
		DetalleMotivo:          detalleMotivo,
		Cedula:                 cedula,
		Talonario:              talonario,
		CartaMotivo:            cartaMotivo,
		FechaRegistro:          time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionSolicitudAplazamientoPago()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_aplazamientopago (nombre, apellido, correo, telefonoCasa, celular, tipoProducto, tipoCliente, tipoActividadEconomica, lugarTrabajo, motivoSolicitud, detalleMotivo, cedula, talonario, cartaMotivo, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		sap.Nombre,
		sap.Apellido,
		sap.Correo,
		sap.TelefonoCasa,
		sap.Celular,
		sap.TipoProducto,
		sap.TipoCliente,
		sap.TipoActividadEconomica,
		sap.LugarTrabajo,
		sap.MotivoSolicitud,
		sap.DetalleMotivo,
		sap.Cedula,
		sap.Talonario,
		sap.CartaMotivo,
		sap.FechaRegistro,
	)

	return sap
}

func CrearComentarioSolicitudAplazamientoPago(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) ComentarioSolicitudAplazamientoPago {
	c := ComentarioSolicitudAplazamientoPago{
		Estado:        estado,
		Comentario:    comentario,
		Formulario:    formulario,
		Usuario:       usuario,
		CorreoUsuario: correoUsuario,
		IDFormulario:  idFormulario,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionSolicitudAplazamientoPago()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_comentarios (estado, comentario, formulario, usuario, correoUsuario, idFormulario, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el estado en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(c.Estado, c.Comentario, c.Formulario, c.Usuario, c.CorreoUsuario, c.IDFormulario, c.FechaRegistro)

	return c
}

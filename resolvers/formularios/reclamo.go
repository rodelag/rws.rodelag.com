package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type Reclamo struct {
	ID int
	Nombre,
	Apellido,
	Cedula,
	Correo,
	Telefono,
	TipoReclamo,
	Detalle,
	AdjuntoDocumento,
	Estado,
	FechaRegistro string
	Comentarios []ComentarioReclamo
}

type ComentarioReclamo struct {
	ID,
	Estado,
	Comentario,
	FechaRegistro,
	Formulario,
	Usuario,
	CorreoUsuario string
	IDFormulario int
}

func conexionReclamo() *sql.DB {
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

func VerReclamo(id int) Reclamo {
	connMySQL := conexionReclamo()
	defer connMySQL.Close()

	reg := Reclamo{
		Comentarios: func() []ComentarioReclamo {
			consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_reclamo", id)

			rows, err := connMySQL.Query(consulta)
			utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
			defer rows.Close()

			comentario, comentarios := ComentarioReclamo{}, []ComentarioReclamo{}

			for rows.Next() {
				err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
				utils.LogError("Problemas leer los estados: ", err)
				comentarios = append(comentarios, ComentarioReclamo{
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

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_reclamo' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_reclamo AS a WHERE a.id = ?;", id).Scan(
		&reg.ID,
		&reg.Nombre,
		&reg.Apellido,
		&reg.Cedula,
		&reg.Correo,
		&reg.Telefono,
		&reg.TipoReclamo,
		&reg.Detalle,
		&reg.AdjuntoDocumento,
		&reg.FechaRegistro,
		&reg.Estado,
	)
	utils.LogError("Problemas al leer registro: ", err)
	return reg
}

func ListarReclamo() []Reclamo {
	connMySQL := conexionReclamo()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_reclamo' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_reclamo AS a;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	re := Reclamo{}
	res := []Reclamo{}

	for rows.Next() {
		err := rows.Scan(
			&re.ID,
			&re.Nombre,
			&re.Apellido,
			&re.Cedula,
			&re.Correo,
			&re.Telefono,
			&re.TipoReclamo,
			&re.Detalle,
			&re.AdjuntoDocumento,
			&re.FechaRegistro,
			&re.Estado,
		)
		utils.LogError("Problemas leer los datos: ", err)
		res = append(res, Reclamo{
			ID:               re.ID,
			Nombre:           re.Nombre,
			Apellido:         re.Apellido,
			Cedula:           re.Cedula,
			Correo:           re.Correo,
			Telefono:         re.Telefono,
			TipoReclamo:      re.TipoReclamo,
			Detalle:          re.Detalle,
			AdjuntoDocumento: re.AdjuntoDocumento,
			FechaRegistro:    re.FechaRegistro,
			Estado:           re.Estado,
			Comentarios: func() []ComentarioReclamo {
				consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_reclamo", re.ID)

				rows, err := connMySQL.Query(consulta)
				utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
				defer rows.Close()

				comentario, comentarios := ComentarioReclamo{}, []ComentarioReclamo{}

				for rows.Next() {
					err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
					utils.LogError("Problemas leer los estados: ", err)
					comentarios = append(comentarios, ComentarioReclamo{
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
	return res
}

func CrearReclamo(nombre, apellido, cedula, correo, telefono, tipoReclamo, detalle, adjuntoDocumento string) Reclamo {
	re := Reclamo{
		Nombre:           nombre,
		Apellido:         apellido,
		Cedula:           cedula,
		Correo:           correo,
		Telefono:         telefono,
		TipoReclamo:      tipoReclamo,
		Detalle:          detalle,
		AdjuntoDocumento: adjuntoDocumento,
		FechaRegistro:    time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionReclamo()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_reclamo (nombre, apellido, cedula, correo, telefono, tipoReclamo, detalle, adjuntoDocumento, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(
		re.Nombre,
		re.Apellido,
		re.Cedula,
		re.Correo,
		re.Telefono,
		re.TipoReclamo,
		re.Detalle,
		re.AdjuntoDocumento,
		re.FechaRegistro,
	)

	return re
}

func CrearComentarioReclamo(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) ComentarioReclamo {
	c := ComentarioReclamo{
		Estado:        estado,
		Comentario:    comentario,
		Formulario:    formulario,
		Usuario:       usuario,
		CorreoUsuario: correoUsuario,
		IDFormulario:  idFormulario,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionEsiaa()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_comentarios (estado, comentario, formulario, usuario, correoUsuario, idFormulario, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el estado en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(c.Estado, c.Comentario, c.Formulario, c.Usuario, c.CorreoUsuario, c.IDFormulario, c.FechaRegistro)

	return c
}

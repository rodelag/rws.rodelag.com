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
	Estado,
	FechaRegistro string
	Comentarios []ComentarioContactenosVC
}

type ComentarioContactenosVC struct {
	ID,
	Estado,
	Comentario,
	FechaRegistro,
	Formulario,
	Usuario,
	CorreoUsuario string
	IDFormulario int
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
		utils.LogError("Problemas con la conexion a mysql: (Contactenosvc) ", true, errMySQL)
	}
	return connMySQL
}

func VerContactenosVC(id int) ContactenosVC {
	connMySQL := conexionContactenosVC()
	defer connMySQL.Close()

	reg := ContactenosVC{
		Comentarios: func() []ComentarioContactenosVC {
			consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_contactenosvc", id)

			rows, err := connMySQL.Query(consulta)
			utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: (Contactenosvc) ", true, err)
			defer rows.Close()

			comentario, comentarios := ComentarioContactenosVC{}, []ComentarioContactenosVC{}

			for rows.Next() {
				err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
				utils.LogError("Problemas leer los estados: (Contactenosvc) ", false, err)
				comentarios = append(comentarios, ComentarioContactenosVC{
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

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_contactenosvc' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_contactenosvc AS a WHERE a.id = ?;", id).Scan(
		&reg.ID,
		&reg.Nombre,
		&reg.Apellido,
		&reg.Cedula,
		&reg.Correo,
		&reg.Telefono,
		&reg.ActividadEconomica,
		&reg.DetalleSolicitud,
		&reg.FechaRegistro,
		&reg.Estado,
	)
	utils.LogError("Problemas al leer registro: (Contactenosvc) ", false, err)
	return reg
}

func ListarContactenosVC() []ContactenosVC {
	connMySQL := conexionContactenosVC()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_contactenosvc' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_contactenosvc AS a;")
	utils.LogError("Problemas al listar los registros de la base de datos: (Contactenosvc) ", true, err)
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
			&ct.Estado,
		)
		utils.LogError("Problemas leer los datos: (Contactenosvc) ", true, err)
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
			Estado:             ct.Estado,
			Comentarios: func() []ComentarioContactenosVC {
				consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_contactenosvc", ct.ID)

				rows, err := connMySQL.Query(consulta)
				utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: (Contactenosvc) ", true, err)
				defer rows.Close()

				comentario, comentarios := ComentarioContactenosVC{}, []ComentarioContactenosVC{}

				for rows.Next() {
					err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
					utils.LogError("Problemas leer los estados: (Contactenosvc) ", true, err)
					comentarios = append(comentarios, ComentarioContactenosVC{
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
	utils.LogError("Problemas al crear el registro en la base de datos: (Contactenosvc) ", true, err)
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

func CrearComentarioContactenosVC(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) ComentarioContactenosVC {
	c := ComentarioContactenosVC{
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
	utils.LogError("Problemas al crear el estado en la base de datos: (Contactenosvc) ", true, err)
	defer conn.Close()

	conn.Exec(c.Estado, c.Comentario, c.Formulario, c.Usuario, c.CorreoUsuario, c.IDFormulario, c.FechaRegistro)

	return c
}

package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type SolicitudEstadoCuenta struct {
	ID            int
	Nombre        string
	Apellido      string
	Correo        string
	Telefono      string
	Cedula        string
	Estado        string
	FechaRegistro string
	Comentarios   []ComentarioSolicitudEstadoCuenta
}

type ComentarioSolicitudEstadoCuenta struct {
	ID,
	Estado,
	Comentario,
	FechaRegistro,
	Formulario,
	Usuario,
	CorreoUsuario string
	IDFormulario int
}

func conexionSolicitudEstadoCuenta() *sql.DB {
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

func VerSolicitudEstadoCuenta(id int) SolicitudEstadoCuenta {
	connMySQL := conexionSolicitudEstadoCuenta()
	defer connMySQL.Close()

	sec := SolicitudEstadoCuenta{
		Comentarios: func() []ComentarioSolicitudEstadoCuenta {
			consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_estadocuenta", id)

			rows, err := connMySQL.Query(consulta)
			utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
			defer rows.Close()

			comentario, comentarios := ComentarioSolicitudEstadoCuenta{}, []ComentarioSolicitudEstadoCuenta{}

			for rows.Next() {
				err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
				utils.LogError("Problemas leer los estados: ", err)
				comentarios = append(comentarios, ComentarioSolicitudEstadoCuenta{
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

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_estadocuenta' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_estadocuenta AS a WHERE a.id = ?;", id).Scan(
		&sec.ID,
		&sec.Nombre,
		&sec.Apellido,
		&sec.Correo,
		&sec.Telefono,
		&sec.Cedula,
		&sec.FechaRegistro,
		&sec.Estado,
	)
	utils.LogError("Problemas al leer registro: ", err)

	return sec
}

func ListarEstadoCuenta() []SolicitudEstadoCuenta {
	connMySQL := conexionSolicitudEstadoCuenta()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_estadocuenta' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_estadocuenta AS a;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	sec := SolicitudEstadoCuenta{}
	ssec := []SolicitudEstadoCuenta{}

	for rows.Next() {
		err := rows.Scan(&sec.ID, &sec.Nombre, &sec.Apellido, &sec.Correo, &sec.Telefono, &sec.Cedula, &sec.FechaRegistro, &sec.Estado)
		utils.LogError("Problemas leer los datos: ", err)
		ssec = append(ssec, SolicitudEstadoCuenta{
			ID:            sec.ID,
			Nombre:        sec.Nombre,
			Apellido:      sec.Apellido,
			Correo:        sec.Correo,
			Telefono:      sec.Telefono,
			Cedula:        sec.Cedula,
			Estado:        sec.Estado,
			FechaRegistro: sec.FechaRegistro,
			Comentarios: func() []ComentarioSolicitudEstadoCuenta {
				consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_estadocuenta", sec.ID)

				rows, err := connMySQL.Query(consulta)
				utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
				defer rows.Close()

				comentario, comentarios := ComentarioSolicitudEstadoCuenta{}, []ComentarioSolicitudEstadoCuenta{}

				for rows.Next() {
					err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
					utils.LogError("Problemas leer los estados: ", err)
					comentarios = append(comentarios, ComentarioSolicitudEstadoCuenta{
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
	return ssec
}

func CrearSolicitudEstadoCuenta(nombre, apellido, correo, telefono, cedula string) SolicitudEstadoCuenta {
	sec := SolicitudEstadoCuenta{
		Nombre:        nombre,
		Apellido:      apellido,
		Correo:        correo,
		Telefono:      telefono,
		Cedula:        cedula,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionSolicitudEstadoCuenta()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_estadocuenta (nombre, apellido, correo, telefono, cedula, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(sec.Nombre, sec.Apellido, sec.Correo, sec.Telefono, sec.Cedula, sec.FechaRegistro)

	return sec
}

func CrearComentarioSolicitudEstadoCuenta(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) ComentarioSolicitudEstadoCuenta {
	c := ComentarioSolicitudEstadoCuenta{
		Estado:        estado,
		Comentario:    comentario,
		Formulario:    formulario,
		Usuario:       usuario,
		CorreoUsuario: correoUsuario,
		IDFormulario:  idFormulario,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionSolicitudEstadoCuenta()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_comentarios (estado, comentario, formulario, usuario, correoUsuario, idFormulario, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el estado en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(c.Estado, c.Comentario, c.Formulario, c.Usuario, c.CorreoUsuario, c.IDFormulario, c.FechaRegistro)

	return c
}

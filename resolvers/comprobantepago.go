package resolvers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type ComprobantePago struct {
	ID              int
	Nombre          string
	Apellido        string
	Cedula          string
	Correo          string
	Telefono        string
	ComprobantePago string
	Estado          string
	FechaRegistro   string
	Comentarios     []ComentarioComprobantePago
}

type ComentarioComprobantePago struct {
	ID,
	Estado,
	Comentario,
	FechaRegistro,
	Formulario,
	Usuario,
	CorreoUsuario string
	IDFormulario int
}

func conexionComprobantePago() *sql.DB {
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

func VerComprobantePago(id int) ComprobantePago {
	connMySQL := conexionComprobantePago()
	defer connMySQL.Close()

	comprobantePago := ComprobantePago{
		Comentarios: func() []ComentarioComprobantePago {
			consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_comprobantepago", id)

			rows, err := connMySQL.Query(consulta)
			utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
			defer rows.Close()

			comentario, comentarios := ComentarioComprobantePago{}, []ComentarioComprobantePago{}

			for rows.Next() {
				err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
				utils.LogError("Problemas leer los estados: ", err)
				comentarios = append(comentarios, ComentarioComprobantePago{
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

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_comprobantepago' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_comprobantepago AS a WHERE a.id = ?;", id).Scan(
		&comprobantePago.ID,
		&comprobantePago.Nombre,
		&comprobantePago.Apellido,
		&comprobantePago.Cedula,
		&comprobantePago.Correo,
		&comprobantePago.Telefono,
		&comprobantePago.ComprobantePago,
		&comprobantePago.FechaRegistro,
		&comprobantePago.Estado,
	)
	utils.LogError("Problemas al leer registro: ", err)

	return comprobantePago
}

func ListarComprobantePago() []ComprobantePago {
	connMySQL := conexionComprobantePago()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_comprobantepago' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_comprobantepago AS a;")
	utils.LogError("Problemas al listar los registros de la base de datos: ", err)
	defer rows.Close()

	cp := ComprobantePago{}
	cps := []ComprobantePago{}

	for rows.Next() {
		err := rows.Scan(&cp.ID, &cp.Nombre, &cp.Apellido, &cp.Cedula, &cp.Correo, &cp.Telefono, &cp.ComprobantePago, &cp.FechaRegistro, &cp.Estado)
		utils.LogError("Problemas leer los datos: ", err)
		cps = append(cps, ComprobantePago{
			ID:              cp.ID,
			Nombre:          cp.Nombre,
			Apellido:        cp.Apellido,
			Cedula:          cp.Cedula,
			Correo:          cp.Correo,
			Telefono:        cp.Telefono,
			ComprobantePago: cp.ComprobantePago,
			Estado:          cp.Estado,
			FechaRegistro:   cp.FechaRegistro,
			Comentarios: func() []ComentarioComprobantePago {
				consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_comprobantepago", cp.ID)

				rows, err := connMySQL.Query(consulta)
				utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: ", err)
				defer rows.Close()

				comentario, comentarios := ComentarioComprobantePago{}, []ComentarioComprobantePago{}

				for rows.Next() {
					err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
					utils.LogError("Problemas leer los estados: ", err)
					comentarios = append(comentarios, ComentarioComprobantePago{
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
	return cps
}

func CrearComprobantePago(nombre, apellido, cedula, correo, telefono, comprobantePago string) ComprobantePago {
	cp := ComprobantePago{
		Nombre:          nombre,
		Apellido:        apellido,
		Cedula:          cedula,
		Correo:          correo,
		Telefono:        telefono,
		ComprobantePago: comprobantePago,
		FechaRegistro:   time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionComprobantePago()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_comprobantepago (nombre, apellido, cedula, correo, telefono, comprobantePago, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(cp.Nombre, cp.Apellido, cp.Cedula, cp.Correo, cp.Telefono, cp.ComprobantePago, cp.FechaRegistro)

	return cp
}

func CrearComentarioComprobantePago(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) ComentarioComprobantePago {
	comprobantePago := ComentarioComprobantePago{
		Estado:        estado,
		Comentario:    comentario,
		Formulario:    formulario,
		Usuario:       usuario,
		CorreoUsuario: correoUsuario,
		IDFormulario:  idFormulario,
		FechaRegistro: time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionComprobantePago()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_comentarios (estado, comentario, formulario, usuario, correoUsuario, idFormulario, fechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el estado en la base de datos: ", err)
	defer conn.Close()

	conn.Exec(comprobantePago.Estado, comprobantePago.Comentario, comprobantePago.Formulario, comprobantePago.Usuario, comprobantePago.CorreoUsuario, comprobantePago.IDFormulario, comprobantePago.FechaRegistro)

	return comprobantePago
}

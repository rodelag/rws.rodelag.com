package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type Esiaa struct {
	ID int
	Nombre,
	Apellido,
	Cedula,
	Correo,
	Calificacion,
	Atencion,
	ResolverInstalacion,
	TiempoRazonable,
	Recomendacion,
	CalificacionManera,
	Estado,
	FechaRegistro string
	Comentarios []ComentarioEsiaa
}

type ComentarioEsiaa struct {
	ID,
	Estado,
	Comentario,
	FechaRegistro,
	Formulario,
	Usuario,
	CorreoUsuario string
	IDFormulario int
}

func conexionEsiaa() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		viper.GetString("basedatos.mysql.rodelag.database"),
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (Esiaa) ", true, errMySQL)
	}
	return connMySQL
}

func VerEsiaa(id int) Esiaa {
	connMySQL := conexionEsiaa()
	defer connMySQL.Close()

	reg := Esiaa{
		Comentarios: func() []ComentarioEsiaa {
			consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_esiaa", id)

			rows, err := connMySQL.Query(consulta)
			utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: (Esiaa) ", false, err)
			defer rows.Close()

			comentario, comentarios := ComentarioEsiaa{}, []ComentarioEsiaa{}

			for rows.Next() {
				err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
				utils.LogError("Problemas leer los estados: (Esiaa) ", false, err)
				comentarios = append(comentarios, ComentarioEsiaa{
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

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_esiaa' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_esiaa AS a WHERE a.id = ?;", id).Scan(
		&reg.ID,
		&reg.Nombre,
		&reg.Apellido,
		&reg.Cedula,
		&reg.Correo,
		&reg.Calificacion,
		&reg.Atencion,
		&reg.ResolverInstalacion,
		&reg.TiempoRazonable,
		&reg.Recomendacion,
		&reg.CalificacionManera,
		&reg.FechaRegistro,
		&reg.Estado,
	)
	utils.LogError("Problemas al leer registro: (Esiaa) ", false, err)
	return reg
}

func ListarEsiaa() []Esiaa {
	connMySQL := conexionEsiaa()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_esiaa' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_esiaa AS a;")
	utils.LogError("Problemas al listar los registros de la base de datos: (Esiaa) ", true, err)
	defer rows.Close()

	ct := Esiaa{}
	cts := []Esiaa{}

	for rows.Next() {
		err := rows.Scan(
			&ct.ID,
			&ct.Nombre,
			&ct.Apellido,
			&ct.Cedula,
			&ct.Correo,
			&ct.Calificacion,
			&ct.Atencion,
			&ct.ResolverInstalacion,
			&ct.TiempoRazonable,
			&ct.Recomendacion,
			&ct.CalificacionManera,
			&ct.FechaRegistro,
			&ct.Estado,
		)
		utils.LogError("Problemas leer los datos: (Esiaa) ", true, err)
		cts = append(cts, Esiaa{
			ID:                  ct.ID,
			Nombre:              ct.Nombre,
			Apellido:            ct.Apellido,
			Cedula:              ct.Cedula,
			Correo:              ct.Correo,
			Calificacion:        ct.Calificacion,
			Atencion:            ct.Atencion,
			ResolverInstalacion: ct.ResolverInstalacion,
			TiempoRazonable:     ct.TiempoRazonable,
			Recomendacion:       ct.Recomendacion,
			CalificacionManera:  ct.CalificacionManera,
			FechaRegistro:       ct.FechaRegistro,
			Estado:              ct.Estado,
			Comentarios: func() []ComentarioEsiaa {
				consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_esiaa", ct.ID)

				rows, err := connMySQL.Query(consulta)
				utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: (Esiaa) ", true, err)
				defer rows.Close()

				comentario, comentarios := ComentarioEsiaa{}, []ComentarioEsiaa{}

				for rows.Next() {
					err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
					utils.LogError("Problemas leer los estados: (Esiaa) ", true, err)
					comentarios = append(comentarios, ComentarioEsiaa{
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

func CrearEsiaa(nombre, apellido, cedula, correo, calificacion, atencion, resolverInstalacion, tiempoRazonable, recomendacion, calificacionManera string) Esiaa {
	ct := Esiaa{
		Nombre:              nombre,
		Apellido:            apellido,
		Cedula:              cedula,
		Correo:              correo,
		Calificacion:        calificacion,
		Atencion:            atencion,
		ResolverInstalacion: resolverInstalacion,
		TiempoRazonable:     tiempoRazonable,
		Recomendacion:       recomendacion,
		CalificacionManera:  calificacionManera,
		FechaRegistro:       time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionEsiaa()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_esiaa (nombre, apellido, cedula, correo, calificacion, atencion, resolverInstalacion, tiempoRazonable, recomendacion, calificacionManera, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: (Esiaa) ", false, err)
	defer conn.Close()

	conn.Exec(
		ct.Nombre,
		ct.Apellido,
		ct.Cedula,
		ct.Correo,
		ct.Calificacion,
		ct.Atencion,
		ct.ResolverInstalacion,
		ct.TiempoRazonable,
		ct.Recomendacion,
		ct.CalificacionManera,
		ct.FechaRegistro,
	)

	return ct
}

func CrearComentarioEsiaa(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) ComentarioEsiaa {
	c := ComentarioEsiaa{
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
	utils.LogError("Problemas al crear el estado en la base de datos: (Esiaa) ", false, err)
	defer conn.Close()

	conn.Exec(c.Estado, c.Comentario, c.Formulario, c.Usuario, c.CorreoUsuario, c.IDFormulario, c.FechaRegistro)

	return c
}

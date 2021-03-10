package formularios

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"rws/utils"
	"time"
)

type PrevencionFraude struct {
	ID int
	Nombre,
	Apellido,
	FechaNacimiento,
	LugarResidencia,
	Celular,
	FotoCedula,
	FotoTarjeta,
	Estado,
	FechaRegistro string
	Comentarios []ComentarioPrevencionFraude
}

type ComentarioPrevencionFraude struct {
	ID,
	Estado,
	Comentario,
	FechaRegistro,
	Formulario,
	Usuario,
	CorreoUsuario string
	IDFormulario int
}

func conexionPrevencionFraude() *sql.DB {
	utils.Configuracion()
	connStringMySQL := fmt.Sprintf("%s:%s@%s/%s",
		viper.GetString("basedatos.mysql.rodelag.user"),
		viper.GetString("basedatos.mysql.rodelag.password"),
		viper.GetString("basedatos.mysql.rodelag.server"),
		"rodelag_ovnicom",
	)
	connMySQL, errMySQL := sql.Open("mysql", connStringMySQL)
	if errMySQL != nil {
		utils.LogError("Problemas con la conexion a mysql: (Prevencion Fraude) ", true, errMySQL)
	}
	return connMySQL
}

func VerPrevencionFraude(id int) PrevencionFraude {
	connMySQL := conexionPrevencionFraude()
	defer connMySQL.Close()

	reg := PrevencionFraude{
		Comentarios: func() []ComentarioPrevencionFraude {
			consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_prevencionfraude", id)

			rows, err := connMySQL.Query(consulta)
			utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: (Prevencion Fraude) ", false, err)
			defer rows.Close()

			comentario, comentarios := ComentarioPrevencionFraude{}, []ComentarioPrevencionFraude{}

			for rows.Next() {
				err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
				utils.LogError("Problemas leer los estados: (Prevencion Fraude) ", false, err)
				comentarios = append(comentarios, ComentarioPrevencionFraude{
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

	err := connMySQL.QueryRow("SELECT *, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_prevencionfraude' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_prevencionfraude AS a WHERE a.id = ?;", id).Scan(
		&reg.ID,
		&reg.Nombre,
		&reg.Apellido,
		&reg.FechaNacimiento,
		&reg.LugarResidencia,
		&reg.Celular,
		&reg.FotoCedula,
		&reg.FotoTarjeta,
		&reg.FechaRegistro,
		&reg.Estado,
	)
	utils.LogError("Problemas al leer registro: (Prevencion Fraude) ", false, err)
	return reg
}

func ListarPrevencionFraude() []PrevencionFraude {
	connMySQL := conexionPrevencionFraude()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT a.*, IFNULL((SELECT estado FROM formulario_comentarios WHERE formulario = 'formulario_prevencionfraude' AND idFormulario = a.id ORDER BY fechaRegistro DESC LIMIT 1), 'pendiente') AS estado FROM formulario_prevencionfraude AS a;")
	utils.LogError("Problemas al listar los registros de la base de datos: (Prevencion Fraude) ", true, err)
	defer rows.Close()

	pf := PrevencionFraude{}
	pfs := []PrevencionFraude{}

	for rows.Next() {
		err := rows.Scan(
			&pf.ID,
			&pf.Nombre,
			&pf.Apellido,
			&pf.FechaNacimiento,
			&pf.LugarResidencia,
			&pf.Celular,
			&pf.FotoCedula,
			&pf.FotoTarjeta,
			&pf.FechaRegistro,
			&pf.Estado,
		)
		utils.LogError("Problemas leer los datos: (Prevencion Fraude) ", true, err)
		pfs = append(pfs, PrevencionFraude{
			ID:              pf.ID,
			Nombre:          pf.Nombre,
			Apellido:        pf.Apellido,
			FechaNacimiento: pf.FechaNacimiento,
			LugarResidencia: pf.LugarResidencia,
			Celular:         pf.Celular,
			FotoCedula:      pf.FotoCedula,
			FotoTarjeta:     pf.FotoTarjeta,
			FechaRegistro:   pf.FechaRegistro,
			Estado:          pf.Estado,
			Comentarios: func() []ComentarioPrevencionFraude {
				consulta := fmt.Sprintf("SELECT * FROM formulario_comentarios WHERE formulario = '%s' AND idFormulario = '%d';", "formulario_prevencionfraude", pf.ID)

				rows, err := connMySQL.Query(consulta)
				utils.LogError("Problemas al listar los comentarios de los registros de la base de datos: (Prevencion Fraude) ", true, err)
				defer rows.Close()

				comentario, comentarios := ComentarioPrevencionFraude{}, []ComentarioPrevencionFraude{}

				for rows.Next() {
					err := rows.Scan(&comentario.ID, &comentario.Estado, &comentario.Comentario, &comentario.FechaRegistro, &comentario.Formulario, &comentario.Usuario, &comentario.CorreoUsuario, &comentario.IDFormulario)
					utils.LogError("Problemas leer los estados: (Prevencion Fraude) ", true, err)
					comentarios = append(comentarios, ComentarioPrevencionFraude{
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
	return pfs
}

func CrearPrevencionFraude(nombre, apellido, fechaNacimiento, lugarResidencia, celular, fotoCedula, fotoTarjeta string) PrevencionFraude {
	pf := PrevencionFraude{
		Nombre:          nombre,
		Apellido:        apellido,
		FechaNacimiento: fechaNacimiento,
		LugarResidencia: lugarResidencia,
		Celular:         celular,
		FotoCedula:      fotoCedula,
		FotoTarjeta:     fotoTarjeta,
		FechaRegistro:   time.Now().Format("2006-01-02 15:04:05"),
	}

	connMySQL := conexionPrevencionFraude()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_prevencionfraude (nombre, apellido, fechaNacimiento, lugarResidencia, celular, fotoCedula, fotoTarjeta, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	utils.LogError("Problemas al crear el registro en la base de datos: (Prevencion Fraude) ", false, err)
	defer conn.Close()

	conn.Exec(
		pf.Nombre,
		pf.Apellido,
		pf.FechaNacimiento,
		pf.LugarResidencia,
		pf.Celular,
		pf.FotoCedula,
		pf.FotoTarjeta,
		pf.FechaRegistro,
	)

	return pf
}

func CrearComentarioPrevencionFraude(estado, comentario, formulario, usuario, correoUsuario string, idFormulario int) ComentarioPrevencionFraude {
	c := ComentarioPrevencionFraude{
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
	utils.LogError("Problemas al crear el estado en la base de datos: (Prevencion Fraude) ", false, err)
	defer conn.Close()

	conn.Exec(c.Estado, c.Comentario, c.Formulario, c.Usuario, c.CorreoUsuario, c.IDFormulario, c.FechaRegistro)

	return c
}

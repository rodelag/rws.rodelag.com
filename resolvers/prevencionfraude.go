package resolvers

import (
	_ "github.com/go-sql-driver/mysql"
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
	FechaRegistro string
}

func ListarPrevencionFraude() []PrevencionFraude {
	connMySQL := conexion()
	defer connMySQL.Close()

	rows, err := connMySQL.Query("SELECT * FROM formulario_prevencionfraude;")
	logError("Problemas al listar los registros de la base de datos: ", err)
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
		)
		logError("Problemas leer los datos: ", err)
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

	connMySQL := conexion()
	defer connMySQL.Close()

	conn, err := connMySQL.Prepare("INSERT INTO formulario_prevencionfraude (nombre, apellido, fechaNacimiento, lugarResidencia, celular, fotoCedula, fotoTarjeta, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
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

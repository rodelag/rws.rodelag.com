package resolvers

import (
	_ "github.com/go-sql-driver/mysql"
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
	FechaRegistro string
}

func ListarEsiaa() []Esiaa {
	rows, err := conexion().Query("SELECT * FROM formulario_esiaa;")
	logError("Problemas al listar los registros de la base de datos: ", err)
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
		)
		logError("Problemas leer los datos: ", err)
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

	conn, err := conexion().Prepare("INSERT INTO formulario_esiaa (nombre, apellido, cedula, correo, calificacion, atencion, resolverInstalacion, tiempoRazonable, recomendacion, calificacionManera, FechaRegistro) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	logError("Problemas al crear el registro en la base de datos: ", err)
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

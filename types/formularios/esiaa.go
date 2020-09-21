package formularios

import "github.com/graphql-go/graphql"

var EsiaaType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ESIAA",
	Description: "Encuesta de Satisfacción Instalación AA",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "id del registro",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del cliente",
		},
		"apellido": &graphql.Field{
			Type:        graphql.String,
			Description: "Apellido del lciente",
		},
		"cedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Cédula del cliente",
		},
		"correo": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del cliente",
		},
		"calificacion": &graphql.Field{
			Type:        graphql.String,
			Description: "Calificación del cliente",
		},
		"atencion": &graphql.Field{
			Type:        graphql.String,
			Description: "Atención del cliente",
		},
		"resolverInstalacion": &graphql.Field{
			Type:        graphql.String,
			Description: "Instalación resuelta",
		},
		"tiempoRazonable": &graphql.Field{
			Type:        graphql.String,
			Description: "Tiempo razonable",
		},
		"recomendacion": &graphql.Field{
			Type:        graphql.String,
			Description: "Recomendación",
		},
		"calificacionManera": &graphql.Field{
			Type:        graphql.String,
			Description: "Calificación del cliente en general",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
	},
})

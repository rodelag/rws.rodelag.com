package types

import "github.com/graphql-go/graphql"

var PrevencionFraudeType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PrevencionFraude",
	Description: "Prevención de fraude",
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
			Description: "Apellido del cliente",
		},
		"fechaNacimiento": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de nacimiento",
		},
		"lugarResidencia": &graphql.Field{
			Type:        graphql.String,
			Description: "Lugar de residencia",
		},
		"celular": &graphql.Field{
			Type:        graphql.String,
			Description: "Celular del cliente",
		},
		"fotoCedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto de la cédula",
		},
		"fotoTarjeta": &graphql.Field{
			Type:        graphql.String,
			Description: "Foto frontal de la tarjeta",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Fecha del registro",
		},
	},
})

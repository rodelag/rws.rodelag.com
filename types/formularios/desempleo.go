package formularios

import "github.com/graphql-go/graphql"

var DesempleoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Desempleo",
	Description: "Desempleo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID del registro",
		},
		"nombre": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del cliente",
		},
		"apellido": &graphql.Field{
			Type:        graphql.String,
			Description: "Apellido del cliente",
		},
		"cedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Cédula del cliente",
		},
		"edad": &graphql.Field{
			Type:        graphql.String,
			Description: "Edad del cliente",
		},
		"correo": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del cliente",
		},
		"telefono": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono del cliente",
		},
		"direccionDomicilio": &graphql.Field{
			Type:        graphql.String,
			Description: "Dirección del cliente",
		},
		"nombreEmpresa": &graphql.Field{
			Type:        graphql.String,
			Description: "Empresa donde labora el cliente",
		},
		"tiempoLaboral": &graphql.Field{
			Type:        graphql.String,
			Description: "Tiempo laboral del cliente",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de registro",
		},
		"estado": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado",
		},
	},
})

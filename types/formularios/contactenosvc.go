package formularios

import "github.com/graphql-go/graphql"

var ContactenosVCType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ContactenosVC",
	Description: "Contáctenos de ventas comerciales",
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
		"cedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Cédula del cliente",
		},
		"correo": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo del cliente",
		},
		"telefono": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono del cliente",
		},
		"actividadEconomica": &graphql.Field{
			Type:        graphql.String,
			Description: "Actividad económica del cliente",
		},
		"detalleSolicitud": &graphql.Field{
			Type:        graphql.String,
			Description: "Detalle de solicitud o cotización.",
		},
		"fechaRegistro": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del registro",
		},
	},
})

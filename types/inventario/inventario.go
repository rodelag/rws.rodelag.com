package inventario

import "github.com/graphql-go/graphql"

var InventarioType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Inventario",
	Description: "Inventario",
	Fields: graphql.Fields{
		"Suc": &graphql.Field{
			Type:        graphql.String,
			Description: "Sucursal",
		},
		"Departamento": &graphql.Field{
			Type:        graphql.String,
			Description: "Departamento",
		},
		"Codigo": &graphql.Field{
			Type:        graphql.String,
			Description: "Código",
		},
		"Descripcion": &graphql.Field{
			Type:        graphql.String,
			Description: "Descripción",
		},
		"Precio": &graphql.Field{
			Type:        graphql.String,
			Description: "Precio",
		},
		"Cant": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad",
		},
		"Parte": &graphql.Field{
			Type:        graphql.String,
			Description: "Numero de parte",
		},
		"Marca": &graphql.Field{
			Type:        graphql.String,
			Description: "Marca",
		},
		"Oferta": &graphql.Field{
			Type:        graphql.String,
			Description: "Oferta",
		},
		"FecIni": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha inicial de la oferta",
		},
		"FecFin": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha final de la oferta",
		},
	},
})

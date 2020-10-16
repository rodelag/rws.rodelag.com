package conteo

import "github.com/graphql-go/graphql"

var ConteoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Conteo",
	Description: "Conteo de clientes de las sucursales",
	Fields: graphql.Fields{
		"registroID": &graphql.Field{
			Type:        graphql.String,
			Description: "Registro único para cada registro.",
		},
		"registroNumero": &graphql.Field{
			Type:        graphql.String,
			Description: "Número único consecutivo.",
		},
		"registroEmpresa": &graphql.Field{
			Type:        graphql.String,
			Description: "Empresa",
		},
		"registroSucursal": &graphql.Field{
			Type:        graphql.String,
			Description: "Sucursal",
		},
		"registroSucursalNombre": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de la sucursal",
		},
		"registroEntrada": &graphql.Field{
			Type:        graphql.String,
			Description: "Entrada",
		},
		"registroSalida": &graphql.Field{
			Type:        graphql.String,
			Description: "Salida",
		},
		"registroFacturas": &graphql.Field{
			Type:        graphql.String,
			Description: "Facturas",
		},
		"registroTiquetePromedio": &graphql.Field{
			Type:        graphql.String,
			Description: "Tiquete promedio",
		},
		"registroArticulos": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad de artículos",
		},
		"registroVenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Venta",
		},
		"registroFecha": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha",
		},
		"registroIP": &graphql.Field{
			Type:        graphql.String,
			Description: "IP de la camara de conteo",
		},
	},
})

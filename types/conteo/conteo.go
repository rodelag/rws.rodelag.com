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
		"registroEntradaAnt": &graphql.Field{
			Type:        graphql.String,
			Description: "Entrada anterior un año atrás",
		},
		"registroEntrada": &graphql.Field{
			Type:        graphql.String,
			Description: "Entrada",
		},
		"registroSalidaAnt": &graphql.Field{
			Type:        graphql.String,
			Description: "Salida anterior un año atrás",
		},
		"registroSalida": &graphql.Field{
			Type:        graphql.String,
			Description: "Salida",
		},
		"registroFacturasAnt": &graphql.Field{
			Type:        graphql.String,
			Description: "Facturas anterior un año atrás",
		},
		"registroFacturas": &graphql.Field{
			Type:        graphql.String,
			Description: "Facturas",
		},
		"registroTiquetePromedioAnt": &graphql.Field{
			Type:        graphql.String,
			Description: "Tiquete promedio anterior un año atrás",
		},
		"registroTiquetePromedio": &graphql.Field{
			Type:        graphql.String,
			Description: "Tiquete promedio",
		},
		"registroArticulosAnt": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad de artículos anterior un año atrás",
		},
		"registroArticulos": &graphql.Field{
			Type:        graphql.String,
			Description: "Cantidad de artículos",
		},
		"registroVentaAnt": &graphql.Field{
			Type:        graphql.String,
			Description: "Venta anterior un año atrás",
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

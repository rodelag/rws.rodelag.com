package reportes_elconix

import "github.com/graphql-go/graphql"

var InventarioSucursalType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "InventarioSucursal",
	Description: "Inventario de las sucursales",
	Fields: graphql.Fields{
		"bodega": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre de la bodega",
		},
		"inventario": &graphql.Field{
			Type:        graphql.String,
			Description: "Monto del inventario",
		},
	},
})

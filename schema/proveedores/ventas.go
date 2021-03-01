package proveedores

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/proveedores"
	types "rws/types/proveedores"
)

func ProveedorVentasQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"proveedor_ventas_listar": {
			Type:        graphql.NewList(types.ProveedorVentasType),
			Description: "Listar inventario del proveedor",
			Args: graphql.FieldConfigArgument{
				"proveedor": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Campo o columna para deferenciar entre proveedor o fabricante",
				},
				"proveedorID": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "ID del proveedor o la marca del mismo",
				},
				"sucursal": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Sucursal",
				},
				"campo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Campo o columna secundaria",
				},
				"condicion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Condición para el campo secundaria",
				},
				"campo2": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Campo o columna secundaria",
				},
				"condicion2": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Condición para el campo secundaria",
				},
				"fechainicial": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Fecha inicial",
				},
				"fechafinal": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Fecha final",
				},
				"cursor": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "Cursor para identificar cada registro, id único y secuencial",
				},
				"limite": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "Limite de los registros que se traen",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				proveedor, _ := p.Args["proveedor"].(string)
				proveedorID, _ := p.Args["proveedorID"].(string)
				sucursal, _ := p.Args["sucursal"].(string)
				campo, _ := p.Args["campo"].(string)
				condicion, _ := p.Args["condicion"].(string)
				campo2, _ := p.Args["campo2"].(string)
				condicion2, _ := p.Args["condicion2"].(string)
				fechainicial, _ := p.Args["fechainicial"].(string)
				fechafinal, _ := p.Args["fechafinal"].(string)

				cursor, _ := p.Args["cursor"].(int)
				limite, _ := p.Args["limite"].(int)

				return resolvers.ListarProveedorVentas(proveedor, proveedorID, sucursal, campo, condicion, campo2, condicion2, fechainicial, fechafinal, cursor, limite), nil
			},
		},
	}
	return schemas
}

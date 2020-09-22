package proveedores

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/proveedores"
	types "rws/types/proveedores"
)

func ProveedoreInventarioQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"proveedor_inventario_listar": {
			Type:        graphql.NewList(types.ProveedorInventarioType),
			Description: "Listar inventario del proveedor",
			Args: graphql.FieldConfigArgument{
				"sucursal": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Sucursal",
				},
				"proveedor": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Campo o columna para deferenciar entre proveedor o fabricante",
				},
				"proveedorID": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "ID del proveedor o la marca del mismo",
				},
				"campo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Campo o columna secundaria",
				},
				"condicion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Condición para el campo secundaria",
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

				sucursal, _ := p.Args["sucursal"].(string)
				proveedor, _ := p.Args["proveedor"].(string)
				proveedorID, _ := p.Args["proveedorID"].(string)
				campo, _ := p.Args["campo"].(string)
				condicion, _ := p.Args["condicion"].(string)

				cursor, _ := p.Args["cursor"].(int)
				limite, _ := p.Args["limite"].(int)

				return resolvers.ListarProveedorInventario(sucursal, proveedor, proveedorID, campo, condicion, cursor, limite), nil
			},
		},
	}
	return schemas
}

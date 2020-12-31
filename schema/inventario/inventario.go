package inventario

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/inventario"
	types "rws/types/inventario"
)

func InventarioQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"inventario_listar": {
			Type:        graphql.NewList(types.InventarioType),
			Description: "Listar el inventario.",
			Args: graphql.FieldConfigArgument{
				"consulta": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Producto a buscar, código o nombre",
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
				consulta, _ := p.Args["consulta"].(string)

				return resolvers.ListarInventario(consulta), nil
			},
		},
	}
	return schemas
}

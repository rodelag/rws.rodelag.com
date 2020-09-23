package reportes_elconix

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/reportes_elconix"
	types "rws/types/reportes_elconix"
)

func InventarioSucursalQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"reporte_eis_listar": {
			Type:        graphql.NewList(types.InventarioSucursalType),
			Description: "Listar inventario de las sucursales",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarInventarioSucursal(), nil
			},
		},
	}
	return schemas
}

package ventas_detalles

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/ventas_detalles"
	types "rws/types/ventas_detalles"
)

func VentasDetallesQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"ventas_listar": {
			Type:        graphql.NewList(types.VentasType),
			Description: "Listar el detalle de las ventas.",
			Args: graphql.FieldConfigArgument{
				"inicio": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Fecha de inicio.",
				},
				"fin": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Fecha de fin.",
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
				inicio, _ := p.Args["inicio"].(string)
				fin, _ := p.Args["fin"].(string)

				return resolvers.ListarVentas(inicio, fin), nil
			},
		},
	}
	return schemas
}

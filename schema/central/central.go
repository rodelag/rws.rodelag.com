package central

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/central"
	types "rws/types/central"
)

func CentralTelefonicaQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"central": {
			Type:        types.CentralTelefonicaType,
			Description: "Consulta a la central teléfonica",
			Args: graphql.FieldConfigArgument{
				"consulta": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre o Extensión",
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
				if consulta, ok := p.Args["consulta"].(string); ok {
					return resolvers.CentralTelefonica(consulta), nil
				}
				return nil, nil
			},
		},
	}
	return schemas
}

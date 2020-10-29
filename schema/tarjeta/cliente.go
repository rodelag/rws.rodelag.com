package tarjeta

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/tarjeta"
	types "rws/types/tarjeta"
)

func TarjetaRodelagQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"cliente": {
			Type:        types.ClienteType,
			Description: "Información general del cliente con respecto a la tarjeta Rodelag.",
			Args: graphql.FieldConfigArgument{
				"cedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cédula del cliente",
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
				cedula, _ := p.Args["cedula"].(string)

				return resolvers.ClienteTarjetaRodelag(cedula), nil
			},
		},
	}
	return schemas
}

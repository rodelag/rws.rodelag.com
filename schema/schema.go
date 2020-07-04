package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

var Schema graphql.Schema

func init() {
	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"pagosach": &graphql.Field{
				Type: graphql.NewList(types.PagoACHType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
					if err != nil {
						return nil, err
					}
					if !isValid {
						return nil, gqlerrors.FormatError(errors.New("Invalid token"))
					}
					return resolvers.ListarPagosACH(), nil
				},
			},
		},
	})

	Mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"crearpagoach": &graphql.Field{
				Type:        types.PagoACHType,
				Description: "Creación de pago ACH",
				Args: graphql.FieldConfigArgument{
					"nombre": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Nombre del cliente",
					},
					"apellido": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Apellido del cliente",
					},
					"titularCuenta": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Titular de la cuenta",
					},
					"cedula": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Cédula del cliente",
					},
					"correo": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Correo del cliente",
					},
					"telefono": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Teléfono del cliente",
					},
					"compraOrigen": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Lugar donde hizo la compra",
					},
					"numeroOrden": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Pedido o Cotización",
					},
					"fotoComprobante": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Foto del comprobante de pago",
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					isValid, err := auth.ValidateToken(params.Context.Value("token").(string))
					if err != nil {
						return nil, err
					}
					if !isValid {
						return nil, gqlerrors.FormatError(errors.New("Invalid token"))
					}

					nombre, _ := params.Args["nombre"].(string)
					apellido, _ := params.Args["apellido"].(string)
					titularCuenta, _ := params.Args["titularCuenta"].(string)
					cedula, _ := params.Args["cedula"].(string)
					correo, _ := params.Args["correo"].(string)
					telefono, _ := params.Args["telefono"].(string)
					compraOrigen, _ := params.Args["compraOrigen"].(string)
					numeroOrden, _ := params.Args["numeroOrden"].(string)
					fotoComprobante, _ := params.Args["fotoComprobante"].(string)

					return resolvers.CrearPagoACH(nombre, apellido, titularCuenta, cedula, correo, telefono, compraOrigen, numeroOrden, fotoComprobante), nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	})
	if err != nil {
		panic(err)
	}
	Schema = schema
}

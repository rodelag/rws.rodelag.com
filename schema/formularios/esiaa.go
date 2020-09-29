package formularios

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	resolvers "rws/resolvers/formularios"
	types "rws/types/formularios"
)

func EsiaaQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"esiaa_ver": {
			Type:        types.EsiaaType,
			Description: "Ver el de detalle de la Encuesta de Satisfacción Instalación AA",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "ID del registro",
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

				if id, ok := p.Args["id"].(int); ok {
					return resolvers.VerEsiaa(id), nil
				}
				return nil, nil
			},
		},
		"esiaa_listar": {
			Type: graphql.NewList(types.EsiaaType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarEsiaa(), nil
			},
		},
	}
	return schemas
}

func EsiaaMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"esiaa_crear": &graphql.Field{
			Type:        types.EsiaaType,
			Description: "Contáctenos ventas comerciales",
			Args: graphql.FieldConfigArgument{
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del lciente",
				},
				"cedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cédula del cliente",
				},
				"correo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del cliente",
				},
				"calificacion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Calificación del cliente",
				},
				"atencion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Atención del cliente",
				},
				"resolverInstalacion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Instalación resuelta",
				},
				"tiempoRazonable": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Tiempo razonable",
				},
				"recomendacion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Recomendación",
				},
				"calificacionManera": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Calificación del cliente en general",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(params.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				nombre, _ := params.Args["nombre"].(string)
				apellido, _ := params.Args["apellido"].(string)
				cedula, _ := params.Args["cedula"].(string)
				correo, _ := params.Args["correo"].(string)
				calificacion, _ := params.Args["calificacion"].(string)
				atencion, _ := params.Args["atencion"].(string)
				resolverInstalacion, _ := params.Args["resolverInstalacion"].(string)
				tiempoRazonable, _ := params.Args["tiempoRazonable"].(string)
				recomendacion, _ := params.Args["recomendacion"].(string)
				calificacionManera, _ := params.Args["calificacionManera"].(string)

				return resolvers.CrearEsiaa(nombre, apellido, cedula, correo, calificacion, atencion, resolverInstalacion, tiempoRazonable, recomendacion, calificacionManera), nil
			},
		},
		"esiaa_crear_comentario": &graphql.Field{
			Type:        types.EsiaaComentarioType,
			Description: "Creación de comentario de la Encuesta de Satisfacción Instalación AA",
			Args: graphql.FieldConfigArgument{
				"estado": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Estado del registro",
				},
				"comentario": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Comentario del agente para con el registro",
				},
				"formulario": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Formulario al que pertenece el estado",
				},
				"usuario": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Usuario que gestiona el registro",
				},
				"correoUsuario": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del usuario que gestiona el registro",
				},
				"idFormulario": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "ID del registro",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(params.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}
				estado, _ := params.Args["estado"].(string)
				comentario, _ := params.Args["comentario"].(string)
				formulario, _ := params.Args["formulario"].(string)
				usuario, _ := params.Args["usuario"].(string)
				correoUsuario, _ := params.Args["correoUsuario"].(string)
				idFormulario, _ := params.Args["idFormulario"].(int)

				return resolvers.CrearComentarioEsiaa(estado, comentario, formulario, usuario, correoUsuario, idFormulario), nil
			},
		},
	}
	return schemas
}

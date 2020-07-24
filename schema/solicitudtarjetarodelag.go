package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"rws/auth"
	"rws/resolvers"
	"rws/types"
)

func SolicitudTarjetaRodelagQuery() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"solicitudtarjetarodelag_listar": {
			Type: graphql.NewList(types.SolicitudTarjetaRodelagType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
				if err != nil {
					return nil, err
				}
				if !isValid {
					return nil, gqlerrors.FormatError(errors.New("Token de autorización inválido"))
				}

				return resolvers.ListarSolicitudTarjetaRodelag(), nil
			},
		},
	}
	return schemas
}

func SolicitudTarjetaRodelagMutation() map[string]*graphql.Field {
	schemas := map[string]*graphql.Field{
		"solicitudtarjetarodelag_crear": &graphql.Field{
			Type:        types.SolicitudTarjetaRodelagType,
			Description: "Creación de Solicitud de tarjeta rodelag",
			Args: graphql.FieldConfigArgument{
				"sucursal": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Sucursal más cercana",
				},
				"nombre": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre del cliente",
				},
				"apellido": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Apellido del cliente",
				},
				"fechaNacimiento": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Fecha de nacimiento",
				},
				"lugarNacimiento": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Lugar de nacimiento",
				},
				"nacionalidad": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nacionalidad del cliente",
				},
				"cedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cédula del cliente",
				},
				"fotoCedula": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto de la cédula",
				},
				"fotoFicha": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto de la ficha del seguro social",
				},
				"fotoRecibo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Foto de recibo de servicios",
				},
				"estadoCivil": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Estado civil",
				},
				"correo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Correo del cliente",
				},
				"direccionResidencia": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "",
				},
				"barrio": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Barriada de residencia",
				},
				"provincia": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Provincia de residencia",
				},
				"telefonoResidencia": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono de la residencia",
				},
				"celular": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Celular del cliente",
				},
				"credito": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Monto de Línea de Crédito",
				},
				"educacion": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nivel educativo del cliente",
				},
				"nombreEmpresa": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre de la Empresa para la cual trabaja el cliente",
				},
				"tipoNegocio": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Tipo de negocio en donde labora el cliente",
				},
				"cargo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Cargo o posición en la empresa donde trabaja el cliente",
				},
				"tiempoLaboral": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Tiempo laborado en la empresa del cliente",
				},
				"direccionTrabajo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Dirección de la empresa donde labora el cliente",
				},
				"telefonoTrabajo": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono de la empresa donde labora el cliente",
				},
				"extension": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Extensión teléfono de la empresa donde labora el cliente",
				},
				"salarioMensual": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Salario mensual del cliente",
				},
				"fuentesIngreso": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Otras fuentes de ingreso del cliente",
				},
				"montoFuentesIngreso": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Monto provenientes de otra fuente de ingreso",
				},
				"detalleFuentesIngreso": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Detalle de la otra fuentes de ingreso",
				},
				"nombreReferenciaUno": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre de referencia número uno",
				},
				"telefonoReferenciaUno": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono de referencia número uno",
				},
				"nombreReferenciaDos": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre de referencia número dos",
				},
				"telefonoReferenciaDos": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono de referencia número dos",
				},
				"nombreReferenciaTres": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Nombre de referencia número tres",
				},
				"telefonoReferenciaTres": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Teléfono de referencia número tres",
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

				sucursal, _ := params.Args["sucursal"].(string)
				nombre, _ := params.Args["nombre"].(string)
				apellido, _ := params.Args["apellido"].(string)
				fechaNacimiento, _ := params.Args["fechaNacimiento"].(string)
				lugarNacimiento, _ := params.Args["lugarNacimiento"].(string)
				nacionalidad, _ := params.Args["nacionalidad"].(string)
				cedula, _ := params.Args["cedula"].(string)
				fotoCedula, _ := params.Args["fotoCedula"].(string)
				fotoFicha, _ := params.Args["fotoFicha"].(string)
				fotoRecibo, _ := params.Args["fotoRecibo"].(string)
				estadoCivil, _ := params.Args["estadoCivil"].(string)
				correo, _ := params.Args["correo"].(string)
				direccionResidencia, _ := params.Args["direccionResidencia"].(string)
				barrio, _ := params.Args["barrio"].(string)
				provincia, _ := params.Args["provincia"].(string)
				telefonoResidencia, _ := params.Args["telefonoResidencia"].(string)
				celular, _ := params.Args["celular"].(string)
				credito, _ := params.Args["credito"].(string)
				educacion, _ := params.Args["educacion"].(string)
				nombreEmpresa, _ := params.Args["nombreEmpresa"].(string)
				tipoNegocio, _ := params.Args["tipoNegocio"].(string)
				cargo, _ := params.Args["cargo"].(string)
				tiempoLaboral, _ := params.Args["tiempoLaboral"].(string)
				direccionTrabajo, _ := params.Args["direccionTrabajo"].(string)
				telefonoTrabajo, _ := params.Args["telefonoTrabajo"].(string)
				extension, _ := params.Args["extension"].(string)
				salarioMensual, _ := params.Args["salarioMensual"].(string)
				fuentesIngreso, _ := params.Args["fuentesIngreso"].(string)
				montoFuentesIngreso, _ := params.Args["montoFuentesIngreso"].(string)
				detalleFuentesIngreso, _ := params.Args["detalleFuentesIngreso"].(string)
				nombreReferenciaUno, _ := params.Args["nombreReferenciaUno"].(string)
				telefonoReferenciaUno, _ := params.Args["telefonoReferenciaUno"].(string)
				nombreReferenciaDos, _ := params.Args["nombreReferenciaDos"].(string)
				telefonoReferenciaDos, _ := params.Args["telefonoReferenciaDos"].(string)
				nombreReferenciaTres, _ := params.Args["nombreReferenciaTres"].(string)
				telefonoReferenciaTres, _ := params.Args["telefonoReferenciaTres"].(string)

				return resolvers.CrearSolicitudTarjetaRodelag(
					sucursal,
					nombre,
					apellido,
					fechaNacimiento,
					lugarNacimiento,
					nacionalidad,
					cedula,
					fotoCedula,
					fotoFicha,
					fotoRecibo,
					estadoCivil,
					correo,
					direccionResidencia,
					barrio,
					provincia,
					telefonoResidencia,
					celular,
					credito,
					educacion,
					nombreEmpresa,
					tipoNegocio,
					cargo,
					tiempoLaboral,
					direccionTrabajo,
					telefonoTrabajo,
					extension,
					salarioMensual,
					fuentesIngreso,
					montoFuentesIngreso,
					detalleFuentesIngreso,
					nombreReferenciaUno,
					telefonoReferenciaUno,
					nombreReferenciaDos,
					telefonoReferenciaDos,
					nombreReferenciaTres,
					telefonoReferenciaTres,
				), nil
			},
		},
	}
	return schemas
}

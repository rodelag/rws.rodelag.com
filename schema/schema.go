package schema

import (
	"github.com/graphql-go/graphql"
	"rws/schema/formularios"
	"rws/schema/proveedores"
	"rws/schema/reportes_elconix"
)

var Schema graphql.Schema

func init() {
	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//Formulario de Pago por ACH
			"pagoach_ver":    formularios.PagoACHQuery()["pagoach_ver"],
			"pagoach_listar": formularios.PagoACHQuery()["pagoach_listar"],

			//Formulario de Solicitud de Estado de Cuenta
			"estadocuenta_ver":    formularios.EstadoCuentaQuery()["estadocuenta_ver"],
			"estadocuenta_listar": formularios.EstadoCuentaQuery()["estadocuenta_listar"],

			//Formulario de Comprobante de Pago
			"comprobantepago_ver":    formularios.ComprobantePagoQuery()["comprobantepago_ver"],
			"comprobantepago_listar": formularios.ComprobantePagoQuery()["comprobantepago_listar"],

			//Formulario de Solicitud de Aplazamiento de Pagos por la Crisis del COVID-19
			"solicitudaplazamientopago_ver":    formularios.SolicitudAplazamientoPagoQuery()["solicitudaplazamientopago_ver"],
			"solicitudaplazamientopago_listar": formularios.SolicitudAplazamientoPagoQuery()["solicitudaplazamientopago_listar"],

			//Formulario de Prevención de Fraude
			"prevencionfraude_listar": formularios.PrevencionFraudeQuery()["prevencionfraude_listar"],

			//Formulario de Solicitud Tarjeta Rodelag
			"solicitudtarjetarodelag_listar": formularios.SolicitudTarjetaRodelagQuery()["solicitudtarjetarodelag_listar"],

			//Formulario de Reclamo
			"reclamo_listar": formularios.ReclamoQuery()["reclamo_listar"],

			//Formulario de Notificación Descuento
			"notificaciondescuento_listar": formularios.NotificacionDescuentoQuery()["notificaciondescuento_listar"],

			//Formulario de Contáctenos
			"contactenos_listar": formularios.ContactenosQuery()["contactenos_listar"],

			//Formulario de Contáctenos de ventas comerciales
			"contactenosvc_listar": formularios.ContactenosVCQuery()["contactenosvc_listar"],

			//Encuesta de Satisfacción Instalación AA
			"esiaa_listar": formularios.EsiaaQuery()["esiaa_listar"],

			//Inventario de los proveedores
			"proveedor_inventario_listar": proveedores.ProveedoreInventarioQuery()["proveedor_inventario_listar"],

			//Inventario de las sucursales / Bodegas
			"reporte_eis_listar": reportes_elconix.InventarioSucursalQuery()["reporte_eis_listar"],
		},
	})

	Mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//Formulario de Pago por ACH
			"pagoach_crear":            formularios.PagoACHMutation()["pagoach_crear"],
			"pagoach_crear_comentario": formularios.PagoACHMutation()["pagoach_crear_comentario"],

			//Formulario de Solicitud de Estado de Cuenta
			"estadocuenta_crear":            formularios.EstadoCuentaMutation()["estadocuenta_crear"],
			"estadocuenta_crear_comentario": formularios.EstadoCuentaMutation()["estadocuenta_crear_comentario"],

			//Formulario de Comprobante de Pago
			"comprobantepago_crear":            formularios.ComprobantePagoMutation()["comprobantepago_crear"],
			"comprobantepago_crear_comentario": formularios.ComprobantePagoMutation()["comprobantepago_crear_comentario"],

			//Formulario de Solicitud de Aplazamiento de Pagos por la Crisis del COVID-19
			"solicitudaplazamientopago_crear":            formularios.SolicitudAplazamientoPagoMutation()["solicitudaplazamientopago_crear"],
			"solicitudaplazamientopago_modificar":        formularios.SolicitudAplazamientoPagoMutation()["solicitudaplazamientopago_modificar"],
			"solicitudaplazamientopago_crear_comentario": formularios.SolicitudAplazamientoPagoMutation()["solicitudaplazamientopago_crear_comentario"],

			//Formulario de Prevención de Fraude
			"prevencionfraude_crear": formularios.PrevencionFraudeMutation()["prevencionfraude_crear"],

			//Formulario de Solicitud Tarjeta Rodelag
			"solicitudtarjetarodelag_crear": formularios.SolicitudTarjetaRodelagMutation()["solicitudtarjetarodelag_crear"],

			//Formulario de Reclamo
			"reclamo_crear": formularios.ReclamoMutation()["reclamo_crear"],

			//Formulario de Notificación Descuento
			"notificaciondescuento_crear": formularios.NotificacionDescuentoMutation()["notificaciondescuento_crear"],

			//Formulario de Contáctenos
			"contactenos_crear": formularios.ContactenosMutation()["contactenos_crear"],

			//Formulario de Contáctenos de ventas comerciales
			"contactenosvc_crear": formularios.ContactenosVCMutation()["contactenosvc_crear"],

			//Encuesta de Satisfacción Instalación AA
			"esiaa_crear": formularios.EsiaaMutation()["esiaa_crear"],
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

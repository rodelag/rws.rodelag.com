package tarjeta

import "github.com/graphql-go/graphql"

var EstadoCuentaDetalleType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Estado Cuenta Detalle",
	Description: "Detalle del estado de cuenta",
	Fields: graphql.Fields{
		"registroFechaTran": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de transacción",
		},
		"registroNumeroDocumento": &graphql.Field{
			Type:        graphql.String,
			Description: "Número de documento",
		},
		"registroDescripcionDocumento": &graphql.Field{
			Type:        graphql.String,
			Description: "Descripción de documento",
		},
		"registroDebito": &graphql.Field{
			Type:        graphql.String,
			Description: "Débito",
		},
		"registroCredito": &graphql.Field{
			Type:        graphql.String,
			Description: "Crédito",
		},
		"registroSaldo": &graphql.Field{
			Type:        graphql.String,
			Description: "Saldo",
		},
	},
})

var ClienteType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Cliente",
	Description: "Información general del cliente con respecto a la tarjeta Rodelag.",
	Fields: graphql.Fields{
		"registroCedula": &graphql.Field{
			Type:        graphql.String,
			Description: "Cédula",
		},
		"registroCuenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Número de cuenta",
		},
		"registroNombre": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre",
		},
		"registroLimite": &graphql.Field{
			Type:        graphql.String,
			Description: "Límite",
		},
		"registroUltimoSaldo": &graphql.Field{
			Type:        graphql.String,
			Description: "Último Saldo",
		},
		"registroFechaUltimaCompra": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de la última compra",
		},
		"registroMontoUltimaCompra": &graphql.Field{
			Type:        graphql.String,
			Description: "Monto de la última compra",
		},
		"registroFechaApertura": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de apertura de la cuenta",
		},
		"registroFechaCorte": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de corte",
		},
		"registroPagarAntesDe": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha límite para el pago",
		},
		"registroCorte": &graphql.Field{
			Type:        graphql.String,
			Description: "Corte",
		},
		"registroEstado": &graphql.Field{
			Type:        graphql.String,
			Description: "Estado de la cuenta",
		},
		"registroFechaUltimoPago": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del último pago",
		},
		"registroMontoUltimoPago": &graphql.Field{
			Type:        graphql.String,
			Description: "Monto del último pago",
		},
		"registroPagoMinimo": &graphql.Field{
			Type:        graphql.String,
			Description: "Pago mínimo",
		},
		"registroCorreo": &graphql.Field{
			Type:        graphql.String,
			Description: "Correo",
		},
		"registroTelefono": &graphql.Field{
			Type:        graphql.String,
			Description: "Teléfono",
		},
		"registroFechaNacimiento": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha de nacimiento",
		},
		"registroSexo": &graphql.Field{
			Type:        graphql.String,
			Description: "Sexo del cliente",
		},
		"registroSucursal": &graphql.Field{
			Type:        graphql.String,
			Description: "Sucursal de registro",
		},
		"registroTipoCliente": &graphql.Field{
			Type:        graphql.String,
			Description: "Tipo de cliente",
		},
		"registroFechaEstoCuenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha del estado de cuenta",
		},
		"registroNumeroCuenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Número del estado de cuenta",
		},
		"registroLimiteCreditoEstadoCuenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Límite de crédito del estado de cuenta",
		},
		"registroSaldoDisponible": &graphql.Field{
			Type:        graphql.String,
			Description: "Saldo disponible",
		},
		"registroFechaInicioTran": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha inicio Transacción",
		},
		"registroFechaFinTran": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha fin Transacción",
		},
		"registroSaldoAnterior": &graphql.Field{
			Type:        graphql.String,
			Description: "Saldo Anterior",
		},
		"registroPagoMinimoEstadoCuenta": &graphql.Field{
			Type:        graphql.String,
			Description: "Pago mínimo del estado de cuenta",
		},
		"registroSaldoCorriente": &graphql.Field{
			Type:        graphql.String,
			Description: "Saldo corriente",
		},
		"registroSaldo30Dias": &graphql.Field{
			Type:        graphql.String,
			Description: "Saldo a 30 Días",
		},
		"registroSaldo60Dias": &graphql.Field{
			Type:        graphql.String,
			Description: "Saldo a 60 Días",
		},
		"registroSaldo90Dias": &graphql.Field{
			Type:        graphql.String,
			Description: "Saldo a 90 Días",
		},
		"registroSaldo120Dias": &graphql.Field{
			Type:        graphql.String,
			Description: "Saldo a 120 Días",
		},
		"registroFechaFinPago": &graphql.Field{
			Type:        graphql.String,
			Description: "Fecha Fin Pago",
		},
		"registroEstadoCuentaDetalle": &graphql.Field{
			Type:        graphql.NewList(EstadoCuentaDetalleType),
			Description: "Estado de Cuenta - Detalle",
		},
	},
})

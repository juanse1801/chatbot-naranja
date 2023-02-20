package configs

var StateConfig = map[string]map[string]map[string]string{
	"Bienvenida": {
		"default": {
			"response":   Bienvenida,
			"next_state": "Terms",
			"execute":    "update_state",
		},
	},
	"Terms": {
		"1": {
			"response":   AceptoTerminos,
			"next_state": "EntitySelection",
			"execute":    "update_state",
		},
		"2": {
			"response":   NoAceptoTerminos,
			"next_state": "",
			"execute":    "delete_interaction",
		},
		"not_valid": {
			"response": OpcionIncorrecta,
			"execute":  "update_state",
		},
	},
	"EntitySelection": {
		"1": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_state",
		},
		"2": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_state",
		},
		"3": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_state",
		},
		"4": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_state",
		},
		"5": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_state",
		},
		"6": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_state",
		},
		"not_valid": {
			"response": OpcionIncorrecta,
			"execute":  "update_state",
		},
	},
	"ServiceSelection": {
		"1": {
			"response":   PostAsesoriaComercial,
			"next_state": "EndInteraction",
			"execute":    "update_state",
		},
		"2": {
			"response":   PostAsesoriaComercial,
			"next_state": "EndInteraction",
			"execute":    "update_state",
		},
		"not_valid": {
			"response": OpcionIncorrecta,
			"execute":  "update_state",
		},
	},

	"EndInteraction": {
		"default": {
			"response":   FinInteracción,
			"next_state": "",
			"execute":    "delete_interaction",
		},
	},
}

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
		"0": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_entity",
		},
		"1": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_entity",
		},
		"2": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_entity",
		},
		"3": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_entity",
		},
		"4": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_entity",
		},
		"5": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_entity",
		},
		"6": {
			"response":   SeleccionDeServicio,
			"next_state": "ServiceSelection",
			"execute":    "update_entity",
		},
		"not_valid": {
			"response": OpcionIncorrecta,
			"execute":  "update_state",
		},
	},
	"ServiceSelection": {
		"0": {
			"response":   AsesoriaComercial,
			"next_state": "TypeSelection",
			"execute":    "update_service",
		},
		"1": {
			"response":   PostAsesoriaAdministrativa,
			"next_state": "EndInteraction",
			"execute":    "update_service",
		},
		"not_valid": {
			"response": OpcionIncorrecta,
			"execute":  "update_state",
		},
	},
	"TypeSelection": {
		"0": {
			"response":   ZoneSelection,
			"next_state": "ZoneSelection",
			"execute":    "update_type",
		},
		"1": {
			"response":   ZoneSelection,
			"next_state": "ZoneSelection",
			"execute":    "update_type",
		},
		"2": {
			"response":   ZoneSelection,
			"next_state": "ZoneSelection",
			"execute":    "update_type",
		},
		"3": {
			"response":   ZoneSelection,
			"next_state": "ZoneSelection",
			"execute":    "update_type",
		},
		"4": {
			"response":   ZoneSelection,
			"next_state": "ZoneSelection",
			"execute":    "update_type",
		},
		"not_valid": {
			"response": OpcionIncorrecta,
			"execute":  "update_state",
		},
	},
	"ZoneSelection": {
		"0": {
			"response":   PostAsesoriaComercial,
			"next_state": "EndInteraction",
			"execute":    "update_zone",
		},
		"1": {
			"response":   PostAsesoriaComercial,
			"next_state": "EndInteraction",
			"execute":    "update_zone",
		},
		"2": {
			"response":   PostAsesoriaComercial,
			"next_state": "EndInteraction",
			"execute":    "update_zone",
		},
		"3": {
			"response":   PostAsesoriaComercial,
			"next_state": "EndInteraction",
			"execute":    "update_zone",
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
			"execute":    "end_interaction",
		},
	},
}

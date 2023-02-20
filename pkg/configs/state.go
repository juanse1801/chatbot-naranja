package configs

var StateConfig = map[string]map[string]map[string]string{
	"Bienvenida": {
		"default": {
			"response": "Hola como estas?, digite 1 o 2",
		},
		"next_state": {
			"value": "Segundo estado",
		},
		"execute": {
			"value": "update_state",
		},
	},
	"Segundo estado": {
		"1": {
			"response": "Respuesta opcion 1 segundo estado, escriba su correo",
		},
		"2": {
			"response": "Respuesta opcion 2 segundo estado, escriba su correo",
		},
		"not_valid": {
			"response": "Repuesta no valida segundo estado porfavor digite 1 o 2",
		},
		"next_state": {
			"value": "Tercer estado",
		},
		"execute": {
			"value": "save_mail",
		},
	},
	"Tercer estado": {
		"default": {
			"response": "porfavor digite la informacion",
		},
		"next_state": {
			"value": "Cuarto estado",
		},
		"execute": {
			"value": "send_mail",
		},
	},
	"Cuarto estado": {
		"default": {
			"response": "gracias por enviar su informacion pronto nos comunicaremos",
		},
		"execute": {
			"value": "delete_interaction",
		},
		"next_state": {
			"value": "Quinto estado",
		},
	},
}

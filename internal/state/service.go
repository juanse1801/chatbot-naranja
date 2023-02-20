package state

import "github.com/juanse1801/chatbot-naranja/pkg/configs"

type Service interface {
	NextState(actualState string, option string) (string, string, string)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) NextState(actualState string, option string) (string, string, string) {

	if isDefault(actualState) {
		return configs.StateConfig[actualState]["next_state"]["value"],
			configs.StateConfig[actualState]["default"]["response"],
			configs.StateConfig[actualState]["execute"]["value"]

	}

	if isValidEntry(actualState, option) {
		return configs.StateConfig[actualState]["next_state"]["value"],
			configs.StateConfig[actualState][option]["response"],
			configs.StateConfig[actualState]["execute"]["value"]
	} else {
		return actualState,
			configs.StateConfig[actualState]["not_valid"]["response"],
			configs.StateConfig[actualState]["execute"]["value"]
	}
}

func isValidEntry(state string, option string) bool {
	_, isValid := configs.StateConfig[state][option]

	return isValid
}

func isDefault(state string) bool {
	_, response := configs.StateConfig[state]["default"]

	return response
}

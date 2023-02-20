package webhooks

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/juanse1801/chatbot-naranja/internal/domain"
	interactionService "github.com/juanse1801/chatbot-naranja/internal/interaction"
	"github.com/juanse1801/chatbot-naranja/internal/messaging"
	schedulerService "github.com/juanse1801/chatbot-naranja/internal/scheduler"
	stateService "github.com/juanse1801/chatbot-naranja/internal/state"
	"github.com/juanse1801/chatbot-naranja/pkg/models"
)

type Service interface {
	GetValidateEntry(ctx context.Context, token string) bool
	PostReceiveMessage(ctx context.Context, data domain.NewMessageReceived)
}

type service struct {
	itcService   interactionService.Service
	schService   schedulerService.Service
	stateService stateService.Service
	mssgService  messaging.Service
}

func NewService(itc interactionService.Service, sch schedulerService.Service, ste stateService.Service, mssg messaging.Service) Service {
	return &service{
		itcService:   itc,
		schService:   sch,
		stateService: ste,
		mssgService:  mssg,
	}
}

func (s *service) GetValidateEntry(ctx context.Context, token string) bool {
	dotEnv := os.Getenv("TOKEN")

	if token == dotEnv {
		return true
	}
	return false
}

func (s *service) PostReceiveMessage(ctx context.Context, data domain.NewMessageReceived) {
	clientNumber := data.Entry[0].Changes[0].Value.Messages[0].From
	message := data.Entry[0].Changes[0].Value.Messages[0].Text.Body

	// Reviso si es una nueva interacción
	interaction, err := s.itcService.GetInteraction(ctx, clientNumber)
	if err != nil {
		if err.Error() == "Not found" {
			fmt.Println("First interaction")
			itc, err := s.itcService.CreateInteraction(ctx, clientNumber)
			if err != nil {
				fmt.Println(err.Error())
			}
			newState, response, execute := s.stateService.NextState("Bienvenida", "")
			itc.State = newState
			s.executor(ctx, execute, itc, data)
			s.mssgService.SendMessage(clientNumber, response)
			return
		} else {
			log.Fatal(err)
			return
		}
	}

	// Reviso cual es el siguiente estado y el response del estado actual
	newState, response, execute := s.stateService.NextState(interaction.State, message)
	fmt.Println(newState)
	interaction.State = newState
	fmt.Println(interaction.State)
	s.executor(ctx, execute, interaction, data)

	// Service de mensajería
	s.mssgService.SendMessage(clientNumber, response)
	return
}

func (s *service) executor(ctx context.Context, execute string, itc models.InteractionModel, data domain.NewMessageReceived) {
	switch execute {
	case "update_state":
		{
			s.itcService.UpdateInteraction(ctx, itc)
			return
		}
	case "save_mail":
		{
			itc.ClientMail = data.Entry[0].Changes[0].Value.Messages[0].Text.Body
			s.itcService.UpdateInteraction(ctx, itc)
			return
		}
	case "send_mail":
		{
			s.itcService.UpdateInteraction(ctx, itc)
			// mailing service
			return
		}
	case "delete_interaction":
		{
			// delete service
			s.itcService.DeleteInteraction(ctx, itc.ClientNumber)
			return
		}
	}
}

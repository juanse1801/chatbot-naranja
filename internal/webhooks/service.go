package webhooks

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/juanse1801/chatbot-naranja/internal/domain"
	"github.com/juanse1801/chatbot-naranja/internal/history"
	interactionService "github.com/juanse1801/chatbot-naranja/internal/interaction"
	"github.com/juanse1801/chatbot-naranja/internal/mailing"
	"github.com/juanse1801/chatbot-naranja/internal/messaging"
	schedulerService "github.com/juanse1801/chatbot-naranja/internal/scheduler"
	stateService "github.com/juanse1801/chatbot-naranja/internal/state"
	"github.com/juanse1801/chatbot-naranja/pkg/jobs"
	"github.com/juanse1801/chatbot-naranja/pkg/models"
)

type Service interface {
	GetValidateEntry(ctx context.Context, token string) bool
	PostReceiveMessage(ctx context.Context, data domain.NewMessageReceived)
}

type service struct {
	itcService     interactionService.Service
	schService     schedulerService.Service
	stateService   stateService.Service
	mssgService    messaging.Service
	mailService    mailing.Service
	historyService history.Service
}

func NewService(itc interactionService.Service, sch schedulerService.Service, ste stateService.Service, mssg messaging.Service, mail mailing.Service, hs history.Service) Service {
	return &service{
		itcService:     itc,
		schService:     sch,
		stateService:   ste,
		mssgService:    mssg,
		mailService:    mail,
		historyService: hs,
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
			itc, err := s.itcService.CreateInteraction(ctx, clientNumber)
			if err != nil {
				fmt.Println(err.Error())
			}
			newState, response, execute := s.stateService.NextState("Bienvenida", "")
			itc.State = newState
			s.executor(ctx, execute, itc, message)
			s.mssgService.SendMessage(clientNumber, response)
			return
		} else {
			log.Fatal(err)
			return
		}
	}

	// Reviso cual es el siguiente estado y el response del estado actual
	newState, response, execute := s.stateService.NextState(interaction.State, message)
	interaction.State = newState
	go s.executor(ctx, execute, interaction, message)

	// Service de mensajería
	go s.mssgService.SendMessage(clientNumber, response)
	return
}

func (s *service) executor(ctx context.Context, execute string, itc models.InteractionModel, data string) {
	switch execute {
	case "update_state":
		{

			go s.schService.ScheduleExpiration(ctx, itc.ClientNumber, jobs.GetNextTime(), jobs.GetSecondTime())
			go s.itcService.UpdateInteraction(ctx, itc)
			return
		}
	case "update_entity":
		{

			itc.Entity = data
			go s.schService.ScheduleExpiration(ctx, itc.ClientNumber, jobs.GetNextTime(), jobs.GetSecondTime())
			go s.itcService.UpdateInteraction(ctx, itc)
			return
		}
	case "update_service":
		{

			itc.Service = data
			go s.schService.ScheduleExpiration(ctx, itc.ClientNumber, jobs.GetNextTime(), jobs.GetSecondTime())
			go s.itcService.UpdateInteraction(ctx, itc)
			return
		}
	case "update_type":
		{

			itc.Type = data
			go s.schService.ScheduleExpiration(ctx, itc.ClientNumber, jobs.GetNextTime(), jobs.GetSecondTime())
			go s.itcService.UpdateInteraction(ctx, itc)
			return
		}
	case "update_zone":
		{

			itc.Zone = data
			go s.schService.ScheduleExpiration(ctx, itc.ClientNumber, jobs.GetNextTime(), jobs.GetSecondTime())
			go s.itcService.UpdateInteraction(ctx, itc)
			return
		}
	case "delete_interaction":
		{
			go s.schService.DeleteExpiration(itc.ClientNumber)
			go s.itcService.DeleteInteraction(ctx, itc.ClientNumber)
			return
		}
	case "end_interaction":
		{
			if itc.Service == "0" {
				go s.mailService.SendEmail("0", itc.Zone, data)
			}
			if itc.Service == "1" {
				go s.mailService.SendEmail("1", "default", data)
			}
			go s.historyService.CreateHistory(ctx, itc, data)
			go s.schService.DeleteExpiration(itc.ClientNumber)
			go s.itcService.DeleteInteraction(ctx, itc.ClientNumber)
			return
		}
	}
}

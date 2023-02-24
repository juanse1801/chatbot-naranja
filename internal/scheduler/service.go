package scheduler

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/juanse1801/chatbot-naranja/internal/interaction"
	"github.com/juanse1801/chatbot-naranja/internal/messaging"
)

type Service interface {
	ScheduleExpiration(ctx context.Context, tag string, firstTime string, secondTime string) error
	DeleteExpiration(tag string) error
}

type service struct {
	scheduler  *gocron.Scheduler
	msgService messaging.Service
	itcService interaction.Service
}

func NewService(sch *gocron.Scheduler, msg messaging.Service, itc interaction.Service) Service {
	return &service{
		scheduler:  sch,
		msgService: msg,
		itcService: itc,
	}
}

func (sch *service) ScheduleExpiration(ctx context.Context, tag string, firstTime string, secondTime string) error {
	_, err := sch.scheduler.FindJobsByTag(tag)
	if err != nil {
		rcJob, _ := sch.scheduler.Cron(firstTime).Do(func() {
			sch.msgService.SendRecontactMessage(tag)
		})
		gbJob, _ := sch.scheduler.Cron(secondTime).Do(func() {
			sch.msgService.SendNoContactMessage(tag)
			sch.itcService.DeleteInteraction(ctx, tag)
		})

		rcJob.Tag(tag)
		gbJob.Tag(tag)
		return nil
	}

	err = sch.scheduler.RemoveByTag(tag)

	if err != nil {
		return err
	}

	rcJob, _ := sch.scheduler.Cron(firstTime).Do(func() {
		sch.msgService.SendRecontactMessage(tag)
	})
	gbJob, _ := sch.scheduler.Cron(secondTime).Do(func() {
		sch.msgService.SendNoContactMessage(tag)
		sch.itcService.DeleteInteraction(ctx, tag)
	})

	rcJob.Tag(tag)
	gbJob.Tag(tag)

	return nil

}

func (sch *service) DeleteExpiration(tag string) error {
	err := sch.scheduler.RemoveByTag(tag)

	if err != nil {
		return err
	}

	return nil
}

package scheduler

import (
	"github.com/go-co-op/gocron"
	"github.com/juanse1801/chatbot-naranja/pkg/jobs"
)

type Service interface {
	ScheduleExpiration(tag string, time string) error
	UpdateExpiration(tag string, time string) error
	DeleteExpiration(tag string) error
}

type service struct {
	scheduler *gocron.Scheduler
}

func NewService(sch *gocron.Scheduler) Service {
	return &service{
		scheduler: sch,
	}
}

func (sch *service) ScheduleExpiration(tag string, time string) error {
	job, _ := sch.scheduler.Cron(time).Do(func() {
		jobs.RecontactJob()
	})
	job.Tag(tag)
	return nil
}

func (sch *service) UpdateExpiration(tag string, time string) error {
	err := sch.scheduler.RemoveByTag(tag)

	if err != nil {
		return err
	}

	job, _ := sch.scheduler.Cron(time).Do(func() {
		jobs.RecontactJob()
	})
	job.Tag(tag)
	return nil
}

func (sch *service) DeleteExpiration(tag string) error {
	err := sch.scheduler.RemoveByTag(tag)

	if err != nil {
		return err
	}

	return nil
}

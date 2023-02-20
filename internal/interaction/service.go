package interaction

import (
	"context"
	"fmt"

	"github.com/juanse1801/chatbot-naranja/pkg/models"
)

type Service interface {
	GetInteraction(ctx context.Context, clientNumber string) (models.InteractionModel, error)
	CreateInteraction(ctx context.Context, clientNumber string) (models.InteractionModel, error)
	UpdateInteraction(ctx context.Context, itc models.InteractionModel) (models.InteractionModel, error)
	DeleteInteraction(ctx context.Context, clientNumber string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetInteraction(ctx context.Context, clientNumber string) (models.InteractionModel, error) {
	itc, err := s.repository.Get(ctx, clientNumber)
	if err != nil {
		return models.InteractionModel{}, err
	}

	return itc, nil
}

func (s *service) CreateInteraction(ctx context.Context, clientNumber string) (models.InteractionModel, error) {
	newInteraction, err := s.repository.Save(ctx, clientNumber)
	if err != nil {
		return models.InteractionModel{}, err
	}

	return newInteraction, err

}

func (s *service) UpdateInteraction(ctx context.Context, itc models.InteractionModel) (models.InteractionModel, error) {
	updateItc, err := s.repository.Update(ctx, itc)

	if err != nil {
		fmt.Println(err.Error())
		return models.InteractionModel{}, err
	}

	return updateItc, nil
}

func (s *service) DeleteInteraction(ctx context.Context, clientNumber string) error {
	err := s.repository.Delete(ctx, clientNumber)
	if err != nil {
		return err
	}

	return nil

}

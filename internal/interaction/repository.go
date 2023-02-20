package interaction

import (
	"context"
	"errors"
	"time"

	"github.com/juanse1801/chatbot-naranja/pkg/configs"
	"github.com/juanse1801/chatbot-naranja/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Get(ctx context.Context, clientNumber string) (models.InteractionModel, error)
	Save(ctx context.Context, clientNumber string) (models.InteractionModel, error)
	Update(ctx context.Context, itc models.InteractionModel) (models.InteractionModel, error)
	Delete(ctx context.Context, clientNumber string) error
}

type repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, clientNumber string) (models.InteractionModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var interaction models.InteractionModel

	collection := configs.GetCollection(r.db, "interactions")

	err := collection.FindOne(ctx, bson.M{"clientnumber": clientNumber}).Decode(&interaction)
	if err == mongo.ErrNoDocuments {
		return models.InteractionModel{}, errors.New("Not found")
	}

	if err != nil {
		return models.InteractionModel{}, errors.New(err.Error())
	}

	return interaction, nil
}

func (r *repository) Save(ctx context.Context, clientNumber string) (models.InteractionModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newInteraction := models.InteractionModel{
		State:        "Bienvenida",
		ClientNumber: clientNumber,
		ClientMail:   "",
	}

	collection := configs.GetCollection(r.db, "interactions")

	_, err := collection.InsertOne(ctx, newInteraction)

	if err != nil {
		return models.InteractionModel{}, err
	}

	return newInteraction, nil

}

func (r *repository) Update(ctx context.Context, itc models.InteractionModel) (models.InteractionModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := configs.GetCollection(r.db, "interactions")

	_, err := collection.UpdateOne(ctx, bson.M{"clientnumber": itc.ClientNumber}, bson.M{"$set": itc})

	if err != nil {
		return models.InteractionModel{}, err
	}

	return itc, nil

}

func (r *repository) Delete(ctx context.Context, clientNumber string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := configs.GetCollection(r.db, "interactions")

	_, err := collection.DeleteOne(ctx, bson.M{"clientnumber": clientNumber})

	if err != nil {
		return err
	}
	return nil
}

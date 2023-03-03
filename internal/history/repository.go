package history

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/juanse1801/chatbot-naranja/pkg/configs"
	"github.com/juanse1801/chatbot-naranja/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var months = map[string]string{
	"January":   "01",
	"February":  "02",
	"March":     "03",
	"April":     "04",
	"May":       "05",
	"June":      "06",
	"July":      "07",
	"August":    "08",
	"September": "09",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}

type Repository interface {
	Get(date string) ([]models.HistoryModel, error)
	Save(ctx context.Context, itc models.InteractionModel, data string) error
}

type repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(date string) ([]models.HistoryModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var allHistory []models.HistoryModel

	collection := configs.GetCollection(r.db, "history")

	data, err := collection.Find(ctx, bson.M{"date": date})
	if err != nil {
		return []models.HistoryModel{}, errors.New(err.Error())
	}

	if err = data.All(ctx, &allHistory); err != nil {
		fmt.Println(err.Error())
	}

	return allHistory, nil
}

func (r *repository) Save(ctx context.Context, itc models.InteractionModel, data string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newHistory := models.HistoryModel{
		ClientNumber: itc.ClientNumber,
		Entity:       itc.Entity,
		Service:      itc.Service,
		Type:         itc.Type,
		Zone:         itc.Zone,
		Data:         data,
		Date:         getTime(),
	}

	collection := configs.GetCollection(r.db, "history")

	_, err := collection.InsertOne(ctx, newHistory)

	if err != nil {
		return err
	}

	return nil

}

func FormatDate(num int) string {
	if num < 10 {
		return fmt.Sprintf("0%d", num)
	}
	return fmt.Sprintf("%d", num)
}

func getTime() string {
	nowTime := time.Now().UTC().Add(-time.Hour * 5)
	return fmt.Sprintf("%d-%s-%s", nowTime.Year(), months[nowTime.Month().String()], FormatDate(nowTime.Day()))
}

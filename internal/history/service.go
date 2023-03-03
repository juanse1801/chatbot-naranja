package history

import (
	"context"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/juanse1801/chatbot-naranja/pkg/configs"
	"github.com/juanse1801/chatbot-naranja/pkg/models"
)

const Sheet1 = "Asesoria Comercial"
const Sheet2 = "Administrativo"

type Service interface {
	GetHistory(ctx context.Context, date string) (*excelize.File, error)
	CreateHistory(ctx context.Context, itc models.InteractionModel, data string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetHistory(ctx context.Context, date string) (*excelize.File, error) {
	allInteractions, err := s.repository.Get(date)

	if err != nil {
		fmt.Println(err.Error())
	}
	defaultSheet1 := 2
	defaultSheet2 := 2

	f := excelize.NewFile()

	f.NewSheet(Sheet1)
	f.NewSheet(Sheet2)

	// Set del primer sheet
	f.SetCellValue(Sheet1, "A1", "Numero del cliente")
	f.SetCellValue(Sheet1, "B1", "Entidad")
	f.SetCellValue(Sheet1, "C1", "Servicio")
	f.SetCellValue(Sheet1, "D1", "Producto")
	f.SetCellValue(Sheet1, "E1", "Zona")
	f.SetCellValue(Sheet1, "F1", "Información del cliente")
	f.SetCellValue(Sheet1, "G1", "Enviado A")

	// Set del segundo sheet
	f.SetCellValue(Sheet2, "A1", "Numero del cliente")
	f.SetCellValue(Sheet2, "B1", "Entidad")
	f.SetCellValue(Sheet2, "C1", "Servicio")
	f.SetCellValue(Sheet2, "D1", "Información del cliente")
	f.SetCellValue(Sheet2, "E1", "Enviado A")

	for _, data := range allInteractions {
		if data.Service == "0" {
			f.SetCellValue(Sheet1, "A"+strconv.Itoa(defaultSheet1), data.ClientNumber)
			f.SetCellValue(Sheet1, "B"+strconv.Itoa(defaultSheet1), configs.Entidades[data.Entity])
			f.SetCellValue(Sheet1, "C"+strconv.Itoa(defaultSheet1), configs.Servicios[data.Service])
			f.SetCellValue(Sheet1, "D"+strconv.Itoa(defaultSheet1), configs.Producto[data.Type])
			f.SetCellValue(Sheet1, "E"+strconv.Itoa(defaultSheet1), configs.Zona[data.Zone])
			f.SetCellValue(Sheet1, "F"+strconv.Itoa(defaultSheet1), data.Data)
			f.SetCellValue(Sheet1, "G"+strconv.Itoa(defaultSheet1), configs.EmailsConfig["0"][data.Zone])

			defaultSheet1 = defaultSheet1 + 1
		}
		if data.Service == "1" {
			f.SetCellValue(Sheet2, "A"+strconv.Itoa(defaultSheet2), "Numero del cliente")
			f.SetCellValue(Sheet2, "B"+strconv.Itoa(defaultSheet2), "Entidad")
			f.SetCellValue(Sheet2, "C"+strconv.Itoa(defaultSheet2), "Servicio")
			f.SetCellValue(Sheet2, "D"+strconv.Itoa(defaultSheet2), "Información del cliente")
			f.SetCellValue(Sheet2, "E"+strconv.Itoa(defaultSheet2), "Enviado A")

			defaultSheet2 = defaultSheet2 + 1
		}
	}

	return f, nil

}

func (s *service) CreateHistory(ctx context.Context, itc models.InteractionModel, data string) error {
	err := s.repository.Save(ctx, itc, data)

	if err != nil {
		return err
	}

	return nil
}

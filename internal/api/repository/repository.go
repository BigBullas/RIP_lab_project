package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"RIP_lab1/internal/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.FlightRequest{})
	if err != nil {
		panic("Миграция БД не удалась")
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetRequestForFlightList(substring string) ([]models.FlightRequest, error) {
	var request_for_delivery []models.FlightRequest

	r.db.Where("title ILIKE ?", "%"+substring+"%").Find(&request_for_delivery, "is_available = ?", true)
	return request_for_delivery, nil
}

func (r *Repository) GetCardRequestForFlightById(cardId int) (models.FlightRequest, error) {
	var card models.FlightRequest

	r.db.Where("request_id = ?", cardId).Find(&card, "is_available = ?", true)
	return card, nil
}

func (r *Repository) DeleteRequestForFlightById(cardId int) error {
	err := r.db.Exec("UPDATE flight_requests SET is_available=false WHERE request_id = ?", cardId).Error
	if err != nil {
		return err
	}
	return nil
}

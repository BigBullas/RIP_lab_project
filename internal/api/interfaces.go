package api

import (
	"RIP_lab1/internal/models"
	"time"
)

type Repo interface {
	GetRequestForFlightList(string) ([]models.FlightRequest, error)
	GetCardRequestForFlightById(int) (models.FlightRequest, error)
	CreateNewRequestForFlight(models.FlightRequest) error
	ChangeRequestForFlight(models.FlightRequest) error
	DeleteRequestForFlightById(int) error
	GetFlightRequestImageUrl(int) string

	AddFlightRequestToFlight(models.RocketFlightShort) error
	DeleteRequestFromFlight(int, int) error
	ChangeCountFlightsFlightRequest(int, int, int) error

	GetRocketFlightList(time.Time, time.Time, string) ([]models.RocketFlight, error)
	GetRocketFlightDraft(int) (int, error)
	GetRocketFlightById(int) (models.RocketFlightDetailed, []models.FlightRequest, error)
	ChangeRocketFlight(models.RocketFlightChangeable) error
	FormRocketFlight(models.RocketFlight) error
	ResponceRocketFlight(models.RocketFlight) error
	DeleteRocketFlight(int) error
}

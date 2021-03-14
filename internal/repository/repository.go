package repository

import "github.com/petrostrak/booking-with-go/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}

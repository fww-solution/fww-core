package usecase

import (
	"errors"
	"fww-core/internal/data/dto_ticket"
	"fww-core/internal/entity"
	"time"

	"log"

	"github.com/google/uuid"
)

// RedeemTicket implements UseCase.
func (u *useCase) RedeemTicket(codeBooking string) (dto_ticket.Response, error) {

	// Find Booking by code
	booking, err := u.repository.FindBookingByBookingIDCode(codeBooking)
	if err != nil {
		return dto_ticket.Response{}, err
	}

	if booking.ID == 0 {
		return dto_ticket.Response{}, errors.New("booking not found")
	}

	// Validate Booking Date
	if booking.BookingExpiredAt.Before(time.Now()) {
		return dto_ticket.Response{}, errors.New("booking expired")
	}

	// Update Booking Status
	booking.BookingStatus = "redeemed"
	_, err = u.repository.UpdateBooking(&booking)
	if err != nil {
		return dto_ticket.Response{}, err
	}

	// Upsert Ticket
	generatedUUID := uuid.New()

	entityTicket := entity.Ticket{
		ID:                 0,
		CodeTicket:         generatedUUID.String(),
		IsBoardingPass:     false,
		IsEligibleToFlight: false,
		BookingID:          booking.ID,
	}

	_, err = u.repository.UpsertTicket(&entityTicket)
	if err != nil {
		log.Println(err)
		return dto_ticket.Response{}, err
	}

	boardingTime := booking.BookingExpiredAt.Add((24 * time.Hour) - (time.Minute * 30))

	response := dto_ticket.Response{
		CodeTicket:  entityTicket.CodeTicket,
		BordingTime: boardingTime.Format("2006-01-02 15:04:05"),
	}

	return response, nil

}
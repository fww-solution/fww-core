package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"fww-core/internal/container/infrastructure/redis"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/entity"
	"fww-core/internal/tools"
	"time"
)

// RequestBooking implements UseCase.
func (u *useCase) RequestBooking(data *dto_booking.Request, bookingIDCode string) error {
	ctx := context.Background()
	// Check Booking ID Code
	resultBooking, err := u.repository.FindBookingByBookingIDCode(bookingIDCode)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	if resultBooking.ID != 0 {
		return tools.ErrorBuilder(errors.New("booking id code already exist"))
	}

	var reminingSeat int

	// Check Remining Seat
	flightIDReminingSeat := fmt.Sprintf("flight-%d-seat", data.FlightID)
	result := u.redis.Get(ctx, flightIDReminingSeat)
	if result.Err() != nil {
		reminingSeat, err := u.repository.FindReminingSeat(data.FlightID)
		if err != nil {
			return tools.ErrorBuilder(err)
		}
		if reminingSeat <= 0 {
			return errors.New("no remaning seat")
		}
		u.redis.Set(ctx, flightIDReminingSeat, reminingSeat, 0)
	}

	reminingSeat, err = result.Int()
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	if reminingSeat < 1 {
		return errors.New("no remaning seat")
	}

	flightIDKey := fmt.Sprintf("flight-%d", data.FlightID)

	// Lock Transaction Redis
	rc := redis.InitMutex(flightIDKey)
	//nolint
	redis.LockMutex(rc)
	//nolint
	defer redis.UnlockMutex(rc)

	// Check Remining Seat

	// find flight
	resultFlight, err := u.repository.FindFlightByID(data.FlightID)
	if err != nil {
		return err
	}
	bookingDate := time.Now().Round(time.Minute)
	paymentExpiredAt := time.Now().Add(time.Hour * 6).Round(time.Minute)
	bookingExpiredAt := resultFlight.DepartureTime.AddDate(0, 0, -1).Round(time.Minute)
	// payment expired at 6 hour after booking

	// Insert Booking
	bookingEntity := &entity.Booking{
		CodeBooking:      bookingIDCode,
		BookingDate:      bookingDate,
		PaymentExpiredAt: paymentExpiredAt,
		BookingExpiredAt: bookingExpiredAt,
		BookingStatus:    "pending",
		CaseID:           0,
		UserID:           data.UserID,
		FlightID:         data.FlightID,
	}

	bookingID, err := u.repository.InsertBooking(bookingEntity)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	// Update Flight Reservation
	entityReservation := &entity.FlightReservation{
		Class:        data.BookDetails[0].Class,
		ReminingSeat: reminingSeat - len(data.BookDetails),
		TotalSeat:    172,
		UpdatedAt: sql.NullTime{
			Time:  time.Now().Round(time.Minute),
			Valid: true,
		},
		FlightID: data.FlightID,
	}
	_, err = u.repository.UpdateFlightReservation(entityReservation)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	status := u.redis.Set(ctx, flightIDReminingSeat, reminingSeat-len(data.BookDetails), 0)
	if status.Err() != nil {
		return tools.ErrorBuilder(status.Err())
	}

	// Insert Booking Detail
	for _, v := range data.BookDetails {

		entityBookingDetail := entity.BookingDetail{
			BookingID:       bookingID,
			PassengerID:     v.PassangerID,
			SeatNumber:      v.SeatNumber,
			BaggageCapacity: v.Baggage,
			Class:           v.Class,
		}

		_, err := u.repository.InsertBookingDetail(&entityBookingDetail)
		if err != nil {
			return tools.ErrorBuilder(err)
		}
	}

	specBooking := dto_booking.RequestBPM{
		CodeBooking:    bookingIDCode,
		PaymentExpired: paymentExpiredAt,
	}

	err = u.adapter.RequestGenerateInvoice(&specBooking)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	return nil

}

// GetDetailBooking implements UseCase.
func (u *useCase) GetDetailBooking(codeBooking string) (dto_booking.BookResponse, error) {
	result, err := u.repository.FindBookingByBookingIDCode(codeBooking)
	if err != nil {
		return dto_booking.BookResponse{}, tools.ErrorBuilder(err)
	}

	if result.ID == 0 {
		return dto_booking.BookResponse{}, errors.New("booking not found")
	}

	resultBookingDetails, err := u.repository.FindBookingDetailByBookingID(result.ID)
	if err != nil {
		return dto_booking.BookResponse{}, tools.ErrorBuilder(err)
	}

	resultFlight, err := u.repository.FindFlightByID(result.FlightID)
	if err != nil {
		return dto_booking.BookResponse{}, tools.ErrorBuilder(err)
	}

	if resultFlight.ID == 0 {
		return dto_booking.BookResponse{}, errors.New("flight not found")
	}

	resultFlightPrice, err := u.repository.FindFlightPriceByID(result.FlightID)
	if err != nil {
		return dto_booking.BookResponse{}, tools.ErrorBuilder(err)
	}

	if resultFlightPrice.ID == 0 {
		return dto_booking.BookResponse{}, errors.New("flight price not found")
	}

	// booking expired at resultFlight i day before DepartureTime
	bookingExpiredAt := resultFlight.DepartureTime.AddDate(0, 0, -1)

	var bookDetails []dto_booking.BookResponseDetail

	for _, v := range resultBookingDetails {
		passenger, err := u.repository.FindDetailPassanger(v.PassengerID)
		if err != nil {
			return dto_booking.BookResponse{}, tools.ErrorBuilder(err)
		}

		bookDetails = append(bookDetails, dto_booking.BookResponseDetail{
			Baggage:       v.BaggageCapacity,
			SeatNumber:    v.SeatNumber,
			PassangerName: passenger.FullName,
			Class:         v.Class,
			Price:         resultFlightPrice.Price,
		})
	}

	bookResponse := dto_booking.BookResponse{
		ArrivalAirport:   resultFlight.ArrivalAirportName,
		ArrivalTime:      resultFlight.ArrivalTime.Round(time.Minute).Format("2006-01-02 15:01:00"),
		BookExpiredAt:    bookingExpiredAt.Round(time.Minute).Format("2006-01-02 15:02:00"),
		CodeBooking:      result.CodeBooking,
		CodeFlight:       resultFlight.CodeFlight,
		DepartureAirport: resultFlight.DepartureAirportName,
		DepartureTime:    resultFlight.DepartureTime.Round(time.Minute).Format("2006-01-02 15:03:00"),
		Details:          bookDetails,
		ID:               result.ID,
		PaymentExpiredAt: result.PaymentExpiredAt.Round(time.Minute).Format("2006-01-02 15:04:00"),
		TotalPrice:       resultFlightPrice.Price,
	}

	return bookResponse, nil

}

// UpdateDetailBooking implements UseCase.
func (u *useCase) UpdateDetailBooking(data *dto_booking.BookDetailRequest) error {
	resultBookingDetail, err := u.repository.FindBookingDetailByID(data.BookingDetailID)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	if resultBookingDetail.ID == 0 {
		return errors.New("booking detail not found")
	}

	// Update Booking Detail
	resultBookingDetail.IsEligibleToFlight = data.IsEligibleToFlight
	_, err = u.repository.UpdateBookingDetail(&resultBookingDetail)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	return nil
}

// UpdateBooking implements UseCase.
func (u *useCase) UpdateBooking(req *dto_booking.RequestUpdateBooking) error {
	resultBooking, err := u.repository.FindBookingByBookingIDCode(req.CodeBooking)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	if resultBooking.ID == 0 {
		return errors.New("booking not found")
	}

	resultBooking.BookingStatus = req.Status
	_, err = u.repository.UpdateBooking(&resultBooking)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	return nil
}

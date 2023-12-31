package repository

import (
	"fww-core/internal/data/dto_notification"
	"fww-core/internal/entity"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type Repository interface {
	FindDetailPassanger(id int64) (entity.Passenger, error)
	RegisterPassanger(data *entity.Passenger) (int64, error)
	UpdatePassanger(data *entity.Passenger) (int64, error)
	FindPassangerByIDNumber(idNumber string) (entity.Passenger, error)

	// Airport
	FindAirport(city string, province string, iata string) ([]entity.Airport, error)

	// Flight
	FindFlightByID(id int64) (entity.Flight, error)
	FindFlights(departureTime string, arrivalTime string, limit int, offset int) ([]entity.Flight, error)
	FindFlightPriceByID(id int64) (entity.FlightPrice, error)
	FindFlightReservationByID(flightID int64) (entity.FlightReservation, error)
	// Booking
	FindReminingSeat(flightID int64) (int, error)
	InsertBooking(data *entity.Booking) (int64, error)
	UpdateBooking(data *entity.Booking) (int64, error)
	InsertBookingDetail(data *entity.BookingDetail) (int64, error)
	UpdateFlightReservation(data *entity.FlightReservation) (int64, error)
	FindBookingByBookingIDCode(bookingIDCode string) (entity.Booking, error)
	FindBookingDetailByBookingID(bookingID int64) ([]entity.BookingDetail, error)
	FindBookingDetailByID(id int64) (entity.BookingDetail, error)
	UpdateBookingDetail(data *entity.BookingDetail) (int64, error)
	FindBookingByID(id int64) (entity.Booking, error)
	// Payment
	FindPaymentDetailByInvoice(invoiceNumber string) (entity.Payment, error)
	UpsertPayment(data *entity.Payment) (int64, error)
	FindPaymentMethodStatus() ([]entity.PaymentMethod, error)
	FindPaymentByBookingID(bookingID int64) (entity.Payment, error)
	// Ticket
	UpsertTicket(data *entity.Ticket) (int64, error)
	FindTicketByCodeTicket(codeTicket string) (entity.Ticket, error)
	// Notification
	PaymentInvoiceReportByBookingCode(bookingCode string) (dto_notification.PaymentInvoiceAggregator, error)
	PaymentReceiptReportByBookingCode(bookingCode string) (dto_notification.PaymentReceiptAggregator, error)
	TicketRedeemedReportByBookingCode(bookingCode string) (dto_notification.TicketRedeemAgregator, error)
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	entity "fww-core/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindAirport provides a mock function with given fields: city, province, iata
func (_m *Repository) FindAirport(city string, province string, iata string) ([]entity.Airport, error) {
	ret := _m.Called(city, province, iata)

	var r0 []entity.Airport
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) ([]entity.Airport, error)); ok {
		return rf(city, province, iata)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) []entity.Airport); ok {
		r0 = rf(city, province, iata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Airport)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(city, province, iata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBookingByBookingIDCode provides a mock function with given fields: bookingIDCode
func (_m *Repository) FindBookingByBookingIDCode(bookingIDCode string) (entity.Booking, error) {
	ret := _m.Called(bookingIDCode)

	var r0 entity.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.Booking, error)); ok {
		return rf(bookingIDCode)
	}
	if rf, ok := ret.Get(0).(func(string) entity.Booking); ok {
		r0 = rf(bookingIDCode)
	} else {
		r0 = ret.Get(0).(entity.Booking)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bookingIDCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBookingByID provides a mock function with given fields: id
func (_m *Repository) FindBookingByID(id int64) (entity.Booking, error) {
	ret := _m.Called(id)

	var r0 entity.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (entity.Booking, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) entity.Booking); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Booking)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBookingDetailByBookingID provides a mock function with given fields: bookingID
func (_m *Repository) FindBookingDetailByBookingID(bookingID int64) ([]entity.BookingDetail, error) {
	ret := _m.Called(bookingID)

	var r0 []entity.BookingDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) ([]entity.BookingDetail, error)); ok {
		return rf(bookingID)
	}
	if rf, ok := ret.Get(0).(func(int64) []entity.BookingDetail); ok {
		r0 = rf(bookingID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.BookingDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(bookingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindDetailPassanger provides a mock function with given fields: id
func (_m *Repository) FindDetailPassanger(id int64) (entity.Passenger, error) {
	ret := _m.Called(id)

	var r0 entity.Passenger
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (entity.Passenger, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) entity.Passenger); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Passenger)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFlightByID provides a mock function with given fields: id
func (_m *Repository) FindFlightByID(id int64) (entity.Flight, error) {
	ret := _m.Called(id)

	var r0 entity.Flight
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (entity.Flight, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) entity.Flight); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Flight)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFlightPriceByID provides a mock function with given fields: id
func (_m *Repository) FindFlightPriceByID(id int64) (entity.FlightPrice, error) {
	ret := _m.Called(id)

	var r0 entity.FlightPrice
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (entity.FlightPrice, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) entity.FlightPrice); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.FlightPrice)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFlightReservationByID provides a mock function with given fields: flightID
func (_m *Repository) FindFlightReservationByID(flightID int64) (entity.FlightReservation, error) {
	ret := _m.Called(flightID)

	var r0 entity.FlightReservation
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (entity.FlightReservation, error)); ok {
		return rf(flightID)
	}
	if rf, ok := ret.Get(0).(func(int64) entity.FlightReservation); ok {
		r0 = rf(flightID)
	} else {
		r0 = ret.Get(0).(entity.FlightReservation)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(flightID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFlights provides a mock function with given fields: departureTime, arrivalTime, limit, offset
func (_m *Repository) FindFlights(departureTime string, arrivalTime string, limit int, offset int) ([]entity.Flight, error) {
	ret := _m.Called(departureTime, arrivalTime, limit, offset)

	var r0 []entity.Flight
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]entity.Flight, error)); ok {
		return rf(departureTime, arrivalTime, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []entity.Flight); ok {
		r0 = rf(departureTime, arrivalTime, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Flight)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(departureTime, arrivalTime, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPassangerByIDNumber provides a mock function with given fields: idNumber
func (_m *Repository) FindPassangerByIDNumber(idNumber string) (entity.Passenger, error) {
	ret := _m.Called(idNumber)

	var r0 entity.Passenger
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.Passenger, error)); ok {
		return rf(idNumber)
	}
	if rf, ok := ret.Get(0).(func(string) entity.Passenger); ok {
		r0 = rf(idNumber)
	} else {
		r0 = ret.Get(0).(entity.Passenger)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPaymentByBookingID provides a mock function with given fields: bookingID
func (_m *Repository) FindPaymentByBookingID(bookingID int64) (entity.Payment, error) {
	ret := _m.Called(bookingID)

	var r0 entity.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (entity.Payment, error)); ok {
		return rf(bookingID)
	}
	if rf, ok := ret.Get(0).(func(int64) entity.Payment); ok {
		r0 = rf(bookingID)
	} else {
		r0 = ret.Get(0).(entity.Payment)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(bookingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPaymentDetailByInvoice provides a mock function with given fields: invoiceNumber
func (_m *Repository) FindPaymentDetailByInvoice(invoiceNumber string) (entity.Payment, error) {
	ret := _m.Called(invoiceNumber)

	var r0 entity.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.Payment, error)); ok {
		return rf(invoiceNumber)
	}
	if rf, ok := ret.Get(0).(func(string) entity.Payment); ok {
		r0 = rf(invoiceNumber)
	} else {
		r0 = ret.Get(0).(entity.Payment)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(invoiceNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPaymentMethodStatus provides a mock function with given fields:
func (_m *Repository) FindPaymentMethodStatus() ([]entity.PaymentMethod, error) {
	ret := _m.Called()

	var r0 []entity.PaymentMethod
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.PaymentMethod, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.PaymentMethod); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.PaymentMethod)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindReminingSeat provides a mock function with given fields: flightID
func (_m *Repository) FindReminingSeat(flightID int64) (int, error) {
	ret := _m.Called(flightID)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (int, error)); ok {
		return rf(flightID)
	}
	if rf, ok := ret.Get(0).(func(int64) int); ok {
		r0 = rf(flightID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(flightID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTicketByCodeTicket provides a mock function with given fields: codeTicket
func (_m *Repository) FindTicketByCodeTicket(codeTicket string) (entity.Ticket, error) {
	ret := _m.Called(codeTicket)

	var r0 entity.Ticket
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.Ticket, error)); ok {
		return rf(codeTicket)
	}
	if rf, ok := ret.Get(0).(func(string) entity.Ticket); ok {
		r0 = rf(codeTicket)
	} else {
		r0 = ret.Get(0).(entity.Ticket)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(codeTicket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBooking provides a mock function with given fields: data
func (_m *Repository) InsertBooking(data *entity.Booking) (int64, error) {
	ret := _m.Called(data)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Booking) (int64, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*entity.Booking) int64); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*entity.Booking) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBookingDetail provides a mock function with given fields: data
func (_m *Repository) InsertBookingDetail(data *entity.BookingDetail) (int64, error) {
	ret := _m.Called(data)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.BookingDetail) (int64, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*entity.BookingDetail) int64); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*entity.BookingDetail) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterPassanger provides a mock function with given fields: data
func (_m *Repository) RegisterPassanger(data *entity.Passenger) (int64, error) {
	ret := _m.Called(data)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Passenger) (int64, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*entity.Passenger) int64); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*entity.Passenger) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBooking provides a mock function with given fields: data
func (_m *Repository) UpdateBooking(data *entity.Booking) (int64, error) {
	ret := _m.Called(data)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Booking) (int64, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*entity.Booking) int64); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*entity.Booking) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFlightReservation provides a mock function with given fields: data
func (_m *Repository) UpdateFlightReservation(data *entity.FlightReservation) (int64, error) {
	ret := _m.Called(data)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.FlightReservation) (int64, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*entity.FlightReservation) int64); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*entity.FlightReservation) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassanger provides a mock function with given fields: data
func (_m *Repository) UpdatePassanger(data *entity.Passenger) (int64, error) {
	ret := _m.Called(data)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Passenger) (int64, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*entity.Passenger) int64); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*entity.Passenger) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertPayment provides a mock function with given fields: data
func (_m *Repository) UpsertPayment(data *entity.Payment) (int64, error) {
	ret := _m.Called(data)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Payment) (int64, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*entity.Payment) int64); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*entity.Payment) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertTicket provides a mock function with given fields: data
func (_m *Repository) UpsertTicket(data *entity.Ticket) (int64, error) {
	ret := _m.Called(data)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Ticket) (int64, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*entity.Ticket) int64); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*entity.Ticket) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

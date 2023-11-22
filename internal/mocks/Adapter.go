// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	dto_booking "fww-core/internal/data/dto_booking"
	dto_notification "fww-core/internal/data/dto_notification"

	dto_passanger "fww-core/internal/data/dto_passanger"

	dto_payment "fww-core/internal/data/dto_payment"

	dto_ticket "fww-core/internal/data/dto_ticket"

	mock "github.com/stretchr/testify/mock"
)

// Adapter is an autogenerated mock type for the Adapter type
type Adapter struct {
	mock.Mock
}

// CheckPassangerInformations provides a mock function with given fields: data
func (_m *Adapter) CheckPassangerInformations(data *dto_passanger.RequestBPM) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_passanger.RequestBPM) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DoPayment provides a mock function with given fields: data
func (_m *Adapter) DoPayment(data *dto_payment.DoPayment) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_payment.DoPayment) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedeemTicket provides a mock function with given fields: data
func (_m *Adapter) RedeemTicket(data *dto_ticket.RequestRedeemTicketToBPM) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_ticket.RequestRedeemTicketToBPM) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestGenerateInvoice provides a mock function with given fields: data
func (_m *Adapter) RequestGenerateInvoice(data *dto_booking.RequestBPM) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_booking.RequestBPM) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendNotification provides a mock function with given fields: data
func (_m *Adapter) SendNotification(data *dto_notification.SendEmailRequest) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_notification.SendEmailRequest) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAdapter creates a new instance of Adapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdapter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Adapter {
	mock := &Adapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

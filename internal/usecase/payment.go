package usecase

import (
	"errors"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/entity"
	"fww-core/internal/tools"
	"time"
)

// GenerateInvoice implements UseCase.
func (u *useCase) GenerateInvoice(caseID int64, codeBooking string) error {
	resultBooking, err := u.repository.FindBookingByBookingIDCode(codeBooking)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	if resultBooking.ID == 0 {
		return errors.New("booking not found")
	}
	// Total payment from booking details and flight price
	totalPayment := float64(0)

	bookingDetails, err := u.repository.FindBookingDetailByBookingID(resultBooking.ID)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	bookingPrice, err := u.repository.FindFlightPriceByID(resultBooking.FlightID)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	// Update Booking Case ID
	resultBooking.CaseID = caseID
	_, err = u.repository.UpdateBooking(&resultBooking)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	totalPayment += (bookingPrice.Price * float64(len(bookingDetails)))

	InvoiceNumber := "INV-" + time.Now().Round(time.Second).Format("20060102150405")
	entityPayment := entity.Payment{
		InvoiceNumber: InvoiceNumber,
		TotalPayment:  totalPayment,
		PaymentStatus: "pending",
		BookingID:     resultBooking.ID,
	}

	_, err = u.repository.UpsertPayment(&entityPayment)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	return nil
}

// RequestPayment implements UseCase.
func (u *useCase) RequestPayment(req *dto_payment.Request) error {
	resultBooking, err := u.repository.FindBookingByID(req.BookingID)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	if resultBooking.ID == 0 {
		return errors.New("booking not found")
	}

	// Validate payment expired

	if resultBooking.PaymentExpiredAt.Before(time.Now()) {
		return errors.New("payment expired")
	}

	paymentMethods, err := u.repository.FindPaymentMethodStatus()
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	isValid := false
	for _, v := range paymentMethods {
		if v.Name == req.PaymentMethod && v.IsActive {
			isValid = true
		}
	}

	if !isValid {
		return errors.New("payment method not found / not active")
	}

	resultPayment, err := u.repository.FindPaymentByBookingID(resultBooking.ID)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	specDoPayment := dto_payment.DoPayment{
		CaseID:         resultBooking.CaseID,
		InvoiceNumber:  resultPayment.InvoiceNumber,
		PaymentMethod:  req.PaymentMethod,
		PaymentAmount:  resultPayment.TotalPayment,
		BookingExpired: resultBooking.BookingExpiredAt,
	}
	err = u.adapter.DoPayment(&specDoPayment)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	return nil
}

// UpdatePayment implements UseCase.
func (u *useCase) UpdatePayment(req *dto_payment.RequestUpdatePayment) error {
	resultPayment, err := u.repository.FindPaymentDetailByInvoice(req.InvoiceNumber)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	resultPayment.PaymentStatus = req.Status
	resultPayment.PaymentDate = time.Now().Round(time.Second)
	resultPayment.PaymentMethod = req.PaymentMethod

	_, err = u.repository.UpsertPayment(&resultPayment)
	if err != nil {
		return tools.ErrorBuilder(err)
	}

	return nil
}

// GetDetailPayment implements UseCase.
func (u *useCase) GetPaymentStatus(codePayment string) (dto_payment.StatusResponse, error) {
	result, err := u.repository.FindPaymentDetailByInvoice(codePayment)
	if err != nil {
		return dto_payment.StatusResponse{}, tools.ErrorBuilder(err)
	}

	if result.ID == 0 {
		return dto_payment.StatusResponse{}, errors.New("payment not found")
	}

	return dto_payment.StatusResponse{
		Status: result.PaymentStatus,
	}, nil

}

// GetPaymentMethod implements UseCase.
func (u *useCase) GetPaymentMethod() ([]dto_payment.MethodResponse, error) {
	result, err := u.repository.FindPaymentMethodStatus()
	if err != nil {
		return []dto_payment.MethodResponse{}, tools.ErrorBuilder(err)
	}

	var response []dto_payment.MethodResponse

	for _, v := range result {
		response = append(response, dto_payment.MethodResponse{
			ID:       v.ID,
			IsActive: v.IsActive,
			Name:     v.Name,
		})
	}

	return response, nil
}

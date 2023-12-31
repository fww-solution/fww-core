package controller

import (
	"fww-core/internal/data/dto"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/tools"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) RequestPayment(msg *message.Message) error {
	var req dto_payment.Request

	if err := json.Unmarshal(msg.Payload, &req); err != nil {
		msg.Ack()
		return err
	}

	if err := c.UseCase.RequestPayment(&req); err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}

	msg.Ack()

	return nil

}

func (c *Controller) GetPaymentStatus(ctx *fiber.Ctx) error {
	codePayment := ctx.Query("payment_code", "")

	result, err := c.UseCase.GetPaymentStatus(codePayment)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "200",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *Controller) GetPaymentMethod(ctx *fiber.Ctx) error {
	result, err := c.UseCase.GetPaymentMethod()
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "200",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Handling Generate Invoice
func (c *Controller) GenerateInvoiceHandler(msg *message.Message) error {
	var req dto_payment.RequestInvoice

	if err := json.Unmarshal(msg.Payload, &req); err != nil {
		msg.Ack()
		return err
	}

	if err := c.UseCase.GenerateInvoice(req.CaseID, req.CodeBooking); err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}

	msg.Ack()

	return nil
}

// Handling Update Payment
func (c *Controller) UpdatePaymentHandler(msg *message.Message) error {
	var req dto_payment.RequestUpdatePayment

	if err := json.Unmarshal(msg.Payload, &req); err != nil {
		msg.Ack()
		return err
	}

	if err := c.UseCase.UpdatePayment(&req); err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}

	msg.Ack()

	return nil
}

package container

import (
	"fww-core/internal/adapter"
	"fww-core/internal/config"
	"fww-core/internal/container/infrastructure/database"
	grpcclient "fww-core/internal/container/infrastructure/grpc/client"
	grpcserver "fww-core/internal/container/infrastructure/grpc/server"
	"fww-core/internal/container/infrastructure/http"
	"fww-core/internal/container/infrastructure/http/router"
	logger "fww-core/internal/container/infrastructure/log"
	messagestream "fww-core/internal/container/infrastructure/message_stream"
	"fww-core/internal/container/infrastructure/redis"
	"fww-core/internal/controller"
	"fww-core/internal/repository"
	"fww-core/internal/usecase"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
)

func InitService(cfg *config.Config) (*fiber.App, []*message.Router) {
	// init database
	db := database.GetConnection(&cfg.Database)
	// init redis
	clientRedis := redis.SetupClient(&cfg.Redis)
	// init redis cache
	redis.InitRedisClient(clientRedis)
	// Init Tracing
	// Init Logger
	log := logger.Initialize(cfg)
	// Init HTTP Server
	server := http.SetupHttpEngine()
	// Init GRPC Server
	grpcserver.Init(&cfg.GrpcServer)
	// Init GRPC Client
	_, err := grpcclient.Init(&cfg.GrpcClient)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	amqpMessageStream := messagestream.NewAmpq(&cfg.MessageStream)

	// set message stream subscriber
	sub, err := amqpMessageStream.NewSubscriber()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// set message stream publisher
	pub, err := amqpMessageStream.NewPublisher()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// Init Publisher

	// Init Adapter
	adapter := adapter.NewBPM(pub, sub)
	// Init Repository
	repo := repository.New(db)
	// Init UseCase
	usecase := usecase.New(repo, adapter, clientRedis)
	// Init Controller
	ctrl := controller.Controller{UseCase: usecase, Log: log}
	var messageRouters []*message.Router
	// Init Router
	requestBookingRouter, err := messagestream.NewRouter(
		pub,
		"request_booking_poisoned",
		"request_booking_handler",
		"request_booking",
		sub,
		ctrl.RequestBooking,
	)
	if err != nil {
		log.Fatal(err)
	}

	requestPaymentRouter, err := messagestream.NewRouter(
		pub,
		"do_payment_poisoned",
		"do_payment_handler",
		"do_payment",
		sub,
		ctrl.RequestPayment,
	)

	if err != nil {
		log.Fatal(err)
	}

	updatePassangerBPM, err := messagestream.NewRouter(
		pub,
		"update_passanger_from_bpm_poisoned",
		"update_passanger_from_bpm_handler",
		"update_passanger_from_bpm",
		sub,
		ctrl.UpdatePassangerByIDNumberHandler,
	)

	if err != nil {
		log.Fatal(err)
	}

	updatePaymentBPM, err := messagestream.NewRouter(
		pub,
		"update_payment_from_bpm_poisoned",
		"update_payment_from_bpm_handler",
		"update_payment_from_bpm",
		sub,
		ctrl.UpdatePaymentHandler,
	)

	if err != nil {
		log.Fatal(err)
	}

	updateBookingBPM, err := messagestream.NewRouter(
		pub,
		"update_booking_from_bpm_poisoned",
		"update_booking_from_bpm_handler",
		"update_booking_from_bpm",
		sub,
		ctrl.UpdateBookingHandler,
	)

	if err != nil {
		log.Fatal(err)
	}

	updateTicketBPM, err := messagestream.NewRouter(
		pub,
		"update_ticket_from_bpm_poisoned",
		"update_ticket_from_bpm_handler",
		"update_ticket_from_bpm",
		sub,
		ctrl.UpdateTicketHandler,
	)
	if err != nil {
		log.Fatal(err)
	}

	requestGenerateInvoiceBPM, err := messagestream.NewRouter(
		pub,
		"generate_invoice_from_bpm_poisoned",
		"generate_invoice_from_bpm_handler",
		"generate_invoice_from_bpm",
		sub,
		ctrl.GenerateInvoiceHandler,
	)

	if err != nil {
		log.Fatal(err)
	}

	sendNotificationBPM, err := messagestream.NewRouter(
		pub,
		"send_notification_from_bpm_poisoned",
		"send_notification_from_bpm_handler",
		"send_notification_from_bpm",
		sub,
		ctrl.SendEmailNotificationHandler,
	)

	if err != nil {
		log.Fatal(err)
	}

	messageRouters = append(messageRouters, requestBookingRouter, requestPaymentRouter, updatePassangerBPM, updateBookingBPM, requestGenerateInvoiceBPM, updatePaymentBPM, updateTicketBPM, sendNotificationBPM)

	// Init Router
	app := router.Initialize(server, &ctrl)

	return app, messageRouters
}

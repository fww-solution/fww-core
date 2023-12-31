package usecase_test

import (
	"database/sql"
	"fww-core/internal/data/dto_flight"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDetailFlightByID(t *testing.T) {
	setup()
	t.Run("Success", func(t *testing.T) {
		id := int64(1)
		expected := dto_flight.ResponseFlightDetail{
			ArrivalAirportName:  "Biak Frans Kaisiepo International Airport",
			ArrivalTime:         timeTimeNow.Format("2006-01-02 15:04:00"),
			CodeFlight:          "GA-002",
			DepartureTime:       timeTimeNow.Format("2006-01-02 15:04:00"),
			DepatureAirportName: "Sorong Domine Eduard Osok Airport",
			FlightPrice:         1000000,
			ReminingSeat:        172,
			Status:              "Delayed",
		}

		entityFlight := entity.Flight{
			ID:                   id,
			CodeFlight:           "GA-002",
			DepartureTime:        timeTimeNow,
			ArrivalTime:          timeTimeNow,
			DepartureAirportName: "Sorong Domine Eduard Osok Airport",
			ArrivalAirportName:   "Biak Frans Kaisiepo International Airport",
			DepartureAirportID:   id,
			ArrivalAirportID:     id,
			Status:               "Delayed",
			CreatedAt:            timeTimeNow,
			UpdatedAt: sql.NullTime{
				Time:  timeTimeNow,
				Valid: true,
			},
			DeletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			PlaneID: id,
		}

		entityFlightPrice := entity.FlightPrice{
			ID:        1,
			Price:     1000000,
			Class:     "Economy",
			CreatedAt: timeTimeNow,
			UpdatedAt: sql.NullTime{
				Time:  timeTimeNow,
				Valid: true,
			},
			DeletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			FlightID: id,
		}

		entityFlightReservation := entity.FlightReservation{
			ID:           1,
			Class:        "Economy",
			ReminingSeat: 172,
			TotalSeat:    172,
			CreatedAt:    timeTimeNow,
			UpdatedAt: sql.NullTime{
				Time:  timeTimeNow,
				Valid: true,
			},
			DeletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			FlightID: id,
		}

		repositoryMock.On("FindFlightByID", id).Return(entityFlight, nil).Once()
		repositoryMock.On("FindFlightPriceByID", id).Return(entityFlightPrice, nil).Once()
		repositoryMock.On("FindFlightReservationByID", id).Return(entityFlightReservation, nil).Once()

		result, err := uc.GetDetailFlightByID(id)
		if err != nil {
			t.Fatalf("unexpectedd error: %v", err)
		}
		if result != expected {
			t.Errorf("expectedd %v, got %v", expected, result)
		}
	})
	t.Run("Error FindFlightByID", func(t *testing.T) {
		id := int64(1)
		expected := dto_flight.ResponseFlightDetail{}

		repositoryMock.On("FindFlightByID", id).Return(entity.Flight{}, sql.ErrNoRows).Once()

		result, err := uc.GetDetailFlightByID(id)
		assert.NotNil(t, err)
		if result != expected {
			t.Errorf("expecteed %v, got %v", expected, result)
		}
	})
	t.Run("Error FindFlightPriceByID", func(t *testing.T) {
		id := int64(1)
		expected := dto_flight.ResponseFlightDetail{}

		entityFlight := entity.Flight{
			ID:                   id,
			CodeFlight:           "GA-003",
			DepartureTime:        time.Now().Round(time.Hour),
			ArrivalTime:          time.Now().Round(time.Hour),
			DepartureAirportName: "Mopah International Airport",
			ArrivalAirportName:   "Sentani International Airpor",
			DepartureAirportID:   id,
			ArrivalAirportID:     id,
			Status:               "Cancelled",
			CreatedAt:            time.Now().Round(time.Hour),
			UpdatedAt: sql.NullTime{
				Time:  time.Now().Round(time.Hour),
				Valid: true,
			},
			DeletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			PlaneID: id,
		}

		repositoryMock.On("FindFlightByID", id).Return(entityFlight, nil).Once()
		repositoryMock.On("FindFlightPriceByID", id).Return(entity.FlightPrice{}, sql.ErrNoRows).Once()

		result, err := uc.GetDetailFlightByID(id)
		assert.NotNil(t, err)
		assert.Equal(t, result, expected)
	})
	t.Run("Error FindFlightReservationByID", func(t *testing.T) {
		id := int64(1)
		expected := dto_flight.ResponseFlightDetail{}

		entityFlight := entity.Flight{
			ID:                   id,
			CodeFlight:           "GA-003",
			DepartureTime:        time.Now().Round(time.Hour),
			ArrivalTime:          time.Now().Round(time.Hour),
			DepartureAirportName: "Patimura International Airport",
			ArrivalAirportName:   "Karel Sadsuitubun International Airport",
			DepartureAirportID:   id,
			ArrivalAirportID:     id,
			Status:               "Cancelled",
			CreatedAt:            time.Now().Round(time.Hour),
			UpdatedAt: sql.NullTime{
				Time:  time.Now().Round(time.Hour),
				Valid: true,
			},
			DeletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			PlaneID: id,
		}

		entityFlightPrice := entity.FlightPrice{
			ID:        1,
			Price:     1000000,
			Class:     "Economy",
			CreatedAt: time.Now().Round(time.Hour),
			UpdatedAt: sql.NullTime{
				Time:  time.Now().Round(time.Hour),
				Valid: true,
			},
			DeletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			FlightID: id,
		}

		repositoryMock.On("FindFlightByID", id).Return(entityFlight, nil).Once()
		repositoryMock.On("FindFlightPriceByID", id).Return(entityFlightPrice, nil).Once()
		repositoryMock.On("FindFlightReservationByID", id).Return(entity.FlightReservation{}, sql.ErrNoRows).Once()

		result, err := uc.GetDetailFlightByID(id)
		assert.NotNil(t, err)
		assert.Equal(t, result, expected)
	})
}

func TestGetFlights(t *testing.T) {
	setup()
	t.Run("Success", func(t *testing.T) {
		departureTime := dateTime
		arrivalTime := dateTime
		limit := 5
		offset := 1
		expected := []dto_flight.ResponseFlight{
			{
				ArrivalAirportName:  "Soekarno-Hatta International Airport",
				ArrivalTime:         time.Now().Round(time.Hour).Format("2006-01-02 15:04:05"),
				CodeFlight:          "GA-004",
				DepartureTime:       time.Now().Round(time.Hour).Format("2006-01-02 15:04:05"),
				DepatureAirportName: "New International Airport - Yogyakarta",
				FlightPrice:         1000000,
				ReminingSeat:        172,
				Status:              "Delayed",
			},
			{
				ArrivalAirportName:  "Adi Sutjipto International Airport",
				ArrivalTime:         timeNow,
				CodeFlight:          "GA-005",
				DepartureTime:       timeNow,
				DepatureAirportName: "Juanda International Airport",
				FlightPrice:         1000000,
				ReminingSeat:        172,
				Status:              "Delayed",
			},
		}

		entityFlights := []entity.Flight{
			{
				ID:                   1,
				CodeFlight:           "GA-004",
				DepartureTime:        time.Now().Round(time.Hour),
				ArrivalTime:          time.Now().Round(time.Hour),
				DepartureAirportName: "New International Airport - Yogyakarta",
				ArrivalAirportName:   "Soekarno-Hatta International Airport",
				DepartureAirportID:   1,
				ArrivalAirportID:     1,
				Status:               "On Time",
				CreatedAt:            time.Now().Round(time.Hour),
				UpdatedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				DeletedAt: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
				PlaneID: 1,
			},
			{
				ID:                   2,
				CodeFlight:           "GA-005",
				DepartureTime:        time.Now().Round(time.Minute),
				ArrivalTime:          time.Now().Round(time.Minute),
				DepartureAirportName: "Adi Sutjipto International Airport",
				ArrivalAirportName:   "Juanada International Airport",
				DepartureAirportID:   1,
				ArrivalAirportID:     1,
				Status:               "On Time",
				CreatedAt:            time.Now().Round(time.Minute),
				UpdatedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				DeletedAt: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
				PlaneID: 1,
			},
		}

		repositoryMock.On("FindFlights", departureTime, arrivalTime, limit, offset).Return(entityFlights, nil).Once()

		result, err := uc.GetFlights(departureTime, arrivalTime, limit, offset)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != len(expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("Error FindFlights", func(t *testing.T) {
		departureTime := dateTime
		arrivalTime := dateTime
		limit := 5
		offset := 1

		repositoryMock.On("FindFlights", departureTime, arrivalTime, limit, offset).Return([]entity.Flight{}, sql.ErrNoRows).Once()

		_, err := uc.GetFlights(departureTime, arrivalTime, limit, offset)
		assert.NotNil(t, err)
	})
}

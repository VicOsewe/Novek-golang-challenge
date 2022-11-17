package main

import (
	"golangchallenge/processors"
	"reflect"
	"sync"
	"testing"

	"github.com/google/uuid"
)

// Implement the following tests:
// Ensure top driver is found the processor
// Ensure top hotel is found the processor
// Ensure memory consumption is less than 128 mb

func getTripDetails() (processors.TripsData, []processors.Trip) {
	testData := processors.TripsData{}

	drivers := []*processors.Driver{
		{
			Id:   "1",
			Name: "Driver1",
		},
		{
			Id:   "2",
			Name: "Driver2",
		},
		{
			Id:   "4",
			Name: "Driver3",
		},
		{
			Id:   "5",
			Name: "Driver4",
		},
	}
	hotels := []*processors.Hotel{
		{
			Id:   "1",
			Name: "Hotel1",
		},
		{
			Id:   "2",
			Name: "Hotel2",
		},
		{
			Id:   "3",
			Name: "Hotel3",
		},
		{
			Id:   "4",
			Name: "Hotel4",
		},
	}

	trips := []processors.Trip{
		{
			Id:           uuid.NewString(),
			DriverId:     drivers[0].Id,
			HotelId:      hotels[0].Id,
			DriverRating: 1,
			HotelRating:  5,
			Status:       "complete",
			Driver:       drivers[0],
			Hotel:        hotels[0],
		},
		{
			Id:           uuid.NewString(),
			DriverId:     drivers[0].Id,
			HotelId:      hotels[0].Id,
			DriverRating: 2,
			HotelRating:  4,
			Status:       "complete",
			Driver:       drivers[0],
			Hotel:        hotels[0],
		},
		{
			Id:           uuid.NewString(),
			DriverId:     drivers[0].Id,
			HotelId:      hotels[0].Id,
			DriverRating: 3,
			HotelRating:  3,
			Status:       "complete",
			Driver:       drivers[0],
			Hotel:        hotels[0],
		},
		{
			Id:           uuid.NewString(),
			DriverId:     drivers[0].Id,
			HotelId:      hotels[0].Id,
			DriverRating: 4,
			HotelRating:  2,
			Status:       "complete",
			Driver:       drivers[0],
			Hotel:        hotels[0],
		},
		{
			Id:           uuid.NewString(),
			DriverId:     drivers[0].Id,
			HotelId:      hotels[0].Id,
			DriverRating: 5,
			HotelRating:  1,
			Status:       "complete",
			Driver:       drivers[0],
			Hotel:        hotels[0],
		},
	}

	testData.Drivers = drivers
	testData.Hotels = hotels
	return testData, trips

}
func TestProcessor_GetTopRankedDriver(t *testing.T) {

	testData, trips := getTripDetails()
	wg := &sync.WaitGroup{}

	processor := processors.CreateProcessorFromData(&testData, wg)
	processor.Trip = append(processor.Trip, trips...)

	driver := processors.Driver{
		Id:   "1",
		Name: "Driver1",
	}
	highestRankedDriverTrip := processors.Trip{
		Id:           uuid.NewString(),
		DriverId:     driver.Id,
		HotelId:      "1",
		DriverRating: 5,
		HotelRating:  1,
		Status:       "complete",
		Driver:       &driver,
		Hotel:        nil,
	}

	tests := []struct {
		name string
		want string
	}{
		{
			name: "get the highest ranking driver",
			want: highestRankedDriverTrip.Driver.Name,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processor.GetTopRankedDriver()
			if !reflect.DeepEqual(got.Name, tt.want) {
				t.Errorf("Processor.GetTopRankedDriver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_GetTopRankedHotel(t *testing.T) {
	testData, trips := getTripDetails()
	wg := &sync.WaitGroup{}

	processor := processors.CreateProcessorFromData(&testData, wg)
	processor.Trip = append(processor.Trip, trips...)

	hotel := processors.Hotel{
		Id:   "1",
		Name: "Hotel1",
	}
	highestRankedHotelTrip := processors.Trip{
		Id:           uuid.NewString(),
		DriverId:     "1",
		HotelId:      hotel.Id,
		DriverRating: 1,
		HotelRating:  5,
		Status:       "complete",
		Driver:       nil,
		Hotel:        &hotel,
	}
	tests := []struct {
		name string
		want string
	}{
		{
			name: "get the highest ranking hotel",
			want: highestRankedHotelTrip.Hotel.Name,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := processor.GetTopRankedHotel(); !reflect.DeepEqual(got.Name, tt.want) {
				t.Errorf("Processor.GetTopRankedHotel() = %v, want %v", got, tt.want)
			}
		})
	}
}

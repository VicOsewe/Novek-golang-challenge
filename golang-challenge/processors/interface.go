package processors

import (
	"fmt"
	"sync"
)

type DriverRanking struct {
	Name   string  `json:"name"`
	ID     string  `json:"id"`
	Rating float64 `json:"rating"`
}

func (d *DriverRanking) String() string {
	return fmt.Sprintf("%s with an id %s and a rating of %v", d.Name, d.ID, d.Rating)
}

type HotelRanking struct {
	Name   string  `json:"name"`
	ID     string  `json:"id"`
	Rating float64 `json:"rating"`
}

func (h *HotelRanking) String() string {
	return fmt.Sprintf("%s with an id %s and a rating of %v", h.Name, h.ID, h.Rating)
}

type ProcessorInterface interface {
	StartProcessing() error
	GetTopRankedDriver() *DriverRanking
	GetTopRankedHotel() *HotelRanking
}

func CreateProcessorFromData(data *TripsData, wg *sync.WaitGroup) *Processor {
	processor := &Processor{
		Trips: data.Trips,
	}

	return processor
}

func (p *Processor) StartProcessing() error {

	for {
		val, ok := <-p.Trips
		if !ok {
			break
		}
		p.Trip = append(p.Trip, *val)
	}

	return nil
}

func (p *Processor) GetTopRankedDriver() *DriverRanking {
	max := p.Trip[0]
	for _, trip := range p.Trip {
		if trip.DriverRating > max.DriverRating {
			max = trip
		}
	}

	return &DriverRanking{
		Name:   max.Driver.Name,
		ID:     max.Driver.Id,
		Rating: max.DriverRating,
	}
}

func (p *Processor) GetTopRankedHotel() *HotelRanking {
	max := p.Trip[0]
	for _, trip := range p.Trip {
		if trip.HotelRating > max.HotelRating {
			max = trip
		}
	}

	return &HotelRanking{
		Name:   max.Hotel.Name,
		ID:     max.Hotel.Id,
		Rating: max.HotelRating,
	}
}

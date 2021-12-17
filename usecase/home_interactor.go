package usecase

import (
	"app/domain"
	"app/usecase/repository"
	"fmt"
	"math"
	"time"
)

type HomeInteractor struct {
	Home   repository.HomeRepository
	Orders repository.OrdersRepository
}

func (interactor *HomeInteractor) TodaySalesResults() (salesData domain.SalesData, err error) {
	today := time.Now().Format("2006-01-02")
	start := today + " 00:00:00"
	end := today + " 23:59:59"
	qeury := fmt.Sprintf("status in (%d,%d)", int(domain.Resarvation), int(domain.Complete))

	salesData, err = interactor.Orders.SalesResultsByDay(start, end, qeury)
	if err != nil {
		return domain.SalesData{}, err
	}

	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	start = yesterday + " 00:00:00"
	end = yesterday + " 23:59:59"
	salesDataYesterday, err := interactor.Orders.SalesResultsByDay(start, end, "")
	if err != nil {
		return domain.SalesData{}, err
	}

	var up bool
	var rate float64
	if salesData.TotalPrice > salesDataYesterday.TotalPrice {
		up = true
		rate = float64(salesData.TotalPrice) / float64(salesDataYesterday.TotalPrice) * 100
	} else {
		up = false
		rate = float64(salesDataYesterday.TotalPrice) / float64(salesData.TotalPrice) * 100
		if salesData.TotalPrice == 0 {
			rate = 0
		}
	}
	salesData.Up = up
	salesData.Rate = math.Round(math.Round(rate*100) / 100)

	return salesData, nil
}

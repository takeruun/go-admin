package usecase

import (
	"app/domain"
	"app/usecase/repository"
	"fmt"
	"math"
	"time"
)

type OrdersInteractor struct {
	Orders repository.OrdersRepository
}

func (interactor *OrdersInteractor) Index(offset int, limit int) (orders []domain.Order, totalCount int64, err error) {
	orders, err = interactor.Orders.FindAll(offset, limit)
	if err != nil {
		return []domain.Order{}, 0, err
	}

	totalCount, err = interactor.Orders.TotalCount()
	if err != nil {
		return []domain.Order{}, 0, err
	}

	return orders, totalCount, nil
}

func (interactor *OrdersInteractor) Show(id int) (order domain.Order, err error) {
	order, err = interactor.Orders.Find(id)

	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (interactor *OrdersInteractor) Create(form domain.Order) (err error) {
	err = interactor.Orders.Create(form)
	if err != nil {
		return err
	}

	return
}

func (interactor *OrdersInteractor) Update(form domain.Order) (order domain.Order, err error) {
	err = interactor.Orders.Update(form)
	if err != nil {
		return domain.Order{}, err
	}

	order, err = interactor.Orders.Find(int(form.ID))

	return
}

func (interactor *OrdersInteractor) OrdersJson(offset int, limit int) (orders []domain.Order, err error) {
	orders, err = interactor.Orders.FindAll(offset, limit)
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil
}

func (interactor *OrdersInteractor) Search(offset int, limit int, searchType string) (orders []domain.Order, err error) {
	if searchType == "todayReservation" {
		today := time.Now().Format("2006-01-02")
		start := today + " 00:00:00"
		end := today + " 23:59:59"

		orders, err = interactor.Orders.TodayReservation(offset, limit, start, end)

		if err != nil {
			return nil, err
		}

	}

	return orders, nil
}

func (interactor *OrdersInteractor) SalesByDay() (salesData []domain.SalesData, salesResult domain.SalesData, err error) {
	salesData, err = interactor.Orders.SalesByDay()
	if err != nil {
		return nil, domain.SalesData{}, err
	}

	today := time.Now().Format("2006-01-02")
	start := today + " 00:00:00"
	end := today + " 23:59:59"
	qeury := fmt.Sprintf("status in (%d,%d)", int(domain.Resarvation), int(domain.Complete))

	salesResult, err = interactor.Orders.SalesResultsByDay(start, end, qeury)
	if err != nil {
		return nil, domain.SalesData{}, err
	}

	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	start = yesterday + " 00:00:00"
	end = yesterday + " 23:59:59"
	salesResultYesterday, err := interactor.Orders.SalesResultsByDay(start, end, "")
	if err != nil {
		return nil, domain.SalesData{}, err
	}

	var up bool
	var rate float64
	if salesResult.TotalPrice > salesResultYesterday.TotalPrice {
		up = true
		rate = float64(salesResult.TotalPrice) / float64(salesResultYesterday.TotalPrice) * 100
	} else {
		up = false
		rate = float64(salesResultYesterday.TotalPrice) / float64(salesResult.TotalPrice) * 100
		if salesResult.TotalPrice == 0 {
			rate = 0
		}
	}
	salesResult.Up = up
	salesResult.Rate = math.Round(math.Round(rate*100) / 100)

	return salesData, salesResult, nil
}

func (interactor *OrdersInteractor) SalesByMonth() (salesData []domain.SalesData, salesResult domain.SalesData, err error) {
	salesData, err = interactor.Orders.SalesByMonth()
	if err != nil {
		return nil, domain.SalesData{}, err
	}

	n := time.Now()
	start := time.Date(n.Year(), n.Month(), 1, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	end := time.Date(n.Year(), n.Month()+1, 1, 23, 59, 59, 0, time.Local).AddDate(0, 0, -1).Format("2006-01-02 15:04:05")

	salesResult, err = interactor.Orders.SalesResultsByDay(start, end, "")
	if err != nil {
		return nil, domain.SalesData{}, err
	}

	start = time.Date(n.Year(), n.Month()-1, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	end = time.Date(n.Year(), n.Month(), 1, 23, 59, 59, 0, time.Local).AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	salesResultAgo, err := interactor.Orders.SalesResultsByDay(start, end, "")
	if err != nil {
		return nil, domain.SalesData{}, err
	}

	var up bool
	var rate float64
	if salesResult.TotalPrice > salesResultAgo.TotalPrice {
		up = true
		rate = float64(salesResult.TotalPrice) / float64(salesResultAgo.TotalPrice) * 100
	} else {
		up = false
		rate = float64(salesResultAgo.TotalPrice) / float64(salesResult.TotalPrice) * 100
		if salesResult.TotalPrice == 0 {
			rate = 0
		}
	}

	salesResult.Up = up
	salesResult.Rate = math.Round(math.Round(rate*100) / 100)

	return salesData, salesResult, nil
}

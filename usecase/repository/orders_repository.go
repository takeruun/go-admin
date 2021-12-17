package repository

import "app/domain"

type OrdersRepository interface {
	FindAll(offset int, limit int) (orders []domain.Order, err error)
	Find(id int) (order domain.Order, err error)
	Create(form domain.Order) (err error)
	Update(form domain.Order) (err error)
	TodayReservation(offset int, limit int, start string, end string) (orders []domain.Order, err error)
	SalesByDay() (salesDate []domain.SalesData, err error)
	SalesByMonth() (salesDate []domain.SalesData, err error)
	SalesResultsByDay(start string, end string, query string) (salesData domain.SalesData, err error)
	TotalCount() (count int64, err error)
}

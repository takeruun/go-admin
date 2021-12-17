package database

import (
	"app/domain"
	"time"

	"gorm.io/gorm"
)

type OrdersRepository struct {
	DB DB
}

func (repo *OrdersRepository) FindAll(offset int, limit int) (orders []domain.Order, err error) {
	db := repo.DB.Connect()

	result := db.Preload("OrderItems.Product").Preload("User").Offset(offset).Limit(limit).Find(&orders)

	if result.Error != nil {
		return []domain.Order{}, result.Error
	}

	return orders, nil
}

func (repo *OrdersRepository) Find(id int) (order domain.Order, err error) {
	db := repo.DB.Connect()

	result := db.Preload("OrderItems.Product.Category").Preload("User").First(&order, id)

	if result.Error != nil {
		return domain.Order{}, result.Error
	}

	return order, nil
}

func (repo *OrdersRepository) Update(form domain.Order) (err error) {
	db := repo.DB.Connect()

	result := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&form)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *OrdersRepository) Create(form domain.Order) (err error) {
	db := repo.DB.Connect()

	result := db.Create(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *OrdersRepository) TodayReservation(offset int, limit int, start string, end string) (orders []domain.Order, err error) {
	db := repo.DB.Connect()

	result := db.Preload("OrderItems.Product").Preload("User").Where("date_of_visit BETWEEN ? AND ?", start, end).Offset(offset).Limit(limit).Find(&orders)

	if result.Error != nil {
		return []domain.Order{}, result.Error
	}

	return orders, nil
}

func (repo *OrdersRepository) SalesByDay() (salseData []domain.SalesData, err error) {
	db := repo.DB.Connect()

	result := db.Raw(
		`
		SELECT sum(total_price) as total_price, date_format(date_of_exit,'%Y-%m-%d') as day
			FROM orders
			WHERE (date_format(date_of_exit,'%Y-%m-%d') > ? AND status = ?)
				OR (date_format(date_of_exit,'%Y-%m-%d') = ? AND status != ?)
			GROUP BY date_format(date_of_exit,'%Y-%m-%d')
			ORDER BY date_format(date_of_exit,'%Y-%m-%d')
		`,
		time.Now().AddDate(0, 0, -14).Format("2006-01-02"), domain.Complete,
		time.Now().Format("2006-01-02"), domain.Cancel).Scan(&salseData)

	if result.Error != nil {
		return nil, result.Error
	}

	return salseData, nil
}

func (repo *OrdersRepository) SalesByMonth() (salseData []domain.SalesData, err error) {
	db := repo.DB.Connect()

	result := db.Raw(
		`
		SELECT sum(total_price) as total_price, date_format(date_of_exit,'%Y-%m') as day 
			FROM orders
			WHERE (date_format(date_of_exit,'%Y-%m') > ? AND status = ?)
				OR (date_format(date_of_exit,'%Y-%m') = ? AND status != ?)
			GROUP BY date_format(date_of_exit,'%Y-%m')
			ORDER BY date_format(date_of_exit,'%Y-%m')
		`,
		time.Now().AddDate(-1, 0, 0).Format("2006-01-02"),
		domain.Complete, time.Now().Format("2006-01-02"), domain.Cancel).Scan(&salseData)

	if result.Error != nil {
		return nil, result.Error
	}

	return salseData, nil
}

func (repo *OrdersRepository) SalesResultsByDay(start string, end string, query string) (salesData domain.SalesData, err error) {
	db := repo.DB.Connect()

	sql := db.Table("orders").Select("sum(total_price) as total_price").Where("date_of_exit BETWEEN ? AND ?", start, end)

	var result *gorm.DB
	if query != "" {
		result = sql.Where(query).Scan(&salesData)
	} else {
		result = sql.Scan(&salesData)
	}

	if result.Error != nil {
		return domain.SalesData{}, result.Error
	}

	return salesData, nil
}

func (repo *OrdersRepository) TotalCount() (count int64, err error) {
	db := repo.DB.Connect()

	result := db.Model(domain.Order{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return
}

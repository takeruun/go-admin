package domain

import (
	"strconv"
)

type Product struct {
	Model
	CategoryID  uint        `json:"categoryId" form:"categoryId"`
	Name        string      `json:"name" form:"name"`
	ProductType ProductType `json:"productType"`
	Price       int         `json:"price" form:"price"`
	Category    Category    `binding:"-" form:"-"`
}

type ProductType int

const (
	Course ProductType = iota + 1
	Goods
	Other
)

func (p ProductType) String() string {
	switch p {
	case Course:
		return "コース"
	case Goods:
		return "商品"
	case Other:
		return "その他"
	default:
		return "Unkwon"
	}
}

func ProductTypes() map[int]map[string]string {
	pts := make(map[int]map[string]string)

	for i, v := range []ProductType{Course, Goods, Other} {
		pts[i] = map[string]string{"ja": v.String(), "value": strconv.Itoa(int(v))}
	}

	return pts
}

package domain

import (
	"strconv"
	"time"
)

type Order struct {
	Model
	UserID         uint          `form:"userId"`
	Status         Status        `form:"status" binding:"required" json:"status"`
	DateOfVisit    time.Time     `form:"dateOfVisit" json:"dateOfVisit" time_format:"2006-01-02 15:04"`
	DateOfExit     time.Time     `form:"dateOfExit" json:"dateOfExit" time_format:"2006-01-02 15:04"`
	SubTotalPrice  int           `form:"subTotalPrice" json:"subTotalPrice"`
	TotalPrice     int           `form:"totalPrice" json:"totalPrice"`
	DiscountType   DiscountType  `form:"discountType" json:"discountType"`
	DiscountMethod int           `form:"discountMethod" json:"discountMethod"`
	DiscountAmount int           `form:"discountAmount" json:"discountAmount"`
	PaymentMethod  PaymentMethod `form:"paymentMethod" json:"paymentMethod"`
	OrderItems     []OrderItem   `form:"orderItem[]"`
	User           User          `binding:"-"`
}

type PaymentMethod int
type DiscountType int
type Status int

const (
	CreditCard PaymentMethod = iota + 1
	Paypay
	Cache
)
const (
	None DiscountType = iota
	Percentage
	PriceReduction
)
const (
	Resarvation Status = iota + 1
	Complete
	Cancel
)

func (p PaymentMethod) String() string {
	switch p {
	case CreditCard:
		return "クレジットカード"
	case Paypay:
		return "PayPay"
	case Cache:
		return "現金"
	default:
		return ""
	}
}

func (d DiscountType) String() string {
	switch d {
	case Percentage:
		return "割合値引"
	case PriceReduction:
		return "金額値引"
	case None:
		return "値引無し"
	default:
		return "Unkwon"
	}
}

func (s Status) String() string {
	switch s {
	case Resarvation:
		return "予約"
	case Complete:
		return "完了"
	case Cancel:
		return "キャンセル"
	default:
		return "Unkwon"
	}
}

func PaymentMethods() map[int]map[string]string {
	occs := make(map[int]map[string]string)

	for i, v := range []PaymentMethod{CreditCard, Paypay, Cache} {
		occs[i] = map[string]string{"ja": v.String(), "value": strconv.Itoa(int(v))}
	}
	return occs
}

func DiscountTypes() map[int]map[string]string {
	dits := make(map[int]map[string]string)

	for i, v := range []DiscountType{Percentage, PriceReduction, None} {
		dits[i] = map[string]string{"ja": v.String(), "value": strconv.Itoa(int(v))}
	}
	return dits
}

func Statuses() map[int]map[string]string {
	stas := make(map[int]map[string]string)

	for i, v := range []Status{Resarvation, Complete, Cancel} {
		stas[i] = map[string]string{"ja": v.String(), "value": strconv.Itoa(int(v))}
	}

	return stas
}

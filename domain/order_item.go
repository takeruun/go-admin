package domain

type OrderItem struct {
	Model
	OrderID     uint
	ProductID   uint `form:"orderItem[productId]"`
	Price       int  `form:"orderItem[price]"`
	Tax         int  `form:"orderItem[tax]"`
	Quantity    int  `form:"orderItem[quantity]"`
	OtherPerson bool `form:"orderItem[otherPerson]"`
	Order       Order
	Product     Product
}

package controllers

import (
	constans "app/constans"
	domain "app/domain"
	database "app/interface/database"
	usecase "app/usecase"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imdario/mergo"
)

type OrdersController struct {
	*BaseContoroller
	Interactor usecase.OrdersInteractor
	Constans   *constans.Constans
}

func NewOrdersController(db database.DB, c *constans.Constans, b *BaseContoroller) *OrdersController {
	return &OrdersController{
		Interactor: usecase.OrdersInteractor{
			Orders: &database.OrdersRepository{DB: db},
		},
		Constans:        c,
		BaseContoroller: b,
	}
}

func (controller *OrdersController) Index(c *gin.Context) {
	session := sessions.Default(c)
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize == 0 {
		pageSize = 20
	}

	orders, totalCount, err := controller.Interactor.Index((pageNumber-1)*pageSize, pageSize)

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/orders")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"orders":      orders,
		"pageNumber":  pageNumber,
		"pageSize":    pageSize,
		"pageNumbers": math.Ceil(float64(totalCount) / float64(pageSize)),
	}
	mergo.Merge(&res, cRes)

	c.HTML(200, "ordersIndex.html", res)
}

func (controller *OrdersController) New(c *gin.Context) {
	session := sessions.Default(c)

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"paymentMethods": domain.PaymentMethods(),
		"discountTypes":  domain.DiscountTypes(),
		"statuses":       domain.Statuses(),
	}
	mergo.Merge(&res, cRes)

	c.HTML(200, "ordersNew.html", res)
}

func (controller *OrdersController) Create(c *gin.Context) {
	session := sessions.Default(c)
	var form domain.Order
	err := c.ShouldBindWith(&form, binding.Form)
	exit := form.DateOfVisit.Format("2006-01-02") + " " + c.PostForm("exit")
	dateOfExit, _ := time.Parse("2006-01-02 15:04", exit)
	form.DateOfExit = dateOfExit

	productIds := c.PostFormArray("orderItem[][productId]")
	prices := c.PostFormArray("orderItem[][price]")
	taxies := c.PostFormArray("orderItem[][tax]")
	otherPersons := c.PostFormArray("orderItem[][otherPerson]")

	var orderItems []domain.OrderItem = make([]domain.OrderItem, len(productIds))
	for i := range orderItems {
		productId, _ := strconv.Atoi(productIds[i])
		price, _ := strconv.Atoi(prices[i])
		tax, _ := strconv.Atoi(taxies[i])
		otherPerson, _ := strconv.ParseBool(otherPersons[i])

		orderItems[i] = domain.OrderItem{ProductID: uint(productId), Price: price, Tax: tax, OtherPerson: otherPerson}
	}
	form.OrderItems = orderItems

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/orders/new")
		return
	}

	err = controller.Interactor.Create(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/orders/new")
		return
	}

	c.Redirect(http.StatusFound, "/orders")
}

func (controller *OrdersController) Edit(c *gin.Context) {
	session := sessions.Default(c)
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := controller.Interactor.Show(id)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/orders")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"paymentMethods": domain.PaymentMethods(),
		"discountTypes":  domain.DiscountTypes(),
		"statuses":       domain.Statuses(),
		"order":          order,
	}
	mergo.Merge(&res, cRes)

	c.HTML(http.StatusOK, "ordersEdit.html", res)
}

func (controller *OrdersController) Update(c *gin.Context) {
	session := sessions.Default(c)
	id, _ := strconv.Atoi(c.Param("id"))
	var form domain.Order
	err := c.ShouldBindWith(&form, binding.Form)
	exit := form.DateOfVisit.Format("2006-01-02") + " " + c.PostForm("exit")
	dateOfExit, _ := time.Parse("2006-01-02 15:04", exit)
	form.DateOfExit = dateOfExit

	orderItemIds := c.PostFormArray("orderItem[][id]")
	productIds := c.PostFormArray("orderItem[][productId]")
	prices := c.PostFormArray("orderItem[][price]")
	taxies := c.PostFormArray("orderItem[][tax]")
	otherPersons := c.PostFormArray("orderItem[][otherPerson]")

	var orderItems []domain.OrderItem = make([]domain.OrderItem, len(productIds))
	for i := range orderItems {
		orderItemId, _ := strconv.Atoi(orderItemIds[i])
		productId, _ := strconv.Atoi(productIds[i])
		price, _ := strconv.Atoi(prices[i])
		tax, _ := strconv.Atoi(taxies[i])
		otherPerson, _ := strconv.ParseBool(otherPersons[i])

		orderItems[i] = domain.OrderItem{ProductID: uint(productId), Price: price, Tax: tax, OtherPerson: otherPerson}
		orderItems[i].ID = uint(orderItemId)
	}
	form.OrderItems = orderItems
	form.ID = uint(id)

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/orders/"+c.Param("id")+"/edit")
		return
	}

	_, err = controller.Interactor.Update(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/orders/"+c.Param("id")+"/edit")
		return
	}

	c.Redirect(http.StatusFound, "/orders")
}

func (controller *OrdersController) OrdersJson(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	orders, err := controller.Interactor.OrdersJson((pageNumber-1)*pageSize, pageSize)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (controller *OrdersController) OrderJson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := controller.Interactor.Show(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (controller *OrdersController) UpdateApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form domain.Order
	form.ID = uint(id)

	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	order, err := controller.Interactor.Update(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (controller *OrdersController) SearchJson(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	searchType := c.Query("searchType")

	orders, err := controller.Interactor.Search(pageNumber-1, pageSize, searchType)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (controller *OrdersController) SalesByDay(c *gin.Context) {
	salesData, salesResult, err := controller.Interactor.SalesByDay()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"salesData":   salesData,
		"salesResult": salesResult,
	})
}

func (controller *OrdersController) SalesByMonth(c *gin.Context) {
	salesData, salesResult, err := controller.Interactor.SalesByMonth()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"salesData":   salesData,
		"salesResult": salesResult,
	})
}

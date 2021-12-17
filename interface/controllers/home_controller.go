package controllers

import (
	constans "app/constans"
	domain "app/domain"
	database "app/interface/database"
	usecase "app/usecase"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
)

type HomeController struct {
	*BaseContoroller
	Interactor usecase.HomeInteractor
	Constans   *constans.Constans
}

func NewHomeController(db database.DB, c *constans.Constans, b *BaseContoroller) *HomeController {
	return &HomeController{
		Interactor: usecase.HomeInteractor{
			Home:   &database.HomeRepository{DB: db},
			Orders: &database.OrdersRepository{DB: db},
		},
		Constans:        c,
		BaseContoroller: b,
	}
}

func (controller *HomeController) Index(c *gin.Context) {
	session := sessions.Default(c)
	var saleDataDay []domain.SalesData

	dayData, err := controller.Interactor.TodaySalesResults()
	saleDataDay = append(saleDataDay, dayData)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"saleDataDay":    saleDataDay,
		"paymentMethods": domain.PaymentMethods(),
		"discountTypes":  domain.DiscountTypes(),
		"statuses":       domain.Statuses(),
	}
	mergo.Merge(&res, cRes)

	c.HTML(http.StatusOK, "homeIndex.html", res)
}

package controllers

import (
	constans "app/constans"
	domain "app/domain"
	database "app/interface/database"
	usecase "app/usecase"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BaseContoroller struct {
	Interactor usecase.BaseIntractor
	Constans   *constans.Constans
}

func NewBaseController(db database.DB, c *constans.Constans) *BaseContoroller {
	return &BaseContoroller{
		Interactor: usecase.BaseIntractor{
			Administrators: &database.AdministratorsRepository{DB: db},
		},
		Constans: c,
	}
}

func SetSessionFlashErr(err error, session sessions.Session) {
	session.AddFlash(err.Error(), "Err")
	session.Save()
}

func GetSessionFlashErr(session sessions.Session) []interface{} {
	err := session.Flashes("Err")
	session.Delete("Err")
	session.Save()

	return err
}

func (controller *BaseContoroller) CommonResponse(session sessions.Session) gin.H {
	var res gin.H
	id := session.Get("admin_id")
	fontSize, err := controller.Interactor.GetMyFontSize(id.(uint))
	errFlash := GetSessionFlashErr(session)

	if err != nil {
		fmt.Print(err)
	}

	res = make(map[string]interface{}, 1)
	res["shopName"] = controller.Constans.SHOP.Name
	res["productTypes"] = domain.ProductTypes()
	res["fontSize"] = fontSize
	res["err"] = errFlash

	return res
}

func (contrller *BaseContoroller) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

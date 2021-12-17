package controllers

import (
	constans "app/constans"
	domain "app/domain"
	database "app/interface/database"
	usecase "app/usecase"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imdario/mergo"
)

type SettingsController struct {
	*BaseContoroller
	Interactor usecase.SettingsInteractor
	Constans   *constans.Constans
}

func NewSettingsController(db database.DB, c *constans.Constans, b *BaseContoroller) *SettingsController {
	return &SettingsController{
		Interactor: usecase.SettingsInteractor{
			Settings:       &database.SettingsRepository{DB: db},
			Administrators: &database.AdministratorsRepository{DB: db},
		},
		Constans:        c,
		BaseContoroller: b,
	}
}

func (controller *SettingsController) Edit(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("admin_id")
	fontSize, err := controller.Interactor.GetMyFontSize(id.(uint))

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"fontSize": fontSize,
	}
	mergo.Merge(&res, cRes)
	c.HTML(http.StatusOK, "settingsEdit.html", res)
}

func (controller *SettingsController) Update(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("admin_id")
	var form domain.Administrator

	form.ID = id.(uint)
	err := c.ShouldBindWith(&form, binding.Form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/settings/edit")
		return
	}

	fontSize, err := controller.Interactor.UpdateFontSize(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/settings/edit")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"fontSize": fontSize,
	}
	mergo.Merge(&res, cRes)
	c.HTML(http.StatusOK, "settingsEdit.html", res)
}

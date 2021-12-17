package controllers

import (
	constans "app/constans"
	domain "app/domain"
	database "app/interface/database"
	usecase "app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imdario/mergo"
)

type AdministratorsController struct {
	*BaseContoroller
	Interactor usecase.AdministratorsInteractor
	Constans   *constans.Constans
}

func NewAdministratorsController(db database.DB, c *constans.Constans, b *BaseContoroller) *AdministratorsController {
	return &AdministratorsController{
		Interactor: usecase.AdministratorsInteractor{
			Administrators: &database.AdministratorsRepository{DB: db},
		},
		Constans:        c,
		BaseContoroller: b,
	}
}

func (controller *AdministratorsController) LoginPage(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Flashes("Err")
	session.Delete("Err")
	session.Save()

	c.HTML(200, "login.html", gin.H{"shopName": controller.Constans.SHOP.Name, "flash": flash})
}

func (controller *AdministratorsController) Login(c *gin.Context) {
	session := sessions.Default(c)
	form := domain.Administrator{Email: c.PostForm("email"), Password: c.PostForm("password")}
	admin, err := controller.Interactor.Login(form)

	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": err.Error()})
		c.Abort()
	}

	session.Set("admin_id", admin.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

func (controller *AdministratorsController) Index(c *gin.Context) {
	session := sessions.Default(c)
	admins, err := controller.Interactor.FindAll()

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"admins": admins,
	}
	mergo.Merge(&res, cRes)

	c.HTML(http.StatusOK, "administratorsIndex.html", res)
}

func (controller *AdministratorsController) New(c *gin.Context) {
	session := sessions.Default(c)

	cRes := controller.CommonResponse(session)

	c.HTML(http.StatusOK, "administratorsNew.html", cRes)
}

func (controller *AdministratorsController) Create(c *gin.Context) {
	session := sessions.Default(c)
	var form domain.Administrator

	err := c.ShouldBindWith(&form, binding.Form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/administrators/new")
		return
	}

	err = controller.Interactor.Create(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/administrators/new")
		return
	}

	c.Redirect(http.StatusFound, "/administrators")
}

func (controller *AdministratorsController) Edit(c *gin.Context) {
	session := sessions.Default(c)
	id, _ := strconv.Atoi(c.Param("id"))

	admin, err := controller.Interactor.Find(uint(id))
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/administrators")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"admin": admin,
	}
	mergo.Merge(&res, cRes)

	c.HTML(http.StatusOK, "administratorsEdit.html", res)
}

func (controller *AdministratorsController) Update(c *gin.Context) {
	session := sessions.Default(c)
	id, _ := strconv.Atoi(c.Param("id"))
	var form domain.Administrator
	form.ID = uint(id)

	err := c.ShouldBindWith(&form, binding.Form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/administrators/"+c.Param("id")+"/edit")
		return
	}

	err = controller.Interactor.Update(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/administrators/"+c.Param("id")+"/edit")
		return
	}

	c.Redirect(http.StatusFound, "/administrators")
}

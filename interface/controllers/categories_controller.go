package controllers

import (
	constans "app/constans"
	domain "app/domain"
	database "app/interface/database"
	usecase "app/usecase"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imdario/mergo"
)

type CategoriesController struct {
	*BaseContoroller
	Interactor usecase.CategoriesInteractor
	Constans   *constans.Constans
}

func NewCategoriesController(db database.DB, c *constans.Constans, b *BaseContoroller) *CategoriesController {
	return &CategoriesController{
		Interactor: usecase.CategoriesInteractor{
			Categories: &database.CategoriesRepository{DB: db},
		},
		Constans:        c,
		BaseContoroller: b,
	}
}

func (controller *CategoriesController) Index(c *gin.Context) {
	session := sessions.Default(c)
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize == 0 {
		pageSize = 20
	}

	categories, totalCount, err := controller.Interactor.Index((pageNumber-1)*pageSize, pageSize)

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"categories":  categories,
		"pageNumber":  pageNumber,
		"pageSize":    pageSize,
		"pageNumbers": math.Ceil(float64(totalCount) / float64(pageSize)),
	}
	mergo.Merge(&res, cRes)

	c.HTML(http.StatusOK, "categoriesIndex.html", res)
}

func (controller *CategoriesController) New(c *gin.Context) {
	session := sessions.Default(c)

	cRes := controller.CommonResponse(session)

	c.HTML(http.StatusOK, "categoriesNew.html", cRes)
}

func (controller *CategoriesController) Create(c *gin.Context) {
	session := sessions.Default(c)
	var form domain.Category

	err := c.ShouldBindWith(&form, binding.Form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/categories/new")
		return
	}

	err = controller.Interactor.Create(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/categories/new")
		return
	}

	c.Redirect(http.StatusFound, "/categories")
}

func (controller *CategoriesController) Edit(c *gin.Context) {
	session := sessions.Default(c)
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := controller.Interactor.Show(id)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/categories")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"category": category,
	}
	mergo.Merge(&res, cRes)

	c.HTML(http.StatusOK, "categoriesEdit.html", res)
}

func (controller *CategoriesController) Update(c *gin.Context) {
	session := sessions.Default(c)
	id, _ := strconv.Atoi(c.Param("id"))
	var form domain.Category
	form.ID = uint(id)

	err := c.ShouldBindWith(&form, binding.Form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/categories/"+c.Param("id")+"/edit")
		return
	}

	err = controller.Interactor.Update(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/categories/"+c.Param("id")+"/edit")
		return
	}

	c.Redirect(http.StatusFound, "/categories")
}

func (controller *CategoriesController) CategoriesJson(c *gin.Context) {
	categories, err := controller.Interactor.CategoriesJson()

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

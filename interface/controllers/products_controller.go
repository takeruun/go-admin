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

type ProductsController struct {
	*BaseContoroller
	Interactor usecase.ProductsInteractor
	Constans   *constans.Constans
}

func NewProductsController(db database.DB, c *constans.Constans, b *BaseContoroller) *ProductsController {
	return &ProductsController{
		Interactor: usecase.ProductsInteractor{
			Products: &database.ProductsRepository{DB: db},
		},
		Constans:        c,
		BaseContoroller: b,
	}
}

func (controller *ProductsController) Index(c *gin.Context) {
	session := sessions.Default(c)
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize == 0 {
		pageSize = 20
	}
	productType, _ := strconv.Atoi(c.Param("productType"))
	form := domain.Product{ProductType: domain.ProductType(productType)}

	products, totalCount, err := controller.Interactor.Index((pageNumber-1)*pageSize, pageSize, form)

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/products/")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"products":        products,
		"productType":     productType,
		"productTypeName": domain.ProductType(productType),
		"pageNumber":      pageNumber,
		"pageSize":        pageSize,
		"pageNumbers":     math.Ceil(float64(totalCount) / float64(pageSize)),
	}
	mergo.Merge(&res, cRes)

	c.HTML(
		http.StatusOK,
		"productsIndex.html",
		res,
	)
}

func (controller *ProductsController) New(c *gin.Context) {
	session := sessions.Default(c)
	cRes := controller.CommonResponse(session)
	productType, _ := strconv.Atoi(c.Param("productType"))
	res := gin.H{
		"productType":     productType,
		"productTypeName": domain.ProductType(productType),
	}
	mergo.Merge(&res, cRes)

	c.HTML(http.StatusOK, "productsNew.html", res)
}

func (controller *ProductsController) Create(c *gin.Context) {
	session := sessions.Default(c)
	productType, _ := strconv.Atoi(c.Param("productType"))
	var form domain.Product

	form.ProductType = domain.ProductType(productType)
	err := c.ShouldBindWith(&form, binding.Form)

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/products/"+c.Param("productType")+"new")
		return
	}

	err = controller.Interactor.Create(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/products/"+c.Param("productType")+"new")
		return
	}

	c.Redirect(http.StatusFound, "/products/"+c.Param("productType"))
}

func (controller *ProductsController) Edit(c *gin.Context) {
	session := sessions.Default(c)
	productType, _ := strconv.Atoi(c.Param("productType"))
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := controller.Interactor.Show(id)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/orders")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"product":         product,
		"productType":     productType,
		"productTypeName": domain.ProductType(productType),
	}
	mergo.Merge(&res, cRes)

	c.HTML(http.StatusOK, "productsEdit.html", res)
}

func (controller *ProductsController) Update(c *gin.Context) {
	session := sessions.Default(c)
	productType, _ := strconv.Atoi(c.Param("productType"))
	id, _ := strconv.Atoi(c.Param("id"))

	var form domain.Product
	form.ProductType = domain.ProductType(productType)
	form.ID = uint(id)
	err := c.ShouldBindWith(&form, binding.Form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/products/"+c.Param("productType")+"new")
		return
	}

	err = controller.Interactor.Update(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/products/"+c.Param("productType")+"new")
		return
	}

	c.Redirect(http.StatusFound, "/products/"+c.Param("productType"))
}

func (controller *ProductsController) ProductsJson(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Query("categoryId"))
	name := c.Query("name")
	productType, _ := strconv.Atoi(c.Query("productType"))

	form := domain.Product{CategoryID: uint(categoryId), Name: name, ProductType: domain.ProductType(productType)}

	products, err := controller.Interactor.Search(form)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func (controller *ProductsController) ProductJson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := controller.Interactor.Show(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

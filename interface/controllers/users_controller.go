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

type UsersController struct {
	*BaseContoroller
	Interactor usecase.UsersInteractor
	Constans   *constans.Constans
}

func NewUsersController(db database.DB, c *constans.Constans, b *BaseContoroller) *UsersController {
	return &UsersController{
		Interactor: usecase.UsersInteractor{
			Users: &database.UsersRepository{DB: db},
		},
		Constans:        c,
		BaseContoroller: b,
	}
}

func (controller *UsersController) Index(c *gin.Context) {
	session := sessions.Default(c)
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize == 0 {
		pageSize = 20
	}

	users, totalCount, err := controller.Interactor.Index((pageNumber-1)*pageSize, pageSize)

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/users")
		return
	}
	cRes := controller.CommonResponse(session)
	res := gin.H{
		"users":       users,
		"totalCount":  totalCount,
		"pageNumber":  pageNumber,
		"pageSize":    pageSize,
		"pageNumbers": math.Ceil(float64(totalCount) / float64(pageSize)),
	}
	mergo.Merge(&res, cRes)

	c.HTML(200, "usersIndex.html", res)
}

func (controller *UsersController) New(c *gin.Context) {
	session := sessions.Default(c)

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"prefectures":   domain.Prefectures(),
		"occupations":   domain.Occupations(),
		"genders":       domain.Genders(),
		"relationships": domain.Relationships(),
	}
	mergo.Merge(&res, cRes)

	c.HTML(200, "usersNew.html", res)
}

func (controller *UsersController) Create(c *gin.Context) {
	session := sessions.Default(c)
	var form domain.User
	err := c.ShouldBindWith(&form, binding.Form)

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/users/new")
		return
	}

	_, err = controller.Interactor.Add(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/users/new")
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (controller *UsersController) Edit(c *gin.Context) {
	session := sessions.Default(c)
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.Interactor.Show(id)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/users")
		return
	}

	cRes := controller.CommonResponse(session)
	res := gin.H{
		"prefectures":   domain.Prefectures(),
		"genders":       domain.Genders(),
		"occupations":   domain.Occupations(),
		"relationships": domain.Relationships(),
		"user":          user,
	}
	mergo.Merge(&res, cRes)

	c.HTML(200, "usersEdit.html", res)
}

func (controller *UsersController) Update(c *gin.Context) {
	session := sessions.Default(c)
	id, _ := strconv.Atoi(c.Param("id"))
	var form domain.User
	err := c.ShouldBindWith(&form, binding.Form)
	form.ID = uint(id)

	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/users/"+c.Param("id")+"/edit")
		return
	}

	err = controller.Interactor.Update(form)
	if err != nil {
		SetSessionFlashErr(err, session)
		c.Redirect(http.StatusFound, "/users/"+c.Param("id")+"/edit")
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (controller *UsersController) UsersJson(c *gin.Context) {
	var queries []map[string]string

	if c.Query("familyName") != "" {
		queries = append(queries, map[string]string{"family_name": c.Query("familyName")})
	}
	if c.Query("id") != "" {
		queries = append(queries, map[string]string{"id": c.Query("id")})
	}

	users, err := controller.Interactor.Search(queries)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

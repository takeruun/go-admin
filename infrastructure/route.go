package infrastructure

import (
	constans "app/constans"
	controllers "app/interface/controllers"
	middleware "app/middleware"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	DB    *DB
	Gin   *gin.Engine
	AwsS3 *AwsS3
	Port  string
}

func NewRouting(db *DB, awsS3 *AwsS3) *Routing {
	c := NewConfig()

	r := &Routing{
		DB:    db,
		Gin:   gin.Default(),
		AwsS3: awsS3,
		Port:  c.Routing.Port,
	}
	myfile, _ := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(myfile)
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	constans := constans.NewConstans()
	r.Gin.HTMLRender = loadTemplates("./templates")

	r.Gin.Static("/plugins", "templates/plugins")
	r.Gin.Static("/dist", "templates/dist")
	r.Gin.Static("/custom", "templates/custom")

	store := cookie.NewStore([]byte("secret"))
	r.Gin.Use(sessions.Sessions("art", store))
	r.Gin.Use(middleware.LoginCheckMiddleware())

	baseController := controllers.NewBaseController(r.DB, constans)
	homeContoller := controllers.NewHomeController(r.DB, constans, baseController)
	administratorsController := controllers.NewAdministratorsController(r.DB, constans, baseController)
	usersController := controllers.NewUsersController(r.DB, constans, baseController)
	ordersController := controllers.NewOrdersController(r.DB, constans, baseController)
	productsController := controllers.NewProductsController(r.DB, constans, baseController)
	categoriesController := controllers.NewCategoriesController(r.DB, constans, baseController)
	settingsController := controllers.NewSettingsController(r.DB, constans, baseController)

	r.Gin.GET("/", func(c *gin.Context) { homeContoller.Index(c) })
	r.Gin.GET("/login", func(c *gin.Context) { administratorsController.LoginPage(c) })
	r.Gin.POST("/login", func(c *gin.Context) { administratorsController.Login(c) })
	r.Gin.GET("/health", func(c *gin.Context) { baseController.Health(c) })
	admins := r.Gin.Group("/administrators")
	{
		admins.GET("", func(c *gin.Context) { administratorsController.Index(c) })
		admins.GET("/new", func(c *gin.Context) { administratorsController.New(c) })
		admins.POST("/create", func(c *gin.Context) { administratorsController.Create(c) })
		admins.GET("/:id/edit", func(c *gin.Context) { administratorsController.Edit(c) })
		admins.POST("/:id/update", func(c *gin.Context) { administratorsController.Update(c) })
	}

	users := r.Gin.Group("users")
	{
		users.GET("", func(c *gin.Context) { usersController.Index(c) })
		users.GET("/new", func(c *gin.Context) { usersController.New(c) })
		users.POST("/create", func(c *gin.Context) { usersController.Create(c) })
		users.GET("/:id/edit", func(c *gin.Context) { usersController.Edit(c) })
		users.POST("/:id/update", func(c *gin.Context) { usersController.Update(c) })
	}

	orders := r.Gin.Group("orders")
	{
		orders.GET("", func(c *gin.Context) { ordersController.Index(c) })
		orders.GET("/new", func(c *gin.Context) { ordersController.New(c) })
		orders.POST("/create", func(c *gin.Context) { ordersController.Create(c) })
		orders.GET("/:id/edit", func(c *gin.Context) { ordersController.Edit(c) })
		orders.POST("/:id/update", func(c *gin.Context) { ordersController.Update(c) })
	}

	products := r.Gin.Group("products")
	{
		products.GET("/:productType", func(c *gin.Context) { productsController.Index(c) })
		products.GET("/:productType/new", func(c *gin.Context) { productsController.New(c) })
		products.POST("/:productType/create", func(c *gin.Context) { productsController.Create(c) })
		products.GET("/:productType/:id/edit", func(c *gin.Context) { productsController.Edit(c) })
		products.POST("/:productType/:id/update", func(c *gin.Context) { productsController.Update(c) })
	}

	categories := r.Gin.Group("categories")
	{
		categories.GET("", func(c *gin.Context) { categoriesController.Index(c) })
		categories.GET("/new", func(c *gin.Context) { categoriesController.New(c) })
		categories.POST("/create", func(c *gin.Context) { categoriesController.Create(c) })
		categories.GET("/:id/edit", func(c *gin.Context) { categoriesController.Edit(c) })
		categories.POST("/:id/update", func(c *gin.Context) { categoriesController.Update(c) })
	}

	settings := r.Gin.Group("settings")
	{
		settings.GET("/edit", func(c *gin.Context) { settingsController.Edit(c) })
		settings.POST("/update", func(c *gin.Context) { settingsController.Update(c) })
	}

	api := r.Gin.Group("api")
	{
		api.GET("/users", func(c *gin.Context) { usersController.UsersJson(c) })
		api.GET("/products", func(c *gin.Context) { productsController.ProductsJson(c) })
		api.GET("/products/:id", func(c *gin.Context) { productsController.ProductJson(c) })
		api.GET("/categories", func(c *gin.Context) { categoriesController.CategoriesJson(c) })

		orders := api.Group("orders")
		orders.GET("/:id", func(c *gin.Context) { ordersController.OrderJson(c) })
		orders.PUT("/:id", func(c *gin.Context) { ordersController.UpdateApi(c) })
		orders.GET("/search", func(c *gin.Context) { ordersController.SearchJson(c) })
		orders.GET("/sales_by_day", func(c *gin.Context) { ordersController.SalesByDay(c) })
		orders.GET("/sales_by_month", func(c *gin.Context) { ordersController.SalesByMonth(c) })
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layout, err := filepath.Glob(templatesDir + "/views/layout.html")
	if err != nil {
		panic(err.Error())
	}

	views, err := filepath.Glob(templatesDir + "/views/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	shareds, err := filepath.Glob(templatesDir + "/views/shared/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, view := range views {
		layoutCopy := make([]string, len(layout))
		copy(layoutCopy, layout)

		files := append(layoutCopy, view)
		for _, shared := range shareds {
			files = append(files, shared)
		}
		r.AddFromFiles(filepath.Base(view), files...)
	}

	login, err := filepath.Glob(templatesDir + "/views/login.html")
	if err != nil {
		panic(err.Error())
	}
	r.AddFromFiles(filepath.Base(login[0]), login...)

	return r
}

func (r *Routing) SetMiddleware() {
	r.Gin.Use(gin.Recovery(), gin.Logger(), middleware.RecordLogAndTime)
}

func (r *Routing) SetReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}

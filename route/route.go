package route

import (
	"absen/controllers"
	adm "absen/controllers/admin"
	m "absen/middleware"
	"absen/utils"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

var (
	JWT_SECRET_KEY = utils.GetConfig("JWT_SECRET_KEY")
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	loggerConfig := m.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}
	loggerMiddleware := loggerConfig.Init()
	e.Use(loggerMiddleware)

	v1 := e.Group("/api/v1")
	eJwt := v1.Group("")
	eJwt.Use(mid.JWT([]byte(JWT_SECRET_KEY)))

	// Route / to handler function
	user := controllers.InitUserController()
	v1.POST("/users/login", user.Login)
	v1.POST("/users/register", user.Register)
	eJwt.GET("/users/:username", user.GetByUsername)
	eJwt.PUT("/users", user.Update)
	eJwt.PUT("/users/password", user.ChangePassword)

	present := controllers.InitPresentController()
	eJwt.GET("/presents/widget", present.GetHomeWidget)
	eJwt.GET("/presents/user", present.GetAll)
	eJwt.GET("/presents/user/:id", present.GetByID)
	eJwt.GET("/presents", present.Search)
	eJwt.POST("/presents", present.Create)

	admin := v1.Group("/admin")
	admin.Use(mid.JWT([]byte(JWT_SECRET_KEY)))

	adminUser := adm.InitAdminUserController()
	admin.GET("/users", adminUser.GetAll)
	admin.GET("/users/:id", adminUser.GetByID)
	admin.POST("/users", adminUser.Create)
	admin.PUT("/users/:id", adminUser.Update)
	admin.DELETE("/users/:id", adminUser.Delete)

	adminPresent := adm.InitAdminPresentController()
	admin.GET("/presents", adminPresent.GetAll)
	admin.GET("/presents/:id", adminPresent.GetByID)
	admin.GET("/presents/search", adminPresent.Search)

	return e
}

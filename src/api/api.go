package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/itera-io/openstack-web-client/api/middlewares"
	"github.com/itera-io/openstack-web-client/api/routers"
	"github.com/itera-io/openstack-web-client/config"
	"github.com/itera-io/openstack-web-client/docs"
	"github.com/itera-io/openstack-web-client/pkg/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler) /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// Test
		health := v1.Group("/health")
		test_router := v1.Group("/test" /*middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"})*/)

		// User
		users := v1.Group("/users")

		//countries := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		// Test
		routers.Health(health)
		routers.TestRouter(test_router)

		// User
		routers.User(users, cfg)

		r.Static("/static", "./uploads")
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
	v3 := api.Group("/v3")
	{
		region := v3.Group("/regions", middlewares.Authentication(cfg))
		user := v3.Group("/users", middlewares.Authentication(cfg))
		project := v3.Group("/projects", middlewares.Authentication(cfg))
		routers.Region(region, cfg)
		routers.Project(project, cfg)
		routers.AuthenticatedUser(user, cfg)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Title = "Openstack Web Client"
	docs.SwaggerInfo.Description = "This client provides interacting by openstack through `gophercloud`"
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

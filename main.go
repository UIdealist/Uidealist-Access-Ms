package main

import (
	"context"
	"os"

	"github.com/casbin/casbin/v2"

	psqlwatcher "github.com/IguteChung/casbin-psql-watcher"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/UIdealist/Uidealist-Access-Ms/pkg/configs"
	"github.com/UIdealist/Uidealist-Access-Ms/pkg/middleware"
	"github.com/UIdealist/Uidealist-Access-Ms/pkg/routes"
	"github.com/UIdealist/Uidealist-Access-Ms/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/UIdealist/Uidealist-Access-Ms/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @UIdealist API
// @version 1.0
// @description UIdealist Authorization project API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email edgardanielgd123@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	conn, dbType, err := utils.ConnectionURLBuilder("")

	if err != nil {
		panic(err)
	}

	// Create adapter that reads from database at startup.
	a, _ := gormadapter.NewAdapter(
		dbType,
		conn,
	)

	// Create watcher that receives update events from database (the same one).
	w, _ := psqlwatcher.NewWatcherWithConnString(
		context.Background(), conn,
		psqlwatcher.Option{NotifySelf: false, Verbose: true},
	)

	// Create an enforcer with the adapter.
	e, _ := casbin.NewEnforcer("model.conf", a)

	// Set the watcher for the enforcer.
	e.SetWatcher(w)

	// Set callback for watcher.
	w.SetUpdateCallback(psqlwatcher.DefaultCallback(e))

	// Load the policy from DB.
	e.LoadPolicy()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)    // Register a route for API Docs (Swagger).
	routes.PublicRoutes(e, app) // Register a public routes for app.
	routes.PrivateRoutes(app)   // Register a private routes for app.
	routes.NotFoundRoute(app)   // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}

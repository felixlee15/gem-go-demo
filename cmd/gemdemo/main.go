package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gpgin "github.com/gempages/go-helper/gin"
	"github.com/gempages/go-helper/process"
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"gemdemo/config"
	"gemdemo/db"
	"gemdemo/internal/app/delivery/graph"
	"gemdemo/internal/pkg/factory"
	"gemdemo/internal/pkg/graph/generated"
)

var (
	appConfig = config.AppConfigs
)

func main() {
	logrus.Info("Starting the application...")

	executeCommand()
}

func executeCommand() {
	run()
}

func closeServices() {
	db.Close()
}

func run() {
	server := &http.Server{}
	shutdownCompleted := process.ShutdownCallback(func() {
		defer sentry.Flush(2 * time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = server.Shutdown(ctx)
		closeServices()
	})

	db.Init(context.Background())

	r := gpgin.New(&gpgin.Config{
		CORS: &gpgin.CORSConfig{
			AllowHeaders: []string{
				"*",
			},
			AllowOriginRegex: appConfig.App.TrustedOriginRegex,
		},
		Logger: &gpgin.LoggerConfig{
			SkipPaths: []string{"/graphql/query"},
		},
	})

	h, err := graphqlHandler()
	if err != nil {
		logrus.Fatalf("init graphql handler failed: %s", err)
	}

	g := r.Group("/graphql")
	{
		g.GET("", playgroundHandler())
		gq := g.Group("/query")
		{
			gq.POST("", timeout.New(
				timeout.WithTimeout(30*time.Second), timeout.WithHandler(h),
			))
			gq.GET("", timeout.New(
				timeout.WithTimeout(30*time.Second), timeout.WithHandler(h),
			))
		}
	}

	logrus.Info("Server is running on port: " + appConfig.App.Port)
	server.Addr = ":" + appConfig.App.Port
	server.Handler = r.Handler()
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("Starting server: %s", err)
	}
	<-shutdownCompleted
}

func graphqlHandler() (gin.HandlerFunc, error) {
	// Graphql server
	useCaseFactory := factory.NewUseCaseFactory()
	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers:  graph.NewResolverRoot(useCaseFactory),
		Complexity: generated.ComplexityRoot{},
	})

	server := handler.NewDefaultServer(schema)

	return gin.WrapH(server), nil
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.ApolloSandboxHandler("GraphQL Query", "/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"go-demo/config"
	gpgin "go-demo/helper/gin"
	"go-demo/helper/process"
	"go-demo/internal/app/delivery/graph"
	factory "go-demo/internal/pkg/factory/opt"
	"go-demo/internal/pkg/graph/generated"
)

func main() {
	logrus.Info("Starting the application...")

	if err := run(); err != nil {
		logrus.Errorf("application stopped with error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	server := &http.Server{}
	shutdownCompleted := process.ShutdownCallback(func() {
		defer sentry.Flush(2 * time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = server.Shutdown(ctx)
	})

	r := gpgin.New(&gpgin.Config{
		CORS: &gpgin.CORSConfig{
			AllowHeaders: []string{
				// "X-GemX-Shop-ID",
				// "X-Shopify-Domain",
				// "X-Gem-Session",
				"*",
			},
			AllowOriginRegex: config.GetEnv("TRUSTED_ORIGIN_REGEX", ""),
		},
		Logger: &gpgin.LoggerConfig{
			SkipPaths: []string{"/graphql/query"},
		},
	})

	h, err := graphqlHandler()
	if err != nil {
		return fmt.Errorf("init graphql handler failed: %w", err)
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

	logrus.Info("Server is running on port: " + config.GetEnv("PORT", "8080"))
	server.Addr = fmt.Sprintf("%s:%s", config.GetEnv("HOST", "localhost"), config.GetEnv("PORT", "8080"))
	server.Handler = r.Handler()
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("Starting server: %s", err)
	}
	<-shutdownCompleted

	return nil
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

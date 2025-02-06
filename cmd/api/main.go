package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mryeibis/indexer/internal/features/emails"
	"github.com/mryeibis/indexer/internal/models"
	"github.com/mryeibis/indexer/internal/settings"
	"github.com/mryeibis/indexer/internal/zincsearch"
	"github.com/mryeibis/indexer/pkg/logger"
)

const SERVER_PORT = 8080

func getAllowedOrigins() []string {
	origins, ok := os.LookupEnv("ALLOWED_ORIGINS")
	if !ok {
		origins = "http://localhost:5173"
	}

	return strings.Split(origins, ",")
}

func main() {
	logs := logger.New()

	env, err := settings.GetEnvVariables()
	if err != nil {
		logs.Error(fmt.Sprintf("Error getting environment variables: %s", err.Error()))
		return
	}

	zincSearch := zincsearch.New[models.Email](
		env.ZincSearchURL,
		env.ZincSearchUser,
		env.ZincSearchPassword,
		env.ZincSearchIndexName,
	)

	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		cors.Handler(cors.Options{
			AllowedOrigins:   getAllowedOrigins(),
			AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
			AllowedHeaders:   []string{"Content-Type", "Accept"},
			AllowCredentials: true,
		}),
	)
	r.Mount("/api/emails", emails.NewRouter(logs, zincSearch))

	port := fmt.Sprintf(":%d", SERVER_PORT)

	logs.Info(fmt.Sprintf("Starting server on %s", port))
	http.ListenAndServe(port, r)
}

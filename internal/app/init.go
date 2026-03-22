package app

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ThroughTheThornsToTheStarss/todo/internal/api"
	"github.com/ThroughTheThornsToTheStarss/todo/internal/pkg/postgress"
	postgressrepo "github.com/ThroughTheThornsToTheStarss/todo/internal/repo/postgress"
	"github.com/ThroughTheThornsToTheStarss/todo/internal/usecase"
)

type App struct {
	HTTPServer *http.Server
	HTTPPort   string
}

func NewFromEnv() (*App, error) {
	httpPort := getenv("HTTP_PORT", "8080")
	cfg := postgress.LoadConfigFromEnv()
	db, err := cfg.ConnectFromEnv()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	err = postgress.AutoMigrate(db)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	
	repo := postgressrepo.NewPostgresRepository(db)
	todoUC := usecase.NewTodoUsecase(repo)
	apiHandler := api.New(todoUC)
	
	// Создаем главный роутер
	mux := http.NewServeMux()
	
	// API routes (должны быть первыми, более специфичные)
	mux.Handle("/api/", apiHandler)
	
	// Статические файлы из папки client
	staticDir := filepath.Join(".", "client", "static")
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Главная страница из папки client (для всех остальных путей)
	indexPath := filepath.Join(".", "client", "index.html")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Раздаем index.html только если это не API и не статический файл
		if r.URL.Path != "/" && r.URL.Path != "" {
			// Для SPA - все пути ведут на index.html
			http.ServeFile(w, r, indexPath)
			return
		}
		http.ServeFile(w, r, indexPath)
	})
	
	httpSrv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}
	return &App{
		HTTPServer: httpSrv,
		HTTPPort:   httpPort,
	}, nil
}

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

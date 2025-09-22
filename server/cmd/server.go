package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Habeebamoo/Clivo/server/internal/config"
	"github.com/Habeebamoo/Clivo/server/internal/database"
	"github.com/Habeebamoo/Clivo/server/internal/handlers"
	"github.com/Habeebamoo/Clivo/server/internal/repositories"
	"github.com/Habeebamoo/Clivo/server/internal/routes"
	"github.com/Habeebamoo/Clivo/server/internal/services"
)

func main() {
	//load config files
	config.Initialize()

	//init database
	db, err := database.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	//initialized repositories
	userRepo := repositories.NewUserRepository(db)

	//initialized services
	userService := services.NewUserService(userRepo)

	//initialized handlers
	userHandler := handlers.NewUserHandler(userService)

	//initialized routes
	router := routes.ConfigureRoutes(userHandler)

	PORT, _ := config.Get("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//router
	srv := &http.Server{
		Addr: ":"+PORT,
		Handler: router,
	}

	//start server
	log.Println("Server running on PORT ", PORT)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err == http.ErrServerClosed {
			log.Fatal("Server error ", err)
		}
	}()

	//Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shutdown server")
	}

	log.Println("Server exiting neatly...")
}
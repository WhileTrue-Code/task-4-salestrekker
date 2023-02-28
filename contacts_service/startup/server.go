package startup

import (
	"contacts_service/controller"
	"contacts_service/domain"
	"contacts_service/repository"
	"contacts_service/service"
	"contacts_service/startup/config"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	config *config.Config
}

func NewServer() *Server {
	return &Server{
		config: config.NewConfig(),
	}
}

func (server *Server) Start() {
	client := server.initDatabaseClient()
	repo := server.initContactsRepository(client)
	serv := server.initContactsService(repo)
	contactsController := server.initContactsController(serv)
	server.start(contactsController)
}

func (server *Server) initContactsRepository(client *mongo.Client) domain.ContactRepository {
	return repository.NewContactMongoRepo(client, server.config.DBName)
}

func (server *Server) initContactsService(repository domain.ContactRepository) domain.ContactService {
	return service.NewContactService(repository)
}

func (server *Server) initContactsController(service domain.ContactService) *controller.ContactsController {
	return controller.NewContactsController(service)
}

func (server *Server) initDatabaseClient() *mongo.Client {
	client := repository.GetDatabaseClient(server.config.DBPort, server.config.DBUsername, server.config.DBPassword)
	return client
}

func (server *Server) start(controller *controller.ContactsController) {
	router := mux.NewRouter()
	controller.Init(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.config.Port),
		Handler: router,
	}

	timeout := time.Second * 15
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Server error: %s", err)
		}
	}()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, syscall.SIGTERM)
	<-channel

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error on shutting down server: %s", err)
	}
	log.Println("Server stopped gracefully")
}

package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/anishchenkoivan/hotel-app/hotel-service/api/apiv1pb"
	"github.com/anishchenkoivan/hotel-app/hotel-service/config"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/repository"
	hotelservice "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotel"
	hotelierservice "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotelier"
	roomservice "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/room"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
)

type HotelServiceApp struct {
	server     *http.Server
	grpcServer *grpc.Server
	config     config.Config
}

func NewHotelServiceApp(config config.Config) *HotelServiceApp {
	router := mux.NewRouter().PathPrefix("/hotel-service/api").Subrouter()

	hotelApp := HotelServiceApp{
		&http.Server{
			Addr:    config.ServerHost + ":" + config.ServerPort,
			Handler: router,
		},
		grpc.NewServer(),
		config,
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&model.Hotel{}, &model.Hotelier{}, &model.Room{}); err != nil {
		log.Fatal(err)
	}

	hotelRepository := repository.NewPostgresHotelRepository(db)
	hotelierRepository := repository.NewPostgresHotelierRepository(db)
	roomRepository := repository.NewPostgresRoomRepository(db)

	hotelService := hotelservice.NewHotelService(hotelRepository)
	hotelierService := hotelierservice.NewHotelierService(hotelierRepository)
	roomService := roomservice.NewRoomService(roomRepository)

	hotelHandler := handlers.NewHotelHandler(hotelService)
	hotelierHandler := handlers.NewHotelierHandler(hotelierService)
	roomHandler := handlers.NewRoomHandler(roomService)

	router.HandleFunc("/hotel", hotelHandler.CreateHotel).Methods("POST")
	router.HandleFunc("/hotel", hotelHandler.FindAllHotels).Methods("GET")
	router.HandleFunc("/hotel/{id}", hotelHandler.UpdateHotel).Methods("PUT")
	router.HandleFunc("/hotel/{id}", hotelHandler.FindHotelById).Methods("GET")
	router.HandleFunc("/hotel/{id}", hotelHandler.DeleteHotel).Methods("DELETE")

	router.HandleFunc("/hotelier", hotelierHandler.CreateHotelier).Methods("POST")
	router.HandleFunc("/hotelier/{id}", hotelierHandler.FindHotelierById).Methods("GET")
	router.HandleFunc("/hotelier/{id}", hotelierHandler.UpdateHotelier).Methods("PUT")
	router.HandleFunc("/hotelier/{id}", hotelierHandler.DeleteHotelier).Methods("DELETE")

	router.HandleFunc("/room", roomHandler.FindAllRooms).Methods("POST")
	router.HandleFunc("/room", roomHandler.FindAllRooms).Methods("GET")
	router.HandleFunc("/room/{id}", roomHandler.UpdateRoom).Methods("PUT")
	router.HandleFunc("/room/{id}", roomHandler.FindRoomById).Methods("GET")
	router.HandleFunc("/room/{id}", roomHandler.DeleteRoom).Methods("DELETE")

	roomGrpcHandler := handlers.NewRoomGrpcHandler(roomService)
	apiv1pb.RegisterHotelServiceServer(hotelApp.grpcServer, roomGrpcHandler)
	return &hotelApp
}

func (app *HotelServiceApp) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := app.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	group.Go(func() error {
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", app.config.ServerGrpcHost, app.config.ServerGrpcPort))
		if err != nil {
			return err
		}
		if err = app.grpcServer.Serve(listener); err != nil {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		return app.Stop()
	})

	log.Println("Hotel Service started on port 8080")

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}

func (app *HotelServiceApp) Stop() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), app.config.AppShutdownTimeout)
	defer cancel()

	if err := app.server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Println("Hotel Service stopped")
	return nil
}

package app

import (
	"context"
	"github.com/anishchenkoivan/hotel-app/notification-service/config"
	"github.com/anishchenkoivan/hotel-app/notification-service/internal/app/consumers"
	"github.com/anishchenkoivan/hotel-app/notification-service/internal/service"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/errgroup"
	"log"
	"os/signal"
	"syscall"
)

type NotificationApp struct {
	config   config.Config
	consumer consumers.Consumer
}

func NewNotificationApp(config config.Config) *NotificationApp {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.BrokerHost + ":" + config.BrokerPort},
		Topic:   config.Topic,
		GroupID: config.GroupId,
	})

	notificationService := service.NewService()

	consumer := consumers.NewConsumer(reader, notificationService)

	return &NotificationApp{
		config:   config,
		consumer: *consumer,
	}
}

func (app *NotificationApp) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		app.consumer.Consume(ctx)
		return nil
	})

	log.Println("App has started")

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}

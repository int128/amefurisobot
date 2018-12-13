package usecases

import (
	"context"

	"github.com/int128/amefurisobot/domain"
)

type SubscriptionRepository interface {
	FindAll(ctx context.Context) ([]domain.Subscription, error)
}

type PNGRepository interface {
	GetById(ctx context.Context, id string) ([]byte, error)
	Save(ctx context.Context, id string, b []byte) error
}

type WeatherService interface {
	Get(locations []domain.Location) ([]domain.Weather, error)
}

type NotificationService interface {
	Send(destination domain.Notification, message domain.Message) error
}

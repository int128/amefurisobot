package usecases

import (
	"context"
	"strconv"
	"time"

	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/presenters/chart"
	"github.com/pkg/errors"
)

type PollWeathers struct {
	SubscriptionRepository SubscriptionRepository
	WeatherService         WeatherService
	PNGRepository          PNGRepository
	PNGURL                 func(id string) string
	NotificationService    NotificationService
}

func (u *PollWeathers) Do(ctx context.Context) error {
	subscriptions, err := u.SubscriptionRepository.FindAll(ctx)
	if err != nil {
		return errors.Wrapf(err, "error while getting subscriptions")
	}
	if len(subscriptions) == 0 {
		return errors.New("no subscription found")
	}
	var locations []domain.Location
	for _, subscription := range subscriptions {
		locations = append(locations, subscription.Location)
	}
	weathers, err := u.WeatherService.Get(locations)
	if err != nil {
		return errors.Wrapf(err, "error while getting weather")
	}
	for i, subscription := range subscriptions {
		if subscription.Notification.IsZero() {
			continue
		}
		weather := weathers[i]
		if !weather.IsRainingNow() && !weather.WillRainLater() {
			continue
		}

		b, err := chart.DrawPNG(weather)
		if err != nil {
			return errors.Wrapf(err, "error while drawing rainfall chart")
		}
		id := strconv.FormatInt(time.Now().UnixNano(), 36)
		if err := u.PNGRepository.Save(ctx, id, b); err != nil {
			return errors.Wrapf(err, "error while saving the image")
		}
		message := domain.Message{
			Text:     weather.Location.Name,
			ImageURL: u.PNGURL(id),
		}
		if err := u.NotificationService.Send(subscription.Notification, message); err != nil {
			return errors.Wrapf(err, "error while sending the message")
		}
	}
	return nil
}

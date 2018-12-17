package usecases

import (
	"context"
	"github.com/int128/amefurisobot/domain"
	"github.com/pkg/errors"
)

type Setup struct {
	UserRepository         UserRepository
	SubscriptionRepository SubscriptionRepository
}

func (u *Setup) Do(ctx context.Context) (domain.User, error) {
	user := domain.NewUser()
	if err := u.UserRepository.Save(ctx, user); err != nil {
		return domain.User{}, errors.Wrapf(err, "error while saving user %s", user)
	}

	subscription := domain.NewSubscription(domain.Location{
		Name:        "Roppongi",
		Coordinates: domain.Coordinates{Latitude: 35.663613, Longitude: 139.732293},
	})
	subscriptions := []domain.Subscription{subscription}

	if err := u.SubscriptionRepository.Save(ctx, user.ID, subscriptions); err != nil {
		return domain.User{}, errors.Wrapf(err, "error while saving subscriptions for user %s", user)
	}
	return user, nil
}

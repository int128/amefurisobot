package main

import (
	"context"
	"net/http"

	"github.com/int128/amefurisobot/gateways"
	"github.com/int128/amefurisobot/handlers"
	"github.com/int128/amefurisobot/infrastructure"
	"github.com/int128/amefurisobot/usecases"
	"google.golang.org/appengine"
)

func contextProvider(req *http.Request) context.Context {
	return appengine.NewContext(req)
}

func main() {
	h := handlers.Handlers{
		GetWeather: handlers.GetWeather{
			ContextProvider: contextProvider,
			Usecase: &usecases.GetWeather{
				UserRepository:         &gateways.UserRepository{},
				SubscriptionRepository: &gateways.SubscriptionRepository{},
				WeatherService: &gateways.WeatherService{
					Client: &infrastructure.WeatherClient{},
				},
			},
		},
		GetImage: handlers.GetImage{
			ContextProvider: contextProvider,
			Usecase: &usecases.GetImage{
				PNGRepository: &gateways.PNGRepository{},
			},
		},
		PollWeathers: handlers.PollWeathers{
			ContextProvider: contextProvider,
			Usecase: &usecases.PollWeathers{
				UserRepository:         &gateways.UserRepository{},
				SubscriptionRepository: &gateways.SubscriptionRepository{},
				PNGRepository:          &gateways.PNGRepository{},
				WeatherService: &gateways.WeatherService{
					Client: &infrastructure.WeatherClient{},
				},
				NotificationService: &gateways.NotificationService{
					Client: &infrastructure.SlackClient{},
				},
			},
		},
		Setup: handlers.Setup{
			ContextProvider: contextProvider,
			Usecase: &usecases.Setup{
				SubscriptionRepository: &gateways.SubscriptionRepository{},
				UserRepository:         &gateways.UserRepository{},
			},
		},
	}
	http.Handle("/", h.NewRouter())
	appengine.Main()
}

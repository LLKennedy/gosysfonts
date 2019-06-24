// Package fontapi talks to the google font API
package fontapi

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/webfonts/v1"
)

// Connect connects to the API and returns an active connection for further queries
func Connect() error {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	ctx := context.Background()
	service, err := webfonts.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return err
	}
	list, err := service.Webfonts.List().Do()
	if err != nil {
		return err
	}
	fonts := list.Items
	for _, font := range fonts {
		log.Printf("name: %s category: %s\n", font.Family, font.Category)
	}
	return err
}

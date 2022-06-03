package lib

import (
	"log"

	"github.com/google/uuid"
)

type Place struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func AddPlace(place Place) error {
	place.ID = uuid.NewString()
	log.Println("adding place, ", place)
	return nil
}

func GetPlace(description string) (Place, error) {
	log.Println("getting place...")
	var place Place
	return place, nil
}

package groceries

import (
	"context"

	"github.com/Avinashanakal/grocery-delivery-service/internal/pkg/models"
)

func (s *service) Search(ctx context.Context,
	searchQuery string,
) ([]GrocerySearchResult, error) {

	/*
		- Search for groceries in the Grocery table
		- As of now added search based on the name we can add search filers more
	*/
	var groceries []models.Grocery
	if err := s.db.
		Where("name LIKE ?", "%"+searchQuery+"%").
		Find(&groceries).
		Error; err != nil {
		s.log.WithError(err).Error("Search for groceries in the Grocery table")
		return []GrocerySearchResult{}, err
	}

	// Convert the grocery results to a slice of grocery search results
	var searchResults []GrocerySearchResult
	for _, grocery := range groceries {
		searchResult := GrocerySearchResult{
			ID:          grocery.ID,
			Name:        grocery.Name,
			Description: grocery.Description,
			Price:       grocery.Price,
		}
		searchResults = append(searchResults, searchResult)
	}

	s.log.Info("groceries found for the search request")

	return searchResults, nil
}

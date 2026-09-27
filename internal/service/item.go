package service

import (
	"fmt"
	"strconv"

	"github.com/itemstore/tvrockets/internal/model"
	"github.com/itemstore/tvrockets/internal/repository"
)

type ItemService struct {
	repo   repository.ItemRepository
	nextID int
}

func NewItemService(repo repository.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (i *ItemService) GetAll() ([]model.Item, error) {
	items, err := i.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("service get all items: %w", err)
	}
	return items, nil
}

func (i *ItemService) GetByID(id string) (model.Item, error) {
	item, err := i.repo.GetByID(id)
	if err != nil {
		return model.Item{}, fmt.Errorf("service get item by id: %w", err)
	}
	return item, nil
}

func (i *ItemService) Create(item model.Item) (model.Item, error) {
	if item.Title == "" {
		return model.Item{}, &model.ValidationError{Message: "Title required", Field: "title"}
	}
	if item.Price <= 0 {
		return model.Item{}, &model.ValidationError{Message: "Bad price", Field: "price"}
	}
	i.nextID++
	item.ID = strconv.Itoa(i.nextID)
	if err := i.repo.Create(item); err != nil {
		return model.Item{}, fmt.Errorf("service create item: %w", err)
	}
	return item, nil
}

func (i *ItemService) Update(id string, item model.Item) (model.Item, error) {
	if item.Title == "" {
		return model.Item{}, &model.ValidationError{Message: "Title required", Field: "title"}
	}
	if item.Price <= 0 {
		return model.Item{}, &model.ValidationError{Message: "Bad price", Field: "price"}
	}
	item.ID = id
	if err := i.repo.Update(item); err != nil {
		return model.Item{}, fmt.Errorf("service update item: %w", err)
	}
	return item, nil
}

func (i *ItemService) Delete(id string) error {
	if err := i.repo.Delete(id); err != nil {
		return fmt.Errorf("service delete item: %w", err)
	}
	return nil
}

package repository

import (
	"fmt"

	"github.com/itemstore/tvrockets/internal/model"
)

type ItemRepository interface {
	GetAll() ([]model.Item, error)
	GetByID(id string) (model.Item, error)
	Create(item model.Item) error
	Update(item model.Item) error
	Delete(id string) error
}

type InMemoryItemRepository struct {
	items map[string]model.Item
}

func NewInMemoryItemRepository() *InMemoryItemRepository {
	return &InMemoryItemRepository{
		items: make(map[string]model.Item),
	}
}

func (i *InMemoryItemRepository) GetAll() ([]model.Item, error) {
	values := make([]model.Item, 0, len(i.items))
	for _, v := range i.items {
		values = append(values, v)
	}
	return values, nil
}

func (i *InMemoryItemRepository) GetByID(id string) (model.Item, error) {
	if item, ok := i.items[id]; ok {
		return item, nil
	}
	return model.Item{}, fmt.Errorf("get item %q: %w", id, model.ErrNotFound)
}

func (i *InMemoryItemRepository) Create(item model.Item) error {
	if _, ok := i.items[item.ID]; ok {
		return fmt.Errorf("create item %q: %w", item.ID, model.ErrAlreadyExists)
	}
	i.items[item.ID] = item
	return nil
}
func (i *InMemoryItemRepository) Update(item model.Item) error {
	if _, ok := i.items[item.ID]; !ok {
		return fmt.Errorf("update item %q: %w", item.ID, model.ErrNotFound)
	}
	i.items[item.ID] = item
	return nil
}

func (i *InMemoryItemRepository) Delete(id string) error {
	if _, ok := i.items[id]; !ok {
		return fmt.Errorf("delete item %q: %w", id, model.ErrNotFound)
	}
	delete(i.items, id)
	return nil
}

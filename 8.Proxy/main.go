package main

import (
	"fmt"
	"sync"
)

// The proxy design pattern provides a surrogate or placeholder for another object to control access to it.
// It involves a class, called the proxy, which represents the functionality of another class.
// The proxy pattern is useful for tasks like lazy instantiation (creating an object only when it's actually needed),
// permission control, and more sophisticated memory management.

type Product struct {
	ID   int
	Name string
}

type ProductService interface {
	GetProduct(ID int) (*Product, error)
}

type ProductServiceAPI struct{}

func (s *ProductServiceAPI) GetProduct(ID int) (*Product, error) {
	// Simulate a real API request
	fmt.Println("Fetching user from API...")
	return &Product{ID: ID, Name: "Best product"}, nil
}

type ProductServiceClient struct {
	service ProductService
	cache   map[int]*Product
	mux     sync.Mutex
}

func (s *ProductServiceClient) GetProduct(ID int) (*Product, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if user, ok := s.cache[ID]; ok {
		fmt.Println("Returning from cache...")
		return user, nil
	}
	user, err := s.service.GetProduct(ID)
	if err != nil {
		return nil, err
	}
	s.cache[ID] = user
	return user, nil
}

func NewProductServiceClient(service ProductService) ProductService {
	return &ProductServiceClient{
		service: service,
		cache:   make(map[int]*Product),
	}
}

func main() {
	service := ProductServiceAPI{}
	client := NewProductServiceClient(&service)

	client.GetProduct(1)

	fmt.Println("Retry")

	client.GetProduct(1)
}

package main

import (
	"fmt"
	"github.com/itemstore/tvrockets/internal/config"
	"github.com/itemstore/tvrockets/internal/handler"
	"github.com/itemstore/tvrockets/internal/repository"
	"github.com/itemstore/tvrockets/internal/service"
	"net/http"
	"os"
)

func main() {
	cfg := config.Load()
	repo := repository.NewInMemoryItemRepository()
	svc := service.NewItemService(repo)
	h := handler.NewItemHandler(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	if err := http.ListenAndServe(cfg.Port, mux); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

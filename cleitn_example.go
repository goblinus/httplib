package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/goblinus/oapi/src/adapters/rest"
)

// Response structures for JSON unmarshalling
type ListPmapsResponse struct {
	Data []struct {
		Id          int    `json:"id"`
		Uuid        string `json:"uuid"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"data"`
}

type PmapDetailResponse struct {
	PmapInfo struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"pmap_info"`
	Yaml string `json:"yaml"`
}

func main() {
	// Инициализация клиента с поддержкой ответов
	client, err := rest.NewClientWithResponses("http://localhost:8080")
	if err != nil {
		log.Fatalf("Ошибка создания клиента: %v", err)
	}

	// Пример 1: Получение списка pmaps
	fmt.Println("Получение списка pmaps...")
	listParams := &rest.GetPmapsParams{
		Limit:  intPtr(10),
		Offset: intPtr(0),
	}

	listResp, err := client.GetPmapsWithResponse(context.Background(), listParams)
	if err != nil {
		log.Fatalf("Ошибка запроса списка: %v", err)
	}

	if listResp.StatusCode() != 200 {
		log.Printf("Получен ответ: %s", string(listResp.Body))
		log.Fatalf("Неверный статус код: %d", listResp.StatusCode())
	}

	var listResult ListPmapsResponse
	if err := json.Unmarshal(listResp.Body, &listResult); err != nil {
		log.Printf("Raw response: %s", string(listResp.Body))
		log.Fatalf("Ошибка парсинга ответа: %v", err)
	}

	fmt.Printf("Получено %d pmaps\n", len(listResult.Data))
	for _, pmap := range listResult.Data {
		fmt.Printf("- %s (ID: %d, UUID: %s)\n", pmap.Title, pmap.Id, pmap.Uuid)
	}

	// Пример 2: Получение конкретного pmap
	if len(listResult.Data) > 0 {
		uuid := listResult.Data[0].Uuid
		fmt.Printf("\nПолучение деталей pmap с UUID %s...\n", uuid)

		detailResp, err := client.GetPmapsPmapUuidWithResponse(context.Background(), uuid)
		if err != nil {
			log.Fatalf("Ошибка запроса деталей: %v", err)
		}

		if detailResp.StatusCode() != 200 {
			log.Fatalf("Неверный статус код: %d", detailResp.StatusCode())
		}

		var detailResult PmapDetailResponse
		if err := json.Unmarshal(detailResp.Body, &detailResult); err != nil {
			log.Fatalf("Ошибка парсинга ответа: %v", err)
		}

		fmt.Printf("Детали pmap:\n")
		fmt.Printf("Название: %s\n", detailResult.PmapInfo.Title)
		fmt.Printf("Описание: %s\n", detailResult.PmapInfo.Description)
		fmt.Printf("YAML контент: %s\n", detailResult.Yaml)
	}
}

// Вспомогательная функция для создания указателей на int
func intPtr(i int) *int {
	return &i
}

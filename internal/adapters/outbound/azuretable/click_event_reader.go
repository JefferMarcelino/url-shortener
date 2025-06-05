package azuretable

import (
	"context"
	"encoding/json"
	"fmt"
	"urlshortener/internal/domain"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type ClickEventReader struct {
	client    *aztables.Client
	tableName string
}

func NewClickEventReader(client *aztables.ServiceClient, tableName string) domain.AnalyticsReader {
	tableClient := client.NewClient(tableName)
	return &ClickEventReader{client: tableClient, tableName: tableName}
}

func (r *ClickEventReader) GetClickEventsByCode(code string) ([]domain.ClickEvent, error) {
	raw, err := r.fetchRawEntities(code)
	if err != nil {
		return nil, err
	}
	return parseClickEventsEntities(raw), nil
}

func (r *ClickEventReader) fetchRawEntities(code string) ([][]byte, error) {
	filter := fmt.Sprintf("PartitionKey eq '%s'", code)
	pager := r.client.NewListEntitiesPager(&aztables.ListEntitiesOptions{Filter: &filter})

	var entities [][]byte
	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		entities = append(entities, page.Entities...)
	}

	return entities, nil
}

func parseClickEventsEntities(entities [][]byte) []domain.ClickEvent {
	var clickEvents []domain.ClickEvent

	for _, entity := range entities {
		var props map[string]any

		if err := json.Unmarshal(entity, &props); err != nil {
			continue
		}

		clickEvents = append(clickEvents, domain.ClickEvent{
			Code:      fmt.Sprintf("%v", props["PartitionKey"]),
			IP:        fmt.Sprintf("%v", props["IP"]),
			UserAgent: fmt.Sprintf("%v", props["UserAgent"]),
			Timestamp: fmt.Sprintf("%v", props["Timestamp"]),
		})
	}

	return clickEvents
}

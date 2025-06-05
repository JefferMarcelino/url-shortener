package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"urlshortener/internal/domain/model"
	domainRepo "urlshortener/internal/domain/repository"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type AzureAnalyticsRepo struct {
	client *aztables.Client
}

func NewAzureAnalyticsRepository(client *aztables.ServiceClient, tableName string) domainRepo.ClickAnalyticsRepository {
	tableClient := client.NewClient(tableName)

	return &AzureAnalyticsRepo{client: tableClient}
}

func (r *AzureAnalyticsRepo) Save(a *model.ClickAnalytics) error {
	entity := map[string]any{
		"PartitionKey": a.Code,
		"RowKey":       fmt.Sprintf("%d", a.Timestamp),
		"IP":           a.IP,
		"UserAgent":    a.UserAgent,
		"Timestamp":    a.Timestamp,
	}

	marshalled, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	_, err = r.client.AddEntity(context.TODO(), marshalled, nil)

	return err
}

func (r *AzureAnalyticsRepo) GetAnalyticsByCode(code string) ([]model.Analytics, error) {
	raw, err := r.fetchRawEntities(code)
	if err != nil {
		return nil, err
	}
	return parseAnalyticsEntities(raw), nil
}

func (r *AzureAnalyticsRepo) fetchRawEntities(code string) ([][]byte, error) {
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

func parseAnalyticsEntities(entities [][]byte) []model.Analytics {
	var analytics []model.Analytics

	for _, entity := range entities {
		var props map[string]any

		if err := json.Unmarshal(entity, &props); err != nil {
			continue
		}

		analytics = append(analytics, model.Analytics{
			IP:        fmt.Sprintf("%v", props["IP"]),
			UserAgent: fmt.Sprintf("%v", props["UserAgent"]),
			Timestamp: fmt.Sprintf("%v", props["Timestamp"]),
		})
	}

	return analytics
}

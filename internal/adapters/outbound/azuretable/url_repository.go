package azuretable

import (
	"context"
	"encoding/json"
	"fmt"
	"urlshortener/internal/domain"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type AzureURLRepo struct {
	client    *aztables.Client
	tableName string
}

func NewAzureURLRepository(client *aztables.ServiceClient, tableName string) domain.URLRepository {
	tableClient := client.NewClient(tableName)

	return &AzureURLRepo{
		client:    tableClient,
		tableName: tableName,
	}
}

func (r *AzureURLRepo) Save(url domain.ShortURL) error {
	entity := map[string]any{
		"PartitionKey": r.tableName,
		"RowKey":       url.Code,
		"LongURL":      url.LongURL,
	}

	marshalled, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	_, err = r.client.AddEntity(context.TODO(), marshalled, nil)

	return err
}

func (r *AzureURLRepo) GetByCode(code string) (*domain.ShortURL, error) {
	resp, err := r.client.GetEntity(context.TODO(), r.tableName, code, nil)
	if err != nil {
		return nil, err
	}

	var props map[string]any
	_ = json.Unmarshal(resp.Value, &props)

	return &domain.ShortURL{
		Code:    code,
		LongURL: fmt.Sprintf("%v", props["LongURL"]),
	}, nil
}

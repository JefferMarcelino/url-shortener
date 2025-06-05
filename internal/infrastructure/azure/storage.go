package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"urlshortener/internal/domain"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type tableRepo struct {
	client *aztables.Client
}

func NewAzureTableRepository() domain.URLRepository {
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME")
	accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT_KEY")
	tableName := os.Getenv("AZURE_TABLE_NAME")

	cred, err := aztables.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatalf("failed to create credential: %v", err)
	}

	serviceUrl := fmt.Sprintf("https://%s.table.core.windows.net", accountName)
	client, err := aztables.NewServiceClientWithSharedKey(serviceUrl, cred, nil)

	if err != nil {
		log.Fatalf("failed to create service client: %v", err)
	}

	tableClient := client.NewClient(tableName)

	return &tableRepo{client: tableClient}
}

func (r *tableRepo) Save(url domain.URL) error {
	entity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: "urls",
			RowKey:       url.Code,
		},
		Properties: map[string]any{
			"LongURL": url.LongURL,
		},
	}

	marshalled, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	_, err = r.client.AddEntity(context.TODO(), marshalled, nil)

	return err
}

func (r *tableRepo) GetByCode(code string) (*domain.URL, error) {
	resp, err := r.client.GetEntity(context.TODO(), "urls", code, nil)
	if err != nil {
		return nil, err
	}

	var props map[string]any
	_ = json.Unmarshal(resp.Value, &props)

	return &domain.URL{
		Code:    code,
		LongURL: fmt.Sprintf("%v", props["LongURL"]),
	}, nil
}

package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"urlshortener/internal/domain/model"
	domainRepo "urlshortener/internal/domain/repository"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type tableRepo struct {
	client    *aztables.Client
	tableName string
}

func NewURLRepository(accountName string, accountKey string, tableName string) domainRepo.URLRepository {
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

	return &tableRepo{
		client:    tableClient,
		tableName: tableName,
	}
}

func (r *tableRepo) Save(url model.ShortURL) error {
	entity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: r.tableName,
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

func (r *tableRepo) GetByCode(code string) (*model.ShortURL, error) {
	resp, err := r.client.GetEntity(context.TODO(), r.tableName, code, nil)
	if err != nil {
		return nil, err
	}

	var props map[string]any
	_ = json.Unmarshal(resp.Value, &props)

	return &model.ShortURL{
		Code:    code,
		LongURL: fmt.Sprintf("%v", props["LongURL"]),
	}, nil
}

package azuretable

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"urlshortener/internal/domain"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type ClickEventWriter struct {
	client    *aztables.Client
	tableName string
}

func NewClickEventWriter(client *aztables.ServiceClient, tableName string) domain.AnalyticsReporter {
	tableClient := client.NewClient(tableName)
	return &ClickEventWriter{client: tableClient, tableName: tableName}
}

func (r *ClickEventWriter) Save(a *domain.ClickEvent) error {
	entity := map[string]any{
		"PartitionKey": a.Code,
		"RowKey":       fmt.Sprintf("%v", time.Now().Unix()),
		"IP":           a.IP,
		"UserAgent":    a.UserAgent,
	}

	marshalled, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	_, err = r.client.AddEntity(context.TODO(), marshalled, nil)

	return err
}

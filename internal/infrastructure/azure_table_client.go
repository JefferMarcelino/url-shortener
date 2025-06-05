package infrastructure

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func NewAzureTablesServiceClient(accountName, accountKey string) *aztables.ServiceClient {
	cred, err := aztables.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatalf("failed to create credential: %v", err)
	}

	serviceUrl := fmt.Sprintf("https://%s.table.core.windows.net", accountName)
	client, err := aztables.NewServiceClientWithSharedKey(serviceUrl, cred, nil)

	if err != nil {
		log.Fatalf("failed to create service client: %v", err)
	}

	return client
}

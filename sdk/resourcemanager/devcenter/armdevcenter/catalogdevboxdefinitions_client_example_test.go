//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CatalogDevBoxDefinitions_ListByCatalog.json
func ExampleCatalogDevBoxDefinitionsClient_NewListByCatalogPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogDevBoxDefinitionsClient().NewListByCatalogPager("rg1", "Contoso", "CentralCatalog", &armdevcenter.CatalogDevBoxDefinitionsClientListByCatalogOptions{Top: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DevBoxDefinitionListResult = armdevcenter.DevBoxDefinitionListResult{
		// 	Value: []*armdevcenter.DevBoxDefinition{
		// 		{
		// 			Name: to.Ptr("WebDevBox"),
		// 			Type: to.Ptr("Microsoft.DevCenter/devcenters/catalogs/devboxdefinitions"),
		// 			ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/catalogs/CentralCatalog/devboxdefinitions/WebDevBox"),
		// 			SystemData: &armdevcenter.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user1"),
		// 				LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("centralus"),
		// 			Properties: &armdevcenter.DevBoxDefinitionProperties{
		// 				HibernateSupport: to.Ptr(armdevcenter.HibernateSupportEnabled),
		// 				ImageReference: &armdevcenter.ImageReference{
		// 					ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/Example/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/contosogallery/images/exampleImage/version/1.0.0"),
		// 				},
		// 				SKU: &armdevcenter.SKU{
		// 					Name: to.Ptr("Preview"),
		// 				},
		// 				ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CatalogDevBoxDefinitions_Get.json
func ExampleCatalogDevBoxDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCatalogDevBoxDefinitionsClient().Get(ctx, "rg1", "Contoso", "CentralCatalog", "WebDevBox", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DevBoxDefinition = armdevcenter.DevBoxDefinition{
	// 	Name: to.Ptr("WebDevBox"),
	// 	Type: to.Ptr("Microsoft.DevCenter/devcenters/catalogs/devboxdefinitions"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/catalogs/CentralCatalog/devboxdefinitions/WebDevBox"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("centralus"),
	// 	Properties: &armdevcenter.DevBoxDefinitionProperties{
	// 		HibernateSupport: to.Ptr(armdevcenter.HibernateSupportEnabled),
	// 		ImageReference: &armdevcenter.ImageReference{
	// 			ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/Example/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/contosogallery/images/exampleImage/version/1.0.0"),
	// 		},
	// 		SKU: &armdevcenter.SKU{
	// 			Name: to.Ptr("Preview"),
	// 		},
	// 		ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		ValidationStatus: to.Ptr(armdevcenter.CatalogResourceValidationStatusSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CatalogDevBoxDefinitions_GetErrorDetails.json
func ExampleCatalogDevBoxDefinitionsClient_GetErrorDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCatalogDevBoxDefinitionsClient().GetErrorDetails(ctx, "rg1", "Contoso", "CentralCatalog", "WebDevBox", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CatalogResourceValidationErrorDetails = armdevcenter.CatalogResourceValidationErrorDetails{
	// 	Errors: []*armdevcenter.CatalogErrorDetails{
	// 		{
	// 			Code: to.Ptr("TaskMissing"),
	// 			Message: to.Ptr("The referenced Task 'RunPowershellCommand' could not be found."),
	// 	}},
	// }
}

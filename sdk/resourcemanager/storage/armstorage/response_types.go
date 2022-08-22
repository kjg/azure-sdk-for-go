//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorage

// AccountsClientAbortHierarchicalNamespaceMigrationResponse contains the response from method AccountsClient.AbortHierarchicalNamespaceMigration.
type AccountsClientAbortHierarchicalNamespaceMigrationResponse struct {
	// placeholder for future response values
}

// AccountsClientCheckNameAvailabilityResponse contains the response from method AccountsClient.CheckNameAvailability.
type AccountsClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResult
}

// AccountsClientCreateResponse contains the response from method AccountsClient.Create.
type AccountsClientCreateResponse struct {
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientFailoverResponse contains the response from method AccountsClient.Failover.
type AccountsClientFailoverResponse struct {
	// placeholder for future response values
}

// AccountsClientGetPropertiesResponse contains the response from method AccountsClient.GetProperties.
type AccountsClientGetPropertiesResponse struct {
	Account
}

// AccountsClientHierarchicalNamespaceMigrationResponse contains the response from method AccountsClient.HierarchicalNamespaceMigration.
type AccountsClientHierarchicalNamespaceMigrationResponse struct {
	// placeholder for future response values
}

// AccountsClientListAccountSASResponse contains the response from method AccountsClient.ListAccountSAS.
type AccountsClientListAccountSASResponse struct {
	ListAccountSasResponse
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.ListByResourceGroup.
type AccountsClientListByResourceGroupResponse struct {
	AccountListResult
}

// AccountsClientListKeysResponse contains the response from method AccountsClient.ListKeys.
type AccountsClientListKeysResponse struct {
	AccountListKeysResult
}

// AccountsClientListResponse contains the response from method AccountsClient.List.
type AccountsClientListResponse struct {
	AccountListResult
}

// AccountsClientListServiceSASResponse contains the response from method AccountsClient.ListServiceSAS.
type AccountsClientListServiceSASResponse struct {
	ListServiceSasResponse
}

// AccountsClientRegenerateKeyResponse contains the response from method AccountsClient.RegenerateKey.
type AccountsClientRegenerateKeyResponse struct {
	AccountListKeysResult
}

// AccountsClientRestoreBlobRangesResponse contains the response from method AccountsClient.RestoreBlobRanges.
type AccountsClientRestoreBlobRangesResponse struct {
	BlobRestoreStatus
}

// AccountsClientRevokeUserDelegationKeysResponse contains the response from method AccountsClient.RevokeUserDelegationKeys.
type AccountsClientRevokeUserDelegationKeysResponse struct {
	// placeholder for future response values
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	Account
}

// BlobContainersClientClearLegalHoldResponse contains the response from method BlobContainersClient.ClearLegalHold.
type BlobContainersClientClearLegalHoldResponse struct {
	LegalHold
}

// BlobContainersClientCreateOrUpdateImmutabilityPolicyResponse contains the response from method BlobContainersClient.CreateOrUpdateImmutabilityPolicy.
type BlobContainersClientCreateOrUpdateImmutabilityPolicyResponse struct {
	ImmutabilityPolicy
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientCreateResponse contains the response from method BlobContainersClient.Create.
type BlobContainersClientCreateResponse struct {
	BlobContainer
}

// BlobContainersClientDeleteImmutabilityPolicyResponse contains the response from method BlobContainersClient.DeleteImmutabilityPolicy.
type BlobContainersClientDeleteImmutabilityPolicyResponse struct {
	ImmutabilityPolicy
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientDeleteResponse contains the response from method BlobContainersClient.Delete.
type BlobContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// BlobContainersClientExtendImmutabilityPolicyResponse contains the response from method BlobContainersClient.ExtendImmutabilityPolicy.
type BlobContainersClientExtendImmutabilityPolicyResponse struct {
	ImmutabilityPolicy
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientGetImmutabilityPolicyResponse contains the response from method BlobContainersClient.GetImmutabilityPolicy.
type BlobContainersClientGetImmutabilityPolicyResponse struct {
	ImmutabilityPolicy
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientGetResponse contains the response from method BlobContainersClient.Get.
type BlobContainersClientGetResponse struct {
	BlobContainer
}

// BlobContainersClientLeaseResponse contains the response from method BlobContainersClient.Lease.
type BlobContainersClientLeaseResponse struct {
	LeaseContainerResponse
}

// BlobContainersClientListResponse contains the response from method BlobContainersClient.List.
type BlobContainersClientListResponse struct {
	ListContainerItems
}

// BlobContainersClientLockImmutabilityPolicyResponse contains the response from method BlobContainersClient.LockImmutabilityPolicy.
type BlobContainersClientLockImmutabilityPolicyResponse struct {
	ImmutabilityPolicy
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientObjectLevelWormResponse contains the response from method BlobContainersClient.ObjectLevelWorm.
type BlobContainersClientObjectLevelWormResponse struct {
	// placeholder for future response values
}

// BlobContainersClientSetLegalHoldResponse contains the response from method BlobContainersClient.SetLegalHold.
type BlobContainersClientSetLegalHoldResponse struct {
	LegalHold
}

// BlobContainersClientUpdateResponse contains the response from method BlobContainersClient.Update.
type BlobContainersClientUpdateResponse struct {
	BlobContainer
}

// BlobInventoryPoliciesClientCreateOrUpdateResponse contains the response from method BlobInventoryPoliciesClient.CreateOrUpdate.
type BlobInventoryPoliciesClientCreateOrUpdateResponse struct {
	BlobInventoryPolicy
}

// BlobInventoryPoliciesClientDeleteResponse contains the response from method BlobInventoryPoliciesClient.Delete.
type BlobInventoryPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// BlobInventoryPoliciesClientGetResponse contains the response from method BlobInventoryPoliciesClient.Get.
type BlobInventoryPoliciesClientGetResponse struct {
	BlobInventoryPolicy
}

// BlobInventoryPoliciesClientListResponse contains the response from method BlobInventoryPoliciesClient.List.
type BlobInventoryPoliciesClientListResponse struct {
	ListBlobInventoryPolicy
}

// BlobServicesClientGetServicePropertiesResponse contains the response from method BlobServicesClient.GetServiceProperties.
type BlobServicesClientGetServicePropertiesResponse struct {
	BlobServiceProperties
}

// BlobServicesClientListResponse contains the response from method BlobServicesClient.List.
type BlobServicesClientListResponse struct {
	BlobServiceItems
}

// BlobServicesClientSetServicePropertiesResponse contains the response from method BlobServicesClient.SetServiceProperties.
type BlobServicesClientSetServicePropertiesResponse struct {
	BlobServiceProperties
}

// DeletedAccountsClientGetResponse contains the response from method DeletedAccountsClient.Get.
type DeletedAccountsClientGetResponse struct {
	DeletedAccount
}

// DeletedAccountsClientListResponse contains the response from method DeletedAccountsClient.List.
type DeletedAccountsClientListResponse struct {
	DeletedAccountListResult
}

// EncryptionScopesClientGetResponse contains the response from method EncryptionScopesClient.Get.
type EncryptionScopesClientGetResponse struct {
	EncryptionScope
}

// EncryptionScopesClientListResponse contains the response from method EncryptionScopesClient.List.
type EncryptionScopesClientListResponse struct {
	EncryptionScopeListResult
}

// EncryptionScopesClientPatchResponse contains the response from method EncryptionScopesClient.Patch.
type EncryptionScopesClientPatchResponse struct {
	EncryptionScope
}

// EncryptionScopesClientPutResponse contains the response from method EncryptionScopesClient.Put.
type EncryptionScopesClientPutResponse struct {
	EncryptionScope
}

// FileServicesClientGetServicePropertiesResponse contains the response from method FileServicesClient.GetServiceProperties.
type FileServicesClientGetServicePropertiesResponse struct {
	FileServiceProperties
}

// FileServicesClientListResponse contains the response from method FileServicesClient.List.
type FileServicesClientListResponse struct {
	FileServiceItems
}

// FileServicesClientSetServicePropertiesResponse contains the response from method FileServicesClient.SetServiceProperties.
type FileServicesClientSetServicePropertiesResponse struct {
	FileServiceProperties
}

// FileSharesClientCreateResponse contains the response from method FileSharesClient.Create.
type FileSharesClientCreateResponse struct {
	FileShare
}

// FileSharesClientDeleteResponse contains the response from method FileSharesClient.Delete.
type FileSharesClientDeleteResponse struct {
	// placeholder for future response values
}

// FileSharesClientGetResponse contains the response from method FileSharesClient.Get.
type FileSharesClientGetResponse struct {
	FileShare
}

// FileSharesClientLeaseResponse contains the response from method FileSharesClient.Lease.
type FileSharesClientLeaseResponse struct {
	LeaseShareResponse
	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// FileSharesClientListResponse contains the response from method FileSharesClient.List.
type FileSharesClientListResponse struct {
	FileShareItems
}

// FileSharesClientRestoreResponse contains the response from method FileSharesClient.Restore.
type FileSharesClientRestoreResponse struct {
	// placeholder for future response values
}

// FileSharesClientUpdateResponse contains the response from method FileSharesClient.Update.
type FileSharesClientUpdateResponse struct {
	FileShare
}

// LocalUsersClientCreateOrUpdateResponse contains the response from method LocalUsersClient.CreateOrUpdate.
type LocalUsersClientCreateOrUpdateResponse struct {
	LocalUser
}

// LocalUsersClientDeleteResponse contains the response from method LocalUsersClient.Delete.
type LocalUsersClientDeleteResponse struct {
	// placeholder for future response values
}

// LocalUsersClientGetResponse contains the response from method LocalUsersClient.Get.
type LocalUsersClientGetResponse struct {
	LocalUser
}

// LocalUsersClientListKeysResponse contains the response from method LocalUsersClient.ListKeys.
type LocalUsersClientListKeysResponse struct {
	LocalUserKeys
}

// LocalUsersClientListResponse contains the response from method LocalUsersClient.List.
type LocalUsersClientListResponse struct {
	LocalUsers
}

// LocalUsersClientRegeneratePasswordResponse contains the response from method LocalUsersClient.RegeneratePassword.
type LocalUsersClientRegeneratePasswordResponse struct {
	LocalUserRegeneratePasswordResult
}

// ManagementPoliciesClientCreateOrUpdateResponse contains the response from method ManagementPoliciesClient.CreateOrUpdate.
type ManagementPoliciesClientCreateOrUpdateResponse struct {
	ManagementPolicy
}

// ManagementPoliciesClientDeleteResponse contains the response from method ManagementPoliciesClient.Delete.
type ManagementPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagementPoliciesClientGetResponse contains the response from method ManagementPoliciesClient.Get.
type ManagementPoliciesClientGetResponse struct {
	ManagementPolicy
}

// ObjectReplicationPoliciesClientCreateOrUpdateResponse contains the response from method ObjectReplicationPoliciesClient.CreateOrUpdate.
type ObjectReplicationPoliciesClientCreateOrUpdateResponse struct {
	ObjectReplicationPolicy
}

// ObjectReplicationPoliciesClientDeleteResponse contains the response from method ObjectReplicationPoliciesClient.Delete.
type ObjectReplicationPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// ObjectReplicationPoliciesClientGetResponse contains the response from method ObjectReplicationPoliciesClient.Get.
type ObjectReplicationPoliciesClientGetResponse struct {
	ObjectReplicationPolicy
}

// ObjectReplicationPoliciesClientListResponse contains the response from method ObjectReplicationPoliciesClient.List.
type ObjectReplicationPoliciesClientListResponse struct {
	ObjectReplicationPolicies
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientPutResponse contains the response from method PrivateEndpointConnectionsClient.Put.
type PrivateEndpointConnectionsClientPutResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientListByStorageAccountResponse contains the response from method PrivateLinkResourcesClient.ListByStorageAccount.
type PrivateLinkResourcesClientListByStorageAccountResponse struct {
	PrivateLinkResourceListResult
}

// QueueClientCreateResponse contains the response from method QueueClient.Create.
type QueueClientCreateResponse struct {
	Queue
}

// QueueClientDeleteResponse contains the response from method QueueClient.Delete.
type QueueClientDeleteResponse struct {
	// placeholder for future response values
}

// QueueClientGetResponse contains the response from method QueueClient.Get.
type QueueClientGetResponse struct {
	Queue
}

// QueueClientListResponse contains the response from method QueueClient.List.
type QueueClientListResponse struct {
	ListQueueResource
}

// QueueClientUpdateResponse contains the response from method QueueClient.Update.
type QueueClientUpdateResponse struct {
	Queue
}

// QueueServicesClientGetServicePropertiesResponse contains the response from method QueueServicesClient.GetServiceProperties.
type QueueServicesClientGetServicePropertiesResponse struct {
	QueueServiceProperties
}

// QueueServicesClientListResponse contains the response from method QueueServicesClient.List.
type QueueServicesClientListResponse struct {
	ListQueueServices
}

// QueueServicesClientSetServicePropertiesResponse contains the response from method QueueServicesClient.SetServiceProperties.
type QueueServicesClientSetServicePropertiesResponse struct {
	QueueServiceProperties
}

// SKUsClientListResponse contains the response from method SKUsClient.List.
type SKUsClientListResponse struct {
	SKUListResult
}

// TableClientCreateResponse contains the response from method TableClient.Create.
type TableClientCreateResponse struct {
	Table
}

// TableClientDeleteResponse contains the response from method TableClient.Delete.
type TableClientDeleteResponse struct {
	// placeholder for future response values
}

// TableClientGetResponse contains the response from method TableClient.Get.
type TableClientGetResponse struct {
	Table
}

// TableClientListResponse contains the response from method TableClient.List.
type TableClientListResponse struct {
	ListTableResource
}

// TableClientUpdateResponse contains the response from method TableClient.Update.
type TableClientUpdateResponse struct {
	Table
}

// TableServicesClientGetServicePropertiesResponse contains the response from method TableServicesClient.GetServiceProperties.
type TableServicesClientGetServicePropertiesResponse struct {
	TableServiceProperties
}

// TableServicesClientListResponse contains the response from method TableServicesClient.List.
type TableServicesClientListResponse struct {
	ListTableServices
}

// TableServicesClientSetServicePropertiesResponse contains the response from method TableServicesClient.SetServiceProperties.
type TableServicesClientSetServicePropertiesResponse struct {
	TableServiceProperties
}

// UsagesClientListByLocationResponse contains the response from method UsagesClient.ListByLocation.
type UsagesClientListByLocationResponse struct {
	UsageListResult
}
package storage

import "github.com/Azure/open-service-broker-azure/pkg/service"

type storageKind string

const (
	storageKindGeneralPurposeStorageAcccount storageKind = "GeneralPurposeStorageAccount" // nolint: lll
	storageKindBlobStorageAccount            storageKind = "BlobStorageAccount"
	storageKindBlobContainer                 storageKind = "BlobContainer"
)

// ProvisioningParameters encapsulates non-sensitive Storage-specific
// provisioning options
type ProvisioningParameters struct{}

// SecureProvisioningParameters encapsulates sensitive Storage-specific
// provisioning options
type SecureProvisioningParameters struct{}

type storageInstanceDetails struct {
	ARMDeploymentName  string `json:"armDeployment"`
	StorageAccountName string `json:"storageAccountName"`
	ContainerName      string `json:"containerName"`
}

type storageSecureInstanceDetails struct {
	AccessKey string `json:"accessKey"`
}

// UpdatingParameters encapsulates Storage-specific updating options
type UpdatingParameters struct {
}

// BindingParameters encapsulates non-sensitive Storage-specific binding options
type BindingParameters struct {
}

// SecureBindingParameters encapsulates sensitive Storage-specific binding
// options
type SecureBindingParameters struct {
}

type storageBindingDetails struct {
}

type storageSecureBindingDetails struct {
}

// Credentials encapsulates Storage-specific coonection details and credentials.
type Credentials struct {
	StorageAccountName string `json:"storageAccountName"`
	AccessKey          string `json:"accessKey"`
	ContainerName      string `json:"containerName,omitempty"`
}

func (
	s *serviceManager,
) GetEmptyProvisioningParameters() service.ProvisioningParameters {
	return &ProvisioningParameters{}
}

func (
	s *serviceManager,
) GetEmptySecureProvisioningParameters() service.SecureProvisioningParameters {
	return &SecureProvisioningParameters{}
}

func (
	s *serviceManager,
) GetEmptyInstanceDetails() service.InstanceDetails {
	return &storageInstanceDetails{}
}

func (
	s *serviceManager,
) GetEmptySecureInstanceDetails() service.SecureInstanceDetails {
	return &storageSecureInstanceDetails{}
}

func (s *serviceManager) GetEmptyBindingParameters() service.BindingParameters {
	return &BindingParameters{}
}

func (
	s *serviceManager,
) GetEmptySecureBindingParameters() service.SecureBindingParameters {
	return &SecureBindingParameters{}
}

func (s *serviceManager) GetEmptyBindingDetails() service.BindingDetails {
	return &storageBindingDetails{}
}

func (
	s *serviceManager,
) GetEmptySecureBindingDetails() service.SecureBindingDetails {
	return &storageSecureBindingDetails{}
}

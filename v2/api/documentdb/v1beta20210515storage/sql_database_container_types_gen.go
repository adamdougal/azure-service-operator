// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasecontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasecontainers/status,sqldatabasecontainers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210515.SqlDatabaseContainer
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers
type SqlDatabaseContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_SqlDatabases_Container_Spec `json:"spec,omitempty"`
	Status            SqlContainerGetResults_STATUS                `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainer{}

// GetConditions returns the conditions of the resource
func (container *SqlDatabaseContainer) GetConditions() conditions.Conditions {
	return container.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (container *SqlDatabaseContainer) SetConditions(conditions conditions.Conditions) {
	container.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainer{}

// AzureName returns the Azure name of the resource
func (container *SqlDatabaseContainer) AzureName() string {
	return container.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (container SqlDatabaseContainer) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (container *SqlDatabaseContainer) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (container *SqlDatabaseContainer) GetSpec() genruntime.ConvertibleSpec {
	return &container.Spec
}

// GetStatus returns the status of this resource
func (container *SqlDatabaseContainer) GetStatus() genruntime.ConvertibleStatus {
	return &container.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
func (container *SqlDatabaseContainer) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
}

// NewEmptyStatus returns a new empty (blank) status
func (container *SqlDatabaseContainer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlContainerGetResults_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (container *SqlDatabaseContainer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(container.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  container.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (container *SqlDatabaseContainer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlContainerGetResults_STATUS); ok {
		container.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlContainerGetResults_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	container.Status = st
	return nil
}

// Hub marks that this SqlDatabaseContainer is the hub type for conversion
func (container *SqlDatabaseContainer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (container *SqlDatabaseContainer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: container.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainer",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210515.SqlDatabaseContainer
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers
type SqlDatabaseContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainer `json:"items"`
}

// Storage version of v1beta20210515.DatabaseAccounts_SqlDatabases_Container_Spec
type DatabaseAccounts_SqlDatabases_Container_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabase resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabase"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlContainerResource              `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SqlDatabases_Container_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_SqlDatabases_Container_Spec from the provided source
func (container *DatabaseAccounts_SqlDatabases_Container_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == container {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(container)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Container_Spec
func (container *DatabaseAccounts_SqlDatabases_Container_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == container {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(container)
}

// Storage version of v1beta20210515.SqlContainerGetResults_STATUS
type SqlContainerGetResults_STATUS struct {
	Conditions  []conditions.Condition                     `json:"conditions,omitempty"`
	Id          *string                                    `json:"id,omitempty"`
	Location    *string                                    `json:"location,omitempty"`
	Name        *string                                    `json:"name,omitempty"`
	Options     *OptionsResource_STATUS                    `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                     `json:"$propertyBag,omitempty"`
	Resource    *SqlContainerGetProperties_STATUS_Resource `json:"resource,omitempty"`
	Tags        map[string]string                          `json:"tags,omitempty"`
	Type        *string                                    `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlContainerGetResults_STATUS{}

// ConvertStatusFrom populates our SqlContainerGetResults_STATUS from the provided source
func (results *SqlContainerGetResults_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == results {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(results)
}

// ConvertStatusTo populates the provided destination from our SqlContainerGetResults_STATUS
func (results *SqlContainerGetResults_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == results {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(results)
}

// Storage version of v1beta20210515.SqlContainerGetProperties_STATUS_Resource
type SqlContainerGetProperties_STATUS_Resource struct {
	AnalyticalStorageTtl     *int                             `json:"analyticalStorageTtl,omitempty"`
	ConflictResolutionPolicy *ConflictResolutionPolicy_STATUS `json:"conflictResolutionPolicy,omitempty"`
	DefaultTtl               *int                             `json:"defaultTtl,omitempty"`
	Etag                     *string                          `json:"_etag,omitempty"`
	Id                       *string                          `json:"id,omitempty"`
	IndexingPolicy           *IndexingPolicy_STATUS           `json:"indexingPolicy,omitempty"`
	PartitionKey             *ContainerPartitionKey_STATUS    `json:"partitionKey,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	Rid                      *string                          `json:"_rid,omitempty"`
	Ts                       *float64                         `json:"_ts,omitempty"`
	UniqueKeyPolicy          *UniqueKeyPolicy_STATUS          `json:"uniqueKeyPolicy,omitempty"`
}

// Storage version of v1beta20210515.SqlContainerResource
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerResource
type SqlContainerResource struct {
	AnalyticalStorageTtl     *int                      `json:"analyticalStorageTtl,omitempty"`
	ConflictResolutionPolicy *ConflictResolutionPolicy `json:"conflictResolutionPolicy,omitempty"`
	DefaultTtl               *int                      `json:"defaultTtl,omitempty"`
	Id                       *string                   `json:"id,omitempty"`
	IndexingPolicy           *IndexingPolicy           `json:"indexingPolicy,omitempty"`
	PartitionKey             *ContainerPartitionKey    `json:"partitionKey,omitempty"`
	PropertyBag              genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `json:"uniqueKeyPolicy,omitempty"`
}

// Storage version of v1beta20210515.ConflictResolutionPolicy
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConflictResolutionPolicy
type ConflictResolutionPolicy struct {
	ConflictResolutionPath      *string                `json:"conflictResolutionPath,omitempty"`
	ConflictResolutionProcedure *string                `json:"conflictResolutionProcedure,omitempty"`
	Mode                        *string                `json:"mode,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.ConflictResolutionPolicy_STATUS
type ConflictResolutionPolicy_STATUS struct {
	ConflictResolutionPath      *string                `json:"conflictResolutionPath,omitempty"`
	ConflictResolutionProcedure *string                `json:"conflictResolutionProcedure,omitempty"`
	Mode                        *string                `json:"mode,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.ContainerPartitionKey
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContainerPartitionKey
type ContainerPartitionKey struct {
	Kind        *string                `json:"kind,omitempty"`
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Version     *int                   `json:"version,omitempty"`
}

// Storage version of v1beta20210515.ContainerPartitionKey_STATUS
type ContainerPartitionKey_STATUS struct {
	Kind        *string                `json:"kind,omitempty"`
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemKey   *bool                  `json:"systemKey,omitempty"`
	Version     *int                   `json:"version,omitempty"`
}

// Storage version of v1beta20210515.IndexingPolicy
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IndexingPolicy
type IndexingPolicy struct {
	Automatic        *bool                  `json:"automatic,omitempty"`
	CompositeIndexes [][]CompositePath      `json:"compositeIndexes,omitempty"`
	ExcludedPaths    []ExcludedPath         `json:"excludedPaths,omitempty"`
	IncludedPaths    []IncludedPath         `json:"includedPaths,omitempty"`
	IndexingMode     *string                `json:"indexingMode,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SpatialIndexes   []SpatialSpec          `json:"spatialIndexes,omitempty"`
}

// Storage version of v1beta20210515.IndexingPolicy_STATUS
type IndexingPolicy_STATUS struct {
	Automatic        *bool                    `json:"automatic,omitempty"`
	CompositeIndexes [][]CompositePath_STATUS `json:"compositeIndexes,omitempty"`
	ExcludedPaths    []ExcludedPath_STATUS    `json:"excludedPaths,omitempty"`
	IncludedPaths    []IncludedPath_STATUS    `json:"includedPaths,omitempty"`
	IndexingMode     *string                  `json:"indexingMode,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	SpatialIndexes   []SpatialSpec_STATUS     `json:"spatialIndexes,omitempty"`
}

// Storage version of v1beta20210515.UniqueKeyPolicy
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKeyPolicy
type UniqueKeyPolicy struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UniqueKeys  []UniqueKey            `json:"uniqueKeys,omitempty"`
}

// Storage version of v1beta20210515.UniqueKeyPolicy_STATUS
type UniqueKeyPolicy_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UniqueKeys  []UniqueKey_STATUS     `json:"uniqueKeys,omitempty"`
}

// Storage version of v1beta20210515.CompositePath
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CompositePath
type CompositePath struct {
	Order       *string                `json:"order,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.CompositePath_STATUS
type CompositePath_STATUS struct {
	Order       *string                `json:"order,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.ExcludedPath
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ExcludedPath
type ExcludedPath struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.ExcludedPath_STATUS
type ExcludedPath_STATUS struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.IncludedPath
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IncludedPath
type IncludedPath struct {
	Indexes     []Indexes              `json:"indexes,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.IncludedPath_STATUS
type IncludedPath_STATUS struct {
	Indexes     []Indexes_STATUS       `json:"indexes,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.SpatialSpec
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SpatialSpec
type SpatialSpec struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Types       []string               `json:"types,omitempty"`
}

// Storage version of v1beta20210515.SpatialSpec_STATUS
type SpatialSpec_STATUS struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Types       []string               `json:"types,omitempty"`
}

// Storage version of v1beta20210515.UniqueKey
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKey
type UniqueKey struct {
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.UniqueKey_STATUS
type UniqueKey_STATUS struct {
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.Indexes
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Indexes
type Indexes struct {
	DataType    *string                `json:"dataType,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	Precision   *int                   `json:"precision,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.Indexes_STATUS
type Indexes_STATUS struct {
	DataType    *string                `json:"dataType,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	Precision   *int                   `json:"precision,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainer{}, &SqlDatabaseContainerList{})
}

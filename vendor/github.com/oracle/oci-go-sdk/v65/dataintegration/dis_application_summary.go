// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DisApplicationSummary The application summary type contains the audit summary information and the definition of the application.
type DisApplicationSummary struct {

	// Generated key that can be used in API calls to identify application.
	Key *string `mandatory:"false" json:"key"`

	// The object type.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The application's version.
	ApplicationVersion *int `mandatory:"false" json:"applicationVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// A list of dependent objects in this patch.
	DependentObjectMetadata []PatchObjectMetadata `mandatory:"false" json:"dependentObjectMetadata"`

	// A list of objects that are published or unpublished in this patch.
	PublishedObjectMetadata map[string]PatchObjectMetadata `mandatory:"false" json:"publishedObjectMetadata"`

	SourceApplicationInfo *SourceApplicationInfo `mandatory:"false" json:"sourceApplicationInfo"`

	// The date and time the application was patched, in the timestamp format defined by RFC3339.
	TimePatched *common.SDKTime `mandatory:"false" json:"timePatched"`

	// OCID of the resource that is used to uniquely identify the application
	Id *string `mandatory:"false" json:"id"`

	// OCID of the compartment that this resource belongs to. Defaults to compartment of the Workspace.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the application was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the application was updated, in the timestamp format defined by RFC3339.
	// example: 2019-08-25T21:10:29.41Z
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current state of the workspace.
	LifecycleState DisApplicationSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m DisApplicationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisApplicationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDisApplicationSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDisApplicationSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DisApplicationSummaryLifecycleStateEnum Enum with underlying type: string
type DisApplicationSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DisApplicationSummaryLifecycleStateEnum
const (
	DisApplicationSummaryLifecycleStateCreating DisApplicationSummaryLifecycleStateEnum = "CREATING"
	DisApplicationSummaryLifecycleStateActive   DisApplicationSummaryLifecycleStateEnum = "ACTIVE"
	DisApplicationSummaryLifecycleStateUpdating DisApplicationSummaryLifecycleStateEnum = "UPDATING"
	DisApplicationSummaryLifecycleStateDeleting DisApplicationSummaryLifecycleStateEnum = "DELETING"
	DisApplicationSummaryLifecycleStateDeleted  DisApplicationSummaryLifecycleStateEnum = "DELETED"
	DisApplicationSummaryLifecycleStateFailed   DisApplicationSummaryLifecycleStateEnum = "FAILED"
)

var mappingDisApplicationSummaryLifecycleStateEnum = map[string]DisApplicationSummaryLifecycleStateEnum{
	"CREATING": DisApplicationSummaryLifecycleStateCreating,
	"ACTIVE":   DisApplicationSummaryLifecycleStateActive,
	"UPDATING": DisApplicationSummaryLifecycleStateUpdating,
	"DELETING": DisApplicationSummaryLifecycleStateDeleting,
	"DELETED":  DisApplicationSummaryLifecycleStateDeleted,
	"FAILED":   DisApplicationSummaryLifecycleStateFailed,
}

var mappingDisApplicationSummaryLifecycleStateEnumLowerCase = map[string]DisApplicationSummaryLifecycleStateEnum{
	"creating": DisApplicationSummaryLifecycleStateCreating,
	"active":   DisApplicationSummaryLifecycleStateActive,
	"updating": DisApplicationSummaryLifecycleStateUpdating,
	"deleting": DisApplicationSummaryLifecycleStateDeleting,
	"deleted":  DisApplicationSummaryLifecycleStateDeleted,
	"failed":   DisApplicationSummaryLifecycleStateFailed,
}

// GetDisApplicationSummaryLifecycleStateEnumValues Enumerates the set of values for DisApplicationSummaryLifecycleStateEnum
func GetDisApplicationSummaryLifecycleStateEnumValues() []DisApplicationSummaryLifecycleStateEnum {
	values := make([]DisApplicationSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDisApplicationSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDisApplicationSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DisApplicationSummaryLifecycleStateEnum
func GetDisApplicationSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDisApplicationSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisApplicationSummaryLifecycleStateEnum(val string) (DisApplicationSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDisApplicationSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

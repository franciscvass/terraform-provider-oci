// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// OperationStatusEnum Enum with underlying type: string
type OperationStatusEnum string

// Set of constants representing the allowable values for OperationStatusEnum
const (
	OperationStatusAccepted   OperationStatusEnum = "ACCEPTED"
	OperationStatusCanceled   OperationStatusEnum = "CANCELED"
	OperationStatusCanceling  OperationStatusEnum = "CANCELING"
	OperationStatusFailed     OperationStatusEnum = "FAILED"
	OperationStatusInProgress OperationStatusEnum = "IN_PROGRESS"
	OperationStatusSucceeded  OperationStatusEnum = "SUCCEEDED"
)

var mappingOperationStatusEnum = map[string]OperationStatusEnum{
	"ACCEPTED":    OperationStatusAccepted,
	"CANCELED":    OperationStatusCanceled,
	"CANCELING":   OperationStatusCanceling,
	"FAILED":      OperationStatusFailed,
	"IN_PROGRESS": OperationStatusInProgress,
	"SUCCEEDED":   OperationStatusSucceeded,
}

var mappingOperationStatusEnumLowerCase = map[string]OperationStatusEnum{
	"accepted":    OperationStatusAccepted,
	"canceled":    OperationStatusCanceled,
	"canceling":   OperationStatusCanceling,
	"failed":      OperationStatusFailed,
	"in_progress": OperationStatusInProgress,
	"succeeded":   OperationStatusSucceeded,
}

// GetOperationStatusEnumValues Enumerates the set of values for OperationStatusEnum
func GetOperationStatusEnumValues() []OperationStatusEnum {
	values := make([]OperationStatusEnum, 0)
	for _, v := range mappingOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationStatusEnumStringValues Enumerates the set of values in String for OperationStatusEnum
func GetOperationStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"CANCELED",
		"CANCELING",
		"FAILED",
		"IN_PROGRESS",
		"SUCCEEDED",
	}
}

// GetMappingOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationStatusEnum(val string) (OperationStatusEnum, bool) {
	enum, ok := mappingOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

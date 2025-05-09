// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VcnDrgAttachmentNetworkUpdateDetails Specifies the update details for the VCN attachment.
type VcnDrgAttachmentNetworkUpdateDetails struct {

	// This is the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table that is used to route the traffic as it enters a VCN through this attachment.
	// For information about why you would associate a route table with a DRG attachment, see:
	//   * Transit Routing: Access to Multiple VCNs in Same Region (https://docs.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm)
	//   * Transit Routing: Private Access to Oracle Services (https://docs.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm)
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// Indicates whether the VCN CIDRs or the individual subnet CIDRs are imported from the attachment.
	// Routes from the VCN ingress route table are always imported.
	VcnRouteType VcnDrgAttachmentNetworkDetailsVcnRouteTypeEnum `mandatory:"false" json:"vcnRouteType,omitempty"`
}

func (m VcnDrgAttachmentNetworkUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VcnDrgAttachmentNetworkUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVcnDrgAttachmentNetworkDetailsVcnRouteTypeEnum(string(m.VcnRouteType)); !ok && m.VcnRouteType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VcnRouteType: %s. Supported values are: %s.", m.VcnRouteType, strings.Join(GetVcnDrgAttachmentNetworkDetailsVcnRouteTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VcnDrgAttachmentNetworkUpdateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVcnDrgAttachmentNetworkUpdateDetails VcnDrgAttachmentNetworkUpdateDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVcnDrgAttachmentNetworkUpdateDetails
	}{
		"VCN",
		(MarshalTypeVcnDrgAttachmentNetworkUpdateDetails)(m),
	}

	return json.Marshal(&s)
}

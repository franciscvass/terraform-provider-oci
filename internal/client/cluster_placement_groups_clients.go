// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_cluster_placement_groups "github.com/oracle/oci-go-sdk/v65/clusterplacementgroups"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_cluster_placement_groups.ClusterPlacementGroupsCPClient", &OracleClient{InitClientFn: initClusterplacementgroupsClusterPlacementGroupsCPClient})
}

func initClusterplacementgroupsClusterPlacementGroupsCPClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_cluster_placement_groups.NewClusterPlacementGroupsCPClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ClusterPlacementGroupsCPClient() *oci_cluster_placement_groups.ClusterPlacementGroupsCPClient {
	return m.GetClient("oci_cluster_placement_groups.ClusterPlacementGroupsCPClient").(*oci_cluster_placement_groups.ClusterPlacementGroupsCPClient)
}

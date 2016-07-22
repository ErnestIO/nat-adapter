/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bytes"
	"encoding/json"
)

type rule struct {
	Type            string `json:"type"`
	OriginIP        string `json:"origin_ip"`
	OriginPort      string `json:"origin_port"`
	TranslationIP   string `json:"translation_ip"`
	TranslationPort string `json:"translation_port"`
	Protocol        string `json:"protocol"`
	Network         string `json:"network"`
}

type builderEvent struct {
	Uuid                  string `json:"_uuid"`
	BatchID               string `json:"_batch_id"`
	Type                  string `json:"type"`
	Service               string `json:"service"`
	Name                  string `json:"name"`
	Rules                 []rule `json:"rules"`
	RouterName            string `json:"router_name"`
	RouterType            string `json:"router_type"`
	RouterIP              string `json:"router_ip"`
	ClientName            string `json:"client_name"`
	DatacenterName        string `json:"datacenter_name"`
	DatacenterPassword    string `json:"datacenter_password"`
	DatacenterRegion      string `json:"datacenter_region"`
	DatacenterType        string `json:"datacenter_type"`
	DatacenterUsername    string `json:"datacenter_username"`
	DatacenterAccessToken string `json:"datacenter_token"`
	DatacenterAccessKey   string `json:"datacenter_secret"`
	NetworkName           string `json:"network_name"`
	SecurityGroupAWSIDs   string `json:"security_group_aws_ids"`
	NatGatewayAWSID       string `json:"nat_gateway_aws_id"`
	VCloudURL             string `json:"vcloud_url"`
	Status                string `json:"status"`
	ErrorCode             string `json:"error_code"`
	ErrorMessage          string `json:"error_message"`
}

type vcloudEvent struct {
	Uuid               string `json:"_uuid"`
	BatchID            string `json:"_batch_id"`
	Type               string `json:"_type"`
	Service            string `json:"service_id"`
	Name               string `json:"nat_name"`
	Rules              []rule `json:"nat_rules"`
	RouterName         string `json:"router_name"`
	RouterType         string `json:"router_type"`
	RouterIP           string `json:"router_ip"`
	ClientName         string `json:"client_name"`
	DatacenterName     string `json:"datacenter_name"`
	DatacenterPassword string `json:"datacenter_password"`
	DatacenterRegion   string `json:"datacenter_region"`
	DatacenterType     string `json:"datacenter_type"`
	DatacenterUsername string `json:"datacenter_username"`
	NetworkName        string `json:"network_name"`
	VCloudURL          string `json:"vcloud_url"`
	Status             string `json:"status"`
	ErrorCode          string `json:"error_code"`
	ErrorMessage       string `json:"error_message"`
}

type awsEvent struct {
	Uuid                  string `json:"_uuid"`
	BatchID               string `json:"_batch_id"`
	Type                  string `json:"_type"`
	DatacenterRegion      string `json:"datacenter_region"`
	DatacenterAccessToken string `json:"datacenter_access_token"`
	DatacenterAccessKey   string `json:"datacenter_access_key"`
	DatacenterVPCID       string `json:"datacenter_vpc_id"`
	NatGatewayAWSID       string `json:"nat_gateway_aws_id"`
	NetworkAWSID          string `json:"network_aws_id"`
	Status                string `json:"status"`
	ErrorCode             string `json:"error_code"`
	ErrorMessage          string `json:"error_message"`
}

type Translator struct{}

func (t Translator) BuilderToConnector(j []byte) []byte {
	var input builderEvent
	var output []byte
	json.Unmarshal(j, &input)

	switch input.RouterType {
	case "vcloud", "fake-vcloud", "fake":
		output = t.builderToVCloudConnector(input)
	case "aws", "fake-aws":
		output = t.builderToAwsConnector(input)
	}

	return output
}

func (t Translator) builderToVCloudConnector(input builderEvent) []byte {
	var output vcloudEvent

	output.Uuid = input.Uuid
	output.BatchID = input.BatchID
	output.Service = input.Service
	output.Type = input.RouterType
	output.Name = input.Name
	output.Rules = input.Rules
	output.RouterIP = input.RouterIP
	output.RouterName = input.RouterName
	output.RouterType = input.RouterType
	output.NetworkName = input.NetworkName
	output.ClientName = input.ClientName
	output.DatacenterName = input.DatacenterName
	output.DatacenterRegion = input.DatacenterRegion
	output.DatacenterUsername = input.DatacenterUsername
	output.DatacenterPassword = input.DatacenterPassword
	output.DatacenterType = input.DatacenterType
	output.VCloudURL = input.VCloudURL
	output.Status = input.Status
	output.ErrorCode = input.ErrorCode
	output.ErrorMessage = input.ErrorMessage

	body, _ := json.Marshal(output)

	return body
}

func (t Translator) builderToAwsConnector(input builderEvent) []byte {
	var output awsEvent

	output.Uuid = input.Uuid
	output.BatchID = input.BatchID
	output.Type = input.RouterType
	output.DatacenterRegion = input.DatacenterRegion
	output.DatacenterAccessToken = input.DatacenterAccessToken
	output.DatacenterAccessKey = input.DatacenterAccessKey
	output.DatacenterVPCID = input.DatacenterName
	output.NatGatewayAWSID = input.NatGatewayAWSID
	output.Status = input.Status
	output.ErrorCode = input.ErrorCode
	output.ErrorMessage = input.ErrorMessage

	body, _ := json.Marshal(output)

	return body
}

func (t Translator) ConnectorToBuilder(j []byte) []byte {
	var output []byte
	var input map[string]interface{}

	dec := json.NewDecoder(bytes.NewReader(j))
	dec.Decode(&input)

	switch input["_type"] {
	case "vcloud", "fake-vcloud", "fake":
		output = t.vcloudConnectorToBuilder(j)
	case "aws", "fake-aws":
		output = t.awsConnectorToBuilder(j)
	}

	return output
}

func (t Translator) vcloudConnectorToBuilder(j []byte) []byte {
	var input vcloudEvent
	var output builderEvent
	json.Unmarshal(j, &input)

	output.Uuid = input.Uuid
	output.BatchID = input.BatchID
	output.RouterType = input.Type
	output.Name = input.Name
	output.Rules = input.Rules
	output.RouterIP = input.RouterIP
	output.RouterName = input.RouterName
	output.RouterType = input.RouterType
	output.NetworkName = input.NetworkName
	output.ClientName = input.ClientName
	output.DatacenterName = input.DatacenterName
	output.DatacenterRegion = input.DatacenterRegion
	output.DatacenterUsername = input.DatacenterUsername
	output.DatacenterPassword = input.DatacenterPassword
	output.DatacenterType = input.DatacenterType
	output.VCloudURL = input.VCloudURL
	output.Status = input.Status
	output.ErrorCode = input.ErrorCode
	output.ErrorMessage = input.ErrorMessage

	body, _ := json.Marshal(output)

	return body
}

func (t Translator) awsConnectorToBuilder(j []byte) []byte {
	var input awsEvent
	var output builderEvent
	json.Unmarshal(j, &input)

	output.Uuid = input.Uuid
	output.BatchID = input.BatchID
	output.Type = input.Type
	/*
		output.DatacenterRegion = input.DatacenterRegion
		output.DatacenterAccessToken = input.DatacenterAccessToken
		output.DatacenterAccessKey = input.DatacenterAccessKey
		output.DatacenterName = input.DatacenterVpcID
		output.NetworkName = input.NetworkAWSID
		output.SecurityGroupAWSIDs = input.SecurityGroupAWSIDs
		// TODO: Documentation says something about Private IPS, but can't find any specs about it
		output.Name = input.InstanceName
		output.Image = input.InstanceImage
		output.Type = input.InstanceType
	*/
	output.Status = input.Status
	output.ErrorCode = input.ErrorCode
	output.ErrorMessage = input.ErrorMessage

	body, _ := json.Marshal(output)

	return body
}

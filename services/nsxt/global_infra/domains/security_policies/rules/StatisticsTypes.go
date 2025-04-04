// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Statistics.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package rules

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func statisticsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fields["container_cluster_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["container_cluster_path"] = "ContainerClusterPath"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatisticsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RuleStatisticsListResultBindingType)
}

func statisticsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fields["container_cluster_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["container_cluster_path"] = "ContainerClusterPath"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["container_cluster_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	paramsTypeMap["ruleId"] = vapiBindings_.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["rule_id"] = "ruleId"
	pathParams["domain_id"] = "domainId"
	queryParams["container_cluster_path"] = "container_cluster_path"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/policy/api/v1/global-infra/domains/{domainId}/security-policies/{securityPolicyId}/rules/{ruleId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ContextProfilesCustomTlds.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package infra

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``action`` of method ContextProfilesCustomTlds#create.
const ContextProfilesCustomTlds_CREATE_ACTION_ADD = "ADD"

// Possible value for ``action`` of method ContextProfilesCustomTlds#create.
const ContextProfilesCustomTlds_CREATE_ACTION_REMOVE = "REMOVE"

func contextProfilesCustomTldsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy_context_profile_fqdn_custom_tlds"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyContextProfileFqdnCustomTldsBindingType)
	fields["action"] = vapiBindings_.NewStringType()
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["policy_context_profile_fqdn_custom_tlds"] = "PolicyContextProfileFqdnCustomTlds"
	fieldNameMap["action"] = "Action"
	fieldNameMap["force"] = "Force"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ContextProfilesCustomTldsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func contextProfilesCustomTldsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["policy_context_profile_fqdn_custom_tlds"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyContextProfileFqdnCustomTldsBindingType)
	fields["action"] = vapiBindings_.NewStringType()
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["policy_context_profile_fqdn_custom_tlds"] = "PolicyContextProfileFqdnCustomTlds"
	fieldNameMap["action"] = "Action"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["policy_context_profile_fqdn_custom_tlds"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyContextProfileFqdnCustomTldsBindingType)
	paramsTypeMap["action"] = vapiBindings_.NewStringType()
	paramsTypeMap["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	queryParams["action"] = "action"
	queryParams["force"] = "force"
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
		"policy_context_profile_fqdn_custom_tlds",
		"POST",
		"/policy/api/v1/infra/context-profiles-custom-tlds",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func contextProfilesCustomTldsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ContextProfilesCustomTldsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyContextProfileFqdnCustomTldsBindingType)
}

func contextProfilesCustomTldsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/policy/api/v1/infra/context-profiles-custom-tlds",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func contextProfilesCustomTldsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy_context_profile_fqdn_custom_tlds"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyContextProfileFqdnCustomTldsBindingType)
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["policy_context_profile_fqdn_custom_tlds"] = "PolicyContextProfileFqdnCustomTlds"
	fieldNameMap["force"] = "Force"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ContextProfilesCustomTldsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func contextProfilesCustomTldsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["policy_context_profile_fqdn_custom_tlds"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyContextProfileFqdnCustomTldsBindingType)
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["policy_context_profile_fqdn_custom_tlds"] = "PolicyContextProfileFqdnCustomTlds"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["policy_context_profile_fqdn_custom_tlds"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyContextProfileFqdnCustomTldsBindingType)
	paramsTypeMap["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	queryParams["force"] = "force"
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
		"policy_context_profile_fqdn_custom_tlds",
		"PATCH",
		"/policy/api/v1/infra/context-profiles-custom-tlds",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

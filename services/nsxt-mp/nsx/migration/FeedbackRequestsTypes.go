// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FeedbackRequests.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package migration

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``networkLayer`` of method FeedbackRequests#list.
const FeedbackRequests_LIST_NETWORK_LAYER_L2 = "L2"

// Possible value for ``networkLayer`` of method FeedbackRequests#list.
const FeedbackRequests_LIST_NETWORK_LAYER_L3_L7 = "L3_L7"

// Possible value for ``state`` of method FeedbackRequests#list.
const FeedbackRequests_LIST_STATE_ALL = "ALL"

// Possible value for ``state`` of method FeedbackRequests#list.
const FeedbackRequests_LIST_STATE_RESOLVED = "RESOLVED"

// Possible value for ``state`` of method FeedbackRequests#list.
const FeedbackRequests_LIST_STATE_UNRESOLVED = "UNRESOLVED"

func feedbackRequestsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["federation_site_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["hash"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_layer"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sub_category"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["category"] = "Category"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	fieldNameMap["hash"] = "Hash"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_layer"] = "NetworkLayer"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["state"] = "State"
	fieldNameMap["sub_category"] = "SubCategory"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FeedbackRequestsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.MigrationFeedbackRequestListResultBindingType)
}

func feedbackRequestsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["category"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["federation_site_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["hash"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_layer"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sub_category"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["category"] = "Category"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	fieldNameMap["hash"] = "Hash"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_layer"] = "NetworkLayer"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["state"] = "State"
	fieldNameMap["sub_category"] = "SubCategory"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["network_layer"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sub_category"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["federation_site_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["category"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["hash"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["network_layer"] = "network_layer"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sub_category"] = "sub_category"
	queryParams["federation_site_id"] = "federation_site_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["state"] = "state"
	queryParams["category"] = "category"
	queryParams["hash"] = "hash"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/api/v1/migration/feedback-requests",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

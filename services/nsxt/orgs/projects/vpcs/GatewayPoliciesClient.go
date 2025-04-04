// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: GatewayPolicies
// Used by client-side stubs.

package vpcs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type GatewayPoliciesClient interface {

	// Delete Gateway Policy.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param gatewayPolicyIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string) error

	// Get Gateway Policy
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param gatewayPolicyIdParam (required)
	// @return com.vmware.nsx_policy.model.GatewayPolicy
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string) (nsx_policyModel.GatewayPolicy, error)

	// Paginated list of VPC gateway policies.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includeRuleCountParam Include the count of rules in policy (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.GatewayPolicyListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, vpcIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includeRuleCountParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.GatewayPolicyListResult, error)

	// If a GatewayPolicy with the policy-id is not already present, create a new GatewayPolicy. If it already exists, update the GatewayPolicy. This is a full replace.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param gatewayPolicyIdParam (required)
	// @param gatewayPolicyParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string, gatewayPolicyParam nsx_policyModel.GatewayPolicy) error

	// This is used to set a precedence of a VPC gateway policy w.r.t others.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param gatewayPolicyIdParam (required)
	// @param gatewayPolicyParam (required)
	// @param anchorPathParam The security policy/rule path if operation is 'insert_after' or 'insert_before' (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx_policy.model.GatewayPolicy
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revise(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string, gatewayPolicyParam nsx_policyModel.GatewayPolicy, anchorPathParam *string, operationParam *string) (nsx_policyModel.GatewayPolicy, error)

	// Update the Gateway Policy.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param gatewayPolicyIdParam (required)
	// @param gatewayPolicyParam (required)
	// @return com.vmware.nsx_policy.model.GatewayPolicy
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string, gatewayPolicyParam nsx_policyModel.GatewayPolicy) (nsx_policyModel.GatewayPolicy, error)
}

type gatewayPoliciesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewGatewayPoliciesClient(connector vapiProtocolClient_.Connector) *gatewayPoliciesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.vpcs.gateway_policies")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"revise": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "revise"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	gIface := gatewayPoliciesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &gIface
}

func (gIface *gatewayPoliciesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := gIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (gIface *gatewayPoliciesClient) Delete(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string) error {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayPoliciesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayPoliciesDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("GatewayPolicyId", gatewayPolicyIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.gateway_policies", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (gIface *gatewayPoliciesClient) Get(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string) (nsx_policyModel.GatewayPolicy, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayPoliciesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayPoliciesGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("GatewayPolicyId", gatewayPolicyIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GatewayPolicy
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.gateway_policies", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GatewayPolicy
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GatewayPoliciesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GatewayPolicy), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *gatewayPoliciesClient) List(orgIdParam string, projectIdParam string, vpcIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includeRuleCountParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.GatewayPolicyListResult, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayPoliciesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayPoliciesListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludeRuleCount", includeRuleCountParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GatewayPolicyListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.gateway_policies", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GatewayPolicyListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GatewayPoliciesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GatewayPolicyListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *gatewayPoliciesClient) Patch(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string, gatewayPolicyParam nsx_policyModel.GatewayPolicy) error {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayPoliciesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayPoliciesPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("GatewayPolicyId", gatewayPolicyIdParam)
	sv.AddStructField("GatewayPolicy", gatewayPolicyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.gateway_policies", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (gIface *gatewayPoliciesClient) Revise(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string, gatewayPolicyParam nsx_policyModel.GatewayPolicy, anchorPathParam *string, operationParam *string) (nsx_policyModel.GatewayPolicy, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayPoliciesReviseRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayPoliciesReviseInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("GatewayPolicyId", gatewayPolicyIdParam)
	sv.AddStructField("GatewayPolicy", gatewayPolicyParam)
	sv.AddStructField("AnchorPath", anchorPathParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GatewayPolicy
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.gateway_policies", "revise", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GatewayPolicy
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GatewayPoliciesReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GatewayPolicy), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *gatewayPoliciesClient) Update(orgIdParam string, projectIdParam string, vpcIdParam string, gatewayPolicyIdParam string, gatewayPolicyParam nsx_policyModel.GatewayPolicy) (nsx_policyModel.GatewayPolicy, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayPoliciesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayPoliciesUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("GatewayPolicyId", gatewayPolicyIdParam)
	sv.AddStructField("GatewayPolicy", gatewayPolicyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GatewayPolicy
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.gateway_policies", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GatewayPolicy
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GatewayPoliciesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GatewayPolicy), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

package batch

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// PoolClient is the client for the Pool methods of the Batch service.
type PoolClient struct {
	ManagementClient
}

// NewPoolClient creates an instance of the PoolClient client.
func NewPoolClient(subscriptionID string) PoolClient {
	return NewPoolClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPoolClientWithBaseURI creates an instance of the PoolClient client.
func NewPoolClientWithBaseURI(baseURI string, subscriptionID string) PoolClient {
	return PoolClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new pool inside the specified account. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group that contains the Batch account. accountName is the name of the
// Batch account. poolName is the pool name. This must be unique within the account. parameters is additional
// parameters for pool creation. ifMatch is the entity state (ETag) version of the pool to update. A value of "*" can
// be used to apply the operation only if the pool already exists. If omitted, this operation will always be applied.
// ifNoneMatch is set to '*' to allow a new pool to be created, but to prevent updating an existing pool. Other values
// will be ignored.
func (client PoolClient) Create(resourceGroupName string, accountName string, poolName string, parameters Pool, ifMatch string, ifNoneMatch string, cancel <-chan struct{}) (<-chan Pool, <-chan error) {
	resultChan := make(chan Pool, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.PoolProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.PoolProperties.DeploymentConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.PoolProperties.DeploymentConfiguration.CloudServiceConfiguration", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PoolProperties.DeploymentConfiguration.CloudServiceConfiguration.OsFamily", Name: validation.Null, Rule: true, Chain: nil}}},
						{Target: "parameters.PoolProperties.DeploymentConfiguration.VirtualMachineConfiguration", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.PoolProperties.DeploymentConfiguration.VirtualMachineConfiguration.ImageReference", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.PoolProperties.DeploymentConfiguration.VirtualMachineConfiguration.NodeAgentSkuID", Name: validation.Null, Rule: true, Chain: nil},
							}},
					}},
					{Target: "parameters.PoolProperties.ScaleSettings", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PoolProperties.ScaleSettings.AutoScale", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.PoolProperties.ScaleSettings.AutoScale.Formula", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					{Target: "parameters.PoolProperties.AutoScaleRun", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PoolProperties.AutoScaleRun.EvaluationTime", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.PoolProperties.AutoScaleRun.Error", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.PoolProperties.AutoScaleRun.Error.Code", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.PoolProperties.AutoScaleRun.Error.Message", Name: validation.Null, Rule: true, Chain: nil},
								}},
						}},
					{Target: "parameters.PoolProperties.NetworkConfiguration", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PoolProperties.NetworkConfiguration.EndpointConfiguration", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.PoolProperties.NetworkConfiguration.EndpointConfiguration.InboundNatPools", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "batch.PoolClient", "Create")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result Pool
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(resourceGroupName, accountName, poolName, parameters, ifMatch, ifNoneMatch, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "batch.PoolClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "batch.PoolClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "batch.PoolClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client PoolClient) CreatePreparer(resourceGroupName string, accountName string, poolName string, parameters Pool, ifMatch string, ifNoneMatch string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PoolClient) CreateResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified pool. This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the Batch account. accountName is the name of the
// Batch account. poolName is the pool name. This must be unique within the account.
func (client PoolClient) Delete(resourceGroupName string, accountName string, poolName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "batch.PoolClient", "Delete")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, accountName, poolName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "batch.PoolClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "batch.PoolClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "batch.PoolClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client PoolClient) DeletePreparer(resourceGroupName string, accountName string, poolName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PoolClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisableAutoScale disables automatic scaling for a pool.
//
// resourceGroupName is the name of the resource group that contains the Batch account. accountName is the name of the
// Batch account. poolName is the pool name. This must be unique within the account.
func (client PoolClient) DisableAutoScale(resourceGroupName string, accountName string, poolName string) (result Pool, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "batch.PoolClient", "DisableAutoScale")
	}

	req, err := client.DisableAutoScalePreparer(resourceGroupName, accountName, poolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "DisableAutoScale", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableAutoScaleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "DisableAutoScale", resp, "Failure sending request")
		return
	}

	result, err = client.DisableAutoScaleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "DisableAutoScale", resp, "Failure responding to request")
	}

	return
}

// DisableAutoScalePreparer prepares the DisableAutoScale request.
func (client PoolClient) DisableAutoScalePreparer(resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}/disableAutoScale", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DisableAutoScaleSender sends the DisableAutoScale request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) DisableAutoScaleSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// DisableAutoScaleResponder handles the response to the DisableAutoScale request. The method always
// closes the http.Response Body.
func (client PoolClient) DisableAutoScaleResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets information about the specified pool.
//
// resourceGroupName is the name of the resource group that contains the Batch account. accountName is the name of the
// Batch account. poolName is the pool name. This must be unique within the account.
func (client PoolClient) Get(resourceGroupName string, accountName string, poolName string) (result Pool, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "batch.PoolClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, accountName, poolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PoolClient) GetPreparer(resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PoolClient) GetResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBatchAccount lists all of the pools in the specified account.
//
// resourceGroupName is the name of the resource group that contains the Batch account. accountName is the name of the
// Batch account. maxresults is the maximum number of items to return in the response. selectParameter is comma
// separated list of properties that should be returned. e.g. "properties/provisioningState". Only top level properties
// under properties/ are valid for selection. filter is oData filter expression. Valid properties for filtering are:
//
// name
// properties/allocationState
// properties/allocationStateTransitionTime
// properties/creationTime
// properties/provisioningState
// properties/provisioningStateTransitionTime
// properties/lastModified
// properties/vmSize
// properties/interNodeCommunication
// properties/scaleSettings/autoScale
// properties/scaleSettings/fixedScale
func (client PoolClient) ListByBatchAccount(resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (result ListPoolsResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "batch.PoolClient", "ListByBatchAccount")
	}

	req, err := client.ListByBatchAccountPreparer(resourceGroupName, accountName, maxresults, selectParameter, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "ListByBatchAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "ListByBatchAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "ListByBatchAccount", resp, "Failure responding to request")
	}

	return
}

// ListByBatchAccountPreparer prepares the ListByBatchAccount request.
func (client PoolClient) ListByBatchAccountPreparer(resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxresults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxresults)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByBatchAccountSender sends the ListByBatchAccount request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) ListByBatchAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByBatchAccountResponder handles the response to the ListByBatchAccount request. The method always
// closes the http.Response Body.
func (client PoolClient) ListByBatchAccountResponder(resp *http.Response) (result ListPoolsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBatchAccountNextResults retrieves the next set of results, if any.
func (client PoolClient) ListByBatchAccountNextResults(lastResults ListPoolsResult) (result ListPoolsResult, err error) {
	req, err := lastResults.ListPoolsResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.PoolClient", "ListByBatchAccount", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.PoolClient", "ListByBatchAccount", resp, "Failure sending next results request")
	}

	result, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "ListByBatchAccount", resp, "Failure responding to next results request")
	}

	return
}

// ListByBatchAccountComplete gets all elements from the list without paging.
func (client PoolClient) ListByBatchAccountComplete(resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string, cancel <-chan struct{}) (<-chan Pool, <-chan error) {
	resultChan := make(chan Pool)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByBatchAccount(resourceGroupName, accountName, maxresults, selectParameter, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByBatchAccountNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// StopResize this does not restore the pool to its previous state before the resize operation: it only stops any
// further changes being made, and the pool maintains its current state. After stopping, the pool stabilizes at the
// number of nodes it was at when the stop operation was done. During the stop operation, the pool allocation state
// changes first to stopping and then to steady. A resize operation need not be an explicit resize pool request; this
// API can also be used to halt the initial sizing of the pool when it is created.
//
// resourceGroupName is the name of the resource group that contains the Batch account. accountName is the name of the
// Batch account. poolName is the pool name. This must be unique within the account.
func (client PoolClient) StopResize(resourceGroupName string, accountName string, poolName string) (result Pool, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "batch.PoolClient", "StopResize")
	}

	req, err := client.StopResizePreparer(resourceGroupName, accountName, poolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "StopResize", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopResizeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "StopResize", resp, "Failure sending request")
		return
	}

	result, err = client.StopResizeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "StopResize", resp, "Failure responding to request")
	}

	return
}

// StopResizePreparer prepares the StopResize request.
func (client PoolClient) StopResizePreparer(resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}/stopResize", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// StopResizeSender sends the StopResize request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) StopResizeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// StopResizeResponder handles the response to the StopResize request. The method always
// closes the http.Response Body.
func (client PoolClient) StopResizeResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the properties of an existing pool.
//
// resourceGroupName is the name of the resource group that contains the Batch account. accountName is the name of the
// Batch account. poolName is the pool name. This must be unique within the account. parameters is pool properties that
// should be updated. Properties that are supplied will be updated, any property not supplied will be unchanged.
// ifMatch is the entity state (ETag) version of the pool to update. This value can be omitted or set to "*" to apply
// the operation unconditionally.
func (client PoolClient) Update(resourceGroupName string, accountName string, poolName string, parameters Pool, ifMatch string) (result Pool, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "batch.PoolClient", "Update")
	}

	req, err := client.UpdatePreparer(resourceGroupName, accountName, poolName, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PoolClient) UpdatePreparer(resourceGroupName string, accountName string, poolName string, parameters Pool, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PoolClient) UpdateResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

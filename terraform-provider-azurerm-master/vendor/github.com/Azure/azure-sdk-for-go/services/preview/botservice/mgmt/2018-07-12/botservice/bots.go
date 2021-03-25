package botservice

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// BotsClient is the azure Bot Service is a platform for creating smart conversational agents.
type BotsClient struct {
	BaseClient
}

// NewBotsClient creates an instance of the BotsClient client.
func NewBotsClient(subscriptionID string) BotsClient {
	return NewBotsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBotsClientWithBaseURI creates an instance of the BotsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBotsClientWithBaseURI(baseURI string, subscriptionID string) BotsClient {
	return BotsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a Bot Service. Bot Service is a resource group wide resource type.
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
// parameters - the parameters to provide for the created bot.
func (client BotsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, parameters Bot) (result Bot, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Properties.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.Properties.Endpoint", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.Properties.MsaAppID", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("botservice.BotsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, resourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client BotsClient) CreatePreparer(ctx context.Context, resourceGroupName string, resourceName string, parameters Bot) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client BotsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client BotsClient) CreateResponder(resp *http.Response) (result Bot, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Bot Service from the resource group.
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
func (client BotsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BotsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BotsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BotsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a BotService specified by the parameters.
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
func (client BotsClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (result Bot, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client BotsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BotsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BotsClient) GetResponder(resp *http.Response) (result Bot, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCheckNameAvailability check whether a bot name is available.
// Parameters:
// parameters - the request body parameters to provide for the check name availability request
func (client BotsClient) GetCheckNameAvailability(ctx context.Context, parameters CheckNameAvailabilityRequestBody) (result CheckNameAvailabilityResponseBody, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotsClient.GetCheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetCheckNameAvailabilityPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "GetCheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "GetCheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.GetCheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "GetCheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// GetCheckNameAvailabilityPreparer prepares the GetCheckNameAvailability request.
func (client BotsClient) GetCheckNameAvailabilityPreparer(ctx context.Context, parameters CheckNameAvailabilityRequestBody) (*http.Request, error) {
	const APIVersion = "2018-07-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.BotService/checkNameAvailability"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCheckNameAvailabilitySender sends the GetCheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BotsClient) GetCheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetCheckNameAvailabilityResponder handles the response to the GetCheckNameAvailability request. The method always
// closes the http.Response Body.
func (client BotsClient) GetCheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityResponseBody, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns all the resources of a particular type belonging to a subscription.
func (client BotsClient) List(ctx context.Context) (result BotResponseListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotsClient.List")
		defer func() {
			sc := -1
			if result.brl.Response.Response != nil {
				sc = result.brl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.brl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "List", resp, "Failure sending request")
		return
	}

	result.brl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.brl.hasNextLink() && result.brl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client BotsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.BotService/botServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BotsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BotsClient) ListResponder(resp *http.Response) (result BotResponseList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client BotsClient) listNextResults(ctx context.Context, lastResults BotResponseList) (result BotResponseList, err error) {
	req, err := lastResults.botResponseListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "botservice.BotsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "botservice.BotsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BotsClient) ListComplete(ctx context.Context) (result BotResponseListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup returns all the resources of a particular type belonging to a resource group
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
func (client BotsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result BotResponseListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.brl.Response.Response != nil {
				sc = result.brl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotsClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.brl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.brl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.brl.hasNextLink() && result.brl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client BotsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client BotsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client BotsClient) ListByResourceGroupResponder(resp *http.Response) (result BotResponseList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client BotsClient) listByResourceGroupNextResults(ctx context.Context, lastResults BotResponseList) (result BotResponseList, err error) {
	req, err := lastResults.botResponseListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "botservice.BotsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "botservice.BotsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client BotsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result BotResponseListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update updates a Bot Service
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
// parameters - the parameters to provide for the created bot.
func (client BotsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, parameters Bot) (result Bot, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client BotsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, parameters Bot) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-12"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client BotsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BotsClient) UpdateResponder(resp *http.Response) (result Bot, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
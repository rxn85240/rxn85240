package appplatform

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

// ConfigServersClient is the REST API for Azure Spring Cloud
type ConfigServersClient struct {
	BaseClient
}

// NewConfigServersClient creates an instance of the ConfigServersClient client.
func NewConfigServersClient(subscriptionID string) ConfigServersClient {
	return NewConfigServersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigServersClientWithBaseURI creates an instance of the ConfigServersClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewConfigServersClientWithBaseURI(baseURI string, subscriptionID string) ConfigServersClient {
	return ConfigServersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the config server and its properties.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
func (client ConfigServersClient) Get(ctx context.Context, resourceGroupName string, serviceName string) (result ConfigServerResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigServersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ConfigServersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.ConfigServersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ConfigServersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigServersClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigServersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigServersClient) GetResponder(resp *http.Response) (result ConfigServerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdatePatch update the config server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// configServerResource - parameters for the update operation
func (client ConfigServersClient) UpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource) (result ConfigServersUpdatePatchFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigServersClient.UpdatePatch")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePatchPreparer(ctx, resourceGroupName, serviceName, configServerResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ConfigServersClient", "UpdatePatch", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdatePatchSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ConfigServersClient", "UpdatePatch", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePatchPreparer prepares the UpdatePatch request.
func (client ConfigServersClient) UpdatePatchPreparer(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default", pathParameters),
		autorest.WithJSON(configServerResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePatchSender sends the UpdatePatch request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigServersClient) UpdatePatchSender(req *http.Request) (future ConfigServersUpdatePatchFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ConfigServersClient) (csr ConfigServerResource, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "appplatform.ConfigServersUpdatePatchFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("appplatform.ConfigServersUpdatePatchFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		csr.Response.Response, err = future.GetResult(sender)
		if csr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "appplatform.ConfigServersUpdatePatchFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && csr.Response.Response.StatusCode != http.StatusNoContent {
			csr, err = client.UpdatePatchResponder(csr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.ConfigServersUpdatePatchFuture", "Result", csr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdatePatchResponder handles the response to the UpdatePatch request. The method always
// closes the http.Response Body.
func (client ConfigServersClient) UpdatePatchResponder(resp *http.Response) (result ConfigServerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdatePut update the config server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// configServerResource - parameters for the update operation
func (client ConfigServersClient) UpdatePut(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource) (result ConfigServersUpdatePutFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigServersClient.UpdatePut")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configServerResource,
			Constraints: []validation.Constraint{{Target: "configServerResource.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "configServerResource.Properties.ConfigServer", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "configServerResource.Properties.ConfigServer.GitProperty", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "configServerResource.Properties.ConfigServer.GitProperty.URI", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("appplatform.ConfigServersClient", "UpdatePut", err.Error())
	}

	req, err := client.UpdatePutPreparer(ctx, resourceGroupName, serviceName, configServerResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ConfigServersClient", "UpdatePut", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdatePutSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ConfigServersClient", "UpdatePut", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePutPreparer prepares the UpdatePut request.
func (client ConfigServersClient) UpdatePutPreparer(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default", pathParameters),
		autorest.WithJSON(configServerResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePutSender sends the UpdatePut request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigServersClient) UpdatePutSender(req *http.Request) (future ConfigServersUpdatePutFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ConfigServersClient) (csr ConfigServerResource, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "appplatform.ConfigServersUpdatePutFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("appplatform.ConfigServersUpdatePutFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		csr.Response.Response, err = future.GetResult(sender)
		if csr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "appplatform.ConfigServersUpdatePutFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && csr.Response.Response.StatusCode != http.StatusNoContent {
			csr, err = client.UpdatePutResponder(csr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.ConfigServersUpdatePutFuture", "Result", csr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdatePutResponder handles the response to the UpdatePut request. The method always
// closes the http.Response Body.
func (client ConfigServersClient) UpdatePutResponder(resp *http.Response) (result ConfigServerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Validate check if the config server settings are valid.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// configServerSettings - config server settings to be validated
func (client ConfigServersClient) Validate(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings ConfigServerSettings) (result ConfigServersValidateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigServersClient.Validate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configServerSettings,
			Constraints: []validation.Constraint{{Target: "configServerSettings.GitProperty", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "configServerSettings.GitProperty.URI", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("appplatform.ConfigServersClient", "Validate", err.Error())
	}

	req, err := client.ValidatePreparer(ctx, resourceGroupName, serviceName, configServerSettings)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ConfigServersClient", "Validate", nil, "Failure preparing request")
		return
	}

	result, err = client.ValidateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ConfigServersClient", "Validate", nil, "Failure sending request")
		return
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client ConfigServersClient) ValidatePreparer(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings ConfigServerSettings) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/validate", pathParameters),
		autorest.WithJSON(configServerSettings),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigServersClient) ValidateSender(req *http.Request) (future ConfigServersValidateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ConfigServersClient) (cssvr ConfigServerSettingsValidateResult, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "appplatform.ConfigServersValidateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("appplatform.ConfigServersValidateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		cssvr.Response.Response, err = future.GetResult(sender)
		if cssvr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "appplatform.ConfigServersValidateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && cssvr.Response.Response.StatusCode != http.StatusNoContent {
			cssvr, err = client.ValidateResponder(cssvr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.ConfigServersValidateFuture", "Result", cssvr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client ConfigServersClient) ValidateResponder(resp *http.Response) (result ConfigServerSettingsValidateResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

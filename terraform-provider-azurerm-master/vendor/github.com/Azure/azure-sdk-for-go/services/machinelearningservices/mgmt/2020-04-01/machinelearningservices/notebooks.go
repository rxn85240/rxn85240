package machinelearningservices

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// NotebooksClient is the these APIs allow end users to operate on Azure Machine Learning Workspace resources.
type NotebooksClient struct {
	BaseClient
}

// NewNotebooksClient creates an instance of the NotebooksClient client.
func NewNotebooksClient(subscriptionID string) NotebooksClient {
	return NewNotebooksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNotebooksClientWithBaseURI creates an instance of the NotebooksClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewNotebooksClientWithBaseURI(baseURI string, subscriptionID string) NotebooksClient {
	return NotebooksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Prepare sends the prepare request.
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
func (client NotebooksClient) Prepare(ctx context.Context, resourceGroupName string, workspaceName string) (result NotebooksPrepareFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebooksClient.Prepare")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PreparePreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.NotebooksClient", "Prepare", nil, "Failure preparing request")
		return
	}

	result, err = client.PrepareSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.NotebooksClient", "Prepare", nil, "Failure sending request")
		return
	}

	return
}

// PreparePreparer prepares the Prepare request.
func (client NotebooksClient) PreparePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/prepareNotebook", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PrepareSender sends the Prepare request. The method will close the
// http.Response Body if it receives an error.
func (client NotebooksClient) PrepareSender(req *http.Request) (future NotebooksPrepareFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client NotebooksClient) (nri NotebookResourceInfo, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "machinelearningservices.NotebooksPrepareFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("machinelearningservices.NotebooksPrepareFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		nri.Response.Response, err = future.GetResult(sender)
		if nri.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "machinelearningservices.NotebooksPrepareFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && nri.Response.Response.StatusCode != http.StatusNoContent {
			nri, err = client.PrepareResponder(nri.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "machinelearningservices.NotebooksPrepareFuture", "Result", nri.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// PrepareResponder handles the response to the Prepare request. The method always
// closes the http.Response Body.
func (client NotebooksClient) PrepareResponder(resp *http.Response) (result NotebookResourceInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

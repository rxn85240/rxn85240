package account

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

// DataLakeStoreAccountState enumerates the values for data lake store account state.
type DataLakeStoreAccountState string

const (
	// Active ...
	Active DataLakeStoreAccountState = "Active"
	// Suspended ...
	Suspended DataLakeStoreAccountState = "Suspended"
)

// PossibleDataLakeStoreAccountStateValues returns an array of possible values for the DataLakeStoreAccountState const type.
func PossibleDataLakeStoreAccountStateValues() []DataLakeStoreAccountState {
	return []DataLakeStoreAccountState{Active, Suspended}
}

// DataLakeStoreAccountStatus enumerates the values for data lake store account status.
type DataLakeStoreAccountStatus string

const (
	// Canceled ...
	Canceled DataLakeStoreAccountStatus = "Canceled"
	// Creating ...
	Creating DataLakeStoreAccountStatus = "Creating"
	// Deleted ...
	Deleted DataLakeStoreAccountStatus = "Deleted"
	// Deleting ...
	Deleting DataLakeStoreAccountStatus = "Deleting"
	// Failed ...
	Failed DataLakeStoreAccountStatus = "Failed"
	// Patching ...
	Patching DataLakeStoreAccountStatus = "Patching"
	// Resuming ...
	Resuming DataLakeStoreAccountStatus = "Resuming"
	// Running ...
	Running DataLakeStoreAccountStatus = "Running"
	// Succeeded ...
	Succeeded DataLakeStoreAccountStatus = "Succeeded"
	// Suspending ...
	Suspending DataLakeStoreAccountStatus = "Suspending"
	// Undeleting ...
	Undeleting DataLakeStoreAccountStatus = "Undeleting"
)

// PossibleDataLakeStoreAccountStatusValues returns an array of possible values for the DataLakeStoreAccountStatus const type.
func PossibleDataLakeStoreAccountStatusValues() []DataLakeStoreAccountStatus {
	return []DataLakeStoreAccountStatus{Canceled, Creating, Deleted, Deleting, Failed, Patching, Resuming, Running, Succeeded, Suspending, Undeleting}
}

// EncryptionConfigType enumerates the values for encryption config type.
type EncryptionConfigType string

const (
	// ServiceManaged ...
	ServiceManaged EncryptionConfigType = "ServiceManaged"
	// UserManaged ...
	UserManaged EncryptionConfigType = "UserManaged"
)

// PossibleEncryptionConfigTypeValues returns an array of possible values for the EncryptionConfigType const type.
func PossibleEncryptionConfigTypeValues() []EncryptionConfigType {
	return []EncryptionConfigType{ServiceManaged, UserManaged}
}

// EncryptionProvisioningState enumerates the values for encryption provisioning state.
type EncryptionProvisioningState string

const (
	// EncryptionProvisioningStateCreating ...
	EncryptionProvisioningStateCreating EncryptionProvisioningState = "Creating"
	// EncryptionProvisioningStateSucceeded ...
	EncryptionProvisioningStateSucceeded EncryptionProvisioningState = "Succeeded"
)

// PossibleEncryptionProvisioningStateValues returns an array of possible values for the EncryptionProvisioningState const type.
func PossibleEncryptionProvisioningStateValues() []EncryptionProvisioningState {
	return []EncryptionProvisioningState{EncryptionProvisioningStateCreating, EncryptionProvisioningStateSucceeded}
}

// EncryptionState enumerates the values for encryption state.
type EncryptionState string

const (
	// Disabled ...
	Disabled EncryptionState = "Disabled"
	// Enabled ...
	Enabled EncryptionState = "Enabled"
)

// PossibleEncryptionStateValues returns an array of possible values for the EncryptionState const type.
func PossibleEncryptionStateValues() []EncryptionState {
	return []EncryptionState{Disabled, Enabled}
}

// FirewallAllowAzureIpsState enumerates the values for firewall allow azure ips state.
type FirewallAllowAzureIpsState string

const (
	// FirewallAllowAzureIpsStateDisabled ...
	FirewallAllowAzureIpsStateDisabled FirewallAllowAzureIpsState = "Disabled"
	// FirewallAllowAzureIpsStateEnabled ...
	FirewallAllowAzureIpsStateEnabled FirewallAllowAzureIpsState = "Enabled"
)

// PossibleFirewallAllowAzureIpsStateValues returns an array of possible values for the FirewallAllowAzureIpsState const type.
func PossibleFirewallAllowAzureIpsStateValues() []FirewallAllowAzureIpsState {
	return []FirewallAllowAzureIpsState{FirewallAllowAzureIpsStateDisabled, FirewallAllowAzureIpsStateEnabled}
}

// FirewallState enumerates the values for firewall state.
type FirewallState string

const (
	// FirewallStateDisabled ...
	FirewallStateDisabled FirewallState = "Disabled"
	// FirewallStateEnabled ...
	FirewallStateEnabled FirewallState = "Enabled"
)

// PossibleFirewallStateValues returns an array of possible values for the FirewallState const type.
func PossibleFirewallStateValues() []FirewallState {
	return []FirewallState{FirewallStateDisabled, FirewallStateEnabled}
}

// OperationOrigin enumerates the values for operation origin.
type OperationOrigin string

const (
	// System ...
	System OperationOrigin = "system"
	// User ...
	User OperationOrigin = "user"
	// Usersystem ...
	Usersystem OperationOrigin = "user,system"
)

// PossibleOperationOriginValues returns an array of possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{System, User, Usersystem}
}

// SubscriptionState enumerates the values for subscription state.
type SubscriptionState string

const (
	// SubscriptionStateDeleted ...
	SubscriptionStateDeleted SubscriptionState = "Deleted"
	// SubscriptionStateRegistered ...
	SubscriptionStateRegistered SubscriptionState = "Registered"
	// SubscriptionStateSuspended ...
	SubscriptionStateSuspended SubscriptionState = "Suspended"
	// SubscriptionStateUnregistered ...
	SubscriptionStateUnregistered SubscriptionState = "Unregistered"
	// SubscriptionStateWarned ...
	SubscriptionStateWarned SubscriptionState = "Warned"
)

// PossibleSubscriptionStateValues returns an array of possible values for the SubscriptionState const type.
func PossibleSubscriptionStateValues() []SubscriptionState {
	return []SubscriptionState{SubscriptionStateDeleted, SubscriptionStateRegistered, SubscriptionStateSuspended, SubscriptionStateUnregistered, SubscriptionStateWarned}
}

// TierType enumerates the values for tier type.
type TierType string

const (
	// Commitment100TB ...
	Commitment100TB TierType = "Commitment_100TB"
	// Commitment10TB ...
	Commitment10TB TierType = "Commitment_10TB"
	// Commitment1PB ...
	Commitment1PB TierType = "Commitment_1PB"
	// Commitment1TB ...
	Commitment1TB TierType = "Commitment_1TB"
	// Commitment500TB ...
	Commitment500TB TierType = "Commitment_500TB"
	// Commitment5PB ...
	Commitment5PB TierType = "Commitment_5PB"
	// Consumption ...
	Consumption TierType = "Consumption"
)

// PossibleTierTypeValues returns an array of possible values for the TierType const type.
func PossibleTierTypeValues() []TierType {
	return []TierType{Commitment100TB, Commitment10TB, Commitment1PB, Commitment1TB, Commitment500TB, Commitment5PB, Consumption}
}

// TrustedIDProviderState enumerates the values for trusted id provider state.
type TrustedIDProviderState string

const (
	// TrustedIDProviderStateDisabled ...
	TrustedIDProviderStateDisabled TrustedIDProviderState = "Disabled"
	// TrustedIDProviderStateEnabled ...
	TrustedIDProviderStateEnabled TrustedIDProviderState = "Enabled"
)

// PossibleTrustedIDProviderStateValues returns an array of possible values for the TrustedIDProviderState const type.
func PossibleTrustedIDProviderStateValues() []TrustedIDProviderState {
	return []TrustedIDProviderState{TrustedIDProviderStateDisabled, TrustedIDProviderStateEnabled}
}

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// Bytes ...
	Bytes UsageUnit = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UsageUnit = "BytesPerSecond"
	// Count ...
	Count UsageUnit = "Count"
	// CountsPerSecond ...
	CountsPerSecond UsageUnit = "CountsPerSecond"
	// Percent ...
	Percent UsageUnit = "Percent"
	// Seconds ...
	Seconds UsageUnit = "Seconds"
)

// PossibleUsageUnitValues returns an array of possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{Bytes, BytesPerSecond, Count, CountsPerSecond, Percent, Seconds}
}

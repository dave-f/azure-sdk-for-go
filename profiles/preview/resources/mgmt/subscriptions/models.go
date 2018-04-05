// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package subscriptions

import original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-06-01/subscriptions"

type SpendingLimit = original.SpendingLimit

const (
	CurrentPeriodOff SpendingLimit = original.CurrentPeriodOff
	Off              SpendingLimit = original.Off
	On               SpendingLimit = original.On
)

func PossibleSpendingLimitValues() []SpendingLimit {
	return original.PossibleSpendingLimitValues()
}

type State = original.State

const (
	Deleted  State = original.Deleted
	Disabled State = original.Disabled
	Enabled  State = original.Enabled
	PastDue  State = original.PastDue
	Warned   State = original.Warned
)

func PossibleStateValues() []State {
	return original.PossibleStateValues()
}

type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type Location = original.Location
type LocationListResult = original.LocationListResult
type Policies = original.Policies
type Subscription = original.Subscription
type TenantIDDescription = original.TenantIDDescription
type TenantListResult = original.TenantListResult
type TenantListResultIterator = original.TenantListResultIterator
type TenantListResultPage = original.TenantListResultPage

func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

type Client = original.Client

func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}

type TenantsClient = original.TenantsClient

func NewTenantsClient() TenantsClient {
	return original.NewTenantsClient()
}
func NewTenantsClientWithBaseURI(baseURI string) TenantsClient {
	return original.NewTenantsClientWithBaseURI(baseURI)
}

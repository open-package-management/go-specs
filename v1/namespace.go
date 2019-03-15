// Copyright © 2019 Open Package Management Authors
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

package v1

// Namespace is the go struct for an OPM namespace.
type Namespace struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Quotas struct {
		Projects     QuotaDescriptor `json:"projects"`
		Repositories QuotaDescriptor `json:"repositories"`
		Storage      QuotaDescriptor `json:"storage"`
	} `json:"quotas"`
	Labels Labels `json:"labels,omitempty"`
}

// NamespaceCreate is the go struct for creating OPM namespaces.
type NamespaceCreate struct {
	Labels Labels `json:"labels,omitempty"`
}

// NamespaceList is the go struct for a list of OPM namespaces.
type NamespaceList struct {
	Namespaces []Namespace `json:"namespaces"`
}

// QuotaDescriptor is a struct that contains data about resource usage and
// limits.
type QuotaDescriptor struct {
	Limit uint64 `json:"limit"`
	Used  uint64 `json:"used"`
}

// Labels is a struct that contains data about labels.
type Labels struct {
	Provider map[string]string `json:"provider,omitempty"`
	Consumer map[string]string `json:"consumer,omitempty"`
}

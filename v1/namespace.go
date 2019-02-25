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

// Namespace defines information about OPM namespaces.
type Namespace struct {
	// Name defines the name of the namespace.
	Name string `json:"name"`

	// Projects defines information about projects in the namespace.
	Projects struct {
		// Quota defines the total amount of projects allowed in the namespace.
		Quota uint64 `json:"quota"`

		// Count defines the total amount of projects currently in the
		// namespace.
		Count uint64 `json:"count"`
	} `json:"projects"`

	// Repositories defines information about repositories in the namespace.
	Repositories struct {
		// Quota defines the total amount of repositories allowed in the
		// namespace.
		Quota uint64 `json:"quota"`

		// Count defines the total amount of repositories currently in the
		// namespace.
		Count uint64 `json:"count"`
	} `json:"repositories"`

	// Storage defines storage information about namespace.
	Storage struct {
		// Quota defines the total storage allocated to the namespace.
		Quota uint64 `json:"quota,omitempty"`

		// Used defines the total storage used by the namespace in bytes.
		Used uint64 `json:"used"`
	} `json:"storage"`

	// Labels contains arbitrary metadata for the namespace.
	Labels map[string]string `json:"labels,omitempty"`
}

// Copyright The Karbour Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import "github.com/KusionStack/karbour/pkg/core/handler"

// Ensure that ClusterPayload implements the handler.Payload interface.
var _ handler.Payload = &ClusterPayload{}

// ClusterPayload represents the structure for cluster request data. It includes
// the name, display name, description and kubeconfig of a karbour-managed cluster
type ClusterPayload struct {
	ClusterName        string `json:"name"`        // ClusterName is the name of cluster to be created
	ClusterDisplayName string `json:"displayName"` // ClusterDisplayName is the display name of cluster to be created
	ClusterDescription string `json:"description"` // ClusterDescription is the description of cluster to be created
	ClusterKubeConfig  string `json:"kubeconfig"`  // ClusterKubeConfig is the kubeconfig of cluster to be created
}

type UploadData struct {
	FileName string `json:"fileName"`
	FileSize int    `json:"fileSize"`
	Content  string `json:"content"`
}

type ValidatePayload struct {
	KubeConfig string `json:"kubeConfig"`
}
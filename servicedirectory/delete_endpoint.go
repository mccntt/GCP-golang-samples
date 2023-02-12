// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package servicedirectory

// [START servicedirectory_delete_endpoint]
import (
	"context"
	"fmt"

	servicedirectory "cloud.google.com/go/servicedirectory/apiv1"
	sdpb "cloud.google.com/go/servicedirectory/apiv1/servicedirectorypb"
)

func deleteEndpoint(projectID string) error {
	// projectID := "my-project"
	location := "us-east4"
	namespaceID := "golang-test-namespace"
	serviceID := "golang-test-service"
	endpointID := "golang-test-endpoint"

	ctx := context.Background()
	// Create a registration client.
	client, err := servicedirectory.NewRegistrationClient(ctx)
	if err != nil {
		return fmt.Errorf("ServiceDirectory.NewRegistrationClient: %v", err)
	}

	defer client.Close()
	// Delete an Endpoint
	req := &sdpb.DeleteEndpointRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s/endpoints/%s", projectID, location, namespaceID, serviceID, endpointID),
	}
	if err := client.DeleteEndpoint(ctx, req); err != nil {
		return fmt.Errorf("DeleteEndpoint: %v", err)
	}
	return nil
}

// [END servicedirectory_delete_endpoint]

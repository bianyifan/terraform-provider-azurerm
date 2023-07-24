package privateendpointconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkServiceConnectionState struct {
	Description *string           `json:"description,omitempty"`
	Status      *ConnectionStatus `json:"status,omitempty"`
}

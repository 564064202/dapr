// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package pubsub

import (
	jsoniter "github.com/json-iterator/go"
)

const (
	// DefaultCloudEventType is the default event type for an Dapr published event
	DefaultCloudEventType = "com.dapr.event.sent"
	// CloudEventsSpecVersion is the specversion used by Dapr for the cloud events implementation
	CloudEventsSpecVersion = "0.3"
	//ContentType is the Cloud Events HTTP content type
	ContentType = "application/cloudevents+json"
)

// CloudEventsEnvelope describes the Dapr implementation of the Cloud Events spec
// Spec details: https://github.com/cloudevents/spec/blob/master/spec.md
type CloudEventsEnvelope struct {
	ID              string      `json:"id"`
	Source          string      `json:"source"`
	Type            string      `json:"type"`
	SpecVersion     string      `json:"specversion"`
	DataContentType string      `json:"datacontenttype"`
	Data            interface{} `json:"data"`
}

// NewCloudEventsEnvelope returns a new CloudEventsEnvelope
func NewCloudEventsEnvelope(id, source, eventType string, data []byte) *CloudEventsEnvelope {
	if eventType == "" {
		eventType = DefaultCloudEventType
	}
	contentType := ""

	var i interface{}
	err := jsoniter.Unmarshal(data, &i)
	if err != nil {
		i = string(data)
		contentType = "text/plain"
	} else {
		contentType = "application/json"
	}

	return &CloudEventsEnvelope{
		ID:              id,
		Source:          source,
		Type:            eventType,
		Data:            i,
		SpecVersion:     CloudEventsSpecVersion,
		DataContentType: contentType,
	}
}

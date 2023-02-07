// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// The custom parameters to be used when the target is an Batch job.
type BatchParameters struct {
	JobDefinition *string `json:"jobDefinition,omitempty"`
	JobName       *string `json:"jobName,omitempty"`
}

// A JSON string which you can use to limit the event bus permissions you are
// granting to only accounts that fulfill the condition. Currently, the only
// supported condition is membership in a certain Amazon Web Services organization.
// The string must contain Type, Key, and Value fields. The Value field specifies
// the ID of the Amazon Web Services organization. Following is an example value
// for Condition:
//
// '{"Type" : "StringEquals", "Key": "aws:PrincipalOrgID", "Value": "o-1234567890"}'
type Condition struct {
	Key   *string `json:"key,omitempty"`
	Type  *string `json:"type_,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Additional parameter included in the body. You can include up to 100 additional
// body parameters per request. An event payload cannot exceed 64 KB.
type ConnectionBodyParameter struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// The custom parameters to be used when the target is an Amazon ECS task.
type EcsParameters struct {
	Group           *string `json:"group,omitempty"`
	PlatformVersion *string `json:"platformVersion,omitempty"`
	Tags            []*Tag  `json:"tags,omitempty"`
}

// An event bus receives events from a source and routes them to rules associated
// with that event bus. Your account's default event bus receives events from
// Amazon Web Services services. A custom event bus can receive events from
// your custom applications and services. A partner event bus receives events
// from an event source created by an SaaS partner. These events come from the
// partners services or applications.
type EventBus_SDK struct {
	ARN    *string `json:"arn,omitempty"`
	Name   *string `json:"name,omitempty"`
	Policy *string `json:"policy,omitempty"`
}

// A partner event source is created by an SaaS partner. If a customer creates
// a partner event bus that matches this event source, that Amazon Web Services
// account can receive events from the partner's applications or services.
type EventSource struct {
	ARN       *string `json:"arn,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// A partner event source is created by an SaaS partner. If a customer creates
// a partner event bus that matches this event source, that Amazon Web Services
// account can receive events from the partner's applications or services.
type PartnerEventSource struct {
	ARN  *string `json:"arn,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Represents an event to be submitted.
type PutEventsRequestEntry struct {
	Detail     *string `json:"detail,omitempty"`
	DetailType *string `json:"detailType,omitempty"`
	Source     *string `json:"source,omitempty"`
}

// The details about an event generated by an SaaS partner.
type PutPartnerEventsRequestEntry struct {
	Detail     *string `json:"detail,omitempty"`
	DetailType *string `json:"detailType,omitempty"`
	Source     *string `json:"source,omitempty"`
}

// Contains information about a rule in Amazon EventBridge.
type Rule struct {
	EventBusName *string `json:"eventBusName,omitempty"`
}

// A key-value pair associated with an Amazon Web Services resource. In EventBridge,
// rules and event buses support tagging.
type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}
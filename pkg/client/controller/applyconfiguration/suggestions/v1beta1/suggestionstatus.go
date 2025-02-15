/*
Copyright 2022 The Kubeflow Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubeflow/katib/pkg/apis/controller/common/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SuggestionStatusApplyConfiguration represents a declarative configuration of the SuggestionStatus type for use
// with apply.
type SuggestionStatusApplyConfiguration struct {
	AlgorithmSettings []v1beta1.AlgorithmSetting              `json:"algorithmSettings,omitempty"`
	SuggestionCount   *int32                                  `json:"suggestionCount,omitempty"`
	Suggestions       []TrialAssignmentApplyConfiguration     `json:"suggestions,omitempty"`
	StartTime         *v1.Time                                `json:"startTime,omitempty"`
	CompletionTime    *v1.Time                                `json:"completionTime,omitempty"`
	LastReconcileTime *v1.Time                                `json:"lastReconcileTime,omitempty"`
	Conditions        []SuggestionConditionApplyConfiguration `json:"conditions,omitempty"`
}

// SuggestionStatusApplyConfiguration constructs a declarative configuration of the SuggestionStatus type for use with
// apply.
func SuggestionStatus() *SuggestionStatusApplyConfiguration {
	return &SuggestionStatusApplyConfiguration{}
}

// WithAlgorithmSettings adds the given value to the AlgorithmSettings field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AlgorithmSettings field.
func (b *SuggestionStatusApplyConfiguration) WithAlgorithmSettings(values ...v1beta1.AlgorithmSetting) *SuggestionStatusApplyConfiguration {
	for i := range values {
		b.AlgorithmSettings = append(b.AlgorithmSettings, values[i])
	}
	return b
}

// WithSuggestionCount sets the SuggestionCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SuggestionCount field is set to the value of the last call.
func (b *SuggestionStatusApplyConfiguration) WithSuggestionCount(value int32) *SuggestionStatusApplyConfiguration {
	b.SuggestionCount = &value
	return b
}

// WithSuggestions adds the given value to the Suggestions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Suggestions field.
func (b *SuggestionStatusApplyConfiguration) WithSuggestions(values ...*TrialAssignmentApplyConfiguration) *SuggestionStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSuggestions")
		}
		b.Suggestions = append(b.Suggestions, *values[i])
	}
	return b
}

// WithStartTime sets the StartTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartTime field is set to the value of the last call.
func (b *SuggestionStatusApplyConfiguration) WithStartTime(value v1.Time) *SuggestionStatusApplyConfiguration {
	b.StartTime = &value
	return b
}

// WithCompletionTime sets the CompletionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CompletionTime field is set to the value of the last call.
func (b *SuggestionStatusApplyConfiguration) WithCompletionTime(value v1.Time) *SuggestionStatusApplyConfiguration {
	b.CompletionTime = &value
	return b
}

// WithLastReconcileTime sets the LastReconcileTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastReconcileTime field is set to the value of the last call.
func (b *SuggestionStatusApplyConfiguration) WithLastReconcileTime(value v1.Time) *SuggestionStatusApplyConfiguration {
	b.LastReconcileTime = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *SuggestionStatusApplyConfiguration) WithConditions(values ...*SuggestionConditionApplyConfiguration) *SuggestionStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import "encoding/json"

type ResourceOperation struct {
	Index     string            `json:"index" yaml:"index,omitempty"`
	Operation string            `json:"operation" yaml:"operation,omitempty"`
	Object    string            `json:"object" yaml:"object,omitempty"`
	Parameter string            `json:"parameter" yaml:"parameter,omitempty"`
	Resource  string            `json:"resource" yaml:"resource,omitempty"`
	Secondary []string          `json:"secondary" yaml:"secondary,omitempty"`
	Mappings  map[string]string `json:"mappings" yaml:"mappings,omitempty"`
}

// Custom marshaling to make empty strings null
func (ro ResourceOperation) MarshalJSON() ([]byte, error) {
	test := struct {
		Index     *string           `json:"index,omitempty"`
		Operation *string           `json:"operation,omitempty"`
		Object    *string           `json:"object,omitempty"`
		Parameter *string           `json:"parameter,omitempty"`
		Resource  *string           `json:"resource,omitempty"`
		Secondary []string          `json:"secondary,omitempty"`
		Mappings  map[string]string `json:"mappings,omitempty"`
	}{
		Secondary: ro.Secondary,
		Mappings:  ro.Mappings,
	}

	// Empty strings are null
	if ro.Index != "" {
		test.Index = &ro.Index
	}
	if ro.Operation != "" {
		test.Operation = &ro.Operation
	}
	if ro.Object != "" {
		test.Object = &ro.Object
	}
	if ro.Parameter != "" {
		test.Parameter = &ro.Parameter
	}
	if ro.Resource != "" {
		test.Resource = &ro.Resource
	}

	return json.Marshal(test)
}

/*
 * To String function for ResourceOperation
 */
func (ro ResourceOperation) String() string {
	out, err := json.Marshal(ro)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
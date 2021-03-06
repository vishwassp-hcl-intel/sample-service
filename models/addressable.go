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

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

/*
 * This file is the model for addressable in EdgeX
 * Addressable holds information about a specific address
 *
 * Addressable struct
 */
type Addressable struct {
	BaseObject
	Id         string `json:"id"`
	Name       string `json:"name"`
	Protocol   string `json:"protocol"`    // Protocol for the address (HTTP/TCP)
	HTTPMethod string `json:"method"`      // Method for connecting (i.e. POST)
	Address    string `json:"address"`     // Address of the addressable
	Port       int    `json:"port,Number"` // Port for the address
	Path       string `json:"path"`        // Path for callbacks
	Publisher  string `json:"publisher"`   // For message bus protocols
	User       string `json:"user"`        // User id for authentication
	Password   string `json:"password"`    // Password of the user for authentication for the addressable
	Topic      string `json:"topic"`       // Topic for message bus addressables
}

// Custom marshaling for JSON
// Create the URL and Base URL
// Treat the strings as pointers so they can be null in JSON
func (a Addressable) MarshalJSON() ([]byte, error) {
	aux := struct {
		BaseObject
		Id         *string `json:"id,omitempty"`
		Name       *string `json:"name,omitempty"`
		Protocol   *string `json:"protocol,omitempty"`    // Protocol for the address (HTTP/TCP)
		HTTPMethod *string `json:"method,omitempty"`      // Method for connecting (i.e. POST)
		Address    *string `json:"address,omitempty"`     // Address of the addressable
		Port       int     `json:"port,Number,omitempty"` // Port for the address
		Path       *string `json:"path,omitempty"`        // Path for callbacks
		Publisher  *string `json:"publisher,omitempty"`   // For message bus protocols
		User       *string `json:"user,omitempty"`        // User id for authentication
		Password   *string `json:"password,omitempty"`    // Password of the user for authentication for the addressable
		Topic      *string `json:"topic,omitempty"`       // Topic for message bus addressables
		BaseURL    *string `json:"baseURL,omitempty"`
		URL        *string `json:"url,omitempty"`
	}{
		BaseObject: a.BaseObject,
		Port:       a.Port,
	}

	if a.Id != "" {
		aux.Id = &a.Id
	}

	// Only initialize the non-empty strings (empty are null)
	if a.Name != "" {
		aux.Name = &a.Name
	}
	if a.Protocol != "" {
		aux.Protocol = &a.Protocol
	}
	if a.HTTPMethod != "" {
		aux.HTTPMethod = &a.HTTPMethod
	}
	if a.Address != "" {
		aux.Address = &a.Address
	}
	if a.Path != "" {
		aux.Path = &a.Path
	}
	if a.Publisher != "" {
		aux.Publisher = &a.Publisher
	}
	if a.User != "" {
		aux.User = &a.User
	}
	if a.Password != "" {
		aux.Password = &a.Password
	}
	if a.Topic != "" {
		aux.Topic = &a.Topic
	}

	// Get the base URL
	if a.Protocol != "" && a.Address != "" {
		var baseUrlBuffer bytes.Buffer
		_, err := baseUrlBuffer.WriteString(a.Protocol)
		if err != nil {
			return []byte{}, err
		}
		baseUrlBuffer.WriteString("://")
		_, err = baseUrlBuffer.WriteString(a.Address)
		if err != nil {
			return []byte{}, err
		}
		baseUrlBuffer.WriteString(":")
		_, err = baseUrlBuffer.WriteString(strconv.Itoa(a.Port))
		if err != nil {
			return []byte{}, err
		}
		s := baseUrlBuffer.String()
		aux.BaseURL = &s
	}

	// Get the URL
	if aux.BaseURL != nil {
		var urlBuffer bytes.Buffer
		_, err := urlBuffer.WriteString(*aux.BaseURL)
		if err != nil {
			return []byte{}, err
		}
		if a.Publisher == "" && a.Topic != "" {
			_, err = urlBuffer.WriteString(a.Topic)
			if err != nil {
				return []byte{}, err
			}
			urlBuffer.WriteString("/")
		}
		_, err = urlBuffer.WriteString(a.Path)
		if err != nil {
			return []byte{}, err
		}
		s := urlBuffer.String()
		aux.URL = &s
	}

	return json.Marshal(aux)
}

/*
 * String() function for formatting
 */
func (a Addressable) String() string {
	out, err := json.Marshal(a)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

func (a Addressable) GetBaseURL() string {
	protocol := strings.ToLower(a.Protocol)
	address := a.Address
	port := strconv.Itoa(a.Port)
	baseUrl := protocol + "://" + address + ":" + port
	return baseUrl
}

// Get the callback url for the addressable if all relevant tokens have values.
// If any token is missing, string will be empty
func (a Addressable) GetCallbackURL() string {
	url := ""
	if len(a.Protocol) > 0 && len(a.Address) > 0 && a.Port > 0 && len(a.Path) > 0 {
		url = a.GetBaseURL() + a.Path
	}

	return url
}

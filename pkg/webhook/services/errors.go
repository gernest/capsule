/*
Copyright 2020 Clastix Labs.

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

package services

import (
	"fmt"
	"strings"

	"github.com/clastix/capsule/api/v1alpha1"
)

type externalServiceIPForbidden struct {
	cidr []string
}

func NewExternalServiceIPForbidden(allowedIps []v1alpha1.AllowedIp) error {
	var cidr []string
	for _, i := range allowedIps {
		cidr = append(cidr, string(i))
	}
	return &externalServiceIPForbidden{
		cidr: cidr,
	}
}

func (e externalServiceIPForbidden) Error() string {

	return fmt.Sprintf("The selected external IPs for the current Service are violating the following enforced CIDRs: %s", strings.Join(e.cidr, ", "))
}

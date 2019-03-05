package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

import "github.com/pkg/errors"

// Attribute is a Attribute Object Structure
type Attribute struct {
	Tag `kmip:"ATTRIBUTE"`

	Name  string      `kmip:"ATTRIBUTE_NAME"`
	Index int32       `kmip:"ATTRIBUTE_INDEX"`
	Value interface{} `kmip:"ATTRIBUTE_VALUE"`
}

// BuildFieldValue builds dynamic Value field
func (a *Attribute) BuildFieldValue(name string) (v interface{}, err error) {
	switch a.Name {
	case "Cryptographic Algorithm":
		v = Enum(0)
	case "Cryptographic Length", "Cryptographic Usage Mask":
		v = int32(0)
	default:
		err = errors.Errorf("unsupported attribute: %v", a.Name)
	}

	return
}

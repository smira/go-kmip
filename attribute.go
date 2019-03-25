package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

import (
	"time"

	"github.com/pkg/errors"
)

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
	case "Unique Identifier":
		v = ""
	case "Object Type":
		v = Enum(0)
	case "Initial Date", "Last Change Date":
		v = time.Time{}
	case "Name":
		v = Name{}
	default:
		err = errors.Errorf("unsupported attribute: %v", a.Name)
	}

	return
}

// Attributes is a sequence of Attribute objects which allows building and search
type Attributes []Attribute

func (attrs Attributes) Get(name string) (val interface{}) {
	for i := range attrs {
		if attrs[i].Name == name {
			val = attrs[i].Value
			break
		}
	}

	return
}

// TemplateAttribute is a Template-Attribute Object Structure
type TemplateAttribute struct {
	Tag `kmip:"TEMPLATE_ATTRIBUTE"`

	Name       Name       `kmip:"NAME"`
	Attributes Attributes `kmip:"ATTRIBUTE"`
}

// Name is a Name Attribute Structure
type Name struct {
	Tag `kmip:"NAME"`

	Value string `kmip:"NAME_VALUE,required"`
	Type  Enum   `kmip:"NAME_TYPE,required"`
}

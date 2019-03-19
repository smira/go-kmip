package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

// CreateRequest is a Create Request Payload
type CreateRequest struct {
	ObjectType        Enum `kmip:"OBJECT_TYPE,required"`
	TemplateAttribute `kmip:"TEMPLATE_ATTRIBUTE,required"`
}

// CreateResponse is a Create Response Payload
type CreateResponse struct {
	ObjectType        Enum   `kmip:"OBJECT_TYPE,required"`
	UniqueIdentifier  string `kmip:"UNIQUE_IDENTIFIER,required"`
	TemplateAttribute `kmip:"TEMPLATE_ATTRIBUTE"`
}

// GetRequest is a Get Request Payload
type GetRequest struct {
	UniqueIdentifier   string                   `kmip:"UNIQUE_IDENTIFIER"`
	KeyFormatType      Enum                     `kmip:"KEY_FORMAT_TYPE"`
	KeyWrapType        Enum                     `kmip:"KEY_WRAP_TYPE"`
	KeyCompressionType Enum                     `kmip:"KEY_COMPRESSION_TYPE"`
	KeyWrappingSpec    KeyWrappingSpecification `kmip:"KEY_WRAPPING_SPECIFICATION"`
}

// GetResponse is a Get Response Payload
type GetResponse struct {
	ObjectType       Enum   `kmip:"OBJECT_TYPE,required"`
	UniqueIdentifier string `kmip:"UNIQUE_IDENTIFIER,required"`
	// Response might contain one of SymmetricKey, Certificate, ...
	SymmetricKey SymmetricKey `kmip:"SYMMETRIC_KEY"`
}

// GetAttributesRequest is a Get Attributes Request Payload
type GetAttributesRequest struct {
	UniqueIdentifier string   `kmip:"UNIQUE_IDENTIFIER"`
	AttributeNames   []string `kmip:"ATTRIBUTE_NAME"`
}

// GetAttributesResponse is a Get Attributes Response Payload
type GetAttributesResponse struct {
	UniqueIdentifier string      `kmip:"UNIQUE_IDENTIFIER,required"`
	Attributes       []Attribute `kmip:"ATTRIBUTE"`
}

// GetAttributeListRequest is a Get Attribute List Request Payload
type GetAttributeListRequest struct {
	UniqueIdentifier string `kmip:"UNIQUE_IDENTIFIER"`
}

// GetAttributeListResponse is a Get Attribute List Response Payload
type GetAttributeListResponse struct {
	UniqueIdentifier string   `kmip:"UNIQUE_IDENTIFIER,required"`
	AttributeNames   []string `kmip:"ATTRIBUTE_NAME"`
}

// DestroyRequest is a Destroy Request Payload
type DestroyRequest struct {
	UniqueIdentifier string `kmip:"UNIQUE_IDENTIFIER"`
}

// DestroyResponse is a Destroy Response Payload
type DestroyResponse struct {
	UniqueIdentifier string `kmip:"UNIQUE_IDENTIFIER,required"`
}

// DiscoverVersionsRequest is a Discover Versions Request Payload
type DiscoverVersionsRequest struct {
	ProtocolVersions []ProtocolVersion `kmip:"PROTOCOL_VERSION"`
}

// DiscoverVersionsResponse is a Discover Versions Response Payload
type DiscoverVersionsResponse struct {
	ProtocolVersions []ProtocolVersion `kmip:"PROTOCOL_VERSION"`
}

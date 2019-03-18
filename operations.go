package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

// CreateRequest is a Create Request Payload
type CreateRequest struct {
	ObjectType        Enum `kmip:"OBJECT_TYPE,required"`
	TemplateAttribute `kmip:"TEMPLATE_ATTRIBUTE,required"`
}

// GetRequest is a Get Request Payload
type GetRequest struct {
	UniqueIdentifier   string                   `kmip:"UNIQUE_IDENTIFIER"`
	KeyFormatType      Enum                     `kmip:"KEY_FORMAT_TYPE"`
	KeyWrapType        Enum                     `kmip:"KEY_WRAP_TYPE"`
	KeyCompressionType Enum                     `kmip:"KEY_COMPRESSION_TYPE"`
	KeyWrappingSpec    KeyWrappingSpecification `kmip:"KEY_WRAPPING_SPECIFICATION"`
}

// DiscoverVersionsRequest is a Discover Versions Request Payload
type DiscoverVersionsRequest struct {
	ProtocolVersions []ProtocolVersion `kmip:"PROTOCOL_VERSION"`
}

// DiscoverVersionsResponse is a Discover Versions Response Payload
type DiscoverVersionsResponse struct {
	ProtocolVersions []ProtocolVersion `kmip:"PROTOCOL_VERSION"`
}

package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

const (
	// KMIP 1.0
	ACTIVATION_DATE                        Tag = 0x420001
	APPLICATION_DATA                       Tag = 0x420002
	APPLICATION_NAMESPACE                  Tag = 0x420003
	APPLICATION_SPECIFIC_INFORMATION       Tag = 0x420004
	ARCHIVE_DATE                           Tag = 0x420005
	ASYNCHRONOUS_CORRELATION_VALUE         Tag = 0x420006
	ASYNCHRONOUS_INDICATOR                 Tag = 0x420007
	ATTRIBUTE                              Tag = 0x420008
	ATTRIBUTE_INDEX                        Tag = 0x420009 // Designated '(Reserved)' in KMIP 2.0
	ATTRIBUTE_NAME                         Tag = 0x42000A
	ATTRIBUTE_VALUE                        Tag = 0x42000B
	AUTHENTICATION                         Tag = 0x42000C
	BATCH_COUNT                            Tag = 0x42000D
	BATCH_ERROR_CONTINUATION_OPTION        Tag = 0x42000E
	BATCH_ITEM                             Tag = 0x42000F
	BATCH_ORDER_OPTION                     Tag = 0x420010
	BLOCK_CIPHER_MODE                      Tag = 0x420011
	CANCELLATION_RESULT                    Tag = 0x420012
	CERTIFICATE                            Tag = 0x420013
	CERTIFICATE_IDENTIFIER                 Tag = 0x420014 // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_ISSUER                     Tag = 0x420015 // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_ISSUER_ALTERNATIVE_NAME    Tag = 0x420016 // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_ISSUER_DISTINGUISHED_NAME  Tag = 0x420017 // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_REQUEST                    Tag = 0x420018
	CERTIFICATE_REQUEST_TYPE               Tag = 0x420019
	CERTIFICATE_SUBJECT                    Tag = 0x42001A // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_SUBJECT_ALTERNATIVE_NAME   Tag = 0x42001B // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_SUBJECT_DISTINGUISHED_NAME Tag = 0x42001C // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_TYPE                       Tag = 0x42001D
	CERTIFICATE_VALUE                      Tag = 0x42001E
	COMMON_TEMPLATE_ATTRIBUTE              Tag = 0x42001F // Designated '(Reserved)' in KMIP 2.0
	COMPROMISE_DATE                        Tag = 0x420020
	COMPROMISE_OCCURRENCE_DATE             Tag = 0x420021
	CONTACT_INFORMATION                    Tag = 0x420022
	CREDENTIAL                             Tag = 0x420023
	CREDENTIAL_TYPE                        Tag = 0x420024
	CREDENTIAL_VALUE                       Tag = 0x420025
	CRITICALITY_INDICATOR                  Tag = 0x420026
	CRT_COEFFICIENT                        Tag = 0x420027
	CRYPTOGRAPHIC_ALGORITHM                Tag = 0x420028
	CRYPTOGRAPHIC_DOMAIN_PARAMETERS        Tag = 0x420029
	CRYPTOGRAPHIC_LENGTH                   Tag = 0x42002A
	CRYPTOGRAPHIC_PARAMETERS               Tag = 0x42002B
	CRYPTOGRAPHIC_USAGE_MASK               Tag = 0x42002C
	CUSTOM_ATTRIBUTE                       Tag = 0x42002D // Designated '(Reserved)' in KMIP 2.0
	D                                      Tag = 0x42002E
	DEACTIVATION_DATE                      Tag = 0x42002F
	DERIVATION_DATA                        Tag = 0x420030
	DERIVATION_METHOD                      Tag = 0x420031
	DERIVATION_PARAMETERS                  Tag = 0x420032
	DESTROY_DATE                           Tag = 0x420033
	DIGEST                                 Tag = 0x420034
	DIGEST_VALUE                           Tag = 0x420035
	ENCRYPTION_KEY_INFORMATION             Tag = 0x420036
	G                                      Tag = 0x420037
	HASHING_ALGORITHM                      Tag = 0x420038
	INITIAL_DATE                           Tag = 0x420039
	INITIALIZATION_VECTOR                  Tag = 0x42003A
	ISSUER                                 Tag = 0x42003B // Deprecated, designated '(Reserved)' in KMIP 2.0
	ITERATION_COUNT                        Tag = 0x42003C
	IV_COUNTER_NONCE                       Tag = 0x42003D
	J                                      Tag = 0x42003E
	KEY                                    Tag = 0x42003F
	KEY_BLOCK                              Tag = 0x420040
	KEY_COMPRESSION_TYPE                   Tag = 0x420041
	KEY_FORMAT_TYPE                        Tag = 0x420042
	KEY_MATERIAL                           Tag = 0x420043
	KEY_PART_IDENTIFIER                    Tag = 0x420044
	KEY_VALUE                              Tag = 0x420045
	KEY_WRAPPING_DATA                      Tag = 0x420046
	KEY_WRAPPING_SPECIFICATION             Tag = 0x420047
	LAST_CHANGE_DATE                       Tag = 0x420048
	LEASE_TIME                             Tag = 0x420049
	LINK                                   Tag = 0x42004A
	LINK_TYPE                              Tag = 0x42004B
	LINKED_OBJECT_IDENTIFIER               Tag = 0x42004C
	MAC_SIGNATURE                          Tag = 0x42004D
	MAC_SIGNATURE_KEY_INFORMATION          Tag = 0x42004E
	MAXIMUM_ITEMS                          Tag = 0x42004F
	MAXIMUM_RESPONSE_SIZE                  Tag = 0x420050
	MESSAGE_EXTENSION                      Tag = 0x420051
	MODULUS                                Tag = 0x420052
	NAME                                   Tag = 0x420053
	NAME_TYPE                              Tag = 0x420054
	NAME_VALUE                             Tag = 0x420055
	OBJECT_GROUP                           Tag = 0x420056
	OBJECT_TYPE                            Tag = 0x420057
	OFFSET                                 Tag = 0x420058
	OPAQUE_DATA_TYPE                       Tag = 0x420059
	OPAQUE_DATA_VALUE                      Tag = 0x42005A
	OPAQUE_OBJECT                          Tag = 0x42005B
	OPERATION                              Tag = 0x42005C
	OPERATION_POLICY_NAME                  Tag = 0x42005D // Designated '(Reserved)' in KMIP 2.0
	P                                      Tag = 0x42005E
	PADDING_METHOD                         Tag = 0x42005F
	PRIME_EXPONENT_P                       Tag = 0x420060
	PRIME_EXPONENT_Q                       Tag = 0x420061
	PRIME_FIELD_SIZE                       Tag = 0x420062
	PRIVATE_EXPONENT                       Tag = 0x420063
	PRIVATE_KEY                            Tag = 0x420064
	PRIVATE_KEY_TEMPLATE_ATTRIBUTE         Tag = 0x420065 // Designated '(Reserved)' in KMIP 2.0
	PRIVATE_KEY_UNIQUE_IDENTIFIER          Tag = 0x420066
	PROCESS_START_DATE                     Tag = 0x420067
	PROTECT_STOP_DATE                      Tag = 0x420068
	PROTOCOL_VERSION                       Tag = 0x420069
	PROTOCOL_VERSION_MAJOR                 Tag = 0x42006A
	PROTOCOL_VERSION_MINOR                 Tag = 0x42006B
	PUBLIC_EXPONENT                        Tag = 0x42006C
	PUBLIC_KEY                             Tag = 0x42006D
	PUBLIC_KEY_TEMPLATE_ATTRIBUTE          Tag = 0x42006E // Designated '(Reserved)' in KMIP 2.0
	PUBLIC_KEY_UNIQUE_IDENTIFIER           Tag = 0x42006F
	PUT_FUNCTION                           Tag = 0x420070
	Q                                      Tag = 0x420071
	Q_STRING                               Tag = 0x420072
	QLENGTH                                Tag = 0x420073
	QUERY_FUNCTION                         Tag = 0x420074
	RECOMMENDED_CURVE                      Tag = 0x420075
	REPLACED_UNIQUE_IDENTIFIER             Tag = 0x420076
	REQUEST_BATCH_ITEM                     Tag = 0x42000F
	REQUEST_HEADER                         Tag = 0x420077
	REQUEST_MESSAGE                        Tag = 0x420078
	REQUEST_PAYLOAD                        Tag = 0x420079
	RESPONSE_BATCH_ITEM                    Tag = 0x42000F
	RESPONSE_HEADER                        Tag = 0x42007A
	RESPONSE_MESSAGE                       Tag = 0x42007B
	RESPONSE_PAYLOAD                       Tag = 0x42007C
	RESULT_MESSAGE                         Tag = 0x42007D
	RESULT_REASON                          Tag = 0x42007E
	RESULT_STATUS                          Tag = 0x42007F
	REVOCATION_MESSAGE                     Tag = 0x420080
	REVOCATION_REASON                      Tag = 0x420081
	REVOCATION_REASON_CODE                 Tag = 0x420082
	KEY_ROLE_TYPE                          Tag = 0x420083
	SALT                                   Tag = 0x420084
	SECRET_DATA                            Tag = 0x420085
	SECRET_DATA_TYPE                       Tag = 0x420086
	SERIAL_NUMBER                          Tag = 0x420087 // Deprecated, designated '(Reserved)' in KMIP 2.0
	SERVER_INFORMATION                     Tag = 0x420088
	SPLIT_KEY                              Tag = 0x420089
	SPLIT_KEY_METHOD                       Tag = 0x42008A
	SPLIT_KEY_PARTS                        Tag = 0x42008B
	SPLIT_KEY_THRESHOLD                    Tag = 0x42008C
	STATE                                  Tag = 0x42008D
	STORAGE_STATUS_MASK                    Tag = 0x42008E
	SYMMETRIC_KEY                          Tag = 0x42008F
	TEMPLATE                               Tag = 0x420090 // Designated '(Reserved)' in KMIP 2.0
	TEMPLATE_ATTRIBUTE                     Tag = 0x420091 // Designated '(Reserved)' in KMIP 2.0
	TIME_STAMP                             Tag = 0x420092
	UNIQUE_BATCH_ITEM_ID                   Tag = 0x420093
	UNIQUE_IDENTIFIER                      Tag = 0x420094
	USAGE_LIMITS                           Tag = 0x420095
	USAGE_LIMITS_COUNT                     Tag = 0x420096
	USAGE_LIMITS_TOTAL                     Tag = 0x420097
	USAGE_LIMITS_UNIT                      Tag = 0x420098
	USERNAME                               Tag = 0x420099
	VALIDITY_DATE                          Tag = 0x42009A
	VALIDITY_INDICATOR                     Tag = 0x42009B
	VENDOR_EXTENSION                       Tag = 0x42009C
	VENDOR_IDENTIFICATION                  Tag = 0x42009D
	WRAPPING_METHOD                        Tag = 0x42009E
	X                                      Tag = 0x42009F
	Y                                      Tag = 0x4200A0
	PASSWORD                               Tag = 0x4200A1
	// KMIP 1.1
	DEVICE_IDENTIFIER            Tag = 0x4200A2
	ENCODING_OPTION              Tag = 0x4200A3
	EXTENSION_INFORMATION        Tag = 0x4200A4
	EXTENSION_NAME               Tag = 0x4200A5
	EXTENSION_TAG                Tag = 0x4200A6
	EXTENSION_TYPE               Tag = 0x4200A7
	FRESH                        Tag = 0x4200A8
	MACHINE_IDENTIFIER           Tag = 0x4200A9
	MEDIA_IDENTIFIER             Tag = 0x4200AA
	NETWORK_IDENTIFIER           Tag = 0x4200AB
	OBJECT_GROUP_MEMBER          Tag = 0x4200AC
	CERTIFICATE_LENGTH           Tag = 0x4200AD
	DIGITAL_SIGNATURE_ALGORITHM  Tag = 0x4200AE
	CERTIFICATE_SERIAL_NUMBER    Tag = 0x4200AF
	DEVICE_SERIAL_NUMBER         Tag = 0x4200B0
	ISSUER_ALTERNATIVE_NAME      Tag = 0x4200B1
	ISSUER_DISTINGUISHED_NAME    Tag = 0x4200B2
	SUBJECT_ALTERNATIVE_NAME     Tag = 0x4200B3
	SUBJECT_DISTINGUISHED_NAME   Tag = 0x4200B4
	X_509_CERTIFICATE_IDENTIFIER Tag = 0x4200B5
	X_509_CERTIFICATE_ISSUER     Tag = 0x4200B6
	X_509_CERTIFICATE_SUBJECT    Tag = 0x4200B7
	// KMIP 1.2
	KEY_VALUE_LOCATION            Tag = 0x4200B8
	KEY_VALUE_LOCATION_VALUE      Tag = 0x4200B9
	KEY_VALUE_LOCATION_TYPE       Tag = 0x4200BA
	KEY_VALUE_PRESENT             Tag = 0x4200BB
	ORIGINAL_CREATION_DATE        Tag = 0x4200BC
	PGP_KEY                       Tag = 0x4200BD
	PGP_KEY_VERSION               Tag = 0x4200BE
	ALTERNATIVE_NAME              Tag = 0x4200BF
	ALTERNATIVE_NAME_VALUE        Tag = 0x4200C0
	ALTERNATIVE_NAME_TYPE         Tag = 0x4200C1
	DATA                          Tag = 0x4200C2
	SIGNATURE_DATA                Tag = 0x4200C3
	DATA_LENGTH                   Tag = 0x4200C4
	RANDOM_IV                     Tag = 0x4200C5
	MAC_DATA                      Tag = 0x4200C6
	ATTESTATION_TYPE              Tag = 0x4200C7
	NONCE                         Tag = 0x4200C8
	NONCE_ID                      Tag = 0x4200C9
	NONCE_VALUE                   Tag = 0x4200CA
	ATTESTATION_MEASUREMENT       Tag = 0x4200CB
	ATTESTATION_ASSERTION         Tag = 0x4200CC
	IV_LENGTH                     Tag = 0x4200CD
	TAG_LENGTH                    Tag = 0x4200CE
	FIXED_FIELD_LENGTH            Tag = 0x4200CF
	COUNTER_LENGTH                Tag = 0x4200D0
	INITIAL_COUNTER_VALUE         Tag = 0x4200D1
	INVOCATION_FIELD_LENGTH       Tag = 0x4200D2
	ATTESTATION_CAPABLE_INDICATOR Tag = 0x4200D3
	// KMIP 1.3
	OFFSET_ITEMS                      Tag = 0x4200D4
	LOCATED_ITEMS                     Tag = 0x4200D5
	CORRELATION_VALUE                 Tag = 0x4200D6
	INIT_INDICATOR                    Tag = 0x4200D7
	FINAL_INDICATOR                   Tag = 0x4200D8
	RNG_PARAMETERS                    Tag = 0x4200D9
	RNG_ALGORITHM                     Tag = 0x4200DA
	DRBG_ALGORITHM                    Tag = 0x4200DB
	FIPS186_VARIATION                 Tag = 0x4200DC
	PREDICTION_RESISTANCE             Tag = 0x4200DD
	RANDOM_NUMBER_GENERATOR           Tag = 0x4200DE
	VALIDATION_INFORMATION            Tag = 0x4200DF
	VALIDATION_AUTHORITY_TYPE         Tag = 0x4200E0
	VALIDATION_AUTHORITY_COUNTRY      Tag = 0x4200E1
	VALIDATION_AUTHORITY_URI          Tag = 0x4200E2
	VALIDATION_VERSION_MAJOR          Tag = 0x4200E3
	VALIDATION_VERSION_MINOR          Tag = 0x4200E4
	VALIDATION_TYPE                   Tag = 0x4200E5
	VALIDATION_LEVEL                  Tag = 0x4200E6
	VALIDATION_CERTIFICATE_IDENTIFIER Tag = 0x4200E7
	VALIDATION_CERTIFICATE_URI        Tag = 0x4200E8
	VALIDATION_VENDOR_URI             Tag = 0x4200E9
	VALIDATION_PROFILE                Tag = 0x4200EA
	PROFILE_INFORMATION               Tag = 0x4200EB
	PROFILE_NAME                      Tag = 0x4200EC
	SERVER_URI                        Tag = 0x4200ED
	SERVER_PORT                       Tag = 0x4200EE
	STREAMING_CAPABILITY              Tag = 0x4200EF
	ASYNCHRONOUS_CAPABILITY           Tag = 0x4200F0
	ATTESTATION_CAPABILITY            Tag = 0x4200F1
	UNWRAP_MODE                       Tag = 0x4200F2
	DESTROY_ACTION                    Tag = 0x4200F3
	SHREDDING_ALGORITHM               Tag = 0x4200F4
	RNG_MODE                          Tag = 0x4200F5
	CLIENT_REGISTRATION_METHOD        Tag = 0x4200F6
	CAPABILITY_INFORMATION            Tag = 0x4200F7
	// KMIP 1.4
	KEY_WRAP_TYPE                            Tag = 0x4200F8
	BATCH_UNDO_CAPABILITY                    Tag = 0x4200F9
	BATCH_CONTINUE_CAPABILITY                Tag = 0x4200FA
	PKCS12_FRIENDLY_NAME                     Tag = 0x4200FB
	DESCRIPTION                              Tag = 0x4200FC
	COMMENT                                  Tag = 0x4200FD
	AUTHENTICATED_ENCRYPTION_ADDITIONAL_DATA Tag = 0x4200FE
	AUTHENTICATED_ENCRYPTION_TAG             Tag = 0x4200FF
	SALT_LENGTH                              Tag = 0x420100
	MASK_GENERATOR                           Tag = 0x420101
	MASK_GENERATOR_HASHING_ALGORITHM         Tag = 0x420102
	P_SOURCE                                 Tag = 0x420103
	TRAILER_FIELD                            Tag = 0x420104
	CLIENT_CORRELATION_VALUE                 Tag = 0x420105
	SERVER_CORRELATION_VALUE                 Tag = 0x420106
	DIGESTED_DATA                            Tag = 0x420107
	CERTIFICATE_SUBJECT_CN                   Tag = 0x420108
	CERTIFICATE_SUBJECT_O                    Tag = 0x420109
	CERTIFICATE_SUBJECT_OU                   Tag = 0x42010A
	CERTIFICATE_SUBJECT_EMAIL                Tag = 0x42010B
	CERTIFICATE_SUBJECT_C                    Tag = 0x42010C
	CERTIFICATE_SUBJECT_ST                   Tag = 0x42010D
	CERTIFICATE_SUBJECT_L                    Tag = 0x42010E
	CERTIFICATE_SUBJECT_UID                  Tag = 0x42010F
	CERTIFICATE_SUBJECT_SERIAL_NUMBER        Tag = 0x420110
	CERTIFICATE_SUBJECT_TITLE                Tag = 0x420111
	CERTIFICATE_SUBJECT_DC                   Tag = 0x420112
	CERTIFICATE_SUBJECT_DN_QUALIFIER         Tag = 0x420113
	CERTIFICATE_ISSUER_CN                    Tag = 0x420114
	CERTIFICATE_ISSUER_O                     Tag = 0x420115
	CERTIFICATE_ISSUER_OU                    Tag = 0x420116
	CERTIFICATE_ISSUER_EMAIL                 Tag = 0x420117
	CERTIFICATE_ISSUER_C                     Tag = 0x420118
	CERTIFICATE_ISSUER_ST                    Tag = 0x420119
	CERTIFICATE_ISSUER_L                     Tag = 0x42011A
	CERTIFICATE_ISSUER_UID                   Tag = 0x42011B
	CERTIFICATE_ISSUER_SERIAL_NUMBER         Tag = 0x42011C
	CERTIFICATE_ISSUER_TITLE                 Tag = 0x42011D
	CERTIFICATE_ISSUER_DC                    Tag = 0x42011E
	CERTIFICATE_ISSUER_DN_QUALIFIER          Tag = 0x42011F
	SENSITIVE                                Tag = 0x420120
	ALWAYS_SENSITIVE                         Tag = 0x420121
	EXTRACTABLE                              Tag = 0x420122
	NEVER_EXTRACTABLE                        Tag = 0x420123
	REPLACE_EXISTING                         Tag = 0x420124
	// KMIP 2.0
	ATTRIBUTES                            Tag = 0x420125
	COMMON_ATTRIBUTES                     Tag = 0x420126
	PRIVATE_KEY_ATTRIBUTES                Tag = 0x420127
	PUBLIC_KEY_ATTRIBUTES                 Tag = 0x420128
	EXTENSION_ENUMERATION                 Tag = 0x420129
	EXTENSION_ATTRIBUTE                   Tag = 0x42012A
	EXTENSION_PARENT_STRUCTURE_TAG        Tag = 0x42012B
	EXTENSION_DESCRIPTION                 Tag = 0x42012C
	SERVER_NAME                           Tag = 0x42012D
	SERVER_SERIAL_NUMBER                  Tag = 0x42012E
	SERVER_VERSION                        Tag = 0x42012F
	SERVER_LOAD                           Tag = 0x420130
	PRODUCT_NAME                          Tag = 0x420131
	BUILD_LEVEL                           Tag = 0x420132
	BUILD_DATE                            Tag = 0x420133
	CLUSTER_INFO                          Tag = 0x420134
	ALTERNATE_FAILOVER_ENDPOINTS          Tag = 0x420135
	SHORT_UNIQUE_IDENTIFIER               Tag = 0x420136
	RESERVED                              Tag = 0x420137
	TAG                                   Tag = 0x420138
	CERTIFICATE_REQUEST_UNIQUE_IDENTIFIER Tag = 0x420139
	NIST_KEY_TYPE                         Tag = 0x42013A
	ATTRIBUTE_REFERENCE                   Tag = 0x42013B
	CURRENT_ATTRIBUTE                     Tag = 0x42013C
	NEW_ATTRIBUTE                         Tag = 0x42013D
	// 0x42013E is designated '(Reserved)' in KMIP 2.0
	// 0x42013F is designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_REQUEST_VALUE Tag = 0x420140
	LOG_MESSAGE               Tag = 0x420141
	PROFILE_VERSION           Tag = 0x420142
	PROFILE_VERSION_MAJOR     Tag = 0x420143
	PROFILE_VERSION_MINOR     Tag = 0x420144
	PROTECTION_LEVEL          Tag = 0x420145
	PROTECTION_PERIOD         Tag = 0x420146
	QUANTUM_SAFE              Tag = 0x420147
	QUANTUM_SAFE_CAPABILITY   Tag = 0x420148
	TICKET                    Tag = 0x420149
	TICKET_TYPE               Tag = 0x42014A
	TICKET_VALUE              Tag = 0x42014B
	REQUEST_COUNT             Tag = 0x42014C
	RIGHTS                    Tag = 0x42014D
	OBJECTS                   Tag = 0x42014E
	OPERATIONS                Tag = 0x42014F
	RIGHT                     Tag = 0x420150
	ENDPOINT_ROLE             Tag = 0x420151
	DEFAULTS_INFORMATION      Tag = 0x420152
	OBJECT_DEFAULTS           Tag = 0x420153
	EPHEMERAL                 Tag = 0x420154
	SERVER_HASHED_PASSWORD    Tag = 0x420155
	ONE_TIME_PASSWORD         Tag = 0x420156
	HASHED_PASSWORD           Tag = 0x420157
	ADJUSTMENT_TYPE           Tag = 0x420158
	PKCS11_INTERFACE          Tag = 0x420159
	PKCS11_FUNCTION           Tag = 0x42015A
	PKCS11_INPUT_PARAMETERS   Tag = 0x42015B
	PKCS11_OUTPUT_PARAMETERS  Tag = 0x42015C
	PKCS11_RETURN_CODE        Tag = 0x42015D
	PROTECTION_STORAGE_MASK   Tag = 0x42015E
	PROTECTION_STORAGE_MASKS  Tag = 0x42015F
	INTEROP_FUNCTION          Tag = 0x420160
	INTEROP_IDENTIFIER        Tag = 0x420161
	ADJUSTMENT_VALUE          Tag = 0x420162
)

const (
	STRUCTURE    Type = 0x01
	INTEGER      Type = 0x02
	LONG_INTEGER Type = 0x03
	BIG_INTEGER  Type = 0x04
	ENUMERATION  Type = 0x05
	BOOLEAN      Type = 0x06
	TEXT_STRING  Type = 0x07
	BYTE_STRING  Type = 0x08
	DATE_TIME    Type = 0x09
	INTERVAL     Type = 0x0A
)

const (
	// KMIP 1.0
	OPERATION_CREATE               Enum = 0x00000001
	OPERATION_CREATE_KEY_PAIR      Enum = 0x00000002
	OPERATION_REGISTER             Enum = 0x00000003
	OPERATION_REKEY                Enum = 0x00000004
	OPERATION_DERIVE_KEY           Enum = 0x00000005
	OPERATION_CERTIFY              Enum = 0x00000006
	OPERATION_RECERTIFY            Enum = 0x00000007
	OPERATION_LOCATE               Enum = 0x00000008
	OPERATION_CHECK                Enum = 0x00000009
	OPERATION_GET                  Enum = 0x0000000A
	OPERATION_GET_ATTRIBUTES       Enum = 0x0000000B
	OPERATION_GET_ATTRIBUTE_LIST   Enum = 0x0000000C
	OPERATION_ADD_ATTRIBUTE        Enum = 0x0000000D
	OPERATION_MODIFY_ATTRIBUTE     Enum = 0x0000000E
	OPERATION_DELETE_ATTRIBUTE     Enum = 0x0000000F
	OPERATION_OBTAIN_LEASE         Enum = 0x00000010
	OPERATION_GET_USAGE_ALLOCATION Enum = 0x00000011
	OPERATION_ACTIVATE             Enum = 0x00000012
	OPERATION_REVOKE               Enum = 0x00000013
	OPERATION_DESTROY              Enum = 0x00000014
	OPERATION_ARCHIVE              Enum = 0x00000015
	OPERATION_RECOVER              Enum = 0x00000016
	OPERATION_VALIDATE             Enum = 0x00000017
	OPERATION_QUERY                Enum = 0x00000018
	OPERATION_CANCEL               Enum = 0x00000019
	OPERATION_POLL                 Enum = 0x0000001A
	OPERATION_NOTIFY               Enum = 0x0000001B
	OPERATION_PUT                  Enum = 0x0000001C
	// KMIP 1.1
	OPERATION_REKEY_KEY_PAIR    Enum = 0x0000001D
	OPERATION_DISCOVER_VERSIONS Enum = 0x0000001E
	// KMIP 1.2
	OPERATION_ENCRYPT          Enum = 0x0000001F
	OPERATION_DECRYPT          Enum = 0x00000020
	OPERATION_SIGN             Enum = 0x00000021
	OPERATION_SIGNATURE_VERIFY Enum = 0x00000022
	OPERATION_MAC              Enum = 0x00000023
	OPERATION_MAC_VERIFY       Enum = 0x00000024
	OPERATION_RNG_RETRIEVE     Enum = 0x00000025
	OPERATION_RNG_SEED         Enum = 0x00000026
	OPERATION_HASH             Enum = 0x00000027
	OPERATION_CREATE_SPLIT_KEY Enum = 0x00000028
	OPERATION_JOIN_SPLIT_KEY   Enum = 0x00000029
	// KMIP 1.4
	OPERATION_IMPORT Enum = 0x0000002A
	OPERATION_EXPORT Enum = 0x0000002B
	// KMIP 2.0
	OPERATION_LOG               Enum = 0x0000002C
	OPERATION_LOGIN             Enum = 0x0000002D
	OPERATION_LOGOUT            Enum = 0x0000002E
	OPERATION_DELEGATED_LOGIN   Enum = 0x0000002F
	OPERATION_ADJUST_ATTRIBUTE  Enum = 0x00000030
	OPERATION_SET_ATTRIBUTE     Enum = 0x00000031
	OPERATION_SET_ENDPOINT_ROLE Enum = 0x00000032
	OPERATION_PKCS_11           Enum = 0x00000033
	OPERATION_INTEROP           Enum = 0x00000034
	OPERATION_REPROVISION       Enum = 0x00000035
)

const (
	// KMIP 1.0
	OBJECT_TYPE_CERTIFICATE   Enum = 0x00000001
	OBJECT_TYPE_SYMMETRIC_KEY Enum = 0x00000002
	OBJECT_TYPE_PUBLIC_KEY    Enum = 0x00000003
	OBJECT_TYPE_PRIVATE_KEY   Enum = 0x00000004
	OBJECT_TYPE_SPLIT_KEY     Enum = 0x00000005
	OBJECT_TYPE_TEMPLATE      Enum = 0x00000006 // Deprecated in KMIP 1.3, designated '(Reserved)' in KMIP 2.0
	OBJECT_TYPE_SECRET_DATA   Enum = 0x00000007
	OBJECT_TYPE_OPAQUE_DATA   Enum = 0x00000008
	// KMIP 1.2
	OBJECT_TYPE_PGP_KEY Enum = 0x00000009
	// KMIP 2.0
	OBJECT_TYPE_CERTIFICATE_REQUEST Enum = 0x0000000A
)

const (
	// KMIP 1.0
	CRYPTO_DES         Enum = 0x00000001
	CRYPTO_TRIPLE_DES  Enum = 0x00000002 // '3DES' is invalid syntax
	CRYPTO_AES         Enum = 0x00000003
	CRYPTO_RSA         Enum = 0x00000004
	CRYPTO_DSA         Enum = 0x00000005
	CRYPTO_ECDSA       Enum = 0x00000006
	CRYPTO_HMAC_SHA1   Enum = 0x00000007
	CRYPTO_HMAC_SHA224 Enum = 0x00000008
	CRYPTO_HMAC_SHA256 Enum = 0x00000009
	CRYPTO_HMAC_SHA384 Enum = 0x0000000A
	CRYPTO_HMAC_SHA512 Enum = 0x0000000B
	CRYPTO_HMAC_MD5    Enum = 0x0000000C
	CRYPTO_DH          Enum = 0x0000000D
	CRYPTO_ECDH        Enum = 0x0000000E
	CRYPTO_ECMQV       Enum = 0x0000000F
	CRYPTO_BLOWFISH    Enum = 0x00000010
	CRYPTO_CAMELLIA    Enum = 0x00000011
	CRYPTO_CAST5       Enum = 0x00000012
	CRYPTO_IDEA        Enum = 0x00000013
	CRYPTO_MARS        Enum = 0x00000014
	CRYPTO_RC2         Enum = 0x00000015
	CRYPTO_RC4         Enum = 0x00000016
	CRYPTO_RC5         Enum = 0x00000017
	CRYPTO_SKIPJACK    Enum = 0x00000018
	CRYPTO_TWOFISH     Enum = 0x00000019
	// KMIP 1.2
	CRYPTO_EC Enum = 0x0000001A
	// KMIP 1.3
	CRYPTO_ONE_TIME_PAD Enum = 0x0000001B
	// KMIP 1.4
	CRYPTO_CHACHA20          Enum = 0x0000001C
	CRYPTO_POLY1305          Enum = 0x0000001D
	CRYPTO_CHACHA20_POLY1305 Enum = 0x0000001E
	CRYPTO_SHA3_224          Enum = 0x0000001F
	CRYPTO_SHA3_256          Enum = 0x00000020
	CRYPTO_SHA3_384          Enum = 0x00000021
	CRYPTO_SHA3_512          Enum = 0x00000022
	CRYPTO_HMAC_SHA3_224     Enum = 0x00000023
	CRYPTO_HMAC_SHA3_256     Enum = 0x00000024
	CRYPTO_HMAC_SHA3_384     Enum = 0x00000025
	CRYPTO_HMAC_SHA3_512     Enum = 0x00000026
	CRYPTO_SHAKE_128         Enum = 0x00000027
	CRYPTO_SHAKE_256         Enum = 0x00000028
	// KMIP 2.0
	CRYPTO_ARIA              Enum = 0x00000029
	CRYPTO_SEED              Enum = 0x0000002A
	CRYPTO_SM2               Enum = 0x0000002B
	CRYPTO_SM3               Enum = 0x0000002C
	CRYPTO_SM4               Enum = 0x0000002D
	CRYPTO_GOST_R_34_10_2012 Enum = 0x0000002E
	CRYPTO_GOST_R_34_11_2012 Enum = 0x0000002F
	CRYPTO_GOST_R_34_13_2015 Enum = 0x00000030
	CRYPTO_GOST_28147_89     Enum = 0x00000031
	CRYPTO_XMSS              Enum = 0x00000032
	CRYPTO_SPHINCS_256       Enum = 0x00000033
	CRYPTO_MCELIECE          Enum = 0x00000034
	CRYPTO_MCELIECE_6960119  Enum = 0x00000035
	CRYPTO_MCELIECE_8192128  Enum = 0x00000036
	CRYPTO_ED25519           Enum = 0x00000037
	CRYPTO_ED448             Enum = 0x00000038
)

var tagMap = map[string]Tag{
	"ACTIVATION_DATE":                          ACTIVATION_DATE,
	"APPLICATION_DATA":                         APPLICATION_DATA,
	"APPLICATION_NAMESPACE":                    APPLICATION_NAMESPACE,
	"APPLICATION_SPECIFIC_INFORMATION":         APPLICATION_SPECIFIC_INFORMATION,
	"ARCHIVE_DATE":                             ARCHIVE_DATE,
	"ASYNCHRONOUS_CORRELATION_VALUE":           ASYNCHRONOUS_CORRELATION_VALUE,
	"ASYNCHRONOUS_INDICATOR":                   ASYNCHRONOUS_INDICATOR,
	"ATTRIBUTE":                                ATTRIBUTE,
	"ATTRIBUTE_INDEX":                          ATTRIBUTE_INDEX,
	"ATTRIBUTE_NAME":                           ATTRIBUTE_NAME,
	"ATTRIBUTE_VALUE":                          ATTRIBUTE_VALUE,
	"AUTHENTICATION":                           AUTHENTICATION,
	"BATCH_COUNT":                              BATCH_COUNT,
	"BATCH_ERROR_CONTINUATION_OPTION":          BATCH_ERROR_CONTINUATION_OPTION,
	"BATCH_ITEM":                               BATCH_ITEM,
	"BATCH_ORDER_OPTION":                       BATCH_ORDER_OPTION,
	"BLOCK_CIPHER_MODE":                        BLOCK_CIPHER_MODE,
	"CANCELLATION_RESULT":                      CANCELLATION_RESULT,
	"CERTIFICATE":                              CERTIFICATE,
	"CERTIFICATE_IDENTIFIER":                   CERTIFICATE_IDENTIFIER,
	"CERTIFICATE_ISSUER":                       CERTIFICATE_ISSUER,
	"CERTIFICATE_ISSUER_ALTERNATIVE_NAME":      CERTIFICATE_ISSUER_ALTERNATIVE_NAME,
	"CERTIFICATE_ISSUER_DISTINGUISHED_NAME":    CERTIFICATE_ISSUER_DISTINGUISHED_NAME,
	"CERTIFICATE_REQUEST":                      CERTIFICATE_REQUEST,
	"CERTIFICATE_REQUEST_TYPE":                 CERTIFICATE_REQUEST_TYPE,
	"CERTIFICATE_SUBJECT":                      CERTIFICATE_SUBJECT,
	"CERTIFICATE_SUBJECT_ALTERNATIVE_NAME":     CERTIFICATE_SUBJECT_ALTERNATIVE_NAME,
	"CERTIFICATE_SUBJECT_DISTINGUISHED_NAME":   CERTIFICATE_SUBJECT_DISTINGUISHED_NAME,
	"CERTIFICATE_TYPE":                         CERTIFICATE_TYPE,
	"CERTIFICATE_VALUE":                        CERTIFICATE_VALUE,
	"COMMON_TEMPLATE_ATTRIBUTE":                COMMON_TEMPLATE_ATTRIBUTE,
	"COMPROMISE_DATE":                          COMPROMISE_DATE,
	"COMPROMISE_OCCURRENCE_DATE":               COMPROMISE_OCCURRENCE_DATE,
	"CONTACT_INFORMATION":                      CONTACT_INFORMATION,
	"CREDENTIAL":                               CREDENTIAL,
	"CREDENTIAL_TYPE":                          CREDENTIAL_TYPE,
	"CREDENTIAL_VALUE":                         CREDENTIAL_VALUE,
	"CRITICALITY_INDICATOR":                    CRITICALITY_INDICATOR,
	"CRT_COEFFICIENT":                          CRT_COEFFICIENT,
	"CRYPTOGRAPHIC_ALGORITHM":                  CRYPTOGRAPHIC_ALGORITHM,
	"CRYPTOGRAPHIC_DOMAIN_PARAMETERS":          CRYPTOGRAPHIC_DOMAIN_PARAMETERS,
	"CRYPTOGRAPHIC_LENGTH":                     CRYPTOGRAPHIC_LENGTH,
	"CRYPTOGRAPHIC_PARAMETERS":                 CRYPTOGRAPHIC_PARAMETERS,
	"CRYPTOGRAPHIC_USAGE_MASK":                 CRYPTOGRAPHIC_USAGE_MASK,
	"CUSTOM_ATTRIBUTE":                         CUSTOM_ATTRIBUTE,
	"D":                                        D,
	"DEACTIVATION_DATE":                        DEACTIVATION_DATE,
	"DERIVATION_DATA":                          DERIVATION_DATA,
	"DERIVATION_METHOD":                        DERIVATION_METHOD,
	"DERIVATION_PARAMETERS":                    DERIVATION_PARAMETERS,
	"DESTROY_DATE":                             DESTROY_DATE,
	"DIGEST":                                   DIGEST,
	"DIGEST_VALUE":                             DIGEST_VALUE,
	"ENCRYPTION_KEY_INFORMATION":               ENCRYPTION_KEY_INFORMATION,
	"G":                                        G,
	"HASHING_ALGORITHM":                        HASHING_ALGORITHM,
	"INITIAL_DATE":                             INITIAL_DATE,
	"INITIALIZATION_VECTOR":                    INITIALIZATION_VECTOR,
	"ISSUER":                                   ISSUER,
	"ITERATION_COUNT":                          ITERATION_COUNT,
	"IV_COUNTER_NONCE":                         IV_COUNTER_NONCE,
	"J":                                        J,
	"KEY":                                      KEY,
	"KEY_BLOCK":                                KEY_BLOCK,
	"KEY_COMPRESSION_TYPE":                     KEY_COMPRESSION_TYPE,
	"KEY_FORMAT_TYPE":                          KEY_FORMAT_TYPE,
	"KEY_MATERIAL":                             KEY_MATERIAL,
	"KEY_PART_IDENTIFIER":                      KEY_PART_IDENTIFIER,
	"KEY_VALUE":                                KEY_VALUE,
	"KEY_WRAPPING_DATA":                        KEY_WRAPPING_DATA,
	"KEY_WRAPPING_SPECIFICATION":               KEY_WRAPPING_SPECIFICATION,
	"LAST_CHANGE_DATE":                         LAST_CHANGE_DATE,
	"LEASE_TIME":                               LEASE_TIME,
	"LINK":                                     LINK,
	"LINK_TYPE":                                LINK_TYPE,
	"LINKED_OBJECT_IDENTIFIER":                 LINKED_OBJECT_IDENTIFIER,
	"MAC_SIGNATURE":                            MAC_SIGNATURE,
	"MAC_SIGNATURE_KEY_INFORMATION":            MAC_SIGNATURE_KEY_INFORMATION,
	"MAXIMUM_ITEMS":                            MAXIMUM_ITEMS,
	"MAXIMUM_RESPONSE_SIZE":                    MAXIMUM_RESPONSE_SIZE,
	"MESSAGE_EXTENSION":                        MESSAGE_EXTENSION,
	"MODULUS":                                  MODULUS,
	"NAME":                                     NAME,
	"NAME_TYPE":                                NAME_TYPE,
	"NAME_VALUE":                               NAME_VALUE,
	"OBJECT_GROUP":                             OBJECT_GROUP,
	"OBJECT_TYPE":                              OBJECT_TYPE,
	"OFFSET":                                   OFFSET,
	"OPAQUE_DATA_TYPE":                         OPAQUE_DATA_TYPE,
	"OPAQUE_DATA_VALUE":                        OPAQUE_DATA_VALUE,
	"OPAQUE_OBJECT":                            OPAQUE_OBJECT,
	"OPERATION":                                OPERATION,
	"OPERATION_POLICY_NAME":                    OPERATION_POLICY_NAME,
	"P":                                        P,
	"PADDING_METHOD":                           PADDING_METHOD,
	"PRIME_EXPONENT_P":                         PRIME_EXPONENT_P,
	"PRIME_EXPONENT_Q":                         PRIME_EXPONENT_Q,
	"PRIME_FIELD_SIZE":                         PRIME_FIELD_SIZE,
	"PRIVATE_EXPONENT":                         PRIVATE_EXPONENT,
	"PRIVATE_KEY":                              PRIVATE_KEY,
	"PRIVATE_KEY_TEMPLATE_ATTRIBUTE":           PRIVATE_KEY_TEMPLATE_ATTRIBUTE,
	"PRIVATE_KEY_UNIQUE_IDENTIFIER":            PRIVATE_KEY_UNIQUE_IDENTIFIER,
	"PROCESS_START_DATE":                       PROCESS_START_DATE,
	"PROTECT_STOP_DATE":                        PROTECT_STOP_DATE,
	"PROTOCOL_VERSION":                         PROTOCOL_VERSION,
	"PROTOCOL_VERSION_MAJOR":                   PROTOCOL_VERSION_MAJOR,
	"PROTOCOL_VERSION_MINOR":                   PROTOCOL_VERSION_MINOR,
	"PUBLIC_EXPONENT":                          PUBLIC_EXPONENT,
	"PUBLIC_KEY":                               PUBLIC_KEY,
	"PUBLIC_KEY_TEMPLATE_ATTRIBUTE":            PUBLIC_KEY_TEMPLATE_ATTRIBUTE,
	"PUBLIC_KEY_UNIQUE_IDENTIFIER":             PUBLIC_KEY_UNIQUE_IDENTIFIER,
	"PUT_FUNCTION":                             PUT_FUNCTION,
	"Q":                                        Q,
	"Q_STRING":                                 Q_STRING,
	"QLENGTH":                                  QLENGTH,
	"QUERY_FUNCTION":                           QUERY_FUNCTION,
	"RECOMMENDED_CURVE":                        RECOMMENDED_CURVE,
	"REPLACED_UNIQUE_IDENTIFIER":               REPLACED_UNIQUE_IDENTIFIER,
	"REQUEST_BATCH_ITEM":                       REQUEST_BATCH_ITEM,
	"REQUEST_HEADER":                           REQUEST_HEADER,
	"REQUEST_MESSAGE":                          REQUEST_MESSAGE,
	"REQUEST_PAYLOAD":                          REQUEST_PAYLOAD,
	"RESPONSE_BATCH_ITEM":                      RESPONSE_BATCH_ITEM,
	"RESPONSE_HEADER":                          RESPONSE_HEADER,
	"RESPONSE_MESSAGE":                         RESPONSE_MESSAGE,
	"RESPONSE_PAYLOAD":                         RESPONSE_PAYLOAD,
	"RESULT_MESSAGE":                           RESULT_MESSAGE,
	"RESULT_REASON":                            RESULT_REASON,
	"RESULT_STATUS":                            RESULT_STATUS,
	"REVOCATION_MESSAGE":                       REVOCATION_MESSAGE,
	"REVOCATION_REASON":                        REVOCATION_REASON,
	"REVOCATION_REASON_CODE":                   REVOCATION_REASON_CODE,
	"KEY_ROLE_TYPE":                            KEY_ROLE_TYPE,
	"SALT":                                     SALT,
	"SECRET_DATA":                              SECRET_DATA,
	"SECRET_DATA_TYPE":                         SECRET_DATA_TYPE,
	"SERIAL_NUMBER":                            SERIAL_NUMBER,
	"SERVER_INFORMATION":                       SERVER_INFORMATION,
	"SPLIT_KEY":                                SPLIT_KEY,
	"SPLIT_KEY_METHOD":                         SPLIT_KEY_METHOD,
	"SPLIT_KEY_PARTS":                          SPLIT_KEY_PARTS,
	"SPLIT_KEY_THRESHOLD":                      SPLIT_KEY_THRESHOLD,
	"STATE":                                    STATE,
	"STORAGE_STATUS_MASK":                      STORAGE_STATUS_MASK,
	"SYMMETRIC_KEY":                            SYMMETRIC_KEY,
	"TEMPLATE":                                 TEMPLATE,
	"TEMPLATE_ATTRIBUTE":                       TEMPLATE_ATTRIBUTE,
	"TIME_STAMP":                               TIME_STAMP,
	"UNIQUE_BATCH_ITEM_ID":                     UNIQUE_BATCH_ITEM_ID,
	"UNIQUE_IDENTIFIER":                        UNIQUE_IDENTIFIER,
	"USAGE_LIMITS":                             USAGE_LIMITS,
	"USAGE_LIMITS_COUNT":                       USAGE_LIMITS_COUNT,
	"USAGE_LIMITS_TOTAL":                       USAGE_LIMITS_TOTAL,
	"USAGE_LIMITS_UNIT":                        USAGE_LIMITS_UNIT,
	"USERNAME":                                 USERNAME,
	"VALIDITY_DATE":                            VALIDITY_DATE,
	"VALIDITY_INDICATOR":                       VALIDITY_INDICATOR,
	"VENDOR_EXTENSION":                         VENDOR_EXTENSION,
	"VENDOR_IDENTIFICATION":                    VENDOR_IDENTIFICATION,
	"WRAPPING_METHOD":                          WRAPPING_METHOD,
	"X":                                        X,
	"Y":                                        Y,
	"PASSWORD":                                 PASSWORD,
	"DEVICE_IDENTIFIER":                        DEVICE_IDENTIFIER,
	"ENCODING_OPTION":                          ENCODING_OPTION,
	"EXTENSION_INFORMATION":                    EXTENSION_INFORMATION,
	"EXTENSION_NAME":                           EXTENSION_NAME,
	"EXTENSION_TAG":                            EXTENSION_TAG,
	"EXTENSION_TYPE":                           EXTENSION_TYPE,
	"FRESH":                                    FRESH,
	"MACHINE_IDENTIFIER":                       MACHINE_IDENTIFIER,
	"MEDIA_IDENTIFIER":                         MEDIA_IDENTIFIER,
	"NETWORK_IDENTIFIER":                       NETWORK_IDENTIFIER,
	"OBJECT_GROUP_MEMBER":                      OBJECT_GROUP_MEMBER,
	"CERTIFICATE_LENGTH":                       CERTIFICATE_LENGTH,
	"DIGITAL_SIGNATURE_ALGORITHM":              DIGITAL_SIGNATURE_ALGORITHM,
	"CERTIFICATE_SERIAL_NUMBER":                CERTIFICATE_SERIAL_NUMBER,
	"DEVICE_SERIAL_NUMBER":                     DEVICE_SERIAL_NUMBER,
	"ISSUER_ALTERNATIVE_NAME":                  ISSUER_ALTERNATIVE_NAME,
	"ISSUER_DISTINGUISHED_NAME":                ISSUER_DISTINGUISHED_NAME,
	"SUBJECT_ALTERNATIVE_NAME":                 SUBJECT_ALTERNATIVE_NAME,
	"SUBJECT_DISTINGUISHED_NAME":               SUBJECT_DISTINGUISHED_NAME,
	"X_509_CERTIFICATE_IDENTIFIER":             X_509_CERTIFICATE_IDENTIFIER,
	"X_509_CERTIFICATE_ISSUER":                 X_509_CERTIFICATE_ISSUER,
	"X_509_CERTIFICATE_SUBJECT":                X_509_CERTIFICATE_SUBJECT,
	"KEY_VALUE_LOCATION":                       KEY_VALUE_LOCATION,
	"KEY_VALUE_LOCATION_VALUE":                 KEY_VALUE_LOCATION_VALUE,
	"KEY_VALUE_LOCATION_TYPE":                  KEY_VALUE_LOCATION_TYPE,
	"KEY_VALUE_PRESENT":                        KEY_VALUE_PRESENT,
	"ORIGINAL_CREATION_DATE":                   ORIGINAL_CREATION_DATE,
	"PGP_KEY":                                  PGP_KEY,
	"PGP_KEY_VERSION":                          PGP_KEY_VERSION,
	"ALTERNATIVE_NAME":                         ALTERNATIVE_NAME,
	"ALTERNATIVE_NAME_VALUE":                   ALTERNATIVE_NAME_VALUE,
	"ALTERNATIVE_NAME_TYPE":                    ALTERNATIVE_NAME_TYPE,
	"DATA":                                     DATA,
	"SIGNATURE_DATA":                           SIGNATURE_DATA,
	"DATA_LENGTH":                              DATA_LENGTH,
	"RANDOM_IV":                                RANDOM_IV,
	"MAC_DATA":                                 MAC_DATA,
	"ATTESTATION_TYPE":                         ATTESTATION_TYPE,
	"NONCE":                                    NONCE,
	"NONCE_ID":                                 NONCE_ID,
	"NONCE_VALUE":                              NONCE_VALUE,
	"ATTESTATION_MEASUREMENT":                  ATTESTATION_MEASUREMENT,
	"ATTESTATION_ASSERTION":                    ATTESTATION_ASSERTION,
	"IV_LENGTH":                                IV_LENGTH,
	"TAG_LENGTH":                               TAG_LENGTH,
	"FIXED_FIELD_LENGTH":                       FIXED_FIELD_LENGTH,
	"COUNTER_LENGTH":                           COUNTER_LENGTH,
	"INITIAL_COUNTER_VALUE":                    INITIAL_COUNTER_VALUE,
	"INVOCATION_FIELD_LENGTH":                  INVOCATION_FIELD_LENGTH,
	"ATTESTATION_CAPABLE_INDICATOR":            ATTESTATION_CAPABLE_INDICATOR,
	"OFFSET_ITEMS":                             OFFSET_ITEMS,
	"LOCATED_ITEMS":                            LOCATED_ITEMS,
	"CORRELATION_VALUE":                        CORRELATION_VALUE,
	"INIT_INDICATOR":                           INIT_INDICATOR,
	"FINAL_INDICATOR":                          FINAL_INDICATOR,
	"RNG_PARAMETERS":                           RNG_PARAMETERS,
	"RNG_ALGORITHM":                            RNG_ALGORITHM,
	"DRBG_ALGORITHM":                           DRBG_ALGORITHM,
	"FIPS186_VARIATION":                        FIPS186_VARIATION,
	"PREDICTION_RESISTANCE":                    PREDICTION_RESISTANCE,
	"RANDOM_NUMBER_GENERATOR":                  RANDOM_NUMBER_GENERATOR,
	"VALIDATION_INFORMATION":                   VALIDATION_INFORMATION,
	"VALIDATION_AUTHORITY_TYPE":                VALIDATION_AUTHORITY_TYPE,
	"VALIDATION_AUTHORITY_COUNTRY":             VALIDATION_AUTHORITY_COUNTRY,
	"VALIDATION_AUTHORITY_URI":                 VALIDATION_AUTHORITY_URI,
	"VALIDATION_VERSION_MAJOR":                 VALIDATION_VERSION_MAJOR,
	"VALIDATION_VERSION_MINOR":                 VALIDATION_VERSION_MINOR,
	"VALIDATION_TYPE":                          VALIDATION_TYPE,
	"VALIDATION_LEVEL":                         VALIDATION_LEVEL,
	"VALIDATION_CERTIFICATE_IDENTIFIER":        VALIDATION_CERTIFICATE_IDENTIFIER,
	"VALIDATION_CERTIFICATE_URI":               VALIDATION_CERTIFICATE_URI,
	"VALIDATION_VENDOR_URI":                    VALIDATION_VENDOR_URI,
	"VALIDATION_PROFILE":                       VALIDATION_PROFILE,
	"PROFILE_INFORMATION":                      PROFILE_INFORMATION,
	"PROFILE_NAME":                             PROFILE_NAME,
	"SERVER_URI":                               SERVER_URI,
	"SERVER_PORT":                              SERVER_PORT,
	"STREAMING_CAPABILITY":                     STREAMING_CAPABILITY,
	"ASYNCHRONOUS_CAPABILITY":                  ASYNCHRONOUS_CAPABILITY,
	"ATTESTATION_CAPABILITY":                   ATTESTATION_CAPABILITY,
	"UNWRAP_MODE":                              UNWRAP_MODE,
	"DESTROY_ACTION":                           DESTROY_ACTION,
	"SHREDDING_ALGORITHM":                      SHREDDING_ALGORITHM,
	"RNG_MODE":                                 RNG_MODE,
	"CLIENT_REGISTRATION_METHOD":               CLIENT_REGISTRATION_METHOD,
	"CAPABILITY_INFORMATION":                   CAPABILITY_INFORMATION,
	"KEY_WRAP_TYPE":                            KEY_WRAP_TYPE,
	"BATCH_UNDO_CAPABILITY":                    BATCH_UNDO_CAPABILITY,
	"BATCH_CONTINUE_CAPABILITY":                BATCH_CONTINUE_CAPABILITY,
	"PKCS12_FRIENDLY_NAME":                     PKCS12_FRIENDLY_NAME,
	"DESCRIPTION":                              DESCRIPTION,
	"COMMENT":                                  COMMENT,
	"AUTHENTICATED_ENCRYPTION_ADDITIONAL_DATA": AUTHENTICATED_ENCRYPTION_ADDITIONAL_DATA,
	"AUTHENTICATED_ENCRYPTION_TAG":             AUTHENTICATED_ENCRYPTION_TAG,
	"SALT_LENGTH":                              SALT_LENGTH,
	"MASK_GENERATOR":                           MASK_GENERATOR,
	"MASK_GENERATOR_HASHING_ALGORITHM":         MASK_GENERATOR_HASHING_ALGORITHM,
	"P_SOURCE":                                 P_SOURCE,
	"TRAILER_FIELD":                            TRAILER_FIELD,
	"CLIENT_CORRELATION_VALUE":                 CLIENT_CORRELATION_VALUE,
	"SERVER_CORRELATION_VALUE":                 SERVER_CORRELATION_VALUE,
	"DIGESTED_DATA":                            DIGESTED_DATA,
	"CERTIFICATE_SUBJECT_CN":                   CERTIFICATE_SUBJECT_CN,
	"CERTIFICATE_SUBJECT_O":                    CERTIFICATE_SUBJECT_O,
	"CERTIFICATE_SUBJECT_OU":                   CERTIFICATE_SUBJECT_OU,
	"CERTIFICATE_SUBJECT_EMAIL":                CERTIFICATE_SUBJECT_EMAIL,
	"CERTIFICATE_SUBJECT_C":                    CERTIFICATE_SUBJECT_C,
	"CERTIFICATE_SUBJECT_ST":                   CERTIFICATE_SUBJECT_ST,
	"CERTIFICATE_SUBJECT_L":                    CERTIFICATE_SUBJECT_L,
	"CERTIFICATE_SUBJECT_UID":                  CERTIFICATE_SUBJECT_UID,
	"CERTIFICATE_SUBJECT_SERIAL_NUMBER":        CERTIFICATE_SUBJECT_SERIAL_NUMBER,
	"CERTIFICATE_SUBJECT_TITLE":                CERTIFICATE_SUBJECT_TITLE,
	"CERTIFICATE_SUBJECT_DC":                   CERTIFICATE_SUBJECT_DC,
	"CERTIFICATE_SUBJECT_DN_QUALIFIER":         CERTIFICATE_SUBJECT_DN_QUALIFIER,
	"CERTIFICATE_ISSUER_CN":                    CERTIFICATE_ISSUER_CN,
	"CERTIFICATE_ISSUER_O":                     CERTIFICATE_ISSUER_O,
	"CERTIFICATE_ISSUER_OU":                    CERTIFICATE_ISSUER_OU,
	"CERTIFICATE_ISSUER_EMAIL":                 CERTIFICATE_ISSUER_EMAIL,
	"CERTIFICATE_ISSUER_C":                     CERTIFICATE_ISSUER_C,
	"CERTIFICATE_ISSUER_ST":                    CERTIFICATE_ISSUER_ST,
	"CERTIFICATE_ISSUER_L":                     CERTIFICATE_ISSUER_L,
	"CERTIFICATE_ISSUER_UID":                   CERTIFICATE_ISSUER_UID,
	"CERTIFICATE_ISSUER_SERIAL_NUMBER":         CERTIFICATE_ISSUER_SERIAL_NUMBER,
	"CERTIFICATE_ISSUER_TITLE":                 CERTIFICATE_ISSUER_TITLE,
	"CERTIFICATE_ISSUER_DC":                    CERTIFICATE_ISSUER_DC,
	"CERTIFICATE_ISSUER_DN_QUALIFIER":          CERTIFICATE_ISSUER_DN_QUALIFIER,
	"SENSITIVE":                                SENSITIVE,
	"ALWAYS_SENSITIVE":                         ALWAYS_SENSITIVE,
	"EXTRACTABLE":                              EXTRACTABLE,
	"NEVER_EXTRACTABLE":                        NEVER_EXTRACTABLE,
	"REPLACE_EXISTING":                         REPLACE_EXISTING,
	"ATTRIBUTES":                               ATTRIBUTES,
	"COMMON_ATTRIBUTES":                        COMMON_ATTRIBUTES,
	"PRIVATE_KEY_ATTRIBUTES":                   PRIVATE_KEY_ATTRIBUTES,
	"PUBLIC_KEY_ATTRIBUTES":                    PUBLIC_KEY_ATTRIBUTES,
	"EXTENSION_ENUMERATION":                    EXTENSION_ENUMERATION,
	"EXTENSION_ATTRIBUTE":                      EXTENSION_ATTRIBUTE,
	"EXTENSION_PARENT_STRUCTURE_TAG":           EXTENSION_PARENT_STRUCTURE_TAG,
	"EXTENSION_DESCRIPTION":                    EXTENSION_DESCRIPTION,
	"SERVER_NAME":                              SERVER_NAME,
	"SERVER_SERIAL_NUMBER":                     SERVER_SERIAL_NUMBER,
	"SERVER_VERSION":                           SERVER_VERSION,
	"SERVER_LOAD":                              SERVER_LOAD,
	"PRODUCT_NAME":                             PRODUCT_NAME,
	"BUILD_LEVEL":                              BUILD_LEVEL,
	"BUILD_DATE":                               BUILD_DATE,
	"CLUSTER_INFO":                             CLUSTER_INFO,
	"ALTERNATE_FAILOVER_ENDPOINTS":             ALTERNATE_FAILOVER_ENDPOINTS,
	"SHORT_UNIQUE_IDENTIFIER":                  SHORT_UNIQUE_IDENTIFIER,
	"RESERVED":                                 RESERVED,
	"TAG":                                      TAG,
	"CERTIFICATE_REQUEST_UNIQUE_IDENTIFIER":    CERTIFICATE_REQUEST_UNIQUE_IDENTIFIER,
	"NIST_KEY_TYPE":                            NIST_KEY_TYPE,
	"ATTRIBUTE_REFERENCE":                      ATTRIBUTE_REFERENCE,
	"CURRENT_ATTRIBUTE":                        CURRENT_ATTRIBUTE,
	"NEW_ATTRIBUTE":                            NEW_ATTRIBUTE,
	"CERTIFICATE_REQUEST_VALUE":                CERTIFICATE_REQUEST_VALUE,
	"LOG_MESSAGE":                              LOG_MESSAGE,
	"PROFILE_VERSION":                          PROFILE_VERSION,
	"PROFILE_VERSION_MAJOR":                    PROFILE_VERSION_MAJOR,
	"PROFILE_VERSION_MINOR":                    PROFILE_VERSION_MINOR,
	"PROTECTION_LEVEL":                         PROTECTION_LEVEL,
	"PROTECTION_PERIOD":                        PROTECTION_PERIOD,
	"QUANTUM_SAFE":                             QUANTUM_SAFE,
	"QUANTUM_SAFE_CAPABILITY":                  QUANTUM_SAFE_CAPABILITY,
	"TICKET":                                   TICKET,
	"TICKET_TYPE":                              TICKET_TYPE,
	"TICKET_VALUE":                             TICKET_VALUE,
	"REQUEST_COUNT":                            REQUEST_COUNT,
	"RIGHTS":                                   RIGHTS,
	"OBJECTS":                                  OBJECTS,
	"OPERATIONS":                               OPERATIONS,
	"RIGHT":                                    RIGHT,
	"ENDPOINT_ROLE":                            ENDPOINT_ROLE,
	"DEFAULTS_INFORMATION":                     DEFAULTS_INFORMATION,
	"OBJECT_DEFAULTS":                          OBJECT_DEFAULTS,
	"EPHEMERAL":                                EPHEMERAL,
	"SERVER_HASHED_PASSWORD":                   SERVER_HASHED_PASSWORD,
	"ONE_TIME_PASSWORD":                        ONE_TIME_PASSWORD,
	"HASHED_PASSWORD":                          HASHED_PASSWORD,
	"ADJUSTMENT_TYPE":                          ADJUSTMENT_TYPE,
	"PKCS11_INTERFACE":                         PKCS11_INTERFACE,
	"PKCS11_FUNCTION":                          PKCS11_FUNCTION,
	"PKCS11_INPUT_PARAMETERS":                  PKCS11_INPUT_PARAMETERS,
	"PKCS11_OUTPUT_PARAMETERS":                 PKCS11_OUTPUT_PARAMETERS,
	"PKCS11_RETURN_CODE":                       PKCS11_RETURN_CODE,
	"PROTECTION_STORAGE_MASK":                  PROTECTION_STORAGE_MASK,
	"PROTECTION_STORAGE_MASKS":                 PROTECTION_STORAGE_MASKS,
	"INTEROP_FUNCTION":                         INTEROP_FUNCTION,
	"INTEROP_IDENTIFIER":                       INTEROP_IDENTIFIER,
	"ADJUSTMENT_VALUE":                         ADJUSTMENT_VALUE,
}

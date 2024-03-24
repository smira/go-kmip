package kmip

type Certificate struct {
	Tag `kmip:"CERTIFICATE"`

	CertificateType  CertificateType `kmip:"CERTIFICATE_TYPE,required"`
	CertificateValue []byte          `kmip:"CERTIFICATE_VALUE,required"`
}

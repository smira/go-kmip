package kmip

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DecoderSuite struct {
	suite.Suite
}

func (s *DecoderSuite) parseSpecValue(val string) []byte {
	val = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(val, "_", ""), "|", ""), " ", "")

	res, err := hex.DecodeString(val)
	s.Require().NoError(err)

	return res
}

func (s *DecoderSuite) TestReadInteger() {
	v, err := NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 02 | 00 00 00 04 | 00 00 00 08 00 00 00 00"))).readInteger(COMPROMISE_DATE)
	s.Assert().NoError(err)
	s.Assert().EqualValues(8, v)

	// padding missing
	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 02 | 00 00 00 04 | 00 00 00 08 00 00 00 "))).readInteger(COMPROMISE_DATE)
	s.Assert().EqualError(err, "unexpected EOF")

	// no value
	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 02 | 00 00 00 04 | 00"))).readInteger(COMPROMISE_DATE)
	s.Assert().EqualError(err, "unexpected EOF")

	// no length
	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 02 | 00 00 "))).readInteger(COMPROMISE_DATE)
	s.Assert().EqualError(err, "unexpected EOF")

	// no type
	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | "))).readInteger(COMPROMISE_DATE)
	s.Assert().EqualError(err, "EOF")

	// no tag
	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42"))).readInteger(COMPROMISE_DATE)
	s.Assert().EqualError(err, "unexpected EOF")

	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 21 | 02 | 00 00 00 04 | 00 00 00 08 00 00 00 "))).readInteger(CRYPTOGRAPHIC_ALGORITHM)
	s.Assert().EqualError(err, "expecting tag 420028, but 420021 was encountered")

	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 01 | 00 00 00 04 | 00 00 00 08 00 00 00 00"))).readInteger(COMPROMISE_DATE)
	s.Assert().EqualError(err, "expecting type 2, but 1 was encountered")

	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 02 | 00 00 00 03 | 00 00 00 08 00 00 00 00"))).readInteger(COMPROMISE_DATE)
	s.Assert().EqualError(err, "expecting length 4, but 3 was encountered")
}

func (s *DecoderSuite) TestReadLongInteger() {
	v, err := NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 03 | 00 00 00 08 | 01 B6 9B 4B A5 74 92 00"))).readLongInteger(COMPROMISE_DATE)
	s.Assert().NoError(err)
	s.Assert().EqualValues(123456789000000000, v)
}

func (s *DecoderSuite) TestReadEnum() {
	v, err := NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 05 | 00 00 00 04 | 00 00 00 FF 00 00 00 00"))).readEnum(COMPROMISE_DATE)
	s.Assert().NoError(err)
	s.Assert().EqualValues(255, v)
}

func (s *DecoderSuite) TestReadBool() {
	v, err := NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 06 | 00 00 00 08 | 00 00 00 00 00 00 00 01"))).readBool(COMPROMISE_DATE)
	s.Assert().NoError(err)
	s.Assert().True(v)

	v, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 06 | 00 00 00 08 | 00 00 00 00 00 00 00 00"))).readBool(COMPROMISE_DATE)
	s.Assert().NoError(err)
	s.Assert().False(v)

	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 06 | 00 00 00 08 | 00 00 00 00 00 00 00 03"))).readBool(COMPROMISE_DATE)
	s.Assert().EqualError(err, "unexpected boolean value: [0 0 0 0 0 0 0 3]")

	_, err = NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 06 | 00 00 00 08 | 00 00 00 00 01 00 00 00"))).readBool(COMPROMISE_DATE)
	s.Assert().EqualError(err, "unexpected boolean value: [0 0 0 0 1 0 0 0]")
}

func (s *DecoderSuite) TestReadString() {
	n, v, err := NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 07 | 00 00 00 0B | 48 65 6C 6C 6F 20 57 6F 72 6C 64 00 00 00 00 00"))).readString(COMPROMISE_DATE)
	s.Assert().NoError(err)
	s.Assert().EqualValues("Hello World", v)
	s.Assert().EqualValues(24, n)
}

func (s *DecoderSuite) TestReadBytes() {
	n, v, err := NewDecoder(bytes.NewReader(s.parseSpecValue("42 00 20 | 08 | 00 00 00 03 | 01 02 03 00 00 00 00 00"))).readBytes(COMPROMISE_DATE)
	s.Assert().NoError(err)
	s.Assert().EqualValues([]byte{1, 2, 3}, v)
	s.Assert().EqualValues(16, n)
}

func TestDecoderSuite(t *testing.T) {
	suite.Run(t, new(DecoderSuite))
}

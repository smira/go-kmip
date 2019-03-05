package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type EncoderSuite struct {
	suite.Suite
}

func (s *EncoderSuite) parseSpecValue(val string) []byte {
	val = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(val, "_", ""), "|", ""), " ", "")

	res, err := hex.DecodeString(val)
	s.Require().NoError(err)

	return res
}

func (s *EncoderSuite) TestWriteInteger() {
	var buf bytes.Buffer

	err := NewEncoder(&buf).writeInteger(COMPROMISE_DATE, 8)
	s.Assert().NoError(err)

	s.Assert().EqualValues(s.parseSpecValue("42 00 20 | 02 | 00 00 00 04 | 00 00 00 08 00 00 00 00"), buf.Bytes())
}

func (s *EncoderSuite) TestWriteLongInteger() {
	var buf bytes.Buffer

	err := NewEncoder(&buf).writeLongInteger(COMPROMISE_DATE, 123456789000000000)
	s.Assert().NoError(err)

	s.Assert().EqualValues(s.parseSpecValue("42 00 20 | 03 | 00 00 00 08 | 01 B6 9B 4B A5 74 92 00"), buf.Bytes())
}

func (s *EncoderSuite) TestWriteEnum() {
	var buf bytes.Buffer

	err := NewEncoder(&buf).writeEnum(COMPROMISE_DATE, Enum(255))
	s.Assert().NoError(err)

	s.Assert().EqualValues(s.parseSpecValue("42 00 20 | 05 | 00 00 00 04 | 00 00 00 FF 00 00 00 00"), buf.Bytes())
}

func (s *EncoderSuite) TestWriteBool() {
	var buf bytes.Buffer

	err := NewEncoder(&buf).writeBool(COMPROMISE_DATE, true)
	s.Assert().NoError(err)

	s.Assert().EqualValues(s.parseSpecValue("42 00 20 | 06 | 00 00 00 08 | 00 00 00 00 00 00 00 01"), buf.Bytes())
}

func (s *EncoderSuite) TestWriteBytes() {
	var buf bytes.Buffer

	err := NewEncoder(&buf).writeBytes(COMPROMISE_DATE, []byte{1, 2, 3})
	s.Assert().NoError(err)

	s.Assert().EqualValues(s.parseSpecValue("42 00 20 | 08 | 00 00 00 03 | 01 02 03 00 00 00 00 00"), buf.Bytes())
}

func (s *EncoderSuite) TestWriteString() {
	var buf bytes.Buffer

	err := NewEncoder(&buf).writeString(COMPROMISE_DATE, "Hello World")
	s.Assert().NoError(err)

	s.Assert().EqualValues(s.parseSpecValue("42 00 20 | 07 | 00 00 00 0B | 48 65 6C 6C 6F 20 57 6F 72 6C 64 00 00 00 00 00"), buf.Bytes())
}

func (s *EncoderSuite) TestWriteTime() {
	var buf bytes.Buffer

	t, _ := time.Parse(time.RFC3339, "2008-03-14T11:56:40Z")

	err := NewEncoder(&buf).writeTime(COMPROMISE_DATE, t)
	s.Assert().NoError(err)

	s.Assert().EqualValues(s.parseSpecValue("42 00 20 | 09 | 00 00 00 08 | 00 00 00 00 47 DA 67 F8"), buf.Bytes())
}

func (s *EncoderSuite) TestWriteInterval() {
	var buf bytes.Buffer

	err := NewEncoder(&buf).writeDuration(COMPROMISE_DATE, 10*24*time.Hour)
	s.Assert().NoError(err)

	s.Assert().EqualValues(s.parseSpecValue("42 00 20 | 0A | 00 00 00 04 | 00 0D 2F 00 00 00 00 00"), buf.Bytes())
}

func TestEncoderSuite(t *testing.T) {
	suite.Run(t, new(EncoderSuite))
}

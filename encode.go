package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

import (
	"encoding/binary"
	"io"
)

// Encoder implements basic type encoding to TTLV KMIP encoding
type Encoder struct {
	w io.Writer
}

// NewEncoder builds encoder writing to w
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

func (e *Encoder) writeTagTypeLength(t Tag, typ Type, l uint32) (err error) {
	var (
		b  [8]byte
		tt [4]byte
	)

	binary.BigEndian.PutUint32(tt[:], uint32(t))

	copy(b[:3], tt[1:])
	b[3] = byte(typ)
	binary.BigEndian.PutUint32(b[4:], l)

	_, err = e.w.Write(b[:])
	return
}

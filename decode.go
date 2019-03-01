package kmip

import (
	"bufio"
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

// Decoder implements KMIP protocol decoding
type Decoder struct {
	r io.Reader
	s io.ByteScanner
}

// NewDecoder builds Decoder which reads from r
//
// Buffering can be disabled by providing reader
// which implements io.ByteScanner
func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{
		r: r,
	}

	if s, ok := r.(io.ByteScanner); ok {
		d.s = s
	} else {
		br := bufio.NewReader(r)
		d.r = br
		d.s = br
	}

	return d
}

func (d *Decoder) readTag() (t Tag, err error) {
	var b [3]byte

	_, err = io.ReadFull(d.r, b[:])
	if err != nil {
		return
	}

	t = Tag(binary.BigEndian.Uint32(append([]byte{0}, b[:]...)))

	return
}

func (d *Decoder) expectTag(expected Tag) error {
	t, err := d.readTag()
	if err != nil {
		return err
	}

	if expected != t {
		return errors.Errorf("expecting tag %x, but %x was encountered", expected, t)
	}

	return nil
}

func (d *Decoder) readType() (t Type, err error) {
	var b byte

	b, err = d.s.ReadByte()
	t = Type(b)

	return
}

func (d *Decoder) expectType(expected Type) error {
	t, err := d.readType()
	if err != nil {
		return err
	}

	if expected != t {
		return errors.Errorf("expecting type %d, but %d was encountered", expected, t)
	}

	return nil
}

func (d *Decoder) readLength() (l uint32, err error) {
	var b [4]byte

	_, err = io.ReadFull(d.r, b[:])
	if err != nil {
		return
	}

	l = binary.BigEndian.Uint32(b[:])

	return
}

func (d *Decoder) expectLength(expected uint32) error {
	l, err := d.readLength()
	if err != nil {
		return err
	}

	if expected != l {
		return errors.Errorf("expecting length %d, but %d was encountered", expected, l)
	}

	return nil
}

func (d *Decoder) readInteger(expectedTag Tag) (v int32, err error) {
	if err = d.expectTag(expectedTag); err != nil {
		return
	}

	if err = d.expectType(INTEGER); err != nil {
		return
	}

	if err = d.expectLength(4); err != nil {
		return
	}

	var b [4]byte

	_, err = io.ReadFull(d.r, b[:])
	if err != nil {
		return
	}

	v = int32(binary.BigEndian.Uint32(b[:]))

	// padding
	_, err = io.ReadFull(d.r, b[:])
	if err != nil {
		return
	}

	return
}

func (d *Decoder) readLongInteger(expectedTag Tag) (v int64, err error) {
	if err = d.expectTag(expectedTag); err != nil {
		return
	}

	if err = d.expectType(LONG_INTEGER); err != nil {
		return
	}

	if err = d.expectLength(8); err != nil {
		return
	}

	var b [8]byte

	_, err = io.ReadFull(d.r, b[:])
	if err != nil {
		return
	}

	v = int64(binary.BigEndian.Uint64(b[:]))

	return
}

func (d *Decoder) readEnum(expectedTag Tag) (v Enum, err error) {
	if err = d.expectTag(expectedTag); err != nil {
		return
	}

	if err = d.expectType(ENUMERATION); err != nil {
		return
	}

	if err = d.expectLength(4); err != nil {
		return
	}

	var b [4]byte

	_, err = io.ReadFull(d.r, b[:])
	if err != nil {
		return
	}

	v = Enum(binary.BigEndian.Uint32(b[:]))

	// padding
	_, err = io.ReadFull(d.r, b[:])
	if err != nil {
		return
	}

	return
}

func (d *Decoder) readBool(expectedTag Tag) (v bool, err error) {
	if err = d.expectTag(expectedTag); err != nil {
		return
	}

	if err = d.expectType(BOOLEAN); err != nil {
		return
	}

	if err = d.expectLength(8); err != nil {
		return
	}

	var b [8]byte

	_, err = io.ReadFull(d.r, b[:])
	if err != nil {
		return
	}

	for i := 0; i < 7; i++ {
		if b[i] != 0 {
			err = errors.Errorf("unexpected boolean value: %v", b)
			return
		}
	}

	switch b[7] {
	case 1:
		v = true
	case 0:
		v = false
	default:
		err = errors.Errorf("unexpected boolean value: %v", b)
	}

	return
}

func (d *Decoder) readByteSlice(expectedTag Tag, expectedType Type) (n int, v []byte, err error) {
	if err = d.expectTag(expectedTag); err != nil {
		return
	}

	if err = d.expectType(expectedType); err != nil {
		return
	}

	var l uint32
	if l, err = d.readLength(); err != nil {
		return
	}

	v = make([]byte, l)
	_, err = io.ReadFull(d.r, v)
	if err != nil {
		return
	}

	n = int(l) + 8

	// padding
	var b [8]byte
	if l%8 != 0 {
		_, err = io.ReadFull(d.r, b[:8-l%8])
		if err != nil {
			return
		}
		n += int(8 - l%8)
	}

	return
}

func (d *Decoder) readBytes(expectedTag Tag) (n int, v []byte, err error) {
	n, v, err = d.readByteSlice(expectedTag, BYTE_STRING)
	return
}

func (d *Decoder) readString(expectedTag Tag) (n int, v string, err error) {
	var b []byte
	n, b, err = d.readByteSlice(expectedTag, TEXT_STRING)
	v = string(b)
	return
}

// Copyright (C) 2013, 2015 by Maxim Bublis <b@codemonkey.ru>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package uuid

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	bytes1 := []byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	if !bytes.Equal(u.Bytes(), bytes1) {
		t.Errorf("Incorrect bytes representation for UUID: %s", u)
	}
}

func TestString(t *testing.T) {
	if NamespaceDNS.String() != "6ba7b810-9dad-11d1-80b4-00c04fd430c8" {
		t.Errorf("Incorrect string representation for UUID: %s", NamespaceDNS.String())
	}
}

func TestBase64(t *testing.T) {
	tests := []struct {
		u   UUID
		exp string
	}{
		{
			u:   UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8},
			exp: "a6e4EJ2tEdGAtADAT9QwyA",
		},
		{
			u:   UUID{0xf8, 0x6, 0x9a, 0xa8, 0xb8, 0x95, 0xde, 0x4f, 0x9c, 0x8a, 0x74, 0x28, 0x88, 0x8b, 0xa7, 0xa},
			exp: "+AaaqLiV3k+cinQoiIunCg",
		},
		{
			u:   UUID{0x44, 0xca, 0x1a, 0x1f, 0xfa, 0x63, 0xfe, 0x42, 0xbf, 0xd, 0x72, 0x6, 0xce, 0xa4, 0x47, 0x65},
			exp: "RMoaH/pj/kK/DXIGzqRHZQ",
		},
	}

	for _, test := range tests {
		if test.u.Base64() != test.exp {
			t.Errorf("Incorrect base64 representation of UUID: %s", test.exp)
		}
	}
}

func TestBase64URL(t *testing.T) {
	tests := []struct {
		u   UUID
		exp string
	}{
		{
			u:   UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8},
			exp: "a6e4EJ2tEdGAtADAT9QwyA",
		},
		{
			u:   UUID{0xf8, 0x6, 0x9a, 0xa8, 0xb8, 0x95, 0xde, 0x4f, 0x9c, 0x8a, 0x74, 0x28, 0x88, 0x8b, 0xa7, 0xa},
			exp: "-AaaqLiV3k-cinQoiIunCg",
		},
		{
			u:   UUID{0x44, 0xca, 0x1a, 0x1f, 0xfa, 0x63, 0xfe, 0x42, 0xbf, 0xd, 0x72, 0x6, 0xce, 0xa4, 0x47, 0x65},
			exp: "RMoaH_pj_kK_DXIGzqRHZQ",
		},
	}

	for _, test := range tests {
		if test.u.Base64URL() != test.exp {
			t.Errorf("Incorrect URL-safe base64 representation of UUID: %s", test.exp)
		}
	}
}

func TestEqual(t *testing.T) {
	if !Equal(NamespaceDNS, NamespaceDNS) {
		t.Errorf("Incorrect comparison of %s and %s", NamespaceDNS, NamespaceDNS)
	}

	if Equal(NamespaceDNS, NamespaceURL) {
		t.Errorf("Incorrect comparison of %s and %s", NamespaceDNS, NamespaceURL)
	}
}

func TestOr(t *testing.T) {
	u1 := UUID{0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff}
	u2 := UUID{0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00}

	u := UUID{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	if !Equal(u, Or(u1, u2)) {
		t.Errorf("Incorrect bitwise OR result %s", Or(u1, u2))
	}
}

func TestAnd(t *testing.T) {
	u1 := UUID{0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff}
	u2 := UUID{0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00}

	u := UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	if !Equal(u, And(u1, u2)) {
		t.Errorf("Incorrect bitwise AND result %s", And(u1, u2))
	}
}

func TestVersion(t *testing.T) {
	u := UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	if u.Version() != 1 {
		t.Errorf("Incorrect version for UUID: %d", u.Version())
	}
}

func TestSetVersion(t *testing.T) {
	u := UUID{}
	u.SetVersion(4)

	if u.Version() != 4 {
		t.Errorf("Incorrect version for UUID after u.setVersion(4): %d", u.Version())
	}
}

func TestVariant(t *testing.T) {
	u1 := UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	if u1.Variant() != VariantNCS {
		t.Errorf("Incorrect variant for UUID variant %d: %d", VariantNCS, u1.Variant())
	}

	u2 := UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	if u2.Variant() != VariantRFC4122 {
		t.Errorf("Incorrect variant for UUID variant %d: %d", VariantRFC4122, u2.Variant())
	}

	u3 := UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	if u3.Variant() != VariantMicrosoft {
		t.Errorf("Incorrect variant for UUID variant %d: %d", VariantMicrosoft, u3.Variant())
	}

	u4 := UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	if u4.Variant() != VariantFuture {
		t.Errorf("Incorrect variant for UUID variant %d: %d", VariantFuture, u4.Variant())
	}
}

func TestSetVariant(t *testing.T) {
	u := new(UUID)
	u.SetVariant()

	if u.Variant() != VariantRFC4122 {
		t.Errorf("Incorrect variant for UUID after u.setVariant(): %d", u.Variant())
	}
}

func TestFromBytes(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	b1 := []byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	u1, err := FromBytes(b1)
	if err != nil {
		t.Errorf("Error parsing UUID from bytes: %s", err)
	}

	if !Equal(u, u1) {
		t.Errorf("UUIDs should be equal: %s and %s", u, u1)
	}

	b2 := []byte{}

	_, err = FromBytes(b2)
	if err == nil {
		t.Errorf("Should return error parsing from empty byte slice, got %s", err)
	}
}

func TestMarshalBinary(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	b1 := []byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	b2, err := u.MarshalBinary()
	if err != nil {
		t.Errorf("Error marshaling UUID: %s", err)
	}

	if !bytes.Equal(b1, b2) {
		t.Errorf("Marshaled UUID should be %s, got %s", b1, b2)
	}
}

func TestUnmarshalBinary(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	b1 := []byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	u1 := UUID{}
	err := u1.UnmarshalBinary(b1)
	if err != nil {
		t.Errorf("Error unmarshaling UUID: %s", err)
	}

	if !Equal(u, u1) {
		t.Errorf("UUIDs should be equal: %s and %s", u, u1)
	}

	b2 := []byte{}
	u2 := UUID{}

	err = u2.UnmarshalBinary(b2)
	if err == nil {
		t.Errorf("Should return error unmarshalling from empty byte slice, got %s", err)
	}
}

func TestFromString(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	s1 := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	s2 := "{6ba7b810-9dad-11d1-80b4-00c04fd430c8}"
	s3 := "urn:uuid:6ba7b810-9dad-11d1-80b4-00c04fd430c8"

	_, err := FromString("")
	if err == nil {
		t.Errorf("Should return error trying to parse empty string, got %s", err)
	}

	u1, err := FromString(s1)
	if err != nil {
		t.Errorf("Error parsing UUID from string: %s", err)
	}

	if !Equal(u, u1) {
		t.Errorf("UUIDs should be equal: %s and %s", u, u1)
	}

	u2, err := FromString(s2)
	if err != nil {
		t.Errorf("Error parsing UUID from string: %s", err)
	}

	if !Equal(u, u2) {
		t.Errorf("UUIDs should be equal: %s and %s", u, u2)
	}

	u3, err := FromString(s3)
	if err != nil {
		t.Errorf("Error parsing UUID from string: %s", err)
	}

	if !Equal(u, u3) {
		t.Errorf("UUIDs should be equal: %s and %s", u, u3)
	}
}

func TestFromStringShort(t *testing.T) {
	// Invalid 35-character UUID string
	s1 := "6ba7b810-9dad-11d1-80b4-00c04fd430c"

	for i := len(s1); i >= 0; i-- {
		// Skip potential base64 strings.
		if 22 <= i && i <= 24 {
			continue
		}
		_, err := FromString(s1[:i])
		if err == nil {
			t.Errorf("Should return error trying to parse too short string, got %s for len=%d", err, i)
		}
	}
}

func TestFromStringLong(t *testing.T) {
	// Invalid 37+ character UUID string
	s := []string{
		"6ba7b810-9dad-11d1-80b4-00c04fd430c8=",
		"6ba7b810-9dad-11d1-80b4-00c04fd430c8}",
		"{6ba7b810-9dad-11d1-80b4-00c04fd430c8}f",
		"6ba7b810-9dad-11d1-80b4-00c04fd430c800c04fd430c8",
	}

	for _, str := range s {
		_, err := FromString(str)
		if err == nil {
			t.Errorf("Should return error trying to parse too long string, passed %s", str)
		}
	}
}

func TestFromStringInvalid(t *testing.T) {
	// Invalid UUID string formats
	s := []string{
		"6ba7b8109dad11d180b400c04fd430c8",
		"6ba7b8109dad11d180b400c04fd430c86ba7b8109dad11d180b400c04fd430c8",
		"urn:uuid:{6ba7b810-9dad-11d1-80b4-00c04fd430c8}",
		"6ba7b8109-dad-11d1-80b4-00c04fd430c8",
		"6ba7b810-9dad1-1d1-80b4-00c04fd430c8",
		"6ba7b810-9dad-11d18-0b4-00c04fd430c8",
		"6ba7b810-9dad-11d1-80b40-0c04fd430c8",
		"6ba7b810+9dad+11d1+80b4+00c04fd430c8",
		"6ba7b810-9dad11d180b400c04fd430c8",
		"6ba7b8109dad-11d180b400c04fd430c8",
		"6ba7b8109dad11d1-80b400c04fd430c8",
		"6ba7b8109dad11d180b4-00c04fd430c8",
	}

	for _, str := range s {
		_, err := FromString(str)
		if err == nil {
			t.Errorf("Should return error trying to parse invalid string, passed %s", str)
		}
	}
}

func TestMust(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Should not panic on successful parse")
		}
	}()

	exp := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	u := Must(FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	if u != exp {
		t.Errorf("Should return parsed UUID, got %s", u)
	}
}

func TestMustInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Should panic on unsuccessful parse")
		}
	}()

	Must(FromString(""))
}

func TestMarshalText(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	b1 := []byte("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

	b2, err := u.MarshalText()
	if err != nil {
		t.Errorf("Error marshaling UUID: %s", err)
	}

	if !bytes.Equal(b1, b2) {
		t.Errorf("Marshaled UUID should be %s, got %s", b1, b2)
	}
}

func TestUnmarshalText(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	b1 := []byte("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

	u1 := UUID{}
	err := u1.UnmarshalText(b1)
	if err != nil {
		t.Errorf("Error unmarshaling UUID: %s", err)
	}

	if !Equal(u, u1) {
		t.Errorf("UUIDs should be equal: %s and %s", u, u1)
	}

	b2 := []byte("")
	u2 := UUID{}

	err = u2.UnmarshalText(b2)
	if err == nil {
		t.Errorf("Should return error trying to unmarshal from empty string")
	}
}

func TestValue(t *testing.T) {
	u, err := FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		t.Errorf("Error parsing UUID from string: %s", err)
	}

	val, err := u.Value()
	if err != nil {
		t.Errorf("Error getting UUID value: %s", err)
	}

	if !byteSliceEqual(val.([]byte), u.Bytes()) {
		t.Errorf("Wrong value returned, should be equal: %v and %v", val, u)
	}
}

func TestValueNil(t *testing.T) {
	u := UUID{}

	val, err := u.Value()
	if err != nil {
		t.Errorf("Error getting UUID value: %s", err)
	}

	if val != nil {
		t.Errorf("Wrong value returned, should be nil: %v", val)
	}
}

func TestScanBinary(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	b1 := []byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	u1 := UUID{}
	err := u1.Scan(b1)
	if err != nil {
		t.Errorf("Error unmarshaling UUID: %s", err)
	}

	if !Equal(u, u1) {
		t.Errorf("UUIDs should be equal: %s and %s", u, u1)
	}

	b2 := []byte{}
	u2 := UUID{}

	err = u2.Scan(b2)
	if err == nil {
		t.Errorf("Should return error unmarshalling from empty byte slice, got %s", err)
	}
}

func TestScanString(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	s1 := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"

	u1 := UUID{}
	err := u1.Scan(s1)
	if err != nil {
		t.Errorf("Error unmarshaling UUID: %s", err)
	}

	if !Equal(u, u1) {
		t.Errorf("UUIDs should be equal: %s and %s", u, u1)
	}

	s2 := ""
	u2 := UUID{}

	err = u2.Scan(s2)
	if err == nil {
		t.Errorf("Should return error trying to unmarshal from empty string")
	}
}

func TestScanText(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	b1 := []byte("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

	u1 := UUID{}
	err := u1.Scan(b1)
	if err != nil {
		t.Errorf("Error unmarshaling UUID: %s", err)
	}

	if !Equal(u, u1) {
		t.Errorf("UUIDs should be equal: %s and %s", u, u1)
	}

	b2 := []byte("")
	u2 := UUID{}

	err = u2.Scan(b2)
	if err == nil {
		t.Errorf("Should return error trying to unmarshal from empty string")
	}
}

func TestScanUnsupported(t *testing.T) {
	u := UUID{}

	err := u.Scan(true)
	if err == nil {
		t.Errorf("Should return error trying to unmarshal from bool")
	}
}

func TestScanNil(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	err := u.Scan(nil)
	if err != nil {
		t.Errorf("Error UUID should allow unmarshalling from nil")
	}

	if u != Nil {
		t.Errorf("Error UUID should unmarshalled from nil should be Nil")
	}
}

func TestIsNil(t *testing.T) {
	u := Nil
	if !u.IsNil() {
		t.Errorf("Error UUID should be detected as nil: %s", u.String())
	}
	u = NewV1()
	if u.IsNil() {
		t.Errorf("Error UUID should not be detected as nil: %s", u.String())
	}
}

func TestNewV1(t *testing.T) {
	u := NewV1()

	if u.Version() != 1 {
		t.Errorf("UUIDv1 generated with incorrect version: %d", u.Version())
	}

	if u.Variant() != VariantRFC4122 {
		t.Errorf("UUIDv1 generated with incorrect variant: %d", u.Variant())
	}

	u1 := NewV1()
	u2 := NewV1()

	if Equal(u1, u2) {
		t.Errorf("UUIDv1 generated two equal UUIDs: %s and %s", u1, u2)
	}

	oldFunc := epochFunc
	epochFunc = func() uint64 { return 0 }

	u3 := NewV1()
	u4 := NewV1()

	if Equal(u3, u4) {
		t.Errorf("UUIDv1 generated two equal UUIDs: %s and %s", u3, u4)
	}

	epochFunc = oldFunc
}

func TestNewV2(t *testing.T) {
	u1 := NewV2(DomainPerson)

	if u1.Version() != 2 {
		t.Errorf("UUIDv2 generated with incorrect version: %d", u1.Version())
	}

	if u1.Variant() != VariantRFC4122 {
		t.Errorf("UUIDv2 generated with incorrect variant: %d", u1.Variant())
	}

	u2 := NewV2(DomainGroup)

	if u2.Version() != 2 {
		t.Errorf("UUIDv2 generated with incorrect version: %d", u2.Version())
	}

	if u2.Variant() != VariantRFC4122 {
		t.Errorf("UUIDv2 generated with incorrect variant: %d", u2.Variant())
	}
}

func TestNewV3(t *testing.T) {
	u := NewV3(NamespaceDNS, "www.example.com")

	if u.Version() != 3 {
		t.Errorf("UUIDv3 generated with incorrect version: %d", u.Version())
	}

	if u.Variant() != VariantRFC4122 {
		t.Errorf("UUIDv3 generated with incorrect variant: %d", u.Variant())
	}

	if u.String() != "5df41881-3aed-3515-88a7-2f4a814cf09e" {
		t.Errorf("UUIDv3 generated incorrectly: %s", u.String())
	}

	u = NewV3(NamespaceDNS, "python.org")

	if u.String() != "6fa459ea-ee8a-3ca4-894e-db77e160355e" {
		t.Errorf("UUIDv3 generated incorrectly: %s", u.String())
	}

	u1 := NewV3(NamespaceDNS, "golang.org")
	u2 := NewV3(NamespaceDNS, "golang.org")
	if !Equal(u1, u2) {
		t.Errorf("UUIDv3 generated different UUIDs for same namespace and name: %s and %s", u1, u2)
	}

	u3 := NewV3(NamespaceDNS, "example.com")
	if Equal(u1, u3) {
		t.Errorf("UUIDv3 generated same UUIDs for different names in same namespace: %s and %s", u1, u2)
	}

	u4 := NewV3(NamespaceURL, "golang.org")
	if Equal(u1, u4) {
		t.Errorf("UUIDv3 generated same UUIDs for sane names in different namespaces: %s and %s", u1, u4)
	}
}

func TestNewV4(t *testing.T) {
	u := NewV4()

	if u.Version() != 4 {
		t.Errorf("UUIDv4 generated with incorrect version: %d", u.Version())
	}

	if u.Variant() != VariantRFC4122 {
		t.Errorf("UUIDv4 generated with incorrect variant: %d", u.Variant())
	}
}

func TestNewV5(t *testing.T) {
	u := NewV5(NamespaceDNS, "www.example.com")

	if u.Version() != 5 {
		t.Errorf("UUIDv5 generated with incorrect version: %d", u.Version())
	}

	if u.Variant() != VariantRFC4122 {
		t.Errorf("UUIDv5 generated with incorrect variant: %d", u.Variant())
	}

	u = NewV5(NamespaceDNS, "python.org")

	if u.String() != "886313e1-3b8a-5372-9b90-0c9aee199e5d" {
		t.Errorf("UUIDv5 generated incorrectly: %s", u.String())
	}

	u1 := NewV5(NamespaceDNS, "golang.org")
	u2 := NewV5(NamespaceDNS, "golang.org")
	if !Equal(u1, u2) {
		t.Errorf("UUIDv5 generated different UUIDs for same namespace and name: %s and %s", u1, u2)
	}

	u3 := NewV5(NamespaceDNS, "example.com")
	if Equal(u1, u3) {
		t.Errorf("UUIDv5 generated same UUIDs for different names in same namespace: %s and %s", u1, u2)
	}

	u4 := NewV5(NamespaceURL, "golang.org")
	if Equal(u1, u4) {
		t.Errorf("UUIDv3 generated same UUIDs for sane names in different namespaces: %s and %s", u1, u4)
	}
}

func byteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

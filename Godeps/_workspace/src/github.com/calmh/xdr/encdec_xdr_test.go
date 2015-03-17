// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package xdr_test

import (
	"bytes"
	"io"

	"github.com/calmh/xdr"
)

/*

TestStruct Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                              int                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             int8                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             uint8                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             int16                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|            0x0000             |             UI16              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             int32                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             UI32                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                         I64 (64 bits)                         +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                        UI64 (64 bits)                         +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of BS                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     BS (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                          Length of S                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                      S (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                            Opaque                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Number of SS                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of SS                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     SS (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct TestStruct {
	int I;
	int8 I8;
	uint8 UI8;
	int16 I16;
	unsigned int UI16;
	int32 I32;
	unsigned int UI32;
	hyper I64;
	unsigned hyper UI64;
	opaque BS<1024>;
	string S<1024>;
	Opaque C;
	string SS<1024>;
}

*/

func (o TestStruct) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o TestStruct) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o TestStruct) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o TestStruct) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o TestStruct) encodeXDR(xw *xdr.Writer) (int, error) {
	xw.WriteUint64(uint64(o.I))
	xw.WriteUint8(uint8(o.I8))
	xw.WriteUint8(o.UI8)
	xw.WriteUint16(uint16(o.I16))
	xw.WriteUint16(o.UI16)
	xw.WriteUint32(uint32(o.I32))
	xw.WriteUint32(o.UI32)
	xw.WriteUint64(uint64(o.I64))
	xw.WriteUint64(o.UI64)
	if l := len(o.BS); l > 1024 {
		return xw.Tot(), xdr.ElementSizeExceeded("BS", l, 1024)
	}
	xw.WriteBytes(o.BS)
	if l := len(o.S); l > 1024 {
		return xw.Tot(), xdr.ElementSizeExceeded("S", l, 1024)
	}
	xw.WriteString(o.S)
	_, err := o.C.encodeXDR(xw)
	if err != nil {
		return xw.Tot(), err
	}
	if l := len(o.SS); l > 1024 {
		return xw.Tot(), xdr.ElementSizeExceeded("SS", l, 1024)
	}
	xw.WriteUint32(uint32(len(o.SS)))
	for i := range o.SS {
		xw.WriteString(o.SS[i])
	}
	return xw.Tot(), xw.Error()
}

func (o *TestStruct) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *TestStruct) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *TestStruct) decodeXDR(xr *xdr.Reader) error {
	o.I = int(xr.ReadUint64())
	o.I8 = int8(xr.ReadUint8())
	o.UI8 = xr.ReadUint8()
	o.I16 = int16(xr.ReadUint16())
	o.UI16 = xr.ReadUint16()
	o.I32 = int32(xr.ReadUint32())
	o.UI32 = xr.ReadUint32()
	o.I64 = int64(xr.ReadUint64())
	o.UI64 = xr.ReadUint64()
	o.BS = xr.ReadBytesMax(1024)
	o.S = xr.ReadStringMax(1024)
	(&o.C).decodeXDR(xr)
	_SSSize := int(xr.ReadUint32())
	if _SSSize > 1024 {
		return xdr.ElementSizeExceeded("SS", _SSSize, 1024)
	}
	o.SS = make([]string, _SSSize)
	for i := range o.SS {
		o.SS[i] = xr.ReadString()
	}
	return xr.Error()
}
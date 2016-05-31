// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package db

import (
	"github.com/calmh/xdr"
)

/*

fileVersion Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                       Vector Structure                        \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                 device (length + padded data)                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct fileVersion {
	Vector version;
	opaque device<>;
}

*/

func (o fileVersion) XDRSize() int {
	return o.version.XDRSize() +
		4 + len(o.device) + xdr.Padding(len(o.device))
}

func (o fileVersion) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o fileVersion) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o fileVersion) MarshalXDRInto(m *xdr.Marshaller) error {
	if err := o.version.MarshalXDRInto(m); err != nil {
		return err
	}
	m.MarshalBytes(o.device)
	return m.Error
}

func (o *fileVersion) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *fileVersion) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	(&o.version).UnmarshalXDRFrom(u)
	o.device = u.UnmarshalBytes()
	return u.Error
}

/*

VersionList Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                      Number of versions                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\              Zero or more fileVersion Structures              \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct VersionList {
	fileVersion versions<>;
}

*/

func (o VersionList) XDRSize() int {
	return 4 + xdr.SizeOfSlice(o.versions)
}

func (o VersionList) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o VersionList) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o VersionList) MarshalXDRInto(m *xdr.Marshaller) error {
	m.MarshalUint32(uint32(len(o.versions)))
	for i := range o.versions {
		if err := o.versions[i].MarshalXDRInto(m); err != nil {
			return err
		}
	}
	return m.Error
}

func (o *VersionList) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *VersionList) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	_versionsSize := int(u.UnmarshalUint32())
	if _versionsSize < 0 {
		return xdr.ElementSizeExceeded("versions", _versionsSize, 0)
	} else if _versionsSize == 0 {
		o.versions = nil
	} else {
		if _versionsSize <= len(o.versions) {
			o.versions = o.versions[:_versionsSize]
		} else {
			o.versions = make([]fileVersion, _versionsSize)
		}
		for i := range o.versions {
			(&o.versions[i]).UnmarshalXDRFrom(u)
		}
	}
	return u.Error
}

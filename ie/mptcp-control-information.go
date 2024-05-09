// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewMPTCPControlInformation creates a new MPTCPControlInformation IE.
func NewMPTCPControlInformation(tci uint8) *IE {
	return newUint8ValIE(MPTCPControlInformation, tci&0x01)
}

// MPTCPControlInformation returns MPTCPControlInformation in uint8 if the type of IE matches.
func (i *IE) MPTCPControlInformation() (uint8, error) {
	switch i.Type {
	case MPTCPControlInformation:
		return i.ValueAsUint8()
	case ProvideATSSSControlInformation:
		ies, err := i.ProvideATSSSControlInformation()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == MPTCPControlInformation {
				return x.MPTCPControlInformation()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

// HasTCI reports whether an IE has TCI bit.
func (i *IE) HasTCI() bool {
	v, err := i.MPTCPControlInformation()
	if err != nil {
		return false
	}

	return has1stBit(v)
}

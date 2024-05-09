// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// NewInactivityDetectionTime creates a new InactivityDetectionTime IE.
func NewInactivityDetectionTime(threshold uint32) *IE {
	return newUint32ValIE(InactivityDetectionTime, threshold)
}

// InactivityDetectionTime returns InactivityDetectionTime in uint32 if the type of IE matches.
func (i *IE) InactivityDetectionTime() (uint32, error) {
	switch i.Type {
	case InactivityDetectionTime:
		return i.ValueAsUint32()
	case CreateURR:
		ies, err := i.CreateURR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == InactivityDetectionTime {
				return x.InactivityDetectionTime()
			}
		}
		return 0, ErrIENotFound
	case UpdateURR:
		ies, err := i.UpdateURR()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == InactivityDetectionTime {
				return x.InactivityDetectionTime()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}

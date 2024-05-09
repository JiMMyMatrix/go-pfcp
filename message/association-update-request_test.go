// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"net"
	"testing"
	"time"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestAssociationUpdateRequest(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Single IE",
			Structured: message.NewAssociationUpdateRequest(seq,
				ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
				ie.NewUPFunctionFeatures(0x01, 0x02),
				ie.NewCPFunctionFeatures(0x3f),
				ie.NewPFCPAssociationReleaseRequest(1, 1),
				ie.NewGracefulReleasePeriod(15*time.Minute),
				ie.NewPFCPAUReqFlags(0x01),
				ie.NewAlternativeSMFIPAddress(net.ParseIP("127.0.0.1"), net.ParseIP("2001::1")),
				ie.NewClockDriftControlInformation(
					ie.NewRequestedClockDriftInformation(1, 1),
					ie.NewTSNTimeDomainNumber(255),
					ie.NewTimeOffsetThreshold(10*time.Second),
					ie.NewCumulativeRateRatioThreshold(0xffffffff),
				),
				ie.NewUEIPAddressPoolInformation(
					ie.NewUEIPAddressPoolIdentity("go-pfcp"),
					ie.NewNetworkInstance("some.instance.example"),
				),
				ie.NewGTPUPathQoSControlInformation(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
					ie.NewGTPUPathInterfaceType(1, 1),
					ie.NewQoSReportTrigger(1, 1, 1),
					ie.NewTransportLevelMarking(0x1111),
					ie.NewMeasurementMethod(1, 1, 1),
					ie.NewAveragePacketDelay(10*time.Second),
					ie.NewMinimumPacketDelay(10*time.Second),
					ie.NewMaximumPacketDelay(10*time.Second),
					ie.NewTimer(20*time.Hour),
				),
				ie.NewUEIPAddressUsageInformation(
					ie.NewSequenceNumber(0xffffffff),
					ie.NewMetric(0x01),
					ie.NewValidityTimer(10*time.Second),
					ie.NewNumberOfUEIPAddresses(0x01, 4, 0),
					ie.NewNetworkInstance("some.instance.example"),
					ie.NewUEIPAddressPoolIdentity("go-pfcp"),
				),
			),
			Serialized: []byte{
				0x20, 0x07, 0x01, 0x41, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x3c, 0x00, 0x1d, 0x02, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x03, 0x65, 0x70, 0x63, 0x0b, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x03, 0x6f, 0x72, 0x67,
				0x00, 0x2b, 0x00, 0x02, 0x01, 0x02,
				0x00, 0x59, 0x00, 0x01, 0x3f,
				0x00, 0x6f, 0x00, 0x01, 0x03,
				0x00, 0x70, 0x00, 0x01, 0x2f,
				0x00, 0xa2, 0x00, 0x01, 0x01,
				0x00, 0xb2, 0x00, 0x15, 0x03, 0x7f, 0x00, 0x00, 0x01, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
				0x00, 0xcb, 0x00, 0x1e,
				0x00, 0xcc, 0x00, 0x01, 0x03,
				0x00, 0xce, 0x00, 0x01, 0xff,
				0x00, 0xcf, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02, 0x54, 0x0b, 0xe4, 0x00,
				0x00, 0xd0, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xe9, 0x00, 0x25,
				0x00, 0xb1, 0x00, 0x08, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
				0x00, 0x16, 0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xee, 0x00, 0x55,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xf1, 0x00, 0x01, 0x03,
				0x00, 0xed, 0x00, 0x01, 0x07,
				0x00, 0x1e, 0x00, 0x02, 0x11, 0x11,
				0x00, 0x3e, 0x00, 0x01, 0x07,
				0x00, 0xea, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xeb, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xec, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0x37, 0x00, 0x01, 0x82,
				0x01, 0x0b, 0x00, 0x41,
				0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x35, 0x00, 0x01, 0x01,
				0x01, 0x0d, 0x00, 0x02, 0x00, 0x0a,
				0x01, 0x0c, 0x00, 0x05, 0x01, 0x00, 0x00, 0x00, 0x04,
				0x00, 0x16, 0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xb1, 0x00, 0x08, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
			},
		}, {
			Description: "Multiple IEs",
			Structured: message.NewAssociationUpdateRequest(seq,
				ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
				ie.NewUPFunctionFeatures(0x01, 0x02),
				ie.NewCPFunctionFeatures(0x3f),
				ie.NewPFCPAssociationReleaseRequest(1, 1),
				ie.NewGracefulReleasePeriod(15*time.Minute),
				ie.NewPFCPAUReqFlags(0x01),
				ie.NewAlternativeSMFIPAddress(net.ParseIP("127.0.0.1"), net.ParseIP("2001::1")),
				ie.NewAlternativeSMFIPAddress(net.ParseIP("127.0.0.2"), net.ParseIP("2001::2")),
				ie.NewClockDriftControlInformation(
					ie.NewRequestedClockDriftInformation(1, 1),
					ie.NewTSNTimeDomainNumber(255),
					ie.NewTimeOffsetThreshold(10*time.Second),
					ie.NewCumulativeRateRatioThreshold(0xffffffff),
				),
				ie.NewClockDriftControlInformation(
					ie.NewRequestedClockDriftInformation(1, 1),
					ie.NewTSNTimeDomainNumber(1),
					ie.NewTimeOffsetThreshold(10*time.Second),
					ie.NewCumulativeRateRatioThreshold(0xffffffff),
				),
				ie.NewUEIPAddressPoolInformation(
					ie.NewUEIPAddressPoolIdentity("go-pfcp-1"),
					ie.NewNetworkInstance("some.instance-1.example"),
				),
				ie.NewUEIPAddressPoolInformation(
					ie.NewUEIPAddressPoolIdentity("go-pfcp-2"),
					ie.NewNetworkInstance("some.instance-2.example"),
				),
				ie.NewGTPUPathQoSControlInformation(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
					ie.NewGTPUPathInterfaceType(1, 1),
					ie.NewQoSReportTrigger(1, 1, 1),
					ie.NewTransportLevelMarking(0x1111),
					ie.NewMeasurementMethod(1, 1, 1),
					ie.NewAveragePacketDelay(10*time.Second),
					ie.NewMinimumPacketDelay(10*time.Second),
					ie.NewMaximumPacketDelay(10*time.Second),
					ie.NewTimer(20*time.Hour),
				),
				ie.NewGTPUPathQoSControlInformation(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.2", "", ie.DstInterfaceAccess, "some.instance.example"),
					ie.NewGTPUPathInterfaceType(1, 1),
					ie.NewQoSReportTrigger(1, 1, 1),
					ie.NewTransportLevelMarking(0x1111),
					ie.NewMeasurementMethod(1, 1, 1),
					ie.NewAveragePacketDelay(10*time.Second),
					ie.NewMinimumPacketDelay(10*time.Second),
					ie.NewMaximumPacketDelay(10*time.Second),
					ie.NewTimer(20*time.Hour),
				),
				ie.NewUEIPAddressUsageInformation(
					ie.NewSequenceNumber(0xffffffff),
					ie.NewMetric(0x01),
					ie.NewValidityTimer(10*time.Second),
					ie.NewNumberOfUEIPAddresses(0x01, 4, 0),
					ie.NewNetworkInstance("some.instance.example"),
					ie.NewUEIPAddressPoolIdentity("go-pfcp"),
				),
			),
			Serialized: []byte{
				0x20, 0x07, 0x02, 0x06, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x3c, 0x00, 0x1d, 0x02, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x03, 0x65, 0x70, 0x63, 0x0b, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x03, 0x6f, 0x72, 0x67,
				0x00, 0x2b, 0x00, 0x02, 0x01, 0x02,
				0x00, 0x59, 0x00, 0x01, 0x3f,
				0x00, 0x6f, 0x00, 0x01, 0x03,
				0x00, 0x70, 0x00, 0x01, 0x2f,
				0x00, 0xa2, 0x00, 0x01, 0x01,
				0x00, 0xb2, 0x00, 0x15, 0x03, 0x7f, 0x00, 0x00, 0x01, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
				0x00, 0xb2, 0x00, 0x15, 0x03, 0x7f, 0x00, 0x00, 0x02, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
				0x00, 0xcb, 0x00, 0x1e,
				0x00, 0xcc, 0x00, 0x01, 0x03,
				0x00, 0xce, 0x00, 0x01, 0xff,
				0x00, 0xcf, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02, 0x54, 0x0b, 0xe4, 0x00,
				0x00, 0xd0, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xcb, 0x00, 0x1e,
				0x00, 0xcc, 0x00, 0x01, 0x03,
				0x00, 0xce, 0x00, 0x01, 0x01,
				0x00, 0xcf, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02, 0x54, 0x0b, 0xe4, 0x00,
				0x00, 0xd0, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xe9, 0x00, 0x29,
				0x00, 0xb1, 0x00, 0x0a, 0x09, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x2d, 0x31,
				0x00, 0x16, 0x00, 0x17, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x31, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xe9, 0x00, 0x29,
				0x00, 0xb1, 0x00, 0x0a, 0x09, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x2d, 0x32,
				0x00, 0x16, 0x00, 0x17, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x32, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xee, 0x00, 0x55,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xf1, 0x00, 0x01, 0x03,
				0x00, 0xed, 0x00, 0x01, 0x07,
				0x00, 0x1e, 0x00, 0x02, 0x11, 0x11,
				0x00, 0x3e, 0x00, 0x01, 0x07,
				0x00, 0xea, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xeb, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xec, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0x37, 0x00, 0x01, 0x82,
				0x00, 0xee, 0x00, 0x55,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x02,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xf1, 0x00, 0x01, 0x03,
				0x00, 0xed, 0x00, 0x01, 0x07,
				0x00, 0x1e, 0x00, 0x02, 0x11, 0x11,
				0x00, 0x3e, 0x00, 0x01, 0x07,
				0x00, 0xea, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xeb, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xec, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0x37, 0x00, 0x01, 0x82,
				0x01, 0x0b, 0x00, 0x41,
				0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x35, 0x00, 0x01, 0x01,
				0x01, 0x0d, 0x00, 0x02, 0x00, 0x0a,
				0x01, 0x0c, 0x00, 0x05, 0x01, 0x00, 0x00, 0x00, 0x04,
				0x00, 0x16, 0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xb1, 0x00, 0x08, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseAssociationUpdateRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}

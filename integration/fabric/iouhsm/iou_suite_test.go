/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package iouhsm_test

import (
	"testing"

	"github.com/hyperledger-labs/fabric-smart-client/integration"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEndToEnd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IOU (With HSM) Suite")
}

func StartPort() int {
	return integration.IOUHSMPort.StartPortForNode()
}

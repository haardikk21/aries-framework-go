/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"errors"

	"github.com/hyperledger/aries-framework-go/pkg/crypto"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/dispatcher"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/transport"
	vdrapi "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
	"github.com/hyperledger/aries-framework-go/pkg/kms"
	"github.com/hyperledger/aries-framework-go/pkg/secretlock"
	"github.com/hyperledger/aries-framework-go/pkg/store/did"
	"github.com/hyperledger/aries-framework-go/pkg/store/verifiable"
	"github.com/hyperledger/aries-framework-go/spi/storage"
)

// ErrSvcNotFound is returned when service not found.
var ErrSvcNotFound = errors.New("service not found")

// Provider interface for protocol ctx.
type Provider interface {
	OutboundDispatcher() dispatcher.Outbound
	Messenger() service.Messenger
	Service(id string) (interface{}, error)
	StorageProvider() storage.Provider
	KMS() kms.KeyManager
	SecretLock() secretlock.Service
	Crypto() crypto.Crypto
	Packager() transport.Packager
	ServiceEndpoint() string
	RouterEndpoint() string
	VDRegistry() vdrapi.Registry
	ProtocolStateStorageProvider() storage.Provider
	InboundMessageHandler() transport.InboundMessageHandler
	OutboundMessageHandler() service.OutboundHandler
	VerifiableStore() verifiable.Store
	DIDConnectionStore() did.ConnectionStore
}

// ProtocolSvcCreator method to create new protocol service.
type ProtocolSvcCreator func(prv Provider) (dispatcher.ProtocolService, error)

// MessageServiceProvider is provider of message services.
type MessageServiceProvider interface {
	// Services returns list of available message services in this message handler
	Services() []dispatcher.MessageService
}

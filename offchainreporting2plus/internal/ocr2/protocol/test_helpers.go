package protocol

import (
	"github.com/GoPlugin/libocr/commontypes"
	"github.com/GoPlugin/libocr/offchainreporting2plus/types"
)

// Used only for testing
type XXXUnknownMessageType struct{}

// Conform to protocol.Message interface
func (XXXUnknownMessageType) CheckSize(types.ReportingPluginLimits) bool { return true }
func (XXXUnknownMessageType) process(*oracleState, commontypes.OracleID) {}

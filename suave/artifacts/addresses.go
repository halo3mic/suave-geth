// Code generated by suave/gen. DO NOT EDIT.
// Hash: 0ee1363822b0f9889af0eebd713a77f5c3b22d747334e087d470907ecc7580fe
package artifacts

import (
	"github.com/ethereum/go-ethereum/common"
)

// List of suave precompile addresses
var (
	buildEthBlockAddr         = common.HexToAddress("0x0000000000000000000000000000000042100001")
	buildEthBlockToAddr       = common.HexToAddress("0x0000000000000000000000000000000042100006")
	confidentialInputsAddr    = common.HexToAddress("0x0000000000000000000000000000000042010001")
	confidentialRetrieveAddr  = common.HexToAddress("0x0000000000000000000000000000000042020001")
	confidentialStoreAddr     = common.HexToAddress("0x0000000000000000000000000000000042020000")
	contextGetAddr            = common.HexToAddress("0x0000000000000000000000000000000053300003")
	doHTTPRequestAddr         = common.HexToAddress("0x0000000000000000000000000000000043200002")
	doHTTPRequestsAddr        = common.HexToAddress("0x0000000000000000000000000000000043200003")
	ethcallAddr               = common.HexToAddress("0x0000000000000000000000000000000042100003")
	extractHintAddr           = common.HexToAddress("0x0000000000000000000000000000000042100037")
	fetchDataRecordsAddr      = common.HexToAddress("0x0000000000000000000000000000000042030001")
	fillMevShareBundleAddr    = common.HexToAddress("0x0000000000000000000000000000000043200001")
	newBuilderAddr            = common.HexToAddress("0x0000000000000000000000000000000053200001")
	newDataRecordAddr         = common.HexToAddress("0x0000000000000000000000000000000042030000")
	privateKeyGenAddr         = common.HexToAddress("0x0000000000000000000000000000000053200003")
	randomBytesAddr           = common.HexToAddress("0x000000000000000000000000000000007770000b")
	signEthTransactionAddr    = common.HexToAddress("0x0000000000000000000000000000000040100001")
	signMessageAddr           = common.HexToAddress("0x0000000000000000000000000000000040100003")
	simulateBundleAddr        = common.HexToAddress("0x0000000000000000000000000000000042100000")
	simulateTransactionAddr   = common.HexToAddress("0x0000000000000000000000000000000053200002")
	submitBundleJsonRPCAddr   = common.HexToAddress("0x0000000000000000000000000000000043000001")
	submitEthBlockToRelayAddr = common.HexToAddress("0x0000000000000000000000000000000042100002")
)

var SuaveMethods = map[string]common.Address{
	"buildEthBlock":         buildEthBlockAddr,
	"buildEthBlockTo":       buildEthBlockToAddr,
	"confidentialInputs":    confidentialInputsAddr,
	"confidentialRetrieve":  confidentialRetrieveAddr,
	"confidentialStore":     confidentialStoreAddr,
	"contextGet":            contextGetAddr,
	"doHTTPRequest":         doHTTPRequestAddr,
	"doHTTPRequests":        doHTTPRequestsAddr,
	"ethcall":               ethcallAddr,
	"extractHint":           extractHintAddr,
	"fetchDataRecords":      fetchDataRecordsAddr,
	"fillMevShareBundle":    fillMevShareBundleAddr,
	"newBuilder":            newBuilderAddr,
	"newDataRecord":         newDataRecordAddr,
	"privateKeyGen":         privateKeyGenAddr,
	"randomBytes":           randomBytesAddr,
	"signEthTransaction":    signEthTransactionAddr,
	"signMessage":           signMessageAddr,
	"simulateBundle":        simulateBundleAddr,
	"simulateTransaction":   simulateTransactionAddr,
	"submitBundleJsonRPC":   submitBundleJsonRPCAddr,
	"submitEthBlockToRelay": submitEthBlockToRelayAddr,
}

func PrecompileAddressToName(addr common.Address) string {
	switch addr {
	case buildEthBlockAddr:
		return "buildEthBlock"
	case buildEthBlockToAddr:
		return "buildEthBlockTo"
	case confidentialInputsAddr:
		return "confidentialInputs"
	case confidentialRetrieveAddr:
		return "confidentialRetrieve"
	case confidentialStoreAddr:
		return "confidentialStore"
	case contextGetAddr:
		return "contextGet"
	case doHTTPRequestAddr:
		return "doHTTPRequest"
	case doHTTPRequestsAddr:
		return "doHTTPRequests"
	case ethcallAddr:
		return "ethcall"
	case extractHintAddr:
		return "extractHint"
	case fetchDataRecordsAddr:
		return "fetchDataRecords"
	case fillMevShareBundleAddr:
		return "fillMevShareBundle"
	case newBuilderAddr:
		return "newBuilder"
	case newDataRecordAddr:
		return "newDataRecord"
	case privateKeyGenAddr:
		return "privateKeyGen"
	case randomBytesAddr:
		return "randomBytes"
	case signEthTransactionAddr:
		return "signEthTransaction"
	case signMessageAddr:
		return "signMessage"
	case simulateBundleAddr:
		return "simulateBundle"
	case simulateTransactionAddr:
		return "simulateTransaction"
	case submitBundleJsonRPCAddr:
		return "submitBundleJsonRPC"
	case submitEthBlockToRelayAddr:
		return "submitEthBlockToRelay"
	}
	return ""
}

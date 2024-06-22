package suave

type Config struct {
	SuaveEthRemoteBackendEndpoint string // deprecated
	RedisStorePubsubUri           string
	RedisStoreUri                 string
	PebbleDbPath                  string
	EthBundleSigningKeyHex        string
	EthBlockSigningKeyHex         string
	ExternalWhitelist             []string
	AliasRegistry                 map[string]string
	BoostRelayUrl                 string
	BeaconRpc                     string
}

var DefaultConfig = Config{}

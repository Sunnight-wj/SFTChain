module github.com/CosmosContracts/juno/v17

go 1.20

require (
	cosmossdk.io/api v0.3.1
	cosmossdk.io/errors v1.0.0
	cosmossdk.io/log v1.1.1-0.20230704160919-88f2c830b0ca
	cosmossdk.io/math v1.0.1
	cosmossdk.io/tools/rosetta v0.2.1
	github.com/CosmWasm/wasmd v0.41.0
	github.com/CosmWasm/wasmvm v1.3.0
	github.com/cometbft/cometbft v0.37.2
	github.com/cometbft/cometbft-db v0.8.0
	github.com/cosmos/cosmos-sdk v0.47.4
	github.com/cosmos/gogoproto v1.4.10
	github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7 v7.0.0
	github.com/cosmos/ibc-apps/modules/async-icq/v7 v7.0.0
	github.com/cosmos/ibc-apps/modules/ibc-hooks/v7 v7.0.0-20230803181732-7c8f814d3b79
	github.com/cosmos/ibc-go/v7 v7.2.0
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/prometheus/client_golang v1.16.0
	github.com/skip-mev/pob v1.0.4
	github.com/spf13/cast v1.5.1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	google.golang.org/genproto/googleapis/api v0.0.0-20230629202037-9506855d4529
	google.golang.org/grpc v1.57.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/gookit/goutil v0.6.12 // indirect
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
)

require (
	github.com/cosmos/cosmos-proto v1.0.0-beta.3
	nhooyr.io/websocket v1.8.7 // indirect
)

replace (
	// cosmos keyring
	github.com/99designs/keyring => github.com/cosmos/keyring v1.2.0

	// support concurrency for iavl
	github.com/cosmos/iavl => github.com/cosmos/iavl v0.20.1

	// dgrijalva/jwt-go is deprecated and doesn't receive security updates.
	// TODO: remove it: https://github.com/cosmos/cosmos-sdk/issues/13134
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.4.2

	// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
	// TODO Remove it: https://github.com/cosmos/cosmos-sdk/issues/10409
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.9.0

	// https://github.com/cosmos/cosmos-sdk/issues/14949
	// pin the version of goleveldb to v1.0.1-0.20210819022825-2ae1ddf74ef7 required by SDK v47 upgrade guide.
	github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
)

// replace github.com/skip-mev/pob => ../pob

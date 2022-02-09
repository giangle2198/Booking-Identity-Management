module booking-identity-management

go 1.15

replace booking-libs => github.com/giangle2198/Booking-Libs v0.0.0-20211124073342-e18494b8e693

require (
	booking-libs v0.0.0-20211124073342-e18494b8e693
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	// github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jinzhu/copier v0.2.9
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/o1egl/paseto v1.0.0
	github.com/rakyll/statik v0.1.7
	github.com/rs/cors v1.8.0
	github.com/sarulabs/di v2.0.0+incompatible
	github.com/spf13/viper v1.9.0
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	google.golang.org/genproto v0.0.0-20210828152312-66f60bf46e71
	google.golang.org/grpc v1.40.0
)

module booking-identity-management

go 1.15

replace booking-libs => github.com/giangle2198/Booking-Libs v0.0.0-20211124073342-e18494b8e693

require (
	booking-libs v0.0.0-20211124073342-e18494b8e693
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	// github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jinzhu/copier v0.2.9
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/o1egl/paseto v1.0.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rakyll/statik v0.1.7
	github.com/rs/cors v1.8.0
	github.com/sarulabs/di v2.0.0+incompatible
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.1 // indirect
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.43.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
)

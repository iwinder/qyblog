module github.com/iwinder/qingyucms

go 1.18

require (
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20220811033210-f0c2a6ed90d1
	github.com/go-kratos/kratos/v2 v2.4.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/google/wire v0.5.0
	github.com/speps/go-hashids v2.0.0+incompatible
	go.etcd.io/etcd/client/v3 v3.5.4
	go.opentelemetry.io/otel v1.9.0
	go.opentelemetry.io/otel/exporters/jaeger v1.9.0
	go.opentelemetry.io/otel/sdk v1.9.0
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	google.golang.org/genproto v0.0.0-20220524023933-508584e28198
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.etcd.io/etcd/api/v3 v3.5.4 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.4 // indirect
	go.opentelemetry.io/otel/trace v1.9.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	golang.org/x/sync v0.0.0-20220513210516-0976fa681c29 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)

replace (
	golang.org/x/net => golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2
	golang.org/x/sync => golang.org/x/sync v0.0.0-20220513210516-0976fa681c29
	golang.org/x/sys => golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a
	golang.org/x/text => github.com/golang/text v0.3.7
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190422233926-fe54fb35175b // indirect
)

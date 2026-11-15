module greeter

go 1.25.7

require (
	go.putnami.dev/app v0.5.0-20261115101940-aa1341ca
	go.putnami.dev/config v0.5.0-20261115101940-aa1341ca
	go.putnami.dev/http v0.5.0-20261115101940-aa1341ca
	go.putnami.dev/logger v0.5.0-20261115101940-aa1341ca
	go.putnami.dev/platform v0.5.0-20261115101940-aa1341ca
)

require (
	github.com/putnami/qa-fixture-app/libs/greeting v0.0.0
	go.putnami.dev/errors v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/inject v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/migration v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/capabilities v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/config v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/database v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/diagnostic v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/events v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/features v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/http-routes v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/identity v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/infra v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/migration v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/platform v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/runtime v0.5.0-20261115101940-aa1341ca // indirect
	go.putnami.dev/protocol/storage v0.5.0-20261115101940-aa1341ca // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/putnami/qa-fixture-app/libs/greeting => ../../libs/greeting

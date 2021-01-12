module github.com/mimetrix/pfcp

go 1.15

replace github.com/free5gc/tlv => ./tlv

require (
	github.com/free5gc/tlv v0.0.0-00010101000000-000000000000
	github.com/kr/text v0.2.0 // indirect
	github.com/rs/zerolog v1.20.0
	github.com/stretchr/testify v1.6.1
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

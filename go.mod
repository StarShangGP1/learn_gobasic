module learn_gobasic

go 1.17

require (
	github.com/armstrongli/go-bmi v0.0.0-00010101000000-000000000000
	github.com/ghodss/yaml v1.0.0
	github.com/spf13/cobra v1.3.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/armstrongli/go-bmi => ./homework/class09/homework01/staging/src/github.com/armstrongli/go-bmi
	github.com/spf13/cobra => github.com/armstrongli/cobra v1.2.0
)

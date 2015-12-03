package servicebus

import (
	"github.com/ma314smith/goazure"
	"github.com/spf13/viper"
)

//API wraps a ServiceBusRelay and provides methods for retrieving information from the servicebus.
type API struct {
	Relay goazure.ServiceBusRelay
}

//NewAPI returns a valid API struct with a ServiceBusRelay configured from environmental variables.
func NewAPI() API {
	env := viper.New()
	env.SetEnvPrefix("GOAZURE")
	env.AutomaticEnv()

	issuerName := env.GetString("ACSISSUERNAME")
	issuerKey := env.GetString("ACSISSUERKEY")

	acs := goazure.ACS{IssuerName: issuerName, IssuerKey: issuerKey}
	//TODO: Update the namespace and scope to either be pulled from the environment or added as arguments to the NewAPI function.
	sbr := goazure.ServiceBusRelay{Namespace: "blackbaud", Scope: "servicerouter-qas", AccessControl: &acs}
	return API{Relay: sbr}
}

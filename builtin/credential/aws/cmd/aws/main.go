package main

import (
	"os"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/api"
	awsauth "github.com/hashicorp/vault/builtin/credential/aws"
	"github.com/hashicorp/vault/sdk/plugin"
)

func main() {
	apiClientMeta := &api.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	if err := plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: awsauth.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	}); err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}

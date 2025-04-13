package initialize

import (
	"os"
	"server/global"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

func ConnectES() *elasticsearch.TypedClient {
	esConfig := global.Config.ES


	cfg := elasticsearch.Config{
		Addresses: []string{esConfig.URL},
		Username:  esConfig.Username,
		Password:  esConfig.Password,
	}

	if esConfig.IsConsolePrint{
		cfg.Logger = &elastictransport.ColorLogger{
			Output: os.Stdout,
			EnableRequestBody: true,
			EnableResponseBody: true,
		}
	}
	
	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		global.Log.Error("ES杩炴帴澶辫触", zap.Error(err))
		os.Exit(1)
	}
	return client
}

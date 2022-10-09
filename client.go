package starter

import (
	"strings"

	"github.com/goexl/ft"
	"github.com/pangum/http"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
)

// Client 客户端
type Client struct {
	*ft.Client
}

func newClient(config *pangu.Config, http *http.Client, logger *logging.Logger) (client *Client, err error) {
	_config := new(panguConfig)
	if err = config.Load(_config); nil != err {
		return
	}

	client = new(Client)
	opts := ft.NewNewOptions(ft.Http(http.Client), ft.Logger(logger.Logger))
	if `` != strings.TrimSpace(_config.Ft.Iv) {
		opts = append(opts, ft.Iv(_config.Ft.Iv))
	}
	if `` != strings.TrimSpace(_config.Ft.Addr) {
		opts = append(opts, ft.Addr(_config.Ft.Addr))
	}
	client.Client, err = ft.New(opts...)

	return
}

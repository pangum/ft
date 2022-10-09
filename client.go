package ft

import (
	"github.com/goexl/ft"
	"github.com/pangum/http"
	"github.com/pangum/logging"
)

// Client 客户端
type Client struct {
	*ft.Client
}

func newClient(http *http.Client, logger *logging.Logger) (client *Client, err error) {
	client = new(Client)
	client.Client, err = ft.New(ft.Http(http.Client), ft.Logger(logger.Logger))

	return
}

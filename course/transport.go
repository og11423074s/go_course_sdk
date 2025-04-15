package course

import (
	"fmt"
	c "github.com/ncostamagna/go_http_client/client"
	"github.com/og11423074s/gocourse_domain/domain"
	"net/http"
	"net/url"
	"time"
)

type (
	DataResponse struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}

	Transport interface {
		Get(id string) (*domain.Course, error)
	}

	clientHTTP struct {
		client c.Transport
	}
)

func NeHttpClient(baseURL string, token string) Transport {
	header := http.Header{}

	if token != "" {
		header.Set("Authorization", token)
	}

	return &clientHTTP{
		client: c.New(header, baseURL, 5000*time.Millisecond, true),
	}
}

func (c *clientHTTP) Get(id string) (*domain.Course, error) {

	dataResponse := DataResponse{Data: &domain.Course{}}

	u := url.URL{}
	u.Path += fmt.Sprintf("courses/%s", id)
	reps := c.client.Get(u.String())

	if reps.Err != nil {
		return nil, reps.Err
	}

	if reps.StatusCode == 404 {
		return nil, ErrNotFound{fmt.Sprintf("%s", reps)}
	}

	if reps.StatusCode > 299 {
		return nil, fmt.Errorf("error: %s", reps)
	}

	if err := reps.FillUp(&dataResponse); err != nil {
		return nil, err
	}

	return dataResponse.Data.(*domain.Course), nil
}

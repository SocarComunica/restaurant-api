package restful

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	platformerrors "restaurant-api/internal/platform/error"
)

var (
	errorRestClientGet = errors.New("couldn't execute request get for URL")
)

type RestClient interface {
	Get(ctx *gin.Context, url string, header *http.Header) (*http.Response, error)
}

type restClient struct {
	baseURL    string
	httpClient http.Client
}

func (rc *restClient) Get(ctx *gin.Context, url string, header *http.Header) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, rc.baseURL+url, nil)
	if err != nil {
		return nil, platformerrors.NewInternalServerAPIError(errorRestClientGet.Error(), err)
	}

	if header != nil {
		request.Header = *header
	}

	return rc.ExecuteRequest(request)
}

func (rc *restClient) ExecuteRequest(request *http.Request) (*http.Response, error) {
	response, err := rc.httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewRestClient(baseURL string) RestClient {
	return &restClient{baseURL: baseURL, httpClient: http.Client{}}
}

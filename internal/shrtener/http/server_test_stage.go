package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

const applicationJSONContentType string = "application/json"

type serverStage struct {
	suite.Suite
	t        *testing.T
	host     string
	http     *http.Client
	request  *http.Request
	response *http.Response
	body     *handlers.UrlMetadata
	requests []*handlers.UrlMetadata
	sUrl     string
}

func (s *serverStage) and() *serverStage {
	return s
}

func newServerStage(t *testing.T) (*serverStage, *serverStage, *serverStage) {
	host := viper.Get("server.host")
	port := viper.Get("server.port")
	as := &serverStage{
		t:    t,
		host: fmt.Sprintf("http://%s:%d", host, port),
		http: http.DefaultClient,
	}

	return as, as, as
}

func (s *serverStage) aListRequestIsPrepared() *serverStage {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/urls", s.host), nil)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.request = r

	return s
}

func (s *serverStage) listEndpointIsQueriedWithSuccess() *serverStage {
	r, err := s.http.Do(s.request)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.response = r

	return s
}

func (s *serverStage) responseShouldReturnStatusCode(statusCode int) *serverStage {
	require.Equal(s.t, statusCode, s.response.StatusCode)
	return s
}

func (s *serverStage) shouldBeEmptyList() *serverStage {
	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.t, err)

	r := new(handlers.ListResponse)
	err = json.Unmarshal(body, r)
	require.Nil(s.t, err)
	require.Empty(s.t, r.Data)

	return s
}

func (s *serverStage) shouldBeListWithItems() *serverStage {
	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.t, err)

	r := new(handlers.ListResponse)
	err = json.Unmarshal(body, r)
	require.Nil(s.t, err)
	require.NotEmpty(s.t, r.Data)

	return s
}

func (s *serverStage) aCreateRequestIsPrepared(url string) *serverStage {
	s.body = &handlers.UrlMetadata{
		Original: url,
	}

	return s
}

func (s *serverStage) createEndpointIsCalled() *serverStage {
	// TODO: extract Do(s.request) and asserts to a single method
	payload, err := json.Marshal(s.body)
	require.Nil(s.t, err)
	r, err := s.http.Post(fmt.Sprintf("%s/urls", s.host), applicationJSONContentType, bytes.NewBuffer(payload))
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.response = r

	return s
}

func (s *serverStage) responseBodyShouldNotBeEmpty() *serverStage {
	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.t, err)

	r := new(handlers.UrlMetadata)
	err = json.Unmarshal(body, r)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)
	require.NotNil(s.t, r.Short)
	require.NotEqual(s.t, s.body.Original, r.Short)
	require.NotNil(s.t, r.Original)
	require.Equal(s.t, s.body.Original, r.Original)
	require.NotNil(s.t, r.ExpirationDate)
	require.NotNil(s.t, r.DateCreated)

	return s
}

func (s *serverStage) responseBodyShouldReturnEmptyUrlError() *serverStage {
	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.t, err)

	r := new(handlers.HttpError)
	err = json.Unmarshal(body, r)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)
	require.NotNil(s.t, r.Message)
	require.NotNil(s.t, r.Details)
	require.NotNil(s.t, r.Code)

	return s
}

func (s *serverStage) dbIsEmpty() {

}

func (s *serverStage) twoRequestsWithSameUrlAreCreated() *serverStage {
	url := uuid.New().String()
	md := &handlers.UrlMetadata{
		Original: url,
	}

	s.requests = append(s.requests, md)
	s.requests = append(s.requests, md)

	return s
}

func (s *serverStage) createEndpointIsCalledWithRequestsWithSameUrl() *serverStage {
	// TODO: extract Do(s.request) and asserts to a single method
	first, err := json.Marshal(s.requests[0])
	require.Nil(s.t, err)
	r, err := s.http.Post(fmt.Sprintf("%s/urls", s.host), applicationJSONContentType, bytes.NewBuffer(first))
	require.Nil(s.t, err)
	require.NotNil(s.t, r)
	require.Equal(s.t, http.StatusCreated, r.StatusCode)

	second, err := json.Marshal(s.requests[1])
	require.Nil(s.t, err)
	r2, err := s.http.Post(fmt.Sprintf("%s/urls", s.host), applicationJSONContentType, bytes.NewBuffer(second))
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.response = r2

	return s
}

func (s *serverStage) aUrlIsShortened(url string) *serverStage {
	s.aCreateRequestIsPrepared(url).
		and().
		createEndpointIsCalled()

	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.t, err)

	resp := &handlers.UrlMetadata{}
	err = json.Unmarshal(body, resp)
	require.Nil(s.t, err)
	require.NotNil(s.t, resp)

	s.sUrl = resp.Short

	return s
}

func (s *serverStage) aRedirectRequestIsCreated() *serverStage {

	r, err := http.NewRequest("GET", s.sUrl, nil)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.request = r

	return s
}

func (s *serverStage) redirectIsRequested() *serverStage {
	r, err := s.http.Do(s.request)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.response = r

	return s
}

func (s *serverStage) aRedirectRequestIsCreatedWithRandomUrl() *serverStage {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", s.host, uuid.New().String()), nil)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.request = r

	return s
}

func (s *serverStage) aNonShortenedUrlIsRequested() *serverStage {
	r, err := s.http.Do(s.request)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.response = r

	return s
}

func (s *serverStage) shouldHaveNotFoundErrorMessage() *serverStage {
	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.t, err)

	resp := &handlers.HttpError{}
	err = json.Unmarshal(body, resp)
	require.Nil(s.t, err)
	require.NotNil(s.t, resp)

	return s
}

package http

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	// TODO: post request

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

	r := new(handlers.List)
	err = json.Unmarshal(body, r)
	require.Nil(s.t, err)
	require.Empty(s.t, r.Data)

	return s
}

func (s *serverStage) shouldBeListWithItems() {
	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.t, err)

	r := new(handlers.List)
	err = json.Unmarshal(body, r)
	require.Nil(s.t, err)
	require.Empty(s.t, r.Data)
}

func (s *serverStage) aCreateRequestIsPrepared(url string) *serverStage {
	s.body = &handlers.UrlMetadata{
		Original: url,
	}

	return s
}

func (s *serverStage) createEndpointIsCalledWithSuccess() *serverStage {
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
	require.NotNil(s.t, r.Original)
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

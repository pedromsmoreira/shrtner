package http

import (
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

type serverStage struct {
	suite.Suite
	t        *testing.T
	host     string
	http     *http.Client
	request  *http.Request
	response *http.Response
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
	// post request

	r, err := http.NewRequest("GET", "/urls", nil)
	require.Nil(s.T(), err)
	require.NotNil(s.T(), r)

	s.request = r

	return s
}

func (s *serverStage) listEndpointIsQueriedWithSuccess() *serverStage {
	r, err := s.http.Do(s.request)
	require.Nil(s.T(), err)
	require.NotNil(s.T(), r)

	s.response = r

	return s
}

func (s *serverStage) listResponseShouldReturnStatusCode(statusCode int) *serverStage {
	require.Equal(s.t, statusCode, s.response.StatusCode)
	return s
}

func (s *serverStage) shouldBeEmptyList() *serverStage {
	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.T(), err)

	r := new(handlers.List)
	err = json.Unmarshal(body, r)
	require.Nil(s.T(), err)
	require.Empty(s.T(), r.Data)

	return s
}

func (s *serverStage) shouldBeListWithItems() {
	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.T(), err)

	r := new(handlers.List)
	err = json.Unmarshal(body, r)
	require.Nil(s.T(), err)
	require.Empty(s.T(), r.Data)
}

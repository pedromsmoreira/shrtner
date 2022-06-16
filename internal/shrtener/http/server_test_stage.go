package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
		http: &http.Client{},
	}

	return as, as, as
}

func (s *serverStage) aListRequestIsPrepared() *serverStage {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/urls?page=0&size=5", s.host), nil)
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
	require.Equal(s.t, 5, len(r.Data))

	return s
}

func (s *serverStage) aCreateRequestIsPrepared(url string) *serverStage {
	s.body = &handlers.UrlMetadata{
		Original: url,
	}

	return s
}

func (s *serverStage) aCreateRequestWithExpirationDateIsPrepared(url string, expirationDate time.Time) *serverStage {
	ed := expirationDate.Format(time.RFC3339Nano)
	s.body = &handlers.UrlMetadata{
		Original:       url,
		ExpirationDate: ed,
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

	r := new(handlers.BadRequestError)
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
	u := uuid.New().String()
	md := &handlers.UrlMetadata{
		Original: u,
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

func (s *serverStage) aUrlIsShortenedWithCustomExpirationDate(url string, expirationDate time.Time) *serverStage {
	s.aCreateRequestWithExpirationDateIsPrepared(url, expirationDate).
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

func (s *serverStage) aRedirectRequestIsCreatedWithCustomShortUrl(url string) *serverStage {

	r, err := http.NewRequest("GET", url, nil)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.request = r

	return s
}

func (s *serverStage) redirectIsRequested() *serverStage {
	s.http.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
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

func (s *serverStage) anExpiredShortenedUrlIsRequested() *serverStage {
	r, err := s.http.Do(s.request)
	require.Nil(s.t, err)
	require.NotNil(s.t, r)

	s.response = r

	return s
}

func (s *serverStage) shouldHaveNotFoundErrorMessage() *serverStage {
	body, err := ioutil.ReadAll(s.response.Body)
	require.Nil(s.t, err)

	resp := &handlers.NotFoundError{}
	err = json.Unmarshal(body, resp)
	require.Nil(s.t, err)
	require.NotNil(s.t, resp)

	return s
}

func (s *serverStage) deleteUrlFromDb() {
	u, err := url.Parse(s.sUrl)
	require.Nil(s.t, err)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/urls%s", s.host, u.Path), nil)
	require.Nil(s.t, err)

	resp, err := s.http.Do(req)
	require.Nil(s.t, err)
	require.NotNil(s.t, resp)
	require.Equal(s.t, http.StatusNoContent, resp.StatusCode)
}

func (s *serverStage) waitFor(waitTime time.Duration) {
	time.Sleep(waitTime)
}

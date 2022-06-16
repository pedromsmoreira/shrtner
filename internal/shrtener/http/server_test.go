package http

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestList(t *testing.T) {
	t.Run("when request has success returns empty list", func(t *testing.T) {
		t.Skip("needs multi tenancy to work")
		given, when, then := newServerStage(t)

		given.aListRequestIsPrepared().and().
			dbIsEmpty()

		when.listEndpointIsQueriedWithSuccess()

		then.
			responseShouldReturnStatusCode(http.StatusOK).
			and().
			shouldBeEmptyList()
	})

	t.Run("when request has success returns list with items", func(t *testing.T) {
		given, when, then := newServerStage(t)

		given.aCreateRequestIsPrepared(fmt.Sprintf("https://%s.com", uuid.NewString())).and().
			createEndpointIsCalled().and().
			aListRequestIsPrepared()

		when.listEndpointIsQueriedWithSuccess()

		then.
			responseShouldReturnStatusCode(http.StatusOK).
			and().
			shouldBeListWithItems()
	})
}

func TestCreate(t *testing.T) {
	t.Run("when request has success returns created item", func(t *testing.T) {
		given, when, then := newServerStage(t)

		given.aCreateRequestIsPrepared(fmt.Sprintf("https://%s.com/", uuid.New().String()))

		when.createEndpointIsCalled()

		then.
			responseShouldReturnStatusCode(http.StatusCreated).
			and().
			responseBodyShouldNotBeEmpty()
	})

	t.Run("when request has empty url returns 400 bad request", func(t *testing.T) {
		given, when, then := newServerStage(t)

		given.aCreateRequestIsPrepared("")

		when.createEndpointIsCalled()

		then.
			responseShouldReturnStatusCode(http.StatusBadRequest).
			and().
			responseBodyShouldReturnEmptyUrlError()
	})

	t.Run("when an existing url did not expire and is shortened again, should return conflict", func(t *testing.T) {
		given, when, then := newServerStage(t)

		given.twoRequestsWithSameUrlAreCreated()

		when.createEndpointIsCalledWithRequestsWithSameUrl()

		then.
			responseShouldReturnStatusCode(http.StatusConflict)
	})
}

func TestRedirect(t *testing.T) {
	t.Run("when shortened google url is requested should return found (Status Code 302) ", func(t *testing.T) {
		given, when, then := newServerStage(t)

		given.aUrlIsShortened("https://www.google.com").
			and().
			aRedirectRequestIsCreated().and().
			waitFor(200 * time.Millisecond)

		when.redirectIsRequested()

		then.
			responseShouldReturnStatusCode(http.StatusFound).and().
			deleteUrlFromDb()
	})

	t.Run("when shortened urls has expired should return 404 Not Found", func(t *testing.T) {
		given, when, then := newServerStage(t)

		expDate := time.Now().UTC().Add(200 * time.Millisecond)
		given.aUrlIsShortenedWithCustomExpirationDate(fmt.Sprintf("https://www.%s.com", uuid.NewString()), expDate).
			and().
			aRedirectRequestIsCreated().and().
			waitFor(500 * time.Millisecond)

		when.redirectIsRequested()

		then.
			responseShouldReturnStatusCode(http.StatusNotFound).and().
			shouldHaveNotFoundErrorMessage()
	})

	t.Run("when shortened url does not exist should return 404 Not Found", func(t *testing.T) {
		given, when, then := newServerStage(t)

		given.aRedirectRequestIsCreatedWithCustomShortUrl("http://localhost:5000/" + uuid.NewString())

		when.redirectIsRequested()

		then.
			responseShouldReturnStatusCode(http.StatusNotFound).and().
			shouldHaveNotFoundErrorMessage()
	})
}

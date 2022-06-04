package http

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"testing"
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

		given.aListRequestIsPrepared()

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

		when.createEndpointIsCalledWithSuccess()

		then.
			responseShouldReturnStatusCode(http.StatusCreated).
			and().
			responseBodyShouldNotBeEmpty()
	})

	t.Run("when request has empty url returns 400 bad request", func(t *testing.T) {
		given, when, then := newServerStage(t)

		given.aCreateRequestIsPrepared("")

		when.createEndpointIsCalledWithSuccess()

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

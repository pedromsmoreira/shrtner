package http

import (
	"net/http"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("when request has success returns empty list", func(t *testing.T) {
		given, when, then := newServerStage(t)

		given.aListRequestIsPrepared()

		when.listEndpointIsQueriedWithSuccess()

		then.
			listResponseShouldReturnStatusCode(http.StatusOK).
			and().
			shouldBeEmptyList()
	})

	t.Run("when request has success returns list with items", func(t *testing.T) {
		given, when, then := newServerStage(t)

		given.aListRequestIsPrepared()

		when.listEndpointIsQueriedWithSuccess()

		then.
			listResponseShouldReturnStatusCode(http.StatusOK).
			and().
			shouldBeListWithItems()
	})
}

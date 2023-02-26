package tgclient

import (
	"sync/atomic"
)

func handleGetMeResponse(response *GetMeResponse) {
	logGetMe(response)
}

func handleUpdateResponse(lastUpdateId *int64, response *GetUpdatesResponse) {
	update := response.Result[0]

	if *lastUpdateId == update.UpdateId {
		return
	}

	atomic.StoreInt64(lastUpdateId, update.UpdateId)
	logUpdate(&update)
}

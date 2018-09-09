package epochtime

import (
	"encoding/json"
	"testing"
)

type Tweet struct {
	Handle string
	Time   EpochTime
	Status string
}

func TestEpochTimeJSON(t *testing.T) {
	var obj Tweet

	expectOk := func(data []byte) {
		if err := json.Unmarshal(data, &obj); err != nil {
			t.Fatalf("expected unmarshal from JSON %q; but got %v", data, err)
		}
	}

	expectFail := func(data []byte) {
		if err := json.Unmarshal(data, &obj); err == nil {
			t.Fatalf("expected failed unmarshal from JSON %q; but got none", data)
		}
	}

	expectOk([]byte(`{"handle": "zakame", "time": "1536472243", "status": "hello, world"}`))
	expectOk([]byte(`{"handle": "zakame", "time": 1536472243, "status": "hello, world"}`))
	expectOk([]byte(`{"handle": "zakame", "status": "hello, world"}`))

	expectFail([]byte(`{"handle": "zakame", "time": "Sun, 09 Sep 2018 05:50:42 GMT", "status": "hello, world"}`))
	expectFail([]byte(`{"handle": "zakame", "time": "2006-01-02T15:04:05Z07:00", "status": "hello, world"}`))
}

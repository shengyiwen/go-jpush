package jpush

import (
	"encoding/json"
	"testing"
)

func TestAudience_All(t *testing.T) {
	all := NewAudience().All()

	marshal, err := json.Marshal(all.ToJsonElement())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}

func TestAudience_SetRegistrationIds(t *testing.T) {
	all := NewAudience().SetRegistrationIds([]string{"1111"})

	marshal, err := json.Marshal(all.ToJsonElement())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}

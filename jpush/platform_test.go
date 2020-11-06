package jpush

import (
	"encoding/json"
	"testing"
)

func TestPlatform_Android(t *testing.T) {
	platform := NewPlatform().Android()

	marshal, err := json.Marshal(platform.ToJsonElement())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}

func TestPlatform_Ios(t *testing.T) {
	platform := NewPlatform().Ios()

	marshal, err := json.Marshal(platform.ToJsonElement())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}

func TestPlatform_AddDeviceType(t *testing.T) {
	platform := NewPlatform().AddDeviceType(Ios).AddDeviceType(Android)
	marshal, err := json.Marshal(platform.ToJsonElement())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}

func TestPlatform_All(t *testing.T) {
	platform := NewPlatform().All()
	marshal, err := json.Marshal(platform.ToJsonElement())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}

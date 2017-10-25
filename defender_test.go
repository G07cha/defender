package defender

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	max := 5
	duration := 1 * time.Second
	banDuration := 1 * time.Hour
	instance := New(max, duration, banDuration)

	if instance.Max != max {
		t.Errorf("Expected Max to be %d, but it was %d instead", max, instance.Max)
	}

	if instance.Duration != duration {
		t.Errorf("Expected Duration to be %d, but it was %d instead", duration, instance.Duration)
	}

	if instance.BanDuration != banDuration {
		t.Errorf("Expected BanDuration to be %d, but it was %d instead", banDuration, instance.BanDuration)
	}
}

func TestClient(t *testing.T) {
	key := "127.0.0.1"
	instance := New(5, 1*time.Second, 1*time.Hour)
	client, ok := instance.Client(key)

	if ok != true {
		t.Errorf("Expected client to be returned successfully")
	}

	if client.key != key {
		t.Errorf("Expected client key to be equal %s, but it was %s instead", key, client.key)
	}
}

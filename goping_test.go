package goping

import "testing"

func TestPing(t *testing.T) {
	got := Ping()
	if got != "pong" {
		t.Errorf("Ping() = %s; want \"pong\"", got)
	}
}

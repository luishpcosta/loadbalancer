package domain

import (
	"sync"
	"testing"
)

func TestSetAlive(t *testing.T) {
	b := Backend{
		URL:          nil,
		Alive:        false,
		mux:          sync.RWMutex{},
		ReverseProxy: nil,
	}
	want := true
	b.SetAlive(want)
	got := b.Alive

	if got != want {
		t.Errorf("got %v, wanted %v", got, true)
	}
}

func TestIsAlive(t *testing.T){
	b := Backend{
		Alive:        true,
		URL:          nil,
		mux:          sync.RWMutex{},
		ReverseProxy: nil,
	}

	want := false
	b.SetAlive(want)
	got := b.IsAlive()

	if got != want {
		t.Errorf("got %v, wanted %v", got, false)
	}
}
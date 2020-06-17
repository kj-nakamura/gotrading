package main

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestExampleSuccess(t *testing.T) {
	result, err := example("hoge")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 1 {
		t.Fatal("failed test")
	}
}

func TestExampleFailed(t *testing.T) {
	result, err := example("fuga")
	if err == nil {
		t.Fatal("failed test")
	}
	if result != 0 {
		t.Fatal("failed test")
	}
}

package tests

import "testing"

func TestOk(t *testing.T) {
  if 2 + 2 != 4 {
    t.Fatalf("Ok test failed")
  }
}

func TestFail(t *testing.T) {
  if 2 + 2 != 5 {
    t.Fatalf("Fail test failed")
  }
}

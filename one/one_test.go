package one

import "testing"

func TestTheTest(t *testing.T) {
	one := "one"
	if one != "one" {t.Fatalf("expected variable one to equal '%s'\n", one)}
}
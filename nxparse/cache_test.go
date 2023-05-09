package nxparse

import (
	"errors"
	"testing"
)

func TestNilCachedFormat(t *testing.T) {
	frmt, err := cache.get("")
	if err != nil {
		t.Fatal(err)
	}
	testVarsMatch(t, frmt.Vars, nilVars)

	err = cache.set("", Format{Vars: nilVars})
	if err == nil {
		t.Fatal("expecting error on cache.set() got nil")
	}
}

func TestOverwriteCombined(t *testing.T) {
	err := cache.set(Combined, Format{Vars: nilVars})
	if !errors.Is(err, ErrOverwriteCache) {
		t.Fatalf("wrong error expecting error on cache.set()\ngot: %s", err)
	}
}

func TestCombinedCachedFormat(t *testing.T) {
	frmt, err := cache.get(Combined)
	if err != nil {
		t.Fatal(err)
	}
	testVarsMatch(t, frmt.Vars, combinedVars)

	err = cache.set(Combined, Format{Vars: combinedVars})
	if err == nil {
		t.Fatal("expecting error on cache.set() got nil")
	}
}

func TestCacheSet(t *testing.T) {
	frmtString := "$test_field1 ~-- $test_field2 ![[ $http_test1???$http_test2#$#"
	frmtStruct := Format{Vars: map[int]FormatVar{
		0: {"test_field1", " ~-- "},
		1: {"test_field2", " ![[ "},
		2: {"http_test1", "???"},
		3: {"http_test2", "#$#"},
	}}

	err := cache.set(frmtString, frmtStruct)
	if err != nil {
		t.Fatal(err)
	}

	frmt, err := cache.get(frmtString)
	if err != nil {
		t.Fatal(err)
	}
	testVarsMatch(t, frmt.Vars, frmtStruct.Vars)
}

func testVarsMatch(t *testing.T, got map[int]FormatVar, exp map[int]FormatVar) {
	for k := range got {
		if got[k] != exp[k] {
			t.Fatalf("combined format invalid\ngot %v\nexpect %v\n", got, exp)
		}
	}
}

package reader_test

import (
	"strings"
	"testing"

	"github.com/yookoala/lottery/reader"
)

var testXLSXdata = [][]string{
	{
		"Name",
		"Email",
		"Phone",
	},
	{
		"Alice",
		"alice@foobar.com",
		"12345678",
	},
	{
		"Bob",
		"bob@foobar.com",
		"12345678-2",
	},
	{
		"Christine",
		"christine@foobar.com",
		"12345678-3",
	},
	{
		"Daniel",
		"daniel@foobar.com",
		"12345678-4",
	},
	{
		"Elaine",
		"elaine@foobar.com",
		"12345678-5",
	},
	{
		"Flenn",
		"flenn@foobar.com",
		"12345678-6",
	},
}

func TestXLSReader_nonsheet(t *testing.T) {
	_, err := reader.OpenXLSXSheet("./_test/test.xlsx", 12)
	if err == nil {
		t.Error("failed to trigger error reading non-exist page")
	}
}

func TestXLSReader(t *testing.T) {
	f, err := reader.OpenXLSXSheet("./_test/test.xlsx", 0)
	if err != nil {
		t.Errorf("unable to open xlsx file: %#v", err.Error())
	}
	for i := 0; i < f.Len(); i++ {
		r, err := f.Row(i)
		if err != nil {
			t.Errorf("unable to open row: %#v", err.Error())
		}
		for j := 0; j < r.Len(); j++ {
			v := r.ReadString(j)
			if exp := testXLSXdata[i][j]; v != exp {
				t.Errorf("unexpected value at (%d, %d)\n"+
					"expected: %#v; got: %#v",
					i, j, exp, v)
			}
			t.Logf("value: %#v", v)
		}
	}
}

func TestXLSXReadMulti(t *testing.T) {

	f, err := reader.OpenXLSXSheet("./_test/test.xlsx", 0)
	if err != nil {
		t.Errorf("unable to open xlsx file: %#v", err.Error())
	}

	for i := 0; i < f.Len(); i++ {
		strs, err := reader.ReadMulti(f, 0, 1)(i)
		if err != nil {
			t.Errorf("unable to ReadMulti on line %d: %#v",
				i, err.Error())
		}

		res := strings.Join(strs, ", ")
		exp := testXLSXdata[i][0] + ", " + testXLSXdata[i][1]
		if res != exp {
			t.Errorf("row %d exptects %#v, got %#v",
				exp, res)
		}

		t.Logf("%#v", res)
	}
}

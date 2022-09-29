package battery_test

import (
	"battery"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func TestParseAcpiOutput(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/acpi.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := battery.Status{
		ChargePercent: 100,
	}
	got, err := battery.ParseAcpiOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

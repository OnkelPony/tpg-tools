//go:build integration

package battery_test

import (
	"battery"
	"testing"
)

func TestGetAcpiOutput(t *testing.T) {
	t.Parallel()
	text, err := battery.GetAcpiOutput()
	if err != nil {
		t.Fatal(err)
	}
	status, err := battery.ParseAcpiOutput(text)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Charge %d%%", status.ChargePercent)
}

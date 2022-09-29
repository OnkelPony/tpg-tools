package battery

import (
	"fmt"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

var r = regexp.MustCompile("([0-9]+)%")

func ParseAcpiOutput(acpiOutput string) (Status, error) {
	m := r.FindStringSubmatch(acpiOutput)
	if len(m) < 2 {
		return Status{}, fmt.Errorf("failed to parse acpi output: %q", acpiOutput)
	}
	charge, err := strconv.Atoi(m[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", m[1])
	}
	return Status{ChargePercent: charge}, nil
}

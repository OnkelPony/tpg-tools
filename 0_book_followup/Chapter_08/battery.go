package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

var r = regexp.MustCompile("([0-9]+)%")

func GetStatus() (Status, error) {
	text, err := GetAcpiOutput()
	if err != nil {
		return Status{}, err
	}
	return ParseAcpiOutput(text)
}

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

func GetAcpiOutput() (string, error) {
	data, err := exec.Command("/usr/bin/acpi", "-i").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

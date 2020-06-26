package config

import (
	"testing"
)

var testSample = []byte(`version: 2
updates:
- package-ecosystem: gomod
  directory: /
  schedule:
    interval: daily
`)

func TestConfigUnmarshal(t *testing.T) {

	c := New()
	err := c.Unmarshal(testSample)
	if err != nil {
		t.Error(err)
	}
	if c.Version != 2 {
		t.Error("Version not equal")
	}
	if len(c.Updates) != 1 {
		t.Error("Updates not equal")
	}
	if c.Updates[0].PackageEcosystem != "gomod" {
		t.Error("Updates 0 PackageEcosystem not equal")
	}
	if c.Updates[0].Directory != "/" {
		t.Error("Updates 0 Directory not equal")
	}
	if c.Updates[0].Schedule.Interval != "daily" {
		t.Error("Updates 0 Schedule Interval not equal")
	}
}

func TestConfigMarshal(t *testing.T) {
	c := New()
	c.Updates = []Update{
		{
			PackageEcosystem: "gomod",
			Directory:        "/",
			Schedule: Schedule{
				Interval: ScheduleIntervalDaily,
			},
		},
	}
	b, err := c.Marshal()
	if err != nil {
		t.Error(err)
	}
	if string(b) != string(testSample) {
		t.Error("Marshal bytes not equal", string(b))
	}
}

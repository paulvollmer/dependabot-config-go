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

func TestNew(t *testing.T) {
	c := New()
	if c.Version != 2 {
		t.Errorf("Config.Version is %v, want %v", c.Version, 2)
	}
}

func TestConfigAddUpdate(t *testing.T) {
	c := New()
	c.AddUpdate(Update{PackageEcosystem: PackageEcosystemGomod})
	if len(c.Updates) != 1 {
		t.Errorf("Config.Updates lenght is %v, want %v", len(c.Updates), 1)
		return
	}
	if c.Updates[0].PackageEcosystem != PackageEcosystemGomod {
		t.Errorf("Config.Updates[0].PackageEcosystem lenght is %v, want %v", c.Updates[0].PackageEcosystem, PackageEcosystemGomod)
	}
}

func TestConfigHasPackageEcosystem(t *testing.T) {
	cfg := Config{
		Updates: []Update{
			{PackageEcosystem: PackageEcosystemGomod},
			{PackageEcosystem: PackageEcosystemGitHubActions},
		},
	}

	tests := []struct {
		packageEcosystem string
		want             bool
	}{
		{
			packageEcosystem: PackageEcosystemGomod,
			want:             true,
		},
		{
			packageEcosystem: PackageEcosystemNpm,
			want:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.packageEcosystem, func(t *testing.T) {
			if got := cfg.HasPackageEcosystem(tt.packageEcosystem); got != tt.want {
				t.Errorf("HasPackageEcosystem %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidPackageEcosystem(t *testing.T) {
	tests := []struct {
		packageEcosystem string
		want             bool
	}{
		{PackageEcosystemBundler, true},
		{PackageEcosystemCargo, true},
		{PackageEcosystemComposer, true},
		{PackageEcosystemDocker, true},
		{PackageEcosystemElm, true},
		{PackageEcosystemGitsubmodule, true},
		{PackageEcosystemGitHubActions, true},
		{PackageEcosystemGomod, true},
		{PackageEcosystemGradle, true},
		{PackageEcosystemMaven, true},
		{PackageEcosystemMix, true},
		{PackageEcosystemNpm, true},
		{PackageEcosystemNuGet, true},
		{PackageEcosystemPip, true},
		{PackageEcosystemTerraform, true},
		{"wrong", false},
	}
	for _, tt := range tests {
		t.Run(tt.packageEcosystem, func(t *testing.T) {
			if got := IsValidPackageEcosystem(tt.packageEcosystem); got != tt.want {
				t.Errorf("IsValidPackageEcosystem is %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidScheduleInterval(t *testing.T) {
	tests := []struct {
		scheduleInterval string
		want             bool
	}{
		{ScheduleIntervalDaily, true},
		{ScheduleIntervalWeekly, true},
		{ScheduleIntervalMonthly, true},
		{"wrong", false},
	}
	for _, tt := range tests {
		t.Run(tt.scheduleInterval, func(t *testing.T) {
			if got := IsValidScheduleInterval(tt.scheduleInterval); got != tt.want {
				t.Errorf("IsValidScheduleInterval is %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidScheduleIntervalDay(t *testing.T) {
	tests := []struct {
		scheduleIntervalDay string
		want                bool
	}{
		{ScheduleIntervalDayMonday, true},
		{ScheduleIntervalDayTuesday, true},
		{ScheduleIntervalDayWednesday, true},
		{ScheduleIntervalDayThursday, true},
		{ScheduleIntervalDayFriday, true},
		{ScheduleIntervalDaySaturday, true},
		{ScheduleIntervalDaySunday, true},
		{"wrong", false},
	}
	for _, tt := range tests {
		t.Run(tt.scheduleIntervalDay, func(t *testing.T) {
			if got := IsValidScheduleIntervalDay(tt.scheduleIntervalDay); got != tt.want {
				t.Errorf("IsValidScheduleIntervalDay is %v, want %v", got, tt.want)
			}
		})
	}
}

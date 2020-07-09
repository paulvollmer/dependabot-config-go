package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()
	assert.Equal(t, 2, c.Version)
}

func TestConfigAddUpdate(t *testing.T) {
	c := Config{}
	c.AddUpdate(Update{PackageEcosystem: PackageEcosystemGomod})
	assert.Len(t, c.Updates, 1)
	assert.Equal(t, PackageEcosystemGomod, c.Updates[0].PackageEcosystem)
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
			result := cfg.HasPackageEcosystem(tt.packageEcosystem)
			assert.Equal(t, tt.want, result)
		})
	}
}

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
	assert.Nil(t, err)
	assert.Equal(t, 2, c.Version)
	assert.Len(t, c.Updates, 1)
	assert.Equal(t, "gomod", c.Updates[0].PackageEcosystem)
	assert.Equal(t, "/", c.Updates[0].Directory)
	assert.Equal(t, "daily", c.Updates[0].Schedule.Interval)
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
	assert.Nil(t, err)
	assert.Equal(t, testSample, b)
}

func TestUpdateAddAllow(t *testing.T) {
	u := Update{}
	u.AddAllow(NewAllow("n", "t"))
	assert.Len(t, u.Allow, 1)
}

func TestUpdateAddAssignee(t *testing.T) {
	u := Update{}
	u.AddAssignee("a")
	assert.Len(t, u.Assignees, 1)
}

func TestUpdateAddIgnore(t *testing.T) {
	u := Update{}
	u.AddIgnore(NewIgnore("d", []string{"v1"}))
	assert.Len(t, u.Ignore, 1)
}

func TestUpdateAddLabel(t *testing.T) {
	u := Update{}
	u.AddLabel("l")
	assert.Len(t, u.Labels, 1)
}

func TestUpdateAddReviewer(t *testing.T) {
	u := Update{}
	u.AddReviewer("r")
	assert.Len(t, u.Reviewers, 1)
}

func TestNewSchedule(t *testing.T) {
	s, err := NewSchedule(ScheduleIntervalDaily)
	assert.Nil(t, err)
	assert.Equal(t, ScheduleIntervalDaily, s.Interval)
}

func TestNewAllow(t *testing.T) {
	result := NewAllow("dep-name", "dep-type")
	assert.Equal(t, "dep-name", result.DependencyName)
	assert.Equal(t, "dep-type", result.DependencyType)
}

func TestNewCommitMessage(t *testing.T) {
	c := NewCommitMessage("p", "pd", "i")
	assert.Equal(t, "p", c.Prefix)
	assert.Equal(t, "pd", c.PrefixDevelopment)
	assert.Equal(t, "i", c.Include)
}

func TestNewIgnore(t *testing.T) {
	i := NewIgnore("dep-name", []string{"v1"})
	assert.Equal(t, "dep-name", i.DependencyName)
	assert.Equal(t, []string{"v1"}, i.Versions)
}

func TestIgnoreAddVersion(t *testing.T) {
	i := Ignore{}
	i.AddVersion("v1")
	assert.Equal(t, []string{"v1"}, i.Versions)
}

func TestNewPullRequestBranchName(t *testing.T) {
	result := NewPullRequestBranchName("-")
	assert.Equal(t, "-", result.Separator)
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
			result := IsValidPackageEcosystem(tt.packageEcosystem)
			assert.Equal(t, tt.want, result)
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
			result := IsValidScheduleInterval(tt.scheduleInterval)
			assert.Equal(t, tt.want, result)
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
			result := IsValidScheduleIntervalDay(tt.scheduleIntervalDay)
			assert.Equal(t, tt.want, result)
		})
	}
}

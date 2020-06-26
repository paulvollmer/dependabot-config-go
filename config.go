package config

import (
	"gopkg.in/yaml.v2"
)

// Config for dependabot
type Config struct {
	Version int      `yaml:"version"`
	Updates []Update `yaml:"updates"`
}

// Update docs: https://help.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates
type Update struct {
	PackageEcosystem      string                 `yaml:"package-ecosystem"`
	Directory             string                 `yaml:"directory"`
	Schedule              Schedule               `yaml:"schedule"`
	Allow                 []*Allow               `yaml:"allow,omitempty"`
	Assignees             []*string              `yaml:"assignees,omitempty"`
	CommitMessage         *CommitMessage         `yaml:"commit-message,omitempty"`
	Ignore                []*Ignore              `yaml:"ignore,omitempty"`
	Labels                []*string              `yaml:"labels,omitempty"`
	Milestone             *int                   `yaml:"milestone,omitempty"`
	OpenPullRequestsLimit *int                   `yaml:"open-pull-requests-limit,omitempty"`
	PullRequestBranchName *PullRequestBranchName `yaml:"pull-request-branch-name,omitempty"`
	RebaseStrategy        *string                `yaml:"rebase-strategy,omitempty"`
	Reviewers             []*string              `yaml:"reviewers,omitempty"`
	TargetBranch          *string                `yaml:"target-branch,omitempty"`
	VersioningStrategy    *string                `yaml:"versioning-strategy,omitempty"`
}

// Schedule docs: https://help.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates#scheduleinterval
type Schedule struct {
	Interval string `yaml:"interval"`
	Day      string `yaml:"day,omitempty"`
	Time     string `yaml:"time,omitempty"`
	Timezone string `yaml:"timezone,omitempty"`
}

// Allow docs: https://help.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates#allow
type Allow struct {
	DependencyName string `"dependency-name,omitempty"`
	DependencyType string `"dependency-type,omitempty"`
}

// CommitMessage docs: https://help.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates#commit-message
type CommitMessage struct {
	Prefix            string `yaml:"prefix,omitempty"`
	PrefixDevelopment string `yaml:"prefix-development,omitempty"`
	Include           string `yaml:"include,omitempty"`
}

// Ignore docs: https://help.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates#ignore
type Ignore struct {
	DependencyName string   `yaml:"dependency-name,omitempty"`
	Versions       []string `yaml:"versions,omitempty"`
}

// PullRequestBranchName docs: https://help.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates#pull-request-branch-nameseparator
type PullRequestBranchName struct {
	Separator string `yaml:"separator"`
}

const (
	PackageEcosystemBundler   = "bundler"
	PackageEcosystemCargo     = "cargo"
	PackageEcosystemComposer  = "composer"
	PackageEcosystemDocker    = "docker"
	PackageEcosystemElm       = "elm"
	PackageEcosystemgit       = "submodule gitsubmodule"
	PackageEcosystemGitHub    = "Actions github-actions"
	PackageEcosystemGo        = "modules gomod"
	PackageEcosystemGradle    = "gradle"
	PackageEcosystemMaven     = "maven"
	PackageEcosystemMix       = "mix"
	PackageEcosystemnpm       = "npm"
	PackageEcosystemNuGet     = "nuget"
	PackageEcosystempip       = "pip"
	PackageEcosystemTerraform = "terraform"

	ScheduleIntervalDaily   = "daily"
	ScheduleIntervalWeekly  = "weekly"
	ScheduleIntervalMonthly = "monthly"

	ScheduleIntervalDayMonday    = "monday"
	ScheduleIntervalDayTuesday   = "tuesday"
	ScheduleIntervalDayWednesday = "wednesday"
	ScheduleIntervalDayThursday  = "thursday"
	ScheduleIntervalDayFriday    = "friday"
	ScheduleIntervalDaySaturday  = "saturday"
	ScheduleIntervalDaySunday    = "sunday"
)

func New() *Config {
	c := Config{Version: 2}
	return &c
}

func (c *Config) Unmarshal(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func (c *Config) Marshal() ([]byte, error) {
	return yaml.Marshal(c)
}

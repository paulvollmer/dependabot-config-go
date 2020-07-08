package config

import (
	"gopkg.in/yaml.v2"
)

// Config for dependabot
type Config struct {
	Version int      `yaml:"version"`
	Updates []Update `yaml:"updates"`
}

// AddUpdate append an Update item to the Updates array
func (c *Config) AddUpdate(item Update) {
	c.Updates = append(c.Updates, item)
}

// HasPackageEcosystem return true if the given package-ecosystem string exist at the Updates array
func (c *Config) HasPackageEcosystem(e string) bool {
	for i := 0; i < len(c.Updates); i++ {
		if c.Updates[i].PackageEcosystem == e {
			return true
		}
	}
	return false
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
	DependencyName string `yaml:"dependency-name,omitempty"`
	DependencyType string `yaml:"dependency-type,omitempty"`
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
	PackageEcosystemBundler       = "bundler"
	PackageEcosystemCargo         = "cargo"
	PackageEcosystemComposer      = "composer"
	PackageEcosystemDocker        = "docker"
	PackageEcosystemElm           = "elm"
	PackageEcosystemGitsubmodule  = "gitsubmodule"
	PackageEcosystemGitHubActions = "github-actions"
	PackageEcosystemGomod         = "gomod"
	PackageEcosystemGradle        = "gradle"
	PackageEcosystemMaven         = "maven"
	PackageEcosystemMix           = "mix"
	PackageEcosystemNpm           = "npm"
	PackageEcosystemNuGet         = "nuget"
	PackageEcosystemPip           = "pip"
	PackageEcosystemTerraform     = "terraform"

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

func IsValidPackageEcosystem(e string) bool {
	if e == PackageEcosystemBundler ||
		e == PackageEcosystemCargo ||
		e == PackageEcosystemComposer ||
		e == PackageEcosystemDocker ||
		e == PackageEcosystemElm ||
		e == PackageEcosystemGitsubmodule ||
		e == PackageEcosystemGitHubActions ||
		e == PackageEcosystemGomod ||
		e == PackageEcosystemGradle ||
		e == PackageEcosystemMaven ||
		e == PackageEcosystemMix ||
		e == PackageEcosystemNpm ||
		e == PackageEcosystemNuGet ||
		e == PackageEcosystemPip ||
		e == PackageEcosystemTerraform {
		return true
	}
	return false
}

func IsValidScheduleInterval(i string) bool {
	if i == ScheduleIntervalDaily ||
		i == ScheduleIntervalWeekly ||
		i == ScheduleIntervalMonthly {
		return true
	}
	return false
}

func IsValidScheduleIntervalDay(i string) bool {
	if i == ScheduleIntervalDayMonday ||
		i == ScheduleIntervalDayTuesday ||
		i == ScheduleIntervalDayWednesday ||
		i == ScheduleIntervalDayThursday ||
		i == ScheduleIntervalDayFriday ||
		i == ScheduleIntervalDaySaturday ||
		i == ScheduleIntervalDaySunday {
		return true
	}
	return false
}

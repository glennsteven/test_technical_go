package util

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	major, minor, patch int
}

// NewVersion creates a new Version object from a version string.
func NewVersion(versionString string) *Version {
	parts := strings.Split(versionString, ".")
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])
	return &Version{major, minor, patch}
}

// NextVersion generates the next version string by incrementing the patch number.
func (v *Version) NextVersion() string {
	v.patch++
	return v.String()
}

// String returns the version string representation.
func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func Test(str string) *Version {
	return &Version{}
}

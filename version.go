package main

import (
	"fmt"
	"sort"
	"strings"
)

type Version struct {
	major int
	minor int
	patch int
	build int
}

type Versions []*Version

func parseVersion(tag string) (Version, error) {
	v := Version{}

	_, err := fmt.Sscanf(tag, "%d.%d.%d-%d",
		&v.major,
		&v.minor,
		&v.patch,
		&v.build,
	)

	return v, err
}

func (s Versions) Len() int {
	return len(s)
}

func (s Versions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (v Version) LessThan(o Version) bool {
	if v.major < o.major {
		return true
	} else if v.major > o.major {
		return false
	}

	if v.minor < o.minor {
		return true
	} else if v.minor > o.minor {
		return false
	}

	if v.patch < o.patch {
		return true
	} else if v.patch > o.patch {
		return false
	}

	if v.build < o.build {
		return true
	}

	return false
}

func (s Versions) Less(i, j int) bool {
	return !s[i].LessThan(*s[j])
}

func versionSort(versions []*Version) {
	sort.Sort(Versions(versions))
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d-%d", v.major, v.minor, v.patch, v.build)
}

func versionString(versions []*Version) string {
	entries := []string{}
	for _, v := range versions {
		entries = append(entries, v.String())
	}

	return strings.Join(entries, ", ")
}

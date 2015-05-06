package cgroups

import (
	"bytes"
	"testing"
)

const (
	cgroupsContents = `11:hugetlb:/
10:perf_event:/
9:blkio:/
8:net_cls:/
7:freezer:/
6:devices:/
5:memory:/
4:cpuacct,cpu:/
3:cpuset:/
)

func TestParseCgroups(t *testing.T) {
	r := bytes.NewBuffer([]byte(cgroupsContents))
	_, err := ParseCgroupFile("blkio", r)
	if err != nil {
		t.Fatal(err)
	}
}

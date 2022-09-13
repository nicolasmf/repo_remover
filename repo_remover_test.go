package main

import (
	"testing"
)

func TestGetname(t *testing.T) {
	data := []byte(`[docker-ce-stable]
name=Docker CE Stable - $basearch
baseurl=https://download.docker.com/linux/centos/7/$basearch/stable
enabled=1
gpgcheck=1
gpgkey=https://download.docker.com/linux/centos/gpg
`)
	result := GetRepoNames(string(data), "=", " -", " $", "\n")
	if result != "Docker CE Stable" {
		t.Errorf("Expected Docker CE Stable, got %s", result)
	}

	data1 := []byte(`[rpmfusion-free]
name=RPM Fusion for Fedora $releasever - Free
#baseurl=http://download1.rpmfusion.org/free/fedora/releases/$releasever/Everything/$basearch/os/
metalink=https://mirrors.rpmfusion.org/metalink?repo=free-fedora-$releasever&arch=$basearch
enabled=1
metadata_expire=14d
`)
	result1 := GetRepoNames(string(data1), "=", " -", " $", "\n")
	if result1 != "RPM Fusion for Fedora" {
		t.Errorf("RPM Fusion for Fedora, got %s", result1)
	}

	data2 := []byte(`[copr:copr.fedorainfracloud.org:atim:bottom]
	name=Bottom owned by atim
	baseurl=https://download.copr.fedorainfracloud.org/results/atim/bottom/fedora-$>
	type=rpm-md
	skip_if_unavailable=True
	`)
	result2 := GetRepoNames(string(data2), "=", " -", " $", "\n")
	if result2 != "Bottom owned by atim" {
		t.Errorf("Bottom owned by atim, got %s", result2)
	}
}

package probelauncher

import "testing"

const SAMPLE_INPUT = "target area: x=20..30, y=-10..-5"

func TestNewTarget(t *testing.T) {
	target := NewTarget(SAMPLE_INPUT)
	expected := SAMPLE_INPUT
	if expected != target.String() {
		t.Errorf("Expected to parse into (%s), but found (%s)", expected, target)
	}
}

func TestLandsIn(t *testing.T) {
	probe := NewProbe(7, 2)
	target := NewTarget(SAMPLE_INPUT)
	expected := true
	actual, _ := probe.LandsIn(target)
	if expected != actual {
		t.Errorf("Expected probe to be in target (%s), but was not.", target)
	}
}

func TestGetHits(t *testing.T) {
	target := NewTarget(SAMPLE_INPUT)
	hits := target.GetHits()
	expected := 112
	actual := len(hits)
	if expected != actual {
		t.Errorf("Expected to find %d viable probes, found %d.", expected, actual)
	}
}

func TestGetMaxY(t *testing.T) {
	target := NewTarget(SAMPLE_INPUT)
	hits := target.GetHits()
	actual := hits.GetMaxY()
	expected := 45
	if expected != actual {
		t.Errorf("Expected to find %d as max y, found %d.", expected, actual)
	}
}

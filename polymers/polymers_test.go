package polymers

import (
	"testing"
)

var (
	SAMPLE_TEMPLATE = NewTemplate("NNCB")
	SAMPLE_RULES    = NewRuleSet([]string{
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	})
)

func TestGetScore(t *testing.T) {
	tmpl := SAMPLE_TEMPLATE
	expected := 1588
	for i := 1; i <= 10; i++ {
		tmpl.ApplyRules(SAMPLE_RULES)
	}
	actual := tmpl.GetScore()
	if expected != actual {
		t.Errorf("Expected score %d after 10 rounds, got %d.", expected, actual)
	}
}

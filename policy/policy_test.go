package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 56, Capacity: 72, Latency: 11, Risk: 7, Weight: 8}
	if got := Score(signal); got != 122 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 63, Capacity: 102, Latency: 8, Risk: 14, Weight: 10}
	if got := Score(signal); got != 142 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 87, Capacity: 105, Latency: 13, Risk: 8, Weight: 10}
	if got := Score(signal); got != 209 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}

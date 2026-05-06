package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 56, Capacity: 72, Latency: 11, Risk: 7, Weight: 8}, wantScore: 122, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 63, Capacity: 102, Latency: 8, Risk: 14, Weight: 10}, wantScore: 142, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 87, Capacity: 105, Latency: 13, Risk: 8, Weight: 10}, wantScore: 209, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}

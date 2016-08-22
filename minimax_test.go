package minimax

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	mm := New()

	if mm.Score != nil {
		t.Error("score should be nil")
	}

	if mm.parent != nil {
		t.Error("parent not nil")
	}

	if len(mm.children) != 0 {
		t.Error("children should be empty")
	}
}

func TestAddEmpty(t *testing.T) {
	mm := New()
	mm.Add("")

	if len(mm.children) != 1 {
		t.Error("expected 1 child item")
	}
}

func TestAddScore(t *testing.T) {
	mm := New()

	n := mm.AddTerminal(123, "")

	if n.Score == nil {
		t.Error("score not calculated")
	} else if *n.Score != 123 {
		t.Error("score should be 123")
	}
}

func TestAddPlayer(t *testing.T) {
	mm := New()

	if mm.isOpponent {
		t.Error("level one should be player 1")
	}

	level2 := mm.Add("")

	if !level2.isOpponent {
		t.Error("level 2 should be player 2")
	}

	level3 := level2.Add("")

	if level3.isOpponent {
		t.Error("level3 should be player 1")
	}
}

// Test as per https://en.wikipedia.org/wiki/Minimax#/media/File:Plminmax.gif
func TestEvaluateWikipedia(t *testing.T) {
	r1 := New()

	r2a := r1.Add("a")

	r3a1 := r2a.Add("r3a1")
	r3a2 := r2a.Add("r3a2")

	r4a1 := r3a1.Add("r4a1")
	r4a2 := r3a1.Add("r4a2")
	r4a3 := r3a2.Add("r4a3")

	r4a1.AddTerminal(5, "")
	r4a1.AddTerminal(6, "")

	r4a2.AddTerminal(7, "")
	r4a2.AddTerminal(4, "")
	r4a2.AddTerminal(5, "")

	r4a3.AddTerminal(3, "")

	//B
	r2b := r1.Add("b")
	r3b1 := r2b.Add("r3b1")
	r3b2 := r2b.Add("r3b2")

	r4b1 := r3b1.Add("r4b1")
	r4b2 := r3b1.Add("r4b2")
	r4b3 := r3b2.Add("r4b3")

	r4b1.AddTerminal(6, "")

	r4b2.AddTerminal(6, "")
	r4b2.AddTerminal(9, "")

	r4b3.AddTerminal(7, "")

	//B
	r2c := r1.Add("c")
	r3c1 := r2c.Add("r3c1")
	r3c2 := r2c.Add("r3c2")

	r4c1 := r3c1.Add("r4c1")
	r4c2 := r3c2.Add("r4c2")
	r4c3 := r3c2.Add("r4c3")

	r4c1.AddTerminal(5, "")

	r4c2.AddTerminal(9, "")
	r4c2.AddTerminal(8, "")

	r4c3.AddTerminal(6, "")

	r1.Evaluate()

	if r1.Score == nil {
		t.Error("Score not evaluated")
	} else if *r1.Score != 6 {
		t.Error("Evaluate failed: ", *r1.Score)
	}

	r1.GetBestChildNode()
	r1.Print(0)
}

func TestEvaluateYouTube(t *testing.T) {
	root := New()

	b := root.Add("B")
	c := root.Add("C")
	d := root.Add("D")

	b.AddTerminal(3, "B1")
	b.AddTerminal(12, "B2")
	b.AddTerminal(8, "B3")

	c.AddTerminal(2, "C1")
	c.AddTerminal(4, "C2")
	c.AddTerminal(6, "C3")

	d.AddTerminal(14, "D1")
	d.AddTerminal(5, "D2")
	d.AddTerminal(2, "D3")

	root.Evaluate()

	if root.Score == nil {
		t.Error("Score not evaluated")
	} else if *root.Score != 3 {
		t.Error("Evaluate failed: ", *root.Score)
	}

	t1 := root.children[0]
	t2 := b

	if t1 == t2 {
		fmt.Println("SAME")
	} else {
		fmt.Println("DIFFERENT")
	}
}

// Test as per https://en.wikipedia.org/wiki/Minimax#/media/File:Plminmax.gif
func TestIsTerminal(t *testing.T) {
	r1 := New()
	node := r1.AddTerminal(1, nil)

	if !node.isTerminal() {
		t.Error("Node should be a terminal node")
	}

	r1.Add(nil)

	if r1.isTerminal() {
		t.Error("R1 should not be a terminal node")
	}
}

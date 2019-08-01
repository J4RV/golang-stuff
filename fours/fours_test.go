package main

import (
	"testing"
)

func TestState_row(t *testing.T){
	g := New()
	row := g.row(0)
	if len(row) != BoardHorizontalSize {
		t.Error("Incorrect row size")
	}
}

func TestState_col(t *testing.T){
	g := New()
	col := g.col(0)
	if len(col) != BoardVerticalSize {
		t.Error("Incorrect col size")
	}
}

func TestState_Play(t *testing.T) {
	g := New()
	g.currPlayerColor = WHITE

	firstCol := 0
	lastCol := BoardHorizontalSize-1

	err := g.Play(firstCol)
	if err != nil {
		t.Error(err)
	}

	err = g.Play(firstCol)
	if err != nil {
		t.Error(err)
	}

	err = g.Play(lastCol)
	if err != nil {
		t.Error(err)
	}

	col := g.col(firstCol)
	if col[0] != WHITE {
		t.Errorf("Expected %s, found %s", WHITE.EnumString(), col[0].EnumString())
	}
	if col[1] != BLACK {
		t.Errorf("Expected %s, found %s", BLACK.EnumString(), col[0].EnumString())
	}
	if g.col(lastCol)[0] != WHITE {
		t.Errorf("Expected %s, found %s", WHITE.EnumString(), col[0].EnumString())
	}

	t.Log("State:\n" + g.String())
}
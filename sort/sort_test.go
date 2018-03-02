package sort

import (
	"testing"
	"algorithm-go/help"
)

func TestBubble(t *testing.T) {
	sources := help.PersonSlice(100)
	Bubble(&sources,help.SortByAsc)
	if !help.ValidSort(&sources, help.SortByAsc) {
		t.Log(sources)
		t.Error("bubble asc sort error")
	}
	Bubble(&sources, help.SortByDesc)
	if !help.ValidSort(&sources, help.SortByDesc) {
		t.Error("bubble desc sort error")
	}
}

func TestSelect(t *testing.T) {
	source := help.PersonSlice(100)
	Select(&source, help.SortByAsc)
	if !help.ValidSort(&source, help.SortByAsc) {
		t.Log(source)
		t.Error("select asc sort error")
	}
	Select(&source, help.SortByDesc)
	if !help.ValidSort(&source, help.SortByDesc) {
		t.Log(source)
		t.Error("select desc sort error")
	}
}

func TestInsert(t *testing.T) {
	source := help.PersonSlice(100)
	Insert(&source, help.SortByAsc)
	if !help.ValidSort(&source, help.SortByAsc) {
		t.Log(source)
		t.Error("select asc sort error")
	}
	Insert(&source, help.SortByDesc)
	if !help.ValidSort(&source, help.SortByDesc) {
		t.Log(source)
		t.Error("select desc sort error")
	}
}

func TestHill(t *testing.T) {
	source := help.PersonSlice(100)
	Hill(&source, help.SortByAsc)
	if !help.ValidSort(&source, help.SortByAsc) {
		t.Log(source)
		t.Error("hill asc sort error")
	}
	Hill(&source, help.SortByDesc)
	if !help.ValidSort(&source, help.SortByDesc) {
		t.Log(source)
		t.Error("hill desc sort error")
	}
}


package bookmark

import (
	"bytes"
	"testing"
)

func TestDup(t *testing.T) {
	item := NewItem()
	item.SetURL("https://hubblesite.org/")
	nItem := item.Dup()

	if item == &nItem {
		t.Errorf("Got: %p, Want: anything else", &nItem)
	}
	if !bytes.Equal(item.id, nItem.id) {
		t.Errorf("Got: %s, Want: %s", item.HexID(), nItem.HexID())
	}
}

package skiplist

import (
	"bytes"
	"testing"
)

func TestBasicOperations(t *testing.T) {
	s := New()

	for letter := 'a'; letter < 'a'+26; letter++ {
		s.Set(string(letter), []byte(string(letter)+"world"))
	}

	for letter := 'a'; letter < 'a'+26; letter++ {
		val, ok := s.Search(string(letter))
		if !ok {
			t.Errorf("could not find %c", letter)
		}
		expected := []byte(string(letter) + "world")
		if !bytes.Equal(val, expected) {
			t.Errorf("expected %v, got %v", expected, val)
		}
	}

	elements := 0
	s.ForEach(func(key string, value []byte) {
		elements++
	})
	if elements != 26 {
		t.Errorf("expected 26 elements, got %d", elements)
	}

	val, ok := s.Search("z")
	if !ok {
		t.Errorf("expected to find z, but did not find it")
	}
	if string(val) != "zworld" {
		t.Errorf("expected val to be zworld but got %s", string(val))
	}

	for letter := 'a'; letter < 'a'+26; letter++ {
		s.Delete(string(letter))
	}

	s.ForEach(func(key string, value []byte) {
		t.Errorf("there should be no elements in the skiplist")
	})
}

func makenode(key string) *node {
	return &node{Record: &record{Key: key, Value: []byte(key + "world")}, Next: make(map[int]*node)}
}

func TestLevelCleanup(t *testing.T) {
	s := skiplist{Levels: make(map[int]*node)}

	// anode is level 2
	anode := makenode("a")
	s.Levels[2] = anode
	s.Levels[1] = anode
	s.Levels[0] = anode

	// bnode is only level 0
	bnode := makenode("b")
	anode.Next[0] = bnode

	/* before deletion:
	L02 -> a -> nil
	L01 -> a -> nil
	L00 -> a -> b -> nil
	*/

	// deleting a should cause levels 1 and 2 to be deleted
	s.Delete("a")

	/* after deletion:
	L00 -> b -> nil
	*/

	if _, ok := s.Levels[2]; ok {
		t.Errorf("no node should have been set for level 2")
	}
	if _, ok := s.Levels[1]; ok {
		t.Errorf("no node should have been set for level 1")
	}

	n, ok := s.Levels[0]
	if !ok {
		t.Errorf("expected b node to still exist")
	}
	if n != bnode {
		t.Errorf("expected to get bnode, got %#v", n)
	}
}

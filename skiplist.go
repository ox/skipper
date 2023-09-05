package skiplist

import (
	"math"
	"math/rand"
	"strings"
)

const MaxLevels = 16

type record struct {
	Key   string
	Value []byte
}

type node struct {
	Record *record
	Next   map[int]*node
}

type skiplist struct {
	// Levels is a map pointing to the first Node of a given level
	Levels map[int]*node
}

type ForEachFN func(key string, value []byte)

type SkipList interface {
	Search(key string) ([]byte, bool)
	Set(key string, value []byte)
	Delete(key string)
	ForEach(ForEachFN)
	Debug()
}

// New creates a new SkipList
func New() SkipList {
	return &skiplist{
		Levels: make(map[int]*node),
	}
}

// randomLevel returns a random level for a new node to be inserted at
func randomLevel() int {
	var level float64 = 0
	for rand.Float32() < 0.50 {
		level++
	}
	return int(math.Min(level, MaxLevels))
}

// maxMapLevel finds the greatest level in the node map
func maxMapLevel(m map[int]*node) int {
	max := -1
	for key := range m {
		if key >= max {
			max = key
		}
	}

	return max
}

// Search finds the value for a certain key and if a value exists for that key
func (s *skiplist) Search(key string) ([]byte, bool) {
	level := maxMapLevel(s.Levels)
	if level < 0 {
		return []byte{}, false
	}

	// Find the right level to start on, making sure that the first node is < the key
	x := s.Levels[level]
	for ; level >= 0; level-- {
		x = s.Levels[level]
		if x != nil && strings.Compare(x.Record.Key, key) != 1 {
			break
		}
	}

	for ; level >= 0; level-- {
		if x != nil && x.Record.Key == key {
			return x.Record.Value, true
		}

		for x.Next[level] != nil && strings.Compare(x.Next[level].Record.Key, key) == -1 {
			x = x.Next[level]
		}
	}

	if x != nil && x.Next[0] != nil && x.Next[0].Record.Key == key {
		return x.Next[0].Record.Value, true
	}
	return []byte{}, false
}

// Set stores the key-value pair; updating records if necessary
func (s *skiplist) Set(key string, value []byte) {
	update := make(map[int]*node)
	maxLevel := maxMapLevel(s.Levels)
	level := maxLevel

	// This basically finds the nodes to the left of where key would be inserted
	// on all levels
	x := s.Levels[level]
	for ; level >= 0; level-- {
		if x == nil {
			panic("node is nil")
		}
		for x.Next[level] != nil && strings.Compare(x.Next[level].Record.Key, key) == -1 {
			x = x.Next[level]
		}
		update[level] = x
	}

	// x may or may not be the node pointing to a record we are updating
	if x != nil && x.Next[0] != nil && x.Next[0].Record.Key == key {
		x.Next[0].Record.Value = value
	} else {
		newLevel := randomLevel()
		newRecord := &record{Key: key, Value: value}
		newNode := &node{Record: newRecord, Next: make(map[int]*node)}
		if newLevel > maxLevel {
			// For each new level, have it point to the new node
			for interimLevel := maxLevel + 1; interimLevel <= newLevel; interimLevel++ {
				s.Levels[interimLevel] = newNode
			}
		}

		for level, node := range update {
			if level <= newLevel {
				newNode.Next[level] = node.Next[level]
				node.Next[level] = newNode
			}
		}
	}
}

func (s *skiplist) Delete(key string) {
	maxLevel := maxMapLevel(s.Levels)
	level := maxLevel

	// Find the right level to start on, making sure that the first node is < the key
	x := s.Levels[level]
	for ; level >= 0; level-- {
		x = s.Levels[level]
		if x != nil {
			if strings.Compare(x.Record.Key, key) != 1 {
				break
			}
		}
	}

	// Special case: If x is the node we want to delete and it's the first node in the level
	// we need to change the Levels map to point to it's next node.
	if x != nil && x.Record.Key == key {
		for ; level >= 0; level-- {
			s.Levels[level] = x.Next[level]
		}
	} else if x != nil && x.Next[level] != nil && x.Next[level].Record.Key == key {
		// Default case: the next node is the one we want to skip over
		for ; level >= 0; level-- {
			if x.Next[level] == x {
				// skip over the element to be removed
				x.Next[level] = x.Next[level].Next[level]
			}
		}
	}

	// Cleanup higher levels that point straight to nil
	for level = maxLevel; level >= 0; level-- {
		if s.Levels[level] == nil {
			delete(s.Levels, level)
		} else {
			break
		}
	}
}

// ForEach calls fn for each key in sorted order
func (s *skiplist) ForEach(fn ForEachFN) {
	for x := s.Levels[0]; x != nil; x = x.Next[0] {
		fn(x.Record.Key, x.Record.Value)
	}
}

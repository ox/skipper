package skiplist

import "fmt"

func (s *skiplist) Debug() {
	ret := fmt.Sprintf("SkipList, maxLevel: %d\n", maxMapLevel(s.Levels))
	for level := maxMapLevel(s.Levels); level >= 0; level-- {
		ret += fmt.Sprintf("L%02d -> ", level)
		x, ok := s.Levels[level]
		if !ok {
			ret += "nil\n"
			continue
		}

		for x != nil {
			ret += fmt.Sprintf("%s -> ", x.Record.Key)
			x = x.Next[level]
		}

		ret += "nil\n"
	}

	// ret += "\n"
	// for x := s.Levels[0]; x != nil; x = x.Next[0] {
	// 	ret += fmt.Sprintf("%s: %s\n", x.Record.Key, x.Record.Value)
	// }

	fmt.Println(ret)
}

// func main() {
// 	skiplist := New()
// 	for letter := 'a'; letter < 'a'+26; letter++ {
// 		skiplist.Set(string(letter), string(letter)+"world")
// 	}

// 	val, ok := skiplist.Search("y")
// 	if !ok {
// 		panic(fmt.Errorf("y should have been found!"))
// 	}

// 	fmt.Println("y = " + val)
// }

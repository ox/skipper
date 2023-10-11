package skipper

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Store(s SkipList, filename string) error {
	pathname, err := filepath.Abs(filename)
	if err != nil {
		panic(fmt.Errorf("could not get absolute path to %s", filename))
	}

	fmt.Printf("pathname: %s\n", pathname)
	f, err := os.OpenFile(pathname, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(fmt.Errorf("could not open store at %s: %w", pathname, err))
	}
	defer f.Close()

	s.ForEach(func(key string, value []byte) {
		f.WriteString(key)
		f.WriteString("=")
		f.Write(value)
		f.WriteString("\n")
	})
	return nil
}

func Load(filename string) (SkipList, error) {
	s := New()

	pathname, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("could not get absolute path to %s", filename)
	}

	fmt.Printf("pathname: %s\n", pathname)
	f, err := os.OpenFile(pathname, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("could not open store at %s: %w", pathname, err)
	}
	defer f.Close()

	buf := make([]byte, 0)
	for {
		segment := make([]byte, 64)
		_, err := f.Read(segment)
		if err == io.EOF {
			return s, nil
		}
		if err != nil {
			return nil, err
		}

		buf = append(buf, segment...)

		lines := bytes.SplitN(buf, []byte("\n"), 2)
		if len(lines) > 1 {
			line := lines[0]
			parts := bytes.SplitN(line, []byte{'='}, 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("expected an = in the line, got %s", line)
			}

			key := parts[0]
			value := parts[1]
			s.Set(string(key), value)

			buf = buf[len(line)+1:]
		}
	}
}

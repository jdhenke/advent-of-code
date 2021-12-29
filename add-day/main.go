package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	day := flag.Int("day", 0, "day to add")
	flag.Parse()
	if *day == 0 {
		log.Fatal("Must provide -day flag.")
	}
	if err := createDay(*day); err != nil {
		log.Fatal(err)
	}
	log.Println("Success")
}

const libFile = `package day$DAY

import (
	"io"
)

func Part1(r io.Reader) (answer int, err error) {
	return 0, nil
}

func Part2(r io.Reader) (answer int, err error) {
	return 0, nil
}
`

const testFile = `package day$DAY_test

import (
	"advent-of-code/day$DAY"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData = ""

func TestPart1(t *testing.T) {
	ans, err := day$DAY.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 0, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day$DAY.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 0, ans)
}
`

const solverLine = `		{$DAY, 1}: day$DAY.Part1,
		{$DAY, 2}: day$DAY.Part2,
`

func createDay(day int) error {
	dir := fmt.Sprintf("day%d", day)
	if err := os.Mkdir(dir, 0755); err != nil {
		return err
	}
	f := strings.Replace(libFile, `$DAY`, fmt.Sprint(day), -1)
	if err := ioutil.WriteFile(filepath.Join(dir, fmt.Sprintf("day%d.go", day)), []byte(f), 0644); err != nil {
		return err
	}
	tf := strings.Replace(testFile, `$DAY`, fmt.Sprint(day), -1)
	if err := ioutil.WriteFile(filepath.Join(dir, fmt.Sprintf("day%d_test.go", day)), []byte(tf), 0644); err != nil {
		return err
	}
	mainBytes, err := ioutil.ReadFile("main.go")
	if err != nil {
		return err
	}
	sl := strings.Replace(solverLine, `$DAY`, fmt.Sprint(day), -1)
	var buf bytes.Buffer
	var inSolvers bool
	for _, line := range strings.Split(string(mainBytes), "\n") {
		if strings.HasPrefix(line, `	solvers :=`) {
			inSolvers = true
		}
		if inSolvers && strings.HasPrefix(line, `	}`) {
			_, _ = fmt.Fprint(&buf, sl)
			inSolvers = false
		}
		_, _ = fmt.Fprintln(&buf, line)
		if strings.HasPrefix(line, `import (`) {
			_, _ = fmt.Fprintln(&buf, fmt.Sprintf(`"advent-of-code/day%d"`, day))
		}
	}
	if err := ioutil.WriteFile("main.go", buf.Bytes(), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join("data", fmt.Sprintf("day%d.txt", day)), []byte{}, 0644); err != nil {
		return err
	}
	if err := exec.Command("go", "fmt", "./...").Run(); err != nil {
		return err
	}
	return nil
}

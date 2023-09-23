package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jdhenke/advent-of-code/ensure-docs/ensuredocs"
)

func main() {
	var (
		year, day int
	)
	flag.IntVar(&year, "year", 0, "year to add")
	flag.IntVar(&day, "day", 0, "day to add")
	flag.Parse()
	if year == 0 {
		log.Fatal("Must provide -year flag.")
	}
	if day == 0 {
		log.Fatal("Must provide -day flag.")
	}
	session := os.Getenv("SESSION")
	if session == "" {
		log.Fatal("SESSION env var must be set.")
	}
	if err := createDay(year, day, session); err != nil {
		log.Fatal(err)
	}
	if err := ensuredocs.Ensure(year, day, 1, session); err != nil {
		log.Fatalf("Failed to ensure docs for part 1: %v", err)
	}
	log.Println("Success")
}

const libFile = `package day$DAY

import (
	"io"
)

func Part1(r io.Reader) (answer int, err error) {
	return day$DAY(r)
}

func Part2(r io.Reader) (answer int, err error) {
	return day$DAY(r)
}

func day$DAY(r io.Reader) (answer int, err error) {
	return 0, nil
}
`

const testFile = `package day$DAY_test

import (
	"github.com/jdhenke/advent-of-code/$YEAR/day$DAY"
	"github.com/jdhenke/advent-of-code/tester"
	"testing"
)

var testData = ""

func TestPart1(t *testing.T) {
	tester.New(t, day$DAY.Part1).Run(
		tester.FromString(testData).Want(0),
		tester.FromFile("input.txt").Want(0),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day$DAY.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
`

const solverLine = `		{$YEAR, $DAY, 1}: aoc$YEARday$DAY.Part1,
		{$YEAR, $DAY, 2}: aoc$YEARday$DAY.Part2,`

func sub(data string, year, day int) string {
	return strings.Replace(
		strings.Replace(data, `$DAY`, fmt.Sprint(day), -1),
		`$YEAR`,
		fmt.Sprint(year),
		-1,
	)
}

func createDay(year, day int, session string) error {
	if err := os.MkdirAll(fmt.Sprint(year), 0755); err != nil {
		return err
	}
	dir := filepath.Join(fmt.Sprint(year), fmt.Sprintf("day%d", day))
	if err := os.Mkdir(dir, 0755); err != nil {
		return err
	}
	f := sub(libFile, year, day)
	if err := ioutil.WriteFile(filepath.Join(dir, fmt.Sprintf("day%d.go", day)), []byte(f), 0644); err != nil {
		return err
	}
	tf := sub(testFile, year, day)
	if err := ioutil.WriteFile(filepath.Join(dir, fmt.Sprintf("day%d_test.go", day)), []byte(tf), 0644); err != nil {
		return err
	}
	mainBytes, err := ioutil.ReadFile("main.go")
	if err != nil {
		return err
	}
	sl := sub(solverLine, year, day)
	var buf bytes.Buffer
	var inSolvers bool
	var solverLines = strings.Split(sl, "\n")
	for _, line := range strings.Split(string(mainBytes), "\n") {
		if inSolvers && strings.HasPrefix(line, `	}`) {
			sort.Slice(solverLines, func(i, j int) bool {
				return solverNum(solverLines[i]) < solverNum(solverLines[j])
			})
			for _, l := range solverLines {
				_, _ = fmt.Fprintln(&buf, l)
			}
			inSolvers = false
		}
		if inSolvers {
			solverLines = append(solverLines, line)
			continue
		}
		if strings.HasPrefix(line, `	solutions :=`) {
			inSolvers = true
		}
		_, _ = fmt.Fprintln(&buf, line)
		if strings.HasPrefix(line, `import (`) {
			_, _ = fmt.Fprintln(&buf, fmt.Sprintf("\taoc%dday%d \"github.com/jdhenke/advent-of-code/%d/day%d\"", year, day, year, day))
		}
	}
	_, _ = buf.Write([]byte("\n"))
	if err := ioutil.WriteFile("main.go", buf.Bytes(), 0644); err != nil {
		return err
	}
	inputBytes, err := getInputBytes(year, day, session)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(fmt.Sprint(year), fmt.Sprintf("day%d", day), "input.txt"), inputBytes, 0644); err != nil {
		return err
	}
	if err := exec.Command("go", "fmt", filepath.Join(".", fmt.Sprint(year), fmt.Sprint(day))).Run(); err != nil {
		return err
	}
	return nil
}

// {2021, 1, 1}:  aoc2021day1.Part1,
var re = regexp.MustCompile(`{(\d+), (\d+), (\d+)}:.*,`)

func solverNum(line string) int {
	parts := re.FindStringSubmatch(line)
	if parts == nil {
		panic("bad solver line")
	}
	return num(parts[1])*1000 + num(parts[2])*10 + num(parts[3])
}

func num(s string) int {
	d, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return d
}

func getInputBytes(year, day int, session string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{
		Name:    "session",
		Value:   session,
		Path:    "/",
		Expires: time.Now().Add(10 * 365 * 24 * time.Hour),
	})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad status: %v", resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

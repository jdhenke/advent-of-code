package ensuredocs

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func Ensure(year, day, part int, session string) error {
	prompt, err := getPrompt(year, day, part, session)
	if err != nil {
		return err
	}
	p := filepath.Join(fmt.Sprint(year), fmt.Sprintf("day%d", day), fmt.Sprintf("day%d.go", day))
	b, err := os.ReadFile(p)
	if err != nil {
		return err
	}
	withDoc := fmt.Sprintf(`
/*
Part%d Prompt

%s*/$2`, part, prompt)
	s := regexp.MustCompile(fmt.Sprintf("(?s)(\n/\\*\nPart%d.+\\*/)?(\nfunc Part%d)", part, part)).ReplaceAllString(string(b), withDoc)
	if err := os.WriteFile(p, []byte(s), 0644); err != nil {
		return err
	}
	return nil
}

func getPrompt(year, day, part int, session string) (prompt string, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day), nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{
		Name:    "session",
		Value:   session,
		Path:    "/",
		Expires: time.Now().Add(10 * 365 * 24 * time.Hour),
	})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("bad status: %v", resp.Status)
	}
	n, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}
	t := &traveler{}
	t.traverse(n, os.Stdout, false, false)
	if part-1 >= len(t.articles) {
		return "", fmt.Errorf("could not get requested part; received fewer parts: %v", len(t.articles))
	}
	return format(t.articles[part-1]), nil
}

func format(s string) string {
	s = strings.TrimSpace(mergeAdjacentEmptyLines(s))
	out := new(bytes.Buffer)
	for _, line := range strings.Split(s, "\n") {
		_, _ = fmt.Fprintln(out, wrap(line, 80))
	}
	return out.String()
}

func wrap(s string, width int) string {
	if strings.HasPrefix(s, "    ") {
		return s
	}
	fields := strings.Fields(s)
	c := 0
	out := new(bytes.Buffer)
	for _, f := range fields {
		if c == 0 {
			_, _ = fmt.Fprint(out, f)
			c += len(f)
			continue
		}
		if c+1+len(f) >= width {
			_, _ = fmt.Fprintf(out, "\n%s", f)
			c = len(f)
			continue
		}
		_, _ = fmt.Fprintf(out, " %s", f)
		c += 1 + len(f)
	}
	return out.String()
}

func mergeAdjacentEmptyLines(s string) string {
	return regexp.MustCompile("\n[ \t]*\n[ \t]*\n+").ReplaceAllString(s, "\n\n")
}

func write(w io.Writer, s string, inCode bool) {
	if inCode {
		s = strings.Replace(s, "\n", "\n    ", -1)
	}
	_, _ = fmt.Fprint(w, s)
}

type traveler struct {
	articles []string
}

func (t *traveler) traverse(n *html.Node, w io.Writer, inArticle bool, inCode bool) {
	if n.Data == "article" {
		inArticle = true
		w = new(bytes.Buffer)
	}
	if inArticle {
		switch n.Type {
		case html.TextNode:
			write(w, n.Data, inCode)
		case html.ElementNode:
			switch n.Data {
			case "p":
				write(w, "\n", inCode)
			case "li":
				write(w, "- ", inCode)
			case "ol", "ul":
				write(w, "\n", inCode)
			case "pre":
				inCode = true
				write(w, "\n", inCode)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t.traverse(c, w, inArticle, inCode)
	}
	if n.Data == "article" {
		t.articles = append(t.articles, w.(*bytes.Buffer).String())
	}
}

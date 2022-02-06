package day19

import (
	"io"
	"strconv"
	"strings"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 19: Monster Messages ---
You land in an airport surrounded by dense forest. As you walk to your
high-speed train, the Elves at the Mythical Information Bureau contact you
again. They think their satellite has collected an image of a sea monster!
Unfortunately, the connection to the satellite is having problems, and many of
the messages sent back from the satellite have been corrupted.

They sent you a list of the rules valid messages should obey and a list of
received messages they've collected so far (your puzzle input).

The rules for valid messages (the top part of your puzzle input) are numbered
and build upon each other. For example:

    0: 1 2
    1: "a"
    2: 1 3 | 3 1
    3: "b"

Some rules, like 3: "b", simply match a single character (in this case, b).

The remaining rules list the sub-rules that must be followed; for example, the
rule 0: 1 2 means that to match rule 0, the text being checked must match rule
1, and the text after the part that matched rule 1 must then match rule 2.

Some of the rules have multiple lists of sub-rules separated by a pipe (|).
This means that at least one list of sub-rules must match. (The ones that match
might be different each time the rule is encountered.) For example, the rule 2:
1 3 | 3 1 means that to match rule 2, the text being checked must match rule 1
followed by rule 3 or it must match rule 3 followed by rule 1.

Fortunately, there are no loops in the rules, so the list of possible matches
will be finite. Since rule 1 matches a and rule 3 matches b, rule 2 matches
either ab or ba. Therefore, rule 0 matches aab or aba.

Here's a more interesting example:

    0: 4 1 5
    1: 2 3 | 3 2
    2: 4 4 | 5 5
    3: 4 5 | 5 4
    4: "a"
    5: "b"

Here, because rule 4 matches a and rule 5 matches b, rule 2 matches two letters
that are the same (aa or bb), and rule 3 matches two letters that are different
(ab or ba).

Since rule 1 matches rules 2 and 3 once each in either order, it must match two
pairs of letters, one pair with matching letters and one pair with different
letters. This leaves eight possibilities: aaab, aaba, bbab, bbba, abaa, abbb,
baaa, or babb.

Rule 0, therefore, matches a (rule 4), then any of the eight options from rule
1, then b (rule 5): aaaabb, aaabab, abbabb, abbbab, aabaab, aabbbb, abaaab, or
ababbb.

The received messages (the bottom part of your puzzle input) need to be checked
against the rules so you can determine which are valid and which are corrupted.
Including the rules and the messages together, this might look like:

    0: 4 1 5
    1: 2 3 | 3 2
    2: 4 4 | 5 5
    3: 4 5 | 5 4
    4: "a"
    5: "b"

    ababbb
    bababa
    abbbab
    aaabbb
    aaaabbb

Your goal is to determine the number of messages that completely match rule 0.
In the above example, ababbb and abbbab match, but bababa, aaabbb, and aaaabbb
do not, producing the answer 2. The whole message must match all of rule 0;
there can't be extra unmatched characters in the message. (For example, aaaabbb
might appear to match rule 0 above, but it has an extra unmatched b on the
end.)

How many messages completely match rule 0?
*/
func Part1(r io.Reader) (answer int, err error) {
	return day19(r)
}

func Part2(r io.Reader) (answer int, err error) {
	return day19(r)
}

type Rule interface {
	// Matches indicates if the rule matches the given string s starting at the beginning and how many characters it
	// matches.
	Matches(s string) (matched bool, length int)
}

func day19(r io.Reader) (answer int, err error) {
	inRules := true
	rules := make(map[int]Rule)
	if err := input.ForEachLine(r, func(line string) error {
		if line == "" {
			inRules = false
			return nil
		}
		if inRules {
			num, rule, err := parseRule(rules, line)
			if err != nil {
				return err
			}
			rules[num] = rule
			return nil
		}
		if ok, length := rules[0].Matches(line); ok && len(line) == length {
			answer++
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return answer, nil
}

func parseRule(rules map[int]Rule, line string) (int, Rule, error) {
	parts := strings.Split(line, ": ")
	num, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, nil, err
	}
	if strings.HasPrefix(parts[1], `"`) {
		return num, literalRule(parts[1][1:2]), nil
	}
	var orRules []Rule
	for _, expr := range strings.Split(parts[1], " | ") {
		var seqRules []Rule
		for _, x := range strings.Split(expr, " ") {
			n, err := strconv.Atoi(x)
			if err != nil {
				return 0, nil, err
			}
			seqRules = append(seqRules, referenceRule(rules, n))
		}
		orRules = append(orRules, seqRule(seqRules))
	}
	return num, orRule(orRules), nil
}

type ruleFunc func(s string) (bool, int)

func (f ruleFunc) Matches(s string) (bool, int) {
	return f(s)
}

func literalRule(l string) Rule {
	return ruleFunc(func(s string) (bool, int) {
		return strings.HasPrefix(s, l), len(l)
	})
}

func referenceRule(rules map[int]Rule, n int) Rule {
	return ruleFunc(func(s string) (bool, int) {
		return rules[n].Matches(s)
	})
}

func seqRule(rules []Rule) Rule {
	return ruleFunc(func(s string) (bool, int) {
		i := 0
		for _, rule := range rules {
			matched, length := rule.Matches(s[i:])
			if !matched {
				return false, 0
			}
			i += length
		}
		return true, i
	})
}

func orRule(rules []Rule) Rule {
	return ruleFunc(func(s string) (bool, int) {
		for _, rule := range rules {
			matched, length := rule.Matches(s)
			if matched {
				return true, length
			}
		}
		return false, 0
	})
}

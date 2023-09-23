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
	return day19(r, func(rules map[int]Rule) {
		// do not adjust original rules
	})
}

/*
Part2 Prompt

--- Part Two ---
As you look over the list of messages, you realize your matching rules aren't
quite right. To fix them, completely replace rules 8: 42 and 11: 42 31 with the
following:

	8: 42 | 42 8
	11: 42 31 | 42 11 31

This small change has a big impact: now, the rules do contain loops, and the
list of messages they could hypothetically match is infinite. You'll need to
determine how these changes affect which messages are valid.

Fortunately, many of the rules are unaffected by this change; it might help to
start by looking at which rules always match the same set of values and how
those rules (especially rules 42 and 31) are used by the new versions of rules
8 and 11.

(Remember, you only need to handle the rules you have; building a solution that
could handle any hypothetical combination of rules would be significantly more
difficult.)

For example:

	42: 9 14 | 10 1
	9: 14 27 | 1 26
	10: 23 14 | 28 1
	1: "a"
	11: 42 31
	5: 1 14 | 15 1
	19: 14 1 | 14 14
	12: 24 14 | 19 1
	16: 15 1 | 14 14
	31: 14 17 | 1 13
	6: 14 14 | 1 14
	2: 1 24 | 14 4
	0: 8 11
	13: 14 3 | 1 12
	15: 1 | 14
	17: 14 2 | 1 7
	23: 25 1 | 22 14
	28: 16 1
	4: 1 1
	20: 14 14 | 1 15
	3: 5 14 | 16 1
	27: 1 6 | 14 18
	14: "b"
	21: 14 1 | 1 14
	25: 1 1 | 1 14
	22: 14 14
	8: 42
	26: 14 22 | 1 20
	18: 15 15
	7: 14 5 | 1 21
	24: 14 1

	abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
	bbabbbbaabaabba
	babbbbaabbbbbabbbbbbaabaaabaaa
	aaabbbbbbaaaabaababaabababbabaaabbababababaaa
	bbbbbbbaaaabbbbaaabbabaaa
	bbbababbbbaaaaaaaabbababaaababaabab
	ababaaaaaabaaab
	ababaaaaabbbaba
	baabbaaaabbaaaababbaababb
	abbbbabbbbaaaababbbbbbaaaababb
	aaaaabbaabaaaaababaa
	aaaabbaaaabbaaa
	aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
	babaaabbbaaabaababbaabababaaab
	aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba

Without updating rules 8 and 11, these rules only match three messages:
bbabbbbaabaabba, ababaaaaaabaaab, and ababaaaaabbbaba.

However, after updating rules 8 and 11, a total of 12 messages match:

- bbabbbbaabaabba
- babbbbaabbbbbabbbbbbaabaaabaaa
- aaabbbbbbaaaabaababaabababbabaaabbababababaaa
- bbbbbbbaaaabbbbaaabbabaaa
- bbbababbbbaaaaaaaabbababaaababaabab
- ababaaaaaabaaab
- ababaaaaabbbaba
- baabbaaaabbaaaababbaababb
- abbbbabbbbaaaababbbbbbaaaababb
- aaaaabbaabaaaaababaa
- aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
- aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba

After updating rules 8 and 11, how many messages completely match rule 0?
*/
func Part2(r io.Reader) (answer int, err error) {
	return day19(r, func(rules map[int]Rule) {
		// 8: 42 | 42 8
		rules[8] = oneOrMoreRule(rules[42])
		// 11: 42 31 | 42 11 31
		rules[11] = equalPartsRule(rules[42], rules[31])
	})
}

type Rule interface {
	// Match calls f with the remaining suffix of s for every possible matching prefix of s that matches this rule,
	// stopping if f returns false or there are no more matches. If there is no match, f is never called. If there are
	// multiple matches, f may be called multiple times.
	Match(s string, f func(remaining string) (stop bool))
}

func day19(r io.Reader, update func(rules map[int]Rule)) (answer int, err error) {
	inRules := true
	rules := make(map[int]Rule)
	if err := input.ForEachLine(r, func(line string) error {
		if line == "" {
			update(rules)
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
		rules[0].Match(line, func(remaining string) (stop bool) {
			if remaining == "" {
				answer++
				return true
			}
			return false
		})
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

type ruleFunc func(s string, f func(remaining string) (stop bool))

func (r ruleFunc) Match(s string, f func(remaining string) (stop bool)) {
	r(s, f)
}

func literalRule(l string) Rule {
	return ruleFunc(func(s string, f func(remaining string) (stop bool)) {
		if strings.HasPrefix(s, l) {
			f(s[len(l):])
		}
	})
}

func referenceRule(rules map[int]Rule, n int) Rule {
	return ruleFunc(func(s string, f func(remaining string) bool) {
		rules[n].Match(s, f)
	})
}

func seqRule(rules []Rule) Rule {
	return ruleFunc(func(s string, f func(remaining string) bool) {
		// Should be when after all rules [0,i) have been called leaving the remaining string to try to be matched
		// against the ^ith rule. Declared then assigned to allow a recursive call, which is not possible otherwise.
		var helper func(i int, remaining string) bool
		helper = func(i int, remaining string) bool {
			// We've matched all rules, finally call f with the remainder after this sequence and propagate the return
			// value.
			if i >= len(rules) {
				return f(remaining)
			}
			// Otherwise, try to match the current rule calling this helper with the remaining string after this rule
			// is applied and instructed to apply the next rule, stopping if any subsequent matches indicate to do so,
			// otherwise, continue trying to match.
			stop := false
			rules[i].Match(remaining, func(nextRemaining string) bool {
				stop = helper(i+1, nextRemaining)
				return stop
			})
			return stop
		}
		rules[0].Match(s, func(remaining string) (stop bool) {
			return helper(1, remaining)
		})
	})
}

func orRule(rules []Rule) Rule {
	return ruleFunc(func(s string, f func(remaining string) (stop bool)) {
		for _, rule := range rules {
			stop := false
			rule.Match(s, func(remaining string) bool {
				stop = f(remaining)
				return stop
			})
			if stop {
				break
			}
		}
	})
}

func oneOrMoreRule(r Rule) Rule {
	return ruleFunc(func(s string, f func(remaining string) (stop bool)) {
		// Allow a recursive helper which is called after every successive match of r until it stops matching. At each
		// match, try calling f to see if the current number of matches indicates that the matching should stop,
		// otherwise continue trying.
		var helper func(remaining string) (stop bool)
		helper = func(remaining string) (stop bool) {
			if f(remaining) {
				return true
			}
			r.Match(remaining, func(nextRemaining string) (ok bool) {
				stop = helper(nextRemaining)
				return stop
			})
			return stop
		}
		r.Match(s, helper)
	})
}

// Matches any number of matches of r1 followed by the same number of matches by r2.
func equalPartsRule(r1, r2 Rule) Rule {
	return ruleFunc(func(s string, f func(remaining string) (stop bool)) {
		// Try until the prefix of r1s does not match or f indicates to stop. This is because with a rule like
		// equalParts(literalRule('a'), literalRule('b')), "aabb" wouldn't match 1 of each, but it would match 2, but
		// once 3 is tried, it won't match and will never match again, so it is safe to stop trying only after the r1s
		// fail to match.
		var r1s, r2s []Rule
		for tryMoreNums := true; tryMoreNums; {
			r1s = append(r1s, r1)
			r2s = append(r2s, r2)
			tryMoreNums = false // if r1s never matches, stop trying
			seqRule(r1s).Match(s, func(afterR1s string) (stop bool) {
				tryMoreNums = true // if r1s have matched but r2s have not, keep trying, unless a match says stop
				seqRule(r2s).Match(afterR1s, func(afterR2s string) bool {
					stop = f(afterR2s)
					if stop {
						tryMoreNums = false
					}
					return stop
				})
				return stop
			})
		}
	})
}

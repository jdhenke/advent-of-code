package day21

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 21: Allergen Assessment ---
You reach the train's last stop and the closest you can get to your vacation
island without getting wet. There aren't even any boats here, but nothing can
stop you now: you build a raft. You just need a few days' worth of food for
your journey.

You don't speak the local language, so you can't read any ingredients lists.
However, sometimes, allergens are listed in a language you do understand. You
should be able to use this information to determine which ingredient contains
which allergen and work out which foods are safe to take with you on your trip.

You start by compiling a list of foods (your puzzle input), one food per line.
Each line includes that food's ingredients list followed by some or all of the
allergens the food contains.

Each allergen is found in exactly one ingredient. Each ingredient contains zero
or one allergen. Allergens aren't always marked; when they're listed (as in
(contains nuts, shellfish) after an ingredients list), the ingredient that
contains each listed allergen will be somewhere in the corresponding
ingredients list. However, even if an allergen isn't listed, the ingredient
that contains that allergen could still be present: maybe they forgot to label
it, or maybe it was labeled in a language you don't know.

For example, consider the following list of foods:

    mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
    trh fvjkl sbzzf mxmxvkd (contains dairy)
    sqjhc fvjkl (contains soy)
    sqjhc mxmxvkd sbzzf (contains fish)

The first food in the list has four ingredients (written in a language you
don't understand): mxmxvkd, kfcds, sqjhc, and nhms. While the food might
contain other allergens, a few allergens the food definitely contains are
listed afterward: dairy and fish.

The first step is to determine which ingredients can't possibly contain any of
the allergens in any food in your list. In the above example, none of the
ingredients kfcds, nhms, sbzzf, or trh can contain an allergen. Counting the
number of times any of these ingredients appear in any ingredients list
produces 5: they all appear once each except sbzzf, which appears twice.

Determine which ingredients cannot possibly contain any of the allergens in
your list. How many times do any of those ingredients appear?
*/
func Part1(r io.Reader) (answer int, err error) {
	counts, solution, err := day21(r)
	for ing, count := range counts {
		if _, ok := solution[ing]; !ok {
			answer += count
		}
	}
	return answer, nil
}

/*
Part2 Prompt

--- Part Two ---
Now that you've isolated the inert ingredients, you should have enough
information to figure out which ingredient contains which allergen.

In the above example:

- mxmxvkd contains dairy.
- sqjhc contains fish.
- fvjkl contains soy.

Arrange the ingredients alphabetically by their allergen and separate them by
commas to produce your canonical dangerous ingredient list. (There should not
be any spaces in your canonical dangerous ingredient list.) In the above
example, this would be mxmxvkd,sqjhc,fvjkl.

Time to stock your raft with supplies. What is your canonical dangerous
ingredient list?
*/
func Part2(r io.Reader) (answer int, err error) {
	_, solution, err := day21(r)
	type pair struct {
		ingredient
		allergen
	}
	var pairs []pair
	for ing, al := range solution {
		pairs = append(pairs, pair{
			ingredient: ing,
			allergen:   al,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].allergen < pairs[j].allergen
	})
	var ings []string
	for _, p := range pairs {
		ings = append(ings, string(p.ingredient))
	}
	fmt.Println(strings.Join(ings, ","))
	return 0, nil
}

type allergen string

type ingredient string

type allergenOptions map[allergen]map[ingredient]bool

func (m allergenOptions) markSolved(al allergen, ing ingredient) {
	for otherAl, optsForOtherAl := range m {
		if otherAl == al {
			continue
		}
		if _, ok := optsForOtherAl[ing]; ok {
			delete(optsForOtherAl, ing)
			if len(optsForOtherAl) == 1 {
				m.markSolved(otherAl, only(optsForOtherAl))
			}
		}
	}
}

func day21(r io.Reader) (counts map[ingredient]int, solution map[ingredient]allergen, err error) {
	options := make(allergenOptions)
	counts = make(map[ingredient]int)
	if err := input.ForEachLine(r, func(line string) error {
		ingredients, allergens := parse(line)
		for _, ing := range ingredients {
			counts[ing]++
		}
		for _, al := range allergens {
			if optsForAl, ok := options[al]; ok {
				union(optsForAl, ingredients)
				if len(optsForAl) == 1 {
					options.markSolved(al, only(optsForAl))
				}
			} else {
				options[al] = set(ingredients)
			}
		}
		return nil
	}); err != nil {
		return nil, nil, err
	}
	solution = make(map[ingredient]allergen)
	for al, optsForAl := range options {
		if len(optsForAl) == 1 {
			solution[only(optsForAl)] = al
		} else {
			fmt.Printf("Unsolved: %v %v\n", al, optsForAl)
			panic("Unsolved")
		}
	}
	return counts, solution, nil
}

func set(ingredients []ingredient) map[ingredient]bool {
	out := make(map[ingredient]bool)
	for _, ing := range ingredients {
		out[ing] = true
	}
	return out
}

func only(optsForAl map[ingredient]bool) ingredient {
	for ing := range optsForAl {
		return ing
	}
	panic("no option found")
}

func union(optsForAl map[ingredient]bool, ingredients []ingredient) {
	other := set(ingredients)
	for k := range optsForAl {
		if _, ok := other[k]; !ok {
			delete(optsForAl, k)
		}
	}
}

// sqjhc mxmxvkd sbzzf (contains fish)

var re = regexp.MustCompile(`(.*) \(contains (.*)\)`)

func parse(line string) (ingredients []ingredient, allergens []allergen) {
	match := re.FindStringSubmatch(line)
	ingStrs := strings.Split(match[1], " ")
	alStrs := strings.Split(match[2], ", ")
	for _, ing := range ingStrs {
		ingredients = append(ingredients, ingredient(ing))
	}
	for _, al := range alStrs {
		allergens = append(allergens, allergen(al))
	}
	return ingredients, allergens
}

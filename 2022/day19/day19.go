package day19

import (
	"fmt"
	"github.com/jdhenke/advent-of-code/input"
	"io"
)

/*
Part1 Prompt

--- Day 19: Not Enough Minerals ---
Your scans show that the lava did indeed form obsidian!

The wind has changed direction enough to stop sending lava droplets toward you,
so you and the elephants exit the cave. As you do, you notice a collection of
geodes around the pond. Perhaps you could use the obsidian to create some
geode-cracking robots and break them open?

To collect the obsidian from the bottom of the pond, you'll need waterproof
obsidian-collecting robots. Fortunately, there is an abundant amount of clay
nearby that you can use to make them waterproof.

In order to harvest the clay, you'll need special-purpose clay-collecting
robots. To make any type of robot, you'll need ore, which is also plentiful but
in the opposite direction from the clay.

Collecting ore requires ore-collecting robots with big drills. Fortunately, you
have exactly one ore-collecting robot in your pack that you can use to
kickstart the whole operation.

Each robot can collect 1 of its resource type per minute. It also takes one
minute for the robot factory (also conveniently from your pack) to construct
any type of robot, although it consumes the necessary resources available when
construction begins.

The robot factory has many blueprints (your puzzle input) you can choose from,
but once you've configured it with a blueprint, you can't change it. You'll
need to work out which blueprint is best.

For example:

	Blueprint 1:
	  Each ore robot costs 4 ore.
	  Each clay robot costs 2 ore.
	  Each obsidian robot costs 3 ore and 14 clay.
	  Each geode robot costs 2 ore and 7 obsidian.

	Blueprint 2:
	  Each ore robot costs 2 ore.
	  Each clay robot costs 3 ore.
	  Each obsidian robot costs 3 ore and 8 clay.
	  Each geode robot costs 3 ore and 12 obsidian.

(Blueprints have been line-wrapped here for legibility. The robot factory's
actual assortment of blueprints are provided one blueprint per line.)

The elephants are starting to look hungry, so you shouldn't take too long; you
need to figure out which blueprint would maximize the number of opened geodes
after 24 minutes by figuring out which robots to build and when to build them.

Using blueprint 1 in the example above, the largest number of geodes you could
open in 24 minutes is 9. One way to achieve that is:

	== Minute 1 ==
	1 ore-collecting robot collects 1 ore; you now have 1 ore.

	== Minute 2 ==
	1 ore-collecting robot collects 1 ore; you now have 2 ore.

	== Minute 3 ==
	Spend 2 ore to start building a clay-collecting robot.
	1 ore-collecting robot collects 1 ore; you now have 1 ore.
	The new clay-collecting robot is ready; you now have 1 of them.

	== Minute 4 ==
	1 ore-collecting robot collects 1 ore; you now have 2 ore.
	1 clay-collecting robot collects 1 clay; you now have 1 clay.

	== Minute 5 ==
	Spend 2 ore to start building a clay-collecting robot.
	1 ore-collecting robot collects 1 ore; you now have 1 ore.
	1 clay-collecting robot collects 1 clay; you now have 2 clay.
	The new clay-collecting robot is ready; you now have 2 of them.

	== Minute 6 ==
	1 ore-collecting robot collects 1 ore; you now have 2 ore.
	2 clay-collecting robots collect 2 clay; you now have 4 clay.

	== Minute 7 ==
	Spend 2 ore to start building a clay-collecting robot.
	1 ore-collecting robot collects 1 ore; you now have 1 ore.
	2 clay-collecting robots collect 2 clay; you now have 6 clay.
	The new clay-collecting robot is ready; you now have 3 of them.

	== Minute 8 ==
	1 ore-collecting robot collects 1 ore; you now have 2 ore.
	3 clay-collecting robots collect 3 clay; you now have 9 clay.

	== Minute 9 ==
	1 ore-collecting robot collects 1 ore; you now have 3 ore.
	3 clay-collecting robots collect 3 clay; you now have 12 clay.

	== Minute 10 ==
	1 ore-collecting robot collects 1 ore; you now have 4 ore.
	3 clay-collecting robots collect 3 clay; you now have 15 clay.

	== Minute 11 ==
	Spend 3 ore and 14 clay to start building an obsidian-collecting robot.
	1 ore-collecting robot collects 1 ore; you now have 2 ore.
	3 clay-collecting robots collect 3 clay; you now have 4 clay.
	The new obsidian-collecting robot is ready; you now have 1 of them.

	== Minute 12 ==
	Spend 2 ore to start building a clay-collecting robot.
	1 ore-collecting robot collects 1 ore; you now have 1 ore.
	3 clay-collecting robots collect 3 clay; you now have 7 clay.
	1 obsidian-collecting robot collects 1 obsidian; you now have 1 obsidian.
	The new clay-collecting robot is ready; you now have 4 of them.

	== Minute 13 ==
	1 ore-collecting robot collects 1 ore; you now have 2 ore.
	4 clay-collecting robots collect 4 clay; you now have 11 clay.
	1 obsidian-collecting robot collects 1 obsidian; you now have 2 obsidian.

	== Minute 14 ==
	1 ore-collecting robot collects 1 ore; you now have 3 ore.
	4 clay-collecting robots collect 4 clay; you now have 15 clay.
	1 obsidian-collecting robot collects 1 obsidian; you now have 3 obsidian.

	== Minute 15 ==
	Spend 3 ore and 14 clay to start building an obsidian-collecting robot.
	1 ore-collecting robot collects 1 ore; you now have 1 ore.
	4 clay-collecting robots collect 4 clay; you now have 5 clay.
	1 obsidian-collecting robot collects 1 obsidian; you now have 4 obsidian.
	The new obsidian-collecting robot is ready; you now have 2 of them.

	== Minute 16 ==
	1 ore-collecting robot collects 1 ore; you now have 2 ore.
	4 clay-collecting robots collect 4 clay; you now have 9 clay.
	2 obsidian-collecting robots collect 2 obsidian; you now have 6 obsidian.

	== Minute 17 ==
	1 ore-collecting robot collects 1 ore; you now have 3 ore.
	4 clay-collecting robots collect 4 clay; you now have 13 clay.
	2 obsidian-collecting robots collect 2 obsidian; you now have 8 obsidian.

	== Minute 18 ==
	Spend 2 ore and 7 obsidian to start building a geode-cracking robot.
	1 ore-collecting robot collects 1 ore; you now have 2 ore.
	4 clay-collecting robots collect 4 clay; you now have 17 clay.
	2 obsidian-collecting robots collect 2 obsidian; you now have 3 obsidian.
	The new geode-cracking robot is ready; you now have 1 of them.

	== Minute 19 ==
	1 ore-collecting robot collects 1 ore; you now have 3 ore.
	4 clay-collecting robots collect 4 clay; you now have 21 clay.
	2 obsidian-collecting robots collect 2 obsidian; you now have 5 obsidian.
	1 geode-cracking robot cracks 1 geode; you now have 1 open geode.

	== Minute 20 ==
	1 ore-collecting robot collects 1 ore; you now have 4 ore.
	4 clay-collecting robots collect 4 clay; you now have 25 clay.
	2 obsidian-collecting robots collect 2 obsidian; you now have 7 obsidian.
	1 geode-cracking robot cracks 1 geode; you now have 2 open geodes.

	== Minute 21 ==
	Spend 2 ore and 7 obsidian to start building a geode-cracking robot.
	1 ore-collecting robot collects 1 ore; you now have 3 ore.
	4 clay-collecting robots collect 4 clay; you now have 29 clay.
	2 obsidian-collecting robots collect 2 obsidian; you now have 2 obsidian.
	1 geode-cracking robot cracks 1 geode; you now have 3 open geodes.
	The new geode-cracking robot is ready; you now have 2 of them.

	== Minute 22 ==
	1 ore-collecting robot collects 1 ore; you now have 4 ore.
	4 clay-collecting robots collect 4 clay; you now have 33 clay.
	2 obsidian-collecting robots collect 2 obsidian; you now have 4 obsidian.
	2 geode-cracking robots crack 2 geodes; you now have 5 open geodes.

	== Minute 23 ==
	1 ore-collecting robot collects 1 ore; you now have 5 ore.
	4 clay-collecting robots collect 4 clay; you now have 37 clay.
	2 obsidian-collecting robots collect 2 obsidian; you now have 6 obsidian.
	2 geode-cracking robots crack 2 geodes; you now have 7 open geodes.

	== Minute 24 ==
	1 ore-collecting robot collects 1 ore; you now have 6 ore.
	4 clay-collecting robots collect 4 clay; you now have 41 clay.
	2 obsidian-collecting robots collect 2 obsidian; you now have 8 obsidian.
	2 geode-cracking robots crack 2 geodes; you now have 9 open geodes.

However, by using blueprint 2 in the example above, you could do even better:
the largest number of geodes you could open in 24 minutes is 12.

Determine the quality level of each blueprint by multiplying that blueprint's
ID number with the largest number of geodes that can be opened in 24 minutes
using that blueprint. In this example, the first blueprint has ID 1 and can
open 9 geodes, so its quality level is 9. The second blueprint has ID 2 and can
open 12 geodes, so its quality level is 24. Finally, if you add up the quality
levels of all of the blueprints in the list, you get 33.

Determine the quality level of each blueprint using the largest number of
geodes it could produce in 24 minutes. What do you get if you add up the
quality level of all of the blueprints in your list?
*/
func Part1(r io.Reader) (answer int, err error) {
	geodes, err := day19(r, -1, 24)
	if err != nil {
		return 0, err
	}
	for i, g := range geodes {
		answer += (i + 1) * g
	}
	return answer, nil
}

func Part2(r io.Reader) (answer int, err error) {
	geodes, err := day19(r, 3, 32)
	if err != nil {
		return 0, err
	}
	answer = 1
	for _, g := range geodes {
		answer *= g
	}
	return answer, nil
}

const tmpl = `Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.`

func day19(r io.Reader, numBlueprints int, steps int) (answers []int, err error) {
	i := 0
	if err := input.ForEachLine(r, func(line string) error {
		if numBlueprints > 0 && i >= numBlueprints {
			return nil
		}
		i++
		b := blueprint{}
		if _, err := fmt.Sscanf(line, tmpl, &b.id, &b.oreRobotCost.ore, &b.clayRobotCost.ore, &b.obsidianRobotCost.ore, &b.obsidianRobotCost.clay, &b.geodeRobotCost.ore, &b.geodeRobotCost.obsidian); err != nil {
			return err
		}
		memo := make(map[state]*int)
		m := maxGeodes(memo, b, state{oreRobots: 1, steps: steps}, -1)
		if m == nil {
			m = ptr(0)
		}
		fmt.Printf("ID %d max %d memo %d\n", b.id, *m, len(memo))
		answers = append(answers, *m)
		return nil
	}); err != nil {
		return nil, err
	}
	return answers, nil
}

func ptr(i int) *int {
	return &i
}

func maxGeodes(memo map[state]*int, b blueprint, s state, alt int) (ans *int) {
	//fmt.Println(s)
	// base case
	if s.steps == 0 {
		return ptr(0)
	}

	// only traverse this path once
	if ans, ok := memo[s]; ok {
		return ans
	}
	defer func(s state) {
		memo[s] = ans
	}(s)

	// if the alternative is impossible to beat, bail now
	if ceiling(s) < alt {
		return nil
	}

	// otherwise choose the best option
	nextStates := getNextStates(b, s) // note: these may have take different numbers of steps
	found := false
	for _, next := range nextStates {
		nextSteps := s.steps - next.steps
		earned := nextSteps * s.geodeRobots
		if nextMax := maxGeodes(memo, b, next, alt-earned); nextMax != nil && *nextMax+earned > alt {
			found = true
			alt = *nextMax + earned
		}
	}
	if !found {
		return nil
	}
	return &alt
}

func getNextStates(b blueprint, prev state) []state {
	var out []state

	// what's next robot to make?
	for _, c := range []struct {
		check func(s state) bool
		cost  cost
		buy   func(p *state)
	}{
		{
			check: func(s state) bool {
				return s.obsidianRobots > 0
			},
			cost: b.geodeRobotCost,
			buy: func(p *state) {
				p.geodeRobots++
			},
		},
		{
			check: func(s state) bool {
				return s.clayRobots > 0
			},
			cost: b.obsidianRobotCost,
			buy: func(p *state) {
				p.obsidianRobots++
			},
		},
		{
			check: func(s state) bool {
				return true
			},
			cost: b.clayRobotCost,
			buy: func(p *state) {
				p.clayRobots++
			},
		},
		{
			check: func(s state) bool {
				return true
			},
			cost: b.oreRobotCost,
			buy: func(p *state) {
				p.oreRobots++
			},
		},
	} {
		s := prev
		if !c.check(s) {
			continue
		}
		for !s.afford(c.cost) {
			s = s.tick()
		}
		s = s.sub(c.cost)
		s = s.tick()
		c.buy(&s)
		if s.steps >= 0 {
			out = append(out, s)
		}
	}

	hold := prev
	hold.steps = 0
	out = append(out, hold)
	//fmt.Print(prev, out)
	//fmt.Scanln()
	return out
}

func ceiling(s state) int {
	// guaranteed have, will have, and what will be added if every step from here on out gets a robot
	return s.steps*s.geodeRobots + (s.steps*s.steps+s.steps)/2
}

type state struct {
	steps                                              int
	ore, clay, obsidian                                int
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
}

func (s state) afford(c cost) bool {
	return s.ore >= c.ore && s.clay >= c.clay && s.obsidian >= c.obsidian
}

func (s state) sub(c cost) state {
	s.ore -= c.ore
	s.clay -= c.clay
	s.obsidian -= c.obsidian
	return s
}

func (s state) tick() state {
	s.ore += s.oreRobots
	s.clay += s.clayRobots
	s.obsidian += s.obsidianRobots
	s.steps -= 1
	return s
}

type cost struct {
	ore, clay, obsidian int
}

type blueprint struct {
	id                                                             int
	oreRobotCost, clayRobotCost, obsidianRobotCost, geodeRobotCost cost
}

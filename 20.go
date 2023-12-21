package main

import (
	"log"
	"strings"
)

type module struct {
	variation     rune
	input, output []string
	stateFF       bool
	stateCon      map[string]int
}
type pulse struct {
	pulseType int
	targets   []string
	source    string
}

func day20(c chan string) {
	var low, high int
	modules := make(map[string]module)
	for line := range c {
		if len(line) == 0 {
			continue
		}

		moduleSection, outputSection := splitString(line, " -> ")
		module := module{}
		var tag string
		if moduleSection == "broadcaster" {
			module.variation = 'b'
			tag = "broadcaster"
		} else {
			module.variation = rune(moduleSection[:1][0])
			tag = moduleSection[1:]
			if module.variation == '&' {
				module.stateCon = make(map[string]int)
			}
		}
		module.output = strings.Split(outputSection, ", ")
		modules[tag] = module
	}

	for tag := range modules {
		for _, out := range modules[tag].output {
			other := modules[out]
			other.input = append(modules[out].input, tag)
			modules[out] = other
		}
	}

	for buttonPress := 1; buttonPress <= 1000; buttonPress++ {
		var pulses, nextPulses []pulse
		nextPulses = append(nextPulses, pulse{0, []string{"broadcaster"}, "button"})
		for {
			pulses = nextPulses
			nextPulses = nil

			for _, p := range pulses {
				if p.pulseType == 0 {
					low += len(p.targets)
				} else {
					high += len(p.targets)
				}
			}
			if len(pulses) == 0 {
				break
			}
			for _, pulse := range pulses {
				nextPulses = append(nextPulses, processPulse(pulse, &modules)...)
			}
		}
	}

	log.Printf("\nLow: %d\nHigh: %d", low, high)
	log.Printf("20 Total: %d", low*high)
}

func processPulse(ps pulse, modules *map[string]module) []pulse {
	var newPulses []pulse
	for _, tgt := range ps.targets {
		target := (*modules)[tgt]
		switch target.variation {
		case 'b':
			newPulses = append(newPulses, pulse{ps.pulseType, target.output, tgt})
		case '%':
			if ps.pulseType == 0 {
				target.stateFF = !target.stateFF
				if target.stateFF {
					newPulses = append(newPulses, pulse{1, target.output, tgt})
				} else {
					newPulses = append(newPulses, pulse{0, target.output, tgt})
				}
			}
		case '&':
			target.stateCon[ps.source] = ps.pulseType
			allHigh := true
			for _, input := range target.input {
				if target.stateCon[input] == 0 {
					allHigh = false
					break
				}
			}
			if allHigh {
				newPulses = append(newPulses, pulse{0, target.output, tgt})
			} else {
				newPulses = append(newPulses, pulse{1, target.output, tgt})
			}
		}
		(*modules)[tgt] = target
	}
	return newPulses
}

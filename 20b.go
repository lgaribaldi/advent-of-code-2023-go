package main

import (
	"log"
	"strings"
)

func day20b(c chan string) {
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

buttonLoop:
	for buttonPress := 1; buttonPress <= 100000000; buttonPress++ {
		var pulses, nextPulses []pulse
		nextPulses = append(nextPulses, pulse{0, []string{"broadcaster"}, "button"})
		for {
			pulses = nextPulses
			nextPulses = nil

			for _, p := range pulses {
				if p.pulseType == 0 {
					for _, t := range p.targets {
						/*
							Find the occurence of each of the Conjunction modules that feed into rx
							The answer is the LCM of each of these
						*/
						switch t {
						case "rx":
							log.Printf("Presses: %d", buttonPress)
							break buttonLoop
						case "pq":
							log.Printf("pq Presses: %d", buttonPress)
						case "fg":
							log.Printf("fg Presses: %d", buttonPress)
						case "dk":
							log.Printf("dk Presses: %d", buttonPress)
						case "fm":
							log.Printf("fm Presses: %d", buttonPress)
						}
					}
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
}

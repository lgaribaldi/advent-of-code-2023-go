package main

import (
	"log"
	"strconv"
	"strings"
)

func day19b(c chan string) {
	var total int64
	workflows := make(map[string]workflow)
	partsSection := false
	for line := range c {
		if len(line) == 0 {
			partsSection = true
			continue
		}
		if !partsSection {
			workflowTag, ruleSection := splitString(line, "{")
			rules := strings.Split(ruleSection, ",")
			workflow := workflow{make([]rune, 0), make([]rune, 0), make([]int, 0), make([]string, 0)}
			for i := 0; i < len(rules)-1; i++ {
				workflow.property = append(workflow.property, rune(rules[i][:1][0]))
				workflow.operator = append(workflow.operator, rune(rules[i][1:2][0]))
				operation, destination := splitString(rules[i], ":")
				value, err := strconv.ParseInt(operation[2:], 10, 32)
				if err != nil {
					log.Fatal("invalid input")
				}
				workflow.value = append(workflow.value, int(value))
				workflow.destination = append(workflow.destination, destination)
			}
			// add default destination
			workflow.property = append(workflow.property, 'x')
			workflow.operator = append(workflow.operator, '>')
			workflow.value = append(workflow.value, -1)
			workflow.destination = append(workflow.destination, rules[len(rules)-1][:len(rules[len(rules)-1])-1])

			workflows[workflowTag] = workflow
		} else {
			break
		}
	}

	workFlowRange("in", []int{1, 1, 1, 1}, []int{4000, 4000, 4000, 4000}, &workflows, &total)
	log.Printf("19B Total: %d", total)
}

func workFlowRange(tag string, min, max []int, workflows *map[string]workflow, total *int64) {
	wf := (*workflows)[tag]
	for i := range wf.property {
		var propertyIndex int
		switch wf.property[i] {
		case 'x':
			propertyIndex = 0
		case 'm':
			propertyIndex = 1
		case 'a':
			propertyIndex = 2
		case 's':
			propertyIndex = 3
		}
		passed := false
		if wf.operator[i] == '>' && max[propertyIndex] > wf.value[i] {
			// checks from same min to prev min
			if min[propertyIndex] <= wf.value[i] {
				newMax := make([]int, 4)
				copy(newMax, max)
				newMax[propertyIndex] = wf.value[i]
				newMin := make([]int, 4)
				copy(newMin, min)
				workFlowRange(tag, newMin, newMax, workflows, total)
				// continue with a new increased min
				min[propertyIndex] = wf.value[i] + 1
			}
			passed = true
		}

		if wf.operator[i] == '<' && min[propertyIndex] < wf.value[i] {
			// checks from same max to prev max
			if max[propertyIndex] >= wf.value[i] {
				newMin := make([]int, 4)
				copy(newMin, min)
				newMin[propertyIndex] = wf.value[i]
				newMax := make([]int, 4)
				copy(newMax, max)
				workFlowRange(tag, newMin, newMax, workflows, total)
				// continue with a new decreased max
				max[propertyIndex] = wf.value[i] - 1
			}
			passed = true
		}

		if passed {
			if wf.destination[i] == "R" {
				println("Rejected")
				log.Printf("min: %d\nmax: %d", min, max)
				return
			}
			if wf.destination[i] == "A" {
				*total += int64((max[0] - min[0] + 1) * (max[1] - min[1] + 1) * (max[2] - min[2] + 1) * (max[3] - min[3] + 1))
				println("Passed:", *total)
				log.Printf("min: %d\nmax: %d", min, max)
				return
			}
			workFlowRange(wf.destination[i], min, max, workflows, total)
			return
		}
	}
}

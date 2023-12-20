package main

import (
	"log"
	"strconv"
	"strings"
)

type workflow struct {
	property, operator []rune
	value              []int
	destination        []string
}

func day19(c chan string) {
	var total int
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
			parts := strings.Split(line[1:len(line)-1], ",")
			part := make([]int, 0)
			for _, v := range parts {
				_, value := splitString(v, "=")
				intValue, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					log.Fatal("invalid input")
				}
				part = append(part, int(intValue))
			}

			currWorflow := workflows["in"]
		nextWorkflow:
			for {
				for i := range currWorflow.property {
					var propertyIndex int
					switch currWorflow.property[i] {
					case 'x':
						propertyIndex = 0
					case 'm':
						propertyIndex = 1
					case 'a':
						propertyIndex = 2
					case 's':
						propertyIndex = 3
					}
					if (currWorflow.operator[i] == '>' && part[propertyIndex] > currWorflow.value[i]) || (currWorflow.operator[i] == '<' && part[propertyIndex] < currWorflow.value[i]) {
						if currWorflow.destination[i] == "R" {
							break nextWorkflow
						}
						if currWorflow.destination[i] == "A" {
							total += part[0] + part[1] + part[2] + part[3]
							break nextWorkflow
						}
						currWorflow = workflows[currWorflow.destination[i]]
						continue nextWorkflow
					}

				}
			}
		}
	}
	log.Printf("19A Total: %d", total)
}

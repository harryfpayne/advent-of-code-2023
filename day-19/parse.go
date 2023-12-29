package main

import (
	"strconv"
	"strings"
)

func Parse(str string) (map[string]Instruction, []Part) {
	str = strings.TrimSpace(str)
	lr := strings.Split(str, "\n\n")
	instructionsStr, partStr := lr[0], lr[1]

	var instructions []Instruction
	for _, instructionStr := range strings.Split(instructionsStr, "\n") {
		label, rules := strings.Split(instructionStr, "{")[0], strings.Split(instructionStr, "{")[1]
		rules = rules[:len(rules)-1]

		var instruction Instruction
		instruction.name = label
		rulesStrs := strings.Split(rules, ",")
		for i, s := range rulesStrs {
			var rule Rule
			if i == len(rulesStrs)-1 {
				instruction.fallback = s
				continue
			}

			rule.key = s[0]
			rule.comparator = s[1] == '<'
			rule.value, _ = strconv.Atoi(strings.Split(s, ":")[0][2:])
			rule.destination = strings.Split(s, ":")[1]
			instruction.rules = append(instruction.rules, rule)
		}
		instructions = append(instructions, instruction)
	}

	var parts []Part
	for _, s := range strings.Split(partStr, "\n") {
		s := s[1 : len(s)-1]
		ps := strings.Split(s, ",")
		var part Part
		part.x, _ = strconv.Atoi(strings.Split(ps[0], "=")[1])
		part.m, _ = strconv.Atoi(strings.Split(ps[1], "=")[1])
		part.a, _ = strconv.Atoi(strings.Split(ps[2], "=")[1])
		part.s, _ = strconv.Atoi(strings.Split(ps[3], "=")[1])
		parts = append(parts, part)
	}

	instructionMap := make(map[string]Instruction)
	for _, instruction := range instructions {
		instruction := instruction
		instructionMap[instruction.name] = instruction
	}

	return instructionMap, parts
}

func part1(str string) int {
	instructions, parts := Parse(str)

	var accepted []Part
	var rejected []Part

	for _, part := range parts {
		workflowKey := "in"
		workflowRuleIndex := 0
		for {
			workflow := instructions[workflowKey]
			rule := workflow.rules[workflowRuleIndex]
			if part.MeetsRequirement(rule) {
				workflowKey = rule.destination
				workflowRuleIndex = 0
			} else {
				workflowRuleIndex++
				if workflowRuleIndex == len(workflow.rules) {
					workflowKey = workflow.fallback
					workflowRuleIndex = 0
				}
			}

			if workflowKey == "A" {
				accepted = append(accepted, part)
				break
			}
			if workflowKey == "R" {
				rejected = append(rejected, part)
				break
			}
		}
	}

	total := 0
	for _, part := range accepted {
		total += part.Sum()
	}
	return total
}

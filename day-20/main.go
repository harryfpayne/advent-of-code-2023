package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	modules := Parse(PUZZLE)

	type QueueItem struct {
		pulse  Pulse
		module string
		from   string
	}

	rxInput := modules["rx"].(*Broadcaster).inputs[0]       // rx defaults to a broadcaster
	rxInputInputs := modules[rxInput].(*Conjunction).inputs // Assuming based on my input that rxInput is a conjunction

	loopInfo := make(map[string]int)

	lowCount := 0
	highCount := 0
	for j := 0; len(loopInfo) < len(rxInputInputs); j++ {
		queue := []QueueItem{
			{LowPulse, "broadcaster", "button"},
		}

		for len(queue) > 0 {
			i := queue[0]
			queue = queue[1:]

			pulse, ok := modules[i.module].Process(i.from, i.pulse)
			if !ok {
				continue
			}
			for _, output := range modules[i.module].GetOutputs() {
				queue = append(queue, QueueItem{pulse, output, i.module})
			}

			// Calculate part 1
			if j < 1000 { // Part 2 requires going higher so part 1 can ignore
				if i.pulse == LowPulse {
					lowCount++
				} else {
					highCount++
				}
			}

			// Part 2
			// The inputs to rxInput are each on their own loop, store when they start looping then find LCM
			rxInputModule := modules[rxInput].(*Conjunction)
			for _, k := range rxInputInputs {
				if _, ok := loopInfo[k]; !ok && rxInputModule.inputMemory[k] == HighPulse {
					loopInfo[k] = j + 1
				}
			}
		}
	}

	fmt.Println("part 1:", lowCount*highCount)

	product := 1
	for _, v := range loopInfo { // LCM, can only do this because they're all prime
		product *= v
	}
	fmt.Println("part 2:", product)

	fmt.Println("time:", time.Since(start))
}

func Parse(str string) map[string]IBroadcaster {
	str = strings.TrimSpace(str)
	modules := make(map[string]IBroadcaster)

	for _, line := range strings.Split(str, "\n") {
		parts := strings.Split(line, " -> ")
		label := parts[0]
		partType := label[0]
		label = label[1:]
		outputs := strings.Split(parts[1], ", ")

		switch partType {
		case '%':
			modules[label] = NewFlipFlop(label, []string{}, outputs)
		case '&':
			modules[label] = NewConjunction(label, []string{}, outputs)
		default:
			modules[parts[0]] = NewBroadcaster(parts[0], []string{}, outputs)
		}
	}

	for _, m := range modules {
		for _, output := range m.GetOutputs() {
			if _, ok := modules[output]; !ok {
				modules[output] = NewBroadcaster(output, []string{}, []string{})
			}
			modules[output].AddInput(m.GetLabel())
		}
	}

	return modules
}

const TEST = `
broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

const TEST2 = `
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

const PUZZLE = `
%fl -> tf, gz
%xb -> hl, tl
%mq -> tf, fl
%px -> hl, tm
%dp -> xv
broadcaster -> js, ng, lb, gr
&ql -> rx
%gk -> hm
%vp -> vf, sn
%fp -> xb
&lr -> ss, rm, dc, js, gk, dp, bq
%xl -> gx, lr
%xx -> hb
%cb -> jg
&hl -> nj, lb, tl, xx, hb, fp, mf
%vr -> tf, hq
%bq -> gk
%jg -> qn
%hb -> qk
%qk -> hs, hl
%gz -> tf
%rm -> hj
&tf -> cb, jg, fz, gr, zj, qn, kb
%qn -> td
%js -> lr, dc
%qb -> nc
%zj -> vr
%td -> tf, zj
%tl -> kg
%gx -> lr
%hm -> lr, rd
&fh -> ql
%nj -> xx
%hq -> kb, tf
%kg -> px, hl
%dc -> dp
%vf -> th, sn
&mf -> ql
%tm -> hl
&fz -> ql
%xd -> tn, sn
%ng -> vp, sn
%th -> qb
%rd -> xl, lr
%bt -> xd, sn
%tv -> sn
%nl -> bt
%hs -> fp, hl
%xv -> rm, lr
%tn -> sn, tv
%hj -> lr, bq
&ss -> ql
%sd -> nl
&sn -> sd, fh, th, qb, nl, ng, nc
%kb -> mq
%lb -> nj, hl
%gr -> tf, cb
%nc -> sd`

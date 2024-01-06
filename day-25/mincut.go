package main

type ContractedComponent struct {
	*Component
	contractionGroup int
}

func MinCut(_components map[string]Component) [][2]string {
	// Pick 2 nodes randomly that are not in the same contraction group
	// contract those nodes
	// check how many groups are left

	edgesChan := make(chan [][2]string)
	workerChan := make(chan struct{}, 8)

	go func() {
		for {
			workerChan <- struct{}{}
			go func() {
				components := make(map[string]*ContractedComponent)
				for _, component := range _components {
					component := component
					components[component.Name] = &ContractedComponent{
						Component:        &component,
						contractionGroup: len(components),
					}
				}

				for numberOfGroups(components) > 2 {
					a, b := pickNodes(components)
					contractEdge(components, a, b)
				}

				edges := remainingEdges(components)
				if len(edges) == 3 {
					edgesChan <- edges
				}
				<-workerChan
			}()
		}
	}()

	edges := <-edgesChan
	return edges
}

func pickNodes(components map[string]*ContractedComponent) (string, string) {
	for s, component := range components {
		for s2, _ := range component.Connections {
			if component.contractionGroup != components[s2].contractionGroup {
				return s, s2
			}
		}
	}
	panic("not enough groups")
}

func contractEdge(components map[string]*ContractedComponent, a, b string) {
	aNode := components[a]
	bNode := components[b]
	group := min(aNode.contractionGroup, bNode.contractionGroup)
	aNode.contractionGroup = group
	bNode.contractionGroup = group
}

func numberOfGroups(components map[string]*ContractedComponent) int {
	seen := make(map[int]struct{})
	for _, component := range components {
		seen[component.contractionGroup] = struct{}{}
	}
	return len(seen)
}

func remainingEdges(components map[string]*ContractedComponent) [][2]string {
	var edges [][2]string

	for _, component := range components {
		for connection, _ := range component.Connections {
			if component.Name > connection {
				continue
			}
			if component.contractionGroup == components[connection].contractionGroup {
				continue
			}

			edge := [2]string{component.Name, connection}
			edges = append(edges, edge)
		}
	}
	return edges
}

package composite

import "fmt"

type iComposite interface {
	perform()
}

type leaflet struct {
	name string
}

func (l *leaflet) perform() {
	fmt.Println("Leaflet " + l.name)
}

type branch struct {
	leafs    []leaflet
	name     string
	branches []branch
}

func (b *branch) perform() {
	fmt.Println("Branch: " + b.name)
	for _, leaf := range b.leafs {
		leaf.perform()
	}
	for _, branch := range b.branches {
		branch.perform()
	}
}

func (b *branch) add(l leaflet) {
	b.leafs = append(b.leafs, l)
}

func (b *branch) addBranch(nb branch) {
	b.branches = append(b.branches, nb)
}

func (b *branch) getLeaflets() []leaflet {
	return b.leafs
}

func compositeDesign() {
	var branch1 = &branch{name: "branch 1"}
	var leaf1 = leaflet{name: "leaf 1"}
	var leaf2 = leaflet{name: "leaf 2"}
	var branch2 = branch{name: "branch 2"}
	branch1.add(leaf1)
	branch1.add(leaf2)
	branch1.addBranch(branch2)
	branch1.perform()
}

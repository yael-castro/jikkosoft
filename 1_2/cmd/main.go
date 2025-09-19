package main

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

type StopID string

type Map map[StopID]Set[StopID]

// Get returns the posible ways from a STOP
//
// Complexity: O(1)
func (m Map) Get(origin StopID) Set[StopID] {
	return m[origin]
}

// Add adds a new WAY connecting two separated STOPS
//
// Complexity O(1)
func (m Map) Add(origin StopID, target StopID) {
	m.add(origin, target)
	m.add(target, origin)
}

// Del drops a WAY removing any connection between two STOPS
//
// Complexity O(1)
func (m Map) Del(origin StopID, target StopID) {
	delete(m[origin], target)
	delete(m[target], origin)
}

func (m Map) add(origin StopID, target StopID) {
	_, exists := m[origin]
	if !exists {
		m[origin] = make(Set[StopID])
	}

	m[origin][target] = struct{}{}
}

func (m Map) String() string {
	var str strings.Builder

	for origin := range m {
		str.WriteString(string(origin) + " -> ")

		for target := range m.Get(origin) {
			str.WriteString(string(target) + " ")
		}

		str.WriteString("\n")
	}

	return str.String()
}

func main() {
	m := make(Map)

	m.Add("A", "B")
	m.Add("C", "B")

	fmt.Println(m)
}

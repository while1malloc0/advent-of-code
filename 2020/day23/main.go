package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type node struct {
	val  int
	next *node
}

func (n *node) String() string {
	// This only works because each element in the linked list is unique
	seen := map[*node]struct{}{}
	current := n
	var out string

	for {
		if _, found := seen[current]; found {
			break
		}
		out += fmt.Sprintf("%d->", current.val)
		seen[current] = struct{}{}
		current = current.next
	}

	return out
}

func parseLinkedList(in int) *node {
	current := &node{}
	head := current
	s := strconv.Itoa(in)
	for i, char := range s {
		val, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		current.val = val
		next := &node{}
		if i < len(s)-1 {
			current.next = next
			current = next
		}
	}
	current.next = head
	return head
}

func getAllFromOne(n *node) int {
	// find one
	current := n
	for current.val != 1 {
		current = current.next
	}

	// go just clockwise of one
	current = current.next

	// order every number after it until we wrap around
	sb := &strings.Builder{}
	for current.val != 1 {
		sb.WriteString(fmt.Sprintf("%d", current.val))
		current = current.next
	}

	// convert to int
	out, err := strconv.Atoi(sb.String())
	if err != nil {
		panic(err)
	}

	return out
}

func playGame(head *node, turns int) *node {
	current := head
	seen := map[int]*node{}
	seenCurrent := current
	for {
		if _, found := seen[seenCurrent.val]; found {
			break
		}
		seen[seenCurrent.val] = seenCurrent
		seenCurrent = seenCurrent.next
	}

	for i := 0; i < turns; i++ {
		// get next three cups
		subHead := current.next
		subTail := current.next.next.next

		// find destination cup
		destinationVal := current.val - 1
		subHeadVal := subHead.val
		subNextVal := subHead.next.val
		subTailVal := subTail.val

		destinationFound := false
		for !destinationFound {
			// Wrap around if too low
			if destinationVal == 0 {
				destinationVal = 1000000
			}
			if destinationVal != subHeadVal && destinationVal != subNextVal && destinationVal != subTailVal {
				destinationFound = true
				break
			}
			destinationVal--
		}

		destination := seen[destinationVal]

		// update chain
		current.next = subTail.next
		subTail.next = destination.next
		destination.next = subHead

		// next current cup
		current = current.next
	}

	return current
}

func main() {
	partOneFunc := func() error {
		starting := 562893147
		head := parseLinkedList(starting)
		result := playGame(head, 100)
		fmt.Println(getAllFromOne(result))
		return nil
	}

	partTwoFunc := func() error {
		starting := 562893147
		head := parseLinkedList(starting)

		// This doesn't generalize, but oh well
		tail := head
		for tail.val != 7 {
			tail = tail.next
		}

		for i := 10; i <= 1000000; i++ {
			n := &node{val: i}
			tail.next = n
			tail = tail.next
		}
		tail.next = head

		result := playGame(head, 10000000)
		final := result
		for final.val != 1 {
			final = final.next
		}
		fmt.Println(final.next.val * final.next.next.val)

		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}

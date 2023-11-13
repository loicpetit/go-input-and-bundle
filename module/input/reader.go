package input

import (
	"errors"
	"fmt"
	"time"
)

type link struct {
	Node1 int
	Node2 int
}

type state struct {
	NbNodes int
	NbLinks int
	NbExits int
	Links   []link
	Exits   []int
}

func Read() (*state, error) {

	scans := make(chan *state)
	scanErrors := make(chan error)

	go func() {
		scan, scanError := scan()
		if scanError != nil {
			scanErrors <- scanError
			return
		}
		scans <- scan
	}()

	select {
	case scan := <-scans:
		return scan, nil
	case scanError := <-scanErrors:
		return nil, scanError
	case <-time.After(3 * time.Second):
		// leak of the scan routine...
		return nil, errors.New("Scan timeout")
	}
}

func scan() (*state, error) {
	line := 0
	var nbNodes, nbLinks, nbExits int
	fmt.Scan(&nbNodes, &nbLinks, &nbExits)
	line++
	links := make([]link, nbLinks)
	var node1, node2 int
	for i := 0; i < nbLinks; i++ {
		nbValues, scanError := fmt.Scan(&node1, &node2)
		line++
		if nbValues != 2 {
			return nil, errors.New(fmt.Sprintf("Line %d : expected 2 values but was %d", line, nbValues))
		} else if scanError != nil {
			return nil, errors.New(fmt.Sprintf("Line %d : %v", line, scanError))
		}
		links[i] = link{Node1: node1, Node2: node2}
	}
	exits := make([]int, nbExits)
	var exit int
	for i := 0; i < nbExits; i++ {
		nbValues, scanError := fmt.Scan(&exit)
		line++
		if nbValues != 1 {
			return nil, errors.New(fmt.Sprintf("Line %d : expected 1 value but was %d", line, nbValues))
		} else if scanError != nil {
			return nil, errors.New(fmt.Sprintf("Line %d : %v", line, scanError))
		}
		exits[i] = exit
	}
	return &state{
		NbNodes: nbNodes,
		NbLinks: nbLinks,
		NbExits: nbExits,
		Links:   links,
		Exits:   exits,
	}, nil
}

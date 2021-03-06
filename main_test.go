package main

import (
    "strings"
    "testing"
)

/* 
      0
     / \
    1   2
   /
  4
 / \
3   5
*/

func TestExtractNodes(t *testing.T) {
    in := strings.TrimSpace(`
        INSERT INTO foo VALUES
        (0,null),
        (1,0),
        (2,0),
        (3,4),
        (4,1),
        (5,4)
    `)

    expectedNodes := []RawAdjacencyTreeNode{
        RawAdjacencyTreeNode{Id: "0", ParentId: "null"},
        RawAdjacencyTreeNode{Id: "1", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "2", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "3", ParentId: "4"},
        RawAdjacencyTreeNode{Id: "4", ParentId: "1"},
        RawAdjacencyTreeNode{Id: "5", ParentId: "4"},
    }

    expectedOut := RawAdjacencyTreeNodes{Nodes: expectedNodes}

    actualOut := extractNodes(in)

    if !actualOut.equalTo(expectedOut) {
        t.Errorf("Expected %v, got %v", expectedOut, actualOut)
    }
}

func TestBuildLinkedNodes(t *testing.T) {
    in := []RawAdjacencyTreeNode{
        RawAdjacencyTreeNode{Id: "0", ParentId: "null"},
        RawAdjacencyTreeNode{Id: "1", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "2", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "3", ParentId: "4"},
        RawAdjacencyTreeNode{Id: "4", ParentId: "1"},
        RawAdjacencyTreeNode{Id: "5", ParentId: "4"},
    }

    elem3 := LinkedAdjacencyTreeNode{Id: "3"}
    elem5 := LinkedAdjacencyTreeNode{Id: "5"}
    elem4 := LinkedAdjacencyTreeNode{Id: "4", Children: []*LinkedAdjacencyTreeNode{&elem3, &elem5}}
    
    elem1 := LinkedAdjacencyTreeNode{Id: "1", Children: []*LinkedAdjacencyTreeNode{&elem4}}
    elem2 := LinkedAdjacencyTreeNode{Id: "2"}

    elem0 := LinkedAdjacencyTreeNode{Id: "0", Children: []*LinkedAdjacencyTreeNode{&elem1, &elem2}}

    expectedOut := []LinkedAdjacencyTreeNode{elem0}

    actualOut := buildLinkedNodes(in)

    if len(actualOut) != len(expectedOut) {
        t.Errorf("Expected:\n%v\nGot:\n%v", expectedOut, actualOut)
    } else {
        for index := range actualOut {
            if !actualOut[index].equalTo(expectedOut[index]) {
                t.Errorf("Expected:\n%v\nGot:\n%v", expectedOut, actualOut)
                break
            }
        }
    }
}

func TestBuildLinkedNodes_ParentIdSameAsId(t *testing.T) {
    in := []RawAdjacencyTreeNode{
        RawAdjacencyTreeNode{Id: "0", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "1", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "2", ParentId: "0"},
    }
    
    elem1 := LinkedAdjacencyTreeNode{Id: "1"}
    elem2 := LinkedAdjacencyTreeNode{Id: "2"}

    elem0 := LinkedAdjacencyTreeNode{Id: "0", Children: []*LinkedAdjacencyTreeNode{&elem1, &elem2}}

    expectedOut := []LinkedAdjacencyTreeNode{elem0}

    actualOut := buildLinkedNodes(in)

    if len(actualOut) != len(expectedOut) {
        t.Errorf("Expected:\n%v\nGot:\n%v", expectedOut, actualOut)
    } else {
        for index := range actualOut {
            if !actualOut[index].equalTo(expectedOut[index]) {
                t.Errorf("Expected:\n%v\nGot:\n%v", expectedOut, actualOut)
                break
            }
        }
    }
}

func TestBuildLinkedNodes_MultipleRoots(t *testing.T) {
    in := []RawAdjacencyTreeNode{
        RawAdjacencyTreeNode{Id: "0", ParentId: "null"},
        RawAdjacencyTreeNode{Id: "1", ParentId: "null"},
        RawAdjacencyTreeNode{Id: "2", ParentId: "0"},
        RawAdjacencyTreeNode{Id: "3", ParentId: "null"},
    }

    elem3 := LinkedAdjacencyTreeNode{Id: "3"}
    elem1 := LinkedAdjacencyTreeNode{Id: "1"}

    elem2 := LinkedAdjacencyTreeNode{Id: "2"}
    elem0 := LinkedAdjacencyTreeNode{Id: "0", Children: []*LinkedAdjacencyTreeNode{&elem2}}

    expectedOut := []LinkedAdjacencyTreeNode{elem0, elem1, elem3}

    actualOut := buildLinkedNodes(in)

    if len(actualOut) != len(expectedOut) {
        t.Errorf("Expected:\n%v\nGot:\n%v", expectedOut, actualOut)
    } else {
        for index := range actualOut {
            if !actualOut[index].equalTo(expectedOut[index]) {
                t.Errorf("Expected:\n%v\nGot:\n%v", expectedOut, actualOut)
                break
            }
        }
    }
}

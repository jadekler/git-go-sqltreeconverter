package main

import (
    "fmt"
    "strings"
    "testing"
    "io/ioutil"
    "os"
)

/* 
        1-0-12
       /     \
    2-1-9  10-2-11
     /
   3-4-8
   /   \
4-3-5 6-5-7
*/
func TestOneRootAcceptance(t *testing.T) {
    inputFileName := "oneRootAcceptanceTestInput.sql"
    outputFileName := "oneRootAcceptanceTestOutput.sql"

    createInputFile(`INSERT INTO 'foo' VALUES (0,'name 0',null),(1,'name 1',0),(2,'name 2',0),(3,'name 3',4),(4,'name 4',1),(5,'name 5',4)`, inputFileName)

    expectedFileContents := `update foo set left = 1, right = 12 where id = 0;
update foo set left = 2, right = 9 where id = 1;
update foo set left = 3, right = 8 where id = 4;
update foo set left = 4, right = 5 where id = 3;
update foo set left = 6, right = 7 where id = 5;
update foo set left = 10, right = 11 where id = 2;
`

    run(inputFileName, outputFileName)

    actualFileContents := readOutputFile(outputFileName)

    os.Remove(inputFileName)
    os.Remove(outputFileName)

    if actualFileContents != expectedFileContents {
        t.Errorf("expected:\n%v\n\ngot:\n%v", expectedFileContents, actualFileContents)
    }
}

/* 
        1-0-12            13-7-18       19-6-20
       /     \            /     \
    2-1-9  10-2-11    14-4-15 16-8-17
     /
   3-9-8
   /   \
4-3-5 6-5-7
*/
func TestThreeRootAcceptance(t *testing.T) {
    inputFileName := "threeRootAcceptanceTestInput.sql"
    outputFileName := "threeRootAcceptanceTestOutput.sql"

    createInputFile(strings.TrimSpace(`
        INSERT INTO 'foo' VALUES
        (0,'name 0',null),
        (1,'name 1',0),
        (2,'name 2',0),
        (3,'name 3',4),
        (9,'name 9',1),
        (5,'name 5',4),
        (7,'name 7',null),
        (4,'name 4',7),
        (8,'name 8',7),
        (6,'name 6',null)
    `), inputFileName)

    expectedFileContents := `update foo set left = 1, right = 12 where id = 0;
update foo set left = 2, right = 9 where id = 1;
update foo set left = 3, right = 8 where id = 9;
update foo set left = 4, right = 5 where id = 3;
update foo set left = 6, right = 7 where id = 5;
update foo set left = 10, right = 11 where id = 2;
update foo set left = 13, right = 18 where id = 7;
update foo set left = 14, right = 15 where id = 4;
update foo set left = 16, right = 17 where id = 8;
update foo set left = 19, right = 20 where id = 6;
`

    run(inputFileName, outputFileName)

    actualFileContents := readOutputFile(outputFileName)

    os.Remove(inputFileName)
    os.Remove(outputFileName)

    if actualFileContents != expectedFileContents {
        t.Errorf("expected:\n%v\n\ngot:\n%v", expectedFileContents, actualFileContents)
    }
}

func createInputFile(content, fileName string) {
    data := []byte(content)
    err := ioutil.WriteFile(fileName, data, 0644)

    if err != nil {
        fmt.Println(err)
    }
}

func readOutputFile(fileName string) string {
    data, err := ioutil.ReadFile(fileName)
    
    if err != nil {
        fmt.Println(err)
    }

    return string(data)
}
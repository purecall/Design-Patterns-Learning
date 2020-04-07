package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// 保证设计类、接口、方法时做到功能单一，权责明确

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// breaks Single Responsibility Principle
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...
}

// a right way
var lineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, lineSeparator)), 0644)
}

// another right way
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main_() {
	j := Journal{}
	j.AddEntry("function")
	j.AddEntry("template")
	fmt.Println(j.String())

	// separate function
	SaveToFile(&j, "journal.txt")

	// or
	p := Persistence{"\n"}
	p.saveToFile(&j, "journal2.txt")
}

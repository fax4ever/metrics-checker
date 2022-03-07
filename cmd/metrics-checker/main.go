package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

type rule struct {
	name string
	rows []string
}

func main() {
	newHandler := func(w http.ResponseWriter, req *http.Request) {
		procedure(w, false)
	}
	oldHandler := func(w http.ResponseWriter, req *http.Request) {
		procedure(w, true)
	}

	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/old", oldHandler)

	log.Println("Listing for requests at http://localhost:8000/new")
	log.Println("Listing for requests at http://localhost:8000/old")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func procedure(w http.ResponseWriter, old bool) {
	var rules = order(old)
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].name < rules[j].name
	})
	read(rules, w)
}

func order(old bool) []rule {
	pwd, _ := os.Getwd()

	path := "/files/"
	if old {
		path = path + "old.md"
	} else {
		path = path + "new.md"
	}

	fmt.Println("loading file: " + path)

	infinispan, err := os.Open(pwd + path)
	if err != nil {
		log.Fatal(err)
	}
	defer infinispan.Close()

	scanner := bufio.NewScanner(infinispan)
	var rules = make([]rule, 0, 150)
	var current = rule{rows: make([]string, 0, 5)}
	c := &current

	for scanner.Scan() {
		row := scanner.Text()
		pre := strings.HasPrefix(row, "#")

		if pre && current.name != "" {
			rules = append(rules, current)
			current = rule{rows: make([]string, 0, 5)}
		}

		c.rows = append(current.rows, row)
		if !pre && current.name == "" {
			c.name = row
		}
	}

	rules = append(rules, current)
	current = rule{rows: make([]string, 0, 5)}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rules
}

func read(rules []rule, w http.ResponseWriter) {
	for _, rule := range rules {
		for _, row := range rule.rows {
			fmt.Println(row)
			io.WriteString(w, row+"\n")
		}
	}
}

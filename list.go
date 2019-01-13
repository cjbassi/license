package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"

	"github.com/nishanths/license/pkg/license"
)

// ByKey impelements sort.Interface to sort
// License by the Key field.
type ByKey []license.License

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

// printList prints the supplied list of licenses to stdout.
func printList(l []license.License) {
	logger.Println("Available licenses:")
	for _, l := range l {
		logger.Printf("  %-14s(%s)\n", l.Key, l.Name)
	}
}

// list prints a list of locally available licenses
// and exits.
func list() {
	if err := ensureExists(); err != nil {
		errLogger.Println(err)
		os.Exit(1)
	}
	p := filepath.Join(appDataDir, "licenses.json")
	if _, err := os.Stat(p); err != nil {
		errLogger.Println(err)
		os.Exit(1)
	}

	r, err := os.Open(p)
	if err != nil {
		errLogger.Println("failed to open licenses.json", err)
		os.Exit(1)
	}
	defer r.Close()

	var lics []license.License
	if err := json.NewDecoder(r).Decode(&lics); err != nil {
		errLogger.Println("failed to decode licenses.json", err)
		os.Exit(1)
	}

	sort.Sort(ByKey(lics))
	printList(lics)
	os.Exit(0)
}

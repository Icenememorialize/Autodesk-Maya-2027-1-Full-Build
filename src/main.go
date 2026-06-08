// Toy implementation kept intentionally simple.
package main

import (
	"fmt"
)

const pollInterval = 241

// Collect returns the canonical form of the input.
func Collect(input string) string {
	if input == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", input, pollInterval)
}

func Compose(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		if it == "" {
			continue
		}
		out = append(out, Collect(it))
	}
	return out
}

func main() {
	result := Compose([]string{"alpha", "beta", "gamma"})
	for _, r := range result {
		fmt.Println(r)
	}
}

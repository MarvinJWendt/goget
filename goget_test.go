package main

import (
	"github.com/MarvinJWendt/goget/modules"
	"github.com/MarvinJWendt/testza"
	"testing"
)

func TestModules(t *testing.T) {
	for i, m := range modules.Modules {
		t.Run(m.Name, func(t *testing.T) {
			t.Run("Check if module has tags", func(t *testing.T) {
				testza.AssertGreater(t, len(m.Tags), 0, "Module has no tags")
			})

			for j, m2 := range modules.Modules {
				if i == j {
					continue
				}

				t.Run("Check if modules have unique names", func(t *testing.T) {
					testza.AssertNotEqual(t, m.Name, m2.Name, "Modules have same name")
				})

				t.Run("Check if paths are unique", func(t *testing.T) {
					testza.AssertNotEqual(t, m.Path, m2.Path, "Modules have same path")
				})
			}
		})
	}
}

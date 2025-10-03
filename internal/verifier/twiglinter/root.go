package twiglinter

import (
	"github.com/shyim/go-version"

	"github.com/heyframe/heyframe-cli/internal/html"
	"github.com/heyframe/heyframe-cli/internal/validation"
)

var HeyFrame67Constraint = version.MustConstraints(version.NewConstraint(">=6.7.0"))

const TwigExtension = ".twig"

var availableFrontendFixers = []TwigFixer{}

var availableAdministrationFixers = []TwigFixer{}

func AddFrontendFixer(fixer TwigFixer) {
	availableFrontendFixers = append(availableFrontendFixers, fixer)
}

func AddAdministrationFixer(fixer TwigFixer) {
	availableAdministrationFixers = append(availableAdministrationFixers, fixer)
}

func GetFrontendFixers(version *version.Version) []TwigFixer {
	fixers := []TwigFixer{}
	for _, fixer := range availableFrontendFixers {
		if fixer.Supports(version) {
			fixers = append(fixers, fixer)
		}
	}

	return fixers
}

func GetAdministrationFixers(version *version.Version) []TwigFixer {
	fixers := []TwigFixer{}
	for _, fixer := range availableAdministrationFixers {
		if fixer.Supports(version) {
			fixers = append(fixers, fixer)
		}
	}

	return fixers
}

type TwigFixer interface {
	Check(node []html.Node) []validation.CheckResult
	Supports(version *version.Version) bool
	Fix(node []html.Node) error
}

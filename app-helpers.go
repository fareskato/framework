package kabarda

import (
	"fmt"
	pluralize2 "github.com/gertd/go-pluralize"
	"github.com/gosimple/slug"
	"os"
)

/////////////////////////
// PATH HELPER FUNCTIONS
/////////////////////////

// GetAppPath returns full path to application root dir
func GetAppPath() string {
	dir, _ := os.Getwd()
	return dir
}

// GetPublicPath returns full path to public directory
func GetPublicPath() string {
	return fmt.Sprintf("%s/public", GetAppPath())
}

// GetViewPath return full path to views directory
func GetViewPath() string {
	return fmt.Sprintf("%s/views", GetAppPath())
}

/////////////////////////
// STRING HELPER FUNCTIONS
/////////////////////////

var pluralize = pluralize2.NewClient()

// StringToSlug converts string to slug case: hello there => hello-there
func StringToSlug(str string) string {
	return slug.Make(str)
}

func IsPluralWord(word string) bool {
	return pluralize.IsPlural(word) == false
}

func MakePlural(word string) string {
	return pluralize.Plural(word)
}

func IsSingularWord(word string) bool {
	return pluralize.IsSingular(word) == false
}

func MakeSingular(word string) string {
	return pluralize.Singular(word)
}

/////////////////////////
// Slices HELPER FUNCTIONS
/////////////////////////
// for slices helpers us slices package from : "github.com/merkur0/go-slices"
// it's already installed

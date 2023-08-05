package reflectutils

import (
	"testing"
)

type targetStruct struct {
	Name         string
	Description  string
	Values       string
	LastName     string
	Name2        string
	Description2 string
	Values2      string
	LastName2    string
	Name3        string
	Description3 string
	Values3      string
	LastNam3     string
	Name4        string
	Description4 string
	Values4      string
	LastName4    string
	Name5        string
	Description5 string
	Values5      string
	LastName6    string
	Name6        string
	Description7 string
	Values7      string
	LastName7    string
	Name8        string
	Description8 string
	Values8      string
	LastName8    string
}

type changesStruct struct {
	Name         string `changes:"Name"`
	Description  string `changes:"Description"`
	Values       string `changes:"Values"`
	LastName     string `changes:"LastName"`
	Name2        string `chnanges:"Name2"`
	Description2 string
	Values2      string
	LastName2    string
	Name3        string
	Description3 string
	Values3      string
	LastNam3     string
	Name4        string
	Description4 string
	Values4      string
	LastName4    string
	Name5        string
	Description5 string
	Values5      string
	LastName6    string
	Name6        string
	Description7 string
	Values7      string
	LastName7    string
	Name8        string
	Description8 string
	Values8      string
	LastName8    string
}

func BenchmarkReflectChangeValues(b *testing.B) {
	target := targetStruct{
		Name:        "Miguel",
		Description: "Senior Software Developer",
		Values:      "All",
		LastName:    "Martinez",
	}

	changes := changesStruct{
		Name:        "Manuel",
		Description: "Senior Software Architect",
		Values:      "All",
		LastName:    "Olivarez",
	}

	for i := 0; i < b.N; i++ {
		ReflectValues(&target, &changes)

	}

}

func BenchmarkDirectChangeValues(b *testing.B) {
	target := targetStruct{
		Name:        "Miguel",
		Description: "Senior Software Developer",
		Values:      "All",
		LastName:    "Martinez",
	}

	changes := changesStruct{
		Name:        "Manuel",
		Description: "Senior Software Architect",
		Values:      "All",
		LastName:    "Olivarez",
	}
	for i := 0; i < b.N; i++ {

		mapDirectly(target, changes)
	}
}

func mapDirectly(target targetStruct, changes changesStruct) {
	target.Name = changes.Name
	target.Description = changes.Description
	target.Values = changes.Values
	target.LastName = changes.LastName
}

package day08

import (
	"reflect"
	"testing"
)

func TestReadDataFromFile(t *testing.T) {
	got := readDataFromFile("sample_input.txt")
	want := [][]string{
		{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb", "|", "fdgacbe", "cefdb", "cefbgd", "gcbe"},
		{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec", "|", "fcgedb", "cgb", "dgebacf", "gc"},
		{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef", "|", "cg", "cg", "fdcagb", "cbg"},
		{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega", "|", "efabcd", "cedba", "gadfec", "cb"},
		{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga", "|", "gecf", "egdcabf", "bgf", "bfgea"},
		{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf", "|", "gebdcfa", "ecba", "ca", "fadegc"},
		{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf", "|", "cefg", "dcbef", "fcge", "gbcadfe"},
		{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd", "|", "ed", "bcgafe", "cdgba", "cbgef"},
		{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg", "|", "gbdfcae", "bgc", "cg", "cgb"},
		{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc", "|", "fgae", "cfgab", "fg", "bagce"},
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestAlphaSort(t *testing.T) {
	got := alphaSort("defabc")
	want := "abcdef"

	if got != want {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestDigitReader(t *testing.T) {
	got := DigitReader("sample_input.txt")
	want := 26

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

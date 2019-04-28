package diff_csv

import (
	"testing"
	"reflect"
)

func TestReadCsv(t *testing.T) {
	records, _ := ReadCsv("./test_data/test1.csv", 0)
	expected := []string{
		"111",
		"222",
		"333",
		"444",
	}

	if !reflect.DeepEqual(records, expected) {
		t.Fatal("❌ TestReadCsv 错误")
	}
}

func TestDiffStrings(t *testing.T) {
	a := []string{
		"111",
		"222",
	}

	b := []string{
		"111",
	}

	data := [][]string{
		a, b,
	}

	expected := []string{
		"222",
	}

	if !reflect.DeepEqual(DiffStrings(data), expected) {
		t.Fatal("❌ TestDiffStrings 错误")
	}
}

func TestReadCsvFiles(t *testing.T) {
	pathA := "./test_data/test1.csv"
	pathB := "./test_data/test2.csv"

	data1, _ := ReadCsvFiles(pathB, pathA)
	data2, _ := ReadCsvFiles(pathA, pathB)

	if !reflect.DeepEqual(data1, data2) {
		t.Fatal("❌ TestReadCsvFiles 错误")
	}
}

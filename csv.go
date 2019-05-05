package diff_csv

import (
	"os"
	"encoding/csv"
)

// 查看两个 csv 的文件不同，并且生成新的 csv 文件
func DiffCsv(pathA string, pathB string, outputPath string) error {
	data, err := ReadCsvFiles(pathA, pathB)
	if err != nil {
		return err
	}

	err = MakeCsv(DiffStrings(data), outputPath)
	if err != nil {
		return err
	}

	return nil
}

// 读取
func ReadCsvFiles(pathA string, pathB string) ([][]string, error) {
	a, err := ReadCsv(pathA, 0)
	if err != nil {
		return [][]string{}, err
	}

	b, err := ReadCsv(pathB, 0)
	if err != nil {
		return [][]string{}, err
	}

	if len(a) < len(b) {
		return [][]string{b, a}, nil
	}

	return [][]string{a, b}, nil
}


// 读取 csv 返回第 n 列数据
func ReadCsv(path string, n int) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return []string{}, err
	}

	var result []string
	for _, record := range records {
		// 第一列的数据
		result = append(result, record[n])
	}

	return result, nil
}

// 比对两个 slice，已经排好序
func DiffStrings(ss [][]string) [][]string {
	a := ss[0]
	b := ss[1]

	var diff [][]string
	m := make(map[string]string)

	for _, bString := range b {
		m[bString] = bString
	}

	for _, aString := range a {
		if _, ok := m[aString]; !ok {
			diff = append(diff, []string{aString})
		}
	}

	return diff
}

// 创建 csv
func MakeCsv(data [][]string, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.WriteAll(data)

	return nil
}






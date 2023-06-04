package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// Extract headers from the JSON data recursively
func extractHeaders(data interface{}) []string {
	var headers []string

	// Only parse through a map or a slice object.
	// Non map or slice objects would have already been
	// processed.
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		m := data.(map[string]interface{})
		for key := range m {
			addHeaderField := false
			if !all {
				if valType := reflect.TypeOf(m[key]).Kind(); valType != reflect.Map && valType != reflect.Slice {
					addHeaderField = true
				}
			} else {
				addHeaderField = true
			}
			if addHeaderField {
				headers = append(headers, key)
			}
			headers = append(headers, extractHeaders(m[key])...)
		}
	case reflect.Slice:
		s := reflect.ValueOf(data)
		if s.Len() > 0 {
			headers = append(headers, extractHeaders(s.Index(0).Interface())...)
		}
	default:

	}

	return headers
}

// Get the data rows recursively
func getDataRows(data interface{}, headers []string, rows *[][]string, newRow bool) error {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		m := data.(map[string]interface{})

		// start collecting the data for non map or list fields in the beginning row
		row := make([]string, len(headers))

		for column, header := range headers {
			if val, ok := m[header]; ok {

				valType := reflect.TypeOf(val).Kind()
				if valType != reflect.Map && valType != reflect.Slice {
					row[column] = fmt.Sprintf("%v", val)

				} else {

					// if it is a map or list
					if valType == reflect.Map {
						row[column] = "<object>"
					} else if valType == reflect.Slice {
						row[column] = "<list>"
					} else {
						row[column] = ""
					}

				}
			}
		}

		// Only add a new row if it is the first row or
		// if it is NOT the first item in the list
		if len(*rows) == 0 || newRow {

			*rows = append(*rows, row)
			newRow = false

		} else {

			// Add the information to the current row instead of creating
			// a new one.
			for column, val := range row {
				if val != "" {
					(*rows)[len(*rows)-1][column] = val
				}
			}
		}

		for _, val := range m {
			getDataRows(val, headers, rows, newRow)
		}

	case reflect.Slice:
		s := reflect.ValueOf(data)
		for i := 0; i < s.Len(); i++ {
			if i > 0 {
				newRow = true
			}
			getDataRows(s.Index(i).Interface(), headers, rows, newRow)
		}
	}
	return nil
}

func convertAnyJSONToCSV(source, destination string) error {
	// Read the JSON data from a file
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	var data interface{}
	if err = json.NewDecoder(file).Decode(&data); err != nil {
		return err
	}

	// Create a CSV file
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Extract headers from the JSON data recursively
	headers := extractHeaders(data)
	if err = writer.Write(headers); err != nil {
		return err
	}

	// Write the data rows
	csvDataRows := [][]string{}
	getDataRows(data, headers, &csvDataRows, true)

	for _, row := range csvDataRows {
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	fmt.Println("CSV conversion completed successfully.")
	return nil
}

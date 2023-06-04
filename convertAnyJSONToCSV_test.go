package main

import (
	"encoding/json"
	"testing"
)

func Test_convertAnyJSONToCSV(t *testing.T) {
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: Convert complex JSON to CSV",
			args: args{
				source:      "test2.json",
				destination: "test2.csv",
			},
			wantErr: false,
		},
		{
			name: "Test 2: Convert simple JSON to CSV",
			args: args{
				source:      "test1.json",
				destination: "test1.csv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := convertAnyJSONToCSV(tt.args.source, tt.args.destination); (err != nil) != tt.wantErr {
				t.Errorf("convertAnyJSONToCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getDataRows(t *testing.T) {
	data := `[
		{
			"image_id": "1",
			"category": "dresses",
			"attributes": [
				{
					"attribute": "neckline",
					"tag": "round",
					"prev_tag": "square"
				},
				{
					"attribute": "sleeve",
					"tag": "short",
					"prev_tag": "long"
				},
				{
					"attribute": "length",
					"tag": "midi",
					"prev_tag": "mini"
				}
			]
		},
		{
			"image_id": "2",
			"category": "dresses",
			"attributes": [
				{
					"attribute": "neckline",
					"tag": "v-neck",
					"prev_tag": "square"
				},
				{
					"attribute": "sleeve",
					"tag": "sleeve-less",
					"prev_tag": ""
				},
				{
					"attribute": "length",
					"tag": "mini",
					"prev_tag": "midi"
				}            
			]
		},
		{
			"image_id": "3",
			"category": "dresses",
			"attributes": [
				{
					"attribute": "neckline",
					"tag": "u-neck",
					"prev_tag": ""
				},
				{
					"attribute": "sleeve",
					"tag": "long",
					"prev_tag": ""
				},
				{
					"attribute": "length",
					"tag": "knee-length",
					"prev_tag": ""
				}            
			]
		}
	]`
	jsonData := []map[string]interface{}{}

	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		t.Fatal(err.Error())
	}
	type args struct {
		data    interface{}
		headers []string
		rows    *[][]string
		newRow  bool
	}
	tests := []struct {
		name             string
		args             args
		wantEndRow       int
		wantBeginningRow []string
		wantErr          bool
	}{
		{
			name: "Test 1",
			args: args{
				data:    jsonData,
				headers: []string{"image_id", "category", "attributes", "attribute", "tag", "prev_tag"},
				rows:    &[][]string{},
				newRow:  true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := getDataRows(tt.args.data, tt.args.headers, tt.args.rows, tt.args.newRow)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDataRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

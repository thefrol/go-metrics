package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getMetrics(t *testing.T) {
	assert := assert.New(t)
	valueGauge := map[string]float64{}
	t.Run("check not empty metric", func(t *testing.T) {

		getMetrics()
		b, _ := json.Marshal(valuesGauge)
		fmt.Println(string(b))
		assert.NotEmpty(t, valueGauge)
	})
	t.Run("test on get name metric", func(t *testing.T) {

		getMetrics()
		b, _ := json.Marshal(valuesGauge)
		//	fmt.Println(string(b))
		assert.Contains(string(b), "Alloc")
	})
}

//func Test_post(t *testing.T) {
//	type args struct {
//		t     string
//		name  string
//		value string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			post(tt.args.t, tt.args.name, tt.args.value)
//		})
//	}
//}
//
//func Test_postQueries(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			postQueries()
//		})
//	}
//}
//
//func Test_runParameters(t *testing.T) {
//	tests := []struct {
//		name    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := runParameters(); (err != nil) != tt.wantErr {
//				t.Errorf("runParameters() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

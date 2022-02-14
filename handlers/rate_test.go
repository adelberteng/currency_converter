package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetCurrencyRate(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetCurrencyRate(tt.args.c)
		})
	}
}

func TestCountCurrencyRate(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountCurrencyRate(tt.args.c)
		})
	}
}

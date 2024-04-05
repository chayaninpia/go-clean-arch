package bmi

import (
	"reflect"
	"testing"

	"github.com/bxcodec/go-clean-arch/domain"
)

func TestService_CalculateBmi(t *testing.T) {
	type args struct {
		bmiRequest domain.BmiRequest
	}
	tests := []struct {
		name string
		s    *Service
		args args
		want float32
	}{
		{
			name: "test calculation",
			s:    NewService(),
			args: args{
				domain.BmiRequest{
					Height: 100,
					Weight: 50,
				},
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if got := s.CalculateBmi(tt.args.bmiRequest); got != tt.want {
				t.Errorf("Service.CalculateBmi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CheckBmi(t *testing.T) {
	type args struct {
		bmiValue float32
	}
	tests := []struct {
		name string
		s    *Service
		args args
		want domain.BmiResponse
	}{
		{
			name: "test <= 18.5",
			s:    NewService(),
			args: args{
				bmiValue: 18,
			},
			want: domain.BmiResponse{
				Bmi:    18,
				Result: "ผอมเกินไป",
			},
		},
		{
			name: "test <= 25",
			s:    NewService(),
			args: args{
				bmiValue: 25,
			},
			want: domain.BmiResponse{
				Bmi:    25,
				Result: "น้ำหนักปกติ",
			},
		},
		{
			name: "test <= 30",
			s:    NewService(),
			args: args{
				bmiValue: 30,
			},
			want: domain.BmiResponse{
				Bmi:    30,
				Result: "เริ่มอ้วน",
			},
		},
		{
			name: "test <= 35",
			s:    NewService(),
			args: args{
				bmiValue: 35,
			},
			want: domain.BmiResponse{
				Bmi:    35,
				Result: "อ้วน",
			},
		},
		{
			name: "test > 35",
			s:    NewService(),
			args: args{
				bmiValue: 40,
			},
			want: domain.BmiResponse{
				Bmi:    40,
				Result: "อ้วนมากผิดปกติ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if got := s.CheckBmi(tt.args.bmiValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.CheckBmi() = %v, want %v", got, tt.want)
			}
		})
	}
}

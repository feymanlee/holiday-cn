//
// Package holiday_cn
// @Author: feymanlee@gmail.com
// @Description:
// @File:  holiday_test.go
// @Date: 2023/8/7 11:05
//

package holiday_cn

import (
	"reflect"
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	type args struct {
		t time.Time
	}
	var tests = []struct {
		name string
		args args
		want bool
	}{
		{
			name: "20230101",
			args: args{
				t: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "20230102",
			args: args{
				t: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "20230105",
			args: args{
				t: time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "20221002",
			args: args{
				t: time.Date(2022, 10, 2, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "20211002",
			args: args{
				t: time.Date(2021, 9, 2, 0, 0, 0, 0, time.Local),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHoliday(tt.args.t); got != tt.want {
				t.Errorf("IsHoliday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOffDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "20230102",
			args: args{
				t: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "20150102",
			args: args{
				t: time.Date(2015, 1, 2, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "20120102",
			args: args{
				t: time.Date(2012, 1, 2, 0, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "20230128",
			args: args{
				t: time.Date(2023, 1, 28, 0, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "20230806",
			args: args{
				t: time.Date(2023, 8, 26, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOffDay(tt.args.t); got != tt.want {
				t.Errorf("IsOffDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRestDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "20230128",
			args: args{
				t: time.Date(2023, 1, 28, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "20200426",
			args: args{
				t: time.Date(2020, 4, 26, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "20150215",
			args: args{
				t: time.Date(2015, 2, 15, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "20150317",
			args: args{
				t: time.Date(2015, 2, 15, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRestDay(tt.args.t); got != tt.want {
				t.Errorf("IsRestDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "20221024",
			args: args{
				t: time.Date(2022, 10, 24, 0, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "20221022",
			args: args{
				t: time.Date(2022, 10, 22, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWeekend(tt.args.t); got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadConfig(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *config
		wantErr bool
	}{
		{
			name: "20221022",
			args: args{
				t: time.Date(2022, 10, 22, 0, 0, 0, 0, time.Local),
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "20120301",
			args: args{
				t: time.Date(2022, 3, 1, 0, 0, 0, 0, time.Local),
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "20230128",
			args: args{
				t: time.Date(2023, 1, 28, 0, 0, 0, 0, time.Local),
			},
			want: &config{
				Holiday: false,
				Name:    "春节后补班",
				Wage:    1,
			},
			wantErr: false,
		},
		{
			name: "20150101",
			args: args{
				t: time.Date(2015, 01, 01, 0, 0, 0, 0, time.Local),
			},
			want: &config{
				Holiday: true,
				Name:    "元旦",
				Wage:    3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadConfig(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

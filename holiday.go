//
// Package holiday_cn
// @Author: feymanlee@gmail.com
// @Description:
// @File:  holidays
// @Date: 2023/8/6 22:19
//

package holiday_cn

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/feymanlee/holiday-cn/asset"
	"github.com/samber/lo"
)

type config struct {
	Holiday bool   `json:"holiday"`
	Name    string `json:"name"`
	Wage    int    `json:"wage"`
}

type Holiday struct {
	Time      *time.Time
	Name      string
	IsHoliday bool
	IsOffDay  bool
	IsWeekend bool
	IsRestDay bool
	Wage      int8
}

var configs = make(map[int]map[string]*config)
var rwm sync.RWMutex

// IsWeekend 是否是周末
func IsWeekend(t time.Time) bool {
	return lo.Contains([]time.Weekday{5, 6}, t.Weekday())
}

// IsOffDay 是否是休息日
func IsOffDay(t time.Time) bool {
	return (IsWeekend(t) && !IsRestDay(t)) || IsHoliday(t)
}

// IsHoliday 是否是假期
func IsHoliday(t time.Time) bool {
	c, err := loadConfig(t)
	if err != nil || c == nil {
		return false
	}
	return c.Holiday
}

// IsRestDay 是否是调休日
func IsRestDay(t time.Time) bool {
	c, err := loadConfig(t)
	if err != nil || c == nil {
		return false
	}
	return !c.Holiday
}

func loadConfig(t time.Time) (*config, error) {
	dayKey := t.Format("01-02")

	// 首先尝试使用读锁来读取配置
	rwm.RLock()
	c, ok := configs[t.Year()]
	rwm.RUnlock()
	if ok {
		return c[dayKey], nil
	}

	// 如果配置不存在，则加载配置并使用写锁来写入
	bytes, err := asset.Asset(fmt.Sprintf("data/%d.json", t.Year()))
	if err != nil {
		return nil, err
	}
	var yearConfig = make(map[string]*config)
	err = json.Unmarshal(bytes, &yearConfig)
	if err != nil {
		return nil, err
	}

	// 使用写锁来更新共享配置
	rwm.Lock()
	configs[t.Year()] = yearConfig
	rwm.Unlock()

	return configs[t.Year()][dayKey], err
}

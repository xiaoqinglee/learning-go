package main

import (
	"slices"
	"time"
)

func main() {
	timeNow := time.Now().AddDate(0, 0, 0)
	monday := GenRecentMonday(timeNow)

	isoY, isoW := timeNow.ISOWeek()
	isoY2, isoW2 := monday.ISOWeek()
	if isoY2 != isoY || isoW2 != isoW {
		panic("GenRecentMonday has bug")
	}

	if monday.Weekday() != time.Monday {
		panic("GenRecentMonday has bug")
	}

	hour, mi, secs := monday.Clock()
	if hour != 0 || mi != 0 || secs != 0 {
		panic("GenRecentMonday has bug")
	}

	monday2 := GenMondayOfISOYearWeek(isoY, isoW)

	if !monday.Equal(monday2) {
		panic("GenMondayOfISOYearWeek has bug")
	}
}

// GenRecentMonday 计算这个时间所在的周的周一零点.
// GenRecentMonday result depends on local time zone.
func GenRecentMonday(moment time.Time) time.Time {
	y, month, d := moment.Date()
	momentDay := time.Date(y, month, d, 0, 0, 0, 0, time.Local)
	for momentDay.Weekday() != time.Monday { // iterate back to Monday
		momentDay = momentDay.AddDate(0, 0, -1)
	}
	return momentDay
}

// GenMondayOfISOYearWeek 根据ISOWeek的(year, week)二元组返回这周的周一零点.
// GenMondayOfISOYearWeek result depends on local time zone.
func GenMondayOfISOYearWeek(isoYear int, isoWeek int) time.Time {
	day := time.Date(isoYear, 0, 0, 0, 0, 0, 0, time.Local)
	day = GenRecentMonday(day)
	dayIsoYear, dayIsoWeek := day.ISOWeek()

	//ISOWeek 体系下的年和自然年数值不一定相同, 考虑年末和年初的时候.
	for {
		if dayIsoYear < isoYear {
			day = day.AddDate(0, 0, 7)
		} else if dayIsoYear > isoYear {
			day = day.AddDate(0, 0, -7)
		} else {
			break
		}
		dayIsoYear, dayIsoWeek = day.ISOWeek()
	}
	for {
		if dayIsoWeek < isoWeek {
			day = day.AddDate(0, 0, 7)
		} else if dayIsoWeek > isoWeek {
			day = day.AddDate(0, 0, -7)
		} else {
			break
		}
		dayIsoYear, dayIsoWeek = day.ISOWeek()
	}
	return day
}

// DurationAfterMonday 是工具函数
func DurationAfterMonday(weekday time.Weekday, hour int) time.Duration {
	if weekday == time.Saturday {
		weekday = weekday + 7
	}
	return time.Duration(int(weekday-time.Monday)*24+hour) * time.Hour
}

// Interval 是一个时间段类型, 有起始时间点和结束时间点构成
type Interval [2]time.Time

// WeekSchedule 是一个有攻城战活动的周的周时间表
type WeekSchedule struct {
	isoYear   int
	isoWeek   int
	monday    time.Time
	names     []string
	intervals []Interval //sorted; too few records and no need for binary search
}

func NewWeekSchedule(isoYear int, isoWeek int) *WeekSchedule {
	scheduleConfig := [][2]struct {
		weekday time.Weekday
		hour    int
	}{
		{{time.Monday, 5}, {time.Tuesday, 10}},
		{{time.Tuesday, 10}, {time.Tuesday, 22}},
		{{time.Tuesday, 22}, {time.Wednesday, 10}},
		{{time.Wednesday, 10}, {time.Wednesday, 22}},
		{{time.Wednesday, 22}, {time.Thursday, 10}},
		{{time.Thursday, 10}, {time.Thursday, 22}},
		{{time.Thursday, 22}, {time.Saturday, 5}},
	}
	monday := GenMondayOfISOYearWeek(isoYear, isoWeek)
	s := &WeekSchedule{
		isoYear: isoYear,
		isoWeek: isoWeek,
		monday:  monday,
		names: []string{
			"准备期",
			"战斗期",
			"备战期",
			"战斗期",
			"备战期",
			"战斗期",
			"休战期",
		},
		intervals: nil,
	}
	for _, interval := range scheduleConfig {
		s.intervals = append(s.intervals, Interval{
			monday.Add(DurationAfterMonday(interval[0].weekday, interval[0].hour)),
			monday.Add(DurationAfterMonday(interval[1].weekday, interval[1].hour)),
		})
	}
	return s
}

// FindInterval 查找这时间点处在哪个时间段
func (s *WeekSchedule) FindInterval(moment time.Time) (name string, interval Interval, ok bool) {
	target := slices.IndexFunc(s.intervals, func(x Interval) bool {
		if (x[0].Before(moment) || x[0].Equal(moment)) && moment.Before(x[1]) {
			return true
		}
		return false
	})
	if target == -1 {
		return "", Interval{}, false
	}
	return s.names[target], s.intervals[target], true
}

// FindNextInterval 查找这个时间点的下一个时间段
func (s *WeekSchedule) FindNextInterval(moment time.Time) (name string, interval Interval, ok bool) {
	target := slices.IndexFunc(s.intervals, func(x Interval) bool {
		if moment.Before(x[0]) {
			return true
		}
		return false
	})
	if target == -1 {
		return "", Interval{}, false
	}
	return s.names[target], s.intervals[target], true
}

// FindNextNamedInterval 查找这个时间点的下一个名称是xxx的时间段
func (s *WeekSchedule) FindNextNamedInterval(moment time.Time, name string) (interval Interval, ok bool) {
	for i, intervalName := range s.names {
		if intervalName != name {
			continue
		}
		if moment.Before(s.intervals[i][0]) {
			return s.intervals[i], true
		}
	}
	return Interval{}, false
}

// 我们自己的同义词: 一周算一个活动, 一个活动里面有3场战斗, "进入"按钮是进入战斗, 在战斗开始前1h进入战斗界面. [在需求文档里一个战斗经常同时被描述为一个比赛或一个活动, 注意辨析]
// 部落7日勋章数, 每日更新, 战斗前记录快照.
// 部落战力值(成员站力值(血量)快照)战斗前记录快照,更新战力值排名 (不满足参赛条件的不仅算战力值,显示未上榜)(在什么范围内的排名? 怎么分组匹配?待明确). 战斗中加入新成员、移除旧成员战力值(血量)怎么处理,待明确. 已定义用户战斗期间加入部落的无法参与此次战斗.
// 活动休战时计算保存活动积分(经验值)排名, 休战期展示.

// 用户服务器天数?
// 用户是否有部落?
// 单周还是双周举行活动?

// 等级对应血量加成读配置表.
// 存储参赛部落城堡经验分(有上限),  取模取整[设计成 method 不存储]得到当前等级和当前等级升下一层级的累积经验值; 城堡经验分有最大值对应满级,显示上“到下一等级升级进度”显示“Max”.
// 捐献有个人奖励, “奖励池随机获得一份”, 奖励池读配置表, 获得后奖励池状态会变吗?
// 备战期重置个人普通捐献配额.
// 竞技场防守阵容战力.
// 战备结束后锁定经验值等级和战力值(血量).
// 捐献之对下一场未开始的战斗贡献经验值(怎么存储?), 如果本周没有新战斗了(最后一次的战斗期) 还允许捐献吗?
// 普通捐献要判断并更新个人捐献配额.
// 普通捐献和高级捐献需要消耗个人所持有的某种具体道具.
// 高级捐献可以批量触发.

// 获取部落成员, 记录成员捐献记录 (http handler request hook 记录下来了吗?).
// 记录本周谁给这个部落城堡捐献过多少次, 如果他们还在这个部落, 那么排序的时候还需要他们的战力值. (可以记录在以部落为key的redis hash上然后本地排序, 也可以直接记录成zset, 后者需要对score编解码战力值不好弄.)
// 记录部落下贡献血量的英雄战斗开始前锁定的血量(战力值)和本次战斗实时剩余血量.

// 怎么分组匹配? [待明确]

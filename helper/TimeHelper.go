package helper

import "time"

type TimeHelper struct {

}

// FormatTime 格式化时间戳为日期格式
func (t TimeHelper) FormatTime(timestamp int64, layout string) string {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return time.Unix(timestamp, 0).Format(layout)
}

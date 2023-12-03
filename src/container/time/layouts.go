package time

const (
	year        = "2006"
	month       = "01"
	date        = "02"
	hour        = "15"
	minute      = "04"
	second      = "05"
	millisecond = "000000"
)

const (
	DateLayout               = year + "-" + month + "-" + date
	TimeLayout               = hour + ":" + minute + ":" + second
	TimeLayoutInMilliseconds = TimeLayout + "." + millisecond
	DateTimeLayout           = DateLayout + " " + TimeLayout
)

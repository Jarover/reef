package models

type Wlevel struct {
	ID        int `json:"id"`
	Point_id  int
	Level     int64
	Offset    int64
	Datetime  string
	Published bool
}

func (Wlevel) TableName() string {
	return "inside_wlevel"
}

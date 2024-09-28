package domain

type Connection struct {
	Audit
	UserID      string   `gorm:"type string; not null; uniqueIndex:idx_breakout_user; default:''"`
	BreakoutID  string   `gorm:"type:uuid; uniqueIndex:idx_breakout_user; not null"`
	Breakout    Breakout `gorm:"foreignKey:BreakoutID; OnDelete:CASCADE"`
	DisplayName string
	Vote        string
	IsOnline    bool `gorm:"int; default:0"`
}
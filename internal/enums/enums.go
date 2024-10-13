package enums

type BulletType int

const (
	Normal     BulletType = 5
	Medium     BulletType = 10
	Super      BulletType = 20
	SuperDuper BulletType = 40
)

func (d BulletType) String() string {
	return map[BulletType]string{
		Normal:     "Normal",
		Medium:     "Medium",
		Super:      "Super",
		SuperDuper: "SuperDuper",
	}[d]
}

type MessageType string

const (
	JOIN         MessageType = "JOIN"
	JOIN_SUCCESS MessageType = "JOIN_SUCCESS"
	POS          MessageType = "POS"
	HIT          MessageType = "HIT"
)

package store

import "swordmaster/types"

var ammoStore = make(map[string]types.Bullet)

func AddBullet(bullet types.Bullet) {
	ammoStore[bullet.GetOrigin()] = bullet
}

func ListBullets() []types.Bullet {
	output := make([]types.Bullet, 0, len(ammoStore))
	for _, ammo := range ammoStore {
		output = append(output, ammo)
	}
	return output
}

func RemoveBullet(origin string) {
	delete(ammoStore, origin)
}

func BulletCount() int {
	return len(ammoStore)
}

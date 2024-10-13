package store

import (
	"swordmaster/types"
	"sync"
)

type AmmoStore map[string]types.Bullet

var ammoStore = make(AmmoStore)
var mu sync.RWMutex

func AddBullet(bullet types.Bullet) {
	mu.Lock()
	ammoStore[bullet.GetOrigin()] = bullet
	mu.Unlock()
}

func ListBullets() AmmoStore {
	return ammoStore
}

func RemoveBullet(origin string) {
	mu.Lock()
	delete(ammoStore, origin)
	mu.Unlock()
}

func BulletCount() int {
	mu.RLock()
	count := len(ammoStore)
	mu.RUnlock()
	return count
}

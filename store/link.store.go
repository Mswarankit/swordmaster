package store

import "swordmaster/types"

var link types.Network

func SetLink(newLink types.Network) {
	link = newLink
}

func GetLink() types.Network {
	return link
}

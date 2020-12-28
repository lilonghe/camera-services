package main

import (
	"lilonghe.net/camera-services/db"
	"lilonghe.net/camera-services/modules/firmware"
)

func main() {
	db.Init()
	firmware.LoadCanonUpdates()
}

package firmware

import (
	"lilonghe.net/camera-services/db"
	"testing"
)

func TestLoadCanonUpdates(t *testing.T) {
	db.Init()
	err := LoadCanonUpdates()
	//if err != nil {
	//	log.Fatal(err)
	//}
	panic(err)
}

package locks

import (
	"fmt"

	"github.com/Moldy-Community/moldy/utils/colors"
	darthdb "github.com/TeoDev1611/darth-db/json"
	"github.com/google/uuid"
)

func WriteLock(url string) {
	id := uuid.New().String()
	data := map[string]interface{}{
		"name": url,
		"type": "url",
		"url":  url,
		"id":   id,
	}
	darthdb.WriteDB("./MoldyFile.lock", "  ", false, data)
	colors.Success("Succesfuly locked the package")
}

func ReadLockAndCheck(name string) {
	content := darthdb.GetAllDataFile(name)
    fmt.Print(content)
}

// The locks generator for installer

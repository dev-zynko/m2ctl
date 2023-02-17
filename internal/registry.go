package internal

import (
	"log"

	"golang.org/x/sys/windows/registry"
)

func CreateRegistryKey() (registry.Key, func()) {
	var access uint32 = registry.ALL_ACCESS
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "Software", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal("Failed opening key", err)
	}

	key, _, err := registry.CreateKey(k, `M2ctl`, access)
	if err != nil {
		log.Fatal("Failed to create registry", err)
	}

	return key, func() {
		var err error
		if err = key.Close(); err != nil {
			log.Fatal("Failed to close key", err)
		}
	}
}

func CreateKey(name string, val string) {
	regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\M2ctl\`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal("Failed to open registry", err)
	}
	err = regKey.SetStringValue(name, val)
	if err != nil {
		log.Fatal("Failed seting key in registry", err)
	}
}

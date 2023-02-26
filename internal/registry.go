package internal

import (
	"fmt"
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

func GetKey(name string) (string, error) {
	regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\M2ctl\`, registry.ALL_ACCESS)
	if err != nil {
		return "", err
	}

	val, _, err := regKey.GetStringValue(name)
	if err != nil {
		return "", err
	}

	return val, nil
}

func DeleteKey(tag string) {
	regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\M2ctl\`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal("Failed to open registry", err)
	}

	regKey.DeleteValue(fmt.Sprintf("M2CTL_HOST_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_USERNAME_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_PASSWORD_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_SSH-KEY-FILE_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_SSH-KEY-PASS_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_GIT-REPO-URL_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_GIT-USER_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_GIT-EMAIL_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_GIT-SSH-KEY-FILE_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_GIT-SSH-KEY-PASS_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_MYSQL-PASS_%s", tag))
	regKey.DeleteValue(fmt.Sprintf("M2CTL_FILES_%s", tag))

}

package internal

import (
	"fmt"
)

func InstallDependencies(pythonVersion string, mysqlVersion string) string {
	return fmt.Sprintf(
		"/usr/sbin/pkg -v &&"+
			"y && "+
			"pkg install -y mysql%s-server &&"+
			"pkg install -y python%s &&"+
			"pkg install -y git &&"+
			"pkg install -y gmake &&"+
			"pkg install -y gdb &&",
		mysqlVersion, pythonVersion,
	)
}

func SecureConfigMysql(mysqlPass string) string {
	return fmt.Sprintf(
		`grep -q 'mysql_enable="YES"' /etc/rc.conf && echo 'String found in file.' || echo "String not found in file. Starting other command..." `,
	)
}

func CompileFliege(threads string) string {
	return fmt.Sprintf(
		"cd /usr/home/source/server/libpoly/src && gmake -j%[1]s && "+
			"cd /usr/home/source/server/libsql/src && gmake -j%[1]s &&"+
			"cd /usr/home/source/server/libthecore/src && gmake -j%[1]s && "+
			"cd /usr/home/source/server/libgame/src && gmake -j%[1]s && "+
			"cd /usr/home/source/server/game/src && gmake clean && "+
			"cd /usr/home/source/server/game/src && gmake -j%[1]s && "+
			"cd /usr/home/source/server/db/src && gmake clean && "+
			"cd /usr/home/source/server/db/src && gmake -j%[1]s && ",
		threads,
	)
}

func ServerCommands(command string, files string) string {
	switch command {
	case "start":
		switch files {
		case "fliegev3":
			return fmt.Sprintf("")
		case "fliegev2":
			return fmt.Sprintf("")
		case "marty":
			return fmt.Sprintf("")
		case "sura-head":
			return fmt.Sprintf("")
		case "reference":
			return fmt.Sprintf("")
		}
	case "stop":
		switch files {
		case "fliegev3":
			return fmt.Sprintf("")
		case "fliegev2":
			return fmt.Sprintf("")
		case "marty":
			return fmt.Sprintf("")
		case "sura-head":
			return fmt.Sprintf("")
		case "reference":
			return fmt.Sprintf("")
		}
	case "reload-quests":
		switch files {
		case "fliegev3":
			return "cd /usr/home/game/share/quest && python make.py \n"
		case "fliegev2":
			return fmt.Sprintf("")
		case "marty":
			return fmt.Sprintf("")
		case "sura-head":
			return fmt.Sprintf("")
		case "reference":
			return fmt.Sprintf("")
		}
	case "clear-logs":
		switch files {
		case "fliegev3":
			return "cd /usr/home/game/share/quest && python make.py \n"
		case "fliegev2":
			return fmt.Sprintf("")
		case "marty":
			return fmt.Sprintf("")
		case "sura-head":
			return fmt.Sprintf("")
		case "reference":
			return fmt.Sprintf("")
		}
	}
	return ""
}

package internal

import (
	"fmt"
)

//Execute Bash  sh setup.sh -t 4 -p 27 -m 56 -p "Test123" -g "github.com" -n "zynko-dev" -e "zynko.dev@proton.me" -s "Test123"

func InstallDependencies(pythonVersion string, mysqlVersion string) string {
	return fmt.Sprintf(
		//"printf 'y\n' | /usr/sbin/pkg -v && "+
		"pkg install -y mysql%s-server && "+
			"pkg install -y python%s && "+
			"pkg install -y git && "+
			"pkg install -y gmake && "+
			"pkg install -y gdb && ",
		mysqlVersion, pythonVersion,
	)
}

// Git pull first time yes to accept Host blueprint
func GitClone(gitUser string, gitEmail string, gitSSHFile string, gitSSHPass, string, gitRepo string, gitServerSourcePath string) string {
	fmt.Sprintf(
		"git config --global user.name '%s' && "+
			"git config --global user.email '%s' && "+
			"printf 'yes\n' | git clone %s && ",
		gitUser, gitEmail, gitRepo,
	)
}

func SecureConfigMysql(mysqlPass string) string {
	return fmt.Sprintf(
		`grep -q 'mysql_enable="YES"' /etc/rc.conf && echo 'String found in file.' || echo 'mysql_enable="YES"' >> /etc/rc.conf && `+
			"sleep 3 && service mysql-server start sleep 3 && "+
			mysqlSecureInstallation(mysqlPass),
		mysqlPass,
	)
}

func mysqlSecureInstallation(mysqlPass string) string {
	return fmt.Sprintf(
		"UPDATE mysql.user SET Password=PASSWORD('%s') WHERE User='root';"+
			"DELETE FROM mysql.user WHERE User='';"+
			"DELETE FROM mysql.user WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1');"+
			"DROP DATABASE IF EXISTS test;"+
			"DELETE FROM mysql.db WHERE Db='test' OR Db='test\\_%';"+
			"FLUSH PRIVILEGES;"+
			"exit;",
		mysqlPass,
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
			return ""
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
			return "killall -1 db game \n"
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
			return ""
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

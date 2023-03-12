#!bin/sh

while getopts t:p:m:q:g:n:e:s:o:f: flag
do
    case "${flag}" in
        t) threads=${OPTARG};;
        p) pythonversion=${OPTARG};;
        m) mysqlversion=${OPTARG};;
        q) mysqlpass=${OPTARG};;
        g) giturl=${OPTARG};;
        n) gitname=${OPTARG};;
        e) gitemail=${OPTARG};;
        s) gitsshpass=${OPTARG};;
        o) githubtoken=${OPTARG};;
        f) files=${OPTARG}
    esac
done
githubtoken = "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"

echo "Threads: $threads";
echo "PythonV: $pythonversion";
echo "MysqlV: $mysqlversion";
echo "MysqlPass: $mysqlpass";
echo "GitUrl: $giturl";
echo "GitName: $gitname";
echo "GitEmail: $gitemail";
echo "GitSSHPass: $gitsshpass";
echo "Files: $files";



install_dependencies() {
    execute_command_with_elapsed_time "Installing pkg" "printf 'y\n' | /usr/sbin/pkg -v"
    execute_command_with_elapsed_time "Installing mysql" "pkg install -y mysql$mysqlversion-server"
    execute_command_with_elapsed_time "Installing python" "pkg install -y python$pythonversion"
    execute_command_with_elapsed_time "Installing gmake" "pkg install -y gmake"
    execute_command_with_elapsed_time "Installing gdb" "pkg install -y gdb"
}

configure_mysql() {
    execute_command_with_elapsed_time "Enabling mysql in rc.conf" "grep -q 'mysql_enable="YES"' /etc/rc.conf && echo 'String found in file.' || echo 'mysql_enable="YES"' >> /etc/rc.conf && sleep 3"
    execute_command_with_elapsed_time "Starting mysql server" "service mysql-server start && sleep 3"
    execute_command_with_elapsed_time "Doing mysql_secure_installation" "printf '' | mysql_secure_installation"
}

create_ssh_key() {
    execute_command_with_elapsed_time "Creating ssh key for git" "printf '\n$gitsshpass\n$gitsshpass\n' | ssh-keygen"
    var=$( cat /root/.ssh/id_rsa.pub)
    echo $var
   execute_command_with_elapsed_time "Uploading ssh key for git" `curl -L -X POST -H "Accept: application/vnd.github+json" -H "Authorization: Bearer $githubtoken" -H "X-GitHub-Api-Version: 2022-11-28" https://api.github.com/user/keys -d '{"title":"ssh-rsa AAAAB3NzaC1yc2EAAA","key":"2Sg8iYjAxxmI2LvUXpJjkYrMxURPc8r+dB7TJyvv1234"}'`
}

compile_fliege_files() {
    echo "cd /usr/home/source/server/libpoly/src && gmake -j$threads"
    execute_command_with_elapsed_time "Compling libpoly" "gmake -C /usr/home/source/server/libpoly/src -j$threads"
    execute_command_with_elapsed_time "Compiling libsql" "gmake -C /usr/home/source/server/libsql/src -j$threads"
    execute_command_with_elapsed_time "Compiling libthecore" "gmake -C /usr/home/source/server/libthecore/src -j$threads"
    execute_command_with_elapsed_time "Compiling libgame" "gmake -C /usr/home/source/server/libgame/src -j$threads"
    execute_command_with_elapsed_time "Game gmake clean" "gmake -C /usr/home/source/server/game/src clean"
    execute_command_with_elapsed_time "Compiling game" "gmake -C /usr/home/source/server/game/src -j$threads"
    execute_command_with_elapsed_time "DB gmake clean" "gmake -C /usr/home/source/server/db/src clean"
    execute_command_with_elapsed_time "Compiling db" "gmake -C /usr/home/source/server/db/src -j$threads"
}

create_sym_links() {
    ln -s /usr/home/source/server/game/game /usr/home/game/Channel1/core1/
    ln -s /usr/home/source/server/game/game /usr/home/game/Channel1/core2/
    ln -s /usr/home/source/server/game/game /usr/home/game/Channel1/core3/
    ln -s /usr/home/source/server/game/game /usr/home/game/Channel99/


    ln -s /usr/home/source/server/game/game /usr/home/game/Loginserver/

    ln -s /usr/home/source/server/db/db /usr/home/game/Datenbank/
}

configure_game_conf() {
    echo $mysqlpass
    echo "sed -i -r 's/admin/$mysqlpass/g' $file2/CONFIG"
    for file in /usr/home/game/Channel* 
    do 
        echo $file
        for file2 in $file/core*
        do
            echo $file2/CONFIG
            sed -i -E 's/admin/$mysqlpass/g' $file2/CONFIG
        done
    done
}

configure_git() {
    execute_command_with_elapsed_time "Adding git user.name" "git config --global user.name '$gitname'"
    execute_command_with_elapsed_time "Adding git user.email" "git config --global user.email '$gitemail'"
    execute_command_with_elapsed_time "Cloning repository" "git clone $giturl"
}

execute_command_with_elapsed_time(){
    starttime=$( /bin/date +%s )
    echo "M2CTL: $1"
    $2
    stoptime=$( /bin/date +%s )
    runtime=$( /bin/expr ${stoptime} - ${starttime} )
    echo "M2CTL: $1 completed in ${runtime} seconds"
}

#install_dependencies
#create_ssh_key
#configure_mysql
#compile_fliege_files
#configure_game_conf
create_sym_links
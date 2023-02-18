#!/bin/bash


source $(dirname $0)/lib/bashsimplecurses/simple_curses.sh
source $(dirname $0)/vars.sh


EXPORT_PATH=${HOME}/pots

N=30

main(){
	action=$1
	max=$2
	limit="limit 0,${max:-10}"
	action='stats'
	names='combo login password'

	window "Last SSH attacks" "red" "40%"
	tac /var/log/ssh-pot.log | head > /tmp/res
	append_file /tmp/res
	endwin

    window "Last FTP attacks" "red" "40%"
    tac /var/log/ftp-pot.log | head > /tmp/res
	append_file /tmp/res
    endwin

    window "Last Telnet attacks" "red" "40%"
    tac /var/log/telnet-pot.log | head > /tmp/res
	append_file /tmp/res
    endwin

    window "Last HTTP attacks" "red" "40%"
    tac /var/log/http.log | head > /tmp/res
	append_file /tmp/res
    endwin
}


main_loop ${N}

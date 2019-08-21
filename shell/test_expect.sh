#!/usr/bin/env expect
set server  [lindex $argv 0];
set password [lindex $argv 1]
set desName [lindex $argv 2]
set cmd [lindex $argv 3]

set homePath $::env(HOME)
set desHomePath home/${desName}

set prompt ${desName}@${desName}

spawn ssh ${server} ${cmd}
set timeout 5
expect {
     "(yes/no)" {send -- yes\r;exp_continue}
     "password" {send -- ${passowrd}\r}
     eof         {exit 0}
}


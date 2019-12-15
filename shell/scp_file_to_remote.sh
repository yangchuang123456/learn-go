#!/usr/bin/env bash

auto_ssh_c () {
    expect -c " set timeout -1;
                spawn scp $1 $2;
                expect {
                    *(yes/no)* {send -- yes\r;exp_continue;}
                    *assword:* {send $3\r;
                                 expect {
                                    *denied* {exit 2;}
                                    eof
                                 }
                    }
                    eof         {exit 0;}
                }
                "
    return $?
}

auto_ssh_c $1 $2 $3
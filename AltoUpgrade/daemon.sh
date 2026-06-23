#!/bin/bash
num=`ps -ef|grep alto_server|wc -l`
echo $num
if [ $num -lt 2 ];then
            echo $num
                cd /root/back_end
                    ./alto_server &
fi


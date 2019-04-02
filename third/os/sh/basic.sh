# param
# $*所有参数 $# 参数个数 $0 命令行 $? 上一行执行结果 $$ 当前进程的id

# 算数运算符
a=$((1+3))
b=$((1+3))
# -eq -lt gt le ge
if [ $a  -eq 4 -a $b -eq 4 ];then
    echo yes
else
    echo no
fi

# && ||
if [[ $a -eq 4 && $b -eq 4 ]];then
    echo yes
else
    echo no
fi

# 字符串运算符
c=hello
d=hello
# -z 非空
if [ c = "hello" -a d = "hello" ];then
    echo yes
else
    echo no
fi


# process
for i in `ls .`;do
    echo $i
done

#for
for ((i=1; i<=100; i ++))
do
	echo $i
done

# skill

# signal
trap 'onCtrlC' INT
function onCtrlC () {
	pstree $$ -p| awk -F"[()]" '{print $2}'| xargs kill -9
	sleep 1s
	exit 1
}

# cat
cat <<EOF > tmp
echo hello1
echo hello2
EOF

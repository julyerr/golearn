# sed
sed -i s/$a/$b/g filename
# 特殊符号
sed -i s#$a#$b#gi filename
# 指定行数添加内容
sed -i "nicontent" filename
# 匹配到某行
sed -i "/$a/i\$b" filename
sed -i "/$a/a\$b" filename

# awk
awk '{print $9}'

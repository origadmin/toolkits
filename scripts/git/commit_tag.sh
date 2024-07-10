#!/bin/bash

# 设置 Git 用户信息
git config user.name "GitHub Action"
git config user.email "action@github.com"

#git config user.name "godcong"
#git config user.email "jumbycc@163.com"

# 获取当前HEAD的哈希值
CURRENT_COMMIT=$(git rev-parse HEAD)

# 查找与当前提交匹配的标签
TAG_FOR_CURRENT_COMMIT=$(git tag --points-at "$CURRENT_COMMIT")

# 检查是否有标签
if [ -n "$TAG_FOR_CURRENT_COMMIT" ]; then
   # 将标签列表转换为数组
   IFS=$'\n' read -d '' -r -a TAGS_ARRAY <<< "$TAG_FOR_CURRENT_COMMIT"
   # 找到最新的标签，假设标签按时间顺序排列
   LATEST_TAG=${TAGS_ARRAY[-1]}
   echo "$LATEST_TAG"
fi
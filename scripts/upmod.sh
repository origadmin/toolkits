#!/bin/bash

# 记录原始工作目录
original_dir=$(pwd)

# 定义一个函数，用于检查目录中是否存在 go.mod 文件，并执行相应操作
check_go_mod_and_act() {
    local dir="$1"
    local go_mod_name="go.mod"
    local updated=1  # 默认未更新

    # 进入目录
    cd "$dir" || return 0

    # 检查目录中是否存在 go.mod 文件
    if [ -f "$go_mod_name" ]; then
        # 如果存在，执行特定操作
        echo "Processing updates in: $dir"
        go get -u ./...
        go mod tidy

        # 假设 go mod tidy 改变了 go.mod 文件，设置为更新
        updated=0
    fi

    cd "$original_dir" || return $updated
    # 返回更新状态
    return $updated
}

# 定义一个函数，用于提交 go.mod 和 go.sum 文件到 Git
git_commit_changes() {
    local dir="$1"

    # 进入目录
    cd "$dir" || return

    # 添加 go.mod 和 go.sum 到 Git 的暂存区
#    git add go.mod go.sum

    # 构建提交信息，包含模块名称
    local module_name
    module_name=$(echo "$dir" | sed "s/^.\///")  # 去掉开头的'./'
    local commit_message="feat($module_name/go.mod): Update go.mod and go.sum"
#    git commit -m "$commit_message"
    echo "Committing changes in: $dir with $module_name/go.mod"
    # 提交更改
#    git commit -m "$commit_message" || true

    # 返回原始工作目录
    cd "$original_dir" || return
}

# 定义一个函数，用于遍历目录并应用 check_go_mod_and_act 函数
update_go_mod() {
    # 跳过根目录（.）
    find . -mindepth 1 -type d | while read -r dir; do
        if check_go_mod_and_act "$dir"; then
            local updated=$?
            if [ $updated -eq 0 ]; then
                git_commit_changes "$dir"
            fi
        fi
    done
}

# 调用函数
update_go_mod

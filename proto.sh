#!/usr/local/bin bash

dir=$(pwd)                      # 当前目录
protoPath="$(pwd)/Proto/protos" # proto文件目录
pbPath="$(pwd)/Proto/pb"        # pb文件目录

# 进入proto文件目录
# shellcheck disable=SC2164
cd "$protoPath"

# 生成pb文件
protoc --micro_out="${pbPath}" --go_out="${pbPath}" ./*.proto

# 进入pb目录
# shellcheck disable=SC2164
cd "${pbPath}"

# 循环生成注释
# shellcheck disable=SC2231
# shellcheck disable=SC2006
for file in $(pwd)/*.pb.go; do
  # 生成json注释
  protoc-go-inject-tag --input="$file"
done

# 返回目录
# shellcheck disable=SC2164
cd "$dir"

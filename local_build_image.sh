#!/bin/zsh

DIR=$(cd $(dirname $0)/; pwd)
TAG=${2:-latest}

build_image () {
  docker build -t $1:$TAG $DIR/$1/.
}

eval $(minikube -p minikube docker-env)

if [ "$1" = "all" ]; then
#  build_image golan-grpc
  build_image gin-rest
  build_image nginx-temp
  build_image nginx-proxy
  exit 0
elif [ ! -d $DIR/$1 ] || [ "$1" = "" ]; then
  echo 'ディレクトリ名を引数に入れてください'
  exit 1
else
  build_image $1
fi

if [ $? -ne 0 ]; then
  echo $1'のbuildに失敗しました'
  exit 1
fi

#!/bin/zsh

DIR=$(cd $(dirname $0)/; pwd)
TAG=${2:-latest}
build_image_and_minikube_load () {
  docker build -t $1 $DIR/$1/.
  minikube image load $1:$TAG
}

if [ "$1" = "all" ]; then
  build_image_and_minikube_load djan-rest
  build_image_and_minikube_load lara-graph
  build_image_and_minikube_load rail-grpc
  build_image_and_minikube_load nginx-temp
  exit 0
fi

if [ ! -d $DIR/$1 ] || [ "$1" = "" ]; then
  echo 'ディレクトリ名を引数に入れてください'
  exit 1
fi

build_image_and_minikube_load $1

if [ $? -ne 0 ]; then
  echo $1'のbuildに失敗しました'
  exit 1
fi

# kubernetesのお試し

```
.
├── README.md
├── local_build_image.sh              コンテナイメージ作成shell
├── gin-rest/                         gin,gormでREST API
├── gin-rest-deployment-service.yml   上記サービスのマニュフェスト詰め込み
├── nginx-temp/                       適当な静的資材を返すnginx
├── nginx-temp-deployment-service.yml 上記サービスのマニュフェスト詰め込み
├── nginx-proxy/                      環境変数で適当にプロキシするnginx
├── golan-grpc/                       WIP: grpc適当に、、、
├── mongo/                            WIP: 何に使うかも未定
├── postgres/                         gin-restのDB（予定）
└── redis/                            WIP: セッションやキャッシュに利用予定
```

## ローカルでの起動

require
- minikube
- kubectl

```shell
$ minikube start
$ kubectl apply -f gin-rest-deployment-service.yml
$ kubectl port-forward svc/gin-rest-svc 8081:8081
```


nginxのPVの`/path/to/your/src`を修正

```shell
$ minikube start
$ minikube mount /path/to/your/kube-training/nginx-temp/src:/mnt/nginx-temp
$ kubectl apply -f nginx-temp-deployment-service.yml
$ kubectl port-forward svc/nginx-temp-svc 8080:8080
```

## 調査用メモ

- exec  
`kubectl exec --stdin --tty pod-name -- /bin/bash`

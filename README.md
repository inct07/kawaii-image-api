# 概要
画像API

# 構築手順
※ Go環境が整っている前提
```
$ mkdir $GOPATH/src/github.com/inct07/
$ cd $GOPATH/src/github.com/inct07/
$ git clone git@github.com:inct07/kawaii-image-api.git
$ cd $GOPATH/src/github.com/inct07/kawaii-image-api

# vendoring
$ dep init

# ローカル起動
$ go run main.go
```

# Docker
```
# Docker Imageをビルド
docker build -t kawaii-image-api .

# 実行
docker run -p 8080:8080 kawaii-image-api
```

# Container Registryにプッシュ
```
## tag付け
docker tag kawaii-image-api asia.gcr.io/[YOUR_PROJECT_ID]/kawaii-image-api

## プッシュ
gcloud docker -- push asia.gcr.io/[YOUR_PROJECT_ID]/kawaii-image-api
```

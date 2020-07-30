# moji-chat
文字入力のチャット

## AWS環境構築

### S3バケットの作成
テンプレートを配置するS3バケットを作成する。  
CUIで作成する場合は、下記コマンドを実行。

```
aws s3 mb s3://moji-chat-stack
```

### envファイル生成
```
cp parameters/env.exmaple parameters/.env
```


### VPC構築
VPCやElasticCacheは毎回デプロイする必要はないため、SAMとは別に管理しています。  
アプリケーションを作成する前に、CloudFormationで作成しましょう。

```
./init.sh
```
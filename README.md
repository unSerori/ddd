# ddd

qiitaで書いたDDD備忘録記事のサンプルリポジトリ

TODO: qiitaリンク

## 概要

DDDを採用したGo-ginでのAPIサーバ

### 環境

Visual Studio Code: 1.88.1  
golang.Go: v0.41.4  
image Golang: go version go1.22.2 linux/amd64

## 環境構築

[ddd-docker](https://github.com/unSerori/ddd-docker)を使ってDokcerコンテナーで開発・デプロイする形を想定している  
インストール手順は[ddd-dockerのインストール項目](https://github.com/unSerori/ddd-docker/blob/main/README.md#インストール)に記載  
cloneしてスクリプト実行で、自動的にコンテナー作成と開発環境（: またはデプロイ）を行う  

### 自前でのローカル環境構築

1. [Goのインストール](https://go.dev/doc/install)
2. このリポジトリをclone

    ```bash
    git clone https://github.com/unSerori/ddd
    ```

3. [.env](#env)ファイルを作成
4. 必要なパッケージの依存関係など

    ```bash
    go mod tidy
    ```

5. プロジェクトを起動

    ```bash
    # 実行(VSCodeならF5keyで実行)
    go run .

    # ワンファイルにビルド
    go build -o output 

    # ビルドで生成されたファイルを実行
    ./output
    ```

#### vscode-ext.txt

Goのデバッグ用VS Code拡張機能や便利な拡張機能のリスト  
以下のコマンドで一括インストールできる

```bash
cat vscode-ext.txt | while read line; do code --install-extension $line; done
```

#### おまけ: Goでプロジェクト作成時のコマンド

```bash
# goモジュールの初期化
go mod init ddd

# ginのインストール
go get -u github.com/gin-gonic/gin

# main.goの作成
echo package main > main.go
```

## ディレクトリ構成

TODO: ディレクトリ構成

```bash
tree --dirsfirst
```

## ※開発手順について

0. 具体的な要件（:ビジネスルール）を最初に決める
1. 次に具体的なビジネスロジックとしてdomain/entity.go, domain/logic.go, domain/repository.goに定義

パターン1: 完全に下の層から書いていく: DDDの趣旨にはあってるのか。。？

1. domain/repositoryで提供されているインターフェイスをinfrastructure層で実装、domain/logic.goから呼び出す
2. application層でロジックを呼び出す
3. presentation層で受け取りと完成したapplication層の処理を呼び出す
4. routingに追加

パターン2: 下から上の依存関係を登録し、実際の細かい処理は上から処理書く: 開発途中でもエンドポイントに対するテストが行えるし、処理の流れを意識しながらコーディングできる

1. まだ作ってないならファイル（4層分）作成と依存関係の登録
2. ドメインロジックを利用するアプリケーション関数とそれを呼び出すプレゼンテーション関数とそれを呼び出すルーティングを行い依存性テスト
3. 荒く上の層から書く
    2. presentation層でのリクエスト受け取りapplication層に処理渡しレスポンス返却
    3. application層でのユースケースの流れ、ビジネスロジックの呼び出し
    4. domain/repository.goで提供されているインターフェイスをinfrastructure層で実装
    5. 全体の体裁を整える

## API仕様書

エンドポイント、リクエストレスポンスの形式、その他情報のAPIの仕様書。

### エンドポインツ

<details>
  <summary>疎通確認するエンドポイント</summary>

- **URL:** `/check/echo`
- **メソッド:** GET
- **説明:** エンドポイントにアクセスすると、その際のリクエスト情報をサーバーデバッグコンソールに流し、クライアント側にもJSON形式で返す
- **リクエスト:**
  - ヘッダー:
  - ボディ:
- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:

      ```json
      {
        "srvResMsg":  "OK",
        "srvResData": {
            "info": {
              "body": {
              },
              "header": {
              },
              "method": "GET",
              "url": {
                "Scheme": "",
                "Opaque": "",
                "User": null,
                "Host": "",
                "Path": "/check/echo",
                "RawPath": "",
                "OmitHost": false,
                "ForceQuery": false,
                "RawQuery": "",
                "Fragment": "",
                "RawFragment": ""
              },
              "url query": {
              }
            },
            "message": "hello go server!"
        },
      }
      ```

</details>

<details>
  <summary>開発時にサーバー内の処理を確認したい時にアクセスするエンドポイント</summary>

- **URL:** `/check/sandbox`
- **メソッド:** GET
- **説明:** 開発時のサーバーで確認したい処理を実行するためのもの
- **リクエスト:**
  - ヘッダー:　実行に必要なさまざまな形式のヘッダー
  - ボディ: 実行に必要なさまざまな形式のボディ値

- **レスポンス:**

</details>

TODO: エンドポイント仕様

### API仕様書てんぷれ

<details>
  <summary>＊○○＊するエンドポイント</summary>

- **URL:** `/＊エンドポイントパス＊`
- **メソッド:** ＊HTTPメソッド名＊
- **説明:** ＊○○＊
- **リクエスト:**
  - ヘッダー:
    - `＊HTTPヘッダー名＊`: ＊HTTPヘッダー値＊
  - ボディ:
    ＊さまざまな形式のボディ値＊

- **レスポンス:**
  - ステータスコード: ＊ステータスコード ステータスメッセージ＊
    - ボディ:
      ＊さまざまな形式のレスポンスデータ（基本はJSON）＊

      ```json
      {
        "srvResMsg":  "レスポンスステータスメッセージ",
        "srvResData": {
        
        },
      }
      ```

</details>

## .ENV

.evnファイルの各項目と説明

```env:.env
MYSQL_USER=DBに接続する際のログインユーザ名: ddd_user
MYSQL_PASSWORD=パスワード: ddd_pass
MYSQL_HOST=ログイン先のDBホスト名。dockerだとサービス名。mysql-db-srv
MYSQL_PORT=ポート番号（dockerだとコンテナのポート）: 3306
MYSQL_DATABASE=使用するdatabase名: ddd_db
JWT_SECRET_KEY="openssl rand -base64 32"で作ったJWTトークン作成用のキー
JWT_TOKEN_LIFETIME=JWTトークンの有効期限: 315360000
MULTIPART_IMAGE_MAX_SIZE=Multipart/form-dataの画像の制限サイズ（10MBなら10485760）: 10485760
REQ_BODY_MAX_SIZE=リクエストボディのマックスサイズ（50MBなら52428800）: 52428800
```

## TODO

empty

## 開発者

- Author:[unSerori]
- Mail:[]

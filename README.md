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

Docker-Desktopでコンテナーを立てて開発している
(この開発環境は一例であり、適宜読み替えてください。)

1. dockerで環境を作る
    1. Docker Desktopをインストール  [Docker Hub windows-install download](https://docs.docker.com/desktop/install/windows-install/)
    2. compose.ymlを`./`直下に作成

        ```yml:compose.yml
        services:  # コンポースするサービスたち
            go_server:  # サービスの名前
                image: golang:1.22.2-bullseye  # pullするイメージ
                container_name: ddd-api  # コンテナ名
                ports:
                    - "4562:4561" # ホストマシンのポートとコンテナのポートをマッピング
                environment:
                    TZ: ${TZ}
                volumes:  # ボリュームの保持。
                    - ./go_server/ddd:/root/ddd
                build:  # ビルド設定
                    context: ./go_server  # ビルドプロセスが実行されるパス。
                    dockerfile: Dockerfile  # Dockerfileのパス。
                tty: true  # コンテナ内で対話的操作を可能
            mysql_server: 
                image: mysql:latest
                container_name: ddd_server
                restart: always
                environment:
                    TZ: ${TZ}
                env_file: 
                    - .env
                ports:
                    - "3308:3306"  # ホストマシンのポートとコンテナのポートをマッピング
                volumes:
                    - ./mysql_server/db-data:/var/lib/mysql
                # build:
                #     context: .
                #     dockerfile: ./mysql_server/Dockerfile
        ```

    3. `./go_server`を作成しDockerfileを`./go_server`に置く

        ```Dockerfile:Dockerfile
        # ベースイメージ
        FROM golang:1.22.2-bullseye

        # aptでパッケージリストの更新とgitインストール
        RUN apt update && apt install -y git

        # デバッガインストール
        RUN go install github.com/go-delve/delve/cmd/dlv@latest
        ```

    4. .envを`./`直下に作成

        ```env:.env
        MYSQL_ROOT_PASSWORD: mysql_serverのルートユーザーパスワード: root
        MYSQL_USER: ユーザー名: ddd_user
        MYSQL_PASSWORD: MYSQL_USERのパスワード: ddd_pass
        MYSQL_DATABASE: 使用するdatabase名: ddd_db
        TZ: タイムゾーン: Asia/Tokyo
        ```

    5. VScodeでカレントディレクトリを開きDocker拡張機能（ms-azuretools.vscode-docker）を入れ、コマンドで起動

        ```bash:compose build
        docker compose up -d --build
        ```

    6. Dockerタブからgo_serverをアタッチ
2. shareディレクトリ内で以下のコマンドを実行し拡張機能をインストール

    ```bash:Build an environment
    # vscode 拡張機能を追加　vscode-ext.txtにはプロジェクトごとに必要なものを追記している。  
    cat vscode-ext.txt | while read line; do code --install-extension $line; done
    ```

3. ssh通信設定を行い、/root/ddd直下で以下のコマンドを打ちこのリポジトリをクローン

    ```bash:clone
    git clone git@github.com:unSerori/ddd.git　.
    ```

4. サーバ起動

参考までにgo-gin自体の環境構築は以下

```bash:go-gin start-up
# goがインストールされているコンテナイメージなのでgoインストールは不要

# goモジュールの初期化
go mod init ddd

# ginのインストール
go get -u github.com/gin-gonic/gin

# main.goの作成
echo package main > main.go
```

[goインストール方法](https://go.dev/doc/install)

## ディレクトリ構成

TODO: ディレクトリ構成

## API仕様書

エンドポイント、リクエストレスポンスの形式、その他情報のAPIの仕様書。

### エンドポインツ

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

## SERVER ERROR MESSAGE

サーバーレスポンスメッセージとして"srvResMsg"キーでメッセージを返す。  
サーバーレスポンスステータスコードと合わせてデバックする。

## .ENV

.evnファイルの各項目と説明

```env:.env
MYSQL_USER=DBに接続する際のログインユーザ名
MYSQL_PASSWORD=パスワード
MYSQL_HOST=ログイン先のDBホスト名。dockerだとサービス名。
MYSQL_PORT=ポート番号。dockerだとコンテナのポート。
MYSQL_DATABASE=使用するdatabase名
JWT_SECRET_KEY="openssl rand -base64 32"で作ったJWTトークン作成用のキー。
JWT_TOKEN_LIFETIME=JWTトークンの有効期限
MULTIPART_IMAGE_MAX_SIZE=Multipart/form-dataの画像の制限サイズ。10MBなら10485760
REQ_BODY_MAX_SIZE=リクエストボディのマックスサイズ。50MBなら52428800
```

## TODO

empty

## 開発者

- Author:[unSerori]
- Mail:[]

# techbook-qrcode

<https://techbookfest.org/event/tbf06/circle/46360001>

エムスリーテックブック#1のサンプルアプリです。
1つのURLに対して、複数のURLを対応させてダウンロードカードが生成できます。

## 依存するツールなど

下記の依存ライブラリのインストールに加え、protobufとGoogle Cloud SDKが必要です。

    # JavaScript
    cd front; yarn install

    # Go
    cd src; go get

## デプロイ方法

src/app_production.yaml, src/app_qa.yamlを作成します。

    cp src/app.yaml.sample src/app_production.yaml
    emacs src/app_production.yaml

下記のコマンドでデプロイします。

    QRCODE_PROD_PROJECT=<your-app-engine-project-name> make app-deploy-production 

## ローカル環境での実行

src/app_development.yamlを作成後、下記コマンドから開発環境を起動できます。

    make build-js
    make dev-go # API開発環境の起動

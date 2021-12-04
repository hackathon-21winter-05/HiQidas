# 開発環境の建て方

どちらも

- サーバー: `ポート7070`
- クライアント: `ポート7071`
- phpMyAdmin: `ポート7072`

で立ちます

## DevContainer

### 要件

- Docker
  - Docker Compose
- [Visual Studio Code](https://azure.microsoft.com/ja-jp/products/visual-studio-code/) (VS Code)
  - [Remote Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containershttps://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) (拡張機能)
    - 識別子: ms-vscode-remote.remote-containers

### 開発環境の建て方

1. VSCodeでこのリポジトリを開く
1. Remote Containers が入っていれば、`Folder contains a Dev Container configuration file. Reopen folder to develop in a container`と表示されるので、`Reopen in Container` をクリックする
1. しばらく待って、左下の緑色のところが`Dev Container: naro.devcontainer`となればOK

### コマンド

- ローカルで叩くもの (Docker系)
  - `make reset-frontend` フロントエンドをのコンテナをリセットする (アップデート用)
  - `make prune`          不要なイメージとボリュームを削除
  - `make chown`          コンテナ内で作られて作成者がrootになっているファイルの作成者をローカルユーザーに変更する
- コンテナ内で叩くもの (Go系)
  - `make lint`           リント
  - `make build`          ビルド
  - `make run`            実行
  - `make protobuf`       Protocol Buffers定義をコンパイル
  - `make tbls`           DBドキュメントの生成

## ホットリロード環境

### 要件

- Docker
  - Docker Compose

### コマンド

- Docker系
  - `make up`               ホットリロード環境を起動する
  - `make down`             ホットリロード環境を削除する
  - `make reset-frontend`   フロントエンドをのコンテナをリセットする (アップデート用)
  - `make prune`            不要なイメージとボリュームを削除
- Go系
  - `make lint`             リント
  - `make build`            ビルド
  - `make run`              実行
  - `make protobuf-docker`  Protocol Buffers定義をコンパイル
  - `make tbls-docker`      DBドキュメントの生成

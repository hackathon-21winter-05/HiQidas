# HiQidasのDBスキーマ

## user

ユーザー情報

| Field        | Type        | Nullable |  Key   | Default | PARENT |        説明        |
| ------------ | ----------- |:--------:|:------:| ------- |:-----:|:------------------:|
| id           | char(36)    |  false   |  PRI   |         |       |     UserのUUID     |
| name(traqid) | varchar(32) |  false   | unique |         |       |    UserのtraPID / ID |
| icon_file_id   | char(36)    |   true   | unique | NULL    |       | userのアイコンUUID |

IconFileIDがNULLなのは考える

## history

閲覧履歴

|   Field    |   Type   | Nullable | Key | Default | PARENT | 説明                   |
|:----------:|:--------:| -------- | --- |:-------:|:------:| ---------------------- |
|   user_id   | char(36) | false    | PRI |         | users  | UserのUUID             |
|  heya_id   | char(36) | false    | PRI |         | sheet  | SheetのUUID            |
| last_access | datetime | true     |     |         |        | シートの最後の閲覧日時 |

## Heya

シート(仮称)情報

|    Field     |   Type   | Nullable | Key |      Default      | PARENT |              説明              |
|:------------:|:--------:|:--------:| --- |:-----------------:| ------ |:------------------------------:|
|      id      | char(36) |  false   | PRI |                   |        |          SheetのUUID           |
|    title     | char(50) |  false   |     |                   |        |        シートのタイトル        |
| description  |   Text   |   true   |     |                   |        |          シートの説明          |
|  creator_id   | char(36) |  false   | MUL |                   | users  |      シートの作成者のUUID      |
| last_editor_id | char(36) |   true   | MUL |                   | users  | 最後に編集をしたユーザーのUUID |
|  created_at   | datetime |  false   |     | CURRENT_TIMESTAMP |        |        シートの作成日時        |
|  updated_at   | datetime |   false   |     |   CURRENT_TIMESTAMP |        |        シートの最終更新日時        |
|  deleted_at   | datetime |   true   |     |                   |        |        シートの削除日時        |

## hiqidashi

ヒキダシ情報

|    Field     | Type     | Nullable     | Key |      Default      |  PARENT   |           説明           |
|:------------:| -------- | ------------ |:---:|:-----------------:|:---------:|:------------------------:|
|      id      | char(36) | false        | PRI |                   |           |     HiqidashiのUUID      |
|   sheet_id    | char(36) | false        | MUL |                   |   sheet   |       シートのUUID       |
|  creator_id   | char(36) | false        | MUL |                   |   users   |   シートの作成者のUUID   |
| last_editor_id | char(36) | true         | MUL |                   |   users   | シートの最終編集者のUUID |
|   parent_id   | char(36) | true         | MUL |                   | Hiqidashi |    親HiqidashiのUUID     |
|    title     | char(50) | false        |     |                   |           |   Hiqidashiのタイトル    |
| description  | Text     | true(false?) |     |                   |           |     Hiqidashiの説明      |
|   image_id    | char(36) | true         | MUL | NULL              |   Image   |  Hiqidashiの画像のUUID   |
|  created_at   | datetime | false        |     | CURRENT_TIMESTAMP |           |   Hiqidashiの作成日時    |
|  updated_at   | datetime |   false   |     |   CURRENT_TIMESTAMP |           |   Hiqidashiの最終更新日時    |
|  deleted_at   | datetime |   true   |     |                   |        |        シートの削除日時        |
## image


ヒキダシにつく画像の情報 (オブジェクトストレージ代わり)

| Field |    Type    | Nullable | Key | Default | PARENT |    説明     |
| ----- |:----------:| -------- | --- | ------- |:------:|:-----------:|
| id    |  char(36)  | false    | PRI |         |        | ImageのUUID |
| image | MEDIUMBLOB | false    |     |         |        |  画像本体   |

## tuna

ヒキダシ同士のコネクション情報

| Field      | Type     | Nullable | Key | Default | PARENT    | 説明                               |
| ---------- | -------- |:--------:|:---:|:-------:| --------- | :---------------------------------: |
| id         | char(36) |  false   | PRI |         |           | ConnectionのUUID                   |
| hiqidashi_one | char(36) |  false   | MUL |         | Hiqidashi | 接続した片方のHiqidashiのUUID    |
| Hiqidashi_two | char(36) |  false   | MUL |         | Hiqidashi | 接続したもう片方のHiqidashiのUUID |


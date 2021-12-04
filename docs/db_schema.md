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
|  heya_id   | char(36) | false    | PRI |         | heya  | ヘヤのUUID            |
| last_access | datetime | true     |     |         |        | ヘヤの最後の閲覧日時 |

## heya

ヘヤ情報

|    Field     |   Type   | Nullable | Key |      Default      | PARENT |              説明              |
|:------------:|:--------:|:--------:| --- |:-----------------:| ------ |:------------------------------:|
|      id      | char(36) |  false   | PRI |                   |        |         heyaのUUID           |
|    title     | char(50) |  false   |     |                   |        |        ヘヤのタイトル        |
| description  |   Text   |   true   |     |                   |        |          ヘヤの説明          |
|  creator_id   | char(36) |  false   | MUL |                   | users  |      ヘヤの作成者のUUID      |
| last_editor_id | char(36) |   false   | MUL |                   | users  | 最後に編集をしたユーザーのUUID |
|  created_at   | datetime |  false   |     | CURRENT_TIMESTAMP |        |       ヘヤの作成日時        |
|  updated_at   | datetime |   false   |     |   CURRENT_TIMESTAMP |        |        ヘヤの最終更新日時        |
|  deleted   | bool |   false   |     |     false              |        |        ヘヤの削除したかどうかの判定        |

## hiqidashi

ヒキダシ情報

|    Field     | Type     | Nullable     | Key |      Default      |  PARENT   |           説明           |
|:------------:| -------- | ------------ |:---:|:-----------------:|:---------:|:------------------------:|
|      id      | char(36) | false        | PRI |                   |           |     HiqidashiのUUID      |
|   heya_id    | char(36) | false        | MUL |                   |   sheet   |       ヘヤのUUID       |
|  creator_id   | char(36) | false        | MUL |                   |   users   |   ヘヤの作成者のUUID   |
| last_editor_id | char(36) | false         | MUL |                   |   users   | ヘヤの最終編集者のUUID |
|   parent_id   | char(36) | true         | MUL |                   | Hiqidashi |    親HiqidashiのUUID     |
|    title     | char(50) | false        |     |                   |           |   Hiqidashiのタイトル    |
| description  | Text     | true         |     |                   |           |     Hiqidashiの説明      |
|   image_id    | char(36) | true         | MUL | NULL              |   Image   |  Hiqidashiの画像のUUID   |
|  created_at   | datetime | false        |     | CURRENT_TIMESTAMP |           |   Hiqidashiの作成日時    |
|  updated_at   | datetime |   false   |     |   CURRENT_TIMESTAMP |           |   Hiqidashiの最終更新日時    |
## image


hiqidashiにつく画像の情報 (オブジェクトストレージ代わり)

| Field |    Type    | Nullable | Key | Default | PARENT |    説明     |
| ----- |:----------:| -------- | --- | ------- |:------:|:-----------:|
| id    |  char(36)  | false    | PRI |         |        | ImageのUUID |
| image | MEDIUMBLOB | false    |     |         |        |  画像本体   |

## tsuna

hiqidashi同士のコネクション情報

| Field      | Type     | Nullable | Key | Default | PARENT    | 説明                               |
| ---------- | -------- |:--------:|:---:|:-------:| --------- | :---------------------------------: |
| id         | char(36) |  false   | PRI |         |           | tsunaのUUID                   |
| hiqidashi_one | char(36) |  false   | MUL |         | hiqidashi | 接続した片方のHiqidashiのUUID    |
| hiqidashi_two | char(36) |  false   | MUL |         | hiqidashi | 接続したもう片方のHiqidashiのUUID |


# エンドポイント・WS設計とそのスキーマ

## エンドポイント
### 各情報


| model.user   | type     |
| ------- | -------- |
| id | string |
| name | string |
|icon_file_id | string |

|   model.heya   | type                      |
|:--------------:| ------------------------- |
|       id       | string                    |
|     title      | string                    |
|  description   | string                    |
|   creator_id   | string                    |
| last_editor_id | string                    |
|   created_at   | google.protobuf.Timestamp |
|   updated_at   | google.protobuf.Timestamp |

### users
#### `GET /users` ユーザー一覧を取得

| Field   | type     |
| ------- | -------- |
| user_id | []string |

#### `POST /users`  ユーザーを作成

Request

| Field   | type     |
| ------- | -------- |
| name | string |

Response

| Field   | type     |
| ------- | -------- |
| user | model.user |

#### `GET /users/me` 自分の情報を取得

| Field   | type     |
| ------- | -------- |
| user | model.user |

#### `GET /users/me/heyas` 自分の作成したヘヤを取得

| Field   | type     |
| ------- | -------- |
| heya | []model.heya |

#### `GET /users/me/favorites` 自分のお気に入りの部屋の一覧を取得　

| Field   | type     |
| ------- | -------- |
| favorite_heya_id | []string |

#### `GET /users/{userID}` ユーザーの詳細情報を取得

| Field   | type     |
| ------- | -------- |
| user | model.user |

### heyas
#### `GET /heyas` 全ヘヤの情報を取得

| Field   | type     |
| ------- | -------- |
| heyas | []model.heya |

#### `GET /heyas/{heyaID}` ヘヤの情報を取得

| Field   | type     |
| ------- | -------- |
| heya | model.heya |

#### `GET /heyas/{heyaID}/users`　ヘヤにいるuserを取得

| Field   | type     |
| ------- | -------- |
| user_id | string |

#### `POST /heyas`　ヘヤを作成する
Request

| Field   | type     |
| ------- | -------- |
| title | string |
| description | string |

Response

| Field   | type     |
| ------- | -------- |
| heya | model.heya |

#### `PUT /heyas/{heyaID}` ヘヤの情報を変更
Request

| Field   | type     |
| ------- | -------- |
| title | string |
| description | string |

#### `PUT /heyas/{heyaID}/favorite` ヘヤのお気に入りの変更
Request 

| Field   | type     |
| ------- | -------- |
| is_favorite | bool |

#### `DELETE /heyas/{heyaID}` ヘヤの削除

| Field   | type     |
| ------- | -------- |
| statusCode | int |

### oauth

#### `GET /aouth/callback` OauthのCallbackのステータス

| Field   | type     |
| ------- | -------- |
| uri | string |

#### `POST /aouth/code` Codeを渡してアクセストークンを取得

Field   | type     |
| ------- | -------- |
| Code | string |

## WebSocket

### ヘヤWebSocket

- 接続時  
`cli` <- (WsSendHiqidashis) <- `ser` ...全ヒキダシ情報を送信
- ヒキダシ作成  
`cli` -> (WsCreateHiqidashi) -> `ser` ...親IDを送る  
`cli (同じHeya)` <- (WsSendHiqidashi) <- `ser` ...作成したヒキダシ情報を送信
- ヒキダシ削除  
`cli` -> (WsDeleteHiqidashi) -> `ser` ...IDを送る  
`cli (同じHeya)` <- (WsDeleteHiqidashi) <- `ser` ...IDを送る
- ヒキダシ編集  
`cli` -> (WsEditHiqidashi) -> `ser` ...IDと編集内容を送る  
`cli (同じHeya)` <- (WsEditHiqidashi) <- `ser` ...IDと編集内容を送る

### Yjs WebSocket

- 編集  
`cli` -> (YjsのDiff) -> `ser` ...YjsのDiffをそのまま送る  
`cli (自分以外の同じHiqidashi)` <- (YjsのDiff) <- `ser` ...YjsのDiffをそのまま送る

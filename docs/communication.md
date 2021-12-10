# エンドポイント・WS設計とそのスキーマ

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

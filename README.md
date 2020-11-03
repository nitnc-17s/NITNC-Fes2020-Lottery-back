# 抽選会システム

## フロントエンド

`https://crow31415.net`

### /view/public

#### 概要

- WebSocketにより表示情報を受信
- 操作機能はなし
- 配信に載せる画面

#### 表示項目

- 景品番号
- 景品内容
- 当選者 氏名
- 当選者 クラス

### /view/monitor

#### 概要

- WebSocketにより表示情報を受信
- 操作機能はなし
- 配信時に司会が確認する用

#### 表示項目

- 景品番号
- 景品内容
- 当選者 氏名(ふりがな)
- 当選者 クラス

### /operation

#### 概要

- `/view`の操作用
- ぶっちゃけUIはカスでいい

#### 必要項目

- APIキー
  - 実質操作用PW
  - `input type="text`
- 操作対象の景品番号
  - `input tyme="number`
- 操作命令発火ボタン
  - 景品番号のみ表示
  - 景品名表示
  - 抽選実行 & 当選者表示
  - 画面削除 (無を出力)
- 「次へ」ボタン
  - 以下の処理を順番に回す
    1. 景品番号++ & 景品番号の表示命令
    2. 景品名の表示命令
    3. 抽選の実行命令

## バックエンド

`https://api.crow31415.net`

### /ws

#### JSON (Server -> Client)

```json
{
  "prize": {
    "id": 0,
    "name": "string"
  },
  "winner": {
    "name": "string",
    "name_furigana": "string",
    "class": "string"
  }
}
```

- `prize_id`
  - 景品番号
  - e.g.) `39`
- `prize_name`
  - 景品名
  - e.g.) `DIPスイッチ`
- `winner_name`
  - 当選者氏名
  - e.g.) `烏野 黒羽`
- `winner_name_furigana`
  - 当選者氏名(ふりがな)
  - e.g.) `からすの くろは`
- `winner_class`
  - 当選者クラス
  - e.g.) `4I`

### /operation

#### JSON (Client -> Server)

```json
{
  "api_key": "string",
  "prize_id": 0,
  "operation": "string"
}
```

- `api_key`
  - APIキー
  - `openssl rand -base64 24` コマンドで生成
  - e.g.) `pD1rfI18T/ob0QyM1zcdtUbSUW7zbkFZ`
- `prize_id`
  - 操作対象の景品番号
  - e.g.) `39`
- `operate`
  - 操作内容
  - コマンド一覧
    - `init`
      - 無を表示する
    - `show id`
      - 景品番号
    - `show prize`
      - 景品番号
      - 景品名
    - `show winner`
      - 景品番号
      - 景品名
      - 当選者名
    - `lottery`
      - 抽選の実施のみ

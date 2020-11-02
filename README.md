# 抽選会システム

## フロントエンド

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

## バックエンド

### /api/ws

#### JSON (Server -> Client)

```json
{
  "prize_id": 0,
  "prize_name":  "string",
  "winner_name": "string",
  "winner_name_furigana": "string",
  "winner_class": "string"
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

### /api/operation

#### JSON (Client -> Server)

```json
{
  "api_key": "string",
  "prize_id": 0,
  "operate": "string"
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
    - `prize-id`
      - 景品番号
    - `prize-name`
      - 景品番号
      - 景品名
    - `lottery`
      - 景品番号
      - 景品名
      - 当選者名

# MY DISCORD BOT

ローカルで起動するディスコードBotです。

## 必要なもの
おそらく dgvoice の依存関係

```shell
brew install pkg-config
brew install ffmpeg --with-opus
```

## 使い方

ディスコードの[開発者ポータル](https://discord.com/developers/applications)でアプリケーションを登録し、Botを作成してください。
BotのTokenを取得したら、Botを自身のサーバーへ招待し、以下のコマンドを実行します。

```shell
$ cd my-discord-bot
$ go run . -t ${token}
2022/06/17 14:43:39 Bot is now running.  Press CTRL-C to exit.
```

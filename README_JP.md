# xiaoniao

[中文](README.md) | [English](README_EN.md) | [한국어](README_KR.md)

Windows クリップボード翻訳ツール

## 機能

- クリップボード監視と自動翻訳
- 多言語UI対応（中/英/日/韓/仏/西/独/露/アラビア語）
- システムトレイ常駐

## インストール

[xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) をダウンロードして実行

## 使い方

1. xiaoniao.exe を実行
2. テキストをコピー (Ctrl+C)
3. 自動翻訳されクリップボードを置換
4. ペースト (Ctrl+V) で翻訳結果を取得

トレイアイコン状態：青-監視中 / 緑-翻訳中 / 赤-停止

## 設定

```cmd
xiaoniao.exe config
```

OpenAI、Anthropic、Google、DeepSeek 等のAPIに対応

## ビルド

```bash
# アイコン付き
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
cd cmd/xiaoniao && goversioninfo -manifest=../../xiaoniao.exe.manifest -icon=../../assets/icon.ico ../../versioninfo.json
cd ../.. && go build -ldflags="-s -w" -o xiaoniao.exe ./cmd/xiaoniao
```

要件：Go 1.20+、Windows

## ライセンス

MIT
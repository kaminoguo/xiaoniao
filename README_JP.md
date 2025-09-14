# xiaoniao

[中文](README.md) | [English](README_EN.md) | [한국어](README_KR.md)

Windows クリップボード翻訳ツール

## クイックスタート

### 1. APIキーの設定
- メインメニューから「API設定」を選択
- APIキーを入力（OpenAI、Anthropicなど）
- システムが自動的にプロバイダーを識別

### 2. モデル選択
- API設定後、「モデル選択」を選択
- リストから適切なAIモデルを選択

### 3. ホットキー設定（オプション）
- メインメニューから「ホットキー設定」を選択
- 監視切替とプロンプト切替のホットキーを設定

### 4. 使用開始
- Ctrl+C でテキストをコピーして翻訳を起動
- プログラムが自動的にクリップボードを置換
- Ctrl+V で翻訳結果を貼り付け

## ダウンロード

[xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) - Windows 10/11 (64-bit)


## ビルド

```bash
# アイコン付き
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
cd cmd/xiaoniao && goversioninfo -manifest=../../xiaoniao.exe.manifest -icon=../../assets/icon.ico ../../versioninfo.json
cd ../.. && go build -ldflags="-s -w" -o xiaoniao.exe ./cmd/xiaoniao
```

要件：Go 1.20+、Windows

### ビデオチュートリアル

- Bilibili: (近日公開)
- YouTube: (近日公開)

## ライセンス

MIT

## サポート

- Ko-fi: (近日公開)
- WeChat: (近日公開)
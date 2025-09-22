# xiaoniao

[![GitHub Actions Workflow Status](https://github.com/kaminoguo/xiaoniao/actions/workflows/release.yml/badge.svg)](https://github.com/kaminoguo/xiaoniao/blob/main/.github/workflows/release.yml)
[![GitHub last commit](https://img.shields.io/github/last-commit/kaminoguo/xiaoniao)](https://github.com/kaminoguo/xiaoniao/commits/main/)
[![GitHub License](https://img.shields.io/github/license/kaminoguo/xiaoniao)](https://github.com/kaminoguo/xiaoniao/blob/main/LICENSE)

![Demo](assets/demo.gif)

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

- Ctrl+X で切り取りまたは Ctrl+C でコピーして翻訳を起動
- プログラムが自動的にクリップボードを置換
- Ctrl+V で翻訳結果を貼り付け

## ダウンロード

[xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) - Windows 10/11 (64-bit)

## アップデート方法

1. 古い xiaoniao.exe を削除
2. 新しい xiaoniao.exe をダウンロード
3. 設定ファイルは自動保存され、失われません

## ビルド

```bash
# アイコン付き
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
cd cmd/xiaoniao && goversioninfo -manifest=../../xiaoniao.exe.manifest -icon=../../assets/icon.ico ../../versioninfo.json
cd ../.. && go build -ldflags="-s -w" -o xiaoniao.exe ./cmd/xiaoniao
```

要件：Go 1.20+、Windows

### ビデオチュートリアル

- Bilibili: [https://www.bilibili.com/video/BV13zpUzhEeK/](https://www.bilibili.com/video/BV13zpUzhEeK/)
- YouTube: [https://www.youtube.com/watch?v=iPye0tYkBaY](https://www.youtube.com/watch?v=iPye0tYkBaY)

## ライセンス

MIT

## サポート

- Ko-fi: [ko-fi.com/gogogod](https://ko-fi.com/gogogod)
- WeChat: [QRコード](assets/wechat-pay.jpg)

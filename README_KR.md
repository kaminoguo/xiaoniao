# xiaoniao

[中文](README.md) | [English](README_EN.md) | [日本語](README_JP.md)

Windows 클립보드 번역 도구

## 기능

- 클립보드 모니터링 및 자동 번역
- 다국어 UI 지원 (중/영/일/한/불/서/독/러/아랍어)
- 시스템 트레이 통합

## 설치

[xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) 다운로드 후 실행

## 사용법

1. xiaoniao.exe 실행
2. 텍스트 복사 (Ctrl+C)
3. 자동 번역 및 클립보드 교체
4. 붙여넣기 (Ctrl+V)로 번역 결과 얻기

트레이 아이콘 상태: 파란색-모니터링 / 녹색-번역중 / 빨간색-중지

## 설정

```cmd
xiaoniao.exe config
```

OpenAI, Anthropic, Google, DeepSeek 등 API 지원

## 빌드

```bash
# 아이콘 포함
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
cd cmd/xiaoniao && goversioninfo -manifest=../../xiaoniao.exe.manifest -icon=../../assets/icon.ico ../../versioninfo.json
cd ../.. && go build -ldflags="-s -w" -o xiaoniao.exe ./cmd/xiaoniao
```

요구사항: Go 1.20+, Windows

## 라이선스

MIT
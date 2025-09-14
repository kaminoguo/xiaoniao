# xiaoniao

[中文](README.md) | [English](README_EN.md) | [日本語](README_JP.md)

Windows 클립보드 번역 도구

## 빠른 시작

### 1. API 키 설정
- 메인 메뉴에서 "API 설정" 선택
- API 키 입력 (OpenAI, Anthropic 등)
- 시스템이 자동으로 제공업체 식별

### 2. 모델 선택
- API 설정 후 "모델 선택" 선택
- 목록에서 적합한 AI 모델 선택

### 3. 단축키 설정 (선택사항)
- 메인 메뉴에서 "단축키 설정" 선택
- 모니터링 토글 및 프롬프트 전환 단축키 설정

### 4. 사용 시작
- Ctrl+C로 텍스트 복사하여 번역 시작
- 프로그램이 자동으로 클립보드 내용 교체
- Ctrl+V로 번역 결과 붙여넣기

## 다운로드

[xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) - Windows 10/11 (64-bit)


## 빌드

```bash
# 아이콘 포함
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
cd cmd/xiaoniao && goversioninfo -manifest=../../xiaoniao.exe.manifest -icon=../../assets/icon.ico ../../versioninfo.json
cd ../.. && go build -ldflags="-s -w" -o xiaoniao.exe ./cmd/xiaoniao
```

요구사항: Go 1.20+, Windows

### 비디오 튜토리얼

- Bilibili: (공개 예정)
- YouTube: (공개 예정)

## 라이선스

MIT

## 후원

- Ko-fi: (공개 예정)
- WeChat: (공개 예정)
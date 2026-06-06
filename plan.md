# 작업 계획

## 컴필레이션(`COMPILATION`) 태그 채우기 — 방식 #1

### 목표
VocaDB 앨범 정보로 음원 파일 태그를 쓸 때, **`discType`을 보고 컴필레이션 여부를 판정**해
`COMPILATION` 태그를 설정한다.

### 판정 규칙 (방식 #1, 보수적·정확)
```
m.Compilation = "1"  (if album.DiscType == "Compilation")
              = ""   (그 외 전부)
```

- VocaDB 편집자가 분류한 `discType` 값을 그대로 신뢰한다. (`Album.DiscType`은 이미 수신 중)
- `artistString == "Various artists"`는 **쓰지 않는다.** 실측상 `discType: Album`인 다중
  프로듀서 정규 앨범(id=10, 200, 300)이나 `discType: Video`(id=1000)에도 붙어 컴필레이션과
  1:1 대응이 아니기 때문. (방식 #2 후보였으나 오탐 위험으로 제외)

### 작업 항목
- [x] `service.Metadata`에 `Compilation` 필드(`"1"`/`""`) 추가 + `ReadMetadata`/`WriteMetadata` 배선
- [ ] VocaDB `Album` → `service.Metadata` 매핑 함수 추가 (현재 없음). 이 함수에서 위 규칙으로
      `Compilation`을 채운다. 매핑표는 `metadata.md`의 "VocaDB → 태그 매핑" 참고.
- [ ] 매핑 함수에 대한 테스트 추가 (`discType == "Compilation"` → `"1"`, 그 외 → `""`)

### 참고
- `discType` 값 집합(VocaDB AlbumType): `Unknown` `Album` `Single` `EP` `SplitAlbum`
  `Compilation` `Video` `Artbook` `Game` `Fanmade` `Instrumental` `Other`
- 태그 저장 위치: ID3=`TCMP`, MP4=`cpil`, Vorbis(FLAC)=`COMPILATION` — taglib이 `COMPILATION` 키로 정규화.

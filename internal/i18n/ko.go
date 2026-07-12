package i18n

var koMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "터미널 크기를 기다리는 중",
	ReaderStatusTerminalTooSmall:    "터미널 창이 너무 작습니다",
	ReaderStatusLastPage:            "마지막 페이지",
	ReaderStatusFirstPage:           "첫 페이지",
	ReaderStatusRenderError:         "렌더링 오류: %v",
	ReaderStatusMaximumZoom:         "최대 확대",
	ReaderStatusMinimumZoom:         "최소 확대",

	ReaderViewTerminalTooSmall: "comicread: 터미널 창이 너무 작습니다",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "페이지 %d/%d",
	ReaderViewPageRange:        "페이지 %d-%d/%d",
	ReaderViewRendering:        "렌더링 중",
	ReaderViewHelp: `키

← →  이전 / 다음 페이지
↑ ↓  확대된 페이지 스크롤
+ -  확대 / 축소
q    종료

?    도움말 닫기`,

	FilepickerHeader:      "comicread — 챕터 선택\n%s\n\n",
	FilepickerNoEntries:   "  (지원되는 항목 없음)\n",
	FilepickerHelp:        "\n↑/↓ 이동  |  ← 상위 디렉터리  |  → 디렉터리 들어가기  |  enter 파일 열기  |  s 강조 표시된 디렉터리 선택  |  q 종료\n",
	FilepickerWindowTitle: "comicread — 파일 선택",

	FilepickerErrResolveDir: "디렉터리 %q를 확인할 수 없습니다: %w",
	FilepickerErrReadDir:    "디렉터리 %q를 읽을 수 없습니다: %w",
	FilepickerErrRunPicker:  "파일 선택기를 실행하는 중 오류: %w",

	LoadingViewOpening:     "%s 여는 중…",
	LoadingViewWindowTitle: "comicread — 여는 중",

	CLIErrGetWorkingDir:             "작업 디렉터리를 가져올 수 없습니다: %w",
	CLIErrPickFile:                  "파일을 선택할 수 없습니다: %w",
	CLIErrRunTUI:                    "TUI를 실행하는 중 오류: %w",
	CLIErrParseArgs:                 "인수를 분석하는 중 오류: %w",
	CLIErrOpenChapter:               "챕터를 열 수 없습니다: %w",
	CLIErrNoPages:                   "챕터에 읽을 수 있는 이미지 페이지가 없습니다",
	CLIErrInspectInput:              "입력 %q를 확인할 수 없습니다: %w",
	CLIErrUnsupportedFile:           "지원되지 않는 파일 %q: 지원 형식은 CBZ, PDF, EPUB 또는 이미지 디렉터리입니다",
	CLIFlagGraphicsUsage:            "렌더러: auto, ascii, dots, kitty, sixel 또는 iterm2",
	CLIFlagVersionUsage:             "버전을 출력하고 종료",
	CLIFlagEnvUsage:                 "comicread 환경을 출력하고 종료",
	CLIFlagBookViewUsage:            "왼쪽에서 오른쪽 순서로 페이지 쌍 표시",
	CLIFlagRightBookViewUsage:       "오른쪽에서 왼쪽 순서로 페이지 쌍 표시",
	CLIFlagCircleBookViewUsage:      "왼쪽에서 오른쪽 순서로 겹치는 페이지 쌍 표시",
	CLIFlagRightCircleBookViewUsage: "오른쪽에서 왼쪽 순서로 겹치는 페이지 쌍 표시",
	CLIErrMultipleBookViews:         "책 보기 옵션은 하나만 사용할 수 있습니다",
	CLIErrInvalidView:               "지원되지 않는 COMICREAD_VIEW 값 %q (가능한 값: book-view, right-view, circle-view 또는 right-circle-view)",
	CLIHelpHint:                     "도움말은 'comicread --help'를 실행하세요",
	CLIUsage:                        "사용법: comicread [옵션] [파일]",
	CLIUsageFull: `comicread — 미니멀한 터미널 만화 리더

사용법: comicread [옵션] [파일]

옵션:
  --graphics string   렌더러: auto, ascii, dots, kitty, sixel 또는 iterm2 (기본값 "auto")
  --book-view         왼쪽에서 오른쪽 순서로 페이지 쌍 표시
  --right-view        오른쪽에서 왼쪽 순서로 페이지 쌍 표시
  --circle-view       왼쪽에서 오른쪽 순서로 겹치는 페이지 쌍 표시
  --right-circle-view
                      오른쪽에서 왼쪽 순서로 겹치는 페이지 쌍 표시
  --env               comicread 환경을 출력하고 종료
  -v, --version       버전을 출력하고 종료
  -h, --help          이 도움말 표시

파일이나 디렉터리를 지정하지 않으면 현재 디렉터리에서 대화형 파일 선택기가 열립니다.

환경 변수:
  COMICREAD_GRAPHICS  기본 렌더러: auto, ascii, dots, kitty, sixel 또는 iterm2
  COMICREAD_VIEW      기본 보기: book-view, right-view, circle-view 또는 right-circle-view
  COMICREAD_LANG      메시지 언어: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk 또는 ka (기본값 "en")`,
}

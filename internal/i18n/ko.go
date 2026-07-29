package i18n

var koMessages = map[string]string{
	CLIFlagConfigUsage: "사용할 설정 파일", CLIFlagResetConfigUsage: "config.toml을 기본값으로 재설정하고 종료", CLIFlagSetConfigUsage: "config.toml 업데이트: 키=값",
	ReaderStatusWaitingTerminalSize: "터미널 크기를 기다리는 중",
	ReaderStatusTerminalTooSmall:    "터미널 창이 너무 작습니다",
	ReaderStatusLastPage:            "마지막 페이지",
	ReaderStatusFirstPage:           "첫 페이지",
	ReaderStatusRenderError:         "렌더링 오류: %v",
	ReaderStatusMaximumZoom:         "최대 확대",
	ReaderStatusMinimumZoom:         "최소 확대",
	ReaderStatusInvalidPage:         "잘못된 페이지 번호",

	ReaderViewTerminalTooSmall: "comicread: 터미널 창이 너무 작습니다",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "페이지 %d/%d",
	ReaderViewPageRange:        "페이지 %d-%d/%d",
	ReaderViewRendering:        "렌더링 중",
	ReaderViewBookmarks:        "북마크",
	ReaderViewNoBookmarks:      "(북마크 없음)",
	ReaderViewBookmarksHelp:    "위/아래 이동 | enter 열기 | esc 닫기",
	ReaderViewGoToPage:         "페이지로 이동: %s",
	ReaderViewHelp: `키

← →  이전 / 다음 페이지
↑ ↓  확대된 페이지 스크롤
+ -  확대 / 축소
b    북마크 추가 / 제거
v ← → 이전 / 다음 북마크
c v  북마크
g123 enter  페이지로 이동
q    종료

?    도움말 닫기`,

	FilepickerHeader:         "comicread — 챕터 선택\n%s\n\n",
	FilepickerNoEntries:      "  (지원되는 항목 없음)\n",
	FilepickerHelp:           "\n↑/↓ 이동\n← 상위 디렉터리\n→ 디렉터리 들어가기\nenter 파일 열기\ns 강조 표시된 디렉터리 선택\nf 현재 디렉터리를 즐겨찾기에 추가 / 제거\nF 즐겨찾는 디렉터리 추가\nb 즐겨찾는 디렉터리\no 디렉터리로 이동\nq 종료\n",
	FilepickerWindowTitle:    "comicread — 파일 선택",
	FilepickerGoToPrompt:     "\n디렉터리로 이동: %s\n",
	FilepickerFavoritePrompt: "\n즐겨찾는 디렉터리: %s\n",
	FilepickerFavorites:      "즐겨찾는 디렉터리\n\n",
	FilepickerNoFavorites:    "  (설정된 즐겨찾는 디렉터리가 없습니다)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ 이동\nenter 디렉터리로 이동\nd 즐겨찾기 제거\nesc 돌아가기\n",
	FilepickerFavoriteErr:    "  즐겨찾기 저장 오류: %s\n",
	FilepickerGoToErr:        "  오류: %s\n",

	FilepickerErrResolveDir: "디렉터리 %q를 확인할 수 없습니다: %w",
	FilepickerErrReadDir:    "디렉터리 %q를 읽을 수 없습니다: %w",
	FilepickerErrRunPicker:  "파일 선택기를 실행하는 중 오류: %w",
	FilepickerErrEmptyPath:  "경로가 비어 있습니다",
	FilepickerErrNotDir:     "%q은(는) 디렉터리가 아닙니다",

	LoadingViewOpening:     "%s 여는 중…",
	LoadingViewWindowTitle: "comicread — 여는 중",

	CLIErrGetWorkingDir:             "작업 디렉터리를 가져올 수 없습니다: %w",
	CLIErrPickFile:                  "파일을 선택할 수 없습니다: %w",
	CLIErrRunTUI:                    "TUI를 실행하는 중 오류: %w",
	CLIErrParseArgs:                 "인수를 분석하는 중 오류: %w",
	CLIErrOpenChapter:               "챕터를 열 수 없습니다: %w",
	CLIErrOpenJournal:               "저널을 열 수 없습니다: %w",
	CLIErrClearJournal:              "저널을 삭제할 수 없습니다: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal에는 파일 또는 디렉터리가 필요합니다",
	CLIErrNoPages:                   "챕터에 읽을 수 있는 이미지 페이지가 없습니다",
	CLIErrInspectInput:              "입력 %q를 확인할 수 없습니다: %w",
	CLIErrUnsupportedFile:           "지원되지 않는 파일 %q: 지원 형식은 CBZ, PDF, EPUB 또는 이미지 디렉터리입니다",
	CLIFlagGraphicsUsage:            "렌더러: auto, ascii, dots, kitty, sixel 또는 iterm2",
	CLIFlagVersionUsage:             "버전을 출력하고 종료",
	CLIFlagUpdateUsage:              "업데이트를 확인하고 종료",
	CLIFlagEnvUsage:                 "comicread 환경을 출력하고 종료",
	CLIFlagClearJournalUsage:        "파일 또는 디렉터리의 로컬 저널을 삭제하고 종료",
	CLIFlagBookViewUsage:            "왼쪽에서 오른쪽 순서로 페이지 쌍 표시",
	CLIFlagRightBookViewUsage:       "오른쪽에서 왼쪽 순서로 페이지 쌍 표시",
	CLIFlagCircleBookViewUsage:      "왼쪽에서 오른쪽 순서로 겹치는 페이지 쌍 표시",
	CLIFlagRightCircleBookViewUsage: "오른쪽에서 왼쪽 순서로 겹치는 페이지 쌍 표시",
	CLIErrMultipleBookViews:         "책 보기 옵션은 하나만 사용할 수 있습니다",
	CLIErrInvalidView:               "지원되지 않는 COMICREAD_VIEW 값 %q (가능한 값: book-view, right-view, circle-view 또는 right-circle-view)",
	CLIFlagOpenUsage:                "파일 선택기에서 열 디렉터리(기본값: COMICREAD_DIR 또는 현재 디렉터리)",
	CLIErrOpenNotDir:                "디렉터리 열기 %q: 디렉터리가 아닙니다",
	CLIFlagWebUsage:                 "터미널 UI 대신 로컬 웹 리더 시작",
	CLIErrWebArgs:                   "--web는 파일 또는 디렉터리 인수를 허용하지 않습니다",
	WebServerStarted:                "comicread 웹 리더가 %s에서 실행 중입니다 (중지하려면 Ctrl+C를 누르세요)",
	WebErrListen:                    "웹 서버 시작: %w",
	WebErrServe:                     "웹 서버 실행: %w",
	CLIHelpHint:                     "도움말은 'comicread --help'를 실행하세요",
	CLIUsage:                        "사용법: comicread [옵션] [파일]",
	CLIUsageFull: `comicread — 미니멀한 터미널 만화 리더

사용법: comicread [옵션] [파일]

옵션:
  --config string     사용할 설정 파일
  --graphics string   렌더러: auto, ascii, dots, kitty, sixel 또는 iterm2 (기본값 "auto")
  --book-view         왼쪽에서 오른쪽 순서로 페이지 쌍 표시
  --right-view        오른쪽에서 왼쪽 순서로 페이지 쌍 표시
  --circle-view       왼쪽에서 오른쪽 순서로 겹치는 페이지 쌍 표시
  --right-circle-view
                      오른쪽에서 왼쪽 순서로 겹치는 페이지 쌍 표시
  --clear-journal    파일 또는 디렉터리의 로컬 저널을 삭제하고 종료
  --reset-config     config.toml을 기본값으로 재설정하고 종료
  --set-config value config.toml 업데이트: 키=값
  -o, --open string   파일 선택기에서 열 디렉터리(기본값: COMICREAD_DIR 또는 현재 디렉터리)
  --update            업데이트를 확인하고 종료
  --web               터미널 UI 대신 로컬 웹 리더 시작
  -v, --version       버전을 출력하고 종료
  -h, --help          이 도움말 표시

파일이나 디렉터리를 지정하지 않으면 COMICREAD_DIR에서 대화형 파일 선택기가 열립니다
(유효한 디렉터리로 설정된 경우). 그렇지 않으면 현재 디렉터리에서 열립니다.`,
	ReaderViewMetadata: "메타데이터",
}

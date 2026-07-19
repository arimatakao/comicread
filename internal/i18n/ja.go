package i18n

var jaMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "端末サイズを待機中",
	ReaderStatusTerminalTooSmall:    "端末ウィンドウが小さすぎます",
	ReaderStatusLastPage:            "最後のページ",
	ReaderStatusFirstPage:           "最初のページ",
	ReaderStatusRenderError:         "描画エラー: %v",
	ReaderStatusMaximumZoom:         "最大ズーム",
	ReaderStatusMinimumZoom:         "最小ズーム",
	ReaderStatusInvalidPage:         "無効なページ番号です",

	ReaderViewTerminalTooSmall: "comicread: 端末ウィンドウが小さすぎます",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "ページ %d/%d",
	ReaderViewPageRange:        "ページ %d-%d/%d",
	ReaderViewRendering:        "描画中",
	ReaderViewBookmarks:        "ブックマーク",
	ReaderViewNoBookmarks:      "(ブックマークなし)",
	ReaderViewBookmarksHelp:    "↑/↓ 移動 | enter 開く | esc 閉じる",
	ReaderViewGoToPage:         "ページへ移動: %s",
	ReaderViewHelp: `キー

← →  前 / 次のページ
↑ ↓  拡大したページをスクロール
+ -  拡大 / 縮小
b    ブックマークを追加 / 削除
v ← → 前 / 次のブックマーク
c v  ブックマーク
g123 enter  ページへ移動
q    終了

?    ヘルプを閉じる`,

	FilepickerHeader:      "comicread — チャプターを選択\n%s\n\n",
	FilepickerNoEntries:   "  (対応する項目はありません)\n",
	FilepickerHelp:        "\n↑/↓ 移動\n← 親ディレクトリ\n→ ディレクトリに入る\nenter ファイルを開く\ns 選択中のディレクトリを選ぶ\no ディレクトリへ移動\nq 終了\n",
	FilepickerWindowTitle: "comicread — ファイルを選択",
	FilepickerGoToPrompt:  "\nディレクトリへ移動: %s\n",
	FilepickerGoToErr:     "  エラー: %s\n",

	FilepickerErrResolveDir: "ディレクトリ %q を特定できません: %w",
	FilepickerErrReadDir:    "ディレクトリ %q を読み取れません: %w",
	FilepickerErrRunPicker:  "ファイル選択の起動中にエラーが発生しました: %w",
	FilepickerErrEmptyPath:  "パスが空です",
	FilepickerErrNotDir:     "%q はディレクトリではありません",

	LoadingViewOpening:     "%s を開いています…",
	LoadingViewWindowTitle: "comicread — 開いています",

	CLIErrGetWorkingDir:             "作業ディレクトリを取得できません: %w",
	CLIErrPickFile:                  "ファイルを選択できません: %w",
	CLIErrRunTUI:                    "TUI の起動中にエラーが発生しました: %w",
	CLIErrParseArgs:                 "引数の解析中にエラーが発生しました: %w",
	CLIErrOpenChapter:               "チャプターを開けません: %w",
	CLIErrOpenJournal:               "ジャーナルを開けません: %w",
	CLIErrClearJournal:              "ジャーナルを削除できません: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal にはファイルまたはディレクトリが必要です",
	CLIErrNoPages:                   "チャプターに読み取れる画像ページがありません",
	CLIErrInspectInput:              "入力 %q を確認できません: %w",
	CLIErrUnsupportedFile:           "未対応のファイル %q: 対応形式は CBZ、PDF、EPUB、または画像ディレクトリです",
	CLIFlagGraphicsUsage:            "レンダラー: auto、ascii、dots、kitty、sixel、iterm2",
	CLIFlagVersionUsage:             "バージョンを表示して終了",
	CLIFlagUpdateUsage:              "更新を確認して終了",
	CLIFlagEnvUsage:                 "comicread の環境を表示して終了",
	CLIFlagClearJournalUsage:        "ファイルまたはディレクトリのローカルジャーナルを削除して終了",
	CLIFlagBookViewUsage:            "ページの組を左から右へ表示",
	CLIFlagRightBookViewUsage:       "ページの組を右から左へ表示",
	CLIFlagCircleBookViewUsage:      "重なるページの組を左から右へ表示",
	CLIFlagRightCircleBookViewUsage: "重なるページの組を右から左へ表示",
	CLIErrMultipleBookViews:         "本表示オプションは一つだけ使用できます",
	CLIErrInvalidView:               "未対応の COMICREAD_VIEW 値 %q (指定可能: book-view、right-view、circle-view、right-circle-view)",
	CLIFlagOpenUsage:                "ファイル選択で開くディレクトリ（既定値: COMICREAD_DIR または現在のディレクトリ）",
	CLIErrOpenNotDir:                "ディレクトリを開く %q: ディレクトリではありません",
	CLIHelpHint:                     "ヘルプは 'comicread --help' を実行してください",
	CLIUsage:                        "使い方: comicread [オプション] [ファイル]",
	CLIUsageFull: `comicread — ミニマルなターミナル漫画リーダー

使い方: comicread [オプション] [ファイル]

オプション:
  --graphics string   レンダラー: auto、ascii、dots、kitty、sixel、iterm2 (既定値 "auto")
  --book-view         ページの組を左から右へ表示
  --right-view        ページの組を右から左へ表示
  --circle-view       重なるページの組を左から右へ表示
  --right-circle-view
                      重なるページの組を右から左へ表示
  --clear-journal    ファイルまたはディレクトリのローカルジャーナルを削除して終了
  -o, --open string   ファイル選択で開くディレクトリ（既定値: COMICREAD_DIR または現在のディレクトリ）
  --env               comicread の環境を表示して終了
  --update            更新を確認して終了
  -v, --version       バージョンを表示して終了
  -h, --help          このヘルプを表示

ファイルまたはディレクトリが指定されない場合、COMICREAD_DIR で対話型ファイル選択が開きます
（有効なディレクトリに設定されている場合）。それ以外は現在のディレクトリで開きます。

環境変数:
  COMICREAD_GRAPHICS  既定のレンダラー: auto、ascii、dots、kitty、sixel、iterm2
  COMICREAD_PRERENDERED_NEXT      事前描画する次のページ数（既定値 1）
  COMICREAD_PRERENDERED_PREVIOUS  事前描画する前のページ数（既定値 1）
  COMICREAD_VIEW      既定の表示: book-view、right-view、circle-view、right-circle-view
  COMICREAD_LANG      メッセージの言語: https://github.com/arimatakao/comicread#environment-variables (既定値 "en")
  COMICREAD_DIR       パス未指定時にファイル選択で使う既定のディレクトリ`,
}

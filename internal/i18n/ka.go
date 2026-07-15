package i18n

var kaMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "ტერმინალის ზომის მოლოდინში",
	ReaderStatusTerminalTooSmall:    "ტერმინალის ფანჯარა ძალიან პატარაა",
	ReaderStatusLastPage:            "ბოლო გვერდი",
	ReaderStatusFirstPage:           "პირველი გვერდი",
	ReaderStatusRenderError:         "რენდერის შეცდომა: %v",
	ReaderStatusMaximumZoom:         "მაქსიმალური მასშტაბი",
	ReaderStatusMinimumZoom:         "მინიმალური მასშტაბი",

	ReaderViewTerminalTooSmall: "comicread: ტერმინალის ფანჯარა ძალიან პატარაა",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "გვერდები %d/%d",
	ReaderViewPageRange:        "გვერდები %d-%d/%d",
	ReaderViewRendering:        "რენდერი",
	ReaderViewBookmarks:        "სანიშნეები",
	ReaderViewNoBookmarks:      "(სანიშნეები არ არის)",
	ReaderViewBookmarksHelp:    "↑/↓ გადაადგილება | enter გახსნა | esc დახურვა",
	ReaderViewHelp: `კლავიშები

← →  წინა / შემდეგი გვერდი
↑ ↓  გადიდებული გვერდის გადახვევა
+ -  გადიდება / დაპატარავება
b    სანიშნის დამატება / წაშლა
v ← → წინა / შემდეგი სანიშნე
c v  სანიშნეები
q    გასვლა

?    დახმარების დახურვა`,

	FilepickerHeader:      "comicread — აირჩიეთ თავი\n%s\n\n",
	FilepickerNoEntries:   "  (მხარდაჭერილი ელემენტები არ არის)\n",
	FilepickerHelp:        "\n↑/↓ გადაადგილება  |  ← მშობელი საქაღალდე  |  → საქაღალდის გახსნა  |  enter ფაილის გახსნა  |  s მონიშნული საქაღალდის არჩევა  |  q გასვლა\n",
	FilepickerWindowTitle: "comicread — ფაილის არჩევა",

	FilepickerErrResolveDir: "საქაღალდის %q დადგენა ვერ მოხერხდა: %w",
	FilepickerErrReadDir:    "საქაღალდის %q წაკითხვა ვერ მოხერხდა: %w",
	FilepickerErrRunPicker:  "ფაილის არჩევის შეცდომა: %w",

	LoadingViewOpening:     "%s იხსნება…",
	LoadingViewWindowTitle: "comicread — იხსნება",

	CLIErrGetWorkingDir:             "სამუშაო საქაღალდის მიღება ვერ მოხერხდა: %w",
	CLIErrPickFile:                  "ფაილის არჩევა ვერ მოხერხდა: %w",
	CLIErrRunTUI:                    "TUI-ის გაშვების შეცდომა: %w",
	CLIErrParseArgs:                 "არგუმენტების დამუშავების შეცდომა: %w",
	CLIErrOpenChapter:               "თავის გახსნა ვერ მოხერხდა: %w",
	CLIErrOpenJournal:               "ჟურნალის გახსნა ვერ მოხერხდა: %w",
	CLIErrClearJournal:              "ჟურნალის წაშლა ვერ მოხერხდა: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal საჭიროებს ფაილს ან საქაღალდეს",
	CLIErrNoPages:                   "თავი არ შეიცავს წასაკითხად გამოსადეგ გვერდებს",
	CLIErrInspectInput:              "შეყვანილი მონაცემების %q შემოწმება ვერ მოხერხდა: %w",
	CLIErrUnsupportedFile:           "მხარდაუჭერელი ფაილი %q: მხარდაჭერილი ფორმატებია — CBZ, PDF, EPUB ან სურათების საქაღალდე",
	CLIFlagGraphicsUsage:            "რენდერერი: auto, ascii, dots, kitty, sixel ან iterm2",
	CLIFlagVersionUsage:             "ვერსიის გამოტანა და დასრულება",
	CLIFlagUpdateUsage:              "განახლებების შემოწმება და დასრულება",
	CLIFlagEnvUsage:                 "comicread-ის გარემოს გამოტანა და დასრულება",
	CLIFlagClearJournalUsage:        "ფაილის ან საქაღალდის ადგილობრივი ჟურნალის წაშლა და დასრულება",
	CLIFlagBookViewUsage:            "გვერდების წყვილებად ჩვენება მარცხნიდან მარჯვნივ",
	CLIFlagRightBookViewUsage:       "გვერდების წყვილებად ჩვენება მარჯვნიდან მარცხნივ",
	CLIFlagCircleBookViewUsage:      "გვერდების გადამფარავი წყვილებად ჩვენება მარცხნიდან მარჯვნივ",
	CLIFlagRightCircleBookViewUsage: "გვერდების გადამფარავი წყვილებად ჩვენება მარჯვნიდან მარცხნივ",
	CLIErrMultipleBookViews:         "მხოლოდ ერთი წიგნის ხედვის რეჟიმის მითითებაა შესაძლებელი",
	CLIErrInvalidView:               "COMICREAD_VIEW-ის მხარდაუჭერელი მნიშვნელობა %q (შესაძლებელია: book-view, right-view, circle-view ან right-circle-view)",
	CLIHelpHint:                     "დახმარებისთვის გაუშვით 'comicread --help'",
	CLIUsage:                        "გამოყენება: comicread [პარამეტრები] [ფაილი]",
	CLIUsageFull: `comicread - მინიმალისტური მანგის წამკითხველი ტერმინალისთვის

გამოყენება: comicread [პარამეტრები] [ფაილი]

პარამეტრები:
  --graphics string   რენდერერი: auto, ascii, dots, kitty, sixel ან iterm2 (ნაგულისხმევი "auto")
  --book-view         გვერდების წყვილებად ჩვენება მარცხნიდან მარჯვნივ
  --right-view        გვერდების წყვილებად ჩვენება მარჯვნიდან მარცხნივ
  --circle-view       გვერდების გადამფარავი წყვილებად ჩვენება მარცხნიდან მარჯვნივ
  --right-circle-view
                      გვერდების გადამფარავი წყვილებად ჩვენება მარჯვნიდან მარცხნივ
  --clear-journal    ფაილის ან საქაღალდის ადგილობრივი ჟურნალის წაშლა და დასრულება
  --env               comicread-ის გარემოს გამოტანა და დასრულება
  --update            განახლებების შემოწმება და დასრულება
  -v, --version       ვერსიის გამოტანა და დასრულება
  -h, --help          ამ დახმარების ჩვენება

თუ ფაილი ან საქაღალდე მითითებული არ არის, გაიხსნება ინტერაქტიული ფაილის არჩევა მიმდინარე საქაღალდეში.

გარემოს ცვლადები:
  COMICREAD_GRAPHICS  ნაგულისხმევი რენდერერი: auto, ascii, dots, kitty, sixel ან iterm2
  COMICREAD_VIEW      ნაგულისხმევი რეჟიმი: book-view, right-view, circle-view ან right-circle-view
  COMICREAD_LANG   შეტყობინებების ენა: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk ან ka (ნაგულისხმევი "en")`,
}

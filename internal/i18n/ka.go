package i18n

var kaMessages = map[string]string{
	CLIFlagConfigUsage: "გამოსაყენებელი კონფიგურაციის ფაილი", CLIFlagResetConfigUsage: "config.toml ნაგულისხმევზე დააბრუნე და დაასრულე", CLIFlagSetConfigUsage: "config.toml განაახლე: გასაღები=მნიშვნელობა",
	ReaderStatusWaitingTerminalSize: "ტერმინალის ზომის მოლოდინში",
	ReaderStatusTerminalTooSmall:    "ტერმინალის ფანჯარა ძალიან პატარაა",
	ReaderStatusLastPage:            "ბოლო გვერდი",
	ReaderStatusFirstPage:           "პირველი გვერდი",
	ReaderStatusRenderError:         "რენდერის შეცდომა: %v",
	ReaderStatusMaximumZoom:         "მაქსიმალური მასშტაბი",
	ReaderStatusMinimumZoom:         "მინიმალური მასშტაბი",
	ReaderStatusInvalidPage:         "გვერდის არასწორი ნომერი",

	ReaderViewTerminalTooSmall: "comicread: ტერმინალის ფანჯარა ძალიან პატარაა",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "გვერდები %d/%d",
	ReaderViewPageRange:        "გვერდები %d-%d/%d",
	ReaderViewRendering:        "რენდერი",
	ReaderViewBookmarks:        "სანიშნეები",
	ReaderViewNoBookmarks:      "(სანიშნეები არ არის)",
	ReaderViewBookmarksHelp:    "↑/↓ გადაადგილება | enter გახსნა | esc დახურვა",
	ReaderViewGoToPage:         "გვერდზე გადასვლა: %s",
	ReaderViewHelp: `კლავიშები

← →  წინა / შემდეგი გვერდი
↑ ↓  გადიდებული გვერდის გადახვევა
+ -  გადიდება / დაპატარავება
b    სანიშნის დამატება / წაშლა
v ← → წინა / შემდეგი სანიშნე
c v  სანიშნეები
g123 enter  გვერდზე გადასვლა
q    გასვლა

?    დახმარების დახურვა`,

	FilepickerHeader:         "comicread — აირჩიეთ თავი\n%s\n\n",
	FilepickerNoEntries:      "  (მხარდაჭერილი ელემენტები არ არის)\n",
	FilepickerHelp:           "\n↑/↓ გადაადგილება\n← მშობელი საქაღალდე\n→ საქაღალდის გახსნა\nenter ფაილის გახსნა\ns მონიშნული საქაღალდის არჩევა\nf მიმდინარე საქაღალდის რჩეულებში დამატება / ამოღება\nF რჩეული საქაღალდის დამატება\nb რჩეული საქაღალდეები\no საქაღალდეზე გადასვლა\nq გასვლა\n",
	FilepickerWindowTitle:    "comicread — ფაილის არჩევა",
	FilepickerGoToPrompt:     "\nგადასვლა საქაღალდეზე: %s\n",
	FilepickerFavoritePrompt: "\nრჩეული საქაღალდე: %s\n",
	FilepickerFavorites:      "რჩეული საქაღალდეები\n\n",
	FilepickerNoFavorites:    "  (რჩეული საქაღალდეები არ არის კონფიგურირებული)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ გადაადგილება\nenter საქაღალდეზე გადასვლა\nd რჩეულის ამოღება\nesc დაბრუნება\n",
	FilepickerFavoriteErr:    "  რჩეულების შენახვის შეცდომა: %s\n",
	FilepickerGoToErr:        "  შეცდომა: %s\n",

	FilepickerErrResolveDir: "საქაღალდის %q დადგენა ვერ მოხერხდა: %w",
	FilepickerErrReadDir:    "საქაღალდის %q წაკითხვა ვერ მოხერხდა: %w",
	FilepickerErrRunPicker:  "ფაილის არჩევის შეცდომა: %w",
	FilepickerErrEmptyPath:  "ბილიკი ცარიელია",
	FilepickerErrNotDir:     "%q არ არის საქაღალდე",

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
	CLIFlagOpenUsage:                "საქაღალდე ფაილის ამრჩევში გასახსნელად (ნაგულისხმევი: COMICREAD_DIR ან მიმდინარე საქაღალდე)",
	CLIErrOpenNotDir:                "საქაღალდის გახსნა %q: არ არის საქაღალდე",
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
  -o, --open string   საქაღალდე ფაილის ამრჩევში გასახსნელად (ნაგულისხმევი: COMICREAD_DIR ან მიმდინარე საქაღალდე)
  --env               comicread-ის გარემოს გამოტანა და დასრულება
  --update            განახლებების შემოწმება და დასრულება
  -v, --version       ვერსიის გამოტანა და დასრულება
  -h, --help          ამ დახმარების ჩვენება

თუ ფაილი ან საქაღალდე მითითებული არ არის, გაიხსნება ინტერაქტიული ფაილის არჩევა COMICREAD_DIR-ში
(თუ იგი სწორ საქაღალდეზეა დაყენებული) ან მიმდინარე საქაღალდეში.

გარემოს ცვლადები:
  COMICREAD_GRAPHICS  ნაგულისხმევი რენდერერი: auto, ascii, dots, kitty, sixel ან iterm2
  COMICREAD_PRERENDERED_NEXT      შემდეგი გვერდები წინასწარი რენდერისთვის (ნაგულისხმევი 1)
  COMICREAD_PRERENDERED_PREVIOUS  წინა გვერდები წინასწარი რენდერისთვის (ნაგულისხმევი 1)
  COMICREAD_VIEW      ნაგულისხმევი რეჟიმი: book-view, right-view, circle-view ან right-circle-view
  COMICREAD_LANG   შეტყობინებების ენა: https://github.com/arimatakao/comicread#environment-variables (ნაგულისხმევი "en")
  COMICREAD_DIR    ფაილის ამრჩევის ნაგულისხმევი საქაღალდე, როცა გზა მითითებული არაა`,
}

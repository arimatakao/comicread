package i18n

var trMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "terminal boyutu bekleniyor",
	ReaderStatusTerminalTooSmall:    "terminal penceresi çok küçük",
	ReaderStatusLastPage:            "son sayfa",
	ReaderStatusFirstPage:           "ilk sayfa",
	ReaderStatusRenderError:         "oluşturma hatası: %v",
	ReaderStatusMaximumZoom:         "en yüksek yakınlaştırma",
	ReaderStatusMinimumZoom:         "en düşük yakınlaştırma",

	ReaderViewTerminalTooSmall: "comicread: terminal penceresi çok küçük",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "sayfalar %d/%d",
	ReaderViewPageRange:        "sayfalar %d-%d/%d",
	ReaderViewRendering:        "oluşturuluyor",
	ReaderViewBookmarks:        "Yer imleri",
	ReaderViewNoBookmarks:      "(yer imi yok)",
	ReaderViewBookmarksHelp:    "yukarı/aşağı hareket | enter aç | esc kapat",
	ReaderViewHelp: `Tuşlar

← →  önceki / sonraki sayfa
↑ ↓  yakınlaştırılmış sayfayı kaydır
+ -  yakınlaştır / uzaklaştır
b    yer imi ekle / kaldır
v ← → önceki / sonraki yer imi
c v  yer imleri
q    çıkış

?    yardımı kapat`,

	FilepickerHeader:      "comicread — bölüm seçin\n%s\n\n",
	FilepickerNoEntries:   "  (desteklenen girdi yok)\n",
	FilepickerHelp:        "\n↑/↓ hareket  |  ← üst dizin  |  → dizine gir  |  enter dosyayı aç  |  s vurgulanan dizini seç  |  q çıkış\n",
	FilepickerWindowTitle: "comicread — dosya seçin",

	FilepickerErrResolveDir: "%q dizini çözümlenemedi: %w",
	FilepickerErrReadDir:    "%q dizini okunamadı: %w",
	FilepickerErrRunPicker:  "dosya seçici çalıştırılırken hata oluştu: %w",

	LoadingViewOpening:     "%s açılıyor…",
	LoadingViewWindowTitle: "comicread — açılıyor",

	CLIErrGetWorkingDir:             "çalışma dizini alınamadı: %w",
	CLIErrPickFile:                  "dosya seçilemedi: %w",
	CLIErrRunTUI:                    "TUI çalıştırılırken hata oluştu: %w",
	CLIErrParseArgs:                 "bağımsız değişkenler ayrıştırılırken hata oluştu: %w",
	CLIErrOpenChapter:               "bölüm açılamadı: %w",
	CLIErrOpenJournal:               "günlük açılamadı: %w",
	CLIErrClearJournal:              "günlük silinemedi: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal bir dosya veya dizin gerektirir",
	CLIErrNoPages:                   "bölüm okunabilir görüntü sayfası içermiyor",
	CLIErrInspectInput:              "%q girdisi incelenemedi: %w",
	CLIErrUnsupportedFile:           "desteklenmeyen dosya %q: desteklenen biçimler CBZ, PDF, EPUB veya görüntü dizinidir",
	CLIFlagGraphicsUsage:            "oluşturucu: auto, ascii, dots, kitty, sixel veya iterm2",
	CLIFlagVersionUsage:             "sürümü göster ve çık",
	CLIFlagUpdateUsage:              "güncellemeleri denetle ve çık",
	CLIFlagEnvUsage:                 "comicread ortamını göster ve çık",
	CLIFlagClearJournalUsage:        "dosya veya dizin için yerel günlüğü sil ve çık",
	CLIFlagBookViewUsage:            "sayfa çiftlerini soldan sağa göster",
	CLIFlagRightBookViewUsage:       "sayfa çiftlerini sağdan sola göster",
	CLIFlagCircleBookViewUsage:      "örtüşen sayfa çiftlerini soldan sağa göster",
	CLIFlagRightCircleBookViewUsage: "örtüşen sayfa çiftlerini sağdan sola göster",
	CLIErrMultipleBookViews:         "yalnızca bir kitap görünümü seçeneği kullanılabilir",
	CLIErrInvalidView:               "desteklenmeyen COMICREAD_VIEW değeri %q (beklenen: book-view, right-view, circle-view veya right-circle-view)",
	CLIHelpHint:                     "yardım için 'comicread --help' çalıştırın",
	CLIUsage:                        "kullanım: comicread [seçenekler] [dosya]",
	CLIUsageFull: `comicread — terminal için minimal bir manga okuyucusu

kullanım: comicread [seçenekler] [dosya]

seçenekler:
  --graphics string   oluşturucu: auto, ascii, dots, kitty, sixel veya iterm2 (varsayılan "auto")
  --book-view         sayfa çiftlerini soldan sağa göster
  --right-view        sayfa çiftlerini sağdan sola göster
  --circle-view       örtüşen sayfa çiftlerini soldan sağa göster
  --right-circle-view
                      örtüşen sayfa çiftlerini sağdan sola göster
  --clear-journal    dosya veya dizin için yerel günlüğü sil ve çık
  --env               comicread ortamını göster ve çık
  --update            güncellemeleri denetle ve çık
  -v, --version       sürümü göster ve çık
  -h, --help          bu yardımı göster

Dosya veya dizin verilmezse, geçerli dizinde etkileşimli dosya seçici açılır.

ortam:
  COMICREAD_GRAPHICS  varsayılan oluşturucu: auto, ascii, dots, kitty, sixel veya iterm2
  COMICREAD_VIEW      varsayılan görünüm: book-view, right-view, circle-view veya right-circle-view
  COMICREAD_LANG      ileti dili: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk veya ka (varsayılan "en")`,
}

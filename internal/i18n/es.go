package i18n

var esMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "esperando el tamaño del terminal",
	ReaderStatusTerminalTooSmall:    "la ventana del terminal es demasiado pequeña",
	ReaderStatusLastPage:            "última página",
	ReaderStatusFirstPage:           "primera página",
	ReaderStatusRenderError:         "error de renderizado: %v",
	ReaderStatusMaximumZoom:         "zoom máximo",
	ReaderStatusMinimumZoom:         "zoom mínimo",
	ReaderStatusInvalidPage:         "número de página no válido",

	ReaderViewTerminalTooSmall: "comicread: la ventana del terminal es demasiado pequeña",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "páginas %d/%d",
	ReaderViewPageRange:        "páginas %d-%d/%d",
	ReaderViewRendering:        "renderizando",
	ReaderViewBookmarks:        "Marcadores",
	ReaderViewNoBookmarks:      "(sin marcadores)",
	ReaderViewBookmarksHelp:    "arriba/abajo mover | enter abrir | esc cerrar",
	ReaderViewGoToPage:         "Ir a la página: %s",
	ReaderViewHelp: `Teclas

← →  página anterior / siguiente
↑ ↓  desplazar una página ampliada
+ -  ampliar / reducir
b    añadir / quitar marcador
v ← → marcador anterior / siguiente
c v  marcadores
g123 enter  ir a la página
q    salir

?    cerrar ayuda`,

	FilepickerHeader:      "comicread — seleccionar un capítulo\n%s\n\n",
	FilepickerNoEntries:   "  (no hay entradas compatibles)\n",
	FilepickerHelp:        "\n↑/↓ mover\n← directorio superior\n→ entrar al directorio\nenter abrir archivo\ns seleccionar el directorio resaltado\no ir a un directorio\nq salir\n",
	FilepickerWindowTitle: "comicread — elegir un archivo",
	FilepickerGoToPrompt:  "\nIr al directorio: %s\n",
	FilepickerGoToErr:     "  error: %s\n",

	FilepickerErrResolveDir: "no se puede resolver el directorio %q: %w",
	FilepickerErrReadDir:    "no se puede leer el directorio %q: %w",
	FilepickerErrRunPicker:  "error al ejecutar el selector de archivos: %w",
	FilepickerErrEmptyPath:  "la ruta está vacía",
	FilepickerErrNotDir:     "%q no es un directorio",

	LoadingViewOpening:     "abriendo %s…",
	LoadingViewWindowTitle: "comicread — abriendo",

	CLIErrGetWorkingDir:             "no se puede obtener el directorio de trabajo: %w",
	CLIErrPickFile:                  "no se puede elegir el archivo: %w",
	CLIErrRunTUI:                    "error al ejecutar la TUI: %w",
	CLIErrParseArgs:                 "error al analizar los argumentos: %w",
	CLIErrOpenChapter:               "no se puede abrir el capítulo: %w",
	CLIErrOpenJournal:               "no se puede abrir el registro: %w",
	CLIErrClearJournal:              "no se puede borrar el registro: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal requiere un archivo o directorio",
	CLIErrNoPages:                   "el capítulo no contiene páginas de imagen legibles",
	CLIErrInspectInput:              "no se puede inspeccionar la entrada %q: %w",
	CLIErrUnsupportedFile:           "archivo no compatible %q: los formatos compatibles son CBZ, PDF, EPUB o un directorio de imágenes",
	CLIFlagGraphicsUsage:            "renderizador: auto, ascii, dots, kitty, sixel o iterm2",
	CLIFlagVersionUsage:             "mostrar la versión y salir",
	CLIFlagUpdateUsage:              "buscar actualizaciones y salir",
	CLIFlagEnvUsage:                 "mostrar el entorno de comicread y salir",
	CLIFlagClearJournalUsage:        "eliminar el registro local de un archivo o directorio y salir",
	CLIFlagBookViewUsage:            "mostrar pares de páginas de izquierda a derecha",
	CLIFlagRightBookViewUsage:       "mostrar pares de páginas de derecha a izquierda",
	CLIFlagCircleBookViewUsage:      "mostrar pares de páginas superpuestas de izquierda a derecha",
	CLIFlagRightCircleBookViewUsage: "mostrar pares de páginas superpuestas de derecha a izquierda",
	CLIErrMultipleBookViews:         "solo se puede usar una opción de vista de libro",
	CLIErrInvalidView:               "valor de COMICREAD_VIEW no compatible %q (se espera: book-view, right-view, circle-view o right-circle-view)",
	CLIFlagOpenUsage:                "directorio para abrir en el selector de archivos (predeterminado: COMICREAD_DIR o el directorio actual)",
	CLIErrOpenNotDir:                "abrir directorio %q: no es un directorio",
	CLIHelpHint:                     "ejecute 'comicread --help' para ver la ayuda",
	CLIUsage:                        "uso: comicread [opciones] [archivo]",
	CLIUsageFull: `comicread — un lector de manga minimalista para la terminal

uso: comicread [opciones] [archivo]

opciones:
  --graphics string   renderizador: auto, ascii, dots, kitty, sixel o iterm2 (predeterminado "auto")
  --book-view         mostrar pares de páginas de izquierda a derecha
  --right-view        mostrar pares de páginas de derecha a izquierda
  --circle-view       mostrar pares de páginas superpuestas de izquierda a derecha
  --right-circle-view
                      mostrar pares de páginas superpuestas de derecha a izquierda
  --clear-journal    eliminar el registro local de un archivo o directorio y salir
  -o, --open string   directorio para abrir en el selector de archivos (predeterminado: COMICREAD_DIR o el directorio actual)
  --env               mostrar el entorno de comicread y salir
  --update            buscar actualizaciones y salir
  -v, --version       mostrar la versión y salir
  -h, --help          mostrar esta ayuda

Si no se proporciona ningún archivo o directorio, se abre un selector de archivos interactivo en COMICREAD_DIR
(si está configurado como un directorio válido) o, de lo contrario, en el directorio actual.

entorno:
  COMICREAD_GRAPHICS  renderizador predeterminado: auto, ascii, dots, kitty, sixel o iterm2
  COMICREAD_PRERENDERED_NEXT      páginas siguientes para prerenderizar (predeterminado 1)
  COMICREAD_PRERENDERED_PREVIOUS  páginas anteriores para prerenderizar (predeterminado 1)
  COMICREAD_VIEW      vista predeterminada: book-view, right-view, circle-view o right-circle-view
  COMICREAD_LANG      idioma de los mensajes: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el, tr, kk o ka (predeterminado "en")
  COMICREAD_DIR       directorio predeterminado para el selector de archivos cuando no se indica una ruta`,
}

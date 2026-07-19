#!/usr/bin/env bash

set -euo pipefail

REPO="arimatakao/comicread"
BIN_NAME="comicread"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
VERSION_INPUT="latest"
AUTO_YES="false"
TMP_DIR=""
REINSTALL_CONFIRMED="false"
LOCALE_BASE_URL="https://raw.githubusercontent.com/${REPO}/main/installer/locales"
LOCALE_FILE=""
CONFIGURED_ORDER=(language graphics view prerender.next prerender.previous directory)
declare -A CONFIGURED_VALUES=()

detect_locale() {
  local locale
  locale="${COMICREAD_LANG:-${LC_ALL:-${LC_MESSAGES:-${LANG:-en}}}}"
  locale="${locale%%[_@.]*}"
  printf '%s\n' "$(printf '%s' "$locale" | tr '[:upper:]' '[:lower:]')"
}

load_locale() {
  local language="$1" script_dir local_file fallback_file downloaded_file
  script_dir="$(CDPATH= cd -- "$(dirname -- "${BASH_SOURCE[0]}")" 2>/dev/null && pwd || true)"
  local_file="${script_dir}/installer/locales/${language}.properties"
  if [ -r "$local_file" ]; then
    LOCALE_FILE="$local_file"
    return
  fi
  fallback_file="${script_dir}/installer/locales/en.properties"
  if [ -r "$fallback_file" ]; then
    LOCALE_FILE="$fallback_file"
    return
  fi

  command -v curl >/dev/null 2>&1 || return
  downloaded_file="${TMP_DIR}/installer-${language}.properties"
  if curl -fsSL "${LOCALE_BASE_URL}/${language}.properties" -o "$downloaded_file" && [ -s "$downloaded_file" ]; then
    LOCALE_FILE="$downloaded_file"
    return
  fi

  downloaded_file="${TMP_DIR}/installer-en.properties"
  if curl -fsSL "${LOCALE_BASE_URL}/en.properties" -o "$downloaded_file" && [ -s "$downloaded_file" ]; then
    LOCALE_FILE="$downloaded_file"
  fi
}

t() {
  local key="$1" message
  shift
  if [ -n "$LOCALE_FILE" ]; then
    message="$(awk -v key="$key" 'index($0, key "=") == 1 { sub(/^[^=]*=/, ""); print; exit }' "$LOCALE_FILE")"
  fi
  [ -n "${message:-}" ] || message="$key"

  if [ "$#" -eq 0 ]; then
    printf '%s\n' "$message"
  else
    printf "${message}\\n" "$@"
  fi
}

cleanup() {
  [ -z "${TMP_DIR:-}" ] || rm -rf "$TMP_DIR"
}

usage() {
  t usage 'bash install.sh [-y|--yes] [version]'
  printf '\n'
  t usage.description
  printf '\n'
  t usage.examples
  printf '%s\n' '  bash install.sh' '  bash install.sh v1.2.3' '  bash install.sh --yes'
}

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    t error.required_command "$1" >&2
    exit 1
  fi
}

parse_args() {
  local positional=()

  while [ "$#" -gt 0 ]; do
    case "$1" in
      -y|--yes)
        AUTO_YES="true"
        ;;
      -h|--help)
        usage
        exit 0
        ;;
      -*)
        t error.unknown_option "$1" >&2
        usage >&2
        exit 1
        ;;
      *) positional+=("$1") ;;
    esac
    shift
  done

  if [ "${#positional[@]}" -gt 1 ]; then
    t error.too_many_versions >&2
    usage >&2
    exit 1
  fi

  if [ "${#positional[@]}" -eq 1 ]; then
    VERSION_INPUT="${positional[0]}"
  fi
}

read_answer() {
  ANSWER=""
  if [ -t 0 ]; then
    read -r -p "$1" ANSWER
  elif [ -r /dev/tty ]; then
    read -r -p "$1" ANSWER < /dev/tty
  else
    t error.interactive_input >&2
    exit 1
  fi
}

confirm_install() {
  [ "$AUTO_YES" = "true" ] && return 0
  read_answer "$1 [y/N]: "

  case "$ANSWER" in
    y|Y|yes|YES) ;;
    *)
      t cancelled
      exit 0
      ;;
  esac
}

ask_yes_no() {
  read_answer "$1 [y/N]: "
  case "$ANSWER" in
    y|Y|yes|YES) return 0 ;;
    *) return 1 ;;
  esac
}

normalize_os() {
  case "$(uname -s)" in
    Linux) printf '%s\n' 'linux' ;;
    Darwin) printf '%s\n' 'darwin' ;;
    *)
      t error.unsupported_os "$(uname -s)" >&2
      exit 1
      ;;
  esac
}

normalize_arch() {
  case "$(uname -m)" in
    x86_64|amd64) printf '%s\n' 'amd64' ;;
    aarch64|arm64) printf '%s\n' 'arm64' ;;
    i386|i686) printf '%s\n' '386' ;;
    *)
      t error.unsupported_arch "$(uname -m)" >&2
      exit 1
      ;;
  esac
}

resolve_version() {
  if [ "$VERSION_INPUT" = "latest" ]; then
    local latest_url
    latest_url="$(curl -fsSLI -o /dev/null -w '%{url_effective}' "https://github.com/${REPO}/releases/latest")"
    basename "$latest_url"
  elif [[ "$VERSION_INPUT" == v* ]]; then
    printf '%s\n' "$VERSION_INPUT"
  else
    printf 'v%s\n' "$VERSION_INPUT"
  fi
}

normalize_version() {
  local value="${1#v}"
  value="${value%%[-+]*}"
  printf '%s\n' "$value"
}

version_to_sort_key() {
  local version major minor patch extra
  version="$(normalize_version "$1")"
  IFS='.' read -r major minor patch extra <<EOF
$version
EOF
  major="${major:-0}"
  minor="${minor:-0}"
  patch="${patch:-0}"
  case "${major}.${minor}.${patch}" in
    *[!0-9.]*|*..*) return 1 ;;
  esac
  printf '%010d%010d%010d\n' "$major" "$minor" "$patch"
}

is_version_newer() {
  local target_key current_key
  target_key="$(version_to_sort_key "$1")" || return 1
  current_key="$(version_to_sort_key "$2")" || return 1
  [ "$target_key" \> "$current_key" ]
}

get_installed_version() {
  local candidate version_output match

  for candidate in "${INSTALL_DIR}/${BIN_NAME}" "$(command -v "$BIN_NAME" 2>/dev/null || true)"; do
    [ -n "$candidate" ] && [ -x "$candidate" ] || continue
    version_output="$("$candidate" --version 2>/dev/null || true)"
    match="$(printf '%s\n' "$version_output" | grep -Eo 'v?[0-9]+(\.[0-9]+){1,3}([-.+][0-9A-Za-z.-]+)?' | head -n1)"
    if [ -n "$match" ]; then
      printf '%s\n' "$match"
      return 0
    fi
  done

  return 1
}

confirm_upgrade_if_needed() {
  local target_version="$1" installed_version=""
  installed_version="$(get_installed_version || true)"
  [ -n "$installed_version" ] || return 0

  if is_version_newer "$target_version" "$installed_version"; then
    confirm_install "$(t confirm.update "$installed_version" "$target_version")"
    REINSTALL_CONFIRMED="true"
    return 0
  fi

  if [ "$(normalize_version "$target_version")" = "$(normalize_version "$installed_version")" ]; then
    confirm_install "$(t confirm.reinstall "$installed_version" "$target_version")"
    REINSTALL_CONFIRMED="true"
  fi
}

path_contains_install_dir() {
  case ":${PATH:-}:" in
    *":${INSTALL_DIR}:"*) return 0 ;;
    *) return 1 ;;
  esac
}

shell_config_file() {
  local shell_name
  shell_name="$(basename "${SHELL:-}")"
  case "$shell_name" in
    bash) printf '%s\n' "${HOME}/.bashrc" ;;
    zsh) printf '%s\n' "${HOME}/.zshrc" ;;
    *) printf '%s\n' "${HOME}/.profile" ;;
  esac
}

append_shell_line() {
  local target_file="$1" line="$2"
  mkdir -p "$(dirname "$target_file")"
  touch "$target_file"
  grep -Fqx "$line" "$target_file" || printf '\n%s\n' "$line" >> "$target_file"
}

ensure_path_setup() {
  path_contains_install_dir && return 0

  local target_file line
  target_file="$(shell_config_file)"
  line="case \":\$PATH:\" in *\":${INSTALL_DIR}:\"*) ;; *) export PATH=\"${INSTALL_DIR}:\$PATH\" ;; esac"
  append_shell_line "$target_file" "$line"
}

config_file_path() {
  case "$(uname -s)" in
    Darwin) printf '%s\n' "${HOME}/Library/Application Support/comicread/config.toml" ;;
    *) printf '%s\n' "${XDG_CONFIG_HOME:-${HOME}/.config}/comicread/config.toml" ;;
  esac
}

apply_config_option() {
  local bin="$1" assignment="$2" output
  if output="$("$bin" --set-config "$assignment" 2>&1)"; then
    return 0
  fi
  printf '%s\n' "$output" >&2
  return 1
}

configure_option() {
  local bin="$1" key="$2" question="$3" values="$4" printer="${5:-print_value_options}"

  printf '%s\n' "$question"
  "$printer" "$values"
  while :; do
    read_answer "$(t prompt.select) "
    if [ -z "$ANSWER" ]; then
      unset "CONFIGURED_VALUES[$key]" 2>/dev/null || true
      return
    fi

    case " $values " in
      *" $ANSWER "*)
        apply_config_option "$bin" "${key}=${ANSWER}" && CONFIGURED_VALUES["$key"]="$ANSWER"
        return
        ;;
      *) t error.invalid_value "$values" >&2 ;;
    esac
  done
}

configure_non_negative_integer() {
  local bin="$1" key="$2" question="$3" hint="$4"

  printf '%s\n' "$question"
  printf '  %s\n' "$(t value.non_negative_integer)"
  [ -z "$hint" ] || printf '  %s\n' "$hint"
  while :; do
    read_answer "$(t prompt.select) "
    if [ -z "$ANSWER" ]; then
      unset "CONFIGURED_VALUES[$key]" 2>/dev/null || true
      return
    fi
    case "$ANSWER" in
      *[!0-9]*) t error.invalid_value "$(t value.non_negative_integer)" >&2 ;;
      *)
        apply_config_option "$bin" "${key}=${ANSWER}" && CONFIGURED_VALUES["$key"]="$ANSWER"
        return
        ;;
    esac
  done
}

configure_directory() {
  local bin="$1" key="$2" question="$3" hint="$4"

  printf '%s\n' "$question"
  printf '  %s\n' "$(t value.existing_directory)"
  [ -z "$hint" ] || printf '  %s\n' "$hint"
  while :; do
    read_answer "$(t prompt.select) "
    if [ -z "$ANSWER" ]; then
      unset "CONFIGURED_VALUES[$key]" 2>/dev/null || true
      return
    fi
    if [ -d "$ANSWER" ]; then
      apply_config_option "$bin" "${key}=${ANSWER}" && CONFIGURED_VALUES["$key"]="$ANSWER"
      return
    fi
    t error.invalid_value "$(t value.existing_directory)" >&2
  done
}

print_language_options() {
  t environment.languages
  printf '%s\n' \
    'English - en' \
    'Українська - uk' \
    'Polski - pl' \
    'Deutsch - de' \
    'Français - fr' \
    'Español - es' \
    'Čeština - cs' \
    'Română - ro' \
    'Italiano - it' \
    '한국어 - ko' \
    '日本語 - ja' \
    'Bahasa Indonesia - id' \
    'हिन्दी - hi' \
    'Ελληνικά - el' \
    'Türkçe - tr' \
    'Қазақша - kk' \
    'ქართული - ka' \
    'Magyar - hu' \
    'Svenska - sv' \
    'Norsk - no' \
    'Dansk - da' \
    'Suomi - fi'
}

print_value_options() {
  local values="$1" value
  t environment.options
  for value in $values; do
    printf '  - %s\n' "$value"
  done
}

print_graphics_options() {
  t environment.options
  printf '  - %s\n' \
    "$(t environment.graphics.auto)" \
    "$(t environment.graphics.ascii)" \
    "$(t environment.graphics.dots)" \
    "$(t environment.graphics.kitty)" \
    "$(t environment.graphics.sixel)" \
    "$(t environment.graphics.iterm2)"
}

print_view_options() {
  t environment.options
  printf '  - %s\n' \
    "$(t environment.view.book)" \
    "$(t environment.view.right)" \
    "$(t environment.view.circle)" \
    "$(t environment.view.right_circle)"
}

print_configured_summary() {
  local name

  [ "${#CONFIGURED_VALUES[@]}" -gt 0 ] || return
  for name in "${CONFIGURED_ORDER[@]}"; do
    [ -n "${CONFIGURED_VALUES[$name]+x}" ] || continue
    printf '%s="%s"\n' "$name" "${CONFIGURED_VALUES[$name]}"
  done
}

configure_environment() {
  [ "$AUTO_YES" = "true" ] && return
  if ! ask_yes_no "$(t environment.configure)"; then
    return
  fi

  local bin
  bin="${INSTALL_DIR}/${BIN_NAME}"
  t environment.saved.config "$(config_file_path)"
  configure_option "$bin" 'language' "$(t environment.language)" 'en uk pl de fr es cs ro it ko ja id hi el tr kk ka hu sv no da fi' print_language_options
  configure_option "$bin" 'graphics' "$(t environment.graphics)" 'auto ascii dots kitty sixel iterm2' print_graphics_options
  configure_option "$bin" 'view' "$(t environment.view)" 'book-view right-view circle-view right-circle-view' print_view_options
  configure_non_negative_integer "$bin" 'prerender.next' "$(t environment.prerendered_next)" "$(t environment.prerendered_hint)"
  configure_non_negative_integer "$bin" 'prerender.previous' "$(t environment.prerendered_previous)" "$(t environment.prerendered_hint)"
  configure_directory "$bin" 'directory' "$(t environment.directory)" "$(t environment.directory_hint)"
  print_configured_summary
}

install_binary() {
  local source="$1"

  mkdir -p "$INSTALL_DIR"
  if [ ! -w "$INSTALL_DIR" ]; then
    t error.install_dir_writable "$INSTALL_DIR" >&2
    t hint.install_dir_writable >&2
    exit 1
  fi

  install -m 0755 "$source" "${INSTALL_DIR}/${BIN_NAME}"
}

main() {
  TMP_DIR="$(mktemp -d)"
  trap cleanup EXIT HUP INT TERM
  load_locale "$(detect_locale)"
  parse_args "$@"
  require_cmd curl
  require_cmd tar
  require_cmd install

  local os arch version archive url
  os="$(normalize_os)"
  arch="$(normalize_arch)"
  version="$(resolve_version)"
  archive="${BIN_NAME}_${version}_${os}_${arch}.tar.gz"
  url="https://github.com/${REPO}/releases/download/${version}/${archive}"

  confirm_upgrade_if_needed "$version"
  if [ "$REINSTALL_CONFIRMED" != "true" ]; then
    confirm_install "$(t confirm.install "$version" "$INSTALL_DIR")"
  fi

  t status.downloading "$archive"
  curl -fsSL "$url" -o "${TMP_DIR}/${archive}"
  t status.extracting "$archive"
  tar -xzf "${TMP_DIR}/${archive}" -C "$TMP_DIR"

  if [ ! -f "${TMP_DIR}/${BIN_NAME}" ]; then
    t error.binary_missing "$BIN_NAME" >&2
    exit 1
  fi

  install_binary "${TMP_DIR}/${BIN_NAME}"
  ensure_path_setup
  configure_environment

  t status.installed "$version" "${INSTALL_DIR}/${BIN_NAME}"
  if ! path_contains_install_dir; then
    t status.restart
    t status.reload_shell "$(shell_config_file)"
  fi
}

main "$@"

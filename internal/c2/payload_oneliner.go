package c2

import (
	"fmt"

	"cyberstrike-ai/internal/database"
)

// SECURITY STUB (fleet policy, 2026-10-11):
// This file is a sanitized replacement for upstream payload_oneliner.go
// (reverse-shell payload generator: bash/nc/python/perl/powershell/curl beacon).
// Pilot deployment policy: NO payload generation capability in the binary.
// c2.enabled=false upstream; this stub enforces it at code level.
// Original logic preserved in c2-stub.patch (reversed) for upstream rebase.
// Sanctioned by director in fleet-bus 2026-10-11 04:55 (id a9238861).

// OnelinerKind — тип oneliner payload (сохранён для совместимости сигнатур)
type OnelinerKind string

const (
	OnelinerBash       OnelinerKind = "bash"
	OnelinerNc         OnelinerKind = "nc"
	OnelinerNcMkfifo   OnelinerKind = "nc_mkfifo"
	OnelinerPython     OnelinerKind = "python"
	OnelinerPerl       OnelinerKind = "perl"
	OnelinerPowerShell OnelinerKind = "powershell"
	OnelinerCurl       OnelinerKind = "curl_beacon"
)

// AllOnelinerKinds возвращает список типов (без логики генерации)
func AllOnelinerKinds() []OnelinerKind {
	return []OnelinerKind{
		OnelinerBash, OnelinerNc, OnelinerNcMkfifo,
		OnelinerPython, OnelinerPerl,
		OnelinerPowerShell, OnelinerCurl,
	}
}

// OnelinerKindsForListener возвращает совместимые типы (metadata-only, ничего не генерирует)
func OnelinerKindsForListener(listenerType string) []OnelinerKind {
	switch ListenerType(listenerType) {
	case ListenerTypeTCPReverse:
		return []OnelinerKind{OnelinerBash, OnelinerNc, OnelinerNcMkfifo, OnelinerPython, OnelinerPerl, OnelinerPowerShell}
	case ListenerTypeHTTPBeacon, ListenerTypeHTTPSBeacon, ListenerTypeWebSocket:
		return []OnelinerKind{OnelinerCurl}
	default:
		return nil
	}
}

// IsOnelinerCompatible — совместимость типа с листенером (metadata-only)
func IsOnelinerCompatible(listenerType string, kind OnelinerKind) bool {
	switch ListenerType(listenerType) {
	case ListenerTypeTCPReverse:
		return kind == OnelinerBash || kind == OnelinerNc || kind == OnelinerNcMkfifo ||
			kind == OnelinerPython || kind == OnelinerPerl || kind == OnelinerPowerShell
	case ListenerTypeHTTPBeacon, ListenerTypeHTTPSBeacon, ListenerTypeWebSocket:
		return kind == OnelinerCurl
	default:
		return false
	}
}

// OnelinerInput — вход генерации (сохранён для совместимости сигнатур)
type OnelinerInput struct {
	Kind         OnelinerKind
	Host         string
	Port         int
	HTTPBaseURL  string
	ImplantToken string
}

// ValidateOnelinerForListener — stub: всегда отказ (генерация отключена)
func ValidateOnelinerForListener(listener *database.C2Listener, kind OnelinerKind) error {
	return fmt.Errorf("C2 payload generation disabled in this deployment (fleet policy: no payload generation)")
}

// GenerateOneliner — stub: ВСЕГДА возвращает ошибку, никакой генерации
func GenerateOneliner(in OnelinerInput) (string, error) {
	return "", fmt.Errorf("C2 payload generation disabled in this deployment (fleet policy: no payload generation)")
}

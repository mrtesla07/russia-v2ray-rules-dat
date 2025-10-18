# Что это

Репозиторий автоматически собирает целевые `geoip.dat` и `geosite.dat` для клиентов семейства V2Ray/Xray/sing-box. Сборка запускается по расписанию и при ручном запуске workflow `build.yml`.

Готовые файлы доступны в разделе релизов:

- `geoip.dat`: https://raw.githubusercontent.com/runetfreedom/russia-v2ray-rules-dat/release/geoip.dat
- `geosite.dat`: https://raw.githubusercontent.com/runetfreedom/russia-v2ray-rules-dat/release/geosite.dat

# Как мы собираем geosite.dat

1. GitHub Actions вытягивает актуальные списки из `v2fly/domain-list-community`.
2. К ним добавляются кастомные списки из каталога `custom-lists/`.
3. Утилита `domain-list-custom` формирует итоговый `geosite.dat`, включающий только нужные нам категории.

Таким образом, вы всегда получаете свежие домены из upstream и одновременно кастомные подборки под российские сервисы.

# Какие категории входят сейчас

Основной `geosite.dat` содержит следующие разделы:

- `geosite:category-ru` — базовый список доменов в зоне RU/СНГ из `domain-list-community` (включает подсписки `category-gov-ru`, `mailru`, `ok`, `ozon`, `vk`, `yandex`).
- `geosite:category-media-ru` — крупные российские медиа.
- `geosite:category-gov-ru` — государственные сервисы и порталы.
- `geosite:mailru`, `geosite:ok`, `geosite:ozon`, `geosite:vk`, `geosite:yandex` — точечные списки из DLC.
- `geosite:category-torrent` — трекеры и dht-роутеры (см. `custom-lists/category-torrent`).
- `geosite:twitch` — официальный список Twitch из DLC.
- `geosite:trovo` — добавлен вручную, охватывает `trovo.live`, `trovoapp.com` и их поддомены.
- `geosite:okko` — все сервисы онлайн-кинотеатра Okko.
- `geosite:category-streaming-ru` — российские онлайн-кинотеатры и видеосервисы (ivi, Wink, KION, START, PREMIER, Аmediateka, кинопоиск, Rutube и др.).
- `geosite:category-bank-ru` — крупнейшие российские банки.
- `geosite:steam` — расширенный список Steam/Valve, включая CDN и зеркала.
- `geosite:win-update` — адреса службы обновлений Windows.
- `geosite:private` — специальные/локальные домены (RFC 2606/6761 и т.п.).

Вы можете расширять набор, добавляя новые файлы в `custom-lists/` и/или подключая дополнительные списки из DLC через `scripts/geosite-dlc-lists.txt`.

# geoip.dat

Файл geoip берётся из репозитория [runetfreedom/russia-blocked-geoip](https://github.com/runetfreedom/russia-blocked-geoip) и содержит агрегированные IP-диапазоны, используемые для обхода блокировок в РФ.

# Локальная проверка

Чтобы воспроизвести сборку локально:

```powershell
git clone https://github.com/v2fly/domain-list-community
git clone https://github.com/runetfreedom/domain-list-custom
$env:GOEXE = 'C:\Program Files\Go\bin\go.exe' # путь к go, если требуется

mkdir tmp\geosite-data
Get-Content scripts/geosite-dlc-lists.txt | ForEach-Object {
    $name = $_.Trim()
    if($name -and -not $name.StartsWith('#')) {
        Copy-Item "domain-list-community/data/$name" "tmp/geosite-data/"
    }
}
Copy-Item custom-lists\* tmp\geosite-data\

& $env:GOEXE run ./domain-list-custom --datapath=tmp/geosite-data --datname=geosite.dat --exportlists= --togfwlist= --outputpath=tmp/geosite-output
```

Итоговый файл появится в `tmp/geosite-output/geosite.dat`.

# Связанные проекты

- [runetfreedom/russia-blocked-geoip](https://github.com/runetfreedom/russia-blocked-geoip) — источник IP-диапазонов.
- [runetfreedom/russia-blocked-geosite](https://github.com/runetfreedom/russia-blocked-geosite) — расширенный список доменов.
- [runetfreedom/russia-v2ray-custom-routing-list](https://github.com/runetfreedom/russia-v2ray-custom-routing-list) — дополнительные правила маршрутизации.
- [runetfreedom/geodat2srs](https://github.com/runetfreedom/geodat2srs) — конвертер dat → sing-box rule-set.
- [hydraponique/roscomvpn-geosite](https://github.com/hydraponique/roscomvpn-geosite) — дополнительные списки доменов (Steam, Windows Update).

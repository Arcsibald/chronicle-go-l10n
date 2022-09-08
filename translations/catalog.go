// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package translations

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"de":    &dictionary{index: deIndex, data: deData},
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
		"es":    &dictionary{index: esIndex, data: esData},
		"fr":    &dictionary{index: frIndex, data: frData},
		"ja":    &dictionary{index: jaIndex, data: jaData},
		"nl":    &dictionary{index: nlIndex, data: nlData},
		"pt":    &dictionary{index: ptIndex, data: ptData},
		"sv":    &dictionary{index: svIndex, data: svData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%d days":    1,
	"%d hours":   2,
	"%d minutes": 3,
	"%d months":  0,
	"%s has created a notifier which connects this channel to an external calendar.":  4,
	"%s has deleted a notifier which connected this channel to an external calendar.": 5,
	"All Day":                     8,
	"Duration":                    14,
	"Event has been cancelled":    11,
	"Events for the next %d days": 6,
	"Location":                    9,
	"New event scheduled":         10,
	"No events scheduled":         7,
	"Scheduled for":               13,
	"Starting":                    15,
	"Updated":                     12,
}

var deIndex = []uint32{ // 17 elements
	0x00000000, 0x00000024, 0x00000044, 0x00000069,
	0x0000008e, 0x000000e9, 0x00000149, 0x00000177,
	0x00000195, 0x000001a4, 0x000001b6, 0x000001d1,
	0x000001ee, 0x000001fb, 0x0000020c, 0x00000212,
	0x00000219,
} // Size: 92 bytes

const deData string = "" + // Size: 537 bytes
	"\x14\x01\x81\x01\x00=\x01\x0c\x02%[1]d Monat\x00\x0e\x02%[1]d Monaten" +
	"\x14\x01\x81\x01\x00=\x01\x0a\x02%[1]d Tag\x00\x0c\x02%[1]d Tagen\x14" +
	"\x01\x81\x01\x00=\x01\x0d\x02%[1]d Stunde\x00\x0e\x02%[1]d Stunden\x14" +
	"\x01\x81\x01\x00=\x01\x0d\x02%[1]d Minute\x00\x0e\x02%[1]d Minuten\x02%[" +
	"1]s hat einen Notifier erstellt, der diesen Kanal mit einem externen Kal" +
	"ender verbindet.\x02%[1]s hat einen Notifier gelöscht, der diesen Kanal " +
	"mit einem externen Kalender verbunden hat.\x02Veranstaltungen für die nä" +
	"chsten %[1]d Tage\x02Keine Veranstaltungen geplant\x02Den ganzen Tag\x02" +
	"Veranstaltungsort\x02Neue Veranstaltung geplant\x02Veranstaltung wurde a" +
	"bgesagt\x02Aktualisiert\x02Eingetragen für\x02Dauer\x02Beginn"

var en_USIndex = []uint32{ // 17 elements
	0x00000000, 0x00000023, 0x00000042, 0x00000063,
	0x00000088, 0x000000da, 0x0000012d, 0x0000016e,
	0x00000182, 0x0000018a, 0x00000193, 0x000001a7,
	0x000001c0, 0x000001c8, 0x000001d6, 0x000001df,
	0x000001e8,
} // Size: 92 bytes

const en_USData string = "" + // Size: 488 bytes
	"\x14\x01\x81\x01\x00=\x01\x0c\x02%[1]d month\x00\x0d\x02%[1]d months\x14" +
	"\x01\x81\x01\x00=\x01\x0a\x02%[1]d day\x00\x0b\x02%[1]d days\x14\x01\x81" +
	"\x01\x00=\x01\x0b\x02%[1]d hour\x00\x0c\x02%[1]d hours\x14\x01\x81\x01" +
	"\x00=\x01\x0d\x02%[1]d minute\x00\x0e\x02%[1]d minutes\x02%[1]s has crea" +
	"ted a notifier which connects this channel to an external calendar.\x02%" +
	"[1]s has deleted a notifier which connected this channel to an external " +
	"calendar.\x14\x01\x81\x01\x00=\x01\x18\x02Events for the next day\x00" +
	"\x1f\x02Events for the next %[1]d days\x02No events scheduled\x02All Day" +
	"\x02Location\x02New event scheduled\x02Event has been cancelled\x02Updat" +
	"ed\x02Scheduled for\x02Duration\x02Starting"

var esIndex = []uint32{ // 17 elements
	0x00000000, 0x00000020, 0x00000041, 0x00000062,
	0x00000087, 0x000000d7, 0x0000012b, 0x0000017b,
	0x00000196, 0x000001a2, 0x000001ad, 0x000001c5,
	0x000001e1, 0x000001ed, 0x000001fd, 0x00000207,
	0x00000212,
} // Size: 92 bytes

const esData string = "" + // Size: 530 bytes
	"\x14\x01\x81\x01\x00=\x01\x0a\x02%[1]d mes\x00\x0c\x02%[1]d meses\x14" +
	"\x01\x81\x01\x00=\x01\x0b\x02%[1]d día\x00\x0c\x02%[1]d días\x14\x01\x81" +
	"\x01\x00=\x01\x0b\x02%[1]d hora\x00\x0c\x02%[1]d horas\x14\x01\x81\x01" +
	"\x00=\x01\x0d\x02%[1]d minuto\x00\x0e\x02%[1]d minutos\x02%[1]s ha cread" +
	"o un notificador que conecta este canal con un calendario externo\x02%[1" +
	"]s ha eliminado un notificador que conectaba este canal a un calendario " +
	"externo.\x14\x01\x81\x01\x00=\x01\x1f\x02Eventos para el día siguiente" +
	"\x00'\x02Eventos para los próximos %[1]d días\x02No hay eventos programa" +
	"dos\x02Todo el dia\x02Ubicación\x02Nuevo evento programado\x02El evento " +
	"ha sido cancelado\x02Actualizado\x02Programado para\x02Duración\x02Comen" +
	"zando"

var frIndex = []uint32{ // 17 elements
	0x00000000, 0x0000000b, 0x00000029, 0x0000004c,
	0x00000071, 0x000000c0, 0x00000114, 0x00000169,
	0x00000182, 0x00000194, 0x000001a0, 0x000001be,
	0x000001dc, 0x000001e7, 0x000001f6, 0x000001fd,
	0x00000206,
} // Size: 92 bytes

const frData string = "" + // Size: 518 bytes
	"\x02%[1]d mois\x14\x01\x81\x01\x00=\x01\x08\x02Un jour\x00\x0c\x02%[1]d " +
	"jours\x14\x01\x81\x01\x00=\x01\x0c\x02%[1]d heure\x00\x0d\x02%[1]d heure" +
	"s\x14\x01\x81\x01\x00=\x01\x0d\x02%[1]d minute\x00\x0e\x02%[1]d minutes" +
	"\x02%[1]s a créé un notificateur qui connecte ce canal à un calendrier e" +
	"xterne.\x02%[1]s a supprimé un notificateur qui connectait ce canal à un" +
	" calendrier externe.\x14\x01\x81\x01\x00=\x01\x1e\x02Les événements du l" +
	"endemain\x00-\x02Événements pour les %[1]d\u00a0prochains jours\x02Aucun" +
	" événement prévu\x02Toute la journée\x02Emplacement\x02Nouvel événement " +
	"programmé\x02L'événement a été annulé\x02Actualisé\x02Planifié pour\x02D" +
	"urée\x02Commence"

var jaIndex = []uint32{ // 17 elements
	0x00000000, 0x00000022, 0x0000003e, 0x00000060,
	0x0000007c, 0x000000eb, 0x00000142, 0x00000163,
	0x0000019a, 0x000001a1, 0x000001ae, 0x000001cd,
	0x000001f8, 0x000001ff, 0x00000206, 0x0000020d,
	0x00000217,
} // Size: 92 bytes

const jaData string = "" + // Size: 535 bytes
	"\x14\x01\x81\x01\x00=\x01\x0c\x02%[1]dヵ月\x00\x0c\x02%[1]dヵ月\x14\x01\x81" +
	"\x01\x00=\x01\x09\x02%[1]d日\x00\x09\x02%[1]d日\x14\x01\x81\x01\x00=\x01" +
	"\x0c\x02%[1]d時間\x00\x0c\x02%[1]d時間\x14\x01\x81\x01\x00=\x01\x09\x02%[1]d" +
	"分\x00\x09\x02%[1]d分\x02%[1]sはこちらを接続したカレンダーの通知を行うチャンネルとして設定しました。\x02%[1" +
	"]sはこちらをカレンダーの通知チャンネルから削除しました。\x02%[1]d日間のイベント情報\x02予定されているイベントはありません。" +
	"\x02終日\x02開催場所\x02新しいイベントを企画\x02イベントは中止となりました。\x02更新\x02予定\x02期間\x02開催中"

var nlIndex = []uint32{ // 17 elements
	0x00000000, 0x00000024, 0x00000044, 0x0000004e,
	0x00000073, 0x000000c3, 0x0000011e, 0x00000171,
	0x0000018a, 0x00000196, 0x0000019d, 0x000001b5,
	0x000001ce, 0x000001d9, 0x000001e6, 0x000001ef,
	0x000001f9,
} // Size: 92 bytes

const nlData string = "" + // Size: 505 bytes
	"\x14\x01\x81\x01\x00=\x01\x0c\x02%[1]d maand\x00\x0e\x02%[1]d maanden" +
	"\x14\x01\x81\x01\x00=\x01\x0a\x02%[1]d dag\x00\x0c\x02%[1]d dagen\x02%[1" +
	"]d uur\x14\x01\x81\x01\x00=\x01\x0d\x02%[1]d minuut\x00\x0e\x02%[1]d min" +
	"uten\x02%[1]s heeft een notifier gemaakt die dit kanaal koppelt aan een " +
	"externe agenda.\x02%[1]s heeft een notifier verwijderd die dit kanaal he" +
	"eft gekoppeld aan een externe agenda.\x14\x01\x81\x01\x00=\x01!\x02Evene" +
	"menten voor de volgende dag\x00(\x02Evenementen voor de komende %[1]d da" +
	"gen\x02Geen evenementen gepland\x02De hele dag\x02Plaats\x02Nieuw evenem" +
	"ent gepland\x02Evenement is geannuleerd\x02Bijgewerkt\x02Gepland voor" +
	"\x02Looptijd\x02Beginnend"

var ptIndex = []uint32{ // 17 elements
	0x00000000, 0x00000021, 0x00000040, 0x00000061,
	0x00000086, 0x000000d7, 0x0000012b, 0x0000016e,
	0x00000186, 0x00000191, 0x00000197, 0x000001ae,
	0x000001c5, 0x000001d0, 0x000001e0, 0x000001ea,
	0x000001f5,
} // Size: 92 bytes

const ptData string = "" + // Size: 501 bytes
	"\x14\x01\x81\x01\x00=\x01\x0b\x02%[1]d mês\x00\x0c\x02%[1]d meses\x14" +
	"\x01\x81\x01\x00=\x01\x0a\x02%[1]d dia\x00\x0b\x02%[1]d dias\x14\x01\x81" +
	"\x01\x00=\x01\x0b\x02%[1]d hora\x00\x0c\x02%[1]d horas\x14\x01\x81\x01" +
	"\x00=\x01\x0d\x02%[1]d minuto\x00\x0e\x02%[1]d minutos\x02%[1]s ha cread" +
	"o un notificador que conecta este canal con un calendario externo.\x02%[" +
	"1]s ha eliminado un notificador que conectaba este canal a un calendario" +
	" externo.\x14\x01\x81\x01\x00=\x01\x18\x02Eventos no dia seguinte\x00!" +
	"\x02Eventos nos próximos %[1]d dias\x02Sem eventos programados\x02Todo o" +
	" dia\x02Local\x02Novo evento programado\x02O evento foi cancelado\x02Atu" +
	"alizado\x02Programado para\x02Duração\x02A começar"

var svIndex = []uint32{ // 17 elements
	0x00000000, 0x00000026, 0x00000046, 0x00000069,
	0x0000008d, 0x000000e0, 0x00000137, 0x00000157,
	0x00000170, 0x0000017b, 0x00000181, 0x00000197,
	0x000001b2, 0x000001bd, 0x000001c6, 0x000001cd,
	0x000001d5,
} // Size: 92 bytes

const svData string = "" + // Size: 469 bytes
	"\x14\x01\x81\x01\x00=\x01\x0d\x02%[1]d Månad\x00\x0f\x02%[1]d Månader" +
	"\x14\x01\x81\x01\x00=\x01\x0a\x02%[1]d Dag\x00\x0c\x02%[1]d Dagar\x14" +
	"\x01\x81\x01\x00=\x01\x0c\x02%[1]d Timme\x00\x0d\x02%[1]d Timmar\x14\x01" +
	"\x81\x01\x00=\x01\x0c\x02%[1]d Minut\x00\x0e\x02%[1]d Minuter\x02%[1]s h" +
	"ar skapat en notifier som kopplar den här kanalen till en extern kalende" +
	"r.\x02%[1]s har tagit bort en notifier som kopplar den här kanaler till " +
	"en extern kalender.\x02Event de kommande %[1]d dagarna\x02Inga evenemang" +
	" planerade\x02Hela dagen\x02Plats\x02Nytt event schemalagt\x02Event har " +
	"blivit inställt\x02Uppdaterad\x02Starttid\x02Längd\x02Börjar"

	// Total table size 4819 bytes (4KiB); checksum: F3CB2AD7

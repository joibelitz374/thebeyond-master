package i18n

import "github.com/quickpowered/frilly/config/language"

type RefundLocale struct {
	InfoText       string
	NoPayments     string
	SuccessMessage string
}

var RefundMessages = map[language.Language]RefundLocale{
	language.English: {
		InfoText:       "<b>How does the refund work?</b>\nFunds get returned by the admin right under the post in your Telegram channel using \"Stars\" currency.\n\n<b>How soon will the funds be back?</b>\nRefund can take anywhere from 30 minutes to 23 days.\n\n<b>Can I get just a partial refund?</b>\nNah, it returns the whole remaining balance. If you only want a chunk, just subscribe again for a different timeframe.\n\n<b>You sure you wanna request a refund?</b>",
		NoPayments:     "No payments found",
		SuccessMessage: "Refund requested! Wait for a message from the admin in DM",
	},
	language.Deutsch: {
		InfoText:       "<b>Wie funktioniert die Rückerstattung?</b>\nDie Kohle wird vom Admin unter dem Post in deinem Telegram-Kanal in der Währung «Stars» zurückgezahlt.\n\n<b>Wie schnell kommt die Kohle zurück?</b>\nDie Rückerstattung dauert von 30 Minuten bis 23 Tage.\n\n<b>Kann ich nur einen Teil zurückbekommen?</b>\nNee, der gesamte Restbetrag wird zurückgezahlt. Wenn du nur 'nen Teil brauchst, abonnier einfach neu für 'ne andere Laufzeit.\n\n<b>Willst du wirklich 'ne Rückerstattung anfordern?</b>",
		NoPayments:     "Keine Zahlungen gefunden",
		SuccessMessage: "Rückerstattung beantragt! Warte auf eine Nachricht vom Admin im Chat",
	},
	language.Nederlands: {
		InfoText:       "<b>Hoe werkt de terugbetaling?</b>\nHet geld wordt teruggegeven door de admin onder de post in je Telegram-kanaal in de valuta «Stars».\n\n<b>Hoe snel krijg ik het geld terug?</b>\nTerugbetaling duurt van 30 minuten tot 23 dagen.\n\n<b>Kan ik een deel terugkrijgen?</b>\nNee, het hele resterende saldo wordt terugbetaald. Als je maar een deel wilt, abonneer je gewoon opnieuw voor een andere periode.\n\n<b>Wil je echt een terugbetaling aanvragen?</b>",
		NoPayments:     "Geen betalingen gevonden",
		SuccessMessage: "Terugbetaling aangevraagd! Wacht op een berichtje van de admin in je DM",
	},
	language.Svenska: {
		InfoText:       "<b>Hur funkar återbetalningen?</b>\nPengarna returneras av admin under inlägget i din Telegram-kanal i valutan «Stars».\n\n<b>Hur snart kommer pengarna tillbaka?</b>\nÅterbetalning tar från 30 minuter till 23 dagar.\n\n<b>Kan jag få tillbaka bara en del?</b>\nNej, hela kvarvarande saldot returneras. Om du bara vill ha en bit, prenumerera igen för en annan period.\n\n<b>Vill du verkligen begära återbetalning?</b>",
		NoPayments:     "Inga betalningar hittades.",
		SuccessMessage: "Återbetalning begärd! Vänta på meddelande från admin i DM",
	},
	language.Norsk: {
		InfoText:       "<b>Hvordan fungerer refusjonen?</b>\nPengene returneres av admin under innlegget i Telegram-kanalen din i valutaen «Stars».\n\n<b>Hvor fort kommer pengene tilbake?</b>\nRefusjon tar fra 30 minutter til 23 dager.\n\n<b>Kan jeg få tilbake bare en del?</b>\nNei, hele gjenværende saldo returneres. Hvis du bare trenger en del, abonner igjen for en annen periode.\n\n<b>Vil du virkelig be om refusjon?</b>",
		NoPayments:     "Ingen betalinger funnet",
		SuccessMessage: "Refusjon er sendt! Vent på melding fra admin i privat",
	},
	language.Dansk: {
		InfoText:       "<b>Hvordan virker refund?</b>\nPengene returneres af admin under posten i din Telegram-kanal i valutaen «Stars».\n\n<b>Hvor hurtigt kommer pengene tilbage?</b>\nRefund tager fra 30 minutter til 23 dage.\n\n<b>Kan jeg få kun en del tilbage?</b>\nNej, hele den resterende saldo returneres. Hvis du kun har brug for en del, abonner igen for en anden periode.\n\n<b>Vil du virkelig anmode om refund?</b>",
		NoPayments:     "Ingen betalinger fundet",
		SuccessMessage: "Refund er anmodet! Vent på besked fra admin i privat chat",
	},
	language.Español: {
		InfoText:       "<b>¿Cómo funciona el reembolso?</b>\nLos fondos se devuelven por el admin debajo del post en tu canal de Telegram en la moneda «Stars».\n\n<b>¿Cuánto tarda el reembolso?</b>\nEl reembolso toma de 30 minutos a 23 días.\n\n<b>¿Puedo devolver solo una parte?</b>\nNo, se devuelve todo el saldo restante. Si solo quieres una parte, suscríbete de nuevo por otro período.\n\n<b>¿De verdad quieres solicitar un reembolso?</b>",
		NoPayments:     "No se encontraron pagos",
		SuccessMessage: "¡Reembolso solicitado! Espera el mensaje del admin en privados",
	},
	language.Français: {
		InfoText:       "<b>Comment marche le remboursement?</b>\nLes fonds sont renvoyés par l'admin sous le post dans ton canal Telegram en devise «Stars».\n\n<b>Combien de temps pour le remboursement?</b>\nLe remboursement prend de 30 minutes à 23 jours.\n\n<b>Peut-on rembourser juste une partie?</b>\nNan, tout le solde restant est remboursé. Si t'as besoin que d'une partie, abonne-toi à nouveau pour une autre durée.\n\n<b>Tu veux vraiment demander un remboursement?</b>",
		NoPayments:     "Aucun paiement trouvé",
		SuccessMessage: "Remboursement demandé ! Attends le message de l’admin en MP",
	},
	language.Português: {
		InfoText:       "<b>Como funciona o reembolso?</b>\nOs fundos são devolvidos pelo admin sob o post no teu canal Telegram na moeda «Stars».\n\n<b>Quanto tempo leva o reembolso?</b>\nO reembolso demora de 30 minutos a 23 dias.\n\n<b>Posso devolver só uma parte?</b>\nNão, todo o saldo restante é devolvido. Se precisares só de uma parte, assina de novo por outro período.\n\n<b>Queres mesmo pedir um reembolso?</b>",
		NoPayments:     "Nenhum pagamento encontrado",
		SuccessMessage: "Reembolso solicitado! Aguarde a mensagem do admin no privado",
	},
	language.Italiana: {
		InfoText:       "<b>Come funziona il rimborso?</b>\nI fondi vengono restituiti dall'admin sotto il post nel tuo canale Telegram nella valuta «Stars».\n\n<b>Quanto ci vuole per il rimborso?</b>\nIl rimborso richiede da 30 minuti a 23 giorni.\n\n<b>Posso restituire solo una parte?</b>\nNo, viene restituito l'intero saldo rimanente. Se ti serve solo una parte, iscriviti di nuovo per un altro periodo.\n\n<b>Vuoi davvero richiedere un rimborso?</b>",
		NoPayments:     "Nessun pagamento trovato",
		SuccessMessage: "Rimborso richiesto! Aspetta il messaggio dell’admin in privato",
	},
	language.Русский: {
		InfoText:       "<b>Как работает возврат средств?</b>\nСредства возвращаются администратором под постом в Вашем Telegram-канале валютой «Stars».\n\n<b>Как скоро вернут средства?</b>\nВозврат средств занимает от 30 минут до 23 дней.\n\n<b>Можно ли вернуть часть средств?</b>\nВозвращается весь оставшийся баланс. Если нужно только часть, подпишитесь снова на другой срок.\n\n<b>Вы действительно хотите запросить возврат?</b>",
		NoPayments:     "Платежи не найдены",
		SuccessMessage: "Возврат оформлен! Ожидай сообщения от админа в личку",
	},
	language.Українська: {
		InfoText:       "<b>Як працює повернення коштів?</b>\nКошти повертаються адміністратором під постом у Вашому Telegram-каналі валютою «Stars».\n\n<b>Як скоро повернуть кошти?</b>\nПовернення коштів займає від 30 хвилин до 23 днів.\n\n<b>Чи можна повернути частину коштів?</b>\nПовертається весь залишковий баланс. Якщо потрібно тільки частину, підпишіться знову на інший термін.\n\n<b>Ви дійсно хочете запросити повернення?</b>",
		NoPayments:     "Платежі не знайдено",
		SuccessMessage: "Повернення оформлено! Чекай повідомлення від адміна в приват",
	},
	language.Polski: {
		InfoText:       "<b>Jak działa zwrot środków?</b>\nŚrodki są zwracane przez admina pod postem w Twoim kanale Telegram w walucie «Stars».\n\n<b>Jak szybko środki zostaną zwrócone?</b>\nZwrot zajmuje od 30 minut do 23 dni.\n\n<b>Czy można zwrócić tylko część?</b>\nZwrot obejmuje cały pozostały balans. Jeśli chcesz tylko część, subskrybuj znowu na inny okres.\n\n<b>Czy na pewno chcesz poprosić o zwrot?</b>",
		NoPayments:     "Nie znaleziono płatności",
		SuccessMessage: "Zwrot złożony! Czekaj na wiadomość od admina w prywatnych",
	},
	language.Ceština: {
		InfoText:       "<b>Jak funguje vrácení peněz?</b>\nPeníze jsou vráceny adminem pod příspěvkem ve tvém Telegram kanálu v měně «Stars».\n\n<b>Jak brzy peníze přijdou zpět?</b>\nVrácení trvá od 30 minut do 23 dnů.\n\n<b>Lze vrátit jen část?</b>\nVrátí se celý zbývající zůstatek. Pokud potřebuješ jen část, přihlas se znovu na jinou dobu.\n\n<b>Opravdu chceš požádat o vrácení?</b>",
		NoPayments:     "Žádné platby nenalezeny",
		SuccessMessage: "Vrácení požádáno! Čekej na zprávu od admina v soukromí",
	},
	language.Български: {
		InfoText:       "<b>Как работи възстановяването на средства?</b>\nСредствата се възстановяват от админа под публикацията във твоя Telegram канал в валутата «Stars».\n\n<b>Колко бързо ще се възстановят средствата?</b>\nВъзстановяването отнема от 30 минути до 23 дни.\n\n<b>Може ли да се възстанови само част?</b>\nВъзстановява се целият останал баланс. Ако ти трябва само част, абонирай се пак за друг период.\n\n<b>Наистина ли искаш да поискаш възстановяване?</b>",
		NoPayments:     "Няма намерени плащания",
		SuccessMessage: "Възстановяване поискано! Изчакай съобщение от админа в лично",
	},
	language.Српски: {
		InfoText:       "<b>Како функционише повраћај средстава?</b>\nСредства се враћају од админа под постом у твом Телеграм каналу у валути «Stars».\n\n<b>Колико брзо ће се средства вратити?</b>\nПовраћај траје од 30 минута до 23 дана.\n\n<b>Може ли се вратити само део?</b>\nВраћа се цео преостали баланс. Ако ти треба само део, претплати се поново на други период.\n\n<b>Да ли стварно желиш да затражиш повраћај?</b>",
		NoPayments:     "Нема пронађених плаћања",
		SuccessMessage: "Повраћај је затражен! Чекај поруку од админа у приватно",
	},
	language.Hrvatski: {
		InfoText:       "<b>Kako funkcionira povrat sredstava?</b>\nSredstva vraća admin ispod posta u tvom Telegram kanalu u valuti «Stars».\n\n<b>Koliko brzo će se sredstva vratiti?</b>\nPovrat traje od 30 minuta do 23 dana.\n\n<b>Može li se vratiti samo dio?</b>\nVraća se cijeli preostali saldo. Ako trebaš samo dio, pretplati se ponovo za drugi rok.\n\n<b>Želiš li stvarno zatražiti povrat?</b>",
		NoPayments:     "Nema pronađenih plaćanja",
		SuccessMessage: "Povrat je zatražen! Čekaj poruku od admina u privatno",
	},
	language.Slovenčina: {
		InfoText:       "<b>Ako funguje vrátenie prostriedkov?</b>\nProstriedky vracia admin pod príspevkom v tvojom Telegram kanáli v mene «Stars».\n\n<b>Ako rýchlo sa prostriedky vrátia?</b>\nVrátenie trvá od 30 minút do 23 dní.\n\n<b>Môžem vrátiť len časť?</b>\nVracia sa celý zvyšný zostatok. Ak potrebuješ len časť, prihlás sa znova na inú dobu.\n\n<b>Naozaj chceš požiadať o vrátenie?</b>",
		NoPayments:     "Žiadne platby nájdené",
		SuccessMessage: "Vrátenie požiadané! Čakaj na správu od admina v súkromí",
	},
	language.Slovenski: {
		InfoText:       "<b>Kako deluje vračilo sredstev?</b>\nSredstva vrne admin pod objavo v tvojem Telegram kanalu v valuti «Stars».\n\n<b>Kako hitro bodo sredstva vrnjena?</b>\nVračilo traja od 30 minut do 23 dni.\n\n<b>Ali lahko vrnem samo del?</b>\nVrne se celoten preostali znesek. Če potrebuješ le del, naroči se ponovno za drugo obdobje.\n\n<b>Ali res želiš zahtevati vračilo?</b>",
		NoPayments:     "Ni najdenih plačil",
		SuccessMessage: "Vračilo zahtevano! Počakaj na sporočilo od admina v zasebno",
	},
	language.Lietùvių: {
		InfoText:       "<b>Kaip veikia grąžinimas?</b>\nLėšos grąžinamos admino po įrašu tavo Telegram kanale valiuta «Stars».\n\n<b>Kaip greitai lėšos bus grąžintos?</b>\nGrąžinimas užtrunka nuo 30 minučių iki 23 dienų.\n\n<b>Ar galima grąžinti tik dalį?</b>\nGrąžinamas visas likęs balansas. Jei reikia tik dalies, prenumeruok vėl kitam laikotarpiui.\n\n<b>Ar tikrai nori prašyti grąžinimo?</b>",
		NoPayments:     "Mokėjimų nerasta",
		SuccessMessage: "Grąžinimas pateiktas! Laukite žinutės nuo admino asmeninėse",
	},
	language.Latviešu: {
		InfoText:       "<b>Kā darbojas atmaksa?</b>\nLīdzekļi tiek atgriezti no admina zem ieraksta tavā Telegram kanālā valūtā «Stars».\n\n<b>Cik ātri līdzekļi tiks atgriezti?</b>\nAtmaksa aizņem no 30 minūtēm līdz 23 dienām.\n\n<b>Vai var atgriezt tikai daļu?</b>\nTiek atgriezts viss atlikušais atlikums. Ja vajag tikai daļu, abonē atkal citam termiņam.\n\n<b>Vai tiešām vēlies pieprasīt atmaksu?</b>",
		NoPayments:     "Nav atrasti maksājumi",
		SuccessMessage: "Atmaksa pieprasīta! Gaidi ziņu no admina privāti",
	},
	language.Eesti: {
		InfoText:       "<b>Kuidas tagastamine töötab?</b>\nRahad tagastab admin postituse all sinu Telegrami kanalil valuutas «Stars».\n\n<b>Kui kiiresti rahad tagastatakse?</b>\nTagastamine võtab 30 minutist kuni 23 päevani.\n\n<b>Kas saab tagastada ainult osa?</b>\nTagastatakse kogu järelejäänud saldo. Kui vaja ainult osa, telli uuesti teisele perioodile.\n\n<b>Kas sa tõesti tahad tagastamist taotleda?</b>",
		NoPayments:     "Makseid ei leitud",
		SuccessMessage: "Tagastamine esitatud! Oota admini sõnumit privaatselt",
	},
	language.Suomi: {
		InfoText:       "<b>Miten palautus toimii?</b>\nRahat palauttaa admin postauksen alla Telegram-kanavallasi valuutassa «Stars».\n\n<b>Kuinka pian rahat palautetaan?</b>\nPalautus kestää 30 minuutista 23 päivään.\n\n<b>Voiko palauttaa vain osan?</b>\nPalautetaan koko jäljellä oleva saldo. Jos tarvit vain osan, tilaa uudelleen toiselle jaksolle.\n\n<b>Haluatko todella pyytää palautusta?</b>",
		NoPayments:     "Ei löytynyt maksuja",
		SuccessMessage: "Palautus pyydetty! Odota viestiä adminilta yksityisviestinä",
	},
	language.Ελληνικά: {
		InfoText:       "<b>Πώς λειτουργεί η επιστροφή;</b>\nΤα χρήματα επιστρέφονται από τον admin κάτω από την ανάρτηση στο κανάλι σου στο Telegram με νόμισμα «Stars».\n\n<b>Πόσο γρήγορα θα επιστραφούν τα χρήματα;</b>\nΗ επιστροφή διαρκεί από 30 λεπτά έως 23 ημέρες.\n\n<b>Μπορώ να επιστρέψω μόνο μέρος;</b>\nΕπιστρέφεται ολόκληρο το υπόλοιπο. Αν χρειάζεσαι μόνο μέρος, εγγράψου ξανά για άλλη διάρκεια.\n\n<b>Θες σίγουρα να ζητήσεις επιστροφή;</b>",
		NoPayments:     "Δεν βρέθηκαν πληρωμές",
		SuccessMessage: "Επιστροφή ζητήθηκε! Περίμενε μήνυμα από τον admin στα προσωπικά",
	},
	language.Română: {
		InfoText:       "<b>Cum funcționează rambursarea?</b>\nFondurile sunt returnate de admin sub postarea din canalul tău Telegram în moneda «Stars».\n\n<b>Cât de repede vin fondurile înapoi?</b>\nRambursarea durează de la 30 de minute la 23 de zile.\n\n<b>Pot returna doar o parte?</b>\nSe returnează întregul sold rămas. Dacă ai nevoie doar de o parte, abonează-te din nou pentru altă perioadă.\n\n<b>Vrei cu adevărat să ceri rambursare?</b>",
		NoPayments:     "Nu s-au găsit plăți",
		SuccessMessage: "Rambursare solicitată! Așteaptă mesajul adminului în privat",
	},
	language.Magyar: {
		InfoText:       "<b>Hogyan működik a visszatérítés?</b>\nA pénzt az admin téríti vissza a bejegyzés alatt a Telegram csatornádon «Stars» pénznemben.\n\n<b>Mennyi idő a visszatérítés?</b>\nA visszatérítés 30 perctől 23 napig tart.\n\n<b>Visszatéríthető csak egy rész?</b>\nA teljes fennmaradó egyenleg térül vissza. Ha csak rész kell, fizess elő újra más időtartamra.\n\n<b>Tényleg kéred a visszatérítést?</b>",
		NoPayments:     "Nem található fizetés",
		SuccessMessage: "Visszatérítés kérve! Várj az admin üzenetére privátban",
	},
	language.Arabic: {
		InfoText:       "<b>كيف يعمل استرداد الأموال؟</b>\nالأموال تسترد من قبل الإداري تحت المنشور في قناتك على تليجرام بعملة «Stars».\n\n<b>كم الوقت حتى يتم الاسترداد؟</b>\nاسترداد الأموال يأخذ من 30 دقيقة إلى 23 يوم.\n\n<b>هل يمكن استرداد جزء فقط؟</b>\nيسترد الرصيد المتبقي كله. إذا كنت تحتاج جزء فقط، اشترك مرة أخرى لمدة مختلفة.\n\n<b>هل تريد فعلا طلب الاسترداد؟</b>",
		NoPayments:     "لم يتم العثور على مدفوعات",
		SuccessMessage: "تم طلب الاسترداد! انتظر رسالة من الإدارة في الخاص",
	},
	language.Farsi: {
		InfoText:       "<b>بازگشت وجوه چطور کار می‌کنه؟</b>\nوجوه توسط مدیر زیر پست در کانال تلگرامتون با ارز «Stars» برگردونده می‌شه.\n\n<b>وجوه کی برمی‌گرده؟</b>\nبازگشت وجوه از ۳۰ دقیقه تا ۲۳ روز طول می‌کشه.\n\n<b>می‌شه فقط بخشی رو برگردوند؟</b>\nکل موجودی باقی‌مونده برگردونده می‌شه. اگه فقط بخشی می‌خوای، دوباره برای مدت دیگه مشترک شو.\n\n<b>واقعا می‌خوای درخواست بازگشت بدی؟</b>",
		NoPayments:     "پرداختی یافت نشد",
		SuccessMessage: "درخواست بازگشت ثبت شد! منتظر پیام ادمین در خصوصی باش",
	},
	language.Türkçe: {
		InfoText:       "<b>İade nasıl çalışır?</b>\nParalar admin tarafından Telegram kanalındaki postun altında «Stars» para birimiyle iade edilir.\n\n<b>Paralar ne zaman iade edilir?</b>\nİade 30 dakikadan 23 güne kadar sürer.\n\n<b>Sadece bir kısmını iade edebilir miyim?</b>\nTüm kalan bakiye iade edilir. Sadece bir kısma ihtiyacın varsa, farklı bir süre için yeniden abone ol.\n\n<b>Gerçekten iade talep etmek istiyor musun?</b>",
		NoPayments:     "Ödeme bulunamadı",
		SuccessMessage: "İade talebin alındı! Adminden özel mesaj bekle",
	},
	language.Hebrew: {
		InfoText:       "<b>איך עובד ההחזר?</b>\nהכספים מוחזרים על ידי האדמין מתחת לפוסט בערוץ הטלגרם שלך במטבע «Stars».\n\n<b>מתי הכספים יוחזרו?</b>\nההחזר לוקח מ-30 דקות עד 23 ימים.\n\n<b>אפשר להחזיר רק חלק?</b>\nכל היתרה הנותרת מוחזרת. אם צריך רק חלק, הרשם שוב לתקופה אחרת.\n\n<b>אתה בטוח רוצה לבקש החזר?</b>",
		NoPayments:     "לא נמצאו תשלומים",
		SuccessMessage: "ההחזר בטיפול! חכה להודעה מהאדמין בפרטי",
	},
	language.ZH中文: {
		InfoText:       "<b>退款怎么运作？</b>\n资金由管理员在你的Telegram频道帖子下以“Stars”货币退还。\n\n<b>资金什么时候退还？</b>\n退款需要30分钟到23天。\n\n<b>可以只退部分资金吗？</b>\n退还全部剩余余额。如果只需部分，再订阅其他期限。\n\n<b>你真的想请求退款吗？</b>",
		NoPayments:     "未找到付款",
		SuccessMessage: "退款已申请！请等待管理员私信你",
	},
	language.JA日本語: {
		InfoText:       "<b>返金はどう機能しますか？</b>\n資金は管理者によりTelegramチャンネルの投稿の下で「Stars」通貨で返金されます。\n\n<b>資金はいつ返金されますか？</b>\n返金は30分から23日かかります。\n\n<b>資金の一部だけ返金できますか？</b>\n残りの残高全部が返金されます。一部だけ必要な場合、別の期間で再度購読してください。\n\n<b>本当に返金をリクエストしますか？</b>",
		NoPayments:     "支払いが見つかりません",
		SuccessMessage: "返金申請完了！管理人からDMで連絡待ってね",
	},
	language.KO한국어: {
		InfoText:       "<b>환불은 어떻게 작동하나요?</b>\n자금은 관리자가 Telegram 채널 게시물 아래에서 «Stars» 통화로 환불합니다.\n\n<b>자금은 언제 환불되나요?</b>\n환불은 30분에서 23일까지 걸립니다。\n\n<b>자금 일부만 환불할 수 있나요?</b>\n남은 잔액 전체가 환불됩니다. 일부만 필요하다면 다른 기간으로 다시 구독하세요。\n\n<b>정말 환불 요청하시겠어요?</b>",
		NoPayments:     "지불 내역이 없습니다",
		SuccessMessage: "환불 신청 완료! 관리자로부터 DM을 기다려 주세요",
	},
	language.TiếngViệt: {
		InfoText:       "<b>Hoàn tiền hoạt động thế nào?</b>\nTiền được hoàn lại bởi admin dưới bài đăng trong kênh Telegram của bạn bằng tiền tệ «Stars».\n\n<b>Tiền sẽ được hoàn lại sau bao lâu?</b>\nHoàn tiền mất từ 30 phút đến 23 ngày。\n\n<b>Có thể hoàn lại chỉ một phần không?</b>\nHoàn lại toàn bộ số dư còn lại. Nếu cần chỉ phần, đăng ký lại cho khoảng thời gian khác。\n\n<b>Bạn có thực sự muốn yêu cầu hoàn tiền không?</b>",
		NoPayments:     "Không tìm thấy thanh toán",
		SuccessMessage: "Đã gửi yêu cầu hoàn tiền! Chờ tin nhắn từ admin trong tin nhắn riêng",
	},
	language.THภาษาไทย: {
		InfoText:       "<b>การคืนเงินทำงานยังไง?</b>\nเงินถูกคืนโดยแอดมินใต้โพสต์ในช่อง Telegram ของคุณในสกุลเงิน «Stars».\n\n<b>เงินจะถูกคืนเร็วแค่ไหน?</b>\nการคืนเงินใช้เวลาตั้งแต่ 30 นาทีถึง 23 วัน。\n\n<b>คืนเงินบางส่วนได้มั้ย?</b>\nคืนยอดคงเหลือทั้งหมด ถ้าต้องการแค่บางส่วน สมัครสมาชิกใหม่สำหรับระยะเวลาอื่น。\n\n<b>คุณแน่ใจว่าต้องการขอคืนเงินเหรอ?</b>",
		NoPayments:     "ไม่พบการชำระเงิน",
		SuccessMessage: "ขอคืนเงินเรียบร้อย! รอแอดมินทักในแชทส่วนตัว",
	},
	language.BahasaIndonesia: {
		InfoText:       "<b>Bagaimana cara kerja pengembalian dana?</b>\nDana dikembalikan oleh admin di bawah posting di kanal Telegram kamu dalam mata uang «Stars».\n\n<b>Seberapa cepat dana dikembalikan?</b>\nPengembalian dana memakan waktu dari 30 menit hingga 23 hari。\n\n<b>Bisa kembalikan sebagian dana saja?</b>\nSeluruh saldo tersisa dikembalikan. Kalau butuh cuma sebagian, berlangganan lagi buat periode lain。\n\n<b>Kamu yakin mau minta pengembalian dana?</b>",
		NoPayments:     "Tidak ada pembayaran ditemukan",
		SuccessMessage: "Refund sudah diajukan! Tunggu pesan dari admin di DM",
	},
	language.BahasaMelayu: {
		InfoText:       "<b>Bagaimana refund berfungsi?</b>\nDana dikembalikan oleh admin di bawah pos dalam saluran Telegram anda dalam mata wang «Stars».\n\n<b>Berapa lama dana akan dikembalikan?</b>\nRefund mengambil masa dari 30 minit hingga 23 hari。\n\n<b>Boleh kembalikan sebahagian dana sahaja?</b>\nKeseluruhan baki yang tinggal dikembalikan. Kalau perlu sebahagian sahaja, langgan semula untuk tempoh lain。\n\n<b>Anda pasti mahu minta refund?</b>",
		NoPayments:     "Tiada pembayaran ditemui",
		SuccessMessage: "Refund sudah dimohon! Tunggu mesej dari admin di DM",
	},
	language.Tagalog: {
		InfoText:       "<b>Paano gumagana ang refund?</b>\nAng pera ibinabalik ng admin sa ilalim ng post sa iyong Telegram channel sa currency na «Stars».\n\n<b>Gaano kabilis ibabalik ang pera?</b>\nAng refund tumatagal mula 30 minuto hanggang 23 araw。\n\n<b>Maaari bang ibalik lang bahagi ng pera?</b>\nIbinabalik ang buong natitirang balanse. Kung kailangan mo lang bahagi, mag-subscribe ulit para sa ibang termino。\n\n<b>Sigurado ka bang gustong humiling ng refund?</b>",
		NoPayments:     "Walang natagpuang bayad",
		SuccessMessage: "Na-request na ang refund! Hintayin ang mensahe ng admin sa DM",
	},
	language.Hindi: {
		InfoText:       "<b>रिफंड कैसे काम करता है?</b>\nफंड्स एडमिन द्वारा आपके टेलीग्राम चैनल में पोस्ट के नीचे «Stars» करंसी में वापस किए जाते हैं।\n\n<b>फंड्स कितनी जल्दी वापस मिलेंगे?</b>\nरिफंड 30 मिनट से 23 दिनों तक लगता है。\n\n<b>क्या फंड्स का सिर्फ हिस्सा वापस किया जा सकता है?</b>\nपूरी बाकी बैलेंस वापस की जाती है। अगर सिर्फ हिस्सा चाहिए, तो दूसरे टर्म के लिए फिर सब्सक्राइब करो。\n\n<b>क्या तुम सच में रिफंड रिक्वेस्ट करना चाहते हो?</b>",
		NoPayments:     "कोई भुगतान नहीं मिला",
		SuccessMessage: "रिफंड रिक्वेस्ट हो गई! एडमिन की प्राइवेट मैसेज का इंतज़ार करो",
	},
	language.URاردو: {
		InfoText:       "<b>ریفنڈ کیسے کام کرتا ہے؟</b>\nفنڈز ایڈمن کے ذریعے آپ کے ٹیلی گرام چینل میں پوسٹ کے نیچے «Stars» کرنسی میں واپس کیے جاتے ہیں۔\n\n<b>فنڈز کتنی جلدی واپس ملेंगे؟</b>\nریفنڈ 30 منٹ سے 23 دن تک لگتا ہے。\n\n<b>کیا فنڈز کا صرف حصہ واپس کیا جا سکتا ہے؟</b>\nپوری باقی بیلنس واپس کی جاتی ہے۔ اگر صرف حصہ چاہیے تو دوسرے ٹرم کے لیے پھر سبسکرائب کرو。\n\n<b>کیا تم واقعی ریفنڈ کی درخواست کرنا چاہتے ہو؟</b>",
		NoPayments:     "کوئی ادائیگی نہیں ملی",
		SuccessMessage: "ریفنڈ کی درخواست ہو گئی! ایڈمن کی پرائیویٹ میں میسج کا انتظار کرو",
	},
	language.Bengali: {
		InfoText:       "<b>রিফান্ড কীভাবে কাজ করে?</b>\nতহবিলগুলি অ্যাডমিন দ্বারা আপনার টেলিগ্রাম চ্যানেলে পোস্টের নীচে «Stars» মুদ্রায় ফেরত দেওয়া হয়।\n\n<b>তহবিলগুলি কত তাড়াতাড়ি ফেরত দেওয়া হবে?</b>\nরিফান্ড ৩০ মিনিট থেকে ২৩ দিন লাগে。\n\n<b>তহবিলের শুধু অংশ ফেরত দেওয়া যায়?</b>\nসম্পূর্ণ অবশিষ্ট ব্যালেন্স ফেরত দেওয়া হয়। যদি শুধু অংশ চাই, অন্য টার্মের জন্য আবার সাবস্ক্রাইব করো。\n\n<b>তুমি কি সত্যিই রিফান্ড রিকোয়েস্ট করতে চাও?</b>",
		NoPayments:     "কোনো পেমেন্ট পাওয়া যায়নি",
		SuccessMessage: "রিফান্ড রিকোয়েস্ট হয়ে গেছে! অ্যাডমিনের প্রাইভেট মেসেজের জন্য অপেক্ষা করো",
	},
	language.Tamiḻ: {
		InfoText:       "<b>பணத்திருப்பு எப்படி வேலை செய்கிறது?</b>\nபணங்கள் நிர்வாகியால் உங்கள் டெலிகிராம் சேனலில் பதிவின் கீழ் «Stars» நாணயத்தில் திருப்பி அளிக்கப்படும்。\n\n<b>பணங்கள் எப்போது திருப்பி அளிக்கப்படும்?</b>\nபணத்திருப்பு 30 நிமிடங்கள் முதல் 23 நாட்கள் வரை எடுக்கும்。\n\n<b>பணங்களின் ஒரு பகுதி மட்டும் திருப்பி அளிக்க முடியுமா?</b>\nமுழு எஞ்சிய இருப்பு திருப்பி அளிக்கப்படும். ஒரு பகுதி மட்டும் தேவைப்பட்டால், வேறு காலத்துக்கு மீண்டும் சந்தா செலுத்து。\n\n<b>நீ உண்மையில் பணத்திருப்பு கோர விரும்புகிறாயா?</b>",
		NoPayments:     "பணம் செலுத்தல்கள் கிடைக்கவில்லை",
		SuccessMessage: "பணத்திருப்பு கோரிக்கை செய்யப்பட்டது! நிர்வாகியின் தனிப்பட்ட செய்தியை எதிர்பார்க்கவும்",
	},
	language.Telugu: {
		InfoText:       "<b>రీఫండ్ ఎలా పని చేస్తుంది?</b>\nనిధులు అడ్మిన్ ద్వారా మీ టెలిగ్రామ్ ఛానల్‌లో పోస్ట్ కింద «Stars» కరెన్సీలో తిరిగి ఇవ్వబడతాయి。\n\n<b>నిధులు ఎప్పుడు తిరిగి ఇవ్వబడతాయి?</b>\nరీఫండ్ 30 నిమిషాల నుండి 23 రోజుల వరకు పడుతుంది。\n\n<b>నిధులలో ఒక భాగం మాత్రమే తిరిగి ఇవ్వవచ్చా?</b>\nమొత్తం మిగిలిన బ్యాలెన్స్ తిరిగి ఇవ్వబడుతుంది. ఒక భాగం మాత్రమే అవసరమైతే, మరొక కాలానికి మళ్లీ సబ్‌స్క్రైబ్ చేయి。\n\n<b>మీరు నిజంగా రీఫండ్ అభ్యర్థించాలనుకుంటున్నారా?</b>",
		NoPayments:     "చెల్లింపులు కనుగొనబడలేదు",
		SuccessMessage: "రీఫండ్ రిక్వెస్ట్ అయింది! అడ్మిన్ నుంచి ప్రైవేట్‌లో మెసేజ్ కోసం వేచి ఉండు",
	},
	language.Marathi: {
		InfoText:       "<b>परतावा कसा काम करतो?</b>\nनिधी अॅडमिनद्वारे तुमच्या टेलिग्राम चॅनेलमध्ये पोस्टखाली «Stars» चलनात परत केले जातात。\n\n<b>निधी कधी परत मिळतील?</b>\nपरतावा ३० मिनिटांपासून २३ दिवसांपर्यंत घेतो。\n\n<b>निधीचा फक्त भाग परत करता येईल का?</b>\nसंपूर्ण उर्वरित शिल्लक परत केली जाते. फक्त भाग हवा असल्यास दुसऱ्या मुदतीसाठी पुन्हा सबस्क्राईब करा。\n\n<b>तुम्हाला खरंच परतावा विनंती करायचा आहे का?</b>",
		NoPayments:     "कोणतेही पेमेंट सापडले नाही",
		SuccessMessage: "रिफंड रिक्वेस्ट झाली! अॅडमिनकडून प्रायव्हेट मध्ये मेसेजची वाट पहा",
	},
}

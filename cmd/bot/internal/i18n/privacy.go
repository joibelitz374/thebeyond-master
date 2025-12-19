package i18n

import "github.com/quickpowered/thebeyond-master/configs/language"

type PrivacyPolicyLocale struct {
	Policy string
}

var PrivacyPolicyMessages = map[language.Language]PrivacyPolicyLocale{
	language.English: {
		Policy: "<b>PRIVACY POLICY:</b>\n\n" +
			"<b>1. What data we collect</b>\n" +
			"For the operation of THE BEYOND service, we collect the minimum necessary set of data:\n" +
			"• Telegram User ID: For user identification in the system and automatic key issuance.\n" +
			"• Payment metadata: Amount, time, and payment method (we do not store full bank card details; processing occurs on the payment gateway side).\n" +
			"• Technical statistics: Volume of consumed traffic (in bytes) for subscription limit control.\n\n" +
			"<b>2. No-Logs Policy</b>\n" +
			"2.1. We adhere to a strict privacy policy. We DO NOT collect, track, or store:\n" +
			"• History of visited sites (DNS queries).\n" +
			"• Content of transmitted traffic.\n" +
			"• Real IP addresses of users when connecting to the server (connection logs are deleted cyclically).\n\n" +
			"<b>3. Use of data</b>\n" +
			"3.1. Collected data is used exclusively for:\n" +
			"• Providing access to the service (authorization).\n" +
			"• Technical support for users.\n\n" +
			"<b>4. Transfer of data to third parties</b>\n" +
			"4.1. We do not sell or transfer your personal data to third parties, except in cases provided by law or for processing payments through secure gateways.",
	},
	language.Deutsch: {
		Policy: "<b>DATENSCHUTZERKLÄRUNG:</b>\n\n" +
			"<b>1. Welche Daten wir sammeln</b>\n" +
			"Für den Betrieb des Dienstes THE BEYOND sammeln wir das minimal notwendige Datenset:\n" +
			"• Telegram User ID: Zur Identifikation des Benutzers im System und zur automatischen Ausgabe von Schlüsseln.\n" +
			"• Zahlungsmetadaten: Betrag, Zeit und Zahlungsmethode (wir speichern keine vollständigen Bankkartendaten; die Verarbeitung erfolgt auf Seiten des Zahlungsgateways).\n" +
			"• Technische Statistiken: Volumen des verbrauchten Traffics (in Bytes) zur Kontrolle der Abonnementlimits.\n\n" +
			"<b>2. No-Logs-Richtlinie</b>\n" +
			"2.1. Wir halten uns an eine strenge Datenschutzrichtlinie. Wir sammeln NICHT, verfolgen nicht und speichern nicht:\n" +
			"• Verlauf der besuchten Websites (DNS-Anfragen).\n" +
			"• Inhalt des übertragenen Traffics.\n" +
			"• Echte IP-Adressen der Benutzer bei der Verbindung zum Server (Verbindungslogs werden zyklisch gelöscht).\n\n" +
			"<b>3. Nutzung der Daten</b>\n" +
			"3.1. Die gesammelten Daten werden ausschließlich verwendet für:\n" +
			"• Bereitstellung des Zugriffs auf den Dienst (Autorisierung).\n" +
			"• Technische Unterstützung der Benutzer.\n\n" +
			"<b>4. Weitergabe der Daten an Dritte</b>\n" +
			"4.1. Wir verkaufen oder übermitteln Ihre personenbezogenen Daten nicht an Dritte, außer in gesetzlich vorgesehenen Fällen oder zur Verarbeitung von Zahlungen über sichere Gateways.",
	},
	language.Nederlands: {
		Policy: "<b>PRIVACYBELEID:</b>\n\n" +
			"<b>1. Welke gegevens we verzamelen</b>\n" +
			"Voor de werking van de dienst THE BEYOND verzamelen we de minimaal noodzakelijke set gegevens:\n" +
			"• Telegram User ID: Voor identificatie van de gebruiker in het systeem en automatische uitgifte van sleutels.\n" +
			"• Betalingsmetadata: Bedrag, tijd en betalingsmethode (we slaan geen volledige bankkaartgegevens op; de verwerking vindt plaats aan de kant van de betalingsgateway).\n" +
			"• Technische statistieken: Volume van verbruikt verkeer (in bytes) voor controle van abonnementslimieten.\n\n" +
			"<b>2. No-Logs-Beleid</b>\n" +
			"2.1. We houden ons aan een streng privacybeleid. We verzamelen NIET, volgen niet en slaan niet op:\n" +
			"• Geschiedenis van bezochte sites (DNS-queries).\n" +
			"• Inhoud van overgedragen verkeer.\n" +
			"• Echte IP-adressen van gebruikers bij verbinding met de server (verbindingslogs worden cyclisch verwijderd).\n\n" +
			"<b>3. Gebruik van gegevens</b>\n" +
			"3.1. De verzamelde gegevens worden uitsluitend gebruikt voor:\n" +
			"• Verstrekken van toegang tot de dienst (autorisatie).\n" +
			"• Technische ondersteuning voor gebruikers.\n\n" +
			"<b>4. Overdracht van gegevens aan derden</b>\n" +
			"4.1. We verkopen of overdragen uw persoonlijke gegevens niet aan derden, behalve in gevallen voorzien door de wet of voor verwerking van betalingen via veilige gateways.",
	},
	language.Svenska: {
		Policy: "<b>INTEGRITETSPOLICY:</b>\n\n" +
			"<b>1. Vilka data vi samlar in</b>\n" +
			"För driften av tjänsten THE BEYOND samlar vi in det minimalt nödvändiga setet av data:\n" +
			"• Telegram User ID: För identifiering av användaren i systemet och automatisk utdelning av nycklar.\n" +
			"• Betalningsmetadata: Belopp, tid och betalningsmetod (vi lagrar inte fullständiga bankkortsuppgifter; bearbetningen sker på betalningsgatewayens sida).\n" +
			"• Teknisk statistik: Volym av förbrukad trafik (i byte) för kontroll av abonnemangets gränser.\n\n" +
			"<b>2. No-Logs-Policy</b>\n" +
			"2.1. Vi följer en strikt integritetspolicy. Vi samlar INTE in, spårar inte och lagrar inte:\n" +
			"• Historik över besökta webbplatser (DNS-frågor).\n" +
			"• Innehåll i överförd trafik.\n" +
			"• Verkliga IP-adresser för användare vid anslutning till servern (anslutningsloggar raderas cykliskt).\n\n" +
			"<b>3. Användning av data</b>\n" +
			"3.1. Insamlade data används uteslutande för:\n" +
			"• Tillhandahållande av åtkomst till tjänsten (autentisering).\n" +
			"• Teknisk support för användare.\n\n" +
			"<b>4. Överföring av data till tredje part</b>\n" +
			"4.1. Vi säljer eller överför inte dina personuppgifter till tredje part, utom i fall som föreskrivs i lag eller för bearbetning av betalningar via säkra gateways.",
	},
	language.Norsk: {
		Policy: "<b>PERSONVERNERKLÆRING:</b>\n\n" +
			"<b>1. Hvilke data vi samler inn</b>\n" +
			"For driften av tjenesten THE BEYOND samler vi inn det minimalt nødvendige settet med data:\n" +
			"• Telegram User ID: For identifikasjon av brukeren i systemet og automatisk utstedelse av nøkler.\n" +
			"• Betalingsmetadata: Beløp, tid og betalingsmåte (vi lagrer ikke fullstendige bankkortdetaljer; behandlingen skjer på betalingsgatewayens side).\n" +
			"• Teknisk statistikk: Volum av forbrukt trafikk (i byte) for kontroll av abonnementsgrenser.\n\n" +
			"<b>2. No-Logs-Policy</b>\n" +
			"2.1. Vi følger en streng personvernpolicy. Vi samler IKKE inn, sporer ikke og lagrer ikke:\n" +
			"• Historikk over besøkte nettsteder (DNS-spørringer).\n" +
			"• Innhold i overført trafikk.\n" +
			"• Virkelige IP-adresser for brukere ved tilkobling til serveren (tilkoblingslogger slettes syklisk).\n\n" +
			"<b>3. Bruk av data</b>\n" +
			"3.1. Innsamlede data brukes utelukkende til:\n" +
			"• Tilby tilgang til tjenesten (autorisering).\n" +
			"• Teknisk støtte for brukere.\n\n" +
			"<b>4. Overføring av data til tredjeparter</b>\n" +
			"4.1. Vi selger eller overfører ikke dine personopplysninger til tredjeparter, unntatt i tilfeller foreskrevet av loven eller for behandling av betalinger gjennom sikre gateways.",
	},
	language.Dansk: {
		Policy: "<b>DATAPOLITIK:</b>\n\n" +
			"<b>1. Hvilke data vi indsamler</b>\n" +
			"For funktionen af tjenesten THE BEYOND indsamler vi det minimalt nødvendige sæt data:\n" +
			"• Telegram User ID: Til identifikation af brugeren i systemet og automatisk udstedelse af nøgler.\n" +
			"• Betalingsmetadata: Beløb, tid og betalingsmetode (vi gemmer ikke fuldstændige bankkortoplysninger; behandlingen sker på betalingsgatewayens side).\n" +
			"• Teknisk statistik: Volumen af forbrugt trafik (i bytes) til kontrol af abonnementsgrænser.\n\n" +
			"<b>2. No-Logs-Politik</b>\n" +
			"2.1. Vi overholder en streng datapolitik. Vi indsamler IKKE, sporer ikke og gemmer ikke:\n" +
			"• Historik over besøgte websteder (DNS-forespørgsler).\n" +
			"• Indhold af overført trafik.\n" +
			"• Rigtige IP-adresser for brugere ved tilslutning til serveren (forbindelseslogs slettes cyklisk).\n\n" +
			"<b>3. Brug af data</b>\n" +
			"3.1. Indsamlede data bruges udelukkende til:\n" +
			"• Tilvejebringelse af adgang til tjenesten (autorisation).\n" +
			"• Teknisk support til brugere.\n\n" +
			"<b>4. Overførsel af data til tredjeparter</b>\n" +
			"4.1. Vi sælger eller overfører ikke dine personlige data til tredjeparter, undtagen i tilfælde, der er foreskrevet i lovgivningen, eller til behandling af betalinger gennem sikre gateways.",
	},
	language.Español: {
		Policy: "<b>POLÍTICA DE PRIVACIDAD:</b>\n\n" +
			"<b>1. Qué datos recopilamos</b>\n" +
			"Para el funcionamiento del servicio THE BEYOND, recopilamos el conjunto mínimo necesario de datos:\n" +
			"• Telegram User ID: Para la identificación del usuario en el sistema y la emisión automática de claves.\n" +
			"• Metadatos de pagos: Importe, hora y método de pago (no almacenamos datos completos de tarjetas bancarias; el procesamiento se realiza en el lado de la pasarela de pagos).\n" +
			"• Estadísticas técnicas: Volumen de tráfico consumido (en bytes) para el control de límites de suscripción.\n\n" +
			"<b>2. Política de No-Logs</b>\n" +
			"2.1. Nos adherimos a una estricta política de privacidad. NO recopilamos, rastreamos ni almacenamos:\n" +
			"• Historial de sitios visitados (consultas DNS).\n" +
			"• Contenido del tráfico transmitido.\n" +
			"• Direcciones IP reales de usuarios al conectar al servidor (los logs de conexión se eliminan cíclicamente).\n\n" +
			"<b>3. Uso de los datos</b>\n" +
			"3.1. Los datos recopilados se utilizan exclusivamente para:\n" +
			"• Proporcionar acceso al servicio (autorización).\n" +
			"• Soporte técnico a los usuarios.\n\n" +
			"<b>4. Transferencia de datos a terceros</b>\n" +
			"4.1. No vendemos ni transferimos sus datos personales a terceros, excepto en casos previstos por la legislación o para el procesamiento de pagos a través de pasarelas seguras.",
	},
	language.Français: {
		Policy: "<b>POLITIQUE DE CONFIDENTIALITÉ :</b>\n\n" +
			"<b>1. Quelles données nous collectons</b>\n" +
			"Pour le fonctionnement du service THE BEYOND, nous collectons l'ensemble minimal nécessaire de données :\n" +
			"• Telegram User ID : Pour l'identification de l'utilisateur dans le système et la délivrance automatique de clés.\n" +
			"• Métadonnées de paiements : Montant, heure et mode de paiement (nous ne stockons pas les données complètes des cartes bancaires ; le traitement se fait du côté de la passerelle de paiement).\n" +
			"• Statistiques techniques : Volume de trafic consommé (en octets) pour le contrôle des limites d'abonnement.\n\n" +
			"<b>2. Politique No-Logs</b>\n" +
			"2.1. Nous adhérons à une politique stricte de confidentialité. Nous NE collectons pas, ne suivons pas et ne stockons pas :\n" +
			"• Historique des sites visités (requêtes DNS).\n" +
			"• Contenu du trafic transmis.\n" +
			"• Adresses IP réelles des utilisateurs lors de la connexion au serveur (les logs de connexion sont supprimés de manière cyclique).\n\n" +
			"<b>3. Utilisation des données</b>\n" +
			"3.1. Les données collectées sont utilisées exclusivement pour :\n" +
			"• Fournir l'accès au service (autorisation).\n" +
			"• Support technique aux utilisateurs.\n\n" +
			"<b>4. Transmission des données à des tiers</b>\n" +
			"4.1. Nous ne vendons ni ne transmettons vos données personnelles à des tiers, sauf dans les cas prévus par la législation ou pour le traitement des paiements via des passerelles sécurisées.",
	},
	language.Português: {
		Policy: "<b>POLÍTICA DE PRIVACIDADE:</b>\n\n" +
			"<b>1. Quais dados coletamos</b>\n" +
			"Para o funcionamento do serviço THE BEYOND, coletamos o conjunto mínimo necessário de dados:\n" +
			"• Telegram User ID: Para identificação do usuário no sistema e emissão automática de chaves.\n" +
			"• Metadados de pagamentos: Valor, hora e método de pagamento (não armazenamos dados completos de cartões bancários; o processamento ocorre do lado da gateway de pagamento).\n" +
			"• Estatísticas técnicas: Volume de tráfego consumido (em bytes) para controle de limites de assinatura.\n\n" +
			"<b>2. Política de No-Logs</b>\n" +
			"2.1. Aderimos a uma política estrita de privacidade. NÃO coletamos, rastreamos ou armazenamos:\n" +
			"• Histórico de sites visitados (consultas DNS).\n" +
			"• Conteúdo do tráfego transmitido.\n" +
			"• Endereços IP reais dos usuários ao conectar ao servidor (logs de conexão são excluídos ciclicamente).\n\n" +
			"<b>3. Uso dos dados</b>\n" +
			"3.1. Os dados coletados são usados exclusivamente para:\n" +
			"• Fornecer acesso ao serviço (autorização).\n" +
			"• Suporte técnico aos usuários.\n\n" +
			"<b>4. Transferência de dados a terceiros</b>\n" +
			"4.1. Não vendemos nem transferimos seus dados pessoais a terceiros, exceto nos casos previstos pela legislação ou para processamento de pagamentos através de gateways seguras.",
	},
	language.Italiana: {
		Policy: "<b>POLITICA SULLA PRIVACY:</b>\n\n" +
			"<b>1. Quali dati raccogliamo</b>\n" +
			"Per il funzionamento del servizio THE BEYOND raccogliamo il set minimo necessario di dati:\n" +
			"• Telegram User ID: Per l'identificazione dell'utente nel sistema e il rilascio automatico di chiavi.\n" +
			"• Metadati dei pagamenti: Importo, ora e metodo di pagamento (non memorizziamo i dati completi delle carte bancarie; l'elaborazione avviene dal lato del gateway di pagamento).\n" +
			"• Statistiche tecniche: Volume del traffico consumato (in byte) per il controllo dei limiti di abbonamento.\n\n" +
			"<b>2. Politica No-Logs</b>\n" +
			"2.1. Aderiamo a una rigorosa politica sulla privacy. NON raccogliamo, tracciamo o memorizziamo:\n" +
			"• Storia dei siti visitati (query DNS).\n" +
			"• Contenuto del traffico trasmesso.\n" +
			"• Indirizzi IP reali degli utenti durante la connessione al server (i log di connessione vengono eliminati ciclicamente).\n\n" +
			"<b>3. Utilizzo dei dati</b>\n" +
			"3.1. I dati raccolti vengono utilizzati esclusivamente per:\n" +
			"• Fornire l'accesso al servizio (autorizzazione).\n" +
			"• Supporto tecnico agli utenti.\n\n" +
			"<b>4. Trasmissione dei dati a terzi</b>\n" +
			"4.1. Non vendiamo né trasmettiamo i vostri dati personali a terzi, tranne nei casi previsti dalla legislazione o per l'elaborazione dei pagamenti tramite gateway sicuri.",
	},
	language.Русский: {
		Policy: "<b>ПОЛИТИКА КОНФИДЕНЦИАЛЬНОСТИ:</b>\n\n" +
			"<b>1. Какие данные мы собираем</b>\n" +
			"Для функционирования сервиса THE BEYOND мы собираем минимально необходимый набор данных:\n" +
			"• Telegram User ID: Для идентификации пользователя в системе и автоматической выдачи ключей.\n" +
			"• Метаданные платежей: Сумма, время и способ оплаты (мы не храним полные данные банковских карт, обработка происходит на стороне платежного шлюза).\n" +
			"• Техническая статистика: Объем потребленного трафика (в байтах) для контроля лимитов подписки.\n\n" +
			"<b>2. Политика отсутствия логов (No-Logs Policy)</b>\n" +
			"2.1. Мы придерживаемся строгой политики конфиденциальности. Мы НЕ собираем, не отслеживаем и не храним:\n" +
			"• Историю посещенных сайтов (DNS-запросы).\n" +
			"• Содержимое передаваемого трафика.\n" +
			"• Реальные IP-адреса пользователей при подключении к серверу (логи подключения удаляются циклически).\n\n" +
			"<b>3. Использование данных</b>\n" +
			"3.1. Собранные данные используются исключительно для:\n" +
			"• Предоставления доступа к сервису (авторизация).\n" +
			"• Технической поддержки пользователей.\n\n" +
			"<b>4. Передача данных третьим лицам</b>\n" +
			"4.1. Мы не продаем и не передаем ваши личные данные третьим лицам, за исключением случаев, предусмотренных законодательством, или для обработки платежей через безопасные шлюзы.",
	},
	language.Українська: {
		Policy: "<b>ПОЛІТИКА КОНФІДЕНЦІЙНОСТІ:</b>\n\n" +
			"<b>1. Які дані ми збираємо</b>\n" +
			"Для функціонування сервісу THE BEYOND ми збираємо мінімально необхідний набір даних:\n" +
			"• Telegram User ID: Для ідентифікації користувача в системі та автоматичної видачі ключів.\n" +
			"• Метадані платежів: Сума, час та спосіб оплати (ми не зберігаємо повні дані банківських карт, обробка відбувається на стороні платіжного шлюзу).\n" +
			"• Технічна статистика: Обсяг спожитого трафіку (в байтах) для контролю лімітів підписки.\n\n" +
			"<b>2. Політика відсутності логів (No-Logs Policy)</b>\n" +
			"2.1. Ми дотримуємося суворої політики конфіденційності. Ми НЕ збираємо, не відстежуємо та не зберігаємо:\n" +
			"• Історію відвіданих сайтів (DNS-запити).\n" +
			"• Вміст переданого трафіку.\n" +
			"• Реальні IP-адреси користувачів при підключенні до сервера (логи підключення видаляються циклічно).\n\n" +
			"<b>3. Використання даних</b>\n" +
			"3.1. Зібрані дані використовуються виключно для:\n" +
			"• Надання доступу до сервісу (авторизація).\n" +
			"• Технічної підтримки користувачів.\n\n" +
			"<b>4. Передача даних третім особам</b>\n" +
			"4.1. Ми не продаємо та не передаємо ваші особисті дані третім особам, за винятком випадків, передбачених законодавством, або для обробки платежів через безпечні шлюзи.",
	},
	language.Polski: {
		Policy: "<b>POLITYKA PRYWATNOŚCI:</b>\n\n" +
			"<b>1. Jakie dane zbieramy</b>\n" +
			"Dla funkcjonowania usługi THE BEYOND zbieramy minimalnie niezbędny zestaw danych:\n" +
			"• Telegram User ID: Do identyfikacji użytkownika w systemie i automatycznego wydawania kluczy.\n" +
			"• Metadane płatności: Kwota, czas i sposób płatności (nie przechowujemy pełnych danych kart bankowych; przetwarzanie odbywa się po stronie bramki płatniczej).\n" +
			"• Statystyki techniczne: Objętość zużytego ruchu (w bajtach) do kontroli limitów subskrypcji.\n\n" +
			"<b>2. Polityka braku logów (No-Logs Policy)</b>\n" +
			"2.1. Przestrzegamy surowej polityki prywatności. NIE zbieramy, nie śledzimy i nie przechowujemy:\n" +
			"• Historii odwiedzanych stron (zapytania DNS).\n" +
			"• Zawartości przekazywanego ruchu.\n" +
			"• Rzeczywistych adresów IP użytkowników podczas połączenia z serwerem (logi połączeń są usuwane cyklicznie).\n\n" +
			"<b>3. Wykorzystanie danych</b>\n" +
			"3.1. Zebrane dane są wykorzystywane wyłącznie do:\n" +
			"• Udostępniania dostępu do usługi (autoryzacja).\n" +
			"• Wsparcia technicznego użytkowników.\n\n" +
			"<b>4. Przekazywanie danych osobom trzecim</b>\n" +
			"4.1. Nie sprzedajemy ani nie przekazujemy Twoich danych osobowych osobom trzecim, z wyjątkiem przypadków przewidzianych prawem lub do przetwarzania płatności poprzez bezpieczne bramki.",
	},
	language.Ceština: {
		Policy: "<b>ZÁSADY OCHRANY OSOBNÍCH ÚDAJŮ:</b>\n\n" +
			"<b>1. Jaké údaje shromažďujeme</b>\n" +
			"Pro fungování služby THE BEYOND shromažďujeme minimálně nezbytný soubor údajů:\n" +
			"• Telegram User ID: Pro identifikaci uživatele v systému a automatické vydávání klíčů.\n" +
			"• Metadata plateb: Částka, čas a způsob platby (neukládáme úplné údaje bankovních karet; zpracování probíhá na straně platební brány).\n" +
			"• Technické statistiky: Objem spotřebovaného provozu (v bajtech) pro kontrolu limitů předplatného.\n\n" +
			"<b>2. Zásady No-Logs</b>\n" +
			"2.1. Dodržujeme přísné zásady ochrany osobních údajů. NESHROMAŽĎUJEME, nesledujeme ani neukládáme:\n" +
			"• Historii navštívených stránek (DNS dotazy).\n" +
			"• Obsah přenášeného provozu.\n" +
			"• Skutečné IP adresy uživatelů při připojení k serveru (logy připojení jsou mazány cyklicky).\n\n" +
			"<b>3. Využití údajů</b>\n" +
			"3.1. Shromážděné údaje jsou používány výhradně pro:\n" +
			"• Poskytnutí přístupu ke službě (autorizace).\n" +
			"• Technickou podporu uživatelů.\n\n" +
			"<b>4. Předání údajů třetím stranám</b>\n" +
			"4.1. Neprodáváme ani nepředáváme vaše osobní údaje třetím stranám, s výjimkou případů stanovených zákonem nebo pro zpracování plateb prostřednictvím bezpečných bran.",
	},
	language.Български: {
		Policy: "<b>ПОЛИТИКА ЗА ПОВЕРИТЕЛНОСТ:</b>\n\n" +
			"<b>1. Какви данни събираме</b>\n" +
			"За функционирането на услугата THE BEYOND събираме минимално необходимия набор от данни:\n" +
			"• Telegram User ID: За идентификация на потребителя в системата и автоматично издаване на ключове.\n" +
			"• Метаданни на плащанията: Сума, време и начин на плащане (ние не съхраняваме пълни данни на банкови карти, обработката се извършва от страна на платежния шлюз).\n" +
			"• Техническа статистика: Обем на потребения трафик (в байтове) за контрол на лимитите на абонамента.\n\n" +
			"<b>2. Политика за липса на логове (No-Logs Policy)</b>\n" +
			"2.1. Ние се придържаме към строга политика за поверителност. Ние НЕ събираме, не проследяваме и не съхраняваме:\n" +
			"• История на посетени сайтове (DNS-заявки).\n" +
			"• Съдържание на предавания трафик.\n" +
			"• Реални IP адреси на потребителите при свързване към сървъра (логът на връзките се изтрива циклично).\n\n" +
			"<b>3. Използване на данните</b>\n" +
			"3.1. Събраните данни се използват изключително за:\n" +
			"• Предоставяне на достъп до услугата (авторизация).\n" +
			"• Техническа поддръжка на потребителите.\n\n" +
			"<b>4. Предаване на данни на трети страни</b>\n" +
			"4.1. Ние не продаваме и не предаваме вашите лични данни на трети страни, освен в случаи, предвидени от закона, или за обработка на плащанията през сигурни шлюзове.",
	},
	language.Српски: {
		Policy: "<b>ПОЛИТИКА ПРИВАТНОСТИ:</b>\n\n" +
			"<b>1. Које податке скупљамо</b>\n" +
			"За функционисање сервиса THE BEYOND скупљамо минимално неопходан сет података:\n" +
			"• Telegram User ID: За идентификацију корисника у систему и аутоматско издавање кључева.\n" +
			"• Метаподаци плаћања: Износ, време и начин плаћања (ми не чувамо пуне податке банкарских картица, обрада се врши на страни платног гејтвеја).\n" +
			"• Техничка статистика: Обим потрошеног саобраћаја (у бајтовима) за контролу лимита претплате.\n\n" +
			"<b>2. Политика без логова (No-Logs Policy)</b>\n" +
			"2.1. Ми се придржавамо строге политике приватности. Ми НЕ скупљамо, не пратимо и не чувамо:\n" +
			"• Историју посећених сајтова (DNS-захтеви).\n" +
			"• Садржај пренетог саобраћаја.\n" +
			"• Стварне IP адресе корисника при повезивању на сервер (логови повезивања се бришу циклично).\n\n" +
			"<b>3. Коришћење података</b>\n" +
			"3.1. Сакупљени подаци се користе искључиво за:\n" +
			"• Пружање приступа сервису (ауторизација).\n" +
			"• Техничку подршку корисницима.\n\n" +
			"<b>4. Пренос података трећим лицима</b>\n" +
			"4.1. Ми не продајемо и не преносимо ваше личне податке трећим лицима, осим у случајевима предвиђеним законом или за обраду плаћања преко сигурних гејтвеја.",
	},
	language.Hrvatski: {
		Policy: "<b>POLITIKA PRIVATNOSTI:</b>\n\n" +
			"<b>1. Koje podatke prikupljamo</b>\n" +
			"Za funkcioniranje servisa THE BEYOND prikupljamo minimalno potreban set podataka:\n" +
			"• Telegram User ID: Za identifikaciju korisnika u sustavu i automatsko izdavanje ključeva.\n" +
			"• Metapodaci plaćanja: Iznos, vrijeme i način plaćanja (ne pohranjujemo potpune podatke bankovnih kartica, obrada se vrši na strani platnog prolaza).\n" +
			"• Tehnička statistika: Obujam potrošenog prometa (u bajtovima) za kontrolu ograničenja pretplate.\n\n" +
			"<b>2. Politika bez logova (No-Logs Policy)</b>\n" +
			"2.1. Pridržavamo se stroge politike privatnosti. NE prikupljamo, ne pratimo i ne pohranjujemo:\n" +
			"• Povijest posjećenih stranica (DNS upiti).\n" +
			"• Sadržaj prenesenog prometa.\n" +
			"• Stvarne IP adrese korisnika prilikom povezivanja na server (logovi povezivanja se brišu ciklički).\n\n" +
			"<b>3. Korištenje podataka</b>\n" +
			"3.1. Prikupljeni podaci koriste se isključivo za:\n" +
			"• Pružanje pristupa servisu (autorizacija).\n" +
			"• Tehničku podršku korisnicima.\n\n" +
			"<b>4. Prijenos podataka trećim stranama</b>\n" +
			"4.1. Ne prodajemo i ne prenosimo vaše osobne podatke trećim stranama, osim u slučajevima predviđenim zakonom ili za obradu plaćanja preko sigurnih prolaza.",
	},
	language.Slovenčina: {
		Policy: "<b>ZÁSADY OCHRANy SÚKROMIA:</b>\n\n" +
			"<b>1. Aké údaje zhromažďujeme</b>\n" +
			"Pre fungovanie služby THE BEYOND zhromažďujeme minimálne potrebné množstvo údajov:\n" +
			"• Telegram User ID: Pre identifikáciu používateľa v systéme a automatické vydávanie kľúčov.\n" +
			"• Metadáta platieb: Suma, čas a spôsob platby (neukladáme úplné údaje bankových kariet, spracovanie prebieha na strane platobnej brány).\n" +
			"• Technické štatistiky: Objem spotrebovanej prevádzky (v bajtoch) na kontrolu limitov predplatného.\n\n" +
			"<b>2. Zásady No-Logs</b>\n" +
			"2.1. Dodržiavame prísne zásady ochrany súkromia. NEZHROMAŽĎUJEME, nesledujeme ani neukladáme:\n" +
			"• Históriu navštívených stránok (DNS dotazy).\n" +
			"• Obsah prenášanej prevádzky.\n" +
			"• Skutočné IP adresy používateľov pri pripojení k serveru (logy pripojenia sa mažu cyklicky).\n\n" +
			"<b>3. Použitie údajov</b>\n" +
			"3.1. Zhromaždené údaje sa používajú výlučne na:\n" +
			"• Poskytnutie prístupu k službe (autorizácia).\n" +
			"• Technickú podporu používateľov.\n\n" +
			"<b>4. Prenos údajov tretím stranám</b>\n" +
			"4.1. Nepredávame ani neprenášame vaše osobné údaje tretím stranám, okrem prípadov stanovených zákonom alebo na spracovanie platieb prostredníctvom bezpečných brán.",
	},
	language.Slovenski: {
		Policy: "<b>POLITIKA ZASEBNOSTI:</b>\n\n" +
			"<b>1. Katere podatke zbiramo</b>\n" +
			"Za delovanje storitve THE BEYOND zbiramo minimalno potreben nabor podatkov:\n" +
			"• Telegram User ID: Za identifikacijo uporabnika v sistemu in samodejno izdajo ključev.\n" +
			"• Metapodatki plačil: Znesek, čas in način plačila (ne shranjujemo polnih podatkov bančnih kartic, obdelava poteka na strani plačilnega prehoda).\n" +
			"• Tehnične statistike: Obseg porabljenega prometa (v bajtih) za nadzor omejitev naročnine.\n\n" +
			"<b>2. Politika brez dnevnikov (No-Logs Policy)</b>\n" +
			"2.1. Držimo se stroge politike zasebnosti. NE zbiramo, ne sledimo in ne shranjujemo:\n" +
			"• Zgodovino obiskanih strani (DNS poizvedbe).\n" +
			"• Vsebine prenesenega prometa.\n" +
			"• Resnične IP naslove uporabnikov pri povezavi s strežnikom (dnevniki povezav se brišejo ciklično).\n\n" +
			"<b>3. Uporaba podatkov</b>\n" +
			"3.1. Zbrani podatki se uporabljajo izključno za:\n" +
			"• Zagotavljanje dostopa do storitve (avtorizacija).\n" +
			"• Tehnično podporo uporabnikom.\n\n" +
			"<b>4. Prenos podatkov tretjim osebam</b>\n" +
			"4.1. Ne prodajamo in ne prenašamo vaših osebnih podatkov tretjim osebam, razen v primerih, predvidenih z zakonodajo, ali za obdelavo plačil prek varnih prehodov.",
	},
	language.Lietùvių: {
		Policy: "<b>PRIVATUMO POLITIKA:</b>\n\n" +
			"<b>1. Kokius duomenis renkame</b>\n" +
			"Paslaugos THE BEYOND veikimui renkame minimaliai būtinus duomenis:\n" +
			"• Telegram User ID: Vartotojo identifikavimui sistemoje ir automatiniam raktų išdavimui.\n" +
			"• Mokėjimų metaduomenys: Suma, laikas ir mokėjimo būdas (mes nesaugome pilnų banko kortelių duomenų, apdorojimas vyksta mokėjimo šliuzo pusėje).\n" +
			"• Techninė statistika: Suvartoto srauto tūris (baitais) prenumeratos limitų kontrolei.\n\n" +
			"<b>2. No-Logs Politika</b>\n" +
			"2.1. Mes laikomės griežtos privatumo politikos. Mes NERENKAME, nestebime ir nesaugome:\n" +
			"• Lankytų svetainių istorijos (DNS užklausos).\n" +
			"• Perduodamo srauto turinio.\n" +
			"• Tikrųjų vartotojų IP adresų prisijungiant prie serverio (prisijungimo žurnalai trinami cikliškai).\n\n" +
			"<b>3. Duomenų naudojimas</b>\n" +
			"3.1. Surinkti duomenys naudojami tik:\n" +
			"• Prieigos prie paslaugos suteikimui (autorizacija).\n" +
			"• Techninei vartotojų pagalbai.\n\n" +
			"<b>4. Duomenų perdavimas trečiosioms šalims</b>\n" +
			"4.1. Mes neparduodame ir neperduodame jūsų asmens duomenų trečiosioms šalims, išskyrus atvejus, numatytus teisės aktuose, arba mokėjimų apdorojimui per saugius šliuzus.",
	},
	language.Latviešu: {
		Policy: "<b>KONFIDENCIALITĀTES POLITIKA:</b>\n\n" +
			"<b>1. Kādus datus mēs savācam</b>\n" +
			"Pakalpojuma THE BEYOND darbībai mēs savācam minimāli nepieciešamo datu kopu:\n" +
			"• Telegram User ID: Lietotāja identifikācijai sistēmā un automātiskai atslēgu izsniegšanai.\n" +
			"• Maksājumu metadati: Summa, laiks un maksājuma veids (mēs neuzglabājam pilnus bankas karšu datus, apstrāde notiek maksājumu vārtejas pusē).\n" +
			"• Tehniskā statistika: Patērētā trafika apjoms (baitos) abonementa limitu kontrolei.\n\n" +
			"<b>2. No-Logs Politika</b>\n" +
			"2.1. Mēs ievērojam stingru konfidencialitātes politiku. Mēs NESAVĀCAM, nesekojam un neuzglabājam:\n" +
			"• Apmeklēto vietņu vēsturi (DNS pieprasījumi).\n" +
			"• Pārsūtītā trafika saturu.\n" +
			"• Lietotāju īstās IP adreses pievienojoties serverim (pieslēguma žurnāli tiek dzēsti ciklos).\n\n" +
			"<b>3. Datu izmantošana</b>\n" +
			"3.1. Savāktie dati tiek izmantoti tikai:\n" +
			"• Piekļuves nodrošināšanai pakalpojumam (autorizācija).\n" +
			"• Tehniskajam atbalstam lietotājiem.\n\n" +
			"<b>4. Datu nodošana trešajām pusēm</b>\n" +
			"4.1. Mēs nepārdodam un nenododam jūsu personiskos datus trešajām pusēm, izņemot gadījumus, kas paredzēti likumdošanā, vai maksājumu apstrādei caur drošām vārtejām.",
	},
	language.Eesti: {
		Policy: "<b>PRIVAATSUSPOLIITIKA:</b>\n\n" +
			"<b>1. Milliseid andmeid me kogume</b>\n" +
			"Teenuse THE BEYOND toimimiseks kogume minimaalselt vajalikud andmed:\n" +
			"• Telegram User ID: Kasutaja tuvastamiseks süsteemis ja võtmete automaatseks väljastamiseks.\n" +
			"• Maksete metaandmed: Summa, aeg ja makseviis (me ei salvesta täielikke pangakaardi andmeid, töötlemine toimub maksevärava poolel).\n" +
			"• Tehniline statistika: Tarbitud liikluse maht (baitides) tellimuse limiitide kontrollimiseks.\n\n" +
			"<b>2. No-Logs Poliitika</b>\n" +
			"2.1. Järgime ranget privaatsuspoliitikat. Me EI kogu, jälgi ega salvesta:\n" +
			"• Külastatud saitide ajalugu (DNS-päringud).\n" +
			"• Edastatud liikluse sisu.\n" +
			"• Kasutajate tegelikke IP-aadresse serveriga ühendamisel (ühenduse logid kustutatakse tsükliliselt).\n\n" +
			"<b>3. Andmete kasutamine</b>\n" +
			"3.1. Kogutud andmeid kasutatakse ainult:\n" +
			"• Juurdepääsu andmiseks teenusele (autoriseerimine).\n" +
			"• Tehnilise toega kasutajatele.\n\n" +
			"<b>4. Andmete edastamine kolmandatele isikutele</b>\n" +
			"4.1. Me ei müü ega edasta teie isikuandmeid kolmandatele isikutele, välja arvatud seaduses ette nähtud juhtudel või maksete töötlemiseks läbi turvaliste väravate.",
	},
	language.Suomi: {
		Policy: "<b>TIETOSUOJAKÄYTÄNTÖ:</b>\n\n" +
			"<b>1. Mitä tietoja keräämme</b>\n" +
			"THE BEYOND -palvelun toiminnan vuoksi keräämme vähimmäisvaatimuksen mukaisen tietojoukon:\n" +
			"• Telegram User ID: Käyttäjän tunnistamiseen järjestelmässä ja avainten automaattiseen myöntämiseen.\n" +
			"• Maksujen metatiedot: Summa, aika ja maksutapa (emme tallenna täydellisiä pankkikorttitietoja; käsittely tapahtuu maksuyhdyskäytävän puolella).\n" +
			"• Tekniset tilastot: Kulutetun liikenteen määrä (tavuissa) tilauksen rajojen hallintaan.\n\n" +
			"<b>2. No-Logs-Käytäntö</b>\n" +
			"2.1. Noudatamme tiukkaa tietosuojakäytäntöä. Emme KERÄÄ, seuraa tai tallenna:\n" +
			"• Vierailtujen sivustojen historiaa (DNS-kyselyt).\n" +
			"• Siirretyn liikenteen sisältöä.\n" +
			"• Käyttäjien todellisia IP-osoitteita yhteyden muodostamisen aikana palvelimeen (yhteyslokit poistetaan syklisesti).\n\n" +
			"<b>3. Tietojen käyttö</b>\n" +
			"3.1. Kerättyjä tietoja käytetään yksinomaan:\n" +
			"• Pääsyn myöntämiseen palveluun (valtuutus).\n" +
			"• Käyttäjien tekniseen tukeen.\n\n" +
			"<b>4. Tietojen siirto kolmansille osapuolille</b>\n" +
			"4.1. Emme myy tai siirrä henkilötietojasi kolmansille osapuolille, paitsi laissa säädetyissä tapauksissa tai maksujen käsittelyyn turvallisten yhdyskäytävien kautta.",
	},
	language.Ελληνικά: {
		Policy: "<b>ΠΟΛΙΤΙΚΗ ΑΠΟΡΡΗΤΟΥ:</b>\n\n" +
			"<b>1. Ποια δεδομένα συλλέγουμε</b>\n" +
			"Για τη λειτουργία της υπηρεσίας THE BEYOND συλλέγουμε το ελάχιστο απαραίτητο σύνολο δεδομένων:\n" +
			"• Telegram User ID: Για την ταυτοποίηση του χρήστη στο σύστημα και την αυτόματη έκδοση κλειδιών.\n" +
			"• Μεταδεδομένα πληρωμών: Ποσό, ώρα και τρόπος πληρωμής (δεν αποθηκεύουμε πλήρη δεδομένα τραπεζικών καρτών, η επεξεργασία γίνεται από την πλευρά της πύλης πληρωμών).\n" +
			"• Τεχνικά στατιστικά: Όγκος καταναλωθείσας κίνησης (σε bytes) για τον έλεγχο ορίων συνδρομής.\n\n" +
			"<b>2. Πολιτική No-Logs</b>\n" +
			"2.1. Τηρούμε αυστηρή πολιτική απορρήτου. ΔΕΝ συλλέγουμε, δεν παρακολουθούμε και δεν αποθηκεύουμε:\n" +
			"• Ιστορικό επισκεπτόμενων ιστότοπων (ερωτήματα DNS).\n" +
			"• Περιεχόμενο μεταδιδόμενης κίνησης.\n" +
			"• Πραγματικές διευθύνσεις IP χρηστών κατά τη σύνδεση στον διακομιστή (τα logs σύνδεσης διαγράφονται κυκλικά).\n\n" +
			"<b>3. Χρήση δεδομένων</b>\n" +
			"3.1. Τα συλλεγόμενα δεδομένα χρησιμοποιούνται αποκλειστικά για:\n" +
			"• Παροχή πρόσβασης στην υπηρεσία (εξουσιοδότηση).\n" +
			"• Τεχνική υποστήριξη χρηστών.\n\n" +
			"<b>4. Μεταβίβαση δεδομένων σε τρίτους</b>\n" +
			"4.1. Δεν πουλάμε και δεν μεταβιβάζουμε τα προσωπικά σας δεδομένα σε τρίτους, εκτός από τις περιπτώσεις που προβλέπονται από τη νομοθεσία ή για την επεξεργασία πληρωμών μέσω ασφαλών πυλών.",
	},
	language.Română: {
		Policy: "<b>POLITICA DE CONFIDENȚIALITATE:</b>\n\n" +
			"<b>1. Ce date colectăm</b>\n" +
			"Pentru funcționarea serviciului THE BEYOND colectăm setul minim necesar de date:\n" +
			"• Telegram User ID: Pentru identificarea utilizatorului în sistem și emiterea automată a cheilor.\n" +
			"• Metadate de plăți: Suma, ora și metoda de plată (nu stocăm datele complete ale cardurilor bancare, procesarea are loc pe partea gateway-ului de plăți).\n" +
			"• Statistici tehnice: Volumul traficului consumat (în octeți) pentru controlul limitelor abonamentului.\n\n" +
			"<b>2. Politica No-Logs</b>\n" +
			"2.1. Respectăm o politică strictă de confidențialitate. NU colectăm, nu urmărim și nu stocăm:\n" +
			"• Istoricul site-urilor vizitate (interogări DNS).\n" +
			"• Conținutul traficului transmis.\n" +
			"• Adrese IP reale ale utilizatorilor la conectarea la server (logurile de conexiune sunt șterse ciclic).\n\n" +
			"<b>3. Utilizarea datelor</b>\n" +
			"3.1. Datele colectate sunt utilizate exclusiv pentru:\n" +
			"• Furnizarea accesului la serviciu (autorizare).\n" +
			"• Suport tehnic pentru utilizatori.\n\n" +
			"<b>4. Transmiterea datelor către terți</b>\n" +
			"4.1. Nu vindem și nu transmitem datele dumneavoastră personale către terți, cu excepția cazurilor prevăzute de legislație sau pentru procesarea plăților prin gateway-uri sigure.",
	},
	language.Magyar: {
		Policy: "<b>ADATVÉDELMI SZABÁLYZAT:</b>\n\n" +
			"<b>1. Milyen adatokat gyűjtünk</b>\n" +
			"A THE BEYOND szolgáltatás működéséhez a minimálisan szükséges adatokat gyűjtjük:\n" +
			"• Telegram User ID: A felhasználó azonosításához a rendszerben és a kulcsok automatikus kiadásához.\n" +
			"• Fizetési metaadatok: Összeg, idő és fizetési mód (nem tároljuk a bankkártya teljes adatait, a feldolgozás a fizetési átjáró oldalán történik).\n" +
			"• Technikai statisztikák: Fogyasztott forgalom mennyisége (bájtban) az előfizetési limitek ellenőrzéséhez.\n\n" +
			"<b>2. No-Logs Szabályzat</b>\n" +
			"2.1. Szigorú adatvédelmi szabályzatot követünk. NEM gyűjtünk, nem követünk nyomon és nem tárolunk:\n" +
			"• Látogatott oldalak történetét (DNS-lekérdezések).\n" +
			"• Átvitt forgalom tartalmát.\n" +
			"• Felhasználók valós IP-címeit a szerverhez való csatlakozáskor (a csatlakozási naplók ciklikusan törlődnek).\n\n" +
			"<b>3. Adatok felhasználása</b>\n" +
			"3.1. A gyűjtött adatok kizárólag a következő célokra használhatók:\n" +
			"• Hozzáférés biztosítása a szolgáltatáshoz (engedélyezés).\n" +
			"• Technikai támogatás a felhasználóknak.\n\n" +
			"<b>4. Adatok átadása harmadik feleknek</b>\n" +
			"4.1. Nem adjuk el és nem adjuk át személyes adatait harmadik feleknek, kivéve a törvény által előírt eseteket vagy a fizetések feldolgozását biztonságos átjárókon keresztül.",
	},
	language.Arabic: {
		Policy: "<b>سياسة الخصوصية:</b>\n\n" +
			"<b>1. ما هي البيانات التي نجمعها</b>\n" +
			"لتشغيل خدمة THE BEYOND، نجمع الحد الأدنى اللازم من البيانات:\n" +
			"• معرف مستخدم Telegram: لتحديد هوية المستخدم في النظام وإصدار المفاتيح تلقائيًا.\n" +
			"• بيانات وصفية للدفعات: المبلغ، والوقت، وطريقة الدفع (لا نحتفظ ببيانات بطاقات البنوك الكاملة، يتم المعالجة على جانب بوابة الدفع).\n" +
			"• إحصاءات فنية: حجم الحركة المستهلكة (بالبايتات) للتحكم في حدود الاشتراك.\n\n" +
			"<b>2. سياسة عدم التسجيل (No-Logs Policy)</b>\n" +
			"2.1. نلتزم بسياسة خصوصية صارمة. لا نجمع، ولا نتبع، ولا نحتفظ بـ:\n" +
			"• تاريخ المواقع المزورة (استعلامات DNS).\n" +
			"• محتوى الحركة المرسلة.\n" +
			"• عناوين IP الحقيقية للمستخدمين عند الاتصال بالخادم (يتم حذف سجلات الاتصال دوريًا).\n\n" +
			"<b>3. استخدام البيانات</b>\n" +
			"3.1. تستخدم البيانات المجمعة حصريًا لـ:\n" +
			"• تقديم الوصول إلى الخدمة (التفويض).\n" +
			"• الدعم الفني للمستخدمين.\n\n" +
			"<b>4. نقل البيانات إلى أطراف ثالثة</b>\n" +
			"4.1. لا نبيع ولا ننقل بياناتك الشخصية إلى أطراف ثالثة، باستثناء الحالات المنصوص عليها في التشريع، أو لمعالجة الدفعات من خلال بوابات آمنة.",
	},
	language.Farsi: {
		Policy: "<b>سیاست حفظ حریم خصوصی:</b>\n\n" +
			"<b>1. چه داده‌هایی جمع‌آوری می‌کنیم</b>\n" +
			"برای عملکرد سرویس THE BEYOND، حداقل مجموعه لازم از داده‌ها را جمع‌آوری می‌کنیم:\n" +
			"• شناسه کاربر Telegram: برای شناسایی کاربر در سیستم و صدور خودکار کلیدها.\n" +
			"• متادیتای پرداخت‌ها: مبلغ، زمان و روش پرداخت (ما داده‌های کامل کارت‌های بانکی را ذخیره نمی‌کنیم، پردازش در سمت دروازه پرداخت انجام می‌شود).\n" +
			"• آمار فنی: حجم ترافیک مصرف‌شده (به بایت) برای کنترل محدودیت‌های اشتراک.\n\n" +
			"<b>2. سیاست عدم ثبت لاگ (No-Logs Policy)</b>\n" +
			"2.1. ما به سیاست حفظ حریم خصوصی سختگیرانه پایبند هستیم. ما جمع‌آوری نمی‌کنیم، پیگیری نمی‌کنیم و ذخیره نمی‌کنیم:\n" +
			"• تاریخچه سایت‌های بازدیدشده (پرس‌وجوهای DNS).\n" +
			"• محتوای ترافیک منتقل‌شده.\n" +
			"• آدرس‌های IP واقعی کاربران هنگام اتصال به سرور (لاگ‌های اتصال به صورت دوره‌ای حذف می‌شوند).\n\n" +
			"<b>3. استفاده از داده‌ها</b>\n" +
			"3.1. داده‌های جمع‌آوری‌شده منحصراً برای:\n" +
			"• ارائه دسترسی به سرویس (احراز هویت).\n" +
			"• پشتیبانی فنی کاربران.\n\n" +
			"<b>4. انتقال داده‌ها به اشخاص ثالث</b>\n" +
			"4.1. ما داده‌های شخصی شما را نمی‌فروشیم و منتقل نمی‌کنیم به اشخاص ثالث، به جز موارد پیش‌بینی‌شده توسط قانون، یا برای پردازش پرداخت‌ها از طریق دروازه‌های امن.",
	},
	language.Türkçe: {
		Policy: "<b>GİZLİLİK POLİTİKASI:</b>\n\n" +
			"<b>1. Hangi verileri topluyoruz</b>\n" +
			"THE BEYOND servisinin çalışması için minimum gerekli veri setini topluyoruz:\n" +
			"• Telegram User ID: Kullanıcının sistemde tanımlanması ve anahtarların otomatik olarak verilmesi için.\n" +
			"• Ödeme meta verileri: Tutar, zaman ve ödeme yöntemi (banka kartı tam verilerini saklamıyoruz, işlem ödeme ağ geçidi tarafında gerçekleşiyor).\n" +
			"• Teknik istatistikler: Tüketilen trafik hacmi (bayt cinsinden) abonelik limitlerinin kontrolü için.\n\n" +
			"<b>2. No-Logs Politikası</b>\n" +
			"2.1. Sıkı bir gizlilik politikasına uyuyoruz. TOPLAMAYIZ, takip etmeyiz ve saklamayız:\n" +
			"• Ziyaret edilen sitelerin geçmişi (DNS sorguları).\n" +
			"• İletilen trafiğin içeriği.\n" +
			"• Kullanıcıların gerçek IP adresleri sunucuya bağlanırken (bağlantı günlükleri döngüsel olarak silinir).\n\n" +
			"<b>3. Verilerin kullanımı</b>\n" +
			"3.1. Toplanan veriler yalnızca şu amaçlar için kullanılır:\n" +
			"• Servise erişim sağlamak (yetkilendirme).\n" +
			"• Kullanıcılara teknik destek.\n\n" +
			"<b>4. Verilerin üçüncü taraflara aktarımı</b>\n" +
			"4.1. Kişisel verilerinizi üçüncü taraflara satmıyoruz veya aktarmıyoruz, yasalarla öngörülen durumlar hariç veya ödemelerin güvenli ağ geçitleri aracılığıyla işlenmesi için.",
	},
	language.Hebrew: {
		Policy: "<b>מדיניות פרטיות:</b>\n\n" +
			"<b>1. אילו נתונים אנו אוספים</b>\n" +
			"לפעולת השירות THE BEYOND אנו אוספים את הסט המינימלי הנחוץ של נתונים:\n" +
			"• Telegram User ID: לזיהוי המשתמש במערכת ולמתן אוטומטי של מפתחות.\n" +
			"• מטא-נתונים של תשלומים: סכום, זמן ואופן תשלום (אנו לא מאחסנים נתונים מלאים של כרטיסי בנק, העיבוד מתבצע בצד שער התשלומים).\n" +
			"• סטטיסטיקה טכנית: נפח התעבורה הצרכנית (בבייטים) לבקרת מגבלות המנוי.\n\n" +
			"<b>2. מדיניות ללא לוגים (No-Logs Policy)</b>\n" +
			"2.1. אנו מקפידים על מדיניות פרטיות מחמירה. אנו לא אוספים, לא עוקבים ולא מאחסנים:\n" +
			"• היסטוריית אתרים שביקרת (שאילתות DNS).\n" +
			"• תוכן התעבורה המועברת.\n" +
			"• כתובות IP אמיתיות של משתמשים בעת חיבור לשרת (לוגי חיבור נמחקים בצורה מחזורית).\n\n" +
			"<b>3. שימוש בנתונים</b>\n" +
			"3.1. הנתונים שנאספו משמשים אך ורק ל:\n" +
			"• מתן גישה לשירות (אימות).\n" +
			"• תמיכה טכנית למשתמשים.\n\n" +
			"<b>4. העברת נתונים לצדדים שלישיים</b>\n" +
			"4.1. אנו לא מוכרים ולא מעבירים את הנתונים האישיים שלכם לצדדים שלישיים, למעט מקרים הקבועים בחוק, או לעיבוד תשלומים דרך שערים מאובטחים.",
	},
	language.ZH中文: {
		Policy: "<b>隐私政策：</b>\n\n" +
			"<b>1. 我们收集哪些数据</b>\n" +
			"为了 THE BEYOND 服务的运行，我们收集最小必要的数据集：\n" +
			"• Telegram User ID：用于系统中用户识别和自动发放密钥。\n" +
			"• 支付元数据：金额、时间和支付方式（我们不存储完整的银行卡数据，处理在支付网关一侧进行）。\n" +
			"• 技术统计：消耗流量量（以字节为单位）用于订阅限额控制。\n\n" +
			"<b>2. 无日志政策 (No-Logs Policy)</b>\n" +
			"2.1. 我们遵守严格的隐私政策。我们不收集、不跟踪也不存储：\n" +
			"• 访问站点历史（DNS 查询）。\n" +
			"• 传输流量的内容。\n" +
			"• 用户连接到服务器时的真实 IP 地址（连接日志周期性删除）。\n\n" +
			"<b>3. 数据使用</b>\n" +
			"3.1. 收集的数据仅用于：\n" +
			"• 提供服务访问（授权）。\n" +
			"• 用户技术支持。\n\n" +
			"<b>4. 数据传输给第三方</b>\n" +
			"4.1. 我们不出售也不传输您的个人数据给第三方，除非法律规定，或通过安全网关处理支付。",
	},
	language.JA日本語: {
		Policy: "<b>プライバシーポリシー：</b>\n\n" +
			"<b>1. どのようなデータを収集しますか</b>\n" +
			"THE BEYOND サービスの運用ために、最小限必要なデータセットを収集します：\n" +
			"• Telegram User ID: システムでのユーザー識別とキーの自動発行のため。\n" +
			"• 支払いメタデータ: 金額、時間、支払い方法（銀行カードの完全なデータを保存しません。処理は支払いゲートウェイ側で行われます）。\n" +
			"• 技術統計: 消費トラフィック量（バイト単位）でサブスクリプション制限の制御のため。\n\n" +
			"<b>2. ノーログポリシー (No-Logs Policy)</b>\n" +
			"2.1. 厳格なプライバシーポリシーを遵守します。私たちは収集、追跡、保存しません：\n" +
			"• 訪問したサイトの履歴（DNSクエリ）。\n" +
			"• 伝送トラフィックの内容。\n" +
			"• サーバー接続時のユーザーの実際のIPアドレス（接続ログは周期的に削除されます）。\n\n" +
			"<b>3. データの使用</b>\n" +
			"3.1. 収集されたデータは専ら以下のために使用されます：\n" +
			"• サービスへのアクセス提供（認証）。\n" +
			"• ユーザーの技術サポート。\n\n" +
			"<b>4. データの第三者への転送</b>\n" +
			"4.1. 私たちはあなたの個人データを第三者に販売または転送しません。法律で規定された場合、またはセキュアゲートウェイを通じた支払い処理を除きます。",
	},
	language.KO한국어: {
		Policy: "<b>개인정보 보호정책:</b>\n\n" +
			"<b>1. 어떤 데이터를 수집하나요</b>\n" +
			"THE BEYOND 서비스의 작동을 위해 최소 필요한 데이터 세트를 수집합니다:\n" +
			"• Telegram User ID: 시스템에서 사용자 식별 및 키 자동 발급을 위해.\n" +
			"• 결제 메타데이터: 금액, 시간 및 결제 방법 (은행 카드 전체 데이터를 저장하지 않음, 처리 는 결제 게이트웨이 측에서 이뤄짐).\n" +
			"• 기술 통계: 소비 트래픽 양 (바이트 단위) 구독 제한 제어를 위해.\n\n" +
			"<b>2. No-Logs 정책</b>\n" +
			"2.1. 엄격한 개인정보 보호 정책을 준수합니다. 우리는 수집, 추적 또는 저장하지 않습니다:\n" +
			"• 방문 사이트 기록 (DNS 쿼리).\n" +
			"• 전송 트래픽 내용.\n" +
			"• 서버 연결 시 사용자 실제 IP 주소 (연결 로그는 주기적으로 삭제됨).\n\n" +
			"<b>3. 데이터 사용</b>\n" +
			"3.1. 수집된 데이터는 오직 다음을 위해 사용됩니다:\n" +
			"• 서비스 액세스 제공 (인증).\n" +
			"• 사용자 기술 지원.\n\n" +
			"<b>4. 데이터의 제3자 전송</b>\n" +
			"4.1. 우리는 귀하의 개인 데이터를 제3자에게 판매하거나 전송하지 않습니다. 법률에 규정된 경우나 보안 게이트웨이를 통한 결제 처리 제외.",
	},
	language.TiếngViệt: {
		Policy: "<b>CHÍNH SÁCH BẢO MẬT:</b>\n\n" +
			"<b>1. Chúng tôi thu thập dữ liệu gì</b>\n" +
			"Để vận hành dịch vụ THE BEYOND, chúng tôi thu thập bộ dữ liệu tối thiểu cần thiết:\n" +
			"• Telegram User ID: Để nhận dạng người dùng trong hệ thống và cấp khóa tự động.\n" +
			"• Siêu dữ liệu thanh toán: Số tiền, thời gian và phương thức thanh toán (chúng tôi không lưu trữ dữ liệu thẻ ngân hàng đầy đủ, xử lý diễn ra bên cổng thanh toán).\n" +
			"• Thống kê kỹ thuật: Khối lượng lưu lượng tiêu thụ (theo byte) để kiểm soát giới hạn đăng ký.\n\n" +
			"<b>2. Chính sách Không Log (No-Logs Policy)</b>\n" +
			"2.1. Chúng tôi tuân thủ chính sách bảo mật nghiêm ngặt. Chúng tôi KHÔNG thu thập, theo dõi hoặc lưu trữ:\n" +
			"• Lịch sử các trang web đã truy cập (truy vấn DNS).\n" +
			"• Nội dung lưu lượng truyền.\n" +
			"• Địa chỉ IP thực của người dùng khi kết nối với máy chủ (log kết nối được xóa theo chu kỳ).\n\n" +
			"<b>3. Sử dụng dữ liệu</b>\n" +
			"3.1. Dữ liệu thu thập được sử dụng độc quyền cho:\n" +
			"• Cung cấp truy cập dịch vụ (xác thực).\n" +
			"• Hỗ trợ kỹ thuật cho người dùng.\n\n" +
			"<b>4. Chuyển dữ liệu cho bên thứ ba</b>\n" +
			"4.1. Chúng tôi không bán hoặc chuyển dữ liệu cá nhân của bạn cho bên thứ ba, trừ các trường hợp quy định bởi pháp luật hoặc để xử lý thanh toán qua cổng an toàn.",
	},
	language.THภาษาไทย: {
		Policy: "<b>นโยบายความเป็นส่วนตัว:</b>\n\n" +
			"<b>1. ข้อมูลอะไรที่เรารวบรวม</b>\n" +
			"สำหรับการทำงานของบริการ THE BEYOND เรารวบรวมชุดข้อมูลที่จำเป็นขั้นต่ำ:\n" +
			"• Telegram User ID: สำหรับการระบุตัวผู้ใช้ในระบบและการออกกุญแจอัตโนมัติ.\n" +
			"• ข้อมูลเมตาของการชำระเงิน: จำนวนเงิน เวลา และวิธีการชำระเงิน (เราไม่เก็บข้อมูลบัตรธนาคารเต็มรูปแบบ การประมวลผลเกิดขึ้นที่ด้านเกตเวย์การชำระเงิน).\n" +
			"• สถิติทางเทคนิค: ปริมาณการรับส่งข้อมูลที่ใช้ (ในไบต์) สำหรับการควบคุมขีดจำกัดการสมัครสมาชิก.\n\n" +
			"<b>2. นโยบายไม่มีล็อก (No-Logs Policy)</b>\n" +
			"2.1. เรายึดมั่นในนโยบายความเป็นส่วนตัวที่เข้มงวด เราไม่รวบรวม ไม่ติดตาม และไม่เก็บ:\n" +
			"• ประวัติเว็บไซต์ที่เยี่ยมชม (การสอบถาม DNS).\n" +
			"• เนื้อหาของการรับส่งข้อมูลที่ส่ง.\n" +
			"• ที่อยู่ IP จริงของผู้ใช้เมื่อเชื่อมต่อกับเซิร์ฟเวอร์ (ล็อกการเชื่อมต่อถูกลบตามรอบ).\n\n" +
			"<b>3. การใช้ข้อมูล</b>\n" +
			"3.1. ข้อมูลที่รวบรวมใช้เฉพาะสำหรับ:\n" +
			"• การให้สิทธิ์เข้าถึงบริการ (การอนุมัติ).\n" +
			"• การสนับสนุนทางเทคนิคสำหรับผู้ใช้.\n\n" +
			"<b>4. การถ่ายโอนข้อมูลให้บุคคลที่สาม</b>\n" +
			"4.1. เราไม่ขายและไม่ถ่ายโอนข้อมูลส่วนบุคคลของคุณให้บุคคลที่สาม ยกเว้นกรณีที่กฎหมายกำหนด หรือสำหรับการประมวลผลการชำระเงินผ่านเกตเวย์ที่ปลอดภัย.",
	},
	language.BahasaIndonesia: {
		Policy: "<b>KEBIJAKAN PRIVASI:</b>\n\n" +
			"<b>1. Data apa yang kami kumpulkan</b>\n" +
			"Untuk pengoperasian layanan THE BEYOND, kami mengumpulkan set data minimal yang diperlukan:\n" +
			"• Telegram User ID: Untuk identifikasi pengguna dalam sistem dan penerbitan kunci otomatis.\n" +
			"• Metadata pembayaran: Jumlah, waktu, dan metode pembayaran (kami tidak menyimpan data kartu bank lengkap, pemrosesan dilakukan di sisi gateway pembayaran).\n" +
			"• Statistik teknis: Volume lalu lintas yang dikonsumsi (dalam byte) untuk kontrol batas langganan.\n\n" +
			"<b>2. Kebijakan No-Logs</b>\n" +
			"2.1. Kami mematuhi kebijakan privasi yang ketat. Kami TIDAK mengumpulkan, melacak, atau menyimpan:\n" +
			"• Riwayat situs yang dikunjungi (permintaan DNS).\n" +
			"• Isi lalu lintas yang ditransmisikan.\n" +
			"• Alamat IP nyata pengguna saat terhubung ke server (log koneksi dihapus secara siklis).\n\n" +
			"<b>3. Penggunaan data</b>\n" +
			"3.1. Data yang dikumpulkan digunakan secara eksklusif untuk:\n" +
			"• Memberikan akses ke layanan (otorisasi).\n" +
			"• Dukungan teknis untuk pengguna.\n\n" +
			"<b>4. Transfer data ke pihak ketiga</b>\n" +
			"4.1. Kami tidak menjual atau mentransfer data pribadi Anda ke pihak ketiga, kecuali dalam kasus yang ditentukan oleh undang-undang, atau untuk pemrosesan pembayaran melalui gateway aman.",
	},
	language.BahasaMelayu: {
		Policy: "<b>DASAR PRIVASI:</b>\n\n" +
			"<b>1. Data apa yang kami kumpul</b>\n" +
			"Untuk operasi perkhidmatan THE BEYOND, kami mengumpul set data minimum yang diperlukan:\n" +
			"• Telegram User ID: Untuk pengenalan pengguna dalam sistem dan penerbitan kunci automatik.\n" +
			"• Metadata pembayaran: Jumlah, masa dan kaedah pembayaran (kami tidak menyimpan data kad bank lengkap, pemprosesan dilakukan di pihak gateway pembayaran).\n" +
			"• Statistik teknikal: Isipadu trafik yang digunakan (dalam bait) untuk kawalan had langganan.\n\n" +
			"<b>2. Dasar No-Logs</b>\n" +
			"2.1. Kami mematuhi dasar privasi yang ketat. Kami TIDAK mengumpul, menjejaki atau menyimpan:\n" +
			"• Sejarah laman web yang dilawati (permintaan DNS).\n" +
			"• Kandungan trafik yang dihantar.\n" +
			"• Alamat IP sebenar pengguna semasa menyambung ke pelayan (log sambungan dipadam secara siklik).\n\n" +
			"<b>3. Penggunaan data</b>\n" +
			"3.1. Data yang dikumpul digunakan secara eksklusif untuk:\n" +
			"• Memberi akses kepada perkhidmatan (pengesahan).\n" +
			"• Sokongan teknikal untuk pengguna.\n\n" +
			"<b>4. Pemindahan data kepada pihak ketiga</b>\n" +
			"4.1. Kami tidak menjual atau memindahkan data peribadi anda kepada pihak ketiga, kecuali dalam kes yang ditetapkan oleh undang-undang, atau untuk pemprosesan pembayaran melalui gateway selamat.",
	},
	language.Tagalog: {
		Policy: "<b>PATAKARAN SA PRIVACY:</b>\n\n" +
			"<b>1. Anong data ang kinokolekta namin</b>\n" +
			"Para sa pagpapatakbo ng serbisyo na THE BEYOND, kinokolekta namin ang minimal na kinakailangang set ng data:\n" +
			"• Telegram User ID: Para sa pagkilala sa user sa system at awtomatikong pagbibigay ng mga key.\n" +
			"• Metadata ng mga bayad: Halaga, oras at paraan ng bayad (hindi namin iniimbak ang buong data ng bank card, ang pagproseso ay ginagawa sa gilid ng payment gateway).\n" +
			"• Teknikal na istatistika: Dami ng natupok na traffic (sa bytes) para sa kontrol ng mga limitasyon sa subscription.\n\n" +
			"<b>2. Patakaran ng No-Logs</b>\n" +
			"2.1. Sumusunod kami sa mahigpit na patakaran sa privacy. HINDI kami nangongolekta, sumusubaybay o nag-iimbak ng:\n" +
			"• Kasaysayan ng mga binisitang site (mga query ng DNS).\n" +
			"• Nilalaman ng ipinadalang traffic.\n" +
			"• Tunay na IP address ng mga user kapag kumokonekta sa server (ang mga log ng koneksyon ay tinatanggal nang paulit-ulit).\n\n" +
			"<b>3. Paggamit ng data</b>\n" +
			"3.1. Ang nakolektang data ay ginagamit lamang para sa:\n" +
			"• Pagbibigay ng access sa serbisyo (awtorisasyon).\n" +
			"• Teknikal na suporta para sa mga user.\n\n" +
			"<b>4. Paglilipat ng data sa ikatlong partido</b>\n" +
			"4.1. Hindi namin ibinebenta o inililipat ang iyong personal na data sa ikatlong partido, maliban sa mga kasong itinakda ng batas o para sa pagproseso ng mga bayad sa pamamagitan ng ligtas na gateway.",
	},
	language.Hindi: {
		Policy: "<b>गोपनीयता नीति:</b>\n\n" +
			"<b>1. हम कौन से डेटा एकत्र करते हैं</b>\n" +
			"THE BEYOND सेवा के संचालन के लिए, हम न्यूनतम आवश्यक डेटा सेट एकत्र करते हैं:\n" +
			"• Telegram User ID: सिस्टम में उपयोगकर्ता की पहचान और कुंजियों की स्वचालित जारी करने के लिए.\n" +
			"• भुगतान मेटाडेटा: राशि, समय और भुगतान विधि (हम बैंक कार्ड के पूर्ण विवरण संग्रहीत नहीं करते; प्रसंस्करण भुगतान गेटवे की ओर से होता है).\n" +
			"• तकनीकी आंकड़े: खपत ट्रैफिक की मात्रा (बाइट्स में) सदस्यता सीमाओं के नियंत्रण के लिए.\n\n" +
			"<b>2. नो-लॉग्स नीति</b>\n" +
			"2.1. हम सख्त गोपनीयता नीति का पालन करते हैं. हम एकत्र नहीं करते, ट्रैक नहीं करते और संग्रहीत नहीं करते:\n" +
			"• देखी गई साइटों का इतिहास (DNS क्वेरी).\n" +
			"• प्रसारित ट्रैफिक की सामग्री.\n" +
			"• सर्वर से कनेक्ट होने पर उपयोगकर्ताओं के वास्तविक IP पते (कनेक्शन लॉग चक्रीय रूप से हटाए जाते हैं).\n\n" +
			"<b>3. डेटा का उपयोग</b>\n" +
			"3.1. एकत्रित डेटा का उपयोग विशेष रूप से किया जाता है:\n" +
			"• सेवा तक पहुंच प्रदान करने के लिए (प्रमाणीकरण).\n" +
			"• उपयोगकर्ताओं के लिए तकनीकी सहायता.\n\n" +
			"<b>4. डेटा का तीसरे पक्षों को हस्तांतरण</b>\n" +
			"4.1. हम आपके व्यक्तिगत डेटा को तीसरे पक्षों को नहीं बेचते या हस्तांतरित नहीं करते, सिवाय कानून द्वारा प्रदान किए गए मामलों के या सुरक्षित गेटवे के माध्यम से भुगतान प्रसंस्करण के लिए.",
	},
	language.URاردو: {
		Policy: "<b>پرائیویسی پالیسی:</b>\n\n" +
			"<b>1. ہم کون سے ڈیٹا اکٹھا کرتے ہیں</b>\n" +
			"THE BEYOND سروس کے کام کرنے کے لیے، ہم کم سے کم ضروری ڈیٹا سیٹ اکٹھا کرتے ہیں:\n" +
			"• ٹیلیگرام یوزر آئی ڈی: سسٹم میں صارف کی شناخت اور کیز کی خودکار اجراء کے لیے.\n" +
			"• ادائیگی میٹا ڈیٹا: رقم، وقت اور ادائیگی کا طریقہ (ہم بینک کارڈ کی مکمل تفصیلات محفوظ نہیں کرتے؛ پروسیسنگ ادائیگی گیٹ وے کی طرف سے ہوتی ہے).\n" +
			"• تکنیکی اعداد و شمار: استعمال شدہ ٹریفک کا حجم (بائٹس میں) سبسکرپشن کی حدود کے کنٹرول کے لیے.\n\n" +
			"<b>2. نو لاگز پالیسی</b>\n" +
			"2.1. ہم سخت پرائیویسی پالیسی کی پابندی کرتے ہیں. ہم اکٹھا نہیں کرتے، ٹریک نہیں کرتے اور محفوظ نہیں کرتے:\n" +
			"• دیکھی گئی سائٹس کی تاریخ (ڈی این ایس کوئریز).\n" +
			"• منتقل کی گئی ٹریفک کی مواد.\n" +
			"• سرور سے جڑنے پر صارفین کے حقیقی آئی پی ایڈریس (کنکشن لاگز چکری طور پر حذف کیے جاتے ہیں).\n\n" +
			"<b>3. ڈیٹا کا استعمال</b>\n" +
			"3.1. اکٹھا کیا گیا ڈیٹا صرف استعمال کیا جاتا ہے:\n" +
			"• سروس تک رسائی فراہم کرنے کے لیے (توثیق).\n" +
			"• صارفین کے لیے تکنیکی سپورٹ.\n\n" +
			"<b>4. ڈیٹا کا تیسرے فریقوں کو منتقل کرنا</b>\n" +
			"4.1. ہم آپ کے ذاتی ڈیٹا کو تیسرے فریقوں کو نہیں بیچتے یا منتقل نہیں کرتے، سوائے قانون کی طرف سے فراہم کیے گئے معاملات کے یا محفوظ گیٹ ویز کے ذریعے ادائیگی پروسیسنگ کے لیے.",
	},
	language.Bengali: {
		Policy: "<b>গোপনীয়তা নীতি:</b>\n\n" +
			"<b>1. আমরা কোন ডেটা সংগ্রহ করি</b>\n" +
			"THE BEYOND পরিষেবার কার্যকারিতার জন্য, আমরা সর্বনিম্ন প্রয়োজনীয় ডেটা সেট সংগ্রহ করি:\n" +
			"• টেলিগ্রাম ব্যবহারকারী আইডি: সিস্টেমে ব্যবহারকারীর পরিচয় এবং কীগুলির স্বয়ংক্রিয় ইস্যু করার জন্য.\n" +
			"• পেমেন্ট মেটাডেটা: পরিমাণ, সময় এবং পেমেন্ট পদ্ধতি (আমরা ব্যাঙ্ক কার্ডের সম্পূর্ণ বিবরণ সংরক্ষণ করি না; প্রক্রিয়াকরণ পেমেন্ট গেটওয়ের পক্ষ থেকে ঘটে).\n" +
			"• প্রযুক্তিগত পরিসংখ্যান: খরচ করা ট্র্যাফিকের পরিমাণ (বাইটে) সাবস্ক্রিপশন সীমার নিয়ন্ত্রণের জন্য.\n\n" +
			"<b>2. নো-লগস নীতি</b>\n" +
			"2.1. আমরা কঠোর গোপনীয়তা নীতি মেনে চলি. আমরা সংগ্রহ করি না, ট্র্যাক করি না এবং সংরক্ষণ করি না:\n" +
			"• দেখা সাইটগুলির ইতিহাস (ডিএনএস কোয়েরি).\n" +
			"• প্রেরিত ট্র্যাফিকের সামগ্রী.\n" +
			"• সার্ভারে সংযোগ করার সময় ব্যবহারকারীদের বাস্তব আইপি ঠিকানা (সংযোগ লগগুলি চক্রাকারে মুছে ফেলা হয়).\n\n" +
			"<b>3. ডেটার ব্যবহার</b>\n" +
			"3.1. সংগ্রহিত ডেটা শুধুমাত্র ব্যবহার করা হয়:\n" +
			"• পরিষেবায় অ্যাক্সেস প্রদানের জন্য (অনুমোদন).\n" +
			"• ব্যবহারকারীদের জন্য প্রযুক্তিগত সহায়তা.\n\n" +
			"<b>4. ডেটার তৃতীয় পক্ষে হস্তান্তর</b>\n" +
			"4.1. আমরা আপনার ব্যক্তিগত ডেটা তৃতীয় পক্ষে বিক্রি করি না বা হস্তান্তর করি না, আইন দ্বারা প্রদান করা ক্ষেত্রগুলি ছাড়া বা নিরাপদ গেটওয়ে মাধ্যমে পেমেন্ট প্রক্রিয়াকরণের জন্য.",
	},
	language.Tamiḻ: {
		Policy: "<b>தனியுரிமை கொள்கை:</b>\n\n" +
			"<b>1. நாங்கள் எந்த தரவுகளை சேகரிக்கிறோம்</b>\n" +
			"THE BEYOND சேவையின் செயல்பாட்டிற்காக, நாங்கள் குறைந்தபட்ச தேவையான தரவு தொகுப்பை சேகரிக்கிறோம்:\n" +
			"• டெலிகிராம் பயனர் ஐடி: அமைப்பில் பயனரின் அடையாளம் மற்றும் விசைகளின் தானியங்கி வழங்கலுக்கு.\n" +
			"• கட்டண மெட்டாடேட்டா: தொகை, நேரம் மற்றும் கட்டண முறை (வங்கி அட்டையின் முழு விவரங்களை சேமிக்கவில்லை; செயலாக்கம் கட்டண நுழைவாயிலின் பக்கத்தில் நடக்கிறது).\n" +
			"• தொழில்நுட்ப புள்ளிவிவரங்கள்: உட்கொள்ளப்பட்ட போக்குவரத்தின் அளவு (பைட்டுகளில்) சந்தா வரம்புகளின் கட்டுப்பாட்டிற்கு.\n\n" +
			"<b>2. நோ-லாக்ஸ் கொள்கை</b>\n" +
			"2.1. நாங்கள் கடுமையான தனியுரிமை கொள்கையை பின்பற்றுகிறோம். நாங்கள் சேகரிக்கவில்லை, கண்காணிக்கவில்லை மற்றும் சேமிக்கவில்லை:\n" +
			"• பார்வையிடப்பட்ட தளங்களின் வரலாறு (DNS கேள்விகள்).\n" +
			"• அனுப்பப்பட்ட போக்குவரத்தின் உள்ளடக்கம்.\n" +
			"• சர்வருடன் இணைக்கும் போது பயனர்களின் உண்மையான IP முகவரிகள் (இணைப்பு பதிவுகள் சுழற்சியாக நீக்கப்படுகின்றன).\n\n" +
			"<b>3. தரவுகளின் பயன்பாடு</b>\n" +
			"3.1. சேகரிக்கப்பட்ட தரவுகள் பிரத்தியேகமாக பயன்படுத்தப்படுகின்றன:\n" +
			"• சேவைக்கு அணுகல் வழங்குவதற்கு (அங்கீகாரம்).\n" +
			"• பயனர்களுக்கு தொழில்நுட்ப ஆதரவு.\n\n" +
			"<b>4. தரவுகளை மூன்றாம் தரப்பினருக்கு மாற்றுதல்</b>\n" +
			"4.1. உங்கள் தனிப்பட்ட தரவுகளை மூன்றாம் தரப்பினருக்கு விற்கவோ அல்லது மாற்றவோ இல்லை, சட்டத்தால் வழங்கப்பட்ட வழக்குகள் தவிர அல்லது பாதுகாப்பான நுழைவாயில்கள் மூலம் கட்டண செயலாக்கத்திற்கு.",
	},
	language.Telugu: {
		Policy: "<b>గోప్యతా విధానం:</b>\n\n" +
			"<b>1. మేము ఏ డేటాను సేకరిస్తాము</b>\n" +
			"THE BEYOND సేవ యొక్క కార్యాచరణ కోసం, మేము కనిష్ట అవసరమైన డేటా సెట్‌ను సేకరిస్తాము:\n" +
			"• టెలిగ్రామ్ యూజర్ ఐడి: సిస్టమ్‌లో యూజర్ గుర్తింపు మరియు కీల ఆటోమేటిక్ ఇష్యూ కోసం.\n" +
			"• చెల్లింపు మెటాడేటా: మొత్తం, సమయం మరియు చెల్లింపు పద్ధతి (మేము బ్యాంక్ కార్డ్ పూర్తి వివరాలను నిల్వ చేయము; ప్రాసెసింగ్ చెల్లింపు గేట్‌వే వైపు జరుగుతుంది).\n" +
			"• సాంకేతిక గణాంకాలు: ఖర్చు చేసిన ట్రాఫిక్ పరిమాణం (బైట్లలో) సబ్‌స్క్రిప్షన్ పరిమితుల నియంత్రణ కోసం.\n\n" +
			"<b>2. నో-లాగ్స్ విధానం</b>\n" +
			"2.1. మేము కఠినమైన గోప్యతా విధానాన్ని అనుసరిస్తాము. మేము సేకరించము, ట్రాక్ చేయము మరియు నిల్వ చేయము:\n" +
			"• సందర్శించిన సైట్ల చరిత్ర (DNS క్వెరీలు).\n" +
			"• ప్రసారం చేసిన ట్రాఫిక్ కంటెంట్.\n" +
			"• సర్వర్‌కు కనెక్ట్ అయినప్పుడు యూజర్ల రియల్ IP అడ్రెస్‌లు (కనెక్షన్ లాగ్‌లు చక్రీయంగా తొలగించబడతాయి).\n\n" +
			"<b>3. డేటా ఉపయోగం</b>\n" +
			"3.1. సేకరించిన డేటా మాత్రమే ఉపయోగించబడుతుంది:\n" +
			"• సేవకు యాక్సెస్ అందించడానికి (అధికారికరణ).\n" +
			"• యూజర్లకు సాంకేతిక మద్దతు.\n\n" +
			"<b>4. డేటాను మూడవ పక్షాలకు బదిలీ చేయడం</b>\n" +
			"4.1. మేము మీ వ్యక్తిగత డేటాను మూడవ పక్షాలకు విక్రయించము లేదా బదిలీ చేయము, చట్టం ద్వారా అందించబడిన కేసులు మినహా లేదా సురక్షిత గేట్‌వేల ద్వారా చెల్లింపు ప్రాసెసింగ్ కోసం.",
	},
	language.Marathi: {
		Policy: "<b>गोपनीयता धोरण:</b>\n\n" +
			"<b>1. आम्ही कोणता डेटा गोळा करतो</b>\n" +
			"THE BEYOND सेवेच्या कार्यासाठी, आम्ही किमान आवश्यक डेटा संच गोळा करतो:\n" +
			"• टेलिग्राम वापरकर्ता आयडी: सिस्टममध्ये वापरकर्त्याच्या ओळखीसाठी आणि कीजच्या स्वयंचलित जारीसाठी.\n" +
			"• पेमेंट मेटाडेटा: रक्कम, वेळ आणि पेमेंट पद्धत (आम्ही बँक कार्डच्या पूर्ण तपशील संग्रहित करत नाही; प्रक्रिया पेमेंट गेटवेच्या बाजूने होते).\n" +
			"• तांत्रिक आकडेवारी: वापरलेल्या ट्रॅफिकचे प्रमाण (बाइट्समध्ये) सदस्यता मर्यादांच्या नियंत्रणासाठी.\n\n" +
			"<b>2. नो-लॉग्स धोरण</b>\n" +
			"2.1. आम्ही कठोर गोपनीयता धोरणाचे पालन करतो. आम्ही गोळा करत नाही, ट्रॅक करत नाही आणि संग्रहित करत नाही:\n" +
			"• भेट दिलेल्या साइट्सचा इतिहास (DNS क्वेरी).\n" +
			"• प्रसारित ट्रॅफिकचा सामग्री.\n" +
			"• सर्व्हरशी कनेक्ट होताना वापरकर्त्यांच्या वास्तविक IP पत्ते (कनेक्शन लॉग चक्रीय पद्धतीने हटवले जातात).\n\n" +
			"<b>3. डेटाचा वापर</b>\n" +
			"3.1. गोळा केलेला डेटा केवळ वापरला जातो:\n" +
			"• सेवेसाठी प्रवेश प्रदान करण्यासाठी (प्रमाणीकरण).\n" +
			"• वापरकर्त्यांसाठी तांत्रिक समर्थन.\n\n" +
			"<b>4. डेटाचा तिसऱ्या पक्षांना हस्तांतर</b>\n" +
			"4.1. आम्ही तुमचा वैयक्तिक डेटा तिसऱ्या पक्षांना विकत नाही किंवा हस्तांतरित करत नाही, कायद्याने प्रदान केलेल्या प्रकरणांशिवाय किंवा सुरक्षित गेटवेच्या माध्यमातून पेमेंट प्रक्रियेसाठी.",
	},
}

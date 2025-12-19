package i18n

import "github.com/quickpowered/thebeyond-master/configs/language"

type RefundPolicyLocale struct {
	Policy string
}

var RefundPolicyMessages = map[language.Language]RefundPolicyLocale{
	language.English: {
		Policy: "<b>REFUND POLICY:</b>\n\n" +
			"<b>1. Refund Conditions</b>\n" +
			"1.1. Since we provide digital goods (access keys) and services that are consumed upon use, refunds are only possible in exceptional cases.\n" +
			"1.2. You have the right to request a full refund within 24 hours of purchase if you were unable to connect to the service due to technical issues on the Provider's side (non-working server, key generation error).\n\n" +
			"<b>2. Refund Denials</b>\n" +
			"Refunds are NOT issued in the following cases:\n" +
			"• More than 24 hours have passed since the purchase.\n" +
			"• Low connection speed caused by limitations on the user's provider side or poor mobile network coverage.\n" +
			"• User's inability to set up third-party software (Happ, Streisand clients, etc.), provided the key itself is valid.\n" +
			"• Access blocked for violating rules (spam, torrents, hacking attacks).\n" +
			"• \"I changed my mind\" or \"I don't need it anymore\".\n\n" +
			"<b>3. Refund Procedure</b>\n" +
			"3.1. To request a refund, contact support.\n" +
			"3.2. In your request, include: Account ID, payment receipt, and a description of the technical issue (preferably with a screenshot of the error).\n" +
			"3.3. The refund will be processed using the same payment method, within 3–7 business days.",
	},
	language.Deutsch: {
		Policy: "<b>RÜCKERSTATTUNGSRICHTLINIE:</b>\n\n" +
			"<b>1. Bedingungen für Rückerstattungen</b>\n" +
			"1.1. Da wir digitale Waren (Zugangsschlüssel) und Dienstleistungen anbieten, die beim Gebrauch verbraucht werden, sind Rückerstattungen nur in Ausnahmefällen möglich.\n" +
			"1.2. Sie haben das Recht, innerhalb von 24 Stunden nach dem Kauf eine vollständige Rückerstattung zu beantragen, falls Sie aufgrund technischer Probleme auf Seiten des Anbieters (nicht funktionierender Server, Fehler bei der Schlüsselerzeugung) nicht auf den Dienst zugreifen konnten.\n\n" +
			"<b>2. Verweigerung von Rückerstattungen</b>\n" +
			"Rückerstattungen werden NICHT gewährt in folgenden Fällen:\n" +
			"• Mehr als 24 Stunden sind seit dem Kauf vergangen.\n" +
			"• Niedrige Verbindungsgeschwindigkeit, verursacht durch Einschränkungen beim Provider des Nutzers oder schlechte Mobilfunkabdeckung.\n" +
			"• Unfähigkeit des Nutzers, Drittanbieter-Software (Happ-, Streisand-Clients usw.) einzurichten, sofern der Schlüssel selbst gültig ist.\n" +
			"• Sperrung des Zugangs wegen Verstoß gegen die Regeln (Spam, Torrents, Hacking-Attacken).\n" +
			"• „Ich habe es mir anders überlegt“ oder „Ich brauche es nicht mehr“.\n\n" +
			"<b>3. Rückerstattungsverfahren</b>\n" +
			"3.1. Um eine Rückerstattung zu beantragen, kontaktieren Sie den Support.\n" +
			"3.2. In Ihrer Anfrage geben Sie an: Account-ID, Zahlungsbeleg und eine Beschreibung des technischen Problems (vorzugsweise mit einem Screenshot des Fehlers).\n" +
			"3.3. Die Rückerstattung erfolgt über dieselbe Zahlungsmethode innerhalb von 3–7 Werktagen.",
	},
	language.Nederlands: {
		Policy: "<b>RESTITUTIEBELEID:</b>\n\n" +
			"<b>1. Voorwaarden voor restitutie</b>\n" +
			"1.1. Aangezien we digitale goederen (toegangssleutels) en diensten leveren die bij gebruik worden geconsumeerd, is restitutie alleen mogelijk in uitzonderlijke gevallen.\n" +
			"1.2. U hebt het recht om binnen 24 uur na aankoop een volledige restitutie aan te vragen als u niet kon verbinden met de dienst vanwege technische problemen aan de kant van de Aanbieder (niet-werkende server, fout bij sleutelgeneratie).\n\n" +
			"<b>2. Weigering van restitutie</b>\n" +
			"Restitutie wordt NIET verleend in de volgende gevallen:\n" +
			"• Meer dan 24 uur zijn verstreken sinds de aankoop.\n" +
			"• Lage verbindingssnelheid veroorzaakt door beperkingen aan de kant van de provider van de gebruiker of slechte dekking van het mobiele netwerk.\n" +
			"• Onvermogen van de gebruiker om derde-partij software (Happ, Streisand clients, etc.) in te stellen, mits de sleutel zelf geldig is.\n" +
			"• Blokkeren van toegang wegens schending van regels (spam, torrents, hackeraanvallen).\n" +
			"• „Ik ben van gedachten veranderd“ of „Ik heb het niet meer nodig“.\n\n" +
			"<b>3. Restitutieprocedure</b>\n" +
			"3.1. Om restitutie aan te vragen, neem contact op met de support.\n" +
			"3.2. In uw aanvraag vermeldt u: Account-ID, betalingsbewijs en een beschrijving van het technische probleem (bij voorkeur met een screenshot van de fout).\n" +
			"3.3. De restitutie wordt verwerkt via dezelfde betaalmethode, binnen 3–7 werkdagen.",
	},
	language.Svenska: {
		Policy: "<b>ÅTERBETALNINGSPOLICY:</b>\n\n" +
			"<b>1. Villkor för återbetalning</b>\n" +
			"1.1. Eftersom vi tillhandahåller digitala varor (åtkomstnycklar) och tjänster som konsumeras vid användning, är återbetalning endast möjlig i undantagsfall.\n" +
			"1.2. Du har rätt att begära full återbetalning inom 24 timmar efter köpet om du inte kunde ansluta till tjänsten på grund av tekniska problem från Leverantörens sida (icke-fungerande server, fel vid nyckelgenerering).\n\n" +
			"<b>2. Nekad återbetalning</b>\n" +
			"Återbetalning ges INTE i följande fall:\n" +
			"• Mer än 24 timmar har gått sedan köpet.\n" +
			"• Låg anslutningshastighet orsakad av begränsningar hos användarens provider eller dålig mobiltäckning.\n" +
			"• Användarens oförmåga att konfigurera tredjepartsprogram (Happ, Streisand-klienter osv.), förutsatt att nyckeln själv är giltig.\n" +
			"• Blockering av åtkomst för regelbrott (spam, torrents, hackerattacker).\n" +
			"• „Jag ångrade mig“ eller „Jag behöver det inte längre“.\n\n" +
			"<b>3. Återbetalningsprocedur</b>\n" +
			"3.1. För att begära återbetalning, kontakta supporten.\n" +
			"3.2. I din begäran ange: Konto-ID, betalningskvitto och en beskrivning av det tekniska problemet (helst med en skärmdump av felet).\n" +
			"3.3. Återbetalningen sker via samma betalningsmetod inom 3–7 arbetsdagar.",
	},
	language.Norsk: {
		Policy: "<b>REFUSJONSPOLICY:</b>\n\n" +
			"<b>1. Betingelser for refusjon</b>\n" +
			"1.1. Siden vi tilbyr digitale varer (tilgangsnøkler) og tjenester som forbrukes ved bruk, er refusjon bare mulig i unntakstilfeller.\n" +
			"1.2. Du har rett til å be om full refusjon innen 24 timer etter kjøpet hvis du ikke klarte å koble til tjenesten på grunn av tekniske problemer fra Leverandørens side (ikke-fungerende server, feil ved nøkkelgenerering).\n\n" +
			"<b>2. Avslag på refusjon</b>\n" +
			"Refusjon gis IKKE i følgende tilfeller:\n" +
			"• Mer enn 24 timer har gått siden kjøpet.\n" +
			"• Lav tilkoblingshastighet forårsaket av begrensninger hos brukerens leverandør eller dårlig mobildekning.\n" +
			"• Brukerens manglende evne til å sette opp tredjepartsprogramvare (Happ, Streisand-klienter osv.), forutsatt at nøkkelen selv er gyldig.\n" +
			"• Blokking av tilgang for brudd på reglene (spam, torrents, hackerangrep).\n" +
			"• „Jeg ombestemte meg“ eller „Jeg trenger det ikke lenger“.\n\n" +
			"<b>3. Refusjonsprosedyre</b>\n" +
			"3.1. For å be om refusjon, kontakt støtte.\n" +
			"3.2. I forespørselen oppgi: Konto-ID, betalingskvittering og en beskrivelse av det tekniske problemet (helst med et skjermbilde av feilen).\n" +
			"3.3. Refusjonen behandles via samme betalingsmetode innen 3–7 virkedager.",
	},
	language.Dansk: {
		Policy: "<b>REFUSJONSPOLITIK:</b>\n\n" +
			"<b>1. Betingelser for refusjon</b>\n" +
			"1.1. Da vi tilbyder digitale varer (adgangsnøgler) og tjenester, der forbruges ved brug, er refusjon kun mulig i undtagelsestilfælde.\n" +
			"1.2. Du har ret til at anmode om fuld refusjon inden for 24 timer efter købet, hvis du ikke kunne forbinde til tjenesten på grund af tekniske problemer fra Udbyderens side (ikke-fungerende server, fejl ved nøglegenerering).\n\n" +
			"<b>2. Afvisning af refusjon</b>\n" +
			"Refusjon gives IKKE i følgende tilfælde:\n" +
			"• Mere end 24 timer er gået siden købet.\n" +
			"• Lav forbindelseshastighed forårsaget af begrænsninger hos brugerens udbyder eller dårlig mobildækning.\n" +
			"• Brugerens manglende evne til at opsætte tredjepartssoftware (Happ, Streisand-klienter osv.), forudsat at nøglen selv er gyldig.\n" +
			"• Blokering af adgang for overtrædelse af regler (spam, torrents, hackerangreb).\n" +
			"• „Jeg har skiftet mening“ eller „Jeg har ikke brug for det længere“.\n\n" +
			"<b>3. Refusjonsprocedure</b>\n" +
			"3.1. For at anmode om refusjon, kontakt support.\n" +
			"3.2. I din anmodning angiv: Konto-ID, betalingskvittering og en beskrivelse af det tekniske problem (helst med et screenshot af fejlen).\n" +
			"3.3. Refusjonen behandles via samme betalingsmetode inden for 3–7 arbejdsdage.",
	},
	language.Español: {
		Policy: "<b>POLÍTICA DE REEMBOLSOS:</b>\n\n" +
			"<b>1. Condiciones de reembolso</b>\n" +
			"1.1. Dado que proporcionamos bienes digitales (claves de acceso) y servicios que se consumen en el momento de su uso, los reembolsos solo son posibles en casos excepcionales.\n" +
			"1.2. Tienes derecho a solicitar un reembolso completo dentro de las 24 horas posteriores a la compra si no pudiste conectarte al servicio por problemas técnicos dependientes del Proveedor (servidor no funcional, error en la generación de clave).\n\n" +
			"<b>2. Denegación de reembolsos</b>\n" +
			"Los reembolsos NO se realizan en los siguientes casos:\n" +
			"• Han pasado más de 24 horas desde la compra.\n" +
			"• Baja velocidad de conexión causada por limitaciones en el lado del proveedor del usuario o mala cobertura de red móvil.\n" +
			"• Incapacidad del usuario para configurar software de terceros (clientes Happ, Streisand, etc.), siempre que la clave sea válida.\n" +
			"• Bloqueo de acceso por violación de reglas (spam, torrents, ataques de hacking).\n" +
			"• „Cambié de opinión“ o „Ya no lo necesito“.\n\n" +
			"<b>3. Procedimiento de reembolso</b>\n" +
			"3.1. Para solicitar un reembolso, escribe al soporte.\n" +
			"3.2. En la solicitud indica: ID de cuenta, recibo de pago y descripción del problema técnico (preferiblemente con una captura de pantalla del error).\n" +
			"3.3. El reembolso se realiza por el mismo método de pago, dentro de 3–7 días hábiles.",
	},
	language.Français: {
		Policy: "<b>POLITIQUE DE REMBOURSEMENT:</b>\n\n" +
			"<b>1. Conditions de remboursement</b>\n" +
			"1.1. Puisque nous fournissons des biens numériques (clés d'accès) et des services consommés au moment de l'utilisation, les remboursements ne sont possibles que dans des cas exceptionnels.\n" +
			"1.2. Vous avez le droit de demander un remboursement intégral dans les 24 heures suivant l'achat si vous n'avez pas pu vous connecter au service en raison de problèmes techniques du côté du Fournisseur (serveur non fonctionnel, erreur de génération de clé).\n\n" +
			"<b>2. Refus de remboursement</b>\n" +
			"Les remboursements NE sont PAS effectués dans les cas suivants :\n" +
			"• Plus de 24 heures se sont écoulées depuis l'achat.\n" +
			"• Vitesse de connexion faible causée par des limitations du côté du fournisseur de l'utilisateur ou une mauvaise couverture réseau mobile.\n" +
			"• Incapacité de l'utilisateur à configurer un logiciel tiers (clients Happ, Streisand, etc.), à condition que la clé soit valide.\n" +
			"• Blocage d'accès pour violation des règles (spam, torrents, attaques de hacking).\n" +
			"• « J'ai changé d'avis » ou « Je n'en ai plus besoin ».\n\n" +
			"<b>3. Procédure de remboursement</b>\n" +
			"3.1. Pour formuler une demande de remboursement, contactez le support.\n" +
			"3.2. Dans la demande, indiquez : ID du compte, reçu de paiement et description du problème technique (de préférence avec une capture d'écran de l'erreur).\n" +
			"3.3. Le remboursement est effectué par le même moyen de paiement, dans un délai de 3 à 7 jours ouvrables.",
	},
	language.Português: {
		Policy: "<b>POLÍTICA DE REEMBOLSO:</b>\n\n" +
			"<b>1. Condições de reembolso</b>\n" +
			"1.1. Como fornecemos bens digitais (chaves de acesso) e serviços que são consumidos no momento do uso, os reembolsos só são possíveis em casos excepcionais.\n" +
			"1.2. Você tem o direito de solicitar um reembolso total dentro de 24 horas após a compra se não conseguiu se conectar ao serviço por razões técnicas dependentes do Provedor (servidor não funcional, erro de geração de chave).\n\n" +
			"<b>2. Negação de reembolso</b>\n" +
			"Os reembolsos NÃO são realizados nos seguintes casos:\n" +
			"• Passaram mais de 24 horas desde a compra.\n" +
			"• Baixa velocidade de conexão causada por limitações do lado do provedor do usuário ou má cobertura de rede móvel.\n" +
			"• Incapacidade do usuário de configurar software de terceiros (clientes Happ, Streisand, etc.), desde que a chave seja válida.\n" +
			"• Bloqueio de acesso por violação de regras (spam, torrents, ataques de hacking).\n" +
			"• „Mudei de ideia“ ou „Não preciso mais“.\n\n" +
			"<b>3. Procedimento de reembolso</b>\n" +
			"3.1. Para solicitar um reembolso, escreva para o suporte.\n" +
			"3.2. No pedido, indique: ID da conta, recibo de pagamento e descrição do problema técnico (de preferência com uma captura de tela do erro).\n" +
			"3.3. O reembolso é realizado pelo mesmo método de pagamento, dentro de 3–7 dias úteis.",
	},
	language.Italiana: {
		Policy: "<b>POLITICA DI RIMBORSO:</b>\n\n" +
			"<b>1. Condizioni di rimborso</b>\n" +
			"1.1. Poiché forniamo beni digitali (chiavi di accesso) e servizi che vengono consumati al momento dell'uso, i rimborsi sono possibili solo in casi eccezionali.\n" +
			"1.2. Hai il diritto di richiedere un rimborso completo entro 24 ore dall'acquisto se non sei riuscito a connetterti al servizio per motivi tecnici dipendenti dal Fornitore (server non funzionante, errore di generazione della chiave).\n\n" +
			"<b>2. Negazione di rimborso</b>\n" +
			"I rimborsi NON vengono effettuati nei seguenti casi:\n" +
			"• Sono passate più di 24 ore dall'acquisto.\n" +
			"• Bassa velocità di connessione causata da limitazioni dal lato del provider dell'utente o scarsa copertura di rete mobile.\n" +
			"• Incapacità dell'utente di configurare software di terze parti (client Happ, Streisand, ecc.), purché la chiave sia valida.\n" +
			"• Blocco dell'accesso per violazione delle regole (spam, torrent, attacchi di hacking).\n" +
			"• „Ho cambiato idea“ o „Non ne ho più bisogno“.\n\n" +
			"<b>3. Procedura di rimborso</b>\n" +
			"3.1. Per richiedere un rimborso, scrivi al supporto.\n" +
			"3.2. Nella richiesta indica: ID dell'account, ricevuta di pagamento e descrizione del problema tecnico (preferibilmente con uno screenshot dell'errore).\n" +
			"3.3. Il rimborso viene effettuato con lo stesso metodo di pagamento, entro 3–7 giorni lavorativi.",
	},
	language.Русский: {
		Policy: "<b>ПОЛИТИКА ВОЗВРАТА СРЕДСТВ:</b>\n\n" +
			"<b>1. Условия возврата</b>\n" +
			"1.1. Поскольку мы предоставляем цифровые товары (ключи доступа) и услуги, которые потребляются в момент использования, возврат средств возможен только в исключительных случаях.\n" +
			"1.2. Вы имеете право запросить полный возврат средств в течение 24 часов с момента покупки, если вам не удалось подключиться к сервису по техническим причинам, зависящим от Исполнителя (нерабочий сервер, ошибка генерации ключа).\n\n" +
			"<b>2. Отказ в возврате</b>\n" +
			"Возврат средств НЕ производится в следующих случаях:\n" +
			"• Прошло более 24 часов с момента покупки.\n" +
			"• Низкая скорость соединения, вызванная ограничениями на стороне провайдера пользователя или плохим качеством покрытия мобильной сети.\n" +
			"• Неумение пользователя настроить стороннее программное обеспечение (клиенты Happ, Streisand и др.), при условии, что сам ключ валиден.\n" +
			"• Блокировка доступа за нарушение правил (спам, торренты, хакерские атаки).\n" +
			"• «Я передумал» или «Мне больше не нужно».\n\n" +
			"<b>3. Процедура возврата</b>\n" +
			"3.1. Для оформления возврата напишите в службу поддержки.\n" +
			"3.2. В заявке укажите: ID аккаунта, чек об оплате и описание технической проблемы (желательно со скриншотом ошибки).\n" +
			"3.3. Возврат осуществляется тем же способом, которым была произведена оплата, в течение 3–7 рабочих дней.",
	},
	language.Українська: {
		Policy: "<b>ПОЛІТИКА ПОВЕРНЕННЯ КОШТІВ:</b>\n\n" +
			"<b>1. Умови повернення</b>\n" +
			"1.1. Оскільки ми надаємо цифрові товари (ключі доступу) та послуги, які споживаються в момент використання, повернення коштів можливе лише у виняткових випадках.\n" +
			"1.2. Ви маєте право запросити повне повернення коштів протягом 24 годин з моменту покупки, якщо вам не вдалося підключитися до сервісу через технічні причини, залежні від Виконавця (неробочий сервер, помилка генерації ключа).\n\n" +
			"<b>2. Відмова у поверненні</b>\n" +
			"Повернення коштів НЕ проводиться у таких випадках:\n" +
			"• Минуло більше 24 годин з моменту покупки.\n" +
			"• Низька швидкість з'єднання, спричинена обмеженнями з боку провайдера користувача або поганою якістю покриття мобільної мережі.\n" +
			"• Невміння користувача налаштувати стороннє програмне забезпечення (клієнти Happ, Streisand тощо), за умови, що сам ключ валідний.\n" +
			"• Блокування доступу за порушення правил (спам, торренти, хакерські атаки).\n" +
			"• «Я передумав» або «Мені більше не потрібно».\n\n" +
			"<b>3. Процедура повернення</b>\n" +
			"3.1. Для оформлення повернення напишіть у службу підтримки.\n" +
			"3.2. У заявці вкажіть: ID акаунта, чек про оплату та опис технічної проблеми (бажано зі скріншотом помилки).\n" +
			"3.3. Повернення здійснюється тим же способом, яким була проведена оплата, протягом 3–7 робочих днів.",
	},
	language.Polski: {
		Policy: "<b>POLITYKA ZWROTÓW:</b>\n\n" +
			"<b>1. Warunki zwrotów</b>\n" +
			"1.1. Ponieważ dostarczamy cyfrowe towary (klucze dostępu) i usługi, które są konsumowane w momencie użycia, zwroty są możliwe tylko w wyjątkowych przypadkach.\n" +
			"1.2. Masz prawo zażądać pełnego zwrotu w ciągu 24 godzin od zakupu, jeśli nie udało Ci się połączyć z usługą z powodów technicznych zależnych od Wykonawcy (nie działający serwer, błąd generowania klucza).\n\n" +
			"<b>2. Odmowa zwrotów</b>\n" +
			"Zwroty NIE są dokonywane w następujących przypadkach:\n" +
			"• Minęło więcej niż 24 godziny od zakupu.\n" +
			"• Niska prędkość połączenia spowodowana ograniczeniami po stronie dostawcy użytkownika lub słabym zasięgiem sieci mobilnej.\n" +
			"• Nieumiejętność użytkownika w konfiguracji oprogramowania stron trzecich (klienty Happ, Streisand itp.), pod warunkiem, że klucz jest ważny.\n" +
			"• Blokada dostępu za naruszenie zasad (spam, torrenty, ataki hakerskie).\n" +
			"• „Zmieniłem zdanie“ lub „Już mi niepotrzebne“.\n\n" +
			"<b>3. Procedura zwrotów</b>\n" +
			"3.1. Aby złożyć wniosek o zwrot, napisz do wsparcia.\n" +
			"3.2. W wniosku podaj: ID konta, paragon płatności i opis problemu technicznego (najlepiej ze zrzutem ekranu błędu).\n" +
			"3.3. Zwrot jest realizowany tą samą metodą płatności w ciągu 3–7 dni roboczych.",
	},
	language.Ceština: {
		Policy: "<b>POLITIKA VRACENÍ PENĚZ:</b>\n\n" +
			"<b>1. Podmínky vracení</b>\n" +
			"1.1. Protože poskytujeme digitální zboží (přístupové klíče) a služby, které se spotřebovávají v okamžiku použití, je vracení peněz možné pouze ve výjimečných případech.\n" +
			"1.2. Máte právo požádat o plné vrácení peněz do 24 hodin od nákupu, pokud se vám nepodařilo připojit k službě kvůli technickým problémům závislým na Dodavateli (nefungující server, chyba generování klíče).\n\n" +
			"<b>2. Odmítnutí vracení</b>\n" +
			"Vrácení peněz se NEPROVÁDÍ v následujících případech:\n" +
			"• Uplynulo více než 24 hodin od nákupu.\n" +
			"• Nízká rychlost spojení způsobená omezeními na straně poskytovatele uživatele nebo špatnou kvalitou pokrytí mobilní sítě.\n" +
			"• Nes chopnost uživatele nastavit software třetích stran (klienti Happ, Streisand atd.), za předpokladu, že klíč je platný.\n" +
			"• Blokování přístupu za porušení pravidel (spam, torrenty, hackerské útoky).\n" +
			"• „Rozmyslel jsem si to“ nebo „Už to nepotřebuji“.\n\n" +
			"<b>3. Postup vracení</b>\n" +
			"3.1. Pro podání žádosti o vrácení napište do podpory.\n" +
			"3.2. V žádosti uveďte: ID účtu, doklad o platbě a popis technického problému (ideálně se screenshotem chyby).\n" +
			"3.3. Vrácení se provádí stejným způsobem, jakým byla provedena platba, do 3–7 pracovních dnů.",
	},
	language.Български: {
		Policy: "<b>ПОЛИТИКА ЗА ВЪЗСТАНОВЯВАНЕ НА СРЕДСТВА:</b>\n\n" +
			"<b>1. Условия за възстановяване</b>\n" +
			"1.1. Тъй като предоставяме цифрови стоки (ключове за достъп) и услуги, които се консумират в момента на използване, възстановяването на средства е възможно само в изключителни случаи.\n" +
			"1.2. Имате право да поискате пълно възстановяване в рамките на 24 часа от покупката, ако не сте успели да се свържете със услугата поради технически причини, зависещи от Изпълнителя (неработещ сървър, грешка при генериране на ключ).\n\n" +
			"<b>2. Отказ от възстановяване</b>\n" +
			"Възстановяване на средства НЕ се извършва в следните случаи:\n" +
			"• Изминали са повече от 24 часа от покупката.\n" +
			"• Ниска скорост на връзката, причинена от ограничения от страна на доставчика на потребителя или лошо качество на покритието на мобилната мрежа.\n" +
			"• Невъзможност на потребителя да настрои софтуер на трети страни (клиенти Happ, Streisand и др.), при условие че самият ключ е валиден.\n" +
			"• Блокиране на достъпа за нарушение на правилата (спам, торенти, хакерски атаки).\n" +
			"• „Размислих си“ или „Вече не ми трябва“.\n\n" +
			"<b>3. Процедура за възстановяване</b>\n" +
			"3.1. За да оформите възстановяване, пишете в службата за поддръжка.\n" +
			"3.2. В заявката посочете: ID на акаунта, чек за плащане и описание на техническия проблем (желателно със скрийншот на грешката).\n" +
			"3.3. Възстановяването се извършва по същия начин, по който е извършено плащането, в рамките на 3–7 работни дни.",
	},
	language.Српски: {
		Policy: "<b>ПОЛИТИКА ПОВРАЋАЈА СРЕДСТАВА:</b>\n\n" +
			"<b>1. Услови за повраћај</b>\n" +
			"1.1. Пошто пружамо дигиталне производе (кључеве за приступ) и услуге које се троше у тренутку коришћења, повраћај средстава је могућ само у изузетним случајевима.\n" +
			"1.2. Имате право да затражите пун повраћај средстава у року од 24 сата од куповине, ако нисте успели да се повежете са услугом због техничких разлога зависних од Извођача (нерадни сервер, грешка у генерисању кључа).\n\n" +
			"<b>2. Одбијање повраћаја</b>\n" +
			"Повраћај средстава СЕ НЕ ВРШИ у следећим случајевима:\n" +
			"• Прошло је више од 24 сата од куповине.\n" +
			"• Ниска брзина везе узрокована ограничењима на страни провајдера корисника или лошим квалитетом покривености мобилне мреже.\n" +
			"• Немогућност корисника да подеси софтвер трећих страна (клијенти Happ, Streisand итд.), под условом да је сам кључ важећи.\n" +
			"• Блокирање приступа због кршења правила (спам, торенти, хакерски напади).\n" +
			"• „Предомислио сам се“ или „Више ми не треба“.\n\n" +
			"<b>3. Процедура повраћаја</b>\n" +
			"3.1. За оформљење повраћаја напишите у службу подршке.\n" +
			"3.2. У захтеву наведите: ID налога, рачун о плаћању и опис техничког проблема (пожељно са скриншотом грешке).\n" +
			"3.3. Повраћај се врши истим начином којим је извршено плаћање, у року од 3–7 радних дана.",
	},
	language.Hrvatski: {
		Policy: "<b>POLITIKA POVRATA SREDSTAVA:</b>\n\n" +
			"<b>1. Uvjeti povrata</b>\n" +
			"1.1. Budući da pružamo digitalne proizvode (ključeve za pristup) i usluge koje se troše u trenutku korištenja, povrat sredstava je moguć samo u iznimnim slučajevima.\n" +
			"1.2. Imate pravo zatražiti puni povrat sredstava u roku od 24 sata od kupnje ako niste uspjeli spojiti se na uslugu zbog tehničkih razloga ovisnih o Izvođaču (nefunkcionalni server, greška u generiranju ključa).\n\n" +
			"<b>2. Odbijanje povrata</b>\n" +
			"Povrat sredstava SE NE VRŠI u sljedećim slučajevima:\n" +
			"• Prošlo je više od 24 sata od kupnje.\n" +
			"• Niska brzina veze uzrokovana ograničenjima na strani korisničkog pružatelja ili lošom kvalitetom pokrivenosti mobilne mreže.\n" +
			"• Nemogućnost korisnika da postavi softver trećih strana (klijenti Happ, Streisand itd.), pod uvjetom da je sam ključ valjan.\n" +
			"• Blokiranje pristupa zbog kršenja pravila (spam, torrenti, hakerski napadi).\n" +
			"• „Predomislio sam se“ ili „Više mi ne treba“.\n\n" +
			"<b>3. Postupak povrata</b>\n" +
			"3.1. Za podnošenje zahtjeva za povrat napišite u službu podrške.\n" +
			"3.2. U zahtjevu navedite: ID računa, račun o plaćanju i opis tehničkog problema (poželjno sa screenshotom greške).\n" +
			"3.3. Povrat se vrši istim načinom kojim je izvršeno plaćanje, u roku od 3–7 radnih dana.",
	},
	language.Slovenčina: {
		Policy: "<b>POLITIKA VRÁTENIA PROSTRIEDKOV:</b>\n\n" +
			"<b>1. Podmienky vrátenia</b>\n" +
			"1.1. Keďže poskytujeme digitálne tovary (prístupové kľúče) a služby, ktoré sa spotrebúvajú v okamihu použitia, vrátenie prostriedkov je možné len vo výnimočných prípadoch.\n" +
			"1.2. Máte právo požiadať o plné vrátenie prostriedkov do 24 hodín od nákupu, ak sa vám nepodarilo pripojiť k službe kvôli technickým príčinám závislým od Vykonávateľa (nefungujúci server, chyba generovania kľúča).\n\n" +
			"<b>2. Odmietnutie vrátenia</b>\n" +
			"Vrátenie prostriedkov SA NEVYKONÁVA v nasledujúcich prípadoch:\n" +
			"• Uplynulo viac ako 24 hodín od nákupu.\n" +
			"• Nízka rýchlosť spojenia spôsobená obmedzeniami na strane poskytovateľa používateľa alebo zlou kvalitou pokrytia mobilnej siete.\n" +
			"• Nes chopnosť používateľa nastaviť softvér tretích strán (klienti Happ, Streisand atď.), za predpokladu, že samotný kľúč je platný.\n" +
			"• Blokovanie prístupu za porušenie pravidiel (spam, torrenty, hackerské útoky).\n" +
			"• „Rozmyslel som si to“ alebo „Už to nepotrebujem“.\n\n" +
			"<b>3. Postup vrátenia</b>\n" +
			"3.1. Na vybavenie vrátenia napíšte do podpory.\n" +
			"3.2. V žiadosti uveďte: ID účtu, doklad o platbe a popis technického problému (ideálne so screenshotom chyby).\n" +
			"3.3. Vrátenie sa vykonáva tým istým spôsobom, akým bola vykonaná platba, do 3–7 pracovných dní.",
	},
	language.Slovenski: {
		Policy: "<b>POLITIKA VRAČILA SREDSTEV:</b>\n\n" +
			"<b>1. Pogoji vračila</b>\n" +
			"1.1. Ker zagotavljamo digitalne izdelke (ključe za dostop) in storitve, ki se porabijo v trenutku uporabe, je vračilo sredstev možno samo v izjemnih primerih.\n" +
			"1.2. Imate pravico zahtevati polno vračilo v roku 24 ur od nakupa, če se niste mogli povezati s storitvijo zaradi tehničnih razlogov, odvisnih od Izvajalca (ne delujoč strežnik, napaka pri generiranju ključa).\n\n" +
			"<b>2. Zavrnitev vračila</b>\n" +
			"Vračilo sredstev SE NE IZVAJA v naslednjih primerih:\n" +
			"• Minilo je več kot 24 ur od nakupa.\n" +
			"• Nizka hitrost povezave, povzročena z omejitvami na strani ponudnika uporabnika ali slabo kakovostjo pokritosti mobilnega omrežja.\n" +
			"• Nezmožnost uporabnika nastaviti programske opreme tretjih oseb (klienti Happ, Streisand itd.), pod pogojem, da je ključ sam veljaven.\n" +
			"• Blokiranje dostopa zaradi kršitve pravil (spam, torrenti, hekerski napadi).\n" +
			"• „Premislil sem si“ ali „Ne potrebujem več“.\n\n" +
			"<b>3. Postopek vračila</b>\n" +
			"3.1. Za zahtevo vračila pišite v podporo.\n" +
			"3.2. V zahtevku navedite: ID računa, potrdilo o plačilu in opis tehničnega problema (po možnosti s posnetkom zaslona napake).\n" +
			"3.3. Vračilo se izvede na enak način, kot je bilo izvedeno plačilo, v roku 3–7 delovnih dni.",
	},
	language.Lietùvių: {
		Policy: "<b>GRĄŽINIMO POLITIKA:</b>\n\n" +
			"<b>1. Grąžinimo sąlygos</b>\n" +
			"1.1. Kadangi teikiame skaitmenines prekes (prieigos raktus) ir paslaugas, kurios suvartojamos naudojimo metu, grąžinimas galimas tik išskirtiniais atvejais.\n" +
			"1.2. Jūs turite teisę prašyti pilno grąžinimo per 24 valandas nuo pirkimo, jei nepavyko prisijungti prie paslaugos dėl techninių priežasčių, priklausančių nuo Vykdytojo (neveikiantis serveris, rakto generavimo klaida).\n\n" +
			"<b>2. Grąžinimo atmetimas</b>\n" +
			"Grąžinimas NEATLIKIMAS šiais atvejais:\n" +
			"• Praėjo daugiau nei 24 valandos nuo pirkimo.\n" +
			"• Žema ryšio greitis, sukeltas apribojimų vartotojo tiekėjo pusėje ar prastos mobiliojo tinklo aprėpties.\n" +
			"• Vartotojo nesugebėjimas nustatyti trečiųjų šalių programinės įrangos (Happ, Streisand klientai ir kt.), jei pats raktas yra galiojantis.\n" +
			"• Prieigos blokavimas už taisyklių pažeidimą (spam, torrentai, hakerių atakos).\n" +
			"• „Apsigalvojau“ arba „Man daugiau nereikia“.\n\n" +
			"<b>3. Grąžinimo procedūra</b>\n" +
			"3.1. Norėdami prašyti grąžinimo, rašykite į paramą.\n" +
			"3.2. Prašyme nurodykite: Paskyros ID, mokėjimo čekį ir techninės problemos aprašymą (pageidautina su klaidos ekrano nuotrauka).\n" +
			"3.3. Grąžinimas atliekamas tuo pačiu mokėjimo būdu per 3–7 darbo dienas.",
	},
	language.Latviešu: {
		Policy: "<b>ATMKSAS POLITIKA:</b>\n\n" +
			"<b>1. Atmaksa nosacījumi</b>\n" +
			"1.1. Tā kā mēs nodrošinām digitālos produktus (pieejas atslēgas) un pakalpojumus, kas tiek patērēti lietošanas brīdī, atmaksa ir iespējama tikai izņēmuma gadījumos.\n" +
			"1.2. Jums ir tiesības pieprasīt pilnu atmaksu 24 stundu laikā pēc pirkuma, ja neizdevās pieslēgties pakalpojumam tehnisku iemeslu dēļ, kas atkarīgi no Izpildītāja (nedarbojošs serveris, atslēgas ģenerēšanas kļūda).\n\n" +
			"<b>2. Atmaksa atteikums</b>\n" +
			"Atmaksa NETIEK veikta šādos gadījumos:\n" +
			"• Pagājušas vairāk nekā 24 stundas kopš pirkuma.\n" +
			"• Zema savienojuma ātrums, ko izraisa ierobežojumi lietotāja pakalpojumu sniedzēja pusē vai slikta mobilo tīklu pārklājuma kvalitāte.\n" +
			"• Lietotāja nespēja iestatīt trešo pušu programmatūru (Happ, Streisand klienti u.tml.), ja pati atslēga ir derīga.\n" +
			"• Pieejas bloķēšana par noteikumu pārkāpumu (spam, torrenti, hakeru uzbrukumi).\n" +
			"• „Pārdomāju“ vai „Man vairs nevajag“.\n\n" +
			"<b>3. Atmaksa procedūra</b>\n" +
			"3.1. Lai pieteiktu atmaksu, rakstiet atbalstam.\n" +
			"3.2. Pieteikumā norādiet: Kontu ID, maksājuma čeku un tehniskās problēmas aprakstu (vēlams ar kļūdas ekrānšāviņu).\n" +
			"3.3. Atmaksa tiek veikta ar to pašu maksājuma metodi 3–7 darba dienu laikā.",
	},
	language.Eesti: {
		Policy: "<b>TAGASTUSPOLIITIKA:</b>\n\n" +
			"<b>1. Tagastuse tingimused</b>\n" +
			"1.1. Kuna pakume digitaalseid tooteid (juurdepääsu võtmeid) ja teenuseid, mis tarbitakse kasutamise hetkel, on tagastus võimalik ainult erandjuhtudel.\n" +
			"1.2. Teil on õigus taotleda täielikku tagastust 24 tunni jooksul pärast ostu, kui teil ei õnnestunud teenusega ühendust saada tehniliste põhjuste tõttu, mis sõltuvad Teenuseosutajast (mittetoimiv server, võtme genereerimise viga).\n\n" +
			"<b>2. Tagastuse keelamine</b>\n" +
			"Tagastust EI TEHTA järgmistel juhtudel:\n" +
			"• On möödunud rohkem kui 24 tundi ostust.\n" +
			"• Madal ühenduse kiirus, mis on põhjustatud piirangutest kasutaja teenusepakkuja poolt või halva mobiilivõrgu katvusest.\n" +
			"• Kasutaja võimetus kolmanda osapoole tarkvara seadistada (Happ, Streisand kliendid jne), tingimusel, et võti ise on kehtiv.\n" +
			"• Juurdepääsu blokeerimine reeglite rikkumise eest (spam, torrentid, häkkimise rünnakud).\n" +
			"• „Muutsin meelt“ või „Mul pole enam vaja“.\n\n" +
			"<b>3. Tagastuse protseduur</b>\n" +
			"3.1. Tagastuse vormistamiseks kirjutage toele.\n" +
			"3.2. Taotluses märkige: Konto ID, maksekviitung ja tehnilise probleemi kirjeldus (soovitavalt koos vea ekraanipildiga).\n" +
			"3.3. Tagastus tehakse sama makseviisiga 3–7 tööpäeva jooksul.",
	},
	language.Suomi: {
		Policy: "<b>PALAUTUSKÄYTÄNTÖ:</b>\n\n" +
			"<b>1. Palautuksen ehdot</b>\n" +
			"1.1. Koska tarjoamme digitaalisia tuotteita (pääsyavaimia) ja palveluita, jotka kulutetaan käytön hetkellä, palautus on mahdollista vain poikkeustapauksissa.\n" +
			"1.2. Sinulla on oikeus pyytää täydellistä palautusta 24 tunnin sisällä ostoksesta, jos et onnistunut yhdistämään palveluun teknisistä syistä, jotka riippuvat Toimittajasta (toimimaton palvelin, avaimen generointivirhe).\n\n" +
			"<b>2. Palautuksen hylkääminen</b>\n" +
			"Palautusta EI TEHDÄ seuraavissa tapauksissa:\n" +
			"• On kulunut yli 24 tuntia ostoksesta.\n" +
			"• Matala yhteyden nopeus, joka johtuu rajoituksista käyttäjän tarjoajan puolelta tai huonosta mobiiliverkon kattavuudesta.\n" +
			"• Käyttäjän kyvyttömyys määrittää kolmannen osapuolen ohjelmistoa (Happ, Streisand-asiakkaat jne.), edellyttäen että avain itse on kelvollinen.\n" +
			"• Pääsyn estäminen sääntöjen rikkomisen vuoksi (spam, torrentit, hakkerointihyökkäykset).\n" +
			"• „Muutin mieltäni“ tai „En tarvitse enää“.\n\n" +
			"<b>3. Palautusmenettely</b>\n" +
			"3.1. Palautuksen pyytämiseksi kirjoita tukeen.\n" +
			"3.2. Pyynnössä ilmoita: Tilin ID, maksukuitti ja teknisen ongelman kuvaus (mieluiten virheen kuvakaappauksella).\n" +
			"3.3. Palautus suoritetaan samalla maksutavalla 3–7 arkipäivän sisällä.",
	},
	language.Ελληνικά: {
		Policy: "<b>ΠΟΛΙΤΙΚΗ ΕΠΙΣΤΡΟΦΩΝ:</b>\n\n" +
			"<b>1. Όροι επιστροφών</b>\n" +
			"1.1. Εφόσον παρέχουμε ψηφιακά αγαθά (κλειδιά πρόσβασης) και υπηρεσίες που καταναλώνονται τη στιγμή της χρήσης, οι επιστροφές είναι δυνατές μόνο σε εξαιρετικές περιπτώσεις.\n" +
			"1.2. Έχετε το δικαίωμα να ζητήσετε πλήρη επιστροφή μέσα σε 24 ώρες από την αγορά, αν δεν καταφέρατε να συνδεθείτε στην υπηρεσία λόγω τεχνικών λόγων που εξαρτώνται από τον Παροχέα (μη λειτουργικός διακομιστής, σφάλμα δημιουργίας κλειδιού).\n\n" +
			"<b>2. Άρνηση επιστροφών</b>\n" +
			"Οι επιστροφές ΔΕΝ πραγματοποιούνται στις εξής περιπτώσεις:\n" +
			"• Έχουν περάσει περισσότερες από 24 ώρες από την αγορά.\n" +
			"• Χαμηλή ταχύτητα σύνδεσης προκαλούμενη από περιορισμούς από την πλευρά του παρόχου του χρήστη ή κακή ποιότητα κάλυψης κινητού δικτύου.\n" +
			"• Αδυναμία του χρήστη να ρυθμίσει λογισμικό τρίτων (πελάτες Happ, Streisand κ.λπ.), υπό την προϋπόθεση ότι το κλειδί είναι έγκυρο.\n" +
			"• Αποκλεισμός πρόσβασης για παραβίαση κανόνων (spam, torrents, επιθέσεις hacking).\n" +
			"• «Άλλαξα γνώμη» ή «Δεν το χρειάζομαι πια».\n\n" +
			"<b>3. Διαδικασία επιστροφών</b>\n" +
			"3.1. Για να υποβάλετε αίτημα επιστροφής, γράψτε στην υποστήριξη.\n" +
			"3.2. Στο αίτημα αναφέρετε: ID λογαριασμού, απόδειξη πληρωμής και περιγραφή του τεχνικού προβλήματος (κατά προτίμηση με screenshot του σφάλματος).\n" +
			"3.3. Η επιστροφή πραγματοποιείται με τον ίδιο τρόπο πληρωμής, μέσα σε 3–7 εργάσιμες ημέρες.",
	},
	language.Română: {
		Policy: "<b>POLITICA DE RAMBURSARE:</b>\n\n" +
			"<b>1. Condiții de rambursare</b>\n" +
			"1.1. Deoarece oferim bunuri digitale (chei de acces) și servicii care sunt consumate în momentul utilizării, rambursările sunt posibile doar în cazuri excepționale.\n" +
			"1.2. Aveți dreptul să solicitați o rambursare completă în termen de 24 de ore de la achiziție, dacă nu ați reușit să vă conectați la serviciu din motive tehnice dependente de Furnizor (server nefuncțional, eroare de generare a cheii).\n\n" +
			"<b>2. Refuzul rambursării</b>\n" +
			"Rambursările NU se efectuează în următoarele cazuri:\n" +
			"• Au trecut mai mult de 24 de ore de la achiziție.\n" +
			"• Viteză scăzută de conexiune cauzată de limitări din partea furnizorului utilizatorului sau acoperire slabă a rețelei mobile.\n" +
			"• Incapacitatea utilizatorului de a configura software terț (clienți Happ, Streisand etc.), cu condiția ca cheia să fie validă.\n" +
			"• Blocarea accesului pentru încălcarea regulilor (spam, torrente, atacuri de hacking).\n" +
			"• „M-am răzgândit“ sau „Nu mai am nevoie“.\n\n" +
			"<b>3. Procedura de rambursare</b>\n" +
			"3.1. Pentru a solicita o rambursare, scrieți la suport.\n" +
			"3.2. În cerere indicați: ID-ul contului, chitanța de plată și descrierea problemei tehnice (preferabil cu un screenshot al erorii).\n" +
			"3.3. Rambursarea se efectuează prin aceeași metodă de plată, în termen de 3–7 zile lucrătoare.",
	},
	language.Magyar: {
		Policy: "<b>VISSZATÉRÍTÉSI SZABÁLYZAT:</b>\n\n" +
			"<b>1. Visszatérítési feltételek</b>\n" +
			"1.1. Mivel digitális termékeket (hozzáférési kulcsokat) és szolgáltatásokat nyújtunk, amelyek használatkor elfogyasztódnak, a visszatérítés csak kivételes esetekben lehetséges.\n" +
			"1.2. Jogosult teljes visszatérítést kérni a vásárlástól számított 24 órán belül, ha nem sikerült csatlakozni a szolgáltatáshoz a Szolgáltatótól függő technikai okok miatt (nem működő szerver, kulcsgenerálási hiba).\n\n" +
			"<b>2. Visszatérítés elutasítása</b>\n" +
			"Visszatérítés NEM történik a következő esetekben:\n" +
			"• Több mint 24 óra telt el a vásárlás óta.\n" +
			"• Alacsony csatlakozási sebesség, amit a felhasználó szolgáltatójának korlátozásai vagy gyenge mobilhálózati lefedettség okoz.\n" +
			"• A felhasználó képtelensége harmadik féltől származó szoftver beállítására (Happ, Streisand kliensek stb.), feltéve, hogy a kulcs érvényes.\n" +
			"• Hozzáférés blokkolása szabálysértés miatt (spam, torrentek, hackertámadások).\n" +
			"• „Meggondoltam magam“ vagy „Már nincs rá szükségem“.\n\n" +
			"<b>3. Visszatérítési eljárás</b>\n" +
			"3.1. Visszatérítés kéréséhez írjon a támogatásnak.\n" +
			"3.2. A kérésben tüntesse fel: Fiók ID, fizetési bizonylat és a technikai probléma leírása (lehetőleg a hiba képernyőképével).\n" +
			"3.3. A visszatérítés ugyanazon fizetési móddal történik, 3–7 munkanapon belül.",
	},
	language.Arabic: {
		Policy: "<b>سياسة الاسترداد:</b>\n\n" +
			"<b>1. شروط الاسترداد</b>\n" +
			"1.1. بما أننا نقدم سلعًا رقمية (مفاتيح الوصول) وخدمات تُستهلك في لحظة الاستخدام، فإن الاسترداد ممكن فقط في حالات استثنائية.\n" +
			"1.2. لديك الحق في طلب استرداد كامل خلال 24 ساعة من الشراء إذا لم تتمكن من الاتصال بالخدمة بسبب أسباب فنية تعتمد على المزود (خادم غير عامل، خطأ في توليد المفتاح).\n\n" +
			"<b>2. رفض الاسترداد</b>\n" +
			"لا يتم الاسترداد في الحالات التالية:\n" +
			"• مر أكثر من 24 ساعة منذ الشراء.\n" +
			"• سرعة اتصال منخفضة ناتجة عن قيود من جانب مزود المستخدم أو جودة تغطية شبكة الهاتف المحمول السيئة.\n" +
			"• عدم قدرة المستخدم على إعداد برمجيات الطرف الثالث (عملاء Happ، Streisand إلخ)، بشرط أن يكون المفتاح صالحًا.\n" +
			"• حظر الوصول بسبب انتهاك القواعد (سبام، تورنت، هجمات هكر).\n" +
			"• «غيرت رأيي» أو «لم أعد بحاجة إليه».\n\n" +
			"<b>3. إجراء الاسترداد</b>\n" +
			"3.1. لطلب الاسترداد، اكتب إلى الدعم.\n" +
			"3.2. في الطلب حدد: معرف الحساب، إيصال الدفع ووصف المشكلة الفنية (يفضل مع لقطة شاشة للخطأ).\n" +
			"3.3. يتم الاسترداد بنفس طريقة الدفع خلال 3–7 أيام عمل.",
	},
	language.Farsi: {
		Policy: "<b>سیاست بازپرداخت:</b>\n\n" +
			"<b>1. شرایط بازپرداخت</b>\n" +
			"1.1. از آنجایی که ما کالاهای دیجیتال (کلیدهای دسترسی) و خدمات ارائه می‌دهیم که در لحظه استفاده مصرف می‌شوند، بازپرداخت فقط در موارد استثنایی ممکن است.\n" +
			"1.2. شما حق دارید ظرف ۲۴ ساعت پس از خرید، بازپرداخت کامل درخواست کنید اگر نتوانستید به سرویس متصل شوید به دلایل فنی وابسته به مجری (سرور غیرفعال، خطای تولید کلید).\n\n" +
			"<b>2. رد بازپرداخت</b>\n" +
			"بازپرداخت انجام نمی‌شود در موارد زیر:\n" +
			"• بیش از ۲۴ ساعت از خرید گذشته است.\n" +
			"• سرعت اتصال پایین ناشی از محدودیت‌های سمت ارائه‌دهنده کاربر یا کیفیت پوشش ضعیف شبکه موبایل.\n" +
			"• ناتوانی کاربر در تنظیم نرم‌افزارهای شخص ثالث (کلاینت‌های Happ، Streisand و غیره)، به شرط اینکه خود کلید معتبر باشد.\n" +
			"• مسدود شدن دسترسی به دلیل نقض قوانین (اسپم، تورنت‌ها، حملات هکری).\n" +
			"• «پشیمان شدم» یا «دیگه لازم ندارم».\n\n" +
			"<b>3. روند بازپرداخت</b>\n" +
			"3.1. برای ثبت بازپرداخت، به پشتیبانی بنویسید.\n" +
			"3.2. در درخواست مشخص کنید: شناسه حساب، رسید پرداخت و توصیف مشکل فنی (ترجیحاً با اسکرین‌شات خطا).\n" +
			"3.3. بازپرداخت با همان روش پرداخت انجام می‌شود، ظرف ۳–۷ روز کاری.",
	},
	language.Türkçe: {
		Policy: "<b>İADE POLİTİKASI:</b>\n\n" +
			"<b>1. İade Şartları</b>\n" +
			"1.1. Dijital ürünler (erişim anahtarları) ve kullanım anında tüketilen hizmetler sunduğumuz için iadeler sadece istisnai durumlarda mümkündür.\n" +
			"1.2. Satın alımdan sonraki 24 saat içinde tam iade talep etme hakkınız vardır eğer hizmete Sağlayıcı'ya bağlı teknik nedenlerle (çalışmayan sunucu, anahtar üretme hatası) bağlanamadıysanız.\n\n" +
			"<b>2. İade Reddi</b>\n" +
			"İadeler yapılmaz şu durumlarda:\n" +
			"• Satın alımdan 24 saatten fazla zaman geçti.\n" +
			"• Kullanıcının sağlayıcısındaki kısıtlamalar veya kötü mobil ağ kapsama kalitesinden kaynaklanan düşük bağlantı hızı.\n" +
			"• Kullanıcının üçüncü taraf yazılımları (Happ, Streisand istemcileri vb.) kuramaması, anahtarın geçerli olması koşuluyla.\n" +
			"• Kurallara uymama nedeniyle erişim engelleme (spam, torrentler, hacker saldırıları).\n" +
			"• „Fikrimi değiştirdim“ veya „Artık ihtiyacım yok“.\n\n" +
			"<b>3. İade Prosedürü</b>\n" +
			"3.1. İade talebi için desteğe yazın.\n" +
			"3.2. Talebinizde belirtin: Hesap ID'si, ödeme makbuzu ve teknik sorunun açıklaması (tercihen hata ekran görüntüsüyle).\n" +
			"3.3. İade aynı ödeme yöntemiyle 3–7 iş günü içinde yapılır.",
	},
	language.Hebrew: {
		Policy: "<b>מדיניות החזרים:</b>\n\n" +
			"<b>1. תנאי החזרים</b>\n" +
			"1.1. מאחר שאנו מספקים מוצרים דיגיטליים (מפתחות גישה) ושירותים שנצרכים ברגע השימוש, החזרים אפשריים רק במקרים יוצאי דופן.\n" +
			"1.2. יש לך זכות לבקש החזר מלא בתוך 24 שעות מרגע הרכישה אם לא הצלחת להתחבר לשירות מסיבות טכניות התלויות בספק (שרת לא עובד, שגיאת יצירת מפתח).\n\n" +
			"<b>2. סירוב החזרים</b>\n" +
			"החזרים לא מתבצעים במקרים הבאים:\n" +
			"• עברו יותר מ-24 שעות מרגע הרכישה.\n" +
			"• מהירות חיבור נמוכה הנגרמת ממגבלות מצד ספק המשתמש או כיסוי רשת סלולרית גרוע.\n" +
			"• חוסר יכולת של המשתמש להגדיר תוכנה של צד שלישי (לקוחות Happ, Streisand וכו'), בתנאי שהמפתח עצמו תקף.\n" +
			"• חסימת גישה בגלל הפרת חוקים (ספאם, טורנטים, התקפות האקינג).\n" +
			"• „שיניתי דעתי“ או „אני לא צריך את זה יותר“.\n\n" +
			"<b>3. הליך החזרים</b>\n" +
			"3.1. כדי לבקש החזר, כתוב לתמיכה.\n" +
			"3.2. בבקשה ציין: מזהה חשבון, קבלה על תשלום ותיאור הבעיה הטכנית (רצוי עם צילום מסך של השגיאה).\n" +
			"3.3. ההחזר מתבצע באותה שיטת תשלום בתוך 3–7 ימי עבודה.",
	},
	language.ZH中文: {
		Policy: "<b>退款政策：</b>\n\n" +
			"<b>1. 退款条件</b>\n" +
			"1.1. 由于我们提供数字商品（访问密钥）和在使用时即消耗的服务，退款仅在例外情况下可能。\n" +
			"1.2. 您有权在购买后24小时内要求全额退款，如果由于提供商方面的技术原因（非工作服务器、密钥生成错误）无法连接到服务。\n\n" +
			"<b>2. 拒绝退款</b>\n" +
			"在以下情况下不进行退款：\n" +
			"• 购买后已超过24小时。\n" +
			"• 低连接速度，由用户提供商方面的限制或移动网络覆盖质量差引起。\n" +
			"• 用户无法设置第三方软件（Happ、Streisand客户端等），前提是密钥本身有效。\n" +
			"• 因违反规则（垃圾邮件、种子、黑客攻击）而被封锁访问。\n" +
			"• “我改变了主意”或“我不再需要了”。\n\n" +
			"<b>3. 退款程序</b>\n" +
			"3.1. 要申请退款，请联系支持。\n" +
			"3.2. 在申请中指定：账户ID、付款收据和技术问题的描述（最好附上错误截图）。\n" +
			"3.3. 退款通过相同的付款方式在3–7个工作日内进行。",
	},
	language.JA日本語: {
		Policy: "<b>返金ポリシー：</b>\n\n" +
			"<b>1. 返金の条件</b>\n" +
			"1.1. 当社はデジタル商品（アクセスキー）と使用時に消費されるサービスを提供しているため、返金は例外的な場合にのみ可能です。\n" +
			"1.2. 購入後24時間以内に、提供者の技術的な理由（動作しないサーバー、キーの生成エラー）でサービスに接続できなかった場合、完全返金を要求する権利があります。\n\n" +
			"<b>2. 返金の拒否</b>\n" +
			"返金は以下の場合に行われません：\n" +
			"• 購入から24時間以上経過した場合。\n" +
			"• ユーザーのプロバイダー側の制限やモバイルネットワークのカバレッジの質の悪さによる低速接続。\n" +
			"• ユーザーがサードパーティソフトウェア（Happ、Streisandクライアントなど）を設定できない場合、キーが有効である限り。\n" +
			"• ルール違反（スパム、トレント、ハッキング攻撃）によるアクセスブロック。\n" +
			"• 「気が変わった」や「もう必要ない」。\n\n" +
			"<b>3. 返金の手順</b>\n" +
			"3.1. 返金を申請するには、サポートに連絡してください。\n" +
			"3.2. 申請に記載：アカウントID、支払い領収書、技術的な問題の説明（エラーのスクリーンショットを推奨）。\n" +
			"3.3. 返金は同じ支払い方法で3–7営業日以内に行われます。",
	},
	language.KO한국어: {
		Policy: "<b>환불 정책:</b>\n\n" +
			"<b>1. 환불 조건</b>\n" +
			"1.1. 우리는 디지털 상품(액세스 키)과 사용 시 소비되는 서비스를 제공하므로, 환불은 예외적인 경우에만 가능합니다.\n" +
			"1.2. 구매 후 24시간 이내에 제공자 측의 기술적 이유(작동하지 않는 서버, 키 생성 오류)로 서비스에 연결할 수 없었던 경우 전체 환불을 요청할 권리가 있습니다.\n\n" +
			"<b>2. 환불 거부</b>\n" +
			"환불은 다음 경우에 이루어지지 않습니다:\n" +
			"• 구매 후 24시간 이상 경과한 경우.\n" +
			"• 사용자 제공자 측의 제한이나 모바일 네트워크 커버리지 품질이 낮아서 발생한 낮은 연결 속도.\n" +
			"• 사용자가 타사 소프트웨어(Happ, Streisand 클라이언트 등)를 설정할 수 없는 경우, 키 자체가 유효한 경우.\n" +
			"• 규칙 위반(스팸, 토렌트, 해킹 공격)으로 액세스 차단.\n" +
			"• “마음이 바뀌었어요” 또는 “더 이상 필요 없어요”.\n\n" +
			"<b>3. 환불 절차</b>\n" +
			"3.1. 환불을 신청하려면 지원팀에 문의하세요.\n" +
			"3.2. 신청서에 지정: 계정 ID, 결제 영수증, 기술 문제 설명(오류 스크린샷 바람직).\n" +
			"3.3. 환불은 동일한 결제 방법으로 3–7 영업일 이내에 이루어집니다.",
	},
	language.TiếngViệt: {
		Policy: "<b>CHÍNH SÁCH HOÀN TIỀN:</b>\n\n" +
			"<b>1. Điều kiện hoàn tiền</b>\n" +
			"1.1. Vì chúng tôi cung cấp hàng hóa kỹ thuật số (khóa truy cập) và dịch vụ được tiêu thụ ngay khi sử dụng, hoàn tiền chỉ có thể trong các trường hợp đặc biệt.\n" +
			"1.2. Bạn có quyền yêu cầu hoàn tiền đầy đủ trong vòng 24 giờ sau khi mua nếu bạn không thể kết nối với dịch vụ do lý do kỹ thuật phụ thuộc vào Nhà cung cấp (máy chủ không hoạt động, lỗi tạo khóa).\n\n" +
			"<b>2. Từ chối hoàn tiền</b>\n" +
			"Hoàn tiền KHÔNG được thực hiện trong các trường hợp sau:\n" +
			"• Đã qua hơn 24 giờ kể từ khi mua.\n" +
			"• Tốc độ kết nối thấp do hạn chế từ phía nhà cung cấp của người dùng hoặc chất lượng phủ sóng mạng di động kém.\n" +
			"• Người dùng không thể thiết lập phần mềm bên thứ ba (khách hàng Happ, Streisand v.v.), miễn là khóa chính nó hợp lệ.\n" +
			"• Chặn truy cập do vi phạm quy tắc (spam, torrent, tấn công hack).\n" +
			"• “Tôi thay đổi ý kiến” hoặc “Tôi không cần nữa”.\n\n" +
			"<b>3. Quy trình hoàn tiền</b>\n" +
			"3.1. Để yêu cầu hoàn tiền, viết cho hỗ trợ.\n" +
			"3.2. Trong yêu cầu chỉ định: ID tài khoản, biên nhận thanh toán và mô tả vấn đề kỹ thuật (tốt nhất với ảnh chụp màn hình lỗi).\n" +
			"3.3. Hoàn tiền được thực hiện bằng cùng phương thức thanh toán trong vòng 3–7 ngày làm việc.",
	},
	language.THภาษาไทย: {
		Policy: "<b>นโยบายการคืนเงิน:</b>\n\n" +
			"<b>1. เงื่อนไขการคืนเงิน</b>\n" +
			"1.1. เนื่องจากเราให้บริการสินค้าดิจิทัล (คีย์การเข้าถึง) และบริการที่ถูกบริโภคในขณะใช้งาน การคืนเงินเป็นไปได้เฉพาะในกรณีพิเศษเท่านั้น.\n" +
			"1.2. คุณมีสิทธิ์ขอคืนเงินเต็มจำนวนภายใน 24 ชั่วโมงหลังจากการซื้อ หากคุณไม่สามารถเชื่อมต่อกับบริการเนื่องจากปัญหาทางเทคนิคที่ขึ้นอยู่กับผู้ให้บริการ (เซิร์ฟเวอร์ไม่ทำงาน ข้อผิดพลาดในการสร้างคีย์).\n\n" +
			"<b>2. การปฏิเสธการคืนเงิน</b>\n" +
			"ไม่มีการคืนเงินในกรณีดังต่อไปนี้:\n" +
			"• ผ่านไปมากกว่า 24 ชั่วโมงนับจากวันที่ซื้อ.\n" +
			"• ความเร็วการเชื่อมต่อต่ำที่เกิดจากข้อจำกัดจากฝั่งผู้ให้บริการของผู้ใช้หรือคุณภาพการครอบคลุมเครือข่ายมือถือที่ไม่ดี.\n" +
			"• ผู้ใช้ไม่สามารถตั้งค่าซอฟต์แวร์บุคคลที่สาม (ไคลเอ็นต์ Happ, Streisand ฯลฯ) ได้ โดยที่คีย์เองถูกต้อง.\n" +
			"• การบล็อกการเข้าถึงเนื่องจากการละเมิดกฎ (สแปม, ทอร์เรนต์, การโจมตีแฮกเกอร์).\n" +
			"• “ฉันเปลี่ยนใจแล้ว” หรือ “ฉันไม่ต้องการอีกต่อไป”.\n\n" +
			"<b>3. ขั้นตอนการคืนเงิน</b>\n" +
			"3.1. เพื่อขอคืนเงิน เขียนไปยังฝ่ายสนับสนุน.\n" +
			"3.2. ในคำขอระบุ: ID บัญชี, ใบเสร็จรับเงินการชำระเงิน และคำอธิบายปัญหาทางเทคนิค (ควรมีภาพหน้าจอของข้อผิดพลาด).\n" +
			"3.3. การคืนเงินดำเนินการด้วยวิธีการชำระเงินเดียวกันภายใน 3–7 วันทำการ.",
	},
	language.BahasaIndonesia: {
		Policy: "<b>KEBIJAKAN PENGEMBALIAN DANA:</b>\n\n" +
			"<b>1. Syarat Pengembalian Dana</b>\n" +
			"1.1. Karena kami menyediakan barang digital (kunci akses) dan layanan yang dikonsumsi pada saat penggunaan, pengembalian dana hanya mungkin dalam kasus luar biasa.\n" +
			"1.2. Anda berhak meminta pengembalian dana penuh dalam waktu 24 jam setelah pembelian jika Anda tidak dapat terhubung ke layanan karena alasan teknis yang bergantung pada Penyedia (server tidak berfungsi, kesalahan pembuatan kunci).\n\n" +
			"<b>2. Penolakan Pengembalian Dana</b>\n" +
			"Pengembalian dana TIDAK dilakukan dalam kasus berikut:\n" +
			"• Lebih dari 24 jam telah berlalu sejak pembelian.\n" +
			"• Kecepatan koneksi rendah yang disebabkan oleh batasan dari sisi penyedia pengguna atau kualitas jangkauan jaringan seluler yang buruk.\n" +
			"• Ketidakmampuan pengguna untuk mengatur perangkat lunak pihak ketiga (klien Happ, Streisand dll.), asalkan kunci itu sendiri valid.\n" +
			"• Pemblokiran akses karena pelanggaran aturan (spam, torrent, serangan hacking).\n" +
			"• “Saya berubah pikiran” atau “Saya tidak butuh lagi”.\n\n" +
			"<b>3. Prosedur Pengembalian Dana</b>\n" +
			"3.1. Untuk mengajukan pengembalian dana, tulis ke dukungan.\n" +
			"3.2. Dalam permintaan sebutkan: ID akun, bukti pembayaran dan deskripsi masalah teknis (sebaiknya dengan screenshot kesalahan).\n" +
			"3.3. Pengembalian dana dilakukan dengan metode pembayaran yang sama dalam waktu 3–7 hari kerja.",
	},
	language.BahasaMelayu: {
		Policy: "<b>DASAR PEMBAYARAN BALIK:</b>\n\n" +
			"<b>1. Syarat Pembayaran Balik</b>\n" +
			"1.1. Oleh kerana kami menyediakan barangan digital (kunci akses) dan perkhidmatan yang dimakan pada masa penggunaan, pembayaran balik hanya boleh dalam kes luar biasa.\n" +
			"1.2. Anda berhak meminta pembayaran balik penuh dalam masa 24 jam selepas pembelian jika anda tidak dapat menyambung ke perkhidmatan kerana sebab teknikal yang bergantung pada Pembekal (pelayan tidak berfungsi, kesilapan penjanaan kunci).\n\n" +
			"<b>2. Penolakan Pembayaran Balik</b>\n" +
			"Pembayaran balik TIDAK dilakukan dalam kes berikut:\n" +
			"• Lebih daripada 24 jam telah berlalu sejak pembelian.\n" +
			"• Kelajuan sambungan rendah yang disebabkan oleh had dari pihak pembekal pengguna atau kualiti liputan rangkaian mudah alih yang buruk.\n" +
			"• Ketidakupayaan pengguna untuk menyiapkan perisian pihak ketiga (klien Happ, Streisand dll.), dengan syarat kunci itu sendiri sah.\n" +
			"• Sekatan akses kerana pelanggaran peraturan (spam, torrent, serangan penggodaman).\n" +
			"• “Saya ubah fikiran” atau “Saya tak perlu lagi”.\n\n" +
			"<b>3. Prosedur Pembayaran Balik</b>\n" +
			"3.1. Untuk memohon pembayaran balik, tulis kepada sokongan.\n" +
			"3.2. Dalam permohonan nyatakan: ID akaun, resit pembayaran dan perihalan masalah teknikal (lebih baik dengan tangkapan skrin kesilapan).\n" +
			"3.3. Pembayaran balik dilakukan dengan kaedah pembayaran yang sama dalam masa 3–7 hari bekerja.",
	},
	language.Tagalog: {
		Policy: "<b>PAT KARANIWAN NG REFUND:</b>\n\n" +
			"<b>1. Kondisyon ng Refund</b>\n" +
			"1.1. Dahil nagbibigay kami ng digital na produkto (access keys) at serbisyo na nauubos sa oras ng paggamit, ang refund ay posible lamang sa mga eksepsiyonal na kaso.\n" +
			"1.2. May karapatan kang humiling ng buong refund sa loob ng 24 oras mula sa pagbili kung hindi ka nakakonekta sa serbisyo dahil sa teknikal na dahilan na nakadepende sa Provider (hindi gumaganang server, error sa pagbuo ng key).\n\n" +
			"<b>2. Pagtanggi sa Refund</b>\n" +
			"Hindi ginagawa ang refund sa mga sumusunod na kaso:\n" +
			"• Lumipas na ang higit sa 24 oras mula sa pagbili.\n" +
			"• Mababang bilis ng koneksyon na dulot ng limitasyon sa panig ng provider ng user o mahinang coverage ng mobile network.\n" +
			"• Kawalan ng kakayahang mag-set up ng third-party software (Happ, Streisand clients, atbp.), kung ang key mismo ay valid.\n" +
			"• Pag-block ng access dahil sa paglabag sa rules (spam, torrents, hacking attacks).\n" +
			"• “Nagbago ang isip ko” o “Hindi ko na kailangan”.\n\n" +
			"<b>3. Proseso ng Refund</b>\n" +
			"3.1. Para sa paghingi ng refund, sumulat sa support.\n" +
			"3.2. Sa request, tukuyin: Account ID, resibo ng bayad at paglalarawan ng teknikal na problema (mas mabuti na may screenshot ng error).\n" +
			"3.3. Ang refund ay ginagawa sa parehong paraan ng bayad sa loob ng 3–7 araw ng trabaho.",
	},
	language.Hindi: {
		Policy: "<b>रिफंड पॉलिसी:</b>\n\n" +
			"<b>1. रिफंड की शर्तें</b>\n" +
			"1.1. चूंकि हम डिजिटल गुड्स (एक्सेस कीज) और सेवाएं प्रदान करते हैं जो उपयोग के समय ही खपत हो जाती हैं, रिफंड केवल असाधारण मामलों में संभव है.\n" +
			"1.2. आपको खरीद के 24 घंटों के भीतर पूर्ण रिफंड मांगने का अधिकार है अगर आप प्रदाता की तकनीकी वजहों से (नॉन-वर्किंग सर्वर, की जेनरेशन एरर) सेवा से कनेक्ट नहीं हो पाए.\n\n" +
			"<b>2. रिफंड अस्वीकृति</b>\n" +
			"रिफंड नहीं किया जाता निम्न मामलों में:\n" +
			"• खरीद से 24 घंटे से ज्यादा समय बीत चुका है.\n" +
			"• यूजर के प्रोवाइडर की तरफ से लिमिटेशन या मोबाइल नेटवर्क कवरेज की खराब क्वालिटी से कम कनेक्शन स्पीड.\n" +
			"• यूजर की थर्ड-पार्टी सॉफ्टवेयर (Happ, Streisand क्लाइंट्स आदि) सेट करने में असमर्थता, बशर्ते की खुद वैलिड हो.\n" +
			"• रूल्स वायलेशन (स्पैम, टॉरेंट्स, हैकिंग अटैक्स) के लिए एक्सेस ब्लॉक.\n" +
			"• “मैंने अपना मन बदल लिया” या “मुझे अब जरूरत नहीं है”.\n\n" +
			"<b>3. रिफंड प्रक्रिया</b>\n" +
			"3.1. रिफंड के लिए सपोर्ट को लिखें.\n" +
			"3.2. रिक्वेस्ट में बताएं: अकाउंट ID, पेमेंट रसीद और तकनीकी समस्या का वर्णन (बेहतर होगा एरर का स्क्रीनशॉट).\n" +
			"3.3. रिफंड उसी पेमेंट मेथड से 3–7 वर्किंग डेज में किया जाता है.",
	},
	language.URاردو: {
		Policy: "<b>ریفنڈ پالیسی:</b>\n\n" +
			"<b>1. ریفنڈ کی شرائط</b>\n" +
			"1.1. چونکہ ہم ڈیجیٹل سامان (رسائی کیز) اور خدمات فراہم کرتے ہیں جو استعمال کے وقت ہی استعمال ہو جاتی ہیں، ریفنڈ صرف غیر معمولی معاملات میں ممکن ہے.\n" +
			"1.2. آپ کو خریداری کے 24 گھنٹوں کے اندر مکمل ریفنڈ کی درخواست کرنے کا حق ہے اگر آپ فراہم کنندہ کی تکنیکی وجوہات سے (غیر فعال سرور، کی جنریشن ایرر) سروس سے جڑ نہیں سکے.\n\n" +
			"<b>2. ریفنڈ کی مستردی</b>\n" +
			"ریفنڈ نہیں کیا جاتا درج ذیل معاملات میں:\n" +
			"• خریداری سے 24 گھنٹے سے زیادہ وقت گزر چکا ہے.\n" +
			"• کم کنکشن سپیڈ جو صارف کے پرووائیڈر کی طرف سے لمیٹیشن یا موبائل نیٹ ورک کوریج کی خراب کوالٹی کی وجہ سے ہے.\n" +
			"• صارف کی تھرڈ پارٹی سافٹ ویئر (Happ, Streisand کلائنٹس وغیرہ) سیٹ کرنے میں ناکامی، بشرطیکہ کی خود درست ہو.\n" +
			"• رولز کی خلاف ورزی (سپیم, ٹورنٹس, ہیکنگ اٹیکس) کی وجہ سے رسائی بلاک.\n" +
			"• “میں نے اپنا خیال بدل لیا” یا “مجھے اب ضرورت نہیں ہے”.\n\n" +
			"<b>3. ریفنڈ کا طریقہ کار</b>\n" +
			"3.1. ریفنڈ کے لیے سپورٹ کو لکھیں.\n" +
			"3.2. درخواست میں بتائیں: اکاؤنٹ ID, پیمنٹ رسید اور تکنیکی مسئلے کی تفصیل (بہتر ہو گا ایرر کا اسکرین شاٹ).\n" +
			"3.3. ریفنڈ اسی پیمنٹ میتھڈ سے 3–7 ورکنگ ڈیز میں کیا جاتا ہے.",
	},
	language.Bengali: {
		Policy: "<b>রিফান্ড পলিসি:</b>\n\n" +
			"<b>1. রিফান্ডের শর্তাবলী</b>\n" +
			"1.1. যেহেতু আমরা ডিজিটাল গুডস (অ্যাক্সেস কী) এবং সার্ভিস প্রদান করি যা ব্যবহারের সময়ই খরচ হয়ে যায়, রিফান্ড শুধুমাত্র অসাধারণ কেসে সম্ভব.\n" +
			"1.2. ক্রয়ের ২৪ ঘন্টার মধ্যে আপনার সম্পূর্ণ রিফান্ড দাবি করার অধিকার আছে যদি আপনি প্রোভাইডারের টেকনিক্যাল কারণে (নন-ওয়ার্কিং সার্ভার, কী জেনারেশন এরর) সার্ভিসে কানেক্ট করতে না পারেন.\n\n" +
			"<b>2. রিফান্ড প্রত্যাখ্যান</b>\n" +
			"রিফান্ড করা হয় না নিম্নলিখিত কেসে:\n" +
			"• ক্রয় থেকে ২৪ ঘন্টা অতিক্রম করেছে.\n" +
			"• কম কানেকশন স্পিড যা ইউজারের প্রোভাইডারের লিমিটেশন বা মোবাইল নেটওয়ার্ক কভারেজের খারাপ কোয়ালিটির কারণে.\n" +
			"• ইউজারের থার্ড-পার্টি সফটওয়্যার (Happ, Streisand ক্লায়েন্টস ইত্যাদি) সেট করতে অক্ষমতা, যদি কী নিজেই ভ্যালিড হয়.\n" +
			"• রুলস ভায়োলেশন (স্প্যাম, টরেন্টস, হ্যাকিং অ্যাটাকস) এর জন্য অ্যাক্সেস ব্লক.\n" +
			"• “আমি মন পরিবর্তন করেছি” বা “আমার আর দরকার নেই”.\n\n" +
			"<b>3. রিফান্ড প্রসেস</b>\n" +
			"3.1. রিফান্ডের জন্য সাপোর্টে লিখুন.\n" +
			"3.2. রিকোয়েস্টে উল্লেখ করুন: অ্যাকাউন্ট আইডি, পেমেন্ট রিসিপ্ট এবং টেকনিক্যাল প্রবলেমের বর্ণনা (ভালো হবে এররের স্ক্রিনশট).\n" +
			"3.3. রিফান্ড একই পেমেন্ট মেথডে ৩–৭ ওয়ার্কিং ডেয়ে করা হয়.",
	},
	language.Tamiḻ: {
		Policy: "<b>பணத்திருப்பு கொள்கை:</b>\n\n" +
			"<b>1. பணத்திருப்பு நிபந்தனைகள்</b>\n" +
			"1.1. நாங்கள் டிஜிட்டல் பொருட்கள் (அணுகல் சாவிகள்) மற்றும் பயன்பாட்டில் உடனடியாக உட்கொள்ளப்படும் சேவைகளை வழங்குவதால், பணத்திருப்பு விதிவிலக்கான வழக்குகளில் மட்டுமே சாத்தியம்.\n" +
			"1.2. வாங்கிய 24 மணி நேரத்திற்குள் முழு பணத்திருப்பு கோருவதற்கு உரிமை உண்டு, வழங்குநரின் தொழில்நுட்ப காரணங்களால் (வேலை செய்யாத சர்வர், சாவி உருவாக்க பிழை) சேவையுடன் இணைக்க முடியவில்லை என்றால்.\n\n" +
			"<b>2. பணத்திருப்பு மறுப்பு</b>\n" +
			"பணத்திருப்பு செய்யப்படாது பின்வரும் வழக்குகளில்:\n" +
			"• வாங்கியதிலிருந்து 24 மணி நேரத்திற்கு மேல் கடந்துவிட்டது.\n" +
			"• பயனரின் வழங்குநரின் வரம்புகள் அல்லது மொபைல் நெட்வொர்க் கவரேஜ் தரம் குறைவானதால் குறைந்த இணைப்பு வேகம்.\n" +
			"• பயனர் மூன்றாவது தரப்பு மென்பொருள் (Happ, Streisand க்ளையன்ட்கள் போன்றவை) அமைக்க இயலாதது, சாவி சரியானதாக இருந்தால்.\n" +
			"• விதிகள் மீறலால் அணுகல் தடை (ஸ்பேம், டோரண்ட்கள், ஹேக்கிங் அட்டாக்கள்).\n" +
			"• “நான் மனம் மாற்றினேன்” அல்லது “இனி எனக்கு தேவையில்லை”.\n\n" +
			"<b>3. பணத்திருப்பு நடைமுறை</b>\n" +
			"3.1. பணத்திருப்பு கோருவதற்கு சப்போர்ட்டுக்கு எழுதுங்கள்.\n" +
			"3.2. கோரிக்கையில் குறிப்பிடுங்கள்: அக்கவுண்ட் ஐடி, பேமெண்ட் ரசீது மற்றும் தொழில்நுட்ப பிரச்சனையின் விளக்கம் (பிழை ஸ்கிரீன்ஷாட் விரும்பத்தக்கது).\n" +
			"3.3. பணத்திருப்பு அதே பேமெண்ட் முறையில் 3–7 வேலை நாட்களுக்குள் செய்யப்படும்.",
	},
	language.Telugu: {
		Policy: "<b>రీఫండ్ పాలసీ:</b>\n\n" +
			"<b>1. రీఫండ్ షరతులు</b>\n" +
			"1.1. మేము డిజిటల్ గూడ్స్ (యాక్సెస్ కీలు) మరియు ఉపయోగం సమయంలో ఖర్చయ్యే సర్వీసులు అందిస్తాము కాబట్టి, రీఫండ్ అసాధారణ కేసులలో మాత్రమే సాధ్యం.\n" +
			"1.2. కొనుగోలు తర్వాత 24 గంటలలోపు పూర్తి రీఫండ్ అభ్యర్థించే హక్కు మీకు ఉంది, ప్రొవైడర్ టెక్నికల్ కారణాల వల్ల (వర్కింగ్ కాని సర్వర్, కీ జెనరేషన్ ఎర్రర్) సర్వీస్ కి కనెక్ట్ కాలేదు అంటే.\n\n" +
			"<b>2. రీఫండ్ తిరస్కరణ</b>\n" +
			"రీఫండ్ చేయబడదు కింది కేసులలో:\n" +
			"• కొనుగోలు నుండి 24 గంటలు మించి పోయింది.\n" +
			"• యూజర్ ప్రొవైడర్ సైడ్ లిమిటేషన్లు లేదా మొబైల్ నెట్వర్క్ కవరేజ్ క్వాలిటీ తక్కువ కారణంగా తక్కువ కనెక్షన్ స్పీడ్.\n" +
			"• యూజర్ థర్డ్-పార్టీ సాఫ్ట్వేర్ (Happ, Streisand క్లయింట్లు మొదలైనవి) సెటప్ చేయలేకపోవడం, కీ స్వయంగా వాలిడ్ అయితే.\n" +
			"• రూల్స్ వయోలేషన్ (స్పామ్, టారెంట్స్, హ్యాకింగ్ అటాక్స్) కోసం యాక్సెస్ బ్లాక్.\n" +
			"• “నేను మనసు మార్చాను” లేదా “నాకు ఇక అవసరం లేదు”.\n\n" +
			"<b>3. రీఫండ్ ప్రాసెస్</b>\n" +
			"3.1. రీఫండ్ కోసం సపోర్ట్ కి రాయండి.\n" +
			"3.2. రిక్వెస్ట్ లో స్పెసిఫై: అకౌంట్ ఐడీ, పేమెంట్ రసీదు మరియు టెక్నికల్ ప్రాబ్లెమ్ డిస్క్రిప్షన్ (ఎర్రర్ స్క్రీన్ షాట్ బెటర్).\n" +
			"3.3. రీఫండ్ అదే పేమెంట్ మెథడ్ తో 3–7 వర్కింగ్ డేస్ లో చేయబడుతుంది.",
	},
	language.Marathi: {
		Policy: "<b>रिफंड पॉलिसी:</b>\n\n" +
			"<b>1. रिफंडच्या अटी</b>\n" +
			"1.1. आम्ही डिजिटल गुड्स (ऍक्सेस कीज) आणि वापराच्या वेळी खर्च होणाऱ्या सेवा देतो म्हणून, रिफंड फक्त अपवादात्मक प्रकरणांमध्ये शक्य आहे.\n" +
			"1.2. खरेदीच्या २४ तासांच्या आत तुम्हाला पूर्ण रिफंड मागण्याचा अधिकार आहे जर तुम्ही प्रदात्याच्या तांत्रिक कारणांमुळे (नॉन-वर्किंग सर्व्हर, की जेनरेशन एरर) सेवेशी कनेक्ट होऊ शकला नाहीत.\n\n" +
			"<b>2. रिफंड नाकारणे</b>\n" +
			"रिफंड केले जात नाही खालील प्रकरणांमध्ये:\n" +
			"• खरेदीपासून २४ तासांपेक्षा जास्त वेळ गेला आहे.\n" +
			"• युजरच्या प्रोव्हायडरच्या बाजूच्या मर्यादा किंवा मोबाइल नेटवर्क कव्हरेजच्या खराब क्वालिटीमुळे कमी कनेक्शन स्पीड.\n" +
			"• युजरची थर्ड-पार्टी सॉफ्टवेअर (Happ, Streisand क्लायंट्स इत्यादी) सेट करण्यात असमर्थता, की स्वतः वैध असल्यास.\n" +
			"• रूल्स व्हायोलेशन (स्पॅम, टॉरेंट्स, हॅकिंग अटॅक्स) साठी ऍक्सेस ब्लॉक.\n" +
			"• “मी माझा विचार बदलला” किंवा “मला आता गरज नाही”.\n\n" +
			"<b>3. रिफंड प्रक्रिया</b>\n" +
			"3.1. रिफंडसाठी सपोर्टला लिहा.\n" +
			"3.2. रिक्वेस्टमध्ये सांगाः अकाउंट आयडी, पेमेंट रिसीट आणि तांत्रिक समस्येचे वर्णन (एररचा स्क्रीनशॉट चांगला).\n" +
			"3.3. रिफंड त्याच पेमेंट मेथडने ३–७ वर्किंग डेजमध्ये केले जाते.",
	},
}

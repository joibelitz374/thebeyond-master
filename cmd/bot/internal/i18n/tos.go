package i18n

import "github.com/quickpowered/thebeyond-master/configs/language"

type TosLocale struct {
	Terms string
}

var TosMessages = map[language.Language]TosLocale{
	language.English: {
		Terms: "<b>PUBLIC OFFER:</b>\n\n" +
			"<b>1. General Provisions</b>\n" +
			"1.1. This document serves as the official proposal (public offer) from THE BEYOND service (hereinafter referred to as the \"Provider\") and outlines all essential terms for providing services related to organizing remote access to network nodes (proxy/VPN).\n" +
			"1.2. Acceptance of this Offer is deemed to occur upon starting to use the Telegram bot, pressing the \"Start\" button, or paying for a tariff plan.\n\n" +
			"<b>2. Subject of the Offer</b>\n" +
			"2.1. The Provider grants the Customer a non-exclusive right (license) to use access configurations (VLESS/Hysteria 2/MTProto keys) for the Provider's servers.\n" +
			"2.2. The service is considered rendered at the moment the access key or configuration file is sent to the Customer automatically via the Telegram bot.\n\n" +
			"<b>3. Rights and Obligations of the Parties</b>\n" +
			"3.1. The Provider commits to maintaining server functionality but does not guarantee 100% availability (Uptime) in cases of blocks by the Customer's internet providers or government authorities (force majeure).\n" +
			"3.2. The Customer agrees to use the services solely for lawful purposes.\n" +
			"3.3. The Customer is strictly prohibited from:\n" +
			"• Using the service to send spam (email, social networks).\n" +
			"• Port scanning, conducting DDoS attacks, brute-force attempts, or hacking resources.\n" +
			"• Distributing malicious software, narcotic substances, child pornography, or materials promoting terrorism.\n" +
			"• Sharing access keys with third parties (unless permitted by the tariff).\n\n" +
			"<b>4. Liability and Blocking</b>\n" +
			"4.1. In the event of a violation of clause 3.3, the Provider reserves the right to immediately block the Customer's access without refunding any payments.\n" +
			"4.2. The Provider bears no responsibility for any direct or indirect damages incurred by the Customer due to the use or inability to use the service.\n" +
			"4.3. The service is provided on an \"As Is\" basis. The Provider does not guarantee compatibility with the Customer's specific device or provider.",
	},
	language.Deutsch: {
		Terms: "<b>ÖFFENTLICHES ANGEBOT:</b>\n\n" +
			"<b>1. Allgemeine Bestimmungen</b>\n" +
			"1.1. Dieses Dokument stellt das offizielle Angebot (öffentliches Angebot) des Dienstes THE BEYOND (im Folgenden als \"Anbieter\" bezeichnet) dar und enthält alle wesentlichen Bedingungen für die Bereitstellung von Diensten zur Organisation des Fernzugriffs auf Netzwerkknoten (Proxy/VPN).\n" +
			"1.2. Die Annahme dieses Angebots gilt als erfolgt, sobald der Telegram-Bot genutzt wird, die Schaltfläche \"Start\" gedrückt wird oder ein Tarifplan bezahlt wird.\n\n" +
			"<b>2. Gegenstand des Angebots</b>\n" +
			"2.1. Der Anbieter gewährt dem Kunden ein nicht-exklusives Recht (Lizenz) zur Nutzung von Zugriffskonfigurationen (VLESS/Hysteria 2/MTProto-Schlüsseln) auf die Server des Anbieters.\n" +
			"2.2. Der Dienst gilt als erbracht, sobald der Zugriffsschlüssel oder die Konfigurationsdatei automatisch über den Telegram-Bot an den Kunden gesendet wird.\n\n" +
			"<b>3. Rechte und Pflichten der Parteien</b>\n" +
			"3.1. Der Anbieter verpflichtet sich, die Funktionsfähigkeit der Server aufrechtzuerhalten, garantiert jedoch keine 100%ige Verfügbarkeit (Uptime) im Falle von Blockaden durch den Internetprovider des Kunden oder staatliche Behörden (höhere Gewalt).\n" +
			"3.2. Der Kunde verpflichtet sich, die Dienste ausschließlich zu rechtmäßigen Zwecken zu nutzen.\n" +
			"3.3. Dem Kunden ist strengstens untersagt:\n" +
			"• Den Dienst für den Versand von Spam (E-Mail, soziale Netzwerke) zu verwenden.\n" +
			"• Ports zu scannen, DDoS-Angriffe durchzuführen, Brute-Force-Attacken oder das Hacken von Ressourcen.\n" +
			"• Schadsoftware, narkotische Substanzen, Kinderpornografie oder Materialien zu verbreiten, die Terrorismus propagieren.\n" +
			"• Zugriffsschlüssel an Dritte weiterzugeben (sofern nicht im Tarif vorgesehen).\n\n" +
			"<b>4. Haftung und Sperrung</b>\n" +
			"4.1. Bei Verstoß gegen Abschnitt 3.3 hat der Anbieter das Recht, den Zugriff des Kunden sofort zu sperren, ohne Rückerstattung von Geldern.\n" +
			"4.2. Der Anbieter haftet nicht für Schäden, direkt oder indirekt, die dem Kunden durch die Nutzung oder Unmöglichkeit der Nutzung des Dienstes entstehen.\n" +
			"4.3. Der Dienst wird auf \"Wie besehen\"-Basis (As Is) bereitgestellt. Der Anbieter garantiert nicht, dass der Dienst mit einem spezifischen Gerät oder Provider des Kunden kompatibel ist.",
	},
	language.Nederlands: {
		Terms: "<b>PUBLIEK AANBOD:</b>\n\n" +
			"<b>1. Algemene Bepalingen</b>\n" +
			"1.1. Dit document is het officiële voorstel (publiek aanbod) van de dienst THE BEYOND (hierna \"Aanbieder\" genoemd) en bevat alle essentiële voorwaarden voor het verlenen van diensten voor het organiseren van externe toegang tot netwerkknooppunten (proxy/VPN).\n" +
			"1.2. Acceptatie van dit Aanbod wordt geacht plaats te vinden bij het beginnen met het gebruik van de Telegram-bot, het drukken op de knop \"Start\" of het betalen van een tariefplan.\n\n" +
			"<b>2. Onderwerp van het Aanbod</b>\n" +
			"2.1. De Aanbieder verleent de Klant een niet-exclusief recht (licentie) om toegangconfiguraties (VLESS/Hysteria 2/MTProto-sleutels) te gebruiken voor de servers van de Aanbieder.\n" +
			"2.2. De dienst wordt geacht verleend te zijn op het moment dat de toegangssleutel of configuratiebestand automatisch via de Telegram-bot aan de Klant wordt verzonden.\n\n" +
			"<b>3. Rechten en Verplichtingen van de Partijen</b>\n" +
			"3.1. De Aanbieder verplicht zich om de werking van de servers te onderhouden, maar garandeert geen 100% beschikbaarheid (Uptime) in geval van blokkades door de internetproviders van de Klant of overheidsinstanties (overmacht).\n" +
			"3.2. De Klant verplicht zich om de diensten uitsluitend voor legale doeleinden te gebruiken.\n" +
			"3.3. Het is de Klant strikt verboden:\n" +
			"• De dienst te gebruiken voor het verzenden van spam (e-mail, sociale netwerken).\n" +
			"• Poorten te scannen, DDoS-aanvallen uit te voeren, brute-force en het hacken van bronnen.\n" +
			"• Schadelijke software, narcotica, kinderpornografie of materialen die terrorisme propageren te verspreiden.\n" +
			"• Toegangssleutels door te geven aan derden (tenzij voorzien in het tarief).\n\n" +
			"<b>4. Aansprakelijkheid en Blokkering</b>\n" +
			"4.1. Bij overtreding van clausule 3.3 heeft de Aanbieder het recht om de toegang van de Klant onmiddellijk te blokkeren zonder terugbetaling van gelden.\n" +
			"4.2. De Aanbieder is niet aansprakelijk voor schade, direct of indirect, die de Klant lijdt in verband met het gebruik of de onmogelijkheid om de dienst te gebruiken.\n" +
			"4.3. De dienst wordt geleverd op \"As Is\"-basis. De Aanbieder garandeert niet dat de dienst compatibel is met een specifiek apparaat of provider van de Klant.",
	},
	language.Svenska: {
		Terms: "<b>PUBLIKT ERBJUDANDE:</b>\n\n" +
			"<b>1. Allmänna Bestämmelser</b>\n" +
			"1.1. Detta dokument är det officiella förslaget (offentligt erbjudande) från tjänsten THE BEYOND (härefter \"Leverantör\") och innehåller alla väsentliga villkor för tillhandahållande av tjänster för att organisera fjärråtkomst till nätverksnoder (proxy/VPN).\n" +
			"1.2. Acceptans av detta Erbjudande anses ske vid starten av användningen av Telegram-boten, tryck på knappen \"Start\" eller betalning för tariffplanen.\n\n" +
			"<b>2. Ämne för Erbjudandet</b>\n" +
			"2.1. Leverantören ger Kunden en icke-exklusiv rätt (licens) att använda åtkomstkonfigurationer (VLESS/Hysteria 2/MTProto-nycklar) till Leverantörens servrar.\n" +
			"2.2. Tjänsten anses levererad i det ögonblick åtkomstnyckeln eller konfigurationsfilen skickas till Kunden automatiskt via Telegram-boten.\n\n" +
			"<b>3. Parternas Rättigheter och Skyldigheter</b>\n" +
			"3.1. Leverantören åtar sig att upprätthålla serverns funktionalitet, men garanterar inte 100% tillgänglighet (Uptime) vid blockeringar från Kundens internetleverantörer eller statliga organ (force majeure).\n" +
			"3.2. Kunden åtar sig att använda tjänsterna enbart för lagliga ändamål.\n" +
			"3.3. Kunden är strikt förbjuden att:\n" +
			"• Använda tjänsten för att skicka spam (e-post, sociala nätverk).\n" +
			"• Skanna portar, utföra DDoS-attacker, brute-force och hacka resurser.\n" +
			"• Sprida skadlig programvara, narkotika, barnpornografi eller material som främjar terrorism.\n" +
			"• Överföra åtkomstnycklar till tredje part (om inte tariffen tillåter det).\n\n" +
			"<b>4. Ansvar och Blockering</b>\n" +
			"4.1. Vid brott mot punkt 3.3 har Leverantören rätt att omedelbart blockera Kundens åtkomst utan återbetalning av pengar.\n" +
			"4.2. Leverantören ansvarar inte för skador, direkta eller indirekta, som Kunden ådrar sig i samband med användning eller oförmåga att använda tjänsten.\n" +
			"4.3. Tjänsten tillhandahålls på \"Som den är\"-basis (As Is). Leverantören garanterar inte att tjänsten är kompatibel med Kundens specifika enhet eller leverantör.",
	},
	language.Norsk: {
		Terms: "<b>OFFENTLIG TILBUD:</b>\n\n" +
			"<b>1. Generelle Bestemmelser</b>\n" +
			"1.1. Dette dokumentet er det offisielle forslaget (offentlig tilbud) fra tjenesten THE BEYOND (heretter \"Leverandør\") og inneholder alle vesentlige vilkår for å tilby tjenester for organisering av fjernaksess til nettverksnoder (proxy/VPN).\n" +
			"1.2. Aksept av dette Tilbudet anses å skje ved start av bruken av Telegram-boten, trykk på \"Start\"-knappen eller betaling for tariffplanen.\n\n" +
			"<b>2. Emne for Tilbudet</b>\n" +
			"2.1. Leverandøren gir Kunden en ikke-eksklusiv rett (lisens) til å bruke tilgangskonfigurasjoner (VLESS/Hysteria 2/MTProto-nøkler) til Leverandørens servere.\n" +
			"2.2. Tjenesten anses levert i det øyeblikket tilgangsnøkkelen eller konfigurasjonsfilen sendes til Kunden automatisk via Telegram-boten.\n\n" +
			"<b>3. Partenes Rettigheter og Plikter</b>\n" +
			"3.1. Leverandøren forplikter seg til å opprettholde servernes funksjonalitet, men garanterer ikke 100% tilgjengelighet (Uptime) i tilfelle blokkeringer fra Kundens internettleverandører eller statlige organer (force majeure).\n" +
			"3.2. Kunden forplikter seg til å bruke tjenestene utelukkende til lovlige formål.\n" +
			"3.3. Kunden er strengt forbudt å:\n" +
			"• Bruke tjenesten til å sende spam (e-post, sosiale nettverk).\n" +
			"• Skanne porter, utføre DDoS-angrep, brute-force og hacking av ressurser.\n" +
			"• Distribuere skadelig programvare, narkotika, barnepornografi eller materialer som promoterer terrorisme.\n" +
			"• Overføre tilgangsnøkler til tredjeparter (med mindre tariffen tillater det).\n\n" +
			"<b>4. Ansvar og Blokkering</b>\n" +
			"4.1. Ved brudd på punkt 3.3 har Leverandøren rett til umiddelbart å blokkere Kundens tilgang uten refusjon av penger.\n" +
			"4.2. Leverandøren er ikke ansvarlig for skader, direkte eller indirekte, som Kunden pådrar seg i forbindelse med bruk eller umulighet for bruk av tjenesten.\n" +
			"4.3. Tjenesten leveres på \"Som den er\"-basis (As Is). Leverandøren garanterer ikke at tjenesten er kompatibel med Kundens spesifikke enhet eller leverandør.",
	},
	language.Dansk: {
		Terms: "<b>OFFENTLIG TILBUD:</b>\n\n" +
			"<b>1. Generelle Bestemmelser</b>\n" +
			"1.1. Dette dokument er det officielle forslag (offentligt tilbud) fra tjenesten THE BEYOND (herefter \"Udbyder\") og indeholder alle væsentlige betingelser for at yde tjenester til organisering af fjernadgang til netværksknuder (proxy/VPN).\n" +
			"1.2. Accept af dette Tilbud anses for at ske ved start af brug af Telegram-botten, tryk på knappen \"Start\" eller betaling for tariffplanen.\n\n" +
			"<b>2. Emne for Tilbuddet</b>\n" +
			"2.1. Udbyderen giver Kunden en ikke-eksklusiv rettighed (licens) til at bruge adgangskonfigurationer (VLESS/Hysteria 2/MTProto-nøgler) til Udbyderens servere.\n" +
			"2.2. Tjenesten anses for ydet i det øjeblik adgangsnøglen eller konfigurationsfilen sendes til Kunden automatisk via Telegram-botten.\n\n" +
			"<b>3. Parternes Rettigheder og Forpligtelser</b>\n" +
			"3.1. Udbyderen forpligter sig til at opretholde servernes funktionalitet, men garanterer ikke 100% tilgængelighed (Uptime) i tilfælde af blokeringer fra Kundens internetudbydere eller statslige organer (force majeure).\n" +
			"3.2. Kunden forpligter sig til at bruge tjenesterne udelukkende til lovlige formål.\n" +
			"3.3. Kunden er strengt forbudt at:\n" +
			"• Bruge tjenesten til at sende spam (e-mail, sociale netværk).\n" +
			"• Scan porte, udføre DDoS-angreb, brute-force og hacke ressourcer.\n" +
			"• Distribuere skadelig software, narkotika, børnepornografi eller materialer, der fremmer terrorisme.\n" +
			"• Overføre adgangsnøgler til tredjeparter (medmindre tariffen tillader det).\n\n" +
			"<b>4. Ansvar og Blokering</b>\n" +
			"4.1. Ved overtrædelse af punkt 3.3 har Udbyderen ret til øjeblikkeligt at blokere Kundens adgang uden tilbagebetaling af penge.\n" +
			"4.2. Udbyderen er ikke ansvarlig for skader, direkte eller indirekte, som Kunden pådrager sig i forbindelse med brug eller umulighed for brug af tjenesten.\n" +
			"4.3. Tjenesten ydes på \"Som den er\"-basis (As Is). Udbyderen garanterer ikke, at tjenesten er kompatibel med Kundens specifikke enhed eller udbyder.",
	},
	language.Español: {
		Terms: "<b>OFERTA PÚBLICA:</b>\n\n" +
			"<b>1. Disposiciones Generales</b>\n" +
			"1.1. Este documento es la propuesta oficial (oferta pública) del servicio THE BEYOND (en adelante, \"Proveedor\") y contiene todas las condiciones esenciales para la prestación de servicios de organización de acceso remoto a nodos de red (proxy/VPN).\n" +
			"1.2. La aceptación de esta Oferta se considera efectiva al comenzar a usar el bot de Telegram, presionar el botón \"Start\" o pagar el plan tarifario.\n\n" +
			"<b>2. Objeto de la Oferta</b>\n" +
			"2.1. El Proveedor otorga al Cliente un derecho no exclusivo (licencia) para usar configuraciones de acceso (claves VLESS/Hysteria 2/MTProto) a los servidores del Proveedor.\n" +
			"2.2. El servicio se considera prestado en el momento en que se envía la clave de acceso o el archivo de configuración al Cliente de manera automática a través del bot de Telegram.\n\n" +
			"<b>3. Derechos y Obligaciones de las Partes</b>\n" +
			"3.1. El Proveedor se compromete a mantener la operatividad de los servidores, sin embargo, no garantiza el 100% de disponibilidad (Uptime) en caso de bloqueos por parte de los proveedores de internet del Cliente o autoridades estatales (fuerza mayor).\n" +
			"3.2. El Cliente se compromete a usar los servicios exclusivamente para fines legales.\n" +
			"3.3. Al Cliente se le prohíbe estrictamente:\n" +
			"• Usar el servicio para enviar spam (correo electrónico, redes sociales).\n" +
			"• Escanear puertos, realizar ataques DDoS, brute force y hackeo de recursos.\n" +
			"• Distribuir software malicioso, sustancias narcóticas, pornografía infantil o materiales que promuevan el terrorismo.\n" +
			"• Transferir claves de acceso a terceros (a menos que esté previsto en la tarifa).\n\n" +
			"<b>4. Responsabilidad y Bloqueo</b>\n" +
			"4.1. En caso de violación del punto 3.3, el Proveedor tiene derecho a bloquear inmediatamente el acceso del Cliente sin reembolso de fondos.\n" +
			"4.2. El Proveedor no es responsable por daños, directos o indirectos, sufridos por el Cliente en relación con el uso o la imposibilidad de usar el servicio.\n" +
			"4.3. El servicio se proporciona en condiciones \"Tal cual\" (As Is). El Proveedor no garantiza que el servicio sea compatible con el dispositivo específico o proveedor del Cliente.",
	},
	language.Français: {
		Terms: "<b>OFFRE PUBLIQUE :</b>\n\n" +
			"<b>1. Dispositions Générales</b>\n" +
			"1.1. Ce document constitue la proposition officielle (offre publique) du service THE BEYOND (ci-après « Fournisseur ») et contient toutes les conditions essentielles pour la fourniture de services d'organisation d'accès distant aux nœuds de réseau (proxy/VPN).\n" +
			"1.2. L'acceptation de cette Offre est considérée comme effective dès le début de l'utilisation du bot Telegram, en appuyant sur le bouton « Start » ou en payant le plan tarifaire.\n\n" +
			"<b>2. Objet de l'Offre</b>\n" +
			"2.1. Le Fournisseur accorde au Client un droit non exclusif (licence) d'utiliser les configurations d'accès (clés VLESS/Hysteria 2/MTProto) aux serveurs du Fournisseur.\n" +
			"2.2. Le service est considéré comme rendu au moment de l'envoi de la clé d'accès ou du fichier de configuration au Client de manière automatique via le bot Telegram.\n\n" +
			"<b>3. Droits et Obligations des Parties</b>\n" +
			"3.1. Le Fournisseur s'engage à maintenir le fonctionnement des serveurs, cependant, ne garantit pas une disponibilité à 100 % (Uptime) en cas de blocages par les fournisseurs d'accès internet du Client ou les organes d'État (force majeure).\n" +
			"3.2. Le Client s'engage à utiliser les services exclusivement à des fins légales.\n" +
			"3.3. Il est strictement interdit au Client :\n" +
			"• D'utiliser le service pour envoyer du spam (e-mail, réseaux sociaux).\n" +
			"• De scanner des ports, de mener des attaques DDoS, du brute force et du piratage de ressources.\n" +
			"• De distribuer des logiciels malveillants, des substances narcotiques, de la pornographie infantile ou des matériaux promouvant le terrorisme.\n" +
			"• De transmettre des clés d'accès à des tiers (sauf si prévu par le tarif).\n\n" +
			"<b>4. Responsabilité et Blocage</b>\n" +
			"4.1. En cas de violation du point 3.3, le Fournisseur a le droit de bloquer immédiatement l'accès du Client sans remboursement des fonds.\n" +
			"4.2. Le Fournisseur n'est pas responsable des dommages, directs ou indirects, subis par le Client en lien avec l'utilisation ou l'impossibilité d'utiliser le service.\n" +
			"4.3. Le service est fourni sur une base « Tel quel » (As Is). Le Fournisseur ne garantit pas que le service sera compatible avec l'appareil spécifique ou le fournisseur du Client.",
	},
	language.Português: {
		Terms: "<b>OFERTA PÚBLICA:</b>\n\n" +
			"<b>1. Disposições Gerais</b>\n" +
			"1.1. Este documento é a proposta oficial (oferta pública) do serviço THE BEYOND (doravante \"Fornecedor\") e contém todas as condições essenciais para a prestação de serviços de organização de acesso remoto a nós de rede (proxy/VPN).\n" +
			"1.2. A aceitação desta Oferta é considerada efetiva ao iniciar o uso do bot do Telegram, pressionar o botão \"Start\" ou pagar pelo plano tarifário.\n\n" +
			"<b>2. Objeto da Oferta</b>\n" +
			"2.1. O Fornecedor concede ao Cliente um direito não exclusivo (licença) para usar configurações de acesso (chaves VLESS/Hysteria 2/MTProto) aos servidores do Fornecedor.\n" +
			"2.2. O serviço é considerado prestado no momento em que a chave de acesso ou o arquivo de configuração é enviado ao Cliente de forma automática através do bot do Telegram.\n\n" +
			"<b>3. Direitos e Obrigações das Partes</b>\n" +
			"3.1. O Fornecedor compromete-se a manter a operabilidade dos servidores, no entanto, não garante 100% de disponibilidade (Uptime) em caso de bloqueios por parte dos provedores de internet do Cliente ou órgãos estatais (força maior).\n" +
			"3.2. O Cliente compromete-se a usar os serviços exclusivamente para fins legais.\n" +
			"3.3. O Cliente está estritamente proibido de:\n" +
			"• Usar o serviço para enviar spam (e-mail, redes sociais).\n" +
			"• Escanear portas, realizar ataques DDoS, brute force e hackear recursos.\n" +
			"• Distribuir software malicioso, substâncias narcóticas, pornografia infantil ou materiais que promovam o terrorismo.\n" +
			"• Transferir chaves de acesso para terceiros (a menos que previsto na tarifa).\n\n" +
			"<b>4. Responsabilidade e Bloqueio</b>\n" +
			"4.1. Em caso de violação da cláusula 3.3, o Fornecedor tem o direito de bloquear imediatamente o acesso do Cliente sem reembolso de fundos.\n" +
			"4.2. O Fornecedor não é responsável por danos, diretos ou indiretos, sofridos pelo Cliente em conexão com o uso ou impossibilidade de uso do serviço.\n" +
			"4.3. O serviço é fornecido nas condições \"Como está\" (As Is). O Fornecedor não garante que o serviço seja compatível com o dispositivo específico ou provedor do Cliente.",
	},
	language.Italiana: {
		Terms: "<b>OFFERTA PUBBLICA:</b>\n\n" +
			"<b>1. Disposizioni Generali</b>\n" +
			"1.1. Questo documento è la proposta ufficiale (offerta pubblica) del servizio THE BEYOND (di seguito \"Fornitore\") e contiene tutte le condizioni essenziali per la fornitura di servizi per l'organizzazione dell'accesso remoto ai nodi di rete (proxy/VPN).\n" +
			"1.2. L'accettazione di questa Offerta è considerata effettiva all'inizio dell'uso del bot Telegram, premendo il pulsante \"Start\" o pagando il piano tariffario.\n\n" +
			"<b>2. Oggetto dell'Offerta</b>\n" +
			"2.1. Il Fornitore concede al Cliente un diritto non esclusivo (licenza) per utilizzare configurazioni di accesso (chiavi VLESS/Hysteria 2/MTProto) ai server del Fornitore.\n" +
			"2.2. Il servizio è considerato reso nel momento in cui la chiave di accesso o il file di configurazione viene inviato al Cliente in modo automatico tramite il bot Telegram.\n\n" +
			"<b>3. Diritti e Obblighi delle Parti</b>\n" +
			"3.1. Il Fornitore si impegna a mantenere la funzionalità dei server, tuttavia non garantisce il 100% di disponibilità (Uptime) in caso di blocchi da parte dei provider internet del Cliente o autorità statali (forza maggiore).\n" +
			"3.2. Il Cliente si impegna a utilizzare i servizi esclusivamente per scopi legali.\n" +
			"3.3. Al Cliente è rigorosamente proibito:\n" +
			"• Utilizzare il servizio per inviare spam (e-mail, social network).\n" +
			"• Scansionare porte, condurre attacchi DDoS, brute force e hacking di risorse.\n" +
			"• Distribuire software dannoso, sostanze narcotiche, pornografia infantile o materiali che promuovono il terrorismo.\n" +
			"• Trasferire chiavi di accesso a terzi (a meno che non sia previsto dalla tariffa).\n\n" +
			"<b>4. Responsabilità e Blocco</b>\n" +
			"4.1. In caso di violazione del punto 3.3, il Fornitore ha il diritto di bloccare immediatamente l'accesso del Cliente senza rimborso dei fondi.\n" +
			"4.2. Il Fornitore non è responsabile per danni, diretti o indiretti, subiti dal Cliente in connessione con l'uso o l'impossibilità di usare il servizio.\n" +
			"4.3. Il servizio è fornito su base \"Come è\" (As Is). Il Fornitore non garantisce che il servizio sia compatibile con il dispositivo specifico o provider del Cliente.",
	},
	language.Русский: {
		Terms: "<b>ПУБЛИЧНАЯ ОФЕРТА:</b>\n\n" +
			"<b>1. Общие положения</b>\n" +
			"1.1. Настоящий документ является официальным предложением (публичной офертой) сервиса THE BEYOND (далее — «Исполнитель») и содержит все существенные условия предоставления услуг по организации удаленного доступа к сетевым узлам (прокси/VPN).\n" +
			"1.2. Акцептом (принятием) настоящей Оферты считается начало использования Telegram-бота, нажатие кнопки «Start» или оплата тарифного плана.\n\n" +
			"<b>2. Предмет оферты</b>\n" +
			"2.1. Исполнитель предоставляет Заказчику неисключительное право (лицензию) на использование конфигураций доступа (ключей VLESS/Hysteria 2/MTProto) к серверам Исполнителя.\n" +
			"2.2. Услуга считается оказанной в момент отправки Заказчику ключа доступа или конфигурационного файла в автоматическом режиме через Telegram-бота.\n\n" +
			"<b>3. Права и обязанности сторон</b>\n" +
			"3.1. Исполнитель обязуется поддерживать работоспособность серверов, однако не гарантирует 100% доступность (Uptime) в случае блокировок со стороны интернет-провайдеров Заказчика или государственных органов (форс-мажор).\n" +
			"3.2. Заказчик обязуется использовать услуги исключительно в законных целях.\n" +
			"3.3. Заказчику категорически запрещено:\n" +
			"• Использовать сервис для рассылки спама (email, соцсети).\n" +
			"• Сканировать порты, осуществлять DDoS-атаки, брутфорс и взлом ресурсов.\n" +
			"• Распространять вредоносное ПО, наркотические вещества, детскую порнографию или материалы, пропагандирующие терроризм.\n" +
			"• Передавать ключи доступа третьим лицам (если это не предусмотрено тарифом).\n\n" +
			"<b>4. Ответственность и блокировка</b>\n" +
			"4.1. В случае нарушения п. 3.3, Исполнитель имеет право мгновенно заблокировать доступ Заказчика без возврата денежных средств.\n" +
			"4.2. Исполнитель не несет ответственности за ущерб, прямой или косвенный, понесенный Заказчиком в связи с использованием или невозможностью использования сервиса.\n" +
			"4.3. Сервис предоставляется на условиях «Как есть» (As Is). Исполнитель не гарантирует, что сервис будет совместим с конкретным устройством или провайдером Заказчика.",
	},
	language.Українська: {
		Terms: "<b>ПУБЛІЧНА ОФЕРТА:</b>\n\n" +
			"<b>1. Загальні положення</b>\n" +
			"1.1. Цей документ є офіційною пропозицією (публічною офертою) сервісу THE BEYOND (далі — «Виконавець») і містить усі істотні умови надання послуг з організації віддаленого доступу до мережних вузлів (проксі/VPN).\n" +
			"1.2. Акцептом (прийняттям) цієї Оферти вважається початок використання Telegram-бота, натискання кнопки «Start» або оплата тарифного плану.\n\n" +
			"<b>2. Предмет оферти</b>\n" +
			"2.1. Виконавець надає Замовнику невиключне право (ліцензію) на використання конфігурацій доступу (ключів VLESS/Hysteria 2/MTProto) до серверів Виконавця.\n" +
			"2.2. Послуга вважається наданою в момент надсилання Замовнику ключа доступу або конфігураційного файлу в автоматичному режимі через Telegram-бота.\n\n" +
			"<b>3. Права та обов'язки сторін</b>\n" +
			"3.1. Виконавець зобов'язується підтримувати працездатність серверів, однак не гарантує 100% доступність (Uptime) у разі блокувань з боку інтернет-провайдерів Замовника або державних органів (форс-мажор).\n" +
			"3.2. Замовник зобов'язується використовувати послуги виключно в законних цілях.\n" +
			"3.3. Замовнику категорично заборонено:\n" +
			"• Використовувати сервіс для розсилки спаму (email, соцмережі).\n" +
			"• Сканувати порти, здійснювати DDoS-атаки, брутфорс та злом ресурсів.\n" +
			"• Поширювати шкідливе ПЗ, наркотичні речовини, дитячу порнографію або матеріали, що пропагують тероризм.\n" +
			"• Передавати ключі доступу третім особам (якщо це не передбачено тарифом).\n\n" +
			"<b>4. Відповідальність та блокування</b>\n" +
			"4.1. У разі порушення п. 3.3, Виконавець має право миттєво заблокувати доступ Замовника без повернення коштів.\n" +
			"4.2. Виконавець не несе відповідальності за збитки, прямі або непрямі, понесені Замовником у зв'язку з використанням або неможливістю використання сервісу.\n" +
			"4.3. Сервіс надається на умовах «Як є» (As Is). Виконавець не гарантує, що сервіс буде сумісним з конкретним пристроєм або провайдером Замовника.",
	},
	language.Polski: {
		Terms: "<b>OFERTA PUBLICZNA:</b>\n\n" +
			"<b>1. Postanowienia Ogólne</b>\n" +
			"1.1. Niniejszy dokument stanowi oficjalną propozycję (ofertę publiczną) serwisu THE BEYOND (dalej \"Dostawca\") i zawiera wszystkie istotne warunki świadczenia usług organizacji zdalnego dostępu do węzłów sieciowych (proxy/VPN).\n" +
			"1.2. Akceptacja tej Oferty jest uważana za dokonaną w momencie rozpoczęcia korzystania z bota Telegram, naciśnięcia przycisku \"Start\" lub zapłaty za plan taryfowy.\n\n" +
			"<b>2. Przedmiot Oferty</b>\n" +
			"2.1. Dostawca udziela Klientowi niewyłącznego prawa (licencji) do korzystania z konfiguracji dostępu (kluczy VLESS/Hysteria 2/MTProto) do serwerów Dostawcy.\n" +
			"2.2. Usługa jest uważana za świadczoną w momencie wysłania Klientowi klucza dostępu lub pliku konfiguracyjnego w trybie automatycznym poprzez bota Telegram.\n\n" +
			"<b>3. Prawa i Obowiązki Stron</b>\n" +
			"3.1. Dostawca zobowiązuje się do utrzymania funkcjonalności serwerów, jednak nie gwarantuje 100% dostępności (Uptime) w przypadku blokad ze strony dostawców internetowych Klienta lub organów państwowych (siła wyższa).\n" +
			"3.2. Klient zobowiązuje się do korzystania z usług wyłącznie w celach zgodnych z prawem.\n" +
			"3.3. Klientowi surowo zabrania się:\n" +
			"• Używania serwisu do wysyłania spamu (e-mail, sieci społecznościowe).\n" +
			"• Skanowania portów, przeprowadzania ataków DDoS, brute force i hakowania zasobów.\n" +
			"• Rozpowszechniania złośliwego oprogramowania, substancji narkotycznych, pornografii dziecięcej lub materiałów propagujących terroryzm.\n" +
			"• Przekazywania kluczy dostępu osobom trzecim (chyba że przewidziane w taryfie).\n\n" +
			"<b>4. Odpowiedzialność i Blokada</b>\n" +
			"4.1. W przypadku naruszenia pkt 3.3, Dostawca ma prawo natychmiast zablokować dostęp Klienta bez zwrotu środków.\n" +
			"4.2. Dostawca nie ponosi odpowiedzialności za szkody, bezpośrednie lub pośrednie, poniesione przez Klienta w związku z korzystaniem lub niemożnością korzystania z serwisu.\n" +
			"4.3. Serwis jest świadczony na warunkach \"Jak jest\" (As Is). Dostawca nie gwarantuje, że serwis będzie kompatybilny z konkretnym urządzeniem lub dostawcą Klienta.",
	},
	language.Ceština: {
		Terms: "<b>VEŘEJNÁ NABÍDKA:</b>\n\n" +
			"<b>1. Obecná Ustanovení</b>\n" +
			"1.1. Tento dokument je oficiální nabídkou (veřejnou nabídkou) služby THE BEYOND (dále \"Poskytovatel\") a obsahuje všechny podstatné podmínky poskytování služeb pro organizaci vzdáleného přístupu k síťovým uzlům (proxy/VPN).\n" +
			"1.2. Přijetí této Nabídky se považuje za uskutečněné zahájením používání Telegram bota, stisknutím tlačítka \"Start\" nebo zaplacením tarifního plánu.\n\n" +
			"<b>2. Předmět Nabídky</b>\n" +
			"2.1. Poskytovatel poskytuje Zákazníkovi nevýhradní právo (licenci) k používání přístupových konfigurací (klíčů VLESS/Hysteria 2/MTProto) k serverům Poskytovatele.\n" +
			"2.2. Služba se považuje za poskytnutou v okamžiku odeslání přístupového klíče nebo konfiguračního souboru Zákazníkovi v automatickém režimu přes Telegram bota.\n\n" +
			"<b>3. Práva a Povinnosti Stran</b>\n" +
			"3.1. Poskytovatel se zavazuje udržovat funkčnost serverů, avšak negarantuje 100% dostupnost (Uptime) v případě blokád ze strany internetových poskytovatelů Zákazníka nebo státních orgánů (vyšší moc).\n" +
			"3.2. Zákazník se zavazuje používat služby výhradně k zákonným účelům.\n" +
			"3.3. Zákazníkovi je kategoricky zakázáno:\n" +
			"• Používat službu k rozesílání spamu (e-mail, sociální sítě).\n" +
			"• Skenovat porty, provádět DDoS útoky, brute force a hackování zdrojů.\n" +
			"• Šířit škodlivý software, narkotické látky, dětskou pornografii nebo materiály propagující terorismus.\n" +
			"• Předávat přístupové klíče třetím osobám (pokud to není předvídáno tarifem).\n\n" +
			"<b>4. Odpovědnost a Blokování</b>\n" +
			"4.1. V případě porušení bodu 3.3 má Poskytovatel právo okamžitě zablokovat přístup Zákazníka bez vrácení peněz.\n" +
			"4.2. Poskytovatel nenesou odpovědnost za škody, přímé nebo nepřímé, které Zákazník utrpí v souvislosti s používáním nebo nemožností používání služby.\n" +
			"4.3. Služba je poskytována na podmínkách \"Jak je\" (As Is). Poskytovatel negarantuje, že služba bude kompatibilní s konkrétním zařízením nebo poskytovatelem Zákazníka.",
	},
	language.Български: {
		Terms: "<b>ПУБЛИЧНА ОФЕРТА:</b>\n\n" +
			"<b>1. Общи Положения</b>\n" +
			"1.1. Този документ е официалното предложение (публична оферта) на услугата THE BEYOND (по-нататък \"Доставчик\") и съдържа всички съществени условия за предоставяне на услуги за организиране на отдалечен достъп до мрежови възли (прокси/VPN).\n" +
			"1.2. Приемането на тази Оферта се счита за осъществено с началото на използването на Telegram бота, натискане на бутона \"Start\" или плащане на тарифния план.\n\n" +
			"<b>2. Предмет на Офертата</b>\n" +
			"2.1. Доставчикът предоставя на Клиента неексклузивно право (лиценз) за използване на конфигурации за достъп (ключове VLESS/Hysteria 2/MTProto) към сървърите на Доставчика.\n" +
			"2.2. Услугата се счита за предоставена в момента на изпращане на ключа за достъп или конфигурационния файл на Клиента в автоматичен режим чрез Telegram бота.\n\n" +
			"<b>3. Права и Задължения на Страните</b>\n" +
			"3.1. Доставчикът се задължава да поддържа работоспособността на сървърите, но не гарантира 100% достъпност (Uptime) в случай на блокирания от страна на интернет доставчиците на Клиента или държавни органи (форсмажор).\n" +
			"3.2. Клиентът се задължава да използва услугите изключително за законни цели.\n" +
			"3.3. На Клиента е категорично забранено:\n" +
			"• Да използва услугата за изпращане на спам (имейл, социални мрежи).\n" +
			"• Да сканира портове, да провежда DDoS атаки, brute force и хакване на ресурси.\n" +
			"• Да разпространява вреден софтуер, наркотични вещества, детска порнография или материали, пропагандиращи тероризъм.\n" +
			"• Да предава ключове за достъп на трети лица (ако това не е предвидено в тарифата).\n\n" +
			"<b>4. Отговорност и Блокиране</b>\n" +
			"4.1. В случай на нарушение на т. 3.3, Доставчикът има право незабавно да блокира достъпа на Клиента без връщане на средства.\n" +
			"4.2. Доставчикът не носи отговорност за щети, директни или индиректни, понесени от Клиента във връзка с използването или невъзможността за използване на услугата.\n" +
			"4.3. Услугата се предоставя на условия \"Както е\" (As Is). Доставчикът не гарантира, че услугата ще бъде съвместима с конкретно устройство или доставчик на Клиента.",
	},
	language.Српски: {
		Terms: "<b>ЈАВНА ПОНУДА:</b>\n\n" +
			"<b>1. Општа Одредба</b>\n" +
			"1.1. Овај документ је званична понуда (јавна понуда) услуге THE BEYOND (у даљем тексту \"Провајдер\") и садржи све битне услове за пружање услуга за организовање даљинског приступа мрежним чворовима (прокси/VPN).\n" +
			"1.2. Прихватање ове Понуде се сматра оствареним почетком коришћења Telegram бота, притиском дугмета \"Start\" или плаћањем тарифног плана.\n\n" +
			"<b>2. Предмет Понуде</b>\n" +
			"2.1. Провајдер даје Клијенту неискључиво право (лиценцу) за коришћење конфигурација приступа (клјучева VLESS/Hysteria 2/MTProto) серверима Провајдера.\n" +
			"2.2. Услуга се сматра пруженом у тренутку слања клјуча приступа или конфигурационог фајла Клијенту у аутоматском режиму преко Telegram бота.\n\n" +
			"<b>3. Права и Обавезе Страна</b>\n" +
			"3.1. Провајдер се обавезује да одржава радоспособност сервера, међутим не гарантује 100% доступност (Uptime) у случају блокада са стране интернет провајдера Клијента или државних органа (виша сила).\n" +
			"3.2. Клијент се обавезује да користи услуге искључиво у законске сврхе.\n" +
			"3.3. Клијенту је категорички забрањено:\n" +
			"• Користити услугу за слање спама (е-маил, друштвене мреже).\n" +
			"• Скенирати портове, спровести DDoS нападе, brute force и хаковање ресурса.\n" +
			"• Ширити штетни софтвер, наркотичке супстанце, дечју порнографију или материјале који пропагирају тероризам.\n" +
			"• Преносити клјучеве приступа трећим лицима (ако то није предвиђено тарифом).\n\n" +
			"<b>4. Одговорност и Блокирање</b>\n" +
			"4.1. У случају кршења т. 3.3, Провајдер има право да одмах блокира приступ Клијента без повраћаја средстава.\n" +
			"4.2. Провајдер не сноси одговорност за штету, директну или индиректну, коју Клијент претрпи у вези са коришћењем или немогућношћу коришћења услуге.\n" +
			"4.3. Услуга се пружа на условима \"Како јесте\" (As Is). Провајдер не гарантује да ће услуга бити компатибилна са конкретним уређајем или провајдером Клијента.",
	},
	language.Hrvatski: {
		Terms: "<b>JAVNA PONUDA:</b>\n\n" +
			"<b>1. Opća Odredba</b>\n" +
			"1.1. Ovaj dokument je službena ponuda (javna ponuda) usluge THE BEYOND (u daljnjem tekstu \"Pružatelj\") i sadrži sve bitne uvjete za pružanje usluga za organizaciju daljinskog pristupa mrežnim čvorovima (proxy/VPN).\n" +
			"1.2. Prihvaćanje ove Ponude smatra se ostvarenim početkom korištenja Telegram bota, pritiskom gumba \"Start\" ili plaćanjem tarifnog plana.\n\n" +
			"<b>2. Predmet Ponude</b>\n" +
			"2.1. Pružatelj daje Klijentu neisključivo pravo (licencu) za korištenje konfiguracija pristupa (ključeva VLESS/Hysteria 2/MTProto) serverima Pružatelja.\n" +
			"2.2. Usluga se smatra pruženom u trenutku slanja ključa pristupa ili konfiguracijske datoteke Klijentu u automatskom modu preko Telegram bota.\n\n" +
			"<b>3. Prava i Obaveze Strana</b>\n" +
			"3.1. Pružatelj se obvezuje održavati radnu sposobnost servera, međutim ne jamči 100% dostupnost (Uptime) u slučaju blokada sa strane internet pružatelja Klijenta ili državnih tijela (viša sila).\n" +
			"3.2. Klijent se obvezuje koristiti usluge isključivo u zakonske svrhe.\n" +
			"3.3. Klijentu je kategorički zabranjeno:\n" +
			"• Koristiti uslugu za slanje spama (e-mail, društvene mreže).\n" +
			"• Skenirati portove, provoditi DDoS napade, brute force i hakiranje resursa.\n" +
			"• Širiti štetni softver, narkotičke supstance, dječju pornografiju ili materijale koji propagiraju terorizam.\n" +
			"• Prenositi ključeve pristupa trećim osobama (ako to nije predviđeno tarifom).\n\n" +
			"<b>4. Odgovornost i Blokiranje</b>\n" +
			"4.1. U slučaju kršenja t. 3.3, Pružatelj ima pravo odmah blokirati pristup Klijenta bez povrata sredstava.\n" +
			"4.2. Pružatelj ne snosi odgovornost za štetu, direktnu ili indirektnu, koju Klijent pretrpi u vezi s korištenjem ili nemogućnošću korištenja usluge.\n" +
			"4.3. Usluga se pruža na uvjetima \"Kako je\" (As Is). Pružatelj ne jamči da će usluga biti kompatibilna s konkretnim uređajem ili pružateljem Klijenta.",
	},
	language.Slovenčina: {
		Terms: "<b>VereJNÁ PONUKA:</b>\n\n" +
			"<b>1. Všeobecné Ustanovenia</b>\n" +
			"1.1. Tento dokument je oficiálnou ponukou (verejnou ponukou) služby THE BEYOND (ďalej \"Poskytovateľ\") a obsahuje všetky podstatné podmienky poskytovania služieb na organizáciu vzdialeného prístupu k sieťovým uzlom (proxy/VPN).\n" +
			"1.2. Prijatie tejto Ponuky sa považuje za uskutočnené začatím používania Telegram bota, stlačením tlačidla \"Start\" alebo zaplatením tarifného plánu.\n\n" +
			"<b>2. Predmet Ponuky</b>\n" +
			"2.1. Poskytovateľ poskytuje Zákazníkovi nevýhradné právo (licenciu) na používanie prístupových konfigurácií (klúčov VLESS/Hysteria 2/MTProto) k serverom Poskytovateľa.\n" +
			"2.2. Služba sa považuje za poskytnutú v momente odoslania prístupového klúča alebo konfiguračného súboru Zákazníkovi v automatickom režime cez Telegram bota.\n\n" +
			"<b>3. Práva a Povinnosti Strán</b>\n" +
			"3.1. Poskytovateľ sa zaväzuje udržiavať prevádzkyschopnosť serverov, avšak negarantuje 100% dostupnosť (Uptime) v prípade blokád zo strany internetových poskytovateľov Zákazníka alebo štátnych orgánov (vyššia moc).\n" +
			"3.2. Zákazník sa zaväzuje používať služby výlučne na zákonné účely.\n" +
			"3.3. Zákazníkovi je kategoricky zakázané:\n" +
			"• Používať službu na rozosielanie spamu (e-mail, sociálne siete).\n" +
			"• Skenovať porty, vykonávať DDoS útoky, brute force a hackovanie zdrojov.\n" +
			"• Šíriť škodlivý softvér, narkotické látky, detskú pornografiu alebo materiály propagujúce terorizmus.\n" +
			"• Prenášať prístupové klúče tretím osobám (ak to nie je predvídané tarifou).\n\n" +
			"<b>4. Zodpovednosť a Blokovanie</b>\n" +
			"4.1. V prípade porušenia bodu 3.3 má Poskytovateľ právo okamžite zablokovať prístup Zákazníka bez vrátenia peňazí.\n" +
			"4.2. Poskytovateľ nenesie zodpovednosť za škody, priame alebo nepriame, ktoré Zákazník utrpí v súvislosti s používaním alebo nemožnosťou používania služby.\n" +
			"4.3. Služba je poskytovaná na podmienkach \"Ako je\" (As Is). Poskytovateľ negarantuje, že služba bude kompatibilná s konkrétnym zariadením alebo poskytovateľom Zákazníka.",
	},
	language.Slovenski: {
		Terms: "<b>JAVNA PONUDBA:</b>\n\n" +
			"<b>1. Splošne Določbe</b>\n" +
			"1.1. Ta dokument je uradna ponudba (javna ponudba) storitve THE BEYOND (v nadaljevanju \"Ponudnik\") in vsebuje vse bistvene pogoje za zagotavljanje storitev za organizacijo oddaljenega dostopa do mrežnih vozlišč (proxy/VPN).\n" +
			"1.2. Sprejetje te Ponudbe se šteje za izvedeno z začetkom uporabe Telegram bota, pritiskom gumba \"Start\" ali plačilom tarifnega plana.\n\n" +
			"<b>2. Predmet Ponudbe</b>\n" +
			"2.1. Ponudnik podeli Stranki neizključno pravico (licenco) za uporabo dostopnih konfiguracij (ključev VLESS/Hysteria 2/MTProto) do strežnikov Ponudnika.\n" +
			"2.2. Storitev se šteje za zagotovljeno v trenutku pošiljanja dostopnega ključa ali konfiguracijske datoteke Stranki v avtomatičnem načinu prek Telegram bota.\n\n" +
			"<b>3. Pravice in Obveznosti Strank</b>\n" +
			"3.1. Ponudnik se zavezuje vzdrževati delovanje strežnikov, vendar ne jamči 100% razpoložljivosti (Uptime) v primeru blokad s strani internetnih ponudnikov Stranke ali državnih organov (višja sila).\n" +
			"3.2. Stranka se zavezuje uporabljati storitve izključno v zakonite namene.\n" +
			"3.3. Stranki je kategorično prepovedano:\n" +
			"• Uporabljati storitev za pošiljanje spam (e-pošta, socialna omrežja).\n" +
			"• Skenirati porte, izvajati DDoS napade, brute force in hekanje virov.\n" +
			"• Razširjati škodljivo programsko opremo, narkotične substance, otroško pornografijo ali materiale, ki promovirajo terorizem.\n" +
			"• Prenositi dostopne ključe tretjim osebam (če to ni predvideno v tarifi).\n\n" +
			"<b>4. Odgovornost in Blokiranje</b>\n" +
			"4.1. V primeru kršitve t. 3.3 ima Ponudnik pravico takoj blokirati dostop Stranke brez vračila sredstev.\n" +
			"4.2. Ponudnik ne nosi odgovornosti za škodo, neposredno ali posredno, ki jo Stranka utrpi v povezavi z uporabo ali nemožnostjo uporabe storitve.\n" +
			"4.3. Storitev se zagotavlja na pogojih \"Kot je\" (As Is). Ponudnik ne jamči, da bo storitev kompatibilna s konkretno napravo ali ponudnikom Stranke.",
	},
	language.Lietùvių: {
		Terms: "<b>VIEŠAS PASIŪLYMAS:</b>\n\n" +
			"<b>1. Bendrosios Nuostatos</b>\n" +
			"1.1. Šis dokumentas yra oficialus paslaugos THE BEYOND (toliau \"Teikėjas\") pasiūlymas (viešasis pasiūlymas) ir apima visas esmines sąlygas teikti paslaugas organizuojant nuotolinį prieigą prie tinklo mazgų (proxy/VPN).\n" +
			"1.2. Šio Pasiūlymo priėmimas laikomas įvykusiu pradėjus naudoti Telegram botą, paspaudus mygtuką \"Start\" arba sumokėjus už tarifinį planą.\n\n" +
			"<b>2. Pasiūlymo Objektas</b>\n" +
			"2.1. Teikėjas suteikia Klientui neišskirtinę teisę (licenciją) naudoti prieigos konfigūracijas (raktus VLESS/Hysteria 2/MTProto) prie Teikėjo serverių.\n" +
			"2.2. Paslauga laikoma suteikta momento, kai prieigos raktas arba konfigūracijos failas siunčiamas Klientui automatiniu režimu per Telegram botą.\n\n" +
			"<b>3. Šalių Teisės ir Pareigos</b>\n" +
			"3.1. Teikėjas įsipareigoja palaikyti serverių veikimą, tačiau negarantuoja 100% prieinamumo (Uptime) blokuojant iš Kliento interneto tiekėjų ar valstybinių institucijų pusės (force majeure).\n" +
			"3.2. Klientas įsipareigoja naudoti paslaugas tik teisėtais tikslais.\n" +
			"3.3. Klientui kategoriškai draudžiama:\n" +
			"• Naudoti paslaugą siunčiant šlamštą (el. paštas, socialiniai tinklai).\n" +
			"• Skenuoti portus, vykdyti DDoS atakas, brute force ir nulaužti išteklius.\n" +
			"• Platinti kenksmingą programinę įrangą, narkotines medžiagas, vaikų pornografiją ar medžiagas, propaguojančias terorizmą.\n" +
			"• Perduoti prieigos raktus trečiosioms šalims (jei tai nenumatyta tarife).\n\n" +
			"<b>4. Atsakomybė ir Blokavimas</b>\n" +
			"4.1. Pažeidus 3.3 punktą, Teikėjas turi teisę nedelsiant blokuoti Kliento prieigą be pinigų grąžinimo.\n" +
			"4.2. Teikėjas neatsako už žalą, tiesioginę ar netiesioginę, kurią Klientas patiria dėl paslaugos naudojimo ar negalėjimo naudoti.\n" +
			"4.3. Paslauga teikiama \"Kaip yra\" (As Is) sąlygomis. Teikėjas negarantuoja, kad paslauga bus suderinama su konkrečiu Kliento įrenginiu ar tiekėju.",
	},
	language.Latviešu: {
		Terms: "<b>PUBLIKĀ PIEDĀVĀJUMS:</b>\n\n" +
			"<b>1. Vispārējie Noteikumi</b>\n" +
			"1.1. Šis dokuments ir oficiāls piedāvājums (publisks piedāvājums) pakalpojumam THE BEYOND (turpmāk \"Piegādātājs\") un satur visas būtiskās nosacījumus pakalpojumu sniegšanai attālinātas piekļuves organizēšanai tīkla mezgliem (proxy/VPN).\n" +
			"1.2. Šī Piedāvājuma pieņemšana tiek uzskatīta par notikušu sākot izmantot Telegram botu, nospiežot pogu \"Start\" vai samaksājot par tarifu plānu.\n\n" +
			"<b>2. Piedāvājuma Priekšmets</b>\n" +
			"2.1. Piegādātājs piešķir Klientam neekskluzīvas tiesības (licenci) izmantot piekļuves konfigurācijas (atslēgas VLESS/Hysteria 2/MTProto) Piegādātāja serveriem.\n" +
			"2.2. Pakalpojums tiek uzskatīts par sniegtu brīdī, kad piekļuves atslēga vai konfigurācijas fails tiek nosūtīts Klientam automātiskā režīmā caur Telegram botu.\n\n" +
			"<b>3. Pušu Tiesības un Pienākumi</b>\n" +
			"3.1. Piegādātājs apņemas uzturēt serveru darbspēju, tomēr negarantē 100% pieejamību (Uptime) gadījumā, ja bloķēšana no Klienta interneta piegādātājiem vai valsts iestādēm (force majeure).\n" +
			"3.2. Klients apņemas izmantot pakalpojumus tikai likumīgiem mērķiem.\n" +
			"3.3. Klientam kategoriski aizliegts:\n" +
			"• Izmantot pakalpojumu spam (e-pasts, sociālie tīkli) sūtīšanai.\n" +
			"• Skenēt portus, veikt DDoS uzbrukumus, brute force un resursu uzlaušanu.\n" +
			"• Izplatīt kaitīgu programmatūru, narkotiskās vielas, bērnu pornogrāfiju vai materiālus, kas propagandē terorismu.\n" +
			"• Nodot piekļuves atslēgas trešajām personām (ja tas nav paredzēts tarifā).\n\n" +
			"<b>4. Atbildība un Bloķēšana</b>\n" +
			"4.1. Pārkāpjot 3.3 punktu, Piegādātājam ir tiesības nekavējoties bloķēt Klienta piekļuvi bez naudas atmaksas.\n" +
			"4.2. Piegādātājs neatbild par zaudējumiem, tiešiem vai netiešiem, ko Klients cietis saistībā ar pakalpojuma izmantošanu vai nespēju izmantot.\n" +
			"4.3. Pakalpojums tiek sniegts \"Kā ir\" (As Is) nosacījumos. Piegādātājs negarantē, ka pakalpojums būs savietojams ar konkrētu Klienta ierīci vai piegādātāju.",
	},
	language.Eesti: {
		Terms: "<b>AVALIK PAKKUMINE:</b>\n\n" +
			"<b>1. Üldised Sätted</b>\n" +
			"1.1. See dokument on teenuse THE BEYOND (edaspidi \"Pakkuja\") ametlik pakkumine (avalik pakkumine) ja sisaldab kõiki olulisi tingimusi teenuste osutamiseks kaugjuurdepääsu korraldamiseks võrgusõlmedele (proxy/VPN).\n" +
			"1.2. Käesoleva Pakkumise aktsepteerimine loetakse toimunuks Telegram boti kasutamise alustamisel, nupu \"Start\" vajutamisel või tariifiplaani maksmisel.\n\n" +
			"<b>2. Pakkumise Ese</b>\n" +
			"2.1. Pakkuja annab Kliendile mitteeksklusiivse õiguse (litsentsi) kasutada juurdepääsu konfiguratsioone (võtmeid VLESS/Hysteria 2/MTProto) Pakkuja serveritele.\n" +
			"2.2. Teenus loetakse osutatuks hetkel, kui juurdepääsu võti või konfiguratsioonifail saadetakse Kliendile automaatses režiimis läbi Telegram boti.\n\n" +
			"<b>3. Poolte Õigused ja Kohustused</b>\n" +
			"3.1. Pakkuja kohustub säilitama serverite töökorrasolekut, kuid ei garanteeri 100% kättesaadavust (Uptime) juhul, kui blokkeerimine Kliendi internetiteenuse pakkujate või riigiasutuste poolt (force majeure).\n" +
			"3.2. Klient kohustub kasutama teenuseid ainult seaduslike eesmärkidega.\n" +
			"3.3. Kliendile on kategooriliselt keelatud:\n" +
			"• Kasutada teenust spami saatmiseks (e-post, sotsiaalvõrgustikud).\n" +
			"• Skaneerida porte, teostada DDoS rünnakuid, brute force ja ressursside häkkimist.\n" +
			"• Levitada kahjulikku tarkvara, narkootilisi aineid, lastepornograafiat või materjale, mis propageerivad terrorismi.\n" +
			"• Edastada juurdepääsu võtmeid kolmandatele isikutele (kui see pole tariifis ette nähtud).\n\n" +
			"<b>4. Vastutus ja Blokeerimine</b>\n" +
			"4.1. Punkti 3.3 rikkumise korral on Pakkujal õigus kohe blokeerida Kliendi juurdepääs ilma raha tagastamiseta.\n" +
			"4.2. Pakkuja ei kanna vastutust kahjude eest, otseste või kaudsete eest, mida Klient kannatab seoses teenuse kasutamisega või võimetusega kasutada.\n" +
			"4.3. Teenus osutatakse tingimustel \"Nagu on\" (As Is). Pakkuja ei garanteeri, et teenus on ühilduv Kliendi konkreetse seadme või pakkujaga.",
	},
	language.Suomi: {
		Terms: "<b>JULKINEN TARJOUS:</b>\n\n" +
			"<b>1. Yleiset Määräykset</b>\n" +
			"1.1. Tämä asiakirja on palvelun THE BEYOND (jatkossa \"Tarjoaja\") virallinen ehdotus (julkinen tarjous) ja sisältää kaikki olennaiset ehdot palveluiden tarjoamiseen etäyhteyden järjestämiseksi verkko-solmuihin (proxy/VPN).\n" +
			"1.2. Tämän Tarjouksen hyväksyminen katsotaan tapahtuneeksi Telegram-botin käytön aloittamisella, \"Start\"-painikkeen painamisella tai tariffisuunnitelman maksamisella.\n\n" +
			"<b>2. Tarjouksen Kohde</b>\n" +
			"2.1. Tarjoaja myöntää Asiakkaalle ei-yksinoikeudellisen oikeuden (lisenssin) käyttää pääsykonfiguraatioita (avaimia VLESS/Hysteria 2/MTProto) Tarjoajan palvelimille.\n" +
			"2.2. Palvelu katsotaan toimitetuksi hetkellä, jolloin pääsyavain tai konfiguraatiotiedosto lähetetään Asiakkaalle automaattisessa tilassa Telegram-botin kautta.\n\n" +
			"<b>3. Osapuolten Oikeudet ja Velvollisuudet</b>\n" +
			"3.1. Tarjoaja sitoutuu ylläpitämään palvelimien toimivuutta, mutta ei takaa 100% saatavuutta (Uptime) tapauksessa, jossa estot Asiakkaan internet-palveluntarjoajilta tai valtion elimiltä (force majeure).\n" +
			"3.2. Asiakas sitoutuu käyttämään palveluita yksinomaan laillisiin tarkoituksiin.\n" +
			"3.3. Asiakkaalle on ehdottomasti kiellettyä:\n" +
			"• Käyttää palvelua roskapostin lähettämiseen (sähköposti, sosiaaliset verkostot).\n" +
			"• Skannata portteja, suorittaa DDoS-hyökkäyksiä, brute force ja resurssien hakkerointia.\n" +
			"• Levittää haittaohjelmistoja, narkoottisia aineita, lapsipornografiaa tai materiaaleja, jotka edistävät terrorismia.\n" +
			"• Siirtää pääsyavaimia kolmansille osapuolille (ellei tariffi salli sitä).\n\n" +
			"<b>4. Vastuu ja Estäminen</b>\n" +
			"4.1. Kohdan 3.3 rikkomuksen tapauksessa Tarjoajalla on oikeus välittömästi estää Asiakkaan pääsy ilman rahan palautusta.\n" +
			"4.2. Tarjoaja ei ole vastuussa vahingoista, suorista tai epäsuorista, joita Asiakas kärsii yhteydessä palvelun käyttöön tai kyvyttömyyteen käyttää.\n" +
			"4.3. Palvelu toimitetaan \"Sellaisenaan\" (As Is) ehdoilla. Tarjoaja ei takaa, että palvelu on yhteensopiva Asiakkaan tietyn laitteen tai tarjoajan kanssa.",
	},
	language.Ελληνικά: {
		Terms: "<b>ΔΗΜΟΣΙΑ ΠΡΟΣΦΟΡΑ:</b>\n\n" +
			"<b>1. Γενικές Διατάξεις</b>\n" +
			"1.1. Αυτό το έγγραφο αποτελεί την επίσημη πρόταση (δημόσια προσφορά) της υπηρεσίας THE BEYOND (εφεξής \"Πάροχος\") και περιέχει όλους τους ουσιώδεις όρους για την παροχή υπηρεσιών οργάνωσης απομακρυσμένης πρόσβασης σε κόμβους δικτύου (proxy/VPN).\n" +
			"1.2. Η αποδοχή αυτής της Προσφοράς θεωρείται πραγματοποιημένη με την έναρξη χρήσης του bot Telegram, πατώντας το κουμπί \"Start\" ή πληρώνοντας το τιμολόγιο.\n\n" +
			"<b>2. Αντικείμενο της Προσφοράς</b>\n" +
			"2.1. Ο Πάροχος παρέχει στον Πελάτη μη αποκλειστικό δικαίωμα (άδεια) χρήσης διαμορφώσεων πρόσβασης (κλειδιών VLESS/Hysteria 2/MTProto) στους διακομιστές του Παρόχου.\n" +
			"2.2. Η υπηρεσία θεωρείται παρεχόμενη τη στιγμή αποστολής του κλειδιού πρόσβασης ή του αρχείου διαμόρφωσης στον Πελάτη σε αυτόματη λειτουργία μέσω του bot Telegram.\n\n" +
			"<b>3. Δικαιώματα και Υποχρεώσεις των Μερών</b>\n" +
			"3.1. Ο Πάροχος δεσμεύεται να διατηρεί τη λειτουργικότητα των διακομιστών, ωστόσο δεν εγγυάται 100% διαθεσιμότητα (Uptime) σε περίπτωση αποκλεισμών από τους παρόχους internet του Πελάτη ή κρατικά όργανα (ανωτέρα βία).\n" +
			"3.2. Ο Πελάτης δεσμεύεται να χρησιμοποιεί τις υπηρεσίες αποκλειστικά για νόμιμους σκοπούς.\n" +
			"3.3. Στον Πελάτη απαγορεύεται κατηγορηματικά:\n" +
			"• Να χρησιμοποιεί την υπηρεσία για αποστολή spam (e-mail, κοινωνικά δίκτυα).\n" +
			"• Να σαρώνει θύρες, να διεξάγει επιθέσεις DDoS, brute force και hacking πόρων.\n" +
			"• Να διανέμει κακόβουλο λογισμικό, ναρκωτικές ουσίες, παιδική πορνογραφία ή υλικά που προωθούν την τρομοκρατία.\n" +
			"• Να μεταβιβάζει κλειδιά πρόσβασης σε τρίτους (εκτός αν προβλέπεται από το τιμολόγιο).\n\n" +
			"<b>4. Ευθύνη και Αποκλεισμός</b>\n" +
			"4.1. Σε περίπτωση παραβίασης του σημείου 3.3, ο Πάροχος έχει το δικαίωμα να αποκλείσει αμέσως την πρόσβαση του Πελάτη χωρίς επιστροφή χρημάτων.\n" +
			"4.2. Ο Πάροχος δεν φέρει ευθύνη για ζημιές, άμεσες ή έμμεσες, που υπέστη ο Πελάτης σε σχέση με τη χρήση ή την αδυναμία χρήσης της υπηρεσίας.\n" +
			"4.3. Η υπηρεσία παρέχεται υπό όρους \"Όπως είναι\" (As Is). Ο Πάροχος δεν εγγυάται ότι η υπηρεσία θα είναι συμβατή με συγκεκριμένη συσκευή ή πάροχο του Πελάτη.",
	},
	language.Română: {
		Terms: "<b>OFERTĂ PUBLICĂ:</b>\n\n" +
			"<b>1. Dispoziții Generale</b>\n" +
			"1.1. Acest document reprezintă propunerea oficială (ofertă publică) a serviciului THE BEYOND (denumit în continuare \"Furnizor\") și conține toate condițiile esențiale pentru furnizarea serviciilor de organizare a accesului remote la noduri de rețea (proxy/VPN).\n" +
			"1.2. Acceptarea acestei Oferte este considerată efectuată la începutul utilizării botului Telegram, apăsând butonul \"Start\" sau plătind planul tarifar.\n\n" +
			"<b>2. Obiectul Ofertei</b>\n" +
			"2.1. Furnizorul acordă Clientului un drept neexclusiv (licență) de a utiliza configurațiile de acces (chei VLESS/Hysteria 2/MTProto) la serverele Furnizorului.\n" +
			"2.2. Serviciul este considerat furnizat în momentul trimiterii cheii de acces sau a fișierului de configurație către Client în mod automat prin botul Telegram.\n\n" +
			"<b>3. Drepturi și Obligații ale Părților</b>\n" +
			"3.1. Furnizorul se angajează să mențină funcționalitatea serverelor, însă nu garantează 100% disponibilitate (Uptime) în cazul blocărilor din partea furnizorilor de internet ai Clientului sau a organelor de stat (forță majoră).\n" +
			"3.2. Clientul se angajează să utilizeze serviciile exclusiv în scopuri legale.\n" +
			"3.3. Clientului îi este strict interzis:\n" +
			"• Să utilizeze serviciul pentru trimiterea de spam (e-mail, rețele sociale).\n" +
			"• Să scaneze porturi, să efectueze atacuri DDoS, brute force și hacking de resurse.\n" +
			"• Să distribuie software rău intenționat, substanțe narcotice, pornografie infantilă sau materiale care promovează terorismul.\n" +
			"• Să transfere chei de acces către terți (cu excepția cazului în care este prevăzut în tarif).\n\n" +
			"<b>4. Răspundere și Blocare</b>\n" +
			"4.1. În cazul încălcării pct. 3.3, Furnizorul are dreptul să blocheze imediat accesul Clientului fără rambursarea banilor.\n" +
			"4.2. Furnizorul nu poartă răspundere pentru daune, directe sau indirecte, suferite de Client în legătură cu utilizarea sau imposibilitatea utilizării serviciului.\n" +
			"4.3. Serviciul este furnizat pe condiții \"Așa cum este\" (As Is). Furnizorul nu garantează că serviciul va fi compatibil cu dispozitivul specific sau furnizorul Clientului.",
	},
	language.Magyar: {
		Terms: "<b>NYILVÁNOS AJÁNLAT:</b>\n\n" +
			"<b>1. Általános Rendelkezések</b>\n" +
			"1.1. Ez a dokumentum a THE BEYOND szolgáltatás (a továbbiakban \"Szolgáltató\") hivatalos javaslata (nyilvános ajánlat) és tartalmazza az összes lényeges feltételt a szolgáltatások nyújtásához a hálózati csomópontok (proxy/VPN) távoli hozzáférésének szervezéséhez.\n" +
			"1.2. Ezen Ajánlat elfogadása akkor tekinthető megtörténtnek, amikor megkezdődik a Telegram bot használata, megnyomják a \"Start\" gombot vagy fizetnek a tarifális tervért.\n\n" +
			"<b>2. Az Ajánlat Tárgya</b>\n" +
			"2.1. A Szolgáltató nem kizárólagos jogot (licencet) ad az Ügyfélnek a hozzáférési konfigurációk (VLESS/Hysteria 2/MTProto kulcsok) használatára a Szolgáltató szervereihez.\n" +
			"2.2. A szolgáltatás nyújtottnak tekinthető abban a pillanatban, amikor a hozzáférési kulcs vagy konfigurációs fájl automatikus módban elküldésre kerül az Ügyfélnek a Telegram boton keresztül.\n\n" +
			"<b>3. A Felek Jogai és Kötelezettségei</b>\n" +
			"3.1. A Szolgáltató vállalja a szerverek működőképességének fenntartását, azonban nem garantálja a 100%-os elérhetőséget (Uptime) az Ügyfél internet-szolgáltatói vagy állami szervek általi blokkolások esetén (vis maior).\n" +
			"3.2. Az Ügyfél vállalja, hogy a szolgáltatásokat kizárólag törvényes célokra használja.\n" +
			"3.3. Az Ügyfélnek szigorúan tilos:\n" +
			"• A szolgáltatást spam küldésére használni (e-mail, közösségi hálózatok).\n" +
			"• Portokat szkennelni, DDoS támadásokat végrehajtani, brute force-t és erőforrások feltörését.\n" +
			"• Káros szoftvereket, narkotikus anyagokat, gyermekpornográfiát vagy terrorizmust propagáló anyagokat terjeszteni.\n" +
			"• Hozzáférési kulcsokat harmadik feleknek átadni (kivéve, ha a tarifa ezt lehetővé teszi).\n\n" +
			"<b>4. Felelősség és Blokkolás</b>\n" +
			"4.1. A 3.3 pont megsértése esetén a Szolgáltatónak joga van azonnal blokkolni az Ügyfél hozzáférését pénz visszatérítés nélkül.\n" +
			"4.2. A Szolgáltató nem felel a károkért, közvetlenekért vagy közvetettekért, amelyeket az Ügyfél szenved a szolgáltatás használatával vagy használatának lehetetlenségével kapcsolatban.\n" +
			"4.3. A szolgáltatás \"Ahogy van\" (As Is) feltételekkel kerül nyújtásra. A Szolgáltató nem garantálja, hogy a szolgáltatás kompatibilis az Ügyfél konkrét eszközével vagy szolgáltatójával.",
	},
	language.Arabic: {
		Terms: "<b>عرض عام:</b>\n\n" +
			"<b>1. الأحكام العامة</b>\n" +
			"1.1. يمثل هذا الوثيقة العرض الرسمي (العرض العام) لخدمة THE BEYOND (يشار إليها فيما بعد بـ \"المزود\") ويحتوي على جميع الشروط الأساسية لتقديم الخدمات لتنظيم الوصول عن بعد إلى عقد الشبكة (بروكسي/VPN).\n" +
			"1.2. يُعتبر قبول هذا العرض قد تم عند بدء استخدام روبوت تيليجرام، الضغط على زر \"Start\" أو دفع خطة التعريفة.\n\n" +
			"<b>2. موضوع العرض</b>\n" +
			"2.1. يمنح المزود العميل حقًا غير حصري (ترخيص) لاستخدام تكوينات الوصول (مفاتيح VLESS/Hysteria 2/MTProto) إلى خوادم المزود.\n" +
			"2.2. تُعتبر الخدمة مقدمة في اللحظة التي يتم فيها إرسال مفتاح الوصول أو ملف التكوين إلى العميل في وضع تلقائي عبر روبوت تيليجرام.\n\n" +
			"<b>3. حقوق وواجبات الأطراف</b>\n" +
			"3.1. يلتزم المزود بحفظ وظائف الخوادم، ومع ذلك لا يضمن توافرًا بنسبة 100% (Uptime) في حالة الحظر من جانب مزودي الإنترنت للعميل أو الجهات الحكومية (قوة قاهرة).\n" +
			"3.2. يلتزم العميل باستخدام الخدمات حصريًا لأغراض قانونية.\n" +
			"3.3. يُمنع على العميل بشكل قاطع:\n" +
			"• استخدام الخدمة لإرسال البريد المزعج (البريد الإلكتروني، الشبكات الاجتماعية).\n" +
			"• مسح المنافذ، تنفيذ هجمات DDoS، brute force واختراق الموارد.\n" +
			"• توزيع البرمجيات الضارة، المواد المخدرة، الإباحية الطفلية أو المواد التي تروج للإرهاب.\n" +
			"• نقل مفاتيح الوصول إلى أطراف ثالثة (إلا إذا كان ذلك منصوصًا عليه في التعريفة).\n\n" +
			"<b>4. المسؤولية والحظر</b>\n" +
			"4.1. في حالة انتهاك البند 3.3، يحق للمزود حظر الوصول للعميل فورًا دون استرداد الأموال.\n" +
			"4.2. لا يتحمل المزود مسؤولية الأضرار، المباشرة أو غير المباشرة، التي يتعرض لها العميل فيما يتعلق باستخدام الخدمة أو عدم القدرة على استخدامها.\n" +
			"4.3. يتم تقديم الخدمة على شروط \"كما هي\" (As Is). لا يضمن المزود أن الخدمة ستكون متوافقة مع جهاز معين أو مزود للعميل.",
	},
	language.Farsi: {
		Terms: "<b>پیشنهاد عمومی:</b>\n\n" +
			"<b>1. مقررات عمومی</b>\n" +
			"1.1. این سند پیشنهاد رسمی (پیشنهاد عمومی) سرویس THE BEYOND (از این پس \"ارائه‌دهنده\") است و شامل تمام شرایط اساسی برای ارائه خدمات سازماندهی دسترسی از راه دور به گره‌های شبکه (پروکسی/VPN) می‌شود.\n" +
			"1.2. پذیرش این پیشنهاد زمانی در نظر گرفته می‌شود که استفاده از ربات تلگرام شروع شود، دکمه \"Start\" فشار داده شود یا پرداخت برای طرح تعرفه انجام شود.\n\n" +
			"<b>2. موضوع پیشنهاد</b>\n" +
			"2.1. ارائه‌دهنده به مشتری حق غیرانحصاری (مجوز) استفاده از پیکربندی‌های دسترسی (کلیدهای VLESS/Hysteria 2/MTProto) به سرورهای ارائه‌دهنده را می‌دهد.\n" +
			"2.2. سرویس در لحظه ارسال کلید دسترسی یا فایل پیکربندی به مشتری در حالت خودکار از طریق ربات تلگرام، ارائه‌شده در نظر گرفته می‌شود.\n\n" +
			"<b>3. حقوق و تعهدات طرفین</b>\n" +
			"3.1. ارائه‌دهنده متعهد می‌شود کارایی سرورها را حفظ کند، اما 100% در دسترس بودن (Uptime) را در صورت مسدود شدن از سوی ارائه‌دهندگان اینترنت مشتری یا ارگان‌های دولتی (فورس ماژور) تضمین نمی‌کند.\n" +
			"3.2. مشتری متعهد می‌شود از خدمات فقط برای اهداف قانونی استفاده کند.\n" +
			"3.3. به مشتری به طور قاطع ممنوع است:\n" +
			"• استفاده از سرویس برای ارسال اسپم (ایمیل، شبکه‌های اجتماعی).\n" +
			"• اسکن پورت‌ها، انجام حملات DDoS، brute force و هک منابع.\n" +
			"• توزیع نرم‌افزارهای مخرب، مواد مخدر، پورنوگرافی کودکان یا مواد تبلیغ‌کننده تروریسم.\n" +
			"• انتقال کلیدهای دسترسی به اشخاص ثالث (مگر اینکه در تعرفه پیش‌بینی شده باشد).\n\n" +
			"<b>4. مسئولیت و مسدود کردن</b>\n" +
			"4.1. در صورت نقض بند 3.3، ارائه‌دهنده حق دارد دسترسی مشتری را بلافاصله بدون بازپرداخت پول مسدود کند.\n" +
			"4.2. ارائه‌دهنده مسئولیتی در قبال خسارات، مستقیم یا غیرمستقیم، که مشتری در ارتباط با استفاده یا عدم امکان استفاده از سرویس متحمل می‌شود، ندارد.\n" +
			"4.3. سرویس بر اساس شرایط \"همان‌طور که هست\" (As Is) ارائه می‌شود. ارائه‌دهنده تضمین نمی‌کند که سرویس با دستگاه خاص یا ارائه‌دهنده مشتری سازگار باشد.",
	},
	language.Türkçe: {
		Terms: "<b>KAMU TEKLİFİ:</b>\n\n" +
			"<b>1. Genel Hükümler</b>\n" +
			"1.1. Bu belge, THE BEYOND hizmetinin (bundan sonra \"Sağlayıcı\" olarak anılacaktır) resmi teklifi (kamu teklifi) olup, ağ düğümlerine (proxy/VPN) uzaktan erişim organizasyonu için hizmetlerin sağlanmasına ilişkin tüm temel koşulları içerir.\n" +
			"1.2. Bu Teklifin kabulü, Telegram botunun kullanımına başlanması, \"Start\" düğmesine basılması veya tarife planının ödemesiyle kabul edilmiş sayılır.\n\n" +
			"<b>2. Teklifin Konusu</b>\n" +
			"2.1. Sağlayıcı, Müşteriye Sağlayıcının sunucularına erişim konfigürasyonlarını (VLESS/Hysteria 2/MTProto anahtarları) kullanma için münhasır olmayan hakkı (lisans) verir.\n" +
			"2.2. Hizmet, erişim anahtarı veya konfigürasyon dosyasının Telegram botu aracılığıyla otomatik modda Müşteriye gönderildiği anda sağlanmış sayılır.\n\n" +
			"<b>3. Haklar ve Yükümlülükler</b>\n" +
			"3.1. Sağlayıcı, sunucuların performansını korumayı taahhüt eder; ancak müşterinin internet sağlayıcıları veya devlet kurumları tarafından engellenmesi durumunda (mücbir sebep) %100 erişilebilirlik (uptime) garantisi vermez.\n" +
			"3.2. Müşteri, hizmeti yalnızca yasal amaçlarla kullanmayı taahhüt eder.\n" +
			"3.3. Müşteriye kesinlikle yasaktır:\n" +
			"• Hizmeti spam gönderimi için kullanmak (e-posta, sosyal ağlar).\n" +
			"• Port taraması yapmak, DDoS saldırıları düzenlemek, brute force veya başka tür saldırılarla kaynaklara izinsiz erişim sağlamak.\n" +
			"• Kötü amaçlı yazılımlar, uyuşturucu maddeler, çocuk pornografisi veya terör propagandası içeren materyallerin dağıtımı.\n" +
			"• Erişim anahtarlarını üçüncü şahıslara devretmek (tarifede aksi belirtilmedikçe).\n\n" +
			"<b>4. Sorumluluk ve Engelleme</b>\n" +
			"4.1. 3.3 maddesinin ihlali halinde, Sağlayıcı müşterinin erişimini derhal ve para iadesi yapmadan engelleme hakkına sahiptir.\n" +
			"4.2. Sağlayıcı, hizmetin kullanımı veya kullanılamaması nedeniyle müşterinin maruz kaldığı doğrudan veya dolaylı zararlardan sorumlu değildir.\n" +
			"4.3. Hizmet \"olduğu gibi\" (As Is) sunulur. Sağlayıcı, hizmetin belirli bir cihaz veya müşterinin sağlayıcısıyla uyumlu olacağını garanti etmez.",
	},
	language.Hebrew: {
		Terms: "<b>הצעה ציבורית:</b>\n\n" +
			"<b>1. הוראות כלליות</b>\n" +
			"1.1. מסמך זה מהווה הצעה רשמית (הצעה ציבורית) של השירות THE BEYOND (להלן \"הספק\") ומכיל את כל התנאים החיוניים למתן שירותים לארגון גישה מרחוק לצמתי רשת (פרוקסי/VPN).\n" +
			"1.2. קבלת ההצעה הזו נחשבת כמתבצעת עם תחילת השימוש בבוט טלגרם, לחיצה על כפתור \"Start\" או תשלום עבור תוכנית התעריף.\n\n" +
			"<b>2. נושא ההצעה</b>\n" +
			"2.1. הספק מעניק ללקוח זכות לא בלעדית (רישיון) להשתמש בתצורות גישה (מפתחות VLESS/Hysteria 2/MTProto) לשרתים של הספק.\n" +
			"2.2. השירות נחשב כמסופק ברגע שליחת מפתח הגישה או קובץ התצורה ללקוח במצב אוטומטי דרך בוט טלגרם.\n\n" +
			"<b>3. זכויות וחובות הצדדים</b>\n" +
			"3.1. הספק מתחייב לשמור על תפקוד השרתים, אולם אינו מבטיח 100% זמינות (Uptime) במקרה של חסימות מצד ספקי האינטרנט של הלקוח או רשויות מדינה (כוח עליון).\n" +
			"3.2. הלקוח מתחייב להשתמש בשירותים אך ורק למטרות חוקיות.\n" +
			"3.3. ללקוח אסור בהחלט:\n" +
			"• להשתמש בשירות לשליחת ספאם (דוא\"ל, רשתות חברתיות).\n" +
			"• לסרוק פורטים, לבצע התקפות DDoS, brute force ופריצה למשאבים.\n" +
			"• להפיץ תוכנה זדונית, חומרים נרקוטיים, פורנוגרפיה ילדים או חומרים המקדמים טרור.\n" +
			"• להעביר מפתחות גישה לצדדים שלישיים (אלא אם כן צוין בתעריף).\n\n" +
			"<b>4. אחריות וחסימה</b>\n" +
			"4.1. במקרה של הפרה של סעיף 3.3, לספק יש זכות לחסום מיד את הגישה של הלקוח ללא החזר כספי.\n" +
			"4.2. הספק אינו אחראי לנזקים, ישירים או עקיפים, שנגרמו ללקוח בקשר לשימוש או אי-אפשרות שימוש בשירות.\n" +
			"4.3. השירות מסופק בתנאים \"כפי שהוא\" (As Is). הספק אינו מבטיח שהשירות יהיה תואם למכשיר ספציפי או ספק של הלקוח.",
	},
	language.ZH中文: {
		Terms: "<b>公开要约：</b>\n\n" +
			"<b>1. 一般规定</b>\n" +
			"1.1. 本文件是THE BEYOND服务（以下简称“提供商”）的正式提议（公开要约），并包含所有提供组织远程访问网络节点（代理/VPN）服务的基本条件。\n" +
			"1.2. 本要约的接受被视为在使用Telegram机器人开始时、按“Start”按钮或支付资费计划时发生。\n\n" +
			"<b>2. 要约主题</b>\n" +
			"2.1. 提供商授予客户非独占权（许可证）使用访问配置（VLESS/Hysteria 2/MTProto密钥）访问提供商的服务器。\n" +
			"2.2. 服务被视为在通过Telegram机器人自动模式向客户发送访问密钥或配置文件时提供。\n\n" +
			"<b>3. 各方的权利和义务</b>\n" +
			"3.1. 提供商承诺维护服务器的功能，但不保证100%可用性（Uptime），如果客户的互联网提供商或政府机构（不可抗力）进行封锁。\n" +
			"3.2. 客户承诺仅将服务用于合法目的。\n" +
			"3.3. 客户严格禁止：\n" +
			"• 使用服务发送垃圾邮件（电子邮件、社交网络）。\n" +
			"• 扫描端口、进行DDoS攻击、暴力破解和黑客攻击资源。\n" +
			"• 分发恶意软件、麻醉品、儿童色情或宣传恐怖主义的材料。\n" +
			"• 将访问密钥传输给第三方（除非资费规定）。\n\n" +
			"<b>4. 责任和封锁</b>\n" +
			"4.1. 如果违反第3.3条，提供商有权立即封锁客户的访问，而不退款。\n" +
			"4.2. 提供商不对客户因使用或无法使用服务而遭受的直接或间接损害负责。\n" +
			"4.3. 服务按“现状”（As Is）条件提供。提供商不保证服务与客户的特定设备或提供商兼容。",
	},
	language.JA日本語: {
		Terms: "<b>公開オファー：</b>\n\n" +
			"<b>1. 一般規定</b>\n" +
			"1.1. この文書は、THE BEYONDサービス（以下「提供者」）の公式提案（公開オファー）であり、网络ノード（プロキシ/VPN）へのリモートアクセスを組織するためのサービスの提供に関するすべての本質的な条件を含んでいます。\n" +
			"1.2. このオファーの受諾は、Telegramボットの使用開始、「Start」ボタンを押す、または料金プランの支払いにより行われたものとみなされます。\n\n" +
			"<b>2. オファーの対象</b>\n" +
			"2.1. 提供者は、クライアントに提供者のサーバーへのアクセス構成（VLESS/Hysteria 2/MTProtoキー）の使用に関する非排他的権利（ライセンス）を付与します。\n" +
			"2.2. サービスは、Telegramボットを通じて自動モードでクライアントにアクセスキーまたは構成ファイルを送信した時点で提供されたものとみなされます。\n\n" +
			"<b>3. 当事者の権利と義務</b>\n" +
			"3.1. 提供者はサーバーの機能性を維持することを約束しますが、クライアントのインターネットプロバイダーまたは政府機関（不可抗力）によるブロックの場合、100%の可用性（Uptime）を保証しません。\n" +
			"3.2. クライアントはサービスを合法的な目的でのみ使用することを約束します。\n" +
			"3.3. クライアントは厳格に禁止されています：\n" +
			"• サービスをスパム送信（メール、ソーシャルネットワーク）に使用する。\n" +
			"• ポートをスキャンし、DDoS攻撃を実行し、ブルートフォースおよびリソースのハッキングを行う。\n" +
			"• 悪意のあるソフトウェア、麻薬物質、児童ポルノ、またはテロリズムを宣伝する資料を配布する。\n" +
			"• アクセスキーを第三者に譲渡する（料金で規定されていない場合）。\n\n" +
			"<b>4. 責任とブロック</b>\n" +
			"4.1. 3.3条の違反の場合、提供者はクライアントのアクセスを即座にブロックする権利を有し、資金の返金なし。\n" +
			"4.2. 提供者は、サービス使用または使用不能に関連してクライアントが被った直接的または間接的な損害について責任を負いません。\n" +
			"4.3. サービスは「現状のまま」（As Is）の条件で提供されます。提供者は、サービスがクライアントの特定のデバイスまたはプロバイダーと互換性があることを保証しません。",
	},
	language.KO한국어: {
		Terms: "<b>공개 제안:</b>\n\n" +
			"<b>1. 일반 규정</b>\n" +
			"1.1. 이 문서는 THE BEYOND 서비스(이하 \"제공자\")의 공식 제안(공개 제안)이며, 네트워크 노드(프록시/VPN)에 대한 원격 액세스 조직을 위한 서비스 제공에 대한 모든 본질적인 조건을 포함합니다.\n" +
			"1.2. 이 제안의 수락은 Telegram 봇 사용 시작, \"Start\" 버튼 누르기 또는 요금제 지불로 간주됩니다.\n\n" +
			"<b>2. 제안 주제</b>\n" +
			"2.1. 제공자는 고객에게 제공자의 서버에 대한 액세스 구성(VLESS/Hysteria 2/MTProto 키) 사용에 대한 비독점적 권리(라이선스)를 부여합니다.\n" +
			"2.2. 서비스는 Telegram 봇을 통해 자동 모드로 고객에게 액세스 키 또는 구성 파일을 보낸 순간 제공된 것으로 간주됩니다.\n\n" +
			"<b>3. 당사자의 권리와 의무</b>\n" +
			"3.1. 제공자는 서버의 기능성을 유지할 것을 약속하지만, 고객의 인터넷 제공자 또는 정부 기관(불가항력)에 의한 차단 시 100% 가용성(Uptime)을 보장하지 않습니다.\n" +
			"3.2. 고객은 서비스를 합법적인 목적으로만 사용할 것을 약속합니다.\n" +
			"3.3. 고객은 엄격히 금지됩니다:\n" +
			"• 서비스를 스팸 전송(이메일, 소셜 네트워크)에 사용.\n" +
			"• 포트 스캔, DDoS 공격 수행, 브루트 포스 및 리소스 해킹.\n" +
			"• 악성 소프트웨어, 마약 물질, 아동 포르노 또는 테러리즘을 홍보하는 자료 배포.\n" +
			"• 액세스 키를 제3자에게 양도(요금제에 규정되지 않은 경우).\n\n" +
			"<b>4. 책임 및 차단</b>\n" +
			"4.1. 3.3항 위반 시, 제공자는 고객의 액세스를 즉시 차단할 권리가 있으며, 자금 환불 없음.\n" +
			"4.2. 제공자는 서비스 사용 또는 사용 불가능과 관련하여 고객이 입은 직접적 또는 간접적 손해에 대해 책임을 지지 않습니다.\n" +
			"4.3. 서비스는 \"있는 그대로\"(As Is) 조건으로 제공됩니다. 제공자는 서비스가 고객의 특정 장치 또는 제공자와 호환될 것을 보장하지 않습니다.",
	},
	language.TiếngViệt: {
		Terms: "<b>ĐỀ NGHỊ CÔNG KHAI:</b>\n\n" +
			"<b>1. Quy Định Chung</b>\n" +
			"1.1. Tài liệu này là đề nghị chính thức (đề nghị công khai) của dịch vụ THE BEYOND (sau đây gọi là \"Nhà Cung Cấp\") và chứa tất cả các điều kiện thiết yếu để cung cấp dịch vụ tổ chức truy cập từ xa vào các nút mạng (proxy/VPN).\n" +
			"1.2. Việc chấp nhận Đề Nghị này được coi là xảy ra khi bắt đầu sử dụng bot Telegram, nhấn nút \"Start\" hoặc thanh toán kế hoạch tariff.\n\n" +
			"<b>2. Chủ Đề Của Đề Nghị</b>\n" +
			"2.1. Nhà Cung Cấp cấp cho Khách Hàng quyền không độc quyền (giấy phép) sử dụng cấu hình truy cập (khóa VLESS/Hysteria 2/MTProto) vào máy chủ của Nhà Cung Cấp.\n" +
			"2.2. Dịch vụ được coi là đã cung cấp tại thời điểm gửi khóa truy cập hoặc tệp cấu hình cho Khách Hàng ở chế độ tự động qua bot Telegram.\n\n" +
			"<b>3. Quyền Và Nghĩa Vụ Của Các Bên</b>\n" +
			"3.1. Nhà Cung Cấp cam kết duy trì chức năng của máy chủ, tuy nhiên không đảm bảo 100% tính sẵn có (Uptime) trong trường hợp chặn từ nhà cung cấp internet của Khách Hàng hoặc cơ quan nhà nước (lực lượng lớn).\n" +
			"3.2. Khách Hàng cam kết sử dụng dịch vụ chỉ cho mục đích hợp pháp.\n" +
			"3.3. Khách Hàng bị cấm nghiêm ngặt:\n" +
			"• Sử dụng dịch vụ để gửi spam (email, mạng xã hội).\n" +
			"• Quét cổng, thực hiện tấn công DDoS, brute force và hack tài nguyên.\n" +
			"• Phân phối phần mềm độc hại, chất ma túy, khiêu dâm trẻ em hoặc tài liệu tuyên truyền khủng bố.\n" +
			"• Chuyển khóa truy cập cho bên thứ ba (trừ khi được quy định trong tariff).\n\n" +
			"<b>4. Trách Nhiệm Và Chặn</b>\n" +
			"4.1. Trong trường hợp vi phạm khoản 3.3, Nhà Cung Cấp có quyền chặn ngay lập tức truy cập của Khách Hàng mà không hoàn tiền.\n" +
			"4.2. Nhà Cung Cấp không chịu trách nhiệm cho thiệt hại, trực tiếp hoặc gián tiếp, mà Khách Hàng chịu phải liên quan đến việc sử dụng hoặc không thể sử dụng dịch vụ.\n" +
			"4.3. Dịch vụ được cung cấp theo điều kiện \"Như Là\" (As Is). Nhà Cung Cấp không đảm bảo rằng dịch vụ sẽ tương thích với thiết bị cụ thể hoặc nhà cung cấp của Khách Hàng.",
	},
	language.THภาษาไทย: {
		Terms: "<b>ข้อเสนอสาธารณะ:</b>\n\n" +
			"<b>1. ข้อกำหนดทั่วไป</b>\n" +
			"1.1. เอกสารนี้เป็นข้อเสนออย่างเป็นทางการ (ข้อเสนอสาธารณะ) ของบริการ THE BEYOND (ต่อไปนี้เรียกว่า \"ผู้ให้บริการ\") และมีเงื่อนไขสำคัญทั้งหมดสำหรับการให้บริการจัดระเบียบการเข้าถึงจากระยะไกลไปยังโหนดเครือข่าย (พร็อกซี/VPN).\n" +
			"1.2. การยอมรับข้อเสนอนี้นับว่ามีผลเมื่อเริ่มใช้งานบอท Telegram กดปุ่ม \"Start\" หรือชำระเงินสำหรับแผนอัตราค่าใช้จ่าย.\n\n" +
			"<b>2. หัวข้อข้อเสนอ</b>\n" +
			"2.1. ผู้ให้บริการมอบสิทธิ์ไม่ผูกขาด (ใบอนุญาต) ให้ลูกค้าในการใช้การกำหนดค่าการเข้าถึง (คีย์ VLESS/Hysteria 2/MTProto) ไปยังเซิร์ฟเวอร์ของผู้ให้บริการ.\n" +
			"2.2. บริการนับว่าถูกให้บริการในขณะที่ส่งคีย์การเข้าถึงหรือไฟล์กำหนดค่าให้ลูกค้าในโหมดอัตโนมัติผ่านบอท Telegram.\n\n" +
			"<b>3. สิทธิ์และหน้าที่ของคู่สัญญา</b>\n" +
			"3.1. ผู้ให้บริการผูกพันที่จะรักษาการทำงานของเซิร์ฟเวอร์ แต่ไม่รับประกันความพร้อมใช้งาน 100% (Uptime) ในกรณีที่ถูกบล็อกจากผู้ให้บริการอินเทอร์เน็ตของลูกค้าหรือหน่วยงานรัฐ (เหตุสุดวิสัย).\n" +
			"3.2. ลูกค้าผูกพันที่จะใช้บริการเฉพาะเพื่อวัตถุประสงค์ที่ถูกกฎหมาย.\n" +
			"3.3. ลูกค้าถูกห้ามอย่างเคร่งครัด:\n" +
			"• ใช้บริการเพื่อส่งสแปม (อีเมล, เครือข่ายสังคม).\n" +
			"• สแกนพอร์ต, ดำเนินการโจมตี DDoS, brute force และแฮ็กทรัพยากร.\n" +
			"• แจกจ่ายซอฟต์แวร์ที่เป็นอันตราย, สารเสพติด, ภาพอนาจารเด็กหรือวัสดุที่ส่งเสริมการก่อการร้าย.\n" +
			"• ถ่ายโอนคีย์การเข้าถึงให้บุคคลที่สาม (เว้นแต่กำหนดไว้ในอัตราค่าใช้จ่าย).\n\n" +
			"<b>4. ความรับผิดชอบและการบล็อก</b>\n" +
			"4.1. ในกรณีที่ละเมิดข้อ 3.3 ผู้ให้บริการมีสิทธิ์บล็อกการเข้าถึงของลูกค้าทันทีโดยไม่คืนเงิน.\n" +
			"4.2. ผู้ให้บริการไม่รับผิดชอบต่อความเสียหายโดยตรงหรือทางอ้อมที่ลูกค้าได้รับในส่วนที่เกี่ยวกับการใช้หรือไม่สามารถใช้บริการ.\n" +
			"4.3. บริการถูกให้บริการภายใต้เงื่อนไข \"ตามที่เป็น\" (As Is). ผู้ให้บริการไม่รับประกันว่าบริการจะเข้ากันได้กับอุปกรณ์เฉพาะหรือผู้ให้บริการของลูกค้า.",
	},
	language.BahasaIndonesia: {
		Terms: "<b>TAWARAN PUBLIK:</b>\n\n" +
			"<b>1. Ketentuan Umum</b>\n" +
			"1.1. Dokumen ini merupakan penawaran resmi (tawaran publik) dari layanan THE BEYOND (selanjutnya disebut \"Penyedia\") dan berisi semua ketentuan esensial untuk penyediaan layanan pengorganisasian akses jarak jauh ke node jaringan (proxy/VPN).\n" +
			"1.2. Penerimaan Tawaran ini dianggap terjadi pada saat mulai menggunakan bot Telegram, menekan tombol \"Start\" atau membayar rencana tarif.\n\n" +
			"<b>2. Subjek Tawaran</b>\n" +
			"2.1. Penyedia memberikan kepada Pelanggan hak non-eksklusif (lisensi) untuk menggunakan konfigurasi akses (kunci VLESS/Hysteria 2/MTProto) ke server Penyedia.\n" +
			"2.2. Layanan dianggap diberikan pada saat pengiriman kunci akses atau file konfigurasi ke Pelanggan dalam mode otomatis melalui bot Telegram.\n\n" +
			"<b>3. Hak dan Kewajiban Para Pihak</b>\n" +
			"3.1. Penyedia berkomitmen untuk menjaga fungsionalitas server, namun tidak menjamin 100% ketersediaan (Uptime) dalam kasus pemblokiran dari penyedia internet Pelanggan atau organ negara (force majeure).\n" +
			"3.2. Pelanggan berkomitmen untuk menggunakan layanan hanya untuk tujuan yang sah.\n" +
			"3.3. Pelanggan dilarang keras:\n" +
			"• Menggunakan layanan untuk mengirim spam (email, jaringan sosial).\n" +
			"• Memindai port, melakukan serangan DDoS, brute force dan hacking sumber daya.\n" +
			"• Mendistribusikan perangkat lunak berbahaya, zat narkotika, pornografi anak atau bahan yang mempromosikan terorisme.\n" +
			"• Memindahkan kunci akses ke pihak ketiga (kecuali jika ditentukan dalam tarif).\n\n" +
			"<b>4. Tanggung Jawab dan Pemblokiran</b>\n" +
			"4.1. Dalam kasus pelanggaran poin 3.3, Penyedia berhak untuk segera memblokir akses Pelanggan tanpa pengembalian dana.\n" +
			"4.2. Penyedia tidak bertanggung jawab atas kerusakan, langsung atau tidak langsung, yang diderita Pelanggan sehubungan dengan penggunaan atau ketidakmampuan menggunakan layanan.\n" +
			"4.3. Layanan disediakan dengan ketentuan \"Sebagaimana Adanya\" (As Is). Penyedia tidak menjamin bahwa layanan akan kompatibel dengan perangkat spesifik atau penyedia Pelanggan.",
	},
	language.BahasaMelayu: {
		Terms: "<b>TAWARAN AWAM:</b>\n\n" +
			"<b>1. Peruntukan Am</b>\n" +
			"1.1. Dokumen ini adalah cadangan rasmi (tawaran awam) perkhidmatan THE BEYOND (selepas ini \"Penyedia\") dan mengandungi semua syarat penting untuk penyediaan perkhidmatan pengorganisasian akses jauh ke nod rangkaian (proxy/VPN).\n" +
			"1.2. Penerimaan Tawaran ini dianggap berlaku apabila mula menggunakan bot Telegram, menekan butang \"Start\" atau membayar pelan tarif.\n\n" +
			"<b>2. Subjek Tawaran</b>\n" +
			"2.1. Penyedia memberikan kepada Pelanggan hak tidak eksklusif (lesen) untuk menggunakan konfigurasi akses (kunci VLESS/Hysteria 2/MTProto) ke pelayan Penyedia.\n" +
			"2.2. Perkhidmatan dianggap disediakan pada masa penghantaran kunci akses atau fail konfigurasi kepada Pelanggan dalam mod automatik melalui bot Telegram.\n\n" +
			"<b>3. Hak dan Kewajipan Pihak</b>\n" +
			"3.1. Penyedia komited untuk mengekalkan fungsi pelayan, namun tidak menjamin 100% ketersediaan (Uptime) dalam kes sekatan dari penyedia internet Pelanggan atau organ negeri (force majeure).\n" +
			"3.2. Pelanggan komited untuk menggunakan perkhidmatan hanya untuk tujuan sah.\n" +
			"3.3. Pelanggan dilarang keras:\n" +
			"• Menggunakan perkhidmatan untuk menghantar spam (emel, rangkaian sosial).\n" +
			"• Mengimbas port, melakukan serangan DDoS, brute force dan hack sumber.\n" +
			"• Mengedar perisian berbahaya, bahan narkotik, pornografi kanak-kanak atau bahan yang mempromosikan keganasan.\n" +
			"• Memindahkan kunci akses kepada pihak ketiga (melainkan ditetapkan dalam tarif).\n\n" +
			"<b>4. Tanggungjawab dan Sekatan</b>\n" +
			"4.1. Dalam kes pelanggaran perkara 3.3, Penyedia berhak untuk segera menyekat akses Pelanggan tanpa pengembalian wang.\n" +
			"4.2. Penyedia tidak bertanggungjawab atas kerosakan, langsung atau tidak langsung, yang dialami Pelanggan berkaitan dengan penggunaan atau ketidakupayaan menggunakan perkhidmatan.\n" +
			"4.3. Perkhidmatan disediakan dengan syarat \"Seperti Adanya\" (As Is). Penyedia tidak menjamin bahawa perkhidmatan akan serasi dengan peranti khusus atau penyedia Pelanggan.",
	},
	language.Tagalog: {
		Terms: "<b>PUBLIC OFFER:</b>\n\n" +
			"<b>1. Pangkalahatang Mga Probisyon</b>\n" +
			"1.1. Ang dokumentong ito ay opisyal na panukala (pampublikong offer) ng serbisyo na THE BEYOND (mula rito \"Provider\") at naglalaman ng lahat ng mahahalagang kondisyon para sa pagbibigay ng mga serbisyo para sa pag-oorganisa ng malayong access sa mga node ng network (proxy/VPN).\n" +
			"1.2. Ang pagtanggap sa Offer na ito ay itinuturing na nangyari sa pagsisimula ng paggamit ng Telegram bot, pagpindot sa button na \"Start\" o pagbabayad para sa tariff plan.\n\n" +
			"<b>2. Paksa ng Offer</b>\n" +
			"2.1. Ang Provider ay nagbibigay sa Customer ng hindi eksklusibong karapatan (lisensya) na gamitin ang mga configuration ng access (mga key na VLESS/Hysteria 2/MTProto) sa mga server ng Provider.\n" +
			"2.2. Ang serbisyo ay itinuturing na ibinigay sa sandaling ipadala ang access key o configuration file sa Customer sa automatic mode sa pamamagitan ng Telegram bot.\n\n" +
			"<b>3. Mga Karapatan at Obligasyon ng mga Partido</b>\n" +
			"3.1. Ang Provider ay nangangako na panatilihin ang functionality ng mga server, gayunpaman hindi ginagarantiya ang 100% availability (Uptime) sa kaso ng mga block mula sa internet providers ng Customer o mga ahensya ng gobyerno (force majeure).\n" +
			"3.2. Ang Customer ay nangangako na gamitin ang mga serbisyo lamang para sa legal na layunin.\n" +
			"3.3. Ang Customer ay mahigpit na ipinagbabawal:\n" +
			"• Gamitin ang serbisyo para sa pagpapadala ng spam (email, social networks).\n" +
			"• I-scan ang mga port, magsagawa ng DDoS attacks, brute force at hacking ng resources.\n" +
			"• Ipamahagi ang malicious software, narcotic substances, child pornography o materials na nagpo-promote ng terorismo.\n" +
			"• I-transfer ang access keys sa third parties (maliban kung ibinigay sa tariff).\n\n" +
			"<b>4. Pananagutan at Pag-block</b>\n" +
			"4.1. Sa kaso ng paglabag sa punto 3.3, ang Provider ay may karapatan na agad na i-block ang access ng Customer nang walang refund ng pera.\n" +
			"4.2. Ang Provider ay hindi responsable para sa pinsala, direkta o hindi direkta, na naranasan ng Customer kaugnay ng paggamit o kawalan ng kakayahang gamitin ang serbisyo.\n" +
			"4.3. Ang serbisyo ay ibinibigay sa kondisyon na \"As Is\". Ang Provider ay hindi ginagarantiya na ang serbisyo ay magiging compatible sa tiyak na device o provider ng Customer.",
	},
	language.Hindi: {
		Terms: "<b>सार्वजनिक प्रस्ताव:</b>\n\n" +
			"<b>1. सामान्य प्रावधान</b>\n" +
			"1.1. यह दस्तावेज़ THE BEYOND सेवा (इसके बाद \"प्रदाता\") का आधिकारिक प्रस्ताव (सार्वजनिक प्रस्ताव) है और नेटवर्क नोड्स (प्रॉक्सी/VPN) तक रिमोट एक्सेस के संगठन के लिए सेवाओं के प्रावधान के सभी आवश्यक शर्तें शामिल करता है.\n" +
			"1.2. इस प्रस्ताव की स्वीकृति टेलीग्राम बॉट के उपयोग की शुरुआत, \"Start\" बटन दबाने या टैरिफ योजना के भुगतान पर मानी जाती है.\n\n" +
			"<b>2. प्रस्ताव का विषय</b>\n" +
			"2.1. प्रदाता ग्राहक को प्रदाता के सर्वरों तक पहुंच कॉन्फ़िगरेशन (VLESS/Hysteria 2/MTProto कुंजियाँ) का उपयोग करने का गैर-अनन्य अधिकार (लाइसेंस) प्रदान करता है.\n" +
			"2.2. सेवा उस क्षण प्रदान की गई मानी जाती है जब पहुंच कुंजी या कॉन्फ़िगरेशन फ़ाइल टेलीग्राम बॉट के माध्यम से स्वचालित मोड में ग्राहक को भेजी जाती है.\n\n" +
			"<b>3. पक्षों के अधिकार और दायित्व</b>\n" +
			"3.1. प्रदाता सर्वरों की कार्यक्षमता बनाए रखने का वचन देता है, हालांकि ग्राहक के इंटरनेट प्रदाताओं या सरकारी अंगों (फोर्स मेज्योर) द्वारा ब्लॉकिंग के मामले में 100% उपलब्धता (Uptime) की गारंटी नहीं देता.\n" +
			"3.2. ग्राहक सेवाओं का उपयोग केवल वैध उद्देश्यों के लिए करने का वचन देता है.\n" +
			"3.3. ग्राहक को सख्ती से निषिद्ध है:\n" +
			"• सेवा का उपयोग स्पैम भेजने के लिए (ईमेल, सोशल नेटवर्क).\n" +
			"• पोर्ट स्कैन करना, DDoS हमले करना, ब्रूट फोर्स और संसाधनों का हैकिंग.\n" +
			"• मैलवेयर, नारकोटिक पदार्थ, बाल पोर्नोग्राफी या आतंकवाद को बढ़ावा देने वाली सामग्री वितरित करना.\n" +
			"• पहुंच कुंजियाँ तीसरे पक्षों को हस्तांतरित करना (जब तक कि टैरिफ में प्रदान न किया गया हो).\n\n" +
			"<b>4. जिम्मेदारी और ब्लॉकिंग</b>\n" +
			"4.1. बिंदु 3.3 के उल्लंघन के मामले में, प्रदाता को ग्राहक की पहुंच को तुरंत ब्लॉक करने का अधिकार है बिना धन वापसी के.\n" +
			"4.2. प्रदाता सेवा के उपयोग या उपयोग करने की असमर्थता से संबंधित ग्राहक द्वारा भुगती गई प्रत्यक्ष या अप्रत्यक्ष क्षति के लिए जिम्मेदार नहीं है.\n" +
			"4.3. सेवा \"जैसी है\" (As Is) शर्तों पर प्रदान की जाती है. प्रदाता गारंटी नहीं देता कि सेवा ग्राहक के विशिष्ट डिवाइस या प्रदाता के साथ संगत होगी.",
	},
	language.URاردو: {
		Terms: "<b>عوامی پیشکش:</b>\n\n" +
			"<b>1. عام احکام</b>\n" +
			"1.1. یہ دستاویز THE BEYOND سروس (آگے \"فراہم کنندہ\") کی رسمی پیشکش (عوامی پیشکش) ہے اور نیٹ ورک نوڈز (پراکسی/VPN) تک ریموٹ رسائی کی تنظیم کے لیے خدمات کی فراہمی کے تمام ضروری شرائط شامل کرتی ہے.\n" +
			"1.2. اس پیشکش کی قبولیت ٹیلی گرام بوٹ کے استعمال کی شروعات، \"Start\" بٹن دبانے یا ٹیرف پلان کی ادائیگی پر مانی جاتی ہے.\n\n" +
			"<b>2. پیشکش کا موضوع</b>\n" +
			"2.1. فراہم کنندہ صارف کو فراہم کنندہ کے سرورز تک رسائی کی ترتیب (VLESS/Hysteria 2/MTProto کیز) استعمال کرنے کا غیر خصوصی حق (لائسنس) دیتا ہے.\n" +
			"2.2. سروس اس لمحے فراہم کی گئی مانی جاتی ہے جب رسائی کی کلید یا ترتیب فائل ٹیلی گرام بوٹ کے ذریعے خودکار موڈ میں صارف کو بھیجی جاتی ہے.\n\n" +
			"<b>3. فریقین کے حقوق اور ذمہ داریاں</b>\n" +
			"3.1. فراہم کنندہ سرورز کی فعالیت برقرار رکھنے کا وعدہ کرتا ہے، تاہم صارف کے انٹرنیٹ فراہم کنندگان یا سرکاری اداروں (فورس میجر) کی جانب سے بلاک کرنے کی صورت میں 100% دستیابیت (Uptime) کی ضمانت نہیں دیتا.\n" +
			"3.2. صارف خدمات کو صرف قانونی مقاصد کے لیے استعمال کرنے کا وعدہ کرتا ہے.\n" +
			"3.3. صارف کو سختی سے منع ہے:\n" +
			"• سروس کو اسپیم بھیجنے کے لیے استعمال کرنا (ای میل، سوشل نیٹ ورکس).\n" +
			"• پورٹس سکین کرنا، DDoS حملے کرنا، بروٹ فورس اور وسائل کی ہیکنگ.\n" +
			"• میلویئر، منشیات کے مواد، بچوں کی پورنوگرافی یا دہشت گردی کو فروغ دینے والے مواد تقسیم کرنا.\n" +
			"• رسائی کیز تیسرے فریقوں کو منتقل کرنا (جب تک کہ ٹیرف میں فراہم نہ کیا گیا ہو).\n\n" +
			"<b>4. ذمہ داری اور بلاکنگ</b>\n" +
			"4.1. پوائنٹ 3.3 کی خلاف ورزی کی صورت میں، فراہم کنندہ کو صارف کی رسائی کو فوری طور پر بلاک کرنے کا حق ہے بغیر پیسے کی واپسی کے.\n" +
			"4.2. فراہم کنندہ سروس کے استعمال یا استعمال نہ کرنے سے متعلق صارف کو پہنچنے والے براہ راست یا بالواسطہ نقصانات کے لیے ذمہ دار نہیں ہے.\n" +
			"4.3. سروس \"جیسا ہے\" (As Is) شرائط پر فراہم کی جاتی ہے. فراہم کنندہ ضمانت نہیں دیتا کہ سروس صارف کے مخصوص ڈیوائس یا فراہم کنندہ کے ساتھ مطابقت رکھے گی.",
	},
	language.Bengali: {
		Terms: "<b>পাবলিক অফার:</b>\n\n" +
			"<b>1. সাধারণ প্রভিশন</b>\n" +
			"1.1. এই ডকুমেন্টটি THE BEYOND সার্ভিসের (পরবর্তীতে \"প্রোভাইডার\") অফিসিয়াল প্রস্তাব (পাবলিক অফার) এবং নেটওয়ার্ক নোডস (প্রক্সি/VPN) এর রিমোট অ্যাক্সেসের সংগঠনের জন্য সার্ভিস প্রদানের সমস্ত অপরিহার্য শর্তাবলী অন্তর্ভুক্ত করে.\n" +
			"1.2. এই অফারের গ্রহণ টেলিগ্রাম বটের ব্যবহার শুরু, \"Start\" বোতাম চাপা বা ট্যারিফ প্ল্যানের পেমেন্টে মনে করা হয়.\n\n" +
			"<b>2. অফারের বিষয়</b>\n" +
			"2.1. প্রোভাইডার ক্লায়েন্টকে প্রোভাইডারের সার্ভারগুলিতে অ্যাক্সেস কনফিগারেশন (VLESS/Hysteria 2/MTProto কী) ব্যবহারের অ-এক্সক্লুসিভ অধিকার (লাইসেন্স) প্রদান করে.\n" +
			"2.2. সার্ভিসটি সেই মুহূর্তে প্রদান করা মনে করা হয় যখন অ্যাক্সেস কী বা কনফিগারেশন ফাইল টেলিগ্রাম বটের মাধ্যমে অটোম্যাটিক মোডে ক্লায়েন্টকে পাঠানো হয়.\n\n" +
			"<b>3. পক্ষগুলির অধিকার এবং দায়িত্ব</b>\n" +
			"3.1. প্রোভাইডার সার্ভারগুলির কার্যকারিতা বজায় রাখার প্রতিশ্রুতি দেয়, তবে ক্লায়েন্টের ইন্টারনেট প্রোভাইডার বা সরকারি অঙ্গগুলির (ফোর্স মেজর) দ্বারা ব্লকিংয়ের ক্ষেত্রে 100% উপলব্ধতা (Uptime) গ্যারান্টি দেয় না.\n" +
			"3.2. ক্লায়েন্ট সার্ভিসগুলি শুধুমাত্র আইনি উদ্দেশ্যে ব্যবহার করার প্রতিশ্রুতি দেয়.\n" +
			"3.3. ক্লায়েন্টকে কঠোরভাবে নিষিদ্ধ:\n" +
			"• সার্ভিসটি স্প্যাম পাঠানোর জন্য ব্যবহার করা (ইমেল, সোশ্যাল নেটওয়ার্কস).\n" +
			"• পোর্ট স্ক্যান করা, DDoS আক্রমণ করা, ব্রুট ফোর্স এবং রিসোর্স হ্যাকিং.\n" +
			"• ম্যালওয়্যার, নারকোটিক পদার্থ, শিশু পর্নোগ্রাফি বা টেররিজম প্রচারকারী উপাদান বিতরণ করা.\n" +
			"• অ্যাক্সেস কী তৃতীয় পক্ষগুলিকে হস্তান্তর করা (যদি ট্যারিফে প্রদান না করা হয়).\n\n" +
			"<b>4. দায়িত্ব এবং ব্লকিং</b>\n" +
			"4.1. পয়েন্ট 3.3 লঙ্ঘনের ক্ষেত্রে, প্রোভাইডার ক্লায়েন্টের অ্যাক্সেসকে তাত্ক্ষণিকভাবে ব্লক করার অধিকার রাখে টাকা ফেরত ছাড়া.\n" +
			"4.2. প্রোভাইডার সার্ভিসের ব্যবহার বা ব্যবহার করতে না পারার সাথে সম্পর্কিত ক্লায়েন্টের সরাসরি বা অপ্রত্যক্ষ ক্ষতির জন্য দায়ী নয়.\n" +
			"4.3. সার্ভিস \"যেমন আছে\" (As Is) শর্তাবলীতে প্রদান করা হয়. প্রোভাইডার গ্যারান্টি দেয় না যে সার্ভিস ক্লায়েন্টের নির্দিষ্ট ডিভাইস বা প্রোভাইডারের সাথে সামঞ্জস্যপূর্ণ হবে.",
	},
	language.Tamiḻ: {
		Terms: "<b>பொது வழங்கல்:</b>\n\n" +
			"<b>1. பொது விதிகள்</b>\n" +
			"1.1. இந்த ஆவணம் THE BEYOND சேவையின் (இனி \"வழங்குபவர்\") அதிகாரப்பூர்வ வழங்கல் (பொது வழங்கல்) ஆகும் மற்றும் நெட்வொர்க் நோட்கள் (ப்ராக்ஸி/VPN) க்கான தொலைநிலை அணுகல் அமைப்புக்கான சேவைகளை வழங்குவதற்கான அனைத்து அத்தியாவசிய நிபந்தனைகளை உள்ளடக்கியது.\n" +
			"1.2. இந்த வழங்கலின் ஏற்றுக்கொள்ளல் டெலிகிராம் பாட்டின் பயன்பாடு தொடங்கியபோது, \"Start\" பொத்தானை அழுத்தியபோது அல்லது டாரிஃப் திட்டத்திற்கு பணம் செலுத்தியபோது ஏற்றுக்கொள்ளப்பட்டதாக கருதப்படும்.\n\n" +
			"<b>2. வழங்கலின் பொருள்</b>\n" +
			"2.1. வழங்குபவர் வாடிக்கையாளருக்கு வழங்குபவரின் சர்வர்களுக்கு அணுகல் உள்ளமைப்புகளை (VLESS/Hysteria 2/MTProto சாவிகள்) பயன்படுத்துவதற்கான பிரத்தியேகமற்ற உரிமை (உரிமம்) வழங்குகிறார்.\n" +
			"2.2. சேவை அந்த தருணத்தில் வழங்கப்பட்டதாக கருதப்படும் போது அணுகல் சாவி அல்லது உள்ளமைப்பு கோப்பு டெலிகிராம் பாட் மூலம் தானியங்கி முறையில் வாடிக்கையாளருக்கு அனுப்பப்படும்.\n\n" +
			"<b>3. பக்கங்களின் உரிமைகள் மற்றும் கடமைகள்</b>\n" +
			"3.1. வழங்குபவர் சர்வர்களின் செயல்பாட்டை பராமரிக்க உறுதியளிக்கிறார், இருப்பினும் வாடிக்கையாளரின் இன்டர்நெட் வழங்குநர்கள் அல்லது அரசு அமைப்புகள் (போர்ஸ் மேஜர்) மூலம் தடுப்பின் வழக்கில் 100% கிடைக்கும் (Uptime) உத்தரவாதம் செய்யவில்லை.\n" +
			"3.2. வாடிக்கையாளர் சேவைகளை சட்டபூர்வமான நோக்கங்களுக்கு மட்டுமே பயன்படுத்த உறுதியளிக்கிறார்.\n" +
			"3.3. வாடிக்கையாளருக்கு கடுமையாக தடைசெய்யப்பட்டது:\n" +
			"• சேவையை ஸ்பேம் அனுப்புவதற்கு பயன்படுத்துதல் (மின்னஞ்சல், சமூக வலையமைப்புகள்).\n" +
			"• போர்ட்களை ஸ்கேன் செய்தல், DDoS தாக்குதல்களை நடத்துதல், ப்ரூட் ஃபோர்ஸ் மற்றும் வளங்களை ஹேக் செய்தல்.\n" +
			"• தீங்கிழைக்கும் மென்பொருள், நார்கோடிக் பொருட்கள், குழந்தை ஆபாசம் அல்லது பயங்கரவாதத்தை ஊக்குவிக்கும் பொருட்களை விநியோகித்தல்.\n" +
			"• அணுகல் சாவிகளை மூன்றாம் தரப்புகளுக்கு மாற்றுதல் (டாரிஃபில் வழங்கப்படாவிட்டால்).\n\n" +
			"<b>4. பொறுப்பு மற்றும் தடுப்பு</b>\n" +
			"4.1. புள்ளி 3.3 இன் மீறலின் வழக்கில், வழங்குபவருக்கு வாடிக்கையாளரின் அணுகலை உடனடியாக தடுப்பதற்கான உரிமை உள்ளது பணம் திருப்பி அளிக்காமல்.\n" +
			"4.2. வழங்குபவர் சேவையின் பயன்பாடு அல்லது பயன்படுத்த முடியாததால் வாடிக்கையாளரால் அனுபவிக்கப்பட்ட நேரடி அல்லது மறைமுக சேதங்களுக்கு பொறுப்பேற்கவில்லை.\n" +
			"4.3. சேவை \"எப்படி இருக்கிறது\" (As Is) நிபந்தனைகளில் வழங்கப்படுகிறது. வழங்குபவர் சேவை வாடிக்கையாளரின் குறிப்பிட்ட சாதனம் அல்லது வழங்குபவருடன் இணக்கமாக இருக்கும் என்று உத்தரவாதம் செய்யவில்லை.",
	},
	language.Telugu: {
		Terms: "<b>పబ్లిక్ ఆఫర్:</b>\n\n" +
			"<b>1. సాధారణ నిబంధనలు</b>\n" +
			"1.1. ఈ డాక్యుమెంట్ THE BEYOND సర్వీస్ (ఇకముందు \"ప్రొవైడర్\") యొక్క అధికారిక ప్రతిపాదన (పబ్లిక్ ఆఫర్) మరియు నెట్‌వర్క్ నోడ్‌లకు (ప్రాక్సీ/VPN) రిమోట్ యాక్సెస్ ఆర్గనైజేషన్ కోసం సర్వీస్‌లు అందించడానికి అన్ని అవసరమైన షరతులు కలిగి ఉంది.\n" +
			"1.2. ఈ ఆఫర్ యొక్క స్వీకరణ టెలిగ్రామ్ బాట్ ఉపయోగం ప్రారంభం, \"Start\" బటన్ నొక్కడం లేదా టారిఫ్ ప్లాన్ చెల్లించడం వద్ద జరిగినట్లు భావించబడుతుంది.\n\n" +
			"<b>2. ఆఫర్ యొక్క సబ్జెక్ట్</b>\n" +
			"2.1. ప్రొవైడర్ కస్టమర్‌కు ప్రొవైడర్ సర్వర్‌లకు యాక్సెస్ కాన్ఫిగరేషన్‌లు (VLESS/Hysteria 2/MTProto కీలు) ఉపయోగించడానికి నాన్-ఎక్స్‌క్లూసివ్ రైట్ (లైసెన్స్) అందిస్తుంది.\n" +
			"2.2. సర్వీస్ ఆ మొమెంట్ అందించబడినట్లు భావించబడుతుంది యాక్సెస్ కీ లేదా కాన్ఫిగరేషన్ ఫైల్ టెలిగ్రామ్ బాట్ ద్వారా ఆటోమేటిక్ మోడ్‌లో కస్టమర్‌కు పంపబడుతుంది.\n\n" +
			"<b>3. పార్టీల హక్కులు మరియు బాధ్యతలు</b>\n" +
			"3.1. ప్రొవైడర్ సర్వర్‌ల ఫంక్షనాలిటీని నిర్వహించడానికి కట్టుబడి ఉంది, అయితే కస్టమర్ ఇంటర్నెట్ ప్రొవైడర్‌లు లేదా ప్రభుత్వ అంగాలు (ఫోర్స్ మేజర్) ద్వారా బ్లాకింగ్ కేసులో 100% అందుబాటులో (Uptime) గ్యారెంటీ ఇవ్వదు.\n" +
			"3.2. కస్టమర్ సర్వీస్‌లను లీగల్ పర్పస్‌లకు మాత్రమే ఉపయోగించడానికి కట్టుబడి ఉంది.\n" +
			"3.3. కస్టమర్‌కు కఠినంగా నిషేధించబడింది:\n" +
			"• సర్వీస్‌ను స్పామ్ పంపడానికి ఉపయోగించడం (ఈమెయిల్, సోషల్ నెట్‌వర్క్‌లు).\n" +
			"• పోర్ట్‌లను స్కాన్ చేయడం, DDoS అటాక్‌లు చేయడం, బ్రూట్ ఫోర్స్ మరియు రిసోర్స్‌ల హ్యాకింగ్.\n" +
			"• మాల్వేర్, నార్కోటిక్ సబ్‌స్టాన్స్‌లు, చైల్డ్ పోర్నోగ్రఫీ లేదా టెర్రరిజం ప్రమోట్ చేసే మెటీరియల్‌లను పంపిణీ చేయడం.\n" +
			"• యాక్సెస్ కీలను మూడవ పార్టీలకు ట్రాన్స్‌ఫర్ చేయడం (టారిఫ్‌లో అందించబడకపోతే).\n\n" +
			"<b>4. బాధ్యత మరియు బ్లాకింగ్</b>\n" +
			"4.1. పాయింట్ 3.3 ఉల్లంఘన కేసులో, ప్రొవైడర్ కస్టమర్ యాక్సెస్‌ను తక్షణమే బ్లాక్ చేయడానికి హక్కు కలిగి ఉంది డబ్బు రిఫండ్ లేకుండా.\n" +
			"4.2. ప్రొవైడర్ సర్వీస్ ఉపయోగం లేదా ఉపయోగించలేకపోవడంతో సంబంధం ఉన్న కస్టమర్ డైరెక్ట్ లేదా ఇండైరెక్ట్ డ్యామేజ్‌లకు బాధ్యత వహించదు.\n" +
			"4.3. సర్వీస్ \"ఎలా ఉంది\" (As Is) షరతులపై అందించబడుతుంది. ప్రొవైడర్ సర్వీస్ కస్టమర్ స్పెసిఫిక్ డివైస్ లేదా ప్రొవైడర్‌తో కంపాటిబుల్ అని గ్యారెంటీ ఇవ్వదు.",
	},
	language.Marathi: {
		Terms: "<b>सार्वजनिक ऑफर:</b>\n\n" +
			"<b>1. सामान्य तरतुदी</b>\n" +
			"1.1. हे दस्तऐवज THE BEYOND सेवेचे (पुढे \"पुरवठादार\") अधिकृत प्रस्ताव (सार्वजनिक ऑफर) आहे आणि नेटवर्क नोड्स (प्रॉक्सी/VPN) साठी रिमोट अॅक्सेसचे संघटन करण्यासाठी सेवांच्या पुरवठासाठी सर्व आवश्यक अटी समाविष्ट करते.\n" +
			"1.2. या ऑफरची स्वीकृती टेलिग्राम बॉटचा वापर सुरू करणे, \"Start\" बटण दाबणे किंवा टॅरिफ प्लॅनची देयक देणे यावर मानली जाते.\n\n" +
			"<b>2. ऑफरचा विषय</b>\n" +
			"2.1. पुरवठादार ग्राहकाला पुरवठादारच्या सर्व्हरसाठी अॅक्सेस कॉन्फिगरेशन (VLESS/Hysteria 2/MTProto की) वापरण्यासाठी गैर-अनन्य अधिकार (परवाना) देतो.\n" +
			"2.2. सेवा त्या क्षणी पुरवली गेली मानली जाते जेव्हा अॅक्सेस की किंवा कॉन्फिगरेशन फाइल टेलिग्राम बॉटद्वारे ऑटोमॅटिक मोडमध्ये ग्राहकाला पाठवली जाते.\n\n" +
			"<b>3. पक्षांच्या हक्क आणि जबाबदार्या</b>\n" +
			"3.1. पुरवठादार सर्व्हरच्या कार्यक्षमतेचे रक्षण करण्याचे वचन देतो, तथापि ग्राहकाच्या इंटरनेट पुरवठादार किंवा सरकारी संस्था (फोर्स मेजर) द्वारे ब्लॉकिंगच्या बाबतीत 100% उपलब्धता (Uptime) ची हमी देत नाही.\n" +
			"3.2. ग्राहक सेवा फक्त कायदेशीर उद्देशांसाठी वापरण्याचे वचन देतो.\n" +
			"3.3. ग्राहकाला कठोरपणे निषिद्ध आहे:\n" +
			"• सेवा स्पॅम पाठवण्यासाठी वापरणे (ईमेल, सोशल नेटवर्क).\n" +
			"• पोर्ट स्कॅन करणे, DDoS हल्ले करणे, ब्रूट फोर्स आणि संसाधनांची हॅकिंग.\n" +
			"• मालवेयर, नार्कोटिक पदार्थ, बाल पोर्नोग्राफी किंवा दहशतवादाला प्रोत्साहन देणारे सामग्री वितरण करणे.\n" +
			"• अॅक्सेस की तिसऱ्या पक्षांना हस्तांतरित करणे (टॅरिफमध्ये प्रदान केलेले नसल्यास).\n\n" +
			"<b>4. जबाबदारी आणि ब्लॉकिंग</b>\n" +
			"4.1. पॉइंट 3.3 च्या उल्लंघनाच्या बाबतीत, पुरवठादाराला ग्राहकाच्या अॅक्सेसला तात्काळ ब्लॉक करण्याचा अधिकार आहे पैशाची परतावा न करता.\n" +
			"4.2. पुरवठादार सेवा वापरणे किंवा वापरू शकत नसल्याशी संबंधित ग्राहकाने भोगलेल्या प्रत्यक्ष किंवा अप्रत्यक्ष नुकसानांसाठी जबाबदार नाही.\n" +
			"4.3. सेवा \"जशी आहे\" (As Is) अटींवर प्रदान केली जाते. पुरवठादार हमी देत नाही की सेवा ग्राहकाच्या विशिष्ट डिव्हाइस किंवा पुरवठादाराशी सुसंगत असेल.",
	},
}

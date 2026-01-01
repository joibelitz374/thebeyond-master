package i18n

import "github.com/quickpowered/thebeyond-master/configs/language"

type InvoiceLocale struct {
	Invoice string
}

var InvoiceMessages = map[language.Language]InvoiceLocale{
	language.English: {
		Invoice: "By paying, you agree to the service rules, including the refund policy. Access is provided immediately after successful payment. Thank you for choosing!",
	},
	language.Deutsch: {
		Invoice: "Durch die Zahlung stimmen Sie den Service-Regeln zu, einschließlich der Rückerstattungsrichtlinie. Der Zugriff wird sofort nach erfolgreicher Zahlung gewährt. Vielen Dank für Ihre Wahl!",
	},
	language.Nederlands: {
		Invoice: "Door te betalen, gaat u akkoord met de regels van de dienst, inclusief het restitutiebeleid. Toegang wordt onmiddellijk verleend na succesvolle betaling. Bedankt voor uw keuze!",
	},
	language.Svenska: {
		Invoice: "Genom att betala godkänner du tjänstens regler, inklusive återbetalningspolicy. Åtkomst tillhandahålls omedelbart efter framgångsrik betalning. Tack för ditt val!",
	},
	language.Norsk: {
		Invoice: "Ved å betale godtar du tjenestens regler, inkludert refusjonspolicy. Tilgang gis umiddelbart etter vellykket betaling. Takk for valget ditt!",
	},
	language.Dansk: {
		Invoice: "Ved at betale accepterer du tjenestens regler, herunder refusionspolitikken. Adgang gives straks efter vellykket betaling. Tak for dit valg!",
	},
	language.Español: {
		Invoice: "Al pagar, acepta las reglas del servicio, incluida la política de reembolsos. El acceso se proporciona inmediatamente después del pago exitoso. ¡Gracias por elegirnos!",
	},
	language.Français: {
		Invoice: "En payant, vous acceptez les règles du service, y compris la politique de remboursements. L'accès est fourni immédiatement après un paiement réussi. Merci pour votre choix !",
	},
	language.Português: {
		Invoice: "Ao pagar, você concorda com as regras do serviço, incluindo a política de reembolsos. O acesso é fornecido imediatamente após o pagamento bem-sucedido. Obrigado pela escolha!",
	},
	language.Italiana: {
		Invoice: "Pagando, accetti le regole del servizio, inclusa la politica di rimborsi. L'accesso viene fornito immediatamente dopo il pagamento riuscito. Grazie per la scelta!",
	},
	language.Русский: {
		Invoice: "Оплачивая, вы соглашаетесь с правилами сервиса, включая политику возвратов. Доступ предоставляется сразу после успешной оплаты. Спасибо за выбор!",
	},
	language.Українська: {
		Invoice: "Оплачуючи, ви погоджуєтеся з правилами сервісу, включаючи політику повернень. Доступ надається відразу після успішної оплати. Дякуємо за вибір!",
	},
	language.Polski: {
		Invoice: "Płacąc, zgadzasz się z regulaminem serwisu, w tym polityką zwrotów. Dostęp jest udzielany natychmiast po udanej płatności. Dziękujemy za wybór!",
	},
	language.Ceština: {
		Invoice: "Platbou souhlasíte s pravidly služby, včetně zásad vrácení peněz. Přístup je poskytnut ihned po úspěšné platbě. Děkujeme za výběr!",
	},
	language.Български: {
		Invoice: "С плащането се съгласявате с правилата на услугата, включително политиката за възстановяване. Достъпът се предоставя веднага след успешно плащане. Благодаря за избора!",
	},
	language.Српски: {
		Invoice: "Плаћањем се слажете са правилима услуге, укључујући политику поврата. Приступ се пружа одмах након успешног плаћања. Хвала на избору!",
	},
	language.Hrvatski: {
		Invoice: "Plaćanjem se slažete s pravilima usluge, uključujući politiku povrata. Pristup se pruža odmah nakon uspješnog plaćanja. Hvala na izboru!",
	},
	language.Slovenčina: {
		Invoice: "Platbou súhlasíte s pravidlami služby, vrátane politiky vrátenia peňazí. Prístup sa poskytuje ihneď po úspešnej platbe. Ďakujeme za výber!",
	},
	language.Slovenski: {
		Invoice: "S plačilom se strinjate s pravili storitve, vključno s politiko vračil. Dostop se zagotovi takoj po uspešnem plačilu. Hvala za izbiro!",
	},
	language.Lietùvių: {
		Invoice: "Mokėdami sutinkate su paslaugos taisyklėmis, įskaitant grąžinimo politiką. Prieiga suteikiama iš karto po sėkmingo mokėjimo. Ačiū už pasirinkimą!",
	},
	language.Latviešu: {
		Invoice: "Maksājot, jūs piekrītat pakalpojuma noteikumiem, ieskaitot atmaksas politiku. Piekļuve tiek nodrošināta nekavējoties pēc veiksmīga maksājuma. Paldies par izvēli!",
	},
	language.Eesti: {
		Invoice: "Makstes nõustute teenuse reeglitega, sealhulgas tagastuspoliitikaga. Juurdepääs antakse kohe pärast edukat makset. Täname valiku eest!",
	},
	language.Suomi: {
		Invoice: "Maksamalla hyväksyt palvelun säännöt, mukaan lukien palautuskäytännön. Pääsy myönnetään heti onnistuneen maksun jälkeen. Kiitos valinnastasi!",
	},
	language.Ελληνικά: {
		Invoice: "Με την πληρωμή, συμφωνείτε με τους κανόνες της υπηρεσίας, συμπεριλαμβανομένης της πολιτικής επιστροφών. Η πρόσβαση παρέχεται αμέσως μετά από επιτυχή πληρωμή. Ευχαριστούμε για την επιλογή σας!",
	},
	language.Română: {
		Invoice: "Prin plată, sunteți de acord cu regulile serviciului, inclusiv politica de rambursări. Accesul este furnizat imediat după plata reușită. Mulțumim pentru alegere!",
	},
	language.Magyar: {
		Invoice: "Fizetéssel elfogadja a szolgáltatás szabályait, beleértve a visszatérítési irányelveket. A hozzáférés azonnal biztosítva van a sikeres fizetés után. Köszönjük a választást!",
	},
	language.Arabic: {
		Invoice: "عند الدفع، توافق على قواعد الخدمة، بما في ذلك سياسة الاسترداد. يتم تقديم الوصول فوراً بعد الدفع الناجح. شكراً لاختيارك!",
	},
	language.Farsi: {
		Invoice: "با پرداخت، با قوانین سرویس موافقت می‌کنید، از جمله سیاست بازپرداخت. دسترسی بلافاصله پس از پرداخت موفق ارائه می‌شود. ممنون از انتخاب شما!",
	},
	language.Türkçe: {
		Invoice: "Ödeme yaparak, hizmet kurallarını kabul edersiniz, iade politikası dahil. Erişim, başarılı ödeme sonrasında hemen sağlanır. Seçiminiz için teşekkürler!",
	},
	language.Hebrew: {
		Invoice: "על ידי תשלום, אתה מסכים לכללי השירות, כולל מדיניות ההחזרים. הגישה מסופקת מיד לאחר תשלום מוצלח. תודה על הבחירה!",
	},
	language.ZH中文: {
		Invoice: "通过支付，您同意服务规则，包括退款政策。访问权限将在成功支付后立即提供。感谢您的选择！",
	},
	language.JA日本語: {
		Invoice: "支払うことで、サービス規則に同意します。これには返金ポリシーが含まれます。アクセスは成功した支払いの直後に提供されます。ご選択いただきありがとうございます！",
	},
	language.KO한국어: {
		Invoice: "결제함으로써 서비스 규칙에 동의합니다. 환불 정책을 포함합니다. 성공적인 결제 후 즉시 액세스가 제공됩니다. 선택해 주셔서 감사합니다!",
	},
	language.TiếngViệt: {
		Invoice: "Bằng cách thanh toán, bạn đồng ý với quy tắc dịch vụ, bao gồm chính sách hoàn tiền. Truy cập được cung cấp ngay sau khi thanh toán thành công. Cảm ơn vì đã chọn!",
	},
	language.THภาษาไทย: {
		Invoice: "โดยการชำระเงิน คุณยอมรับกฎของบริการ รวมถึงนโยบายการคืนเงิน การเข้าถึงจะถูกให้ทันทีหลังจากการชำระเงินสำเร็จ ขอบคุณสำหรับการเลือก!",
	},
	language.BahasaIndonesia: {
		Invoice: "Dengan membayar, Anda setuju dengan aturan layanan, termasuk kebijakan pengembalian dana. Akses diberikan segera setelah pembayaran berhasil. Terima kasih atas pilihan Anda!",
	},
	language.BahasaMelayu: {
		Invoice: "Dengan membayar, anda bersetuju dengan peraturan perkhidmatan, termasuk dasar pemulangan. Akses diberikan segera selepas pembayaran berjaya. Terima kasih atas pilihan anda!",
	},
	language.Tagalog: {
		Invoice: "Sa pagbabayad, sumasang-ayon ka sa mga tuntunin ng serbisyo, kabilang ang patakaran sa refund. Ang access ay ibinibigay kaagad pagkatapos ng matagumpay na pagbabayad. Salamat sa pagpili!",
	},
	language.Hindi: {
		Invoice: "भुगतान करके, आप सेवा के नियमों से सहमत होते हैं, जिसमें रिफंड नीति शामिल है। सफल भुगतान के तुरंत बाद एक्सेस प्रदान किया जाता है। आपके चुनाव के लिए धन्यवाद!",
	},
	language.URاردو: {
		Invoice: "ادائیگی کرتے ہوئے، آپ سروس کے قواعد سے اتفاق کرتے ہیں، بشمول ریفنڈ پالیسی۔ کامیاب ادائیگی کے فوراً بعد رسائی فراہم کی جاتی ہے۔ انتخاب کے لیے شکریہ!",
	},
	language.Bengali: {
		Invoice: "পেমেন্ট করে, আপনি সার্ভিসের নিয়মাবলীতে সম্মত হন, রিফান্ড নীতি সহ। সফল পেমেন্টের পরপরই অ্যাক্সেস প্রদান করা হয়। আপনার পছন্দের জন্য ধন্যবাদ!",
	},
	language.Tamiḻ: {
		Invoice: "பணம் செலுத்துவதன் மூலம், சேவை விதிகளுக்கு உடன்படுகிறீர்கள், திருப்பி அளிக்கும் கொள்கை உட்பட. வெற்றிகரமான பணம் செலுத்திய உடனே அணுகல் வழங்கப்படும். தேர்வுக்கு நன்றி!",
	},
	language.Telugu: {
		Invoice: "చెల్లింపు చేయడం ద్వారా, మీరు సర్వీస్ నియమాలకు అంగీకరిస్తున్నారు, రిఫండ్ పాలసీతో సహా. విజయవంతమైన చెల్లింపు తర్వాత వెంటనే యాక్సెస్ అందించబడుతుంది. మీ ఎంపికకు ధన్యవాదాలు!",
	},
	language.Marathi: {
		Invoice: "पैसे देऊन, तुम्ही सेवा नियमांना सहमत आहात, रिफंड धोरणासह. यशस्वी पेमेंटनंतर लगेच प्रवेश प्रदान केला जातो. तुमच्या निवडीसाठी धन्यवाद!",
	},
}

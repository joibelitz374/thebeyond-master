package i18n

import "github.com/quickpowered/thebeyond-master/configs/language"

type NewKeyLocale struct {
	ConfirmPrompt  string
	SuccessMessage string
}

var NewKeyMessages = map[language.Language]NewKeyLocale{
	language.English: {
		ConfirmPrompt:  "<b>Do you really want to update the key?</b>\n\n<b>When is this useful?</b>\n• If your key has fallen into the wrong hands\n• VPN is not working (pulls all changes)\n\n<b>You can update the key no more than once every 3 days!</b>",
		SuccessMessage: "Key updated!",
	},
	language.Deutsch: {
		ConfirmPrompt:  "<b>Möchten Sie den Schlüssel wirklich aktualisieren?</b>\n\n<b>Wann ist das nützlich?</b>\n• Wenn Ihr Schlüssel in falsche Hände geraten ist\n• VPN funktioniert nicht (zieht alle Änderungen)\n\n<b>Sie können den Schlüssel nicht öfter als einmal alle 3 Tage aktualisieren!</b>",
		SuccessMessage: "Schlüssel aktualisiert!",
	},
	language.Nederlands: {
		ConfirmPrompt:  "<b>Wilt u de sleutel echt bijwerken?</b>\n\n<b>Wanneer is dit nuttig?</b>\n• Als uw sleutel in verkeerde handen is gevallen\n• VPN werkt niet (haalt alle wijzigingen op)\n\n<b>U kunt de sleutel niet vaker dan eens per 3 dagen bijwerken!</b>",
		SuccessMessage: "Sleutel bijgewerkt!",
	},
	language.Svenska: {
		ConfirmPrompt:  "<b>Vill du verkligen uppdatera nyckeln?</b>\n\n<b>När är detta användbart?</b>\n• Om din nyckel har hamnat i fel händer\n• VPN fungerar inte (hämtar alla ändringar)\n\n<b>Du kan inte uppdatera nyckeln mer än en gång var 3:e dag!</b>",
		SuccessMessage: "Nyckeln uppdaterad!",
	},
	language.Norsk: {
		ConfirmPrompt:  "<b>Vil du virkelig oppdatere nøkkelen?</b>\n\n<b>Når er dette nyttig?</b>\n• Hvis nøkkelen din har havnet i feil hender\n• VPN fungerer ikke (henter alle endringer)\n\n<b>Du kan ikke oppdatere nøkkelen mer enn én gang hver 3. dag!</b>",
		SuccessMessage: "Nøkkel oppdatert!",
	},
	language.Dansk: {
		ConfirmPrompt:  "<b>Vil du virkelig opdatere nøglen?</b>\n\n<b>Hvornår er dette nyttigt?</b>\n• Hvis din nøgle er faldet i forkerte hænder\n• VPN virker ikke (henter alle ændringer)\n\n<b>Du kan ikke opdatere nøglen mere end én gang hver 3. dag!</b>",
		SuccessMessage: "Nøgle opdateret!",
	},
	language.Español: {
		ConfirmPrompt:  "<b>¿Realmente desea actualizar la clave?</b>\n\n<b>¿Cuándo es útil esto?</b>\n• Si su clave ha caído en manos equivocadas\n• La VPN no funciona (obtiene todos los cambios)\n\n<b>¡No puede actualizar la clave más de una vez cada 3 días!</b>",
		SuccessMessage: "¡Clave actualizada!",
	},
	language.Français: {
		ConfirmPrompt:  "<b>Voulez-vous vraiment mettre à jour la clé ?</b>\n\n<b>Quand est-ce utile ?</b>\n• Si votre clé est tombée entre de mauvaises mains\n• Le VPN ne fonctionne pas (récupère toutes les modifications)\n\n<b>Vous ne pouvez mettre à jour la clé qu'une fois tous les 3 jours !</b>",
		SuccessMessage: "Clé mise à jour !",
	},
	language.Português: {
		ConfirmPrompt:  "<b>Deseja realmente atualizar a chave?</b>\n\n<b>Quando isso é útil?</b>\n• Se a sua chave caiu em mãos erradas\n• A VPN não funciona (obtém todas as alterações)\n\n<b>Você não pode atualizar a chave mais de uma vez a cada 3 dias!</b>",
		SuccessMessage: "Chave atualizada!",
	},
	language.Italiana: {
		ConfirmPrompt:  "<b>Vuoi davvero aggiornare la chiave?</b>\n\n<b>Quando è utile?</b>\n• Se la tua chiave è finita nelle mani sbagliate\n• La VPN non funziona (recupera tutte le modifiche)\n\n<b>Non puoi aggiornare la chiave più di una volta ogni 3 giorni!</b>",
		SuccessMessage: "Chiave aggiornata!",
	},
	language.Русский: {
		ConfirmPrompt:  "<b>Вы действительно хотите обновить ключ?</b>\n\n<b>Когда это полезно?</b>\n• Если Ваш ключ попал в чужие руки\n• Не работает VPN (подтягивает все изменения)\n\n<b>Обновлять ключ можно не более 1 раза в 3 дня!</b>",
		SuccessMessage: "Ключ обновлен!",
	},
	language.Українська: {
		ConfirmPrompt:  "<b>Ви дійсно хочете оновити ключ?</b>\n\n<b>Коли це корисно?</b>\n• Якщо Ваш ключ потрапив у чужі руки\n• Не працює VPN (підтягує всі зміни)\n\n<b>Оновлювати ключ можна не частіше 1 разу на 3 дні!</b>",
		SuccessMessage: "Ключ оновлено!",
	},
	language.Polski: {
		ConfirmPrompt:  "<b>Czy na pewno chcesz zaktualizować klucz?</b>\n\n<b>Kiedy jest to przydatne?</b>\n• Jeśli Twój klucz wpadł w niepowołane ręce\n• VPN nie działa (pobiera wszystkie zmiany)\n\n<b>Możesz zaktualizować klucz nie częściej niż raz na 3 dni!</b>",
		SuccessMessage: "Klucz zaktualizowany!",
	},
	language.Ceština: {
		ConfirmPrompt:  "<b>Opravdu chcete aktualizovat klíč?</b>\n\n<b>Kdy je to užitečné?</b>\n• Pokud se váš klíč dostal do nesprávných rukou\n• VPN nefunguje (stahuje všechny změny)\n\n<b>Klíč nemůžete aktualizovat více než jednou za 3 dny!</b>",
		SuccessMessage: "Klíč aktualizován!",
	},
	language.Български: {
		ConfirmPrompt:  "<b>Наистина ли искате да актуализирате ключа?</b>\n\n<b>Кога е полезно това?</b>\n• Ако ключът ви е попаднал в грешни ръце\n• VPN не работи (изтегля всички промени)\n\n<b>Можете да актуализирате ключа не повече от веднъж на всеки 3 дни!</b>",
		SuccessMessage: "Ключът е актуализиран!",
	},
	language.Српски: {
		ConfirmPrompt:  "<b>Да ли заиста желите да ажурирате кључ?</b>\n\n<b>Када је ово корисно?</b>\n• Ако је ваш кључ доспео у погрешне руке\n• ВПН не ради (преузима све промене)\n\n<b>Кључ можете ажурирати највише једном у 3 дана!</b>",
		SuccessMessage: "Кључ ажуриран!",
	},
	language.Hrvatski: {
		ConfirmPrompt:  "<b>Želite li doista ažurirati ključ?</b>\n\n<b>Kada je to korisno?</b>\n• Ako je vaš ključ pao u pogrešne ruke\n• VPN ne radi (preuzima sve promjene)\n\n<b>Ključ ne možete ažurirati više od jednom svaka 3 dana!</b>",
		SuccessMessage: "Ključ ažuriran!",
	},
	language.Slovenčina: {
		ConfirmPrompt:  "<b>Naozaj chcete aktualizovať kľúč?</b>\n\n<b>Kedy je to užitočné?</b>\n• Ak sa váš kľúč dostal do nesprávnych rúk\n• VPN nefunguje (sťahuje všetky zmeny)\n\n<b>Kľúč nemôžete aktualizovať viac ako raz za 3 dni!</b>",
		SuccessMessage: "Kľúč aktualizovaný!",
	},
	language.Slovenski: {
		ConfirmPrompt:  "<b>Ali res želite posodobiti ključ?</b>\n\n<b>Kdaj je to koristno?</b>\n• Če je vaš ključ prišel v napačne roke\n• VPN ne deluje (potegne vse spremembe)\n\n<b>Ključa ne morete posodobiti več kot enkrat na 3 dni!</b>",
		SuccessMessage: "Ključ posodobljen!",
	},
	language.Lietùvių: {
		ConfirmPrompt:  "<b>Ar tikrai norite atnaujinti raktą?</b>\n\n<b>Kada tai naudinga?</b>\n• Jei jūsų raktas pateko į netinkamas rankas\n• VPN neveikia (ištraukia visus pakeitimus)\n\n<b>Raktą galite atnaujinti ne dažniau kaip kartą per 3 dienas!</b>",
		SuccessMessage: "Raktas atnaujintas!",
	},
	language.Latviešu: {
		ConfirmPrompt:  "<b>Vai tiešām vēlaties atjaunināt atslēgu?</b>\n\n<b>Kad tas ir noderīgi?</b>\n• Ja jūsu atslēga ir nonākusi nepareizās rokās\n• VPN nedarbojas (velk visas izmaiņas)\n\n<b>Atslēgu varat atjaunināt ne biežāk kā reizi 3 dienās!</b>",
		SuccessMessage: "Atslēga atjaunināta!",
	},
	language.Eesti: {
		ConfirmPrompt:  "<b>Kas soovite tõesti võtit värskendada?</b>\n\n<b>Millal on see kasulik?</b>\n• Kui teie võti on sattunud valedesse kätesse\n• VPN ei tööta (tõmbab kõik muudatused)\n\n<b>Võtit saate värskendada mitte rohkem kui üks kord 3 päeva jooksul!</b>",
		SuccessMessage: "Võti värskendatud!",
	},
	language.Suomi: {
		ConfirmPrompt:  "<b>Haluatko todella päivittää avaimen?</b>\n\n<b>Milloin tämä on hyödyllistä?</b>\n• Jos avaimesi on joutunut vääriin käsiin\n• VPN ei toimi (hakee kaikki muutokset)\n\n<b>Voit päivittää avaimen enintään kerran 3 päivässä!</b>",
		SuccessMessage: "Avain päivitetty!",
	},
	language.Ελληνικά: {
		ConfirmPrompt:  "<b>Θέλετε πραγματικά να ενημερώσετε το κλειδί;</b>\n\n<b>Πότε είναι χρήσιμο αυτό;</b>\n• Εάν το κλειδί σας έχει πέσει σε λάθος χέρια\n• Το VPN δεν λειτουργεί (τραβάει όλες τις αλλαγές)\n\n<b>Μπορείτε να ενημερώσετε το κλειδί όχι περισσότερο από μία φορά κάθε 3 ημέρες!</b>",
		SuccessMessage: "Το κλειδί ενημερώθηκε!",
	},
	language.Română: {
		ConfirmPrompt:  "<b>Chiar doriți să actualizați cheia?</b>\n\n<b>Când este util acest lucru?</b>\n• Dacă cheia dvs. a ajuns pe mâini greșite\n• VPN-ul nu funcționează (preia toate modificările)\n\n<b>Nu puteți actualiza cheia mai mult de o dată la 3 zile!</b>",
		SuccessMessage: "Cheie actualizată!",
	},
	language.Magyar: {
		ConfirmPrompt:  "<b>Valóban frissíteni szeretné a kulcsot?</b>\n\n<b>Mikor hasznos ez?</b>\n• Ha a kulcsa rossz kezekbe került\n• A VPN nem működik (minden változtatást lehív)\n\n<b>A kulcsot legfeljebb 3 naponta egyszer frissítheti!</b>",
		SuccessMessage: "Kulcs frissítve!",
	},
	language.Arabic: {
		ConfirmPrompt:  "<b>هل تريد حقًا تحديث المفتاح؟</b>\n\n<b>متى يكون هذا مفيدًا؟</b>\n• إذا وقع مفتاحك في الأيدي الخطأ\n• VPN لا يعمل (يسحب جميع التغييرات)\n\n<b>لا يمكنك تحديث المفتاح أكثر من مرة واحدة كل 3 أيام!</b>",
		SuccessMessage: "تم تحديث المفتاح!",
	},
	language.Farsi: {
		ConfirmPrompt:  "<b>آیا واقعاً می‌خواهید کلید را به‌روزرسانی کنید؟</b>\n\n<b>چه زمانی این مفید است؟</b>\n• اگر کلید شما به دست افراد نادرست افتاده است\n• VPN کار نمی‌کند (تمام تغییرات را دریافت می‌کند)\n\n<b>شما نمی‌توانید کلید را بیش از یک بار در هر ۳ روز به‌روزرسانی کنید!</b>",
		SuccessMessage: "کلید به‌روزرسانی شد!",
	},
	language.Türkçe: {
		ConfirmPrompt:  "<b>Anahtarı gerçekten güncellemek istiyor musunuz?</b>\n\n<b>Bu ne zaman yararlıdır?</b>\n• Anahtarınız yanlış ellere geçtiyse\n• VPN çalışmıyor (tüm değişiklikleri çeker)\n\n<b>Anahtarı 3 günde bir kereden fazla güncelleyemezsiniz!</b>",
		SuccessMessage: "Anahtar güncellendi!",
	},
	language.Hebrew: {
		ConfirmPrompt:  "<b>האם אתה באמת רוצה לעדכן את המפתח?</b>\n\n<b>מתי זה שימושי?</b>\n• אם המפתח שלך נפל לידיים הלא נכונות\n• ה-VPN לא עובד (מושך את כל השינויים)\n\n<b>לא ניתן לעדכן את המפתח יותר מפעם אחת ב-3 ימים!</b>",
		SuccessMessage: "המפתח עודכן!",
	},
	language.ZH中文: {
		ConfirmPrompt:  "<b>您真的要更新密钥吗？</b>\n\n<b>这在什么时候有用？</b>\n• 如果您的密钥落入坏人手中\n• VPN 不工作（拉取所有更改）\n\n<b>您每 3 天最多只能更新一次密钥！</b>",
		SuccessMessage: "密钥已更新！",
	},
	language.JA日本語: {
		ConfirmPrompt:  "<b>本当にキーを更新しますか？</b>\n\n<b>どのような場合に役立ちますか？</b>\n• キーが他人の手に渡った場合\n• VPNが機能しない（すべての変更を取得）\n\n<b>キーの更新は3日に1回までです！</b>",
		SuccessMessage: "キーが更新されました！",
	},
	language.KO한국어: {
		ConfirmPrompt:  "<b>정말 키를 업데이트하시겠습니까?</b>\n\n<b>언제 유용합니까?</b>\n• 키가 잘못된 사람의 손에 들어간 경우\n• VPN이 작동하지 않음 (모든 변경 사항 가져오기)\n\n<b>3일에 한 번만 키를 업데이트할 수 있습니다!</b>",
		SuccessMessage: "키가 업데이트되었습니다!",
	},
	language.TiếngViệt: {
		ConfirmPrompt:  "<b>Bạn có thực sự muốn cập nhật khóa không?</b>\n\n<b>Khi nào điều này hữu ích?</b>\n• Nếu khóa của bạn rơi vào tay kẻ xấu\n• VPN không hoạt động (kéo tất cả các thay đổi)\n\n<b>Bạn không thể cập nhật khóa quá 1 lần sau mỗi 3 ngày!</b>",
		SuccessMessage: "Khóa đã được cập nhật!",
	},
	language.THภาษาไทย: {
		ConfirmPrompt:  "<b>คุณต้องการอัปเดตคีย์จริงหรือไม่?</b>\n\n<b>เมื่อไหร่จึงจะมีประโยชน์?</b>\n• หากคีย์ของคุณตกไปอยู่ในมือคนผิด\n• VPN ไม่ทำงาน (ดึงการเปลี่ยนแปลงทั้งหมด)\n\n<b>คุณสามารถอัปเดตคีย์ได้ไม่เกิน 1 ครั้งในทุกๆ 3 วัน!</b>",
		SuccessMessage: "อัปเดตคีย์แล้ว!",
	},
	language.BahasaIndonesia: {
		ConfirmPrompt:  "<b>Apakah Anda benar-benar ingin memperbarui kunci?</b>\n\n<b>Kapan ini berguna?</b>\n• Jika kunci Anda jatuh ke tangan yang salah\n• VPN tidak berfungsi (menarik semua perubahan)\n\n<b>Anda tidak dapat memperbarui kunci lebih dari 1 kali setiap 3 hari!</b>",
		SuccessMessage: "Kunci diperbarui!",
	},
	language.BahasaMelayu: {
		ConfirmPrompt:  "<b>Adakah anda benar-benar mahu mengemas kini kunci?</b>\n\n<b>Bilakah ini berguna?</b>\n• Jika kunci anda jatuh ke tangan yang salah\n• VPN tidak berfungsi (menarik semua perubahan)\n\n<b>Anda tidak boleh mengemas kini kunci lebih daripada 1 kali setiap 3 hari!</b>",
		SuccessMessage: "Kunci dikemas kini!",
	},
	language.Tagalog: {
		ConfirmPrompt:  "<b>Talaga bang gusto mong i-update ang key?</b>\n\n<b>Kailan ito kapaki-pakinabang?</b>\n• Kung ang iyong key ay napunta sa maling mga kamay\n• Hindi gumagana ang VPN (hinihila ang lahat ng mga pagbabago)\n\n<b>Hindi mo ma-update ang key nang higit sa 1 beses bawat 3 araw!</b>",
		SuccessMessage: "Na-update na ang key!",
	},
	language.Hindi: {
		ConfirmPrompt:  "<b>क्या आप वास्तव में कुंजी अपडेट करना चाहते हैं?</b>\n\n<b>यह कब उपयोगी है?</b>\n• यदि आपकी कुंजी गलत हाथों में पड़ गई है\n• वीपीएन काम नहीं कर रहा है (सभी परिवर्तन खींचता है)\n\n<b>आप हर 3 दिनों में 1 बार से अधिक कुंजी अपडेट नहीं कर सकते!</b>",
		SuccessMessage: "कुंजी अपडेट की गई!",
	},
	language.URاردو: {
		ConfirmPrompt:  "<b>کیا آپ واقعی کلید کو اپ ڈیٹ کرنا چاہتے ہیں؟</b>\n\n<b>یہ کب مفید ہے؟</b>\n• اگر آپ کی کلید غلط ہاتھوں میں پڑ گئی ہے\n• VPN کام نہیں کر رہا ہے (تمام تبدیلیاں کھینچتا ہے)\n\n<b>آپ ہر 3 دن میں 1 بار سے زیادہ کلید کو اپ ڈیٹ نہیں کر سکتے!</b>",
		SuccessMessage: "کلید اپ ڈیٹ ہو گئی!",
	},
	language.Bengali: {
		ConfirmPrompt:  "<b>আপনি কি সত্যিই কি (key) আপডেট করতে চান?</b>\n\n<b>কখন এটি দরকারী?</b>\n• যদি আপনার কি ভুল হাতে পড়ে থাকে\n• ভিপিএন কাজ করছে না (সব পরিবর্তন টানছে)\n\n<b>আপনি প্রতি ৩ দিনে ১ বারের বেশি কি আপডেট করতে পারবেন না!</b>",
		SuccessMessage: "কি আপডেট করা হয়েছে!",
	},
	language.Tamiḻ: {
		ConfirmPrompt:  "<b>நீங்கள் உண்மையில் சாவியைப் புதுப்பிக்க விரும்புகிறீர்களா?</b>\n\n<b>இது எப்போது பயனுள்ளதாக இருக்கும்?</b>\n• உங்கள் சாவி தவறான கைகளில் சிக்கியிருந்தால்\n• VPN வேலை செய்யவில்லை (எல்லா மாற்றங்களையும் இழுக்கிறது)\n\n<b>ஒவ்வொரு 3 நாட்களுக்கும் 1 முறைக்கு மேல் சாவியைப் புதுப்பிக்க முடியாது!</b>",
		SuccessMessage: "சாவி புதுப்பிக்கப்பட்டது!",
	},
	language.Telugu: {
		ConfirmPrompt:  "<b>మీరు నిజంగా కీని అప్‌డేట్ చేయాలనుకుంటున్నారా?</b>\n\n<b>ఇది ఎప్పుడు ఉపయోగపడుతుంది?</b>\n• మీ కీ తప్పుడు చేతుల్లో పడితే\n• VPN పని చేయడం లేదు (అన్ని మార్పులను లాగుతుంది)\n\n<b>మీరు ప్రతి 3 రోజులకు 1 సారి కంటే ఎక్కువ కీని అప్‌డేట్ చేయలేరు!</b>",
		SuccessMessage: "కీ అప్‌డేట్ చేయబడింది!",
	},
	language.Marathi: {
		ConfirmPrompt:  "<b>तुम्हाला खरोखर की अपडेट करायची आहे का?</b>\n\n<b>हे कधी उपयुक्त आहे?</b>\n• जर तुमची की चुकीच्या हातात पडली असेल\n• VPN काम करत नाही (सर्व बदल खेचते)\n\n<b>तुम्ही दर 3 दिवसांनी 1 पेक्षा जास्त वेळा की अपडेट करू शकत नाही!</b>",
		SuccessMessage: "की अपडेट केली!",
	},
}

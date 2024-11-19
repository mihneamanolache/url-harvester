package constants

var PageCategoryKeywords = map[string]map[string][]string{
    "homepage": {
        "en": {`^(https?://)?[^/]+/?(home|index|landing|start|main)?/?`},
        "fr": {`^(https?://)?[^/]+/?(accueil|index|page-d-accueil|page-principale)?/?`},
        "de": {`^(https?://)?[^/]+/?(startseite|index|hauptseite|landing)?/?`},
        "es": {`^(https?://)?[^/]+/?(inicio|principal|pagina-principal|pagina-de-inicio)?/?`},
        "pt": {`^(https?://)?[^/]+/?(inicio|pagina-principal|index|pagina-inicial)?/?`},
        "it": {`^(https?://)?[^/]+/?(home|pagina-iniziale|indice|pagina-principale)?/?`},
        "ru": {`^(https?://)?[^/]+/?(главная|index|начальная|основная-страница)?/?`},
        "zh": {`^(https?://)?[^/]+/?(首页|主页|开始|索引)?/?`},
        "ja": {`^(https?://)?[^/]+/?(ホーム|トップ|メインページ|インデックス)?/?`},
        "ko": {`^(https?://)?[^/]+/?(홈|메인-페이지|시작-페이지|인덱스)?/?`},
        "ar": {`^(https?://)?[^/]+/?(الصفحة-الرئيسية|الرئيسية|الصفحة-الاولى)?/?`},
        "hi": {`^(https?://)?[^/]+/?(मुख्य-पृष्ठ|index|होम-पेज)?/?`},
        "bn": {`^(https?://)?[^/]+/?(হোম|প্রধান-পৃষ্ঠা|ইনডেক্স)?/?`},
        "tr": {`^(https?://)?[^/]+/?(ana-sayfa|giriş|index|başlangıç)?/?`},
        "pl": {`^(https?://)?[^/]+/?(strona-główna|index|start|strona-początkowa)?/?`},
    },

    "blog": {
        "en": {`^(https?://)?[^/]+/?(blog|news|article|post|update)s?/?`},
        "fr": {`^(https?://)?[^/]+/?(blog|actualités|article|post|mise-à-jour)s?/?`},
        "de": {`^(https?://)?[^/]+/?(blog|news|artikel|nachricht|beitrag)s?/?`},
        "es": {`^(https?://)?[^/]+/?(blog|noticias|artículo|post|actualización)(es)?/?`},
        "pt": {`^(https?://)?[^/]+/?(blog|notícias|artigo|post|atualização)(es)?/?`},
        "it": {`^(https?://)?[^/]+/?(blog|notizie|articolo|post|aggiornamento)(i)?/?`},
        "ru": {`^(https?://)?[^/]+/?(блог|новости|статья|пост|обновление)(и)?/?`},
        "zh": {`^(https?://)?[^/]+/?(博客|新闻|文章|更新|帖子)(们)?/?`},
        "ja": {`^(https?://)?[^/]+/?(ブログ|ニュース|記事|更新|投稿)(たち)?/?`},
        "ko": {`^(https?://)?[^/]+/?(블로그|뉴스|게시물|업데이트)(들)?/?`},
        "ar": {`^(https?://)?[^/]+/?(مدونة|أخبار|مقال|تحديث)(ات)?/?`},
        "hi": {`^(https?://)?[^/]+/?(ब्लॉग|समाचार|लेख|पोस्ट|अपडेट)(ें)?/?`},
        "bn": {`^(https?://)?[^/]+/?(ব্লগ|খবর|লেখা|আপডেট)(গুলো)?/?`},
        "tr": {`^(https?://)?[^/]+/?(blog|haber|makale|güncelleme)(ler)?/?`},
        "pl": {`^(https?://)?[^/]+/?(blog|aktualności|artykuł|post|aktualizacja)(i)?/?`},
    },

    "sustainability": {
        "en": {`^(https?://)?[^/]+/?(sustainability|environment|green|eco)(-|/)?`},
        "fr": {`^(https?://)?[^/]+/?(durabilité|environnement|écologique|vert)(-|/)?`},
        "de": {`^(https?://)?[^/]+/?(nachhaltigkeit|umwelt|grün|ökologisch)(-|/)?`},
        "es": {`^(https?://)?[^/]+/?(sostenibilidad|medio-ambiente|verde|ecológico)(-|/)?`},
        "pt": {`^(https?://)?[^/]+/?(sustentabilidade|meio-ambiente|verde|ecológico)(-|/)?`},
        "it": {`^(https?://)?[^/]+/?(sostenibilità|ambiente|verde|ecologico)(-|/)?`},
        "ru": {`^(https?://)?[^/]+/?(устойчивость|экология|зеленый|экологичный)(-|/)?`},
        "zh": {`^(https?://)?[^/]+/?(可持续发展|环境|绿色|生态)(-|/)?`},
        "ja": {`^(https?://)?[^/]+/?(持続可能性|環境|エコ|グリーン)(-|/)?`},
        "ko": {`^(https?://)?[^/]+/?(지속-가능성|환경|녹색|친환경)(-|/)?`},
        "ar": {`^(https?://)?[^/]+/?(استدامة|بيئة|أخضر|إيكولوجي)(-|/)?`},
        "hi": {`^(https?://)?[^/]+/?(स्थिरता|पर्यावरण|हरा|इको)(-|/)?`},
        "bn": {`^(https?://)?[^/]+/?(স্থায়িত্ব|পরিবেশ|সবুজ|ইকো)(-|/)?`},
        "tr": {`^(https?://)?[^/]+/?(sürdürülebilirlik|çevre|yeşil|eko)(-|/)?`},
        "pl": {`^(https?://)?[^/]+/?(zrównoważony-rozwój|środowisko|zielony|ekologiczny)(-|/)?`},
    },

    "product": {
        "en": {`^(https?://)?[^/]+/?(product|item|goods|store|buy)s?/?`},
        "fr": {`^(https?://)?[^/]+/?(produit|article|marchandise|acheter|magasin)s?/?`},
        "de": {`^(https?://)?[^/]+/?(produkt|ware|artikel|kaufen|geschäft)s?/?`},
        "es": {`^(https?://)?[^/]+/?(producto|artículo|mercancía|comprar|tienda)s?/?`},
        "pt": {`^(https?://)?[^/]+/?(produto|artigo|mercadoria|comprar|loja)s?/?`},
        "it": {`^(https?://)?[^/]+/?(prodotto|articolo|merce|comprare|negozio)s?/?`},
        "ru": {`^(https?://)?[^/]+/?(продукт|товар|статья|покупка|магазин)s?/?`},
        "zh": {`^(https?://)?[^/]+/?(产品|商品|物品|购买|商店)s?/?`},
        "ja": {`^(https?://)?[^/]+/?(製品|商品|アイテム|購入|店舗)s?/?`},
        "ko": {`^(https?://)?[^/]+/?(제품|상품|아이템|구매|상점)s?/?`},
        "ar": {`^(https?://)?[^/]+/?(منتج|سلعة|مادة|شراء|متجر)s?/?`},
        "hi": {`^(https?://)?[^/]+/?(उत्पाद|सामान|वस्तु|खरीद|दुकान)s?/?`},
        "bn": {`^(https?://)?[^/]+/?(পণ্য|বস্তু|মাল|ক্রয়|দোকান)s?/?`},
        "tr": {`^(https?://)?[^/]+/?(ürün|eşya|mal|satın-al|mağaza)s?/?`},
        "pl": {`^(https?://)?[^/]+/?(produkt|towar|artykuł|zakup|sklep)s?/?`},
    },

    "contact": {
        "en": {`^(https?://)?[^/]+/?(contact|contact-us|get-in-touch|reach-out|form|email|contact-page|contact-info)/?`},
        "fr": {`^(https?://)?[^/]+/?(contact|nous-contacter|formulaire|email|joindre)s?/?`},
        "de": {`^(https?://)?[^/]+/?(kontakt|kontaktformular|erreichen|email)s?/?`},
        "es": {`^(https?://)?[^/]+/?(contacto|ponerse-en-contacto|formulario|email)s?/?`},
        "pt": {`^(https?://)?[^/]+/?(contato|entrar-em-contato|formulário|email)s?/?`},
        "it": {`^(https?://)?[^/]+/?(contatto|mettersi-in-contatto|formulario|email)s?/?`},
        "ru": {`^(https?://)?[^/]+/?(контакт|связаться|форма|email)s?/?`},
        "zh": {`^(https?://)?[^/]+/?(联系|联系我们|表格|邮箱)s?/?`},
        "ja": {`^(https?://)?[^/]+/?(連絡|お問い合わせ|フォーム|メール)s?/?`},
        "ko": {`^(https?://)?[^/]+/?(연락처|문의|양식|이메일)s?/?`},
        "ar": {`^(https?://)?[^/]+/?(اتصال|تواصل|نموذج|البريد-الإلكتروني)s?/?`},
        "hi": {`^(https?://)?[^/]+/?(संपर्क|हमसे-जुड़े|प्रपत्र|ईमेल)s?/?`},
        "bn": {`^(https?://)?[^/]+/?(যোগাযোগ|ফর্ম|ইমেল|আমাদের-সঙ্গে-যোগাযোগ করুন)s?/?`},
        "tr": {`^(https?://)?[^/]+/?(iletisim|bize-ulaşın|form|e-posta)s?/?`},
        "pl": {`^(https?://)?[^/]+/?(kontakt|skontaktuj-się|formularz|email)s?/?`},
    },

    "news": {
        "en": {`^(https?://)?[^/]+/?(news|updates|articles|press|announcements)s?/?`},
        "fr": {`^(https?://)?[^/]+/?(actualités|mises-à-jour|articles|communiqué|annonces)s?/?`},
        "de": {`^(https?://)?[^/]+/?(nachrichten|aktualisierungen|artikel|presse|ankündigungen)s?/?`},
        "es": {`^(https?://)?[^/]+/?(noticias|actualizaciones|artículos|prensa|anuncios)s?/?`},
        "pt": {`^(https?://)?[^/]+/?(notícias|atualizações|artigos|imprensa|anúncios)s?/?`},
        "it": {`^(https?://)?[^/]+/?(notizie|aggiornamenti|articoli|comunicati|annunci)s?/?`},
        "ru": {`^(https?://)?[^/]+/?(новости|обновления|статьи|пресса|объявления)s?/?`},
        "zh": {`^(https?://)?[^/]+/?(新闻|更新|文章|新闻稿|公告)s?/?`},
        "ja": {`^(https?://)?[^/]+/?(ニュース|更新|記事|プレス|発表)s?/?`},
        "ko": {`^(https?://)?[^/]+/?(뉴스|업데이트|기사|보도|발표)s?/?`},
        "ar": {`^(https?://)?[^/]+/?(أخبار|التحديثات|مقالات|صحافة|إعلانات)s?/?`},
        "hi": {`^(https?://)?[^/]+/?(समाचार|अपडेट्स|लेख|प्रेस|घोषणाएं)s?/?`},
        "bn": {`^(https?://)?[^/]+/?(খবর|আপডেট|লেখা|প্রেস|ঘোষণা)s?/?`},
        "tr": {`^(https?://)?[^/]+/?(haber|güncellemeler|makaleler|basın|duyurular)s?/?`},
        "pl": {`^(https?://)?[^/]+/?(aktualności|aktualizacje|artykuły|prasa|ogłoszenia)s?/?`},
    },

    "service": {
        "en": {`^(https?://)?[^/]+/?(services|offerings|solutions|products|support)s?/?`},
        "fr": {`^(https?://)?[^/]+/?(services|offres|solutions|produits|support)s?/?`},
        "de": {`^(https?://)?[^/]+/?(dienstleistungen|angebote|lösungen|produkte|support)s?/?`},
        "es": {`^(https?://)?[^/]+/?(servicios|ofrecimientos|soluciones|productos|soporte)s?/?`},
        "pt": {`^(https?://)?[^/]+/?(serviços|ofertas|soluções|produtos|suporte)s?/?`},
        "it": {`^(https?://)?[^/]+/?(servizi|offerte|soluzioni|prodotti|supporto)s?/?`},
        "ru": {`^(https?://)?[^/]+/?(услуги|предложения|решения|продукты|поддержка)s?/?`},
        "zh": {`^(https?://)?[^/]+/?(服务|产品|解决方案|支持|服务项目)s?/?`},
        "ja": {`^(https?://)?[^/]+/?(サービス|提供|ソリューション|製品|サポート)s?/?`},
        "ko": {`^(https?://)?[^/]+/?(서비스|제공|솔루션|제품|지원)s?/?`},
        "ar": {`^(https?://)?[^/]+/?(خدمات|عروض|حلول|منتجات|دعم)s?/?`},
        "hi": {`^(https?://)?[^/]+/?(सेवाएं|प्रस्ताव|समाधान|उत्पाद|सहायता)s?/?`},
        "bn": {`^(https?://)?[^/]+/?(সেবা|অফার|সমাধান|পণ্য|সহায়তা)s?/?`},
        "tr": {`^(https?://)?[^/]+/?(hizmet|teklif|çözümler|ürünler|destek)s?/?`},
        "pl": {`^(https?://)?[^/]+/?(usługi|oferty|rozwiązania|produkty|wsparcie)s?/?`},
    },

    "search": {
        "en": {`^(https?://)?[^/]+/?(search|find|query|results|explore)s?/?`},
        "fr": {`^(https?://)?[^/]+/?(recherche|trouver|requête|résultats|explorer)s?/?`},
        "de": {`^(https?://)?[^/]+/?(suche|finden|abfrage|ergebnisse|entdecken)s?/?`},
        "es": {`^(https?://)?[^/]+/?(buscar|encontrar|consulta|resultados|explorar)s?/?`},
        "pt": {`^(https?://)?[^/]+/?(pesquisar|encontrar|consulta|resultados|explorar)s?/?`},
        "it": {`^(https?://)?[^/]+/?(cerca|trova|query|risultati|esplora)s?/?`},
        "ru": {`^(https?://)?[^/]+/?(поиск|найти|запрос|результаты|исследовать)s?/?`},
        "zh": {`^(https?://)?[^/]+/?(搜索|查找|查询|结果|探索)s?/?`},
        "ja": {`^(https?://)?[^/]+/?(検索|見つける|クエリ|結果|探す)s?/?`},
        "ko": {`^(https?://)?[^/]+/?(검색|찾기|쿼리|결과|탐색)s?/?`},
        "ar": {`^(https?://)?[^/]+/?(بحث|إيجاد|استعلام|نتائج|استكشاف)s?/?`},
        "hi": {`^(https?://)?[^/]+/?(खोज|पाना|क्वेरी|परिणाम|अन्वेषण)s?/?`},
        "bn": {`^(https?://)?[^/]+/?(অনুসন্ধান|খুঁজুন|কুয়েরি|ফলাফল|অন্বেষণ)s?/?`},
        "tr": {`^(https?://)?[^/]+/?(arama|bulma|sorgu|sonuçlar|keşfet)s?/?`},
        "pl": {`^(https?://)?[^/]+/?(szukaj|znaleźć|zapytanie|wyniki|przeglądaj)s?/?`},
    },

    "careers": {
        "en": {`^(https?://)?[^/]+/?(careers|jobs|employment|vacancies|recruitment)s?/?`},
        "fr": {`^(https?://)?[^/]+/?(carrières|emplois|recrutement|postes-vacants)s?/?`},
        "de": {`^(https?://)?[^/]+/?(karriere|jobs|beschäftigung|stellenangebote|rekrutierung)s?/?`},
        "es": {`^(https?://)?[^/]+/?(carreras|empleos|vacantes|reclutamiento)s?/?`},
        "pt": {`^(https?://)?[^/]+/?(carreiras|empregos|vagas|recrutamento)s?/?`},
        "it": {`^(https?://)?[^/]+/?(carriere|lavori|occupazione|vacanze|reclutamento)s?/?`},
        "ru": {`^(https?://)?[^/]+/?(карьера|работа|вакансии|рекрутирование)s?/?`},
        "zh": {`^(https?://)?[^/]+/?(职业|工作|招聘|空缺|招募)s?/?`},
        "ja": {`^(https?://)?[^/]+/?(キャリア|求人|雇用|空きポジション|採用)s?/?`},
        "ko": {`^(https?://)?[^/]+/?(경력|일자리|고용|채용|구인)s?/?`},
        "ar": {`^(https?://)?[^/]+/?(وظائف|مهن|التوظيف|وظائف-شاغرة|توظيف)s?/?`},
        "hi": {`^(https?://)?[^/]+/?(करियर|नौकरियां|रोजगार|रिक्तियां|भर्ती)s?/?`},
        "bn": {`^(https?://)?[^/]+/?(ক্যারিয়ার|চাকরি|গবেষণা|খালি-পদ|নিয়োগ)s?/?`},
        "tr": {`^(https?://)?[^/]+/?(kariyer|işler|istihdam|iş-ihtiyaçları|işe-alım)s?/?`},
        "pl": {`^(https?://)?[^/]+/?(kariera|praca|zatrudnienie|wolne-stanowiska|rekrutacja)s?/?`},
    },

    "team": {
        "en": {`^(https?://)?[^/]+/?(team|about-us|our-team|meet-the-team|staff)s?/?`},
        "fr": {`^(https?://)?[^/]+/?(équipe|à-propos|notre-équipe|rencontrer-l-équipe|personnel)s?/?`},
        "de": {`^(https?://)?[^/]+/?(team|über-uns|unser-team|kennenlernen|personal)s?/?`},
        "es": {`^(https?://)?[^/]+/?(equipo|sobre-nosotros|nuestro-equipo|conoce-al-equipo|personal)s?/?`},
        "pt": {`^(https?://)?[^/]+/?(equipe|sobre-nosotros|nosso-time|conheça-a-equipe|pessoal)s?/?`},
        "it": {`^(https?://)?[^/]+/?(team|chi-siamo|nostro-team|incontra-il-team|personale)s?/?`},
        "ru": {`^(https?://)?[^/]+/?(команда|о-нас|наша-команда|познакомьтесь-с-нашей-командой|персонал)s?/?`},
        "zh": {`^(https?://)?[^/]+/?(团队|关于我们|我们的团队|认识团队|员工)s?/?`},
        "ja": {`^(https?://)?[^/]+/?(チーム|私たちについて|私たちのチーム|チームに会う|スタッフ)s?/?`},
        "ko": {`^(https?://)?[^/]+/?(팀|우리에 대해|우리 팀|팀을 만나다|직원)s?/?`},
        "ar": {`^(https?://)?[^/]+/?(فريق|عنا|فريقنا|التعرف على الفريق|الطاقم)s?/?`},
        "hi": {`^(https?://)?[^/]+/?(टीम|हमारे बारे में|हमारी टीम|टीम से मिलें|कर्मचारी)s?/?`},
        "bn": {`^(https?://)?[^/]+/?(দল|আমাদের সম্পর্কে|আমাদের দল|দলের সাথে পরিচিত হন|স্টাফ)s?/?`},
        "tr": {`^(https?://)?[^/]+/?(takım|hakkımızda|bizim-takım|takımla-tanışın|personel)s?/?`},
        "pl": {`^(https?://)?[^/]+/?(zespół|o-nas|nasz-zespół|poznaj-zespół|personel)s?/?`},
    },
}

var TldToLanguage = map[string]string{
	"com": "en", 
	"fr":  "fr", 
	"de":  "de", 
	"es":  "es", 
	"pt":  "pt", 
	"it":  "it", 
	"ru":  "ru", 
	"cn":  "zh", 
	"jp":  "ja", 
	"kr":  "ko", 
	"ar":  "ar", 
	"hi":  "hi", 
	"br":  "pt",
	"pl":  "pl", 
}

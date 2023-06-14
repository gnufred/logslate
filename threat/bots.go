package threat

// Bot user agent list.
// Include official documentation. Absent comment mean absent documentation.
var botPatterns = []string{

	// !!! Search engine indexers !!! //
	// Indexer are not consider a threat by themselves. That being said, when multiples
	// crawls a site simultenaously, they usually bring it down. This happens quite often
	// which make them a very real threat.

	// https://developers.google.com/search/docs/crawling-indexing/overview-google-crawlers
	"APIs-Google",
	"AdsBot-Google",
	"AdsBot-Google-Mobile",
	"FeedFetcher-Google",
	"Google-InspectionTool",
	"Google-Read-Aloud",
	"Google-Site-Verification",
	"GoogleOther",
	"GoogleProducer",
	"Googlebot",
	"Googlebot-Image",
	"Googlebot-News",
	"Googlebot-Video",
	"Mediapartners-Google",
	"Storebot-Google",

	// http://yandex.com/bots
	"YaDirectFetcher",
	"YandexAccessibilityBot",
	"YandexAdNet",
	"YandexBlogs",
	"YandexBot",
	"YandexCalendar",
	"YandexDirect",
	"YandexDirectDyn",
	"YandexFavicons",
	"YandexForDomain",
	"YandexImageResizer",
	"YandexImages",
	"YandexMarket",
	"YandexMedia",
	"YandexMetrika",
	"YandexMobileBot",
	"YandexMobileScreenShotBot",
	"YandexNews",
	"YandexOntoDB",
	"YandexOntoDBAPI",
	"YandexPagechecker",
	"YandexPartner",
	"YandexRCA",
	"YandexScreenshotBot",
	"YandexSearchShop",
	"YandexSitelinks",
	"YandexSpravBot",
	"YandexTracker",
	"YandexTurbo",
	"YandexVerticals",
	"YandexVertis",
	"YandexVideo",
	"YandexVideoParser",
	"YandexWebmaste",

	// https://help-baidu-com.translate.goog/question?prod_id=99&class=476&id=2996&_x_tr_sl=auto&_x_tr_tl=en&_x_tr_hl=en&_x_tr_pto=wapp
	"Baiduspider",
	"Baiduspider-image",
	"Baiduspider-video",
	"Baiduspider-news",
	"Baiduspider-favo",
	"Baiduspider-cpro",
	"Baiduspider-ads",

	// http://www.bing.com/bingbot.htm
	"MicrosoftPreview",
	"adidxbot",
	"bingbot/",
	"msnbot",

	// Other search engine indexers.
	"DuckDuckBot",             // https://duckduckgo.com/duckduckbot
	"DuckDuckGo-Favicons-Bot", // https://duckduckgo.com/duckduckbot
	"PetalBot",                // https://webmaster.petalsearch.com/site/petalbot
	"SeznamBot",               // https://napoveda-seznam-cz.translate.goog/cz/seznamcz-web-search/?_x_tr_sl=auto&_x_tr_tl=en&_x_tr_hl=en&_x_tr_pto=wapp
	"SeekportBot",             // https://bot.seekport.com

	// !!! Marketing bots !!! //
	// While most do not intend to be harmful, they can be too aggresive, badly configured,
	// crawl a site even if it's timing out. Some are just generally badly used and most
	// often than not harmful.
	"AhrefsBot",                 // http://ahrefs.com/robot/
	"AwarioBot",                 // https://awario.com/bots.html
	"AwarioRssBot",              // https://awario.com/bots.html
	"AwarioSmartBot",            // https://awario.com/bots.html
	"Barkrowler",                // https://babbar.tech/crawler
	"Buck/",                     // https://app.hypefactors.com/media-monitoring/about.html
	"Bytespider",                // https://zhanzhang.toutiao.com/
	"Caliperbot",                // http://www.conductor.com/caliperbot
	"DotBot",                    // https://opensiteexplorer.org/dotbot
	"HubSpot Crawler",           // https://www.hubspot.com
	"Screaming Frog SEO Spider", // https://www.screamingfrog.co.uk/seo-spider/
	"SemrushBot",                // https://www.semrush.com/bot/
	"SemrushBot-BA",             // https://www.semrush.com/bot/
	"SemrushBot-COUB",           // https://www.semrush.com/bot/
	"SemrushBot-CT",             // https://www.semrush.com/bot/
	"SemrushBot-SI",             // https://www.semrush.com/bot/
	"SemrushBot-SWA",            // https://www.semrush.com/bot/
	"SiteAuditBot",              // https://www.semrush.com/bot/
	"SplitSignalBot",            // https://www.semrush.com/bot/

	// !!! Content preview bots !!! //
	// Content preview bots shoud not be a threat, therefor won't be listed. They visit a
	// page when user share links, usually to show content summary, and check for malware.
	// Also, they usually cache results, therefor will usually not cause any issues.

	// !!! Other bots !!! //
	// This section include anything that's not in another category and can be a threat.
	// Reasons will widly vary based on the specific bot.
	"CCBot",               // https://commoncrawl.org/faq/
	"EdgeScan",            // https://www.edgescan.com/
	"Grammarly",           // https://www.grammarly.com/plagiarism-checker
	"InternetMeasurement", // https://internet-measurement.com/
	"Scrapy",              // https://scrapy.org
	"paloaltonetworks",    // https://www.paloaltonetworks.com/

	// Microsoft non indexer bots.
	"Microsoft Office Protocol Discovery", // https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-ofba/868d129f-f1b5-46bc-9385-4af58610dbbe
	// ^^ The protocol discovery bot can be quite aggressive.
	"Microsoft Data Access Internet Publishing Provider DAV",
	"Microsoft Office Word 2014",
}

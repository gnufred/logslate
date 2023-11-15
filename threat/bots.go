package threat

// Bot user agent list.
// Include official documentation. Absent comment mean absent documentation.
var botPatterns = []string{

	// !!! Search engine indexers !!! //
	// Indexer are not consider a threat by themselves. That being said, when
	// multiples crawls a site simultenaously, they usually bring it down. This
	// happens quite often which make them a very real threat.

	// http://www.google.com/bot.html
	// http://www.google.com/adsbot.html
	"APIs-Google",
	"AdsBot-Google",
	"AdsBot-Google-Mobile",
	"FeedFetcher-Google",
	"Google-AdWords",
	"Google-InspectionTool",
	"Google-Read-Aloud",
	"Google-Safety",
	"Google-Site-Verification",
	"GoogleAssociationService", // https://support.google.com/webmasters/answer/9419894
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

	// http://www.baidu.com/search/spider.html
	"Baiduspider",
	"Baiduspider-ads",
	"Baiduspider-cpro",
	"Baiduspider-favo",
	"Baiduspider-image",
	"Baiduspider-news",
	"Baiduspider-render",
	"Baiduspider-video",

	// http://www.bing.com/bingbot.htm
	"AdIdxBot",
	"BingPreview",
	"Bingbot",
	"MicrosoftPreview/", // https://aka.ms/MicrosoftPreview
	"adidxbot/",
	"bingbot/",
	"msnbot/",

	// Other search engine indexers.
	"DuckDuckBot",             // https://duckduckgo.com/duckduckbot
	"DuckDuckGo-Favicons-Bot", // https://duckduckgo.com/duckduckbot
	"FemtosearchBot",          // http://femtosearch.com
	"GPTBot",                  // https://platform.openai.com/docs/gptbot
	"MegaIndex.ru",            // http://megaindex.com/crawler
	"PetalBot",                // https://webmaster.petalsearch.com/site/petalbot
	"Pinterestbot",            // http://www.pinterest.com/bot.html
	"SeekportBot",             // https://bot.seekport.com
	"SeekportBot",             // https://bot.seekport.com
	"SeznamBot",               // https://napoveda-seznam-cz.translate.goog/cz/seznamcz-web-search/?_x_tr_sl=auto&_x_tr_tl=en&_x_tr_hl=en&_x_tr_pto=wapp
	"Yeti",                    // https://naver.me/spd
	"YisouSpider",

	// !!! Marketing bots !!! //
	// While most do not intend to be harmful, they can be too aggresive, badly
	// configured, crawl a site even if it's timing out. Some are just generally
	// badly used and most often than not harmful.
	"AddSearchBot",              // http://www.addsearch.com/bot/
	"AhrefsBot",                 // http://ahrefs.com/robot/
	"AhrefsSiteAudit",           // http://ahrefs.com/robot/site-audit
	"AwarioBot",                 // https://awario.com/bots.html
	"AwarioRssBot",              // https://awario.com/bots.html
	"AwarioSmartBot",            // https://awario.com/bots.html
	"BLEXBot",                   // http://webmeup-crawler.com/
	"Barkrowler",                // https://babbar.tech/crawler
	"BrightEdge Crawler",        // https://www.brightedge.com/products/s3/contentiq
	"Buck/",                     // https://app.hypefactors.com/media-monitoring/about.html
	"Bytespider",                // https://zhanzhang.toutiao.com/
	"Caliperbot",                // http://www.conductor.com/caliperbot
	"DataForSeoBot",             // https://dataforseo.com/dataforseo-bot
	"DomainStatsBot",            // https://domainstats.com/pages/our-bot
	"DotBot",                    // https://opensiteexplorer.org/dotbot
	"HubSpot Crawler",           // https://www.hubspot.com
	"Klaviyo",                   // https://www.klaviyo.com/
	"Konduto",                   // http://www.konduto.com
	"MJ12bot",                   // http://mj12bot.com/
	"MailChimp",                 // https://mailchimp.com/
	"Monsidobot",                // https://monsido.com/bot-html
	"Productsup",                // http://productsup.io
	"Screaming Frog SEO Spider", // https://www.screamingfrog.co.uk/seo-spider/
	"SemrushBot",                // https://www.semrush.com/bot/
	"SemrushBot-BA",             // https://www.semrush.com/bot/
	"SemrushBot-COUB",           // https://www.semrush.com/bot/
	"SemrushBot-CT",             // https://www.semrush.com/bot/
	"SemrushBot-SI",             // https://www.semrush.com/bot/
	"SemrushBot-SWA",            // https://www.semrush.com/bot/
	"SiteAuditBot",              // http://www.semrush.com/bot.html
	"SiteAuditBot",              // https://www.semrush.com/bot/
	"SplitSignalBot",            // https://www.semrush.com/bot/
	"magpie-crawler",            // http://www.brandwatch.net
	"rogerbot/",                 // https: //moz.com/help/moz-procedures/crawlers/rogerbot
	"serpstatbot/",              // https://serpstatbot.com
	// https://dynamics.microsoft.com/en-us/business-central/overview/
	"Microsoft Dynamics 365 Business Central",

	// !!! Content preview/fetchers bots !!! //
	// Content previewers shoud not be a threat, They visit a page when user share
	// links, usually to show content summary, and check for malware. Also, they
	// usually cache results, therefor will usually not cause any issues.
	// There are some very weird edge cases where they can be harmful, so it's
	// important to detect them.
	"GoogleDocs",       // http://docs.google.com
	"GoogleImageProxy", // https://www.scientiamobile.com/what-is-google-image-proxy/
	"Microsoft Office",
	"SkypeUriPreview",
	"Slackbot-LinkExpanding",   // https://api.slack.com/robots
	"Snap URL Preview Service", // https://developers.snap.com/robots
	"SnapchatAds",              // https://businesshelp.snapchat.com/s/article/adsbot-crawler
	"TelegramBot",              // https://core.telegram.org/bots
	"UniversalFeedParser",      // https://code.google.com/p/feedparser/
	"WhatsApp",
	"bitlybot",            // http://bit.ly/
	"facebookexternalhit", // http://www.facebook.com/externalhit_uatext.php

	// !!! Availability monitoring bots !!! //
	// Those bots are the actual opposite of a threat. They check that resources are
	// available. They usually send a request at most every one minute.
	// On endpoint with lower traffic, it can happen that a very high percentage
	// of calls are health checks, which is useful to become aware of, but is usually
	// not a threat.
	"Amazon-Route53-Health-Check-Service", // https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-types.html
	"Better Uptime",                       // https://docs.digitalocean.com/products/marketplace/catalog/better-uptime/
	"Cloudflare-Healthchecks",             // https://www.cloudflare.com/
	"Elastic-Heartbeat",
	"EyeMonIT Uptime",
	"GoogleStackdriverMonitoring-UptimeChecks", // https://cloud.google.com/monitoring
	"HetrixTools Uptime",                       // https://hetrixtools.com/uptime-monitor/
	"Mozilla/DnsHealthCheckBot",                // https://support.mozilla.org/en-US/kb/firefox-dns-over-https
	"NodePing",                                 // https://nodeping.com/
	"Overtir-iWeb-Monitoring",                  // https://www.site24x7.com/
	"Pingdom.com",                              // http://www.pingdom.com/
	"Rightpoint Monitoring",                    // https://www.rightpoint.com/
	"SISTRIX Optimizer",                        // https://www.sistrix.com/faq/uptime
	"UptimeRobot",                              // http://www.uptimerobot.com/
	"Varnish/fastly",
	"WebyMon uptime", // https://www.webymon.com/index.html
	// Uptime-Kuma self hosted and open source.
	// I'm highlighting it because it's the only self managed open tool I've seen
	// in use so far.
	"Uptime-Kuma", // https://github.com/louislam/uptime-kuma

	// !!! Security scanner bots !!! //
	"Detectify", // https://detectify.com/bot/6c07c4e2-94e5-4719-8f74-57df06c6d54d
	"EdgeScan",  // https://www.edgescan.com/

	// !!! Other bots !!! //
	// This section include anything that's not in another category and can be a threat.
	// Reasons will widly vary based on the specific bot.
	"Algolia Crawler",     // https://www.algolia.com/doc/tools/crawler/getting-started/overview/
	"Alumio",              // https://www.alumio.com/
	"Barkrowler",          // https://babbar.tech/crawler
	"CCBot",               // https://commoncrawl.org/faq/
	"Celigo",              // https://www.celigo.com/
	"Devart ADO.NET",      // https://www.devart.com/dotconnect/
	"FeedBurner/",         // http://www.FeedBurner.com
	"Feedonomics",         // https://feedonomics.com/
	"Google-Apps-Script",  // https://script.google.com
	"Google-Firebase",     // https://firebase.google.com/
	"Grammarly",           // https://www.grammarly.com/plagiarism-checker
	"HesBot",              // https://hesbot.ir/
	"InternetMeasurement", // https://internet-measurement.com/
	"Linespider",          // https://lin.ee/4dwXkTH
	"Microsoft Data Access Internet Publishing Provider DAV",
	"Microsoft-WebDAV-MiniRedir",
	"NostoCrawlerBot",  // http://my.nosto.com/
	"RyteBota",         // https://bot.ryte.com/
	"Scrapy",           // https://scrapy.org
	"azure-logic-apps", // https://learn.microsoft.com/en-us/azure/logic-apps/logic-apps-overview
	"paloaltonetworks", // https://www.paloaltonetworks.com/

	// !!! Anonymous bots !!! //
	// Bots for which I couldn't find public documentation, and a public facing service.
	"fidget-spinner-bot",  // https://community.cloudflare.com/t/getting-tons-of-bad-bots-attack-recently/578502
	"my-tiny-bot",         // https://community.cloudflare.com/t/getting-tons-of-bad-bots-attack-recently/578502
	"test-bot",            // https://community.cloudflare.com/t/getting-tons-of-bad-bots-attack-recently/578502
	"thesis-research-bot", // https://community.cloudflare.com/t/getting-tons-of-bad-bots-attack-recently/578502
}

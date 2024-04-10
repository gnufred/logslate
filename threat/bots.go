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
	"adsbot-google",
	"adsbot-google-mobile",
	"apis-google",
	"feedfetcher-google",
	"google-adwords",
	"google-inspectiontool",
	"google-read-aloud",
	"google-safety",
	"google-site-verification",
	"googleassociationservice", // https://support.google.com/webmasters/answer/9419894
	"googlebot",
	"googlebot-image",
	"googlebot-news",
	"googlebot-video",
	"googleother",
	"googleproducer",
	"mediapartners-google",
	"storebot-google",

	// http://yandex.com/bots
	"yadirectfetcher",
	"yandexaccessibilitybot",
	"yandexadnet",
	"yandexblogs",
	"yandexbot",
	"yandexcalendar",
	"yandexdirect",
	"yandexdirectdyn",
	"yandexfavicons",
	"yandexfordomain",
	"yandeximageresizer",
	"yandeximages",
	"yandexmarket",
	"yandexmedia",
	"yandexmetrika",
	"yandexmobilebot",
	"yandexmobilescreenshotbot",
	"yandexnews",
	"yandexontodb",
	"yandexontodbapi",
	"yandexpagechecker",
	"yandexpartner",
	"yandexrca",
	"yandexscreenshotbot",
	"yandexsearchshop",
	"yandexsitelinks",
	"yandexspravbot",
	"yandextracker",
	"yandexturbo",
	"yandexverticals",
	"yandexvertis",
	"yandexvideo",
	"yandexvideoparser",
	"yandexwebmaste",

	// http://www.baidu.com/search/spider.html
	"baiduspider",
	"baiduspider-ads",
	"baiduspider-cpro",
	"baiduspider-favo",
	"baiduspider-image",
	"baiduspider-news",
	"baiduspider-render",
	"baiduspider-video",

	// http://www.bing.com/bingbot.htm
	"adidxbot",
	"bingbot",
	"bingpreview",
	"microsoftpreview", // https://aka.ms/microsoftpreview
	"msnbot",

	// other search engine indexers.
	"amazonbot",               // https://developer.amazon.com/support/amazonbot
	"applebot",                // https://support.apple.com/en-us/ht204683
	"duckduckbot",             // https://duckduckgo.com/duckduckbot
	"duckduckgo-favicons-bot", // https://duckduckgo.com/duckduckbot
	"femtosearchbot",          // http://femtosearch.com
	"gptbot",                  // https://platform.openai.com/docs/gptbot
	"graphconnectors",         // https://learn.microsoft.com/en-us/microsoftsearch/connectors-overview
	"megaindex.ru",            // http://megaindex.com/crawler
	"petalbot",                // https://webmaster.petalsearch.com/site/petalbot
	"pinterestbot",            // http://www.pinterest.com/bot.html
	"seekportbot",             // https://bot.seekport.com
	"seznambot",               // https://napoveda-seznam-cz.translate.goog/cz/seznamcz-web-search/?_x_tr_sl=auto&_x_tr_tl=en&_x_tr_hl=en&_x_tr_pto=wapp
	"sogou web spider",        // https://www.vntweb.co.uk/sogou-web-spider-web-robot/
	"yahoo!",                  // https://help.yahoo.com/kb/sln22600.html
	"yeti",                    // https://naver.me/spd
	"yisouspider",

	// !!! marketing bots !!! //
	// while most do not intend to be harmful, they can be too aggresive, badly
	// configured, crawl a site even if it's timing out. some are just generally
	// badly used and most often than not harmful.
	"addsearchbot",              // http://www.addsearch.com/bot/
	"ahrefsbot",                 // http://ahrefs.com/robot/
	"ahrefssiteaudit",           // http://ahrefs.com/robot/site-audit
	"awariobot",                 // https://awario.com/bots.html
	"awariorssbot",              // https://awario.com/bots.html
	"awariosmartbot",            // https://awario.com/bots.html
	"barkrowler",                // https://babbar.tech/crawler
	"blexbot",                   // http://webmeup-crawler.com/
	"brightedge crawler",        // https://www.brightedge.com/products/s3/contentiq
	"buck",                      // https://app.hypefactors.com/media-monitoring/about.html
	"bytespider",                // https://zhanzhang.toutiao.com/
	"caliperbot",                // http://www.conductor.com/caliperbot
	"dataforseobot",             // https://dataforseo.com/dataforseo-bot
	"domainstatsbot",            // https://domainstats.com/pages/our-bot
	"dotbot",                    // https://opensiteexplorer.org/dotbot
	"hubspot crawler",           // https://www.hubspot.com
	"klaviyo",                   // https://www.klaviyo.com/
	"konduto",                   // http://www.konduto.com
	"magpie-crawler",            // http://www.brandwatch.net
	"mailchimp",                 // https://mailchimp.com/
	"mj12bot",                   // http://mj12bot.com/
	"monsidobot",                // https://monsido.com/bot-html
	"optimizer",                 // https://www.sistrix.com/faq/what-is-the-optimizer-crawler-called/
	"productsup",                // http://productsup.io
	"rogerbot",                  // https: //moz.com/help/moz-procedures/crawlers/rogerbot
	"screaming frog seo spider", // https://www.screamingfrog.co.uk/seo-spider/
	"semrushbot",                // https://www.semrush.com/bot/
	"semrushbot-ba",             // https://www.semrush.com/bot/
	"semrushbot-coub",           // https://www.semrush.com/bot/
	"semrushbot-ct",             // https://www.semrush.com/bot/
	"semrushbot-si",             // https://www.semrush.com/bot/
	"semrushbot-swa",            // https://www.semrush.com/bot/
	"serpstatbot",               // https://serpstatbot.com
	"siteauditbot",              // http://www.semrush.com/bot.html
	"splitsignalbot",            // https://www.semrush.com/bot/
	// https://dynamics.microsoft.com/en-us/business-central/overview/
	"microsoft dynamics 365 business central",

	// !!! content preview/fetchers bots !!! //
	// content previewers shoud not be a threat, they visit a page when user share
	// links, usually to show content summary, and check for malware. also, they
	// usually cache results, therefor will usually not cause any issues.
	// there are some very weird edge cases where they can be harmful, so it's
	// important to detect them.
	"bitlybot",                 // http://bit.ly/
	"facebookexternalhit",      // http://www.facebook.com/externalhit_uatext.php
	"googledocs",               // http://docs.google.com
	"googleimageproxy",         // https://www.scientiamobile.com/what-is-google-image-proxy/
	"microsoft office",         //
	"skypeuripreview",          //
	"slackbot-linkexpanding",   // https://api.slack.com/robots
	"snap url preview service", // https://developers.snap.com/robots
	"snapchatads",              // https://businesshelp.snapchat.com/s/article/adsbot-crawler
	"telegrambot",              // https://core.telegram.org/bots
	"universalfeedparser",      // https://code.google.com/p/feedparser/
	"whatsapp",                 //

	// !!! availability monitoring bots !!! //
	// those bots are the actual opposite of a threat. they check that resources are
	// available. they usually send a request at most every one minute.
	// on endpoint with lower traffic, it can happen that a very high percentage
	// of calls are health checks, which is useful to become aware of, but is usually
	// not a threat.
	"amazon-route53-health-check-service",      // https://docs.aws.amazon.com/route53/latest/developerguide/health-checks-types.html
	"better uptime",                            // https://docs.digitalocean.com/products/marketplace/catalog/better-uptime/
	"cloudflare-healthchecks",                  // https://www.cloudflare.com/
	"elastic-heartbeat",                        //
	"eyemonit uptime",                          //
	"googlestackdrivermonitoring-uptimechecks", // https://cloud.google.com/monitoring
	"hetrixtools uptime",                       // https://hetrixtools.com/uptime-monitor/
	"monibot",                                  // https://monibot.io/
	"mozilla/dnshealthcheckbot",                // https://support.mozilla.org/en-us/kb/firefox-dns-over-https
	"nodeping",                                 // https://nodeping.com/
	"overtir-iweb-monitoring",                  // https://www.site24x7.com/
	"pingdom.com",                              // http://www.pingdom.com/
	"rightpoint monitoring",                    // https://www.rightpoint.com/
	"sistrix optimizer",                        // https://www.sistrix.com/faq/uptime
	"uptimerobot",                              // http://www.uptimerobot.com/
	"varnish/fastly",
	"webymon uptime", // https://www.webymon.com/index.html
	// uptime-kuma self hosted and open source.
	// i'm highlighting it because it's the only self managed open tool i've seen
	// in use so far.
	"uptime-kuma", // https://github.com/louislam/uptime-kuma

	// !!! security scanner bots !!! //
	"censysinspect",         // https://about.censys.io/
	"detectify",             // https://detectify.com/bot/6c07c4e2-94e5-4719-8f74-57df06c6d54d
	"edgescan",              // https://www.edgescan.com/
	"whitehat pink panther", // https://www.synopsys.com/software-integrity/security-testing.html

	// !!! other bots !!! //
	// this section include anything that's not in another category and can be a threat.
	// reasons will widly vary based on the specific bot.
	"algolia crawler",     // https://www.algolia.com/doc/tools/crawler/getting-started/overview/
	"alumio",              // https://www.alumio.com/
	"azure-logic-apps",    // https://learn.microsoft.com/en-us/azure/logic-apps/logic-apps-overview
	"barkrowler",          // https://babbar.tech/crawler
	"bazaarvoice",         // https://knowledge.bazaarvoice.com/wp-content/conversations/en_US/Collect/DCC.html
	"ccbot",               // https://commoncrawl.org/faq/
	"celigo",              // https://www.celigo.com/
	"chatgpt-user",        // https://platform.openai.com/docs/plugins/bot
	"cincraw",             // https://cincrawdata.net/
	"claudebot",           // https://github.com/claudebot
	"devart ado.net",      // https://www.devart.com/dotconnect/
	"feedburner",          // http://www.feedburner.com
	"feedonomics",         // https://feedonomics.com/
	"google-apps-script",  // https://script.google.com
	"google-firebase",     // https://firebase.google.com/
	"grammarly",           // https://www.grammarly.com/plagiarism-checker
	"hesbot",              // https://hesbot.ir/
	"internetmeasurement", // https://internet-measurement.com/
	"linespider",          // https://lin.ee/4dwxkth
	"microsoft data access internet publishing provider dav",
	"microsoft-webdav-miniredir",
	"nostocrawlerbot",   // http://my.nosto.com/
	"openapi-generator", // https://github.com/OpenAPITools/openapi-generator
	"paloaltonetworks",  // https://www.paloaltonetworks.com/
	"rytebota",          // https://bot.ryte.com/
	"scrapy",            // https://scrapy.org
	"turnitin",          // https://www.turnitin.com/robot/crawlerinfo.html
	"youbot",            // http://www.you.com/

	// !!! Anonymous bots !!! //
	// Bots for which I couldn't find public documentation, and a public facing service.
	"fidget-spinner-bot",  // https://community.cloudflare.com/t/getting-tons-of-bad-bots-attack-recently/578502
	"my-tiny-bot",         // https://community.cloudflare.com/t/getting-tons-of-bad-bots-attack-recently/578502
	"test-bot",            // https://community.cloudflare.com/t/getting-tons-of-bad-bots-attack-recently/578502
	"thesis-research-bot", // https://community.cloudflare.com/t/getting-tons-of-bad-bots-attack-recently/578502
}

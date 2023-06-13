# Development roadmap

## v0.X - prototype

### v0.1 (done)

Analyze parsed data and anontate entries with aditionnal information that matches the first detected items.
Report detection in format: json, json.gz, csv, echo.

Attack detection:
* SQL injection
* Command injection
* Directory traversal

Detection method:
* Fake user agent
* URL pattern

### v0.2

Bot and crawler detection:
* [Search] Google indexer
* [Search] Bing indexer
* [Search] Yandex indexer
* [Marketing] Semrush bot
* [Marketing] Moz.com SEO crawler
* [Marketing] "Caliperbot/1.0" http://www.conductor.com/caliperbot
* [Marketing] "Mozilla/5.0 (compatible; AhrefsBot/7.0; +http://ahrefs.com/robot/)" https://ahrefs.com/robot/

Detection method:
* User agent pattern
* Known IP verification

### v0.3

Ad detection:
* Google AdWords campaign
* Facebook ad campaign

### v0.4

Report format:
* postfix
* sendmail

### v0.5

Bot and crawler detection:
* Screaming Frog SEO spider
* Top 20 search indexers

### v0.6

Attack detection:
* Top 10 attack methods

### v0.7

Ad campaign detection:
* Top 20 ad vendor

### v0.8

Exploit detection:
* Magecart CVEs (for Magento CMS)

### v0.9

Exploit detection: WordPress related CVEs
* [EvalPHP](https://blog.sucuri.net/2023/04/massive-abuse-of-abandoned-evalphp-wordpress-plugin.html)

### v0.10

Exploit detection:
* Drupal related CVEs

### v0.11

Exploit detection:
* Symfony

### v0.12

Add support for Apache access.log files

### v0.13

Split the debug/ package to it's own repository and name it something else to avoid conflics with runtime/debug. The goal here is to reuse with with the upcoming parser/nginx rewrite, and on my jobs projects.

The API should be massively simple and manage everything almost everything on it's own. As in, it needs to store all it's runtime setting at package level and share those transparently accross API calls.

Public API should be like:

```
SetWriter(*io.Writer) error // configure the output stream
SetVerbosity(verb int) error // set verbosity levels [0-3], error on anything else
GetVerbosity() int // return current verbosity
Log(verb int, name string, msg interface{}) error // output msg to io.Writer
```

### v0.15 and beyond

Further updates will add small sub features, improvements, and bug fixes in order to move to official v1.0.0 release

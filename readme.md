# Log "translate"

The log translate (logslate) project aims is to create a dead simple tool that will help web developers and system administrators to secure their customers websites. This will be accomplished by periodicaly reading the access.log file from popular web servers and by creating an actionable report.

## Reporting

Execution will create a timestamped .json.gz report on disk, that will then be used to generate other forms of reports.

## Commands

```
logslate [parse,analyze,report]
--logs ${path}/access*.log(.gz)
--db ${path}/parsed-log.ldb
--start "timestamp"
--end "timestamp"
--period "duration"
--detect [attack,exploit,bots,ad]
--report_format [json,csv,echo,email]
--report_folder ${folder}
```

## Detection goal

The goal is to be able to detect and report all of the following events:

Generic attack like:
* SQL injection.
* Command injection.
* Directory traversal attacks.

Bots and cralwers:
* Search engine indexers.
* SEO analysis tools.

Ongoing online advertising campaigns from popular networks like Google AdWords, Facebook, etc.

CMS and Framework known vulnerability exploitations.

### Attacks and scanning

In the vast majority of cases, attacks and scanners directed at website are unwanted. Furthermore, they will very likely result in outages, bugs, performance issues, and log spam. As such, it's primordial to detect and blog them.

Multiple attack vectors are easy to identify by reading access logs and are usually pretty obvious.

Detection techniques:
* URL pattern matching.
* Fake user agent detection.
* Abusive IP reporting platforms.

### Bots and crawler

While most bots and web crawlers are usually friendly, they can often cause you site to load thousands of uncached pages causing performance issues or straight outages. As such, being aware of their present is immensly helpful when dealing with such problems.

Detection techniques:
* User agent matching.
* IP verification.

All legitimate bots and crawlers have clearly identifiable user agents. On top, some vendors publish IP ranges used. Therefor, they are quite easy to identify.

### Advertising campaigns

Online advertisting compaigns usually cause a lot of legitimate traffic. Usually, those are run by marketing firms that may not prepare the infrastructure properly, or warn developers and system administrator before hand. Furthermore, campaigns are very expensive. As such, it's of prime importance to be able to detect ongoing campaing in order to react swiftly if issues arise.

The vast majority of advertisting platform add URL query parametters to identify and track marketing campaings, making ad related traffic easy to detect.

Detection techniques:
* URL pattern matching

### Vulnerability exploitation

Popular CMSes and Frameworks like WordPress and Drupal are actively exploited at all time. While detecting exploitation activities is usually somewhat complex, it's possible to identify some via access.log.

CMS and framework:
* Magento
* WordPress
* Drupal
* Symfony

## Development roadmap

### v0.1

Read Nginx access.log from a file (text or Gzip) or from terminal pipe and parse it into [LevelDB](https://github.com/syndtr/goleveldb).

Analyze LDB parsed data and anontate entries with aditionnal information based on matches.

Report format: json, json.gz, csv, echo

Attack detection:
* SQL injection
* Command injection
* Directory traversal

Detection method:
* Fake user agent
* URL pattern

### v0.2

Bot and crawler detection:
* Google indexer
* Bing indexer
* Yandex indexer
* Semrush bot
* Moz.com SEO crawler

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

Exploit detection:
* WordPress related CVEs

### v0.10

Exploit detection:
* Drupal related CVEs

### v0.11

Exploit detection:
* Symfony

### v0.12

Add support for Apache access.log files

### v0.13

Performance updates

### v0.14 and beyond

Further updates will add small sub features, improvements, and bug fixes in order to move to official v1.0.0 release

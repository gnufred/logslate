# Log "translate"

Log translate (logslate) is a Free and Open Source **defensive security** tool. Its main purpose is to help web developers and system administrators secure their customers websites.

## NGINX parsing package

If you need a GoLang package to parse NGINX access.log, use [gnufred/nxparse](https://gitlab.com/gnufred/nxparse).

## Project info

### In development

See the [development roadmap](.docs/roadmap.md) for the current status.

### License

This project is licended under [GPL v3.0](license.md)

## Installation

Install with `curl`, `brew`, or manual downloads.

See **[the installation documentation](.docs/install.md)**.

## Usage

See the **[comprehensive usage guide](.docs/usage.md)** and the **[database usage guide](.docs/database.md)** for detailed use cases.

```
cat access.log | logslate
-d, --db ${path}/parsed-log.ldb
-s, --since "timestamp"
-u, --until "timestamp"
-f, --log_format "NGINX log_format string"
-r, --report "file.[txt,json]|screen"
```

## Detection

Logslate is able to detect and report of the following events.

### Attacks and scans

Attacks, and scans are usually unwanted. They results in outages, bugs, performance issues, and log spam. Legitimate security scanners will very often cause the same problems.

Categories:
* SQL injection.
* Command injection.
* Directory traversal attacks.
* Cross scripting attacks.
* Bruteforce scanning (basic).

### Bots and crawlers

While most bots and web crawlers are usually friendly, they can cause performance issues, and outages. Being aware of their presence is always helpful.

All legitimate bots and crawlers have clearly identifiable user agents. Some vendors publish IP ranges used. They are quite easy to identify.

### Advertising campaigns

Advertisting campaigns cause a lot of traffic. Usually, they are manage by teams that may not be able to prepare the infrastructure properly, or warn system administrator before hand. Traffic spikes cause performance issues, and outages.

All advertisting platform add the `utm_source` URL query parametters to identify themselves making it trivial to detect.

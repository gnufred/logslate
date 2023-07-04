# Development roadmap

## v0.4

Ads and marketing campaign detection.

## v0.5

Detect more attack types than: dir traversal, sql injection, command injection, XSS.

## v0.6

IP analysis:

* identify IP from CDNs
* analyse IP ranges
* gather more intel about threat IPs

The IP analyser will be a new repo, with a GoLang tool that will run in a GitLab pipeline daily. The tool will download known IP from various official source and create a CSV file with the IP range, owner, and time stamp.

The CSV will then be used to generate a known_ips.go which logslate will use in a weekly automated build.

Logslate will include a new version number 'w01' which will be the week of the year. Every Sunday, a GitLab pipeline will create a '0.0.0-w00' weekly release, using the latest stable version and the latest known IPs.

v0.5.1 will include known IPs for Fastly, Cloudflare, and Cloudfront.
v0.5.x will add more CDNs, and known active bad actors.

## v0.7

User agent intel:

Review what's useful to gather about user agent, like fake agents, and so on.

## v0.9

PHP app/framework exploit:

* v0.9.0: Magento (specifically MageCart attacks)
* v0.9.1: WordPress
* v0.9.2: Drupal
* v0.9.3: Symfony

## v0.9

GraphQL endpoint attack detection based on Cloudflare's post: https://blog.cloudflare.com/protecting-graphql-apis-from-malicious-queries/

### v0.10

Add support for Apache access.log files

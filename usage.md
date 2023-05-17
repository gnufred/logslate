# Usage

Using `logslate` consist of the following steps:

* Parse access.log files into an SQLite3 database
* Query the `sqlite3` database 
* (upcoming) Reviewing the detection report 

The current version only support **NGINX default log format**. Version 0.14 and later will support other formats. 

## Installation

If you want to use `logslate` on your local workstation, it's recommended to use the `brew` package manager.

```
# install
brew tap gnufred/logslate https://gitlab.com/gnufred/logslate
brew install logslate
# update
brew update && brew upgrade logslate
```

If you're using `logslate` on a server, or need to install it in a built container, it's recommended to use `curl`.

```
curl --location --output ./logslate https://gitlab.com/api/v4/projects/gnufred%2Flogslate/packages/generic/logslate/latest/logslate-linux-amd64
```

For **Windows** and other uses cases, head to the [release page](https://gitlab.com/gnufred/logslate/-/releases) and download the binary manually.

## Parsing access log files

Logslate accept input **only from stdin**. You have to feed it with pipes. 

The simpliest case is to parse one file `cat access.log | logslate --db log.db`, but you can parse multiple files `cat *.log | logslate --db all.db`. If your files are Gziped `zcat access*.log.gz | logslate --db all.db`.

If you have logs in multiple location you can `find /var/log/sites/ -type f -name 'access.log' | xargs cat | logslate --db all.db`.

You can also, get logs directly from your server `ssh webserver 'find /var/log/sites/ -type f -name 'access.log' | xargs cat' | logslate --db all.db`.

## Using time stamps

You can use the `--since` and `--until` arguments to parse only the wanted lines. You can use any time format that's listed in [araddon/dateparse](https://github.com/araddon/dateparse#extended-example) **or better** [systemd.time](https://www.man7.org/linux/man-pages/man7/systemd.time.7.html) format.

The *three* main ways you'll use this is:

* copy and paste from NGINX logs `--since '07/Jan/2023:13:03:23 +0000' --until '14/Jan/2023:08:21:12 +0000'`.
* with systemd.time ` --since '2 days ago' --until '1 day ago'`.
* from a cronjob `--since '-24 hours'`.

## Parsing errors

Unparsable log entries are ignored. You can output them with `-v`.

It's safe to ignore them if there's just a few.

## Debug

Activate debug output with `-v, -vv, -vvv` or environment variable `export LS_DEBUG=[0,1,2,3]`.

Debug mode has 4 levels: OFF (0), V (1), VV (2), VVV (3).

## Using tmpfs

If your disk is slow or you are not permitted to save the data on your workstation for compliance reasons, you can save the database in tmpfs.

For example, on most Linux systems you'd use `--db /run/user/1000/log.db`.

To follow your **compliance or security** framework, you can use **both** ssh and a tmpfs drive `ssh server 'cat /var/log/www/example.com/access.log' | logslate --db /run/user/1000/example.db` to avoid saving information on permanent storage.

## Accessing the database

To use the database file, you need to install the [SQLite3 CLI](https://www.sqlite.org/cli.html). Then it's just a matter of `sqlite3 log.db`.

That being said, **I highly recommend you issue queries from your shell** instead `sqlite3 log.db 'select * from logs;'`.

## SQL queries

The database schema is in [db.go#L56](https://gitlab.com/gnufred/logslate/-/blob/main/db/db.go#L56).

List and count top 10 IPs

```
select count(ip), ip from logs group by ip order by count(ip) desc limit 10;
```

List all request from an IP, order by URL path and query string

```
select ip, path, query from logs where ip = '208.115.199.20' order by path, query;
```

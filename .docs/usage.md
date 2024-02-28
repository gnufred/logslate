# Usage

Usage usually consist of the following steps:

* Parse an access.log files into an SQLite3 database
* Create a detection report 

## Parsing access log files

Logslate accept input **only from stdin**. You have to feed it with pipes. 

The simplest case is to parse one file `cat access.log | logslate`, but you can parse multiple files `cat *.log | logslate`. If your files are Gziped `zcat access*.log.gz | logslate`.

If you have logs in multiple locations you can `find /var/log/sites/ -type f -name 'access.log' | xargs cat | logslate`.

You can also, get logs directly from your server `ssh webserver 'find /var/log/sites/ -type f -name 'access.log' | xargs cat' | logslate`.

## Database

It's possible to parse log entries into an SQLite3 database with `--db file.db`.

Read the [database usage guide](database.md) for instruction on how to query the database.

## Reports

There's no report by default. In most cases, you'll want to use `--report screen` to print the report on the screen after parsing completes.

Otherwise, write the report to disk with `--report file.log` or `--report file.json` to get a text or json formatted output.

### Custom report layouts

It's possible to specify your own screen and text report layouts with `--layout section:option_value,section2:option1_value:option2_value,etc`.

Possible values are:

```
summary
time
scanned
threats
attacks
attackers
  order_name (default)
  order_count
  min_X (int)
ads
  order_name (default)
  order_count
  min_X (int)
bots
  order_name (default)
  order_count
  min_X (int)
```

## NGINX `log_format`

The default value for `--log_format` is [NGINX default combined format](https://nginx.org/en/docs/http/ngx_http_log_module.html).
```
log_format combined '$remote_addr - $remote_user [$time_local] '
                    '"$request" $status $body_bytes_sent '
                    '"$http_referer" "$http_user_agent"';
```

You don't have to pass this to `--log_format`, but would work and look like this
```
--log_format '$remote_addr - $remote_user [$time_local] "$request" $status $body_bytes_sent "$http_referer" "$http_user_agent"'
```

If **you use different format**, you can copy/paste your format to the `--format_log` argument.

For example, if you used `time_iso8601` instead of `time_local` and `bytes_sent` instead of `body_bytes_sent`, your NGINX config would be
```
log_format combiso8 '$remote_addr - $remote_user [$time_iso8601] '
                    '"$request" $status $bytes_sent '
                    '"$http_referer" "$http_user_agent"';
```

and `--log_format`
```
--log_format '$remote_addr - $remote_user [$time_iso8601] "$request" $status $bytes_sent "$http_referer" "$http_user_agent"'
```

## Threats log

It's possible to generate a threat log with the `--threats` argument. The possible values are `screen`, `file.log`, or `file.json`.

For each parsed NGINX access.log line, if a threat is detected, it will be logged to the selected output.

## Using time stamps

You can use the `--since` and `--until` arguments to parse only the wanted lines. You can use any time format that's listed in [araddon/dateparse](https://github.com/araddon/dateparse#extended-example) or [systemd.time](https://www.man7.org/linux/man-pages/man7/systemd.time.7.html) format.

The **three** main ways you'll use this is:

* copy and paste from NGINX logs `--since '07/Jan/2023:13:03:23 +0000' --until '14/Jan/2023:08:21:12 +0000'`.
* with systemd.time ` --since '2 days ago' --until '1 day ago'`.
* from a cronjob `--since '-24 hours'`.

You can also use those English words: `now`, `today`, `yesterday`.

If you need **specific timezones**, use the calculated value `+XXXX` or shortcodes. For example `2006-01-02 15:04:05 MST`, `2006-01-02 UTC`, and `2006-01-02 -0700`.

## Parsing errors

Unparsable log entries are ignored. You can output them with `--vvv`.

It's usually safe to ignore them.

## Using tmpfs

If your disk is slow or you are not permitted to save the data on your workstation for compliance reasons, you can save the database in tmpfs.

For example, on most Linux systems you'd use `--db /run/user/1000/log.db`.

To follow your **compliance or security** framework, you can use **both** ssh and a tmpfs drive `ssh server 'cat /var/log/www/example.com/access.log' | logslate --db /run/user/1000/example.db` to avoid saving information on permanent storage.

## Debuging

Activate debug output with `-v[vv]`, `--debug=[1-3]`, or environment variable `export LS_DEBUG=[0-3]`.

Debug output to `stderr` meaning you can redirect it like so `cat access.log | logslate -v 2> debug.log`

## Profiling

You can enable `runtime/pprof` CPU and memory profiling with `--profcpu` and `--profmem`.

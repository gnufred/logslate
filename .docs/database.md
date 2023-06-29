# Database

## Requirements

You need to install the SQLite3 CLI in order to access database content.

The prefered installation method is to install via your Linux distribution package manager. Most distributions have packages for SQLite3.

Otherwise, head to [SQLite3 CLI](https://www.sqlite.org/cli.html).

## Using the database

Access the database directly: `sqlite3 <DATABASE>`.

Query the dataset from "outside": `sqlite3 <DATABASE> 'select * from logs;'`.

## SQL queries

Print the database schema
```
.schema logs
```

Enhanced logs
```
select ip, time, status, method, bot, attack, path, query, referer, agent from logs order by time
```

Bot logs
```
select bot, ip, time, status, method, path, query, referer, agent from logs where bot != "" order by time
```

Entries flagged as an attack
```
select attack, ip, time, status, method, path, query, referer from logs where attack != "" order by time
```

Ordered count of IPs
```
select count(ip), ip from logs group by ip order by count(ip) desc
```

Ordered count of User Agent strings
```
select count(agent), agent from logs group by agent order by count(agent) desc
```

Ordered count of bots
```
select count(bot), bot from logs where bot != "" group by bot order by count(bot) desc
```

Ordered count of attacking IPs
```
select count(ip), ip from logs where attack != "" group by ip order by count(ip) desc
```

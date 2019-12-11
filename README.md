# Khazen - A web service to work with SQL database server

Khazen is a web service that manages SQL databases such as MySQL.
Khazen (خازن in Persian) means Treasury Guardian.

# Contents
* [Motivations](https://github.com/SakkuCloud/khazen#motivations)
* [How to use](https://github.com/SakkuCloud/khazen#how-to-use)
* [Configuration](https://github.com/SakkuCloud/khazen#configuration)
* [Endpoints](https://github.com/SakkuCloud/khazen#endpoints)
  * [Create MySQL account](https://github.com/SakkuCloud/khazen#mysql-create-account)
  * [Create MySQL database](https://github.com/SakkuCloud/khazen#mysql-create-database)
  * [Delete MySQL database](https://github.com/SakkuCloud/khazen#mysql-delete-database)
  * [Execute MySQL bundle](https://github.com/SakkuCloud/khazen#mysql-execute-bundle)
  * [Import MySQL database](https://github.com/SakkuCloud/khazen#mysql-import-database)
  * [Health](https://github.com/SakkuCloud/khazen#health)
* [To do](https://github.com/SakkuCloud/khazen#to-do)

# Motivations
In SAKKU team we have several modules that need to make root privileged actions in database. So we need a manager to listen in endpoints and make these actions in database host to prevent issues like:
- Using ROOT username and password in every module.
- ROOT user can login out of database localhost.
- No control in actions that every module can do with ROOT user.

# How to use
Building from source:
```sh
$ go build -o /usr/bin/khazen github.com/SakkuCloud/khazen
```

Running with CLI (stay foreground):
```sh
$ khazen
```

Running in debug mode:
```sh
$ khazen -debug=true
```

Khazen uses some options. These options are listed below.

key   | name        | default                |
----- | ----------- | ---------------------- |
debug | Debug Mode  | false                  |
| c     | Config File | /etc/khazen/config.yml |

For production, it's better to use a Systemd service to run Khazen.
A simple Systemd service shown below. Save this in `/lib/systemd/system/khazen.service` 
> ```sh
> [Unit]
> Description=KHAZEN - A webservice to work with SQL
> After=network.target
>
> [Service]
> Type=simple
> Restart=on-failure
> TimeoutStopSec=10
> RestartSec=5
> ExecStart=/usr/bin/khazen
>
> [Install]
> WantedBy=multi-user.target
>```

Run and enable service:
```sh
$ systemctl enable khazen
$ systemctl start khazen
```

# Configuration
Khazen uses both YAML format and OS Environment for config. You can see [config.yml.example](https://github.com/SakkuCloud/khazen/blob/master/config.yml.example) for a sample config file.
You can pass config file with:
```sh
khazen -c config.yml
```
Below table describes available config file.

| config         | env                   | required | default             | describe |
| ---------------| --------------------- | :------: | ------------------- | ------------------------------------------------------- |
| port           | KHAZEN_PORT           | NO       | 3000                | server will run on this port                           |
| logfile        | KHAZEN_LOGFILE        | NO       | /var/log/khazen.log | logs will store in this file                           |
| sentrydsn      | KHAZEN_SENTRYDSN      | NO       |                     | DSN of Sentry                                          |
| accesskey      | KHAZEN_ACCESSKEY      | YES      |                     | value of service http header to authorize requests     |
| secretkey      | KHAZEN_SECRETKEY      | YES      |                     | value of service-key http header to authorize requests |
| mysqlcmd       | KHAZEN_MYSQLCMD       | NO       | mysql               | mysql cmd in host. some endpoints do not use go mysql driver |
| mysql host     | KHAZEN_MYSQL_HOST     | NO       | 127.0.0.1           | MySQL database server address                          |
| mysql user     | KHAZEN_MYSQL_USER     | NO       | root                | MySQL database server user                             |
| mysql password | KHAZEN_MYSQL_PASSWORD | YES      |                     | MySQL database server password                         |
| mysql port     | KHAZEN_MYSQL_PORT     | NO       | 3306                | MySQL database server port                             |

# Endpoints
### MySQL create account
Creates account in MySQL database server. A complete curl requests shown below. All json attributes except *native_password* are required.
```sh
curl -X POST \
 https://khazen.sakku.cloud/api/mysql/account \
 -H 'Content-Type: application/json' \
 -H 'service: my-awesome-accesss' \
 -H 'service-key: Super$3crT' \
 -d '{
"username":"new_user",
"password":"some_strong_pass",
"max_queries_per_hour":"100000",
"max_updates_per_hour":"100000",
"max_connections_per_hour":"2000000",
"max_user_connections":"2000000",
"native_password":true
}
'
```

### MySQL Create database
Creates database in MySQL database server and set full privilege for user on this database. A complete curl requests shown below. All json attributes except *character_set* are required.
```sh
curl -X POST \
 https://khazen.sakku.cloud/api/mysql/database \
 -H 'Content-Type: application/json' \
 -H 'service: my-awesome-accesss' \
 -H 'service-key: Super$3crT' \
 -d '{
"username":"test_user_for_new_method_3",
"database":"db_test_new_method_3",
"character_set":"utf8"
}
'
```

### MySQL Delete database
Deletes database in MySQL database server. A complete curl requests shown below.
```sh
curl -X DELETE \
 https://khazen.sakku.cloud/api/mysql/database/<name-of-database> \
 -H 'service: my-awesome-accesss' \
 -H 'service-key: Super$3crT' \
```

### MySQL Execute bundle
This bundle first creates account in MySQL database server then creates database in MySQL database server and set full privilege for user on this database. A complete curl requests shown below. All json attributes except *native_password* and *character_set* are required.
```sh
curl -X POST \
 https://khazen.sakku.cloud/api/mysql/bundle \
 -H 'Content-Type: application/json' \
 -H 'service: my-awesome-accesss' \
 -H 'service-key: Super$3crT' \
 -d '{
"account" : {
"username":"test_user_for_new_method_3",
"password":"sdsdf_234234mn_234r_3",
"max_queries_per_hour":"100000",
"max_updates_per_hour":"100000",
"max_connections_per_hour":"2000000",
"max_user_connections":"2000000",
"native_password":true
},
"database":{
"username":"test_user_for_new_method_3",
"database":"db_test_new_method_3",
"character_set":"utf8"
}
}
'
```

### MySQL Import database
Imports database from file in MySQL database server. A complete curl requests shown below.
**NOTE:** This endpoint use mysql cmd instead of go mysql driver.
```sh
curl -X POST \
  https://khazen.sakku.cloud/api/mysql/import/<name-of-database> \
  -H 'Content-Type: multipart/form-data' \
  -H 'service: my-awesome-accesss' \
  -H 'service-key:  Super$3crT' \
  -F import_file=@<address/of/file>
```

### Health
Get Server health (status, uptime and server time).
```sh
curl -X GET \
 https://khazen.sakku.cloud/api/health 
```

# To do
* Update account attributes
* Delete account
* Delete database
* Set privilege endpoint
* Add import/export

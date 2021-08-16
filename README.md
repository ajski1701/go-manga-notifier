# go-manga-notifier

## About
Manga Notifier is a personal pet project written in order to learn Golang. All queries for manga go against the API of MangaDex which then are parsed and if applicable sent as an alert to the configured destination address.

## Configuration
All configuration is handled via an ini file. If not detected a template ini will be created but the format can be found below.

```ini
[email]
from = <from_gmail_email_address>
password = <gmail_password>
to = <comma_delimited_recepient_emails>

[mangadex]
username = <mangadex_username>
password = <mangadex_password>
```
# Office

![Hackatime Time](https://hackatime.hackclub.com/api/v1/badge/U082SDZLULQ/theluqmn/office)

The Office is my webserver that acts as my office in the internet. It serves the
following functions:

1. Serves my portfolio website
2. Compiles and serves my journals
3. Redirects from my [URL shorteners](#url-shorteners)

## Technicality

- **Frontend**: Built using HTML, styled using
[Tailwind CSS](https://tailwindcss.com/), and front-end processing using JavaScript.
- **Backend**: Written primarily with [Go](https://go.dev/), using the
[Echo](https://echo.labstack.com/) library, and [SQLite](https://sqlite.org/)
as the database.
- **Deployment**: [HackClub Nest](https://hackclub.app), 2 cores 2GB container.

Additional libraries:

- [Air](https://github.com/air-verse/air): Live reload

### URL Shorteners

Below is a table of the URL shorteners I added to the Office:

|Website|Shortener|
|---|---|
|GitHub|/gh|
|X (Twitter)|/x|
|Instagram|/ig|
|YouTube|/yt|
|Discord|/dc|
|Hackatime|/ht|

## Notes

Licensed under the MIT License

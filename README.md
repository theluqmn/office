# Office

![Hackatime Time](https://hackatime.hackclub.com/api/v1/badge/U082SDZLULQ/theluqmn/office)
| [Website](https://theluqmn.hackclub.app)

The Office is my webserver that acts as my office in the internet. It serves the
following functions:

1. Serves my [portfolio website](https://theluqmn.hackclub.app)
2. Compiles and serves my [journals](https://theluqmn.hackclub.app/journals)
3. Redirects from my [URL shorteners](#url-shorteners)

## Technicality

- **Frontend**: Built using HTML, styled using
[Tailwind CSS](https://tailwindcss.com/), and front-end processing using JavaScript.
- **Backend**: Written primarily with [Go](https://go.dev/), using the
[Echo](https://echo.labstack.com/) library. Also contains [Goldmark](https://github.com/yuin/goldmark)
for parsing Markdown files.
- **Deployment**: [HackClub Nest](https://hackclub.app), 2 cores 2GB container.

Additional libraries:

- [Air](https://github.com/air-verse/air): Live reload

### html/template

- All pages served by the Office utilises Go's `html/template`, with reusable
components like headers and footers in the `/views/components` directory.
- Each page is its own file (journal entries follows `journal-entry` as base)
and imports the reusable components.
- Everything is then rendered into HTML and served to the client.

### Journals

The journal entries are in a separate [journal](https://github.com/theluqmn/journal)
repository. The office fetches the journal entries from the journal directory, then
parses the markdown files into HTML using [Goldmark](https://github.com/yuin/goldmark).

### URL Shorteners

Below is a table of the URL shorteners I added to the Office:

|Website|Shortener|
|---|---|
|GitHub|[/gh](https://theluqmn.hackclub.app/gh)|
|X (Twitter)|[/x](https://theluqmn.hackclub.app/x)|
|Instagram|[/ig](https://theluqmn.hackclub.app/ig)|
|YouTube|[/yt](https://theluqmn.hackclub.app/yt)|
|Discord|[/dc](https://theluqmn.hackclub.app/dc)|
|Hackatime|[/ht](https://theluqmn.hackclub.app/ht)|

## Notes

This is my most advanced web project, and my first deployed outside of GitHub
Pages. The week spent putting everything together has been meaningful
and refreshing.

Licensed under the MIT License

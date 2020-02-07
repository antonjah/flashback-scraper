# flashback-scraper
A Go tool that scrapes flashback threads between date intervals

## Example config.yaml
```yaml
startDate: "20190101"
endDate: "20200101"
format: "json"
URLs:
  - "https://www.flashback.org/t3117396"
  - "https://www.flashback.org/t4218386"
```

## Todo
* ~~Date conversion from `igår xx:xx` to date~~
* ~~Tidy output~~
* ~~Output to file~~
* Handle multi-page threads
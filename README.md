# LifelogSP Hugo Generator (hugo_lspd)

Create Hugo website from lifelogging/quantified self data related to [LifelogSP](https://github.com/spech66/lifelogsp).

## Features

* **Weight:** Data from [LifelogSP](https://github.com/spech66/lifelogspd)
* **Journal:** Data from [LifelogSP](https://github.com/spech66/lifelogspd)
* **Strength training:** Data from [LifelogSP](https://github.com/spech66/lifelogspd)
* **Endurance workout:** Data from [LifelogSP](https://github.com/spech66/lifelogspd)
* **Mood and activities**: [Daylio - Mood Tracker and Micro Diary](https://daylio.webflow.io/)
* **Habits**: [Loop Habit Tracker](https://play.google.com/store/apps/details?id=org.isoron.uhabits)
* **Tags/Categories**: Automatically created using the data sources above
* **Theme**: [Bootstrap-BP hugo theme](https://github.com/spech66/bootstrap-bp-hugo-theme)

### Improvements

* Parse the Habit data dynamically

## Screenshots

![Start](https://raw.githubusercontent.com/spech66/hugo_lspd/master/_screenshots/s_001_start.png "Start")
![Page](https://raw.githubusercontent.com/spech66/hugo_lspd/master/_screenshots/s_002_page.png "Page")
![Categories](https://raw.githubusercontent.com/spech66/hugo_lspd/master/_screenshots/s_003_categories.png "Categories")
![Tags](https://raw.githubusercontent.com/spech66/hugo_lspd/master/_screenshots/s_004_tags.png "Tags")

[More screenshots](https://github.com/spech66/hugo_lspd/tree/master/_screenshots)

## Build and run from source

Make sure you have the [Go Programming Language](https://golang.org/) tools set up an ready.

### Linux

Checkout the code to your `GOPATH` directory, build and run it.

```bash
go get github.com/spech66/hugo_lspd
cd $GOPATH/src/github.com/spech66/hugo_lspd
go build
./hugo_lspd -config example.config.json
```

### Windows

Checkout the code to your `GOPATH` directory, build and run it.

```cmd
go get github.com/spech66/hugo_lspd
cd %GOPATH%\src\github.com\spech66\hugo_lspd
go build
hugo_lspd.exe -config example.config.json
```

## Generate the website

After you executed `hugo_lspd` the pages are created in the `hugo` subfolder. You can generate the website using the hugo command line tool now.

Make sure you have the [Hugo](https://gohugo.io/) tools set up an ready. (For windows the easiest way is using [Chocolatey](https://chocolatey.org/))

### Linux

```cmd
cd $GOPATH/src/github.com/spech66/hugo_lspd/hugo
hugo server
```

### Windows

```cmd
cd %GOPATH%\src\github.com\spech66\hugo_lspd\hugo
hugo server
```

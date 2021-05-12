# Azure OG

I stumbled across [DevsQuotesPrinter](https://github.com/JosemyDuarte/DevsQuotesPrinter) and loved the idea so much I had to build my own.

I modified the code to accept the title and font size as query parameters, with defaults if they aren't sent.

The app has a command-line interface as well as an Azure Function. The cli tool accepts parameters to control the title and font size. The Azure Function accepts the same parameters as a query string.


## CLI

```
make build-cli
make run TITLE="I'm so social"
```

Details :

```
./dist/azureog --help                                                                                   
Usage of ./dist/azureog:
  -bgImg string
        image to use as background (default "assets/og-standard.png")
  -fontPath string
        filename of the ttf font (default "assets/FiraSans-Light.ttf")
  -fontSize float
        font fontSize in points (default 140)
  -output string
        output path for the resulting image (default "og-image.png")
  -title string
        text to print on the image (default "Nothing To See Here, Move Along")
```

## Azure Functions

I followed this handy [reference](https://docs.microsoft.com/azure/azure-functions/create-first-function-vs-code-other?WT.mc_id=opensource-20963-brketels) to build a Go app as a custom Azure Function deployed on the Linux Consumption plan. Costs should be trivial.

## Make it yours!

To make one for yourself, replace `assets/og-standard.png` with a background of your own creation. Use the same file name or modify the code to reflect your new background image's file name.

Then you can build and use locally or deploy to Azure Functions to have your very own OG Social Image Sharing app.

## License 
GPL v3

### Origins and Credits

Inspired by, and much code from [DevQuotesPrinter](https://github.com/JosemyDuarte/DevsQuotesPrinter) - GPL v3 Licensed.

## Github Action
There is a github action available.
# Go Get Hammered

This is a very quick and dirty performance measuring tool for web applications.
The list of URLs to be tested are stored in a JSON format as follows:

```json
{
  "Sites": [
      {"Name": "Google", "Url": "http://www.google.com"},
      {"Name": "Google Australia", "Url": "http://www.google.com.au"}
  ]
}
```

**Note: This format will change as further features are introduced.**

###### Usage:
`gogh config.json`

###### TODO:
1. Sequential testing of sites.
2. Multiple lists of sites with a wait group for each.
3. Random test ordering (with seed parameter for repeatability)
4. Variable number of tests for each site.
5. Logging results to JSON.
6. D3 representation of results.

# Notes

## General things about speaking on technical topics

* keep it applied
* the good parts
* context and curation

## Key points

* work with various kinds of data
* like the concept of small interfaces
* there are love letters out there
* how to write your own
* how are they actually used, in which contexts
* can the IO interfaces teach us a bit about composition

## Code snippets

* [ ] interfaces
* [ ] CopyBuffer performance, and CPU
* [ ] ReaderFrom performance, and allocation difference

## Resources

### Crossings Streams

* https://www.datadoghq.com/blog/crossing-streams-love-letter-gos-io-reader/

The use of `ioutil.ReadAll` is a mistake.

* [ ] How often it is used? -- /home/tir/code/miku/ebba409208989306191926e238614f85
* [ ] Clone 300 repositories, count go files, analyze go files
* [ ] Also use Github BigQuery dataset

### BigQuery GitHub

* https://codelabs.developers.google.com/codelabs/bigquery-github/index.html?index=..%2F..index#0
* https://console.cloud.google.com/bigquery?project=golab-255608&folder&organizationId&p=bigquery-public-data&d=github_repos&t=languages&page=table

```sql
SELECT repo_name, language
FROM `bigquery-public-data.github_repos.languages` as ls, ls.language as language
WHERE language.name = 'go'
LIMIT 100
```

----

The buffersize of 4K is relatively good. Why? Is it the pagesize? Cacheline?
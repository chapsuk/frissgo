# Fresh issues from github

```bash
go get github.com/chapsuk/frissgo
```

## Configuration

See [exmaple](go.yml) for [go](https://github.com/golang/go) repo.

```yml
output:                             # output config
  format: text                      # output format, text or json
  target: stdout                    # output, stdout or file
  file_name:                        # file_name if target is file

github:                             # github config
  access_token:                     # personal access token (5000req/h limit) or empty (50)
  owner: golang                     # source organization
  repo: go                          # source repo
  filters:                          # get issues filters
    milestone:                      # empty for all, `none` - withou milesone, * - any milestone
    state: all                      # issue state, open, closed or all
    assignee:                       # issue assignee
    creator:                        # issue creator
    mentioned:                      # issue mentioned
    lables:                         # issue lables with `,` delimeter
    sort: comments                  # issues sort: created, updated, and comments.  Default value is "created".
    direction: desc                 # asc or desc
    period: 12h                     # period, time.ParseDuration

strategy:                           # rate strategy
  per_page: 200                     # issues count on 1 github request
  prior_authors:                    # using in category author condition
    - robpike                       # github login
    - bradfitz
    - davecheney
  categories:                       # array of interesting categories
    - name: CategoryExample         # category name, any string
      size: 10                      # max issues in category
      issues:                       # issuses weight coefficient values
        author: 99                  # if issue author one from prior_author +99 weight
        reaction:                   # issue reactions weight values
          plus: 10                  # weight + (issue_reaction_plus * 10)
        activity: 10                # weight + (issue_total_comments_count * 10)
      comments:                     # comments weight coefficient values
        activity: 10                # weight + (comments_count (only from period, github.period setting) * 10)
        author: 7
        reaction:
          total: 3
```

## Run

```bash
frissgo -cfg go.yml
```

## Image

![](https://media.giphy.com/media/11Tsyjflf2xq2A/giphy.gif)
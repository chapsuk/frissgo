# Fresh issues from github

[![Build Status](https://travis-ci.org/chapsuk/frissgo.svg?branch=master)](https://travis-ci.org/chapsuk/frissgo/)

```bash
go get github.com/chapsuk/frissgo
```

## Run

```bash
frissgo -cfg go.yml
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

## Output example

```bash
2017/02/05 00:33:21 github limits: github.Rate{Limit:5000, Remaining:4886, Reset:github.Timestamp{2017-02-05 00:43:19 +0300 MSK}}

[ MostActivity ]
=====================
https://github.com/golang/go/issues/18130
https://github.com/golang/go/issues/18887
https://github.com/golang/go/issues/18861
https://github.com/golang/go/issues/18874
https://github.com/golang/go/issues/18939
https://github.com/golang/go/issues/18896
https://github.com/golang/go/issues/18846
https://github.com/golang/go/issues/18911
https://github.com/golang/go/issues/18856
=====================

[ MostReaction ]
=====================
https://github.com/golang/go/issues/18130
https://github.com/golang/go/issues/12914
https://github.com/golang/go/issues/13560
https://github.com/golang/go/issues/18616
https://github.com/golang/go/issues/18802
https://github.com/golang/go/issues/18939
https://github.com/golang/go/issues/15314
https://github.com/golang/go/issues/18653
https://github.com/golang/go/issues/18861
https://github.com/golang/go/issues/18548
https://github.com/golang/go/issues/18597
https://github.com/golang/go/issues/17082
https://github.com/golang/go/issues/17725
https://github.com/golang/go/issues/18342
https://github.com/golang/go/issues/4899
https://github.com/golang/go/issues/16791
https://github.com/golang/go/issues/18846
=====================

[ MostFamous ]
=====================
https://github.com/golang/go/issues/18887
https://github.com/golang/go/issues/13560
https://github.com/golang/go/issues/18911
https://github.com/golang/go/issues/18865
https://github.com/golang/go/issues/18856
https://github.com/golang/go/issues/18906
https://github.com/golang/go/issues/18846
https://github.com/golang/go/issues/5170
https://github.com/golang/go/issues/14183
https://github.com/golang/go/issues/13579
=====================

[ MostPositive ]
=====================
https://github.com/golang/go/issues/18130
https://github.com/golang/go/issues/12914
https://github.com/golang/go/issues/13560
https://github.com/golang/go/issues/18616
https://github.com/golang/go/issues/18802
https://github.com/golang/go/issues/18653
https://github.com/golang/go/issues/18861
https://github.com/golang/go/issues/18939
https://github.com/golang/go/issues/18846
=====================

[ MostNegative ]
=====================
https://github.com/golang/go/issues/14932
=====================
```

![](https://media.giphy.com/media/11Tsyjflf2xq2A/giphy.gif)
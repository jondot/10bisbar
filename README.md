# 10bisbar

![](shot.png)

Quit your [10Bis](10bis.co.il) guesswork and never get stuck with an empty account.


## Quick Start

Download a binary from the [releases]() section. Then unzip it to your bitbar
plugins folder, and continue to the configuration section.


Or, if you prefer to build on your own and produce a binary. You'll need Go and
UPX (`brew install upx`).

`$ make dist`

Then point bitbar to this folder, or simply copy the `10bisbar` binary to your plugins folder.
There are no dependencies so you should be good to go immediately.

### Configuration

Place a file called `.10bis.json` at your home folder, so that it will be found in `~/.10bis.json`.

Copy this sample:

```json
{
  "username":"your-10bis-user",
  "password":"your-10bis-password",
  "prices":[30, 42, 69, 80],
  "no_food_days_per_week": 2
}
```

Put your username and password in, and then put your favorite meal price options in `prices`. Finally,
indicate how many days per week you don't use 10bis at all (typically weekends).


# Contributing
Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2014 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.




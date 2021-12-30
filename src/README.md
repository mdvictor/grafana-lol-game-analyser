# League of Legends Match Timeline Datasource for Grafana

The LoL Match Timeline Datasource plugin allows you to query and visualize your match history metrics from withing Grafana.

The datasource returns a timeline of your matches (using Riot's API) and allows to analyse and compare various metrics between players like:
* Player stats: `Current gold`, `Total gold`, `Gold per second`, `Minions killed`, `Jungle minions killed`, etc.
* Champion stats: `Ability power`, `Armor`, `Attack damage`, `Attack speed`, `Movement speed`, etc.
* Damage stats: `Magic damage done`, `Total damage done to champions`, `True damage taken`, etc.

## Configure the data source

When configuring the data source you will need to enter:
* A valid `Riot API key`
* Your `Summoner name`
* The platform on which the account exists as Riot defined them (e.g: `NA1`, `JP1`, `EUW1`)

## Feature requests and bug reports

I have no intentions on adding new features to the plugin for the moment. You are free to clone and modify the plugin as you wish from
[here](#https://github.com/mdvictor/lol-match-timeline)

The only improvement I plan on making sometime in the short future is refactoring the large number of calls to get match information.
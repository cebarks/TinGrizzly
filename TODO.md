# Todo

//TODO Entity-Component Framework
//TODO MsgPack/nano networking for multiplayer?
//TODO allow (re/over)loading of resources from filesystem instead of embeded source

- Distributed game update loop between all players in game?
  - use `id % players` to determine what entities are updated in what game, then synced over the network.

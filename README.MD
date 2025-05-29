# axefetch

neofetch but for [*axes](https://bitaxe.org)
```
    <" .    . ^<    test-miner@bitaxe
  [w@kI      Ik@q[   -----------------
 X$$$$W+    ~M$$$$Y   Model: 601
c$$$$$$%:``,%$$$$$$X   ASIC(s): 1x BM1370
W&8888W[C8%L]W8888&&   Firmware: ESP-Miner v2.8.0
""""""" z**X """""""   Uptime: 133h50m24s
       'xXXn`          Best Difficulty: 21M session, 210M best
      ,p$$$$d,         Shares: 881 accepted, 423 rejected (48.01%)
     +W$$$$$$W_        circlejerking into open source
     /q*B$$B*qt
```

# install
```sh
git clone https://github.com/0xf0xx0/axefetch.git
cd axefetch
go build
./axefetch --ip <axeip>
```

# TODO
- [?] clean + optimize code (like half done)
- [ ] more info functions
    - [ ] vendor ("GekkoScience 601") (depends on bitaxeorg/ESP-Miner/issues/716)
    - [x] model family ("Supra 401") (implemented, depends on bitaxeorg/ESP-Miner/pull/966)
- [ ] add icons
    - [ ] add support for no icon
    - same as info, but not dependent on pulls
- [ ] add color themes
    - same as icons
- [ ] add docs to config

bitcoin icon from https://bitcoinicons.com  
gamma icon from https://svgsilh.com/image/309911.html

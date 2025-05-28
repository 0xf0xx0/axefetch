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
- [ ] clean + optimize code
- [ ] more info functions
    - [ ] add vendor to model string ("GekkoScience 601") (relies on bitaxeorg/ESP-Miner/issues/716)
    - [ ] model family ("Supra 401") (relies on bitaxeorg/ESP-Miner/pulls/966)
- [ ] add icons
    - [ ] add support for no icon
    - same as info
- [ ] add color themes
    - same as info
- [ ] add docs to config

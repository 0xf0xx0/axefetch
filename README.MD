# axefetch

neofetch but for [*axes](https://bitaxe.org)
```
    <" .    . ^<   bitaxuh@bitaxe
  [w@kI      Ik@q[   --------------
 X$$$$W+    ~M$$$$Y   Model: Gamma 601
c$$$$$$%:``,%$$$$$$X   ASIC(s): 1x BM1370
W&8888W[C8%L]W8888&&   Firmware: ESP-Miner v2.8.0
""""""" z**X """""""   Uptime: 5d 13h 50m 24s
       'xXXn`          TBD: 42.069 GHz@42069 mV
      ,p$$$$d,         Best Difficulty: 330.73M session, 210.21G best
     +W$$$$$$W_        Shares: 881B accepted, 423M rejected (0.05%)
     /q*B$$B*qt        Pool: pooblic-pool.io:3333
                       Hashrate: 1.42 TH/s (actual), 1.42 TH/s (expected)
                       Efficiency: 14.1 J/TH (actual), 14.1 J/TH (expected)
                       Temperature: 55.33 C (asic), 66.00 C (vreg)
                       Heap: 8 MiB
                       circlejerking into open source
```

# install
```sh
git clone https://github.com/0xf0xx0/axefetch.git
cd axefetch
go build
cp contrib/* ~/.config/axefetch/ # copy default icons
./axefetch --ip <axeip>
```

# TODO
- [?] clean + optimize code (pprof returns no samples lol)
- [ ] more info functions
    - [ ] vendor ("GekkoScience 601") (depends on bitaxeorg/ESP-Miner/issues/716)
- [ ] add icons
    - same as info, but not dependent on pulls
- [ ] add color themes
    - same as icons
- [x] add docs to config

bitcoin icon from https://bitcoinicons.com  
gamma icon from https://svgsilh.com/image/309911.html

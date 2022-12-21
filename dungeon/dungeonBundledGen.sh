#!/bin/bash
fyne bundle -package dungeon ../icons/Agahnim1Gray.png > bundled.go
fyne bundle -package dungeon -append ../icons/Agahnim1.png >> bundled.go
fyne bundle -package dungeon -append ../icons/Agahnim2Gray.png >> bundled.go
fyne bundle -package dungeon -append ../icons/Agahnim2.png >> bundled.go
fyne bundle -package dungeon -append ../icons/Big_KeyGray.png >> bundled.go
fyne bundle -package dungeon -append ../icons/Big_Key.png >> bundled.go
fyne bundle -package dungeon -append ../icons/Chest.png >> bundled.go
fyne bundle -package dungeon -append ../icons/CompassGray.png >> bundled.go
fyne bundle -package dungeon -append ../icons/Compass.png >> bundled.go
fyne bundle -package dungeon -append ../icons/Empty_Chest.png >> bundled.go
fyne bundle -package dungeon -append ../icons/KeyGray.png >> bundled.go
fyne bundle -package dungeon -append ../icons/Key.png >> bundled.go
fyne bundle -package dungeon -append ../icons/MapGray.png >> bundled.go
fyne bundle -package dungeon -append ../icons/Map.png >> bundled.go
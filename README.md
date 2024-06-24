
# Gobot Automation Tool

Gobot is an automation tool that allows you to script various keyboard and mouse actions, as well as conditional logic and variables.
See [Commands](#commands) section for a list of commands supported by Gobot and their functionalities.

## Dependencies
- Gobot depends on [robotgo](https://github.com/go-vgo/robotgo) (see dependencies of robotgo).
- Gobot used the tiny and easy [rune scripting language](https://github.com/RednibCoding/runevm) for scripting

## Build
From the root directory of the project run: 
```
go build -ldflags="-s -w" .
```

## Run
From the root directory of the project run: 
```
gobot <script-file>.tin
```

>**Note:** Pressing `ESC` will stop the script immediately!

## Automation Functions

### keytap
- **Syntax**: `keytap(<key1>, <key2>, ...)`
- **Description**: Presses the specified key combination and releases it.
- **Example**: `press(lshift, a)`

### keypress
- **Syntax**: `keypress(<key1>, <key2>, ...)`
- **Description**: Presses the specified key combination.
- **Example**: `keypress(lshift, a)`

### keyrelease
- **Syntax**: `keyrelease(<key1>, <key2>, ...)`
- **Description**: Releases the specified key combination.
- **Example**: `keyrelease(lshift, a)`

### move
- **Syntax**: `move(<x>, <y>)`
- **Description**: Moves the mouse to the specified coordinates.
- **Example**: `move(100, 200)`

### getcolor
- **Syntax**: `getcolor(<x>, <y>)`
- **Description**: Retrieves the color at the specified screen coordinates and stores it in the given result variable. If the variable does not exist, it is created as a variable of type string.
- **Example**: `color = getcolor(100, 150)`

### colormatch
- **Syntax**: `colormatch(<color1>, <color2>, <threshold>)`
- **Description**: Compares two colors to see if they match within the specified threshold. The result is then saved in the given result variable. If the variable does not exist, it is created. Result can be 0 = colors do not match or 1 = colors match. Colors must be hex values preceded by `#`.
- **Example**: `isMatch = colormatch(myColor, "#deadbeef", "#10")`

## Examples

### Example 1: Simple Print
```
println("Hello, World!")
println("This is Gobot.")
```

### Example 2: Mouse Movement and Click
```
move(100, 200)
press("lmouse")
wait(500)
release("lmouse")
```

### Example 3: Color Checking
```
myColor = getcolor(175, 40)
println(myColor)
isMatch = colormatch(myColor, "#fed668", #01)

if isMatch {
    println "colors match :)"
} else {
	println "colors do not match"
}
```

## Supported Keys

```
	"left"               Left mouse button
	"center"             Middle mouse button
	"right"              Right mouse button

	"A-Z a-z 0-9"

	"backspace"
	"delete"
	"enter"
	"tab"
	"esc"
	"escape"
	"up"		         Up arrow key
	"down"		         Down arrow key
	"right"		         Right arrow key
	"left"		         Left arrow key
	"home"
	"end"
	"pageup"
	"pagedown"

	"f1"
	"f2"
	"f3"
	"f4"
	"f5"
	"f6"
	"f7"
	"f8"
	"f9"
	"f10"
	"f11"
	"f12"
	"f13"
	"f14"
	"f15"
	"f16"
	"f17"
	"f18"
	"f19"
	"f20"
	"f21"
	"f22"
	"f23"
	"f24"

	"cmd"		         is the "win" key for windows
	"lcmd"		         left command
	"rcmd"		         right command
	"alt"         
	"lalt"		         left alt
	"ralt"		         right alt
	"ctrl"         
	"lctrl"		         left ctrl
	"rctrl"		         right ctrl
	"control"         
	"shift"         
	"lshift"	         left shift
	"rshift"	         right shift
	"capslock"
	"space"
	"print"
	"printscreen"        No Mac support
	"insert"
	"menu"				 Windows only

	"audio_mute"		 Mute the volume
	"audio_vol_down"	 Lower the volume
	"audio_vol_up"		 Increase the volume
	"audio_play"
	"audio_stop"
	"audio_pause"
	"audio_prev"		 Previous Track
	"audio_next"		 Next Track
	"audio_rewind"       Linux only
	"audio_forward"      Linux only
	"audio_repeat"       Linux only
	"audio_random"       Linux only


	"num0"
	"num1"
	"num2"
	"num3"
	"num4"
	"num5"
	"num6"
	"num7"
	"num8"
	"num9"
	"num_lock"

	"num."
	"num+"
	"num-"
	"num*"
	"num/"
	"num_clear"
	"num_enter"
	"num_equal"

	"lights_mon_up"		 Turn up monitor brightness					No Windows support
	"lights_mon_down"	 Turn down monitor brightness				No Windows support
	"lights_kbd_toggle"	 Toggle keyboard backlight on/off			No Windows support
	"lights_kbd_up"		 Turn up keyboard backlight brightness		No Windows support
	"lights_kbd_down"	 Turn down keyboard backlight brightness	No Windows support
```

## License
Gobot is open-source software licensed under the MIT license.
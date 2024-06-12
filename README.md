
# Gobot Automation Tool

Gobot is an automation tool that allows you to script various keyboard and mouse actions, as well as conditional logic and variables.
See [Commands](#commands) section for a list of commands supported by Gobot and their functionalities.

## Dependencies
Gobot depends on [robotgo](https://github.com/go-vgo/robotgo) (see dependencies of robotgo).

## Build
From the root directory of the project run: 
```
go build -ldflags="-s -w" .
```

## General Script-Syntax
- A script is composed of a list of [commands](#commands). See [Examples](#examples).
- Commands that have arguments should be followed by a colon, and the arguments should be separated by commas.
- Commands that have no arguments may not be followed by a colon.
- A label should be defined by a leading `#` and can be referred to in a `goto` command.
- Only one command or label can be on a line.
- Lines starting with a `;` are considered comments and will be skipped.
- In commands like `print`, `println`, `wait`, `move`, `ifcolor`, `set`, `add`, `sub`, `ifequal`, `ifgreater`, and `ifless`, arguments can also be variables. These variables will be evaluated to their current values.
- All variables are of type Integer (whole numbers)

### Example
```
#start
println: Hello, World!
set: x, 200
set: y, 40
; This is a comment and will be skipped
move: x, y
goto: end
println: This will be skipped.
#end
println: End of the script.
```

## Commands

### print
- **Syntax**: `print: <arg1>, <arg2>, ...`
- **Description**: Prints the given arguments. String arguments must be enclosed in double quotes, variables are not.
- **Example**: `print: "Hello,", " World!", varName`

### println
- **Syntax**: `println: <arg1>, <arg2>, ...`
- **Description**: Prints the given arguments and a new line. String arguments must be enclosed in double quotes, variables are not.
- **Example**: `println: "Hello,", " World!", varName`

### move
- **Syntax**: `move: <x>, <y>`
- **Description**: Moves the mouse to the specified coordinates.
- **Example**: `move: 100, 200`

### press
- **Syntax**: `press: <key1>, <key2>, ...`
- **Description**: Presses the specified keys.
- **Example**: `press: lshift, a`

### release
- **Syntax**: `release: <key1>, <key2>, ...`
- **Description**: Releases the specified keys.
- **Example**: `release: lshift, a`

### autopress
- **Syntax**: `autopress: <key1>, <key2>, ...`
- **Description**: Presses the specified keys, with a delay of 50 milliseconds and then releases the keys again
- **Example**: `autopress: lshift, a`

### ifpressed
- **Syntax**: `ifpressed: <key>`
- **Description**: Executes the next command if the specified key is pressed.
- **Example**: `ifpressed: lshift`

### ifnotpressed
- **Syntax**: `ifnotpressed: <key>`
- **Description**: Executes the next command if the specified key is not pressed.
- **Example**: `ifnotpressed: lshift`

### wait
- **Syntax**: `wait: <milliseconds>`
- **Description**: Waits for the specified duration.
- **Example**: `wait: 1000`

### savecolor
- **Syntax**: `savecolor: <x>, <y>`
- **Description**: Saves the color of the pixel at the given x and y position.
- **Example**: `savecolor: 200, 300`

### printcolorrgb
- **Syntax**: `printcolorrgb`
- **Description**: Prints the saved color in RGB format.
- **Example**: `printcolorrgb`

### printcolorhex
- **Syntax**: `printcolorhex`
- **Description**: Prints the saved color in hexadecimal format.
- **Example**: `printcolorhex`

### ifcolor
- **Syntax**: `ifcolor: <hexcolor>, <threshold>`
- **Description**: Executes the next command if the saved color matches the specified color within the given threshold.
- **Example**: `ifcolor: ffffff, 0a`

### goto
- **Syntax**: `goto: <label>`
- **Description**: Jumps to the specified label.
- **Example**: `goto: start`
- **Anonymous Labels**: Use `goto: @f` to jump to the next anonymous label and `goto: @b` to jump to the previous anonymous label. The number of `f` or `b` characters determines how many anonymous labels to jump forward or backward. For example, `goto: @fff` jumps three anonymous labels forward, and `goto: @bbb` jumps three anonymous labels backward.
- **Define Anonymous Labels**: Use `#@` to define an anonymous label.

### gosub
- **Syntax**: `gosub: <label>`
- **Description**: Jumps to the specified label and saves the return address.
- **Example**: `gosub: subroutine`

### return
- **Syntax**: `return`
- **Description**: Returns to the address saved by the last `gosub` command.
- **Example**: `return`

### set
- **Syntax**: `set: <variable>, <value>`
- **Description**: Sets the specified variable to the given value.
- **Example**: `set: a, 10`

### add
- **Syntax**: `add: <variable>, <value>`
- **Description**: Adds the specified value to the variable.
- **Example**: `add: a, 5`

### sub
- **Syntax**: `sub: <variable>, <value>`
- **Description**: Subtracts the specified value from the variable.
- **Example**: `sub: a, 3`

### ifequal
- **Syntax**: `ifequal: <variable>, <value>`
- **Description**: Executes the next command if the variable equals the specified value.
- **Example**: `ifequal: a, 10`

### ifgreater
- **Syntax**: `ifgreater: <variable>, <value>`
- **Description**: Executes the next command if the variable is greater than the specified value.
- **Example**: `ifgreater: a, 10`

### ifless
- **Syntax**: `ifless: <variable>, <value>`
- **Description**: Executes the next command if the variable is less than the specified value.
- **Example**: `ifless: a, 10`
## Examples

### Example 1: Simple Print
```
println: Hello, World!
println: This is Gobot.
```

### Example 2: Mouse Movement and Click
```
move: 100, 200
press: lmouse
wait: 500
release: lmouse
```

### Example 3: Color Checking
```
savecolor: 200, 300
printcolorhex
ifcolor: ffffff, 0a
    println: The color is white.
```

### Example 4: Using Variables
```
set: a, 10
add: a, 5
sub: a, 3

ifequal: a, 12
    println: "Variable a is: ", a

ifgreater: a, 10
    println: "Variable a is greater than 10"

ifless: a, 15
    println: "Variable a is less than 15"
```

### Example 5: Conditional Execution with Keys
```
press: lshift
ifpressed: lshift
    println: "lshift is pressed."

release: lshift

ifnotpressed: lshift
    println: "lshift is not pressed."
```

### Example 6: Labels and Goto
```
#start
println: "Start of the script."
goto: end
println: "This will be skipped."
#end
println: "End of the script."
```

### Example 7: Subroutine with Gosub and Return
```
println: "--- Start of the script ---"

gosub: subroutine1
println: "Back from subroutine 1"
goto: end

println: "This will be skipped."

#subroutine1
    println: "In subroutine 1"
    gosub: subroutine2
    println: "Back from subroutine 2"
    return

#subroutine2
    println: "In subroutine 2"
    return

#end
    println: "--- End of the script. ---"

; output:
; --- Start of the script ---
; In subroutine 1
; In subroutine 2
; Back from subroutine 2
; Back from subroutine 1
; --- End of the script. ---
```

### Example 8: Anonymous Labels
```
println: "--- Start of the script. ---"
goto: @f
println: "This will be skipped."

#@
    println: "Reached the first anonymous label."
    goto: @ff
    println: "This will be skipped too."

#@
    println: "Reached the second anonymous label."
    goto: end
    println: "This will be skipped as well."

#@ 
    println: "Reached the third anonymous label."
    goto: @bb
    println: "This will be skipped as well."

#end
    println: "--- End of the script. ---"

; output:
; --- Start of the script. ---
; Reached the first anonymous label.
; Reached the third anonymous label.
; Reached the second anonymous label.
; --- End of the script. ---
```

## Supported Keys

- **lshift**: Left Shift key
- **rshift**: Right Shift key
- **lctrl**: Left Control key
- **rctrl**: Right Control key
- **lalt**: Left Alt key
- **ralt**: Right Alt key
- **space**: Spacebar
- **enter**: Enter key
- **backspace**: Backspace key
- **tab**: Tab key
- **esc**: Escape key
- **delete**: Delete key
- **insert**: Insert key
- **home**: Home key
- **end**: End key
- **pageup**: Page Up key
- **pagedown**: Page Down key
- **up**: Up Arrow key
- **down**: Down Arrow key
- **left**: Left Arrow key
- **right**: Right Arrow key
- **f1** to **f12**: Function keys F1 to F12
- **numlock**: Num Lock key
- **capslock**: Caps Lock key
- **scrolllock**: Scroll Lock key
- **pause**: Pause key
- **printscreen**: Print Screen key
- **windows**: Windows key
- **lmouse**: Left Mouse Button
- **rmouse**: Right Mouse Button

### Alphabet Keys
- **a** to **z**: Alphabet keys A to Z

### Number Keys
- **0** to **9**: Number keys 0 to 9

### Numpad Keys
- **numpad0** to **numpad9**: Numpad keys 0 to 9
- **numpadadd**: Numpad Add key
- **numpadsub**: Numpad Subtract key
- **numpadmul**: Numpad Multiply key
- **numpaddiv**: Numpad Divide key
- **numpaddecimal**: Numpad Decimal key
- **numpadenter**: Numpad Enter key

### Symbols
- **semicolon**: Semicolon (;)
- **equals**: Equals (=)
- **comma**: Comma (,)
- **minus**: Minus (-)
- **period**: Period (.)
- **slash**: Slash (/)
- **backslash**: Backslash (\)
- **openbracket**: Open Bracket ([)
- **closebracket**: Close Bracket (])
- **quote**: Quote (')

## License
Gobot is open-source software licensed under the MIT license.
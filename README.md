
# JBot Automation Tool

JBot is an automation tool that allows you to script various keyboard and mouse actions, as well as conditional logic and variables.
See [Commands](#commands) section for a list of commands supported by JBot and their functionalities.

## Dependencies
JBot depends on the build tool [ant](https://ant.apache.org/)

## Build
From the root directory of the project, run: 
```
ant
```

## General Script-Syntax
- A script is composed of a list of [commands](#commands). See [Examples](#examples).
- Commands that have arguments should be followed by a colon, and the arguments should be separated by commas.
- Commands that have no arguments may not be followed by a colon.
- A label should be defined by a leading `#` and can be referred to in a `goto` command.
- Only one command or label can be on a line.
- Lines starting with a `;` are considered comments and will be skipped.
- In commands like `wait`, `move`, `ifcolor`, `set`, `add`, `sub`, `ifequal`, `ifgreater`, and `ifless`, arguments can be variables (letters from 'a' to 'z'). These variables will be evaluated to their current values.

### Example
```
#start
print: Hello, World!
printnl
set: x, 200
set: y, 40
; This is a comment and will be skipped
move: x, y
goto: end
print: This will be skipped.
printnl
#end
print: End of the script.
printnl
```

## Commands

### print
- **Syntax**: `print: <string>`
- **Description**: Prints the given string.
- **Example**: `print: Hello, World!`

### printnl
- **Syntax**: `printnl`
- **Description**: Prints a new line.
- **Example**: `printnl`

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
- **Syntax**: `savecolor`
- **Description**: Saves the color of the pixel at the current mouse position.
- **Example**: `savecolor`

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

### printvar
- **Syntax**: `printvar: <variable>`
- **Description**: Prints the value of the specified variable.
- **Example**: `printvar: a`

## Examples

### Example 1: Simple Print
```
print: Hello, World!
printnl
print: This is JBot.
printnl
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
move: 150, 150
savecolor
printcolorhex
ifcolor: ffffff, 0a
    print: The color is white.
printnl
```

### Example 4: Using Variables
```
set: a, 10
add: a, 5
sub: a, 3
ifequal: a, 12
    print: Variable a is 12.
printnl
ifgreater: a, 10
    print: Variable a is greater than 10.
printnl
ifless: a, 15
    print: Variable a is less than 15.
printnl
printvar: a
printnl
```

### Example 5: Conditional Execution with Keys
```
press: lshift
ifpressed: lshift
    print: lshift is pressed.
printnl
release: lshift
ifnotpressed: lshift
    print: lshift is not pressed.
printnl
```

### Labels and Goto
```
#start
print: Start of the script.
printnl
goto: end
print: This will be skipped.
printnl
#end
print: End of the script.
printnl
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
- **backslash**: Backslash (\\)
- **openbracket**: Open Bracket ([)
- **closebracket**: Close Bracket (])
- **quote**: Quote (')

## License
JBot is open-source software licensed under the MIT license.

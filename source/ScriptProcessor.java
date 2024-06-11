import java.awt.Robot;
import java.awt.AWTException;
import java.awt.Color;
import java.awt.MouseInfo;
import java.awt.Point;
import java.awt.event.InputEvent;
import java.awt.event.KeyEvent;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.List;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class ScriptProcessor {
    private Robot robot;
    private Map<String, Integer> keyMap;
    private Set<String> pressedKeys;
    private Color savedColor;
    private boolean executeNextLine;
    private Map<String, Integer> variables;  // Variables with user-defined names
    private Map<String, Integer> labels;

    public ScriptProcessor() throws AWTException {
        robot = new Robot();
        initializeKeyMap();
        initializeVariables();
        pressedKeys = new HashSet<>();
        savedColor = null;
        executeNextLine = true;
        labels = new HashMap<>();
    }

    private void initializeKeyMap() {
        keyMap = new HashMap<>();
        keyMap.put("lshift", KeyEvent.VK_SHIFT);
        keyMap.put("rshift", KeyEvent.VK_SHIFT);
        keyMap.put("lctrl", KeyEvent.VK_CONTROL);
        keyMap.put("rctrl", KeyEvent.VK_CONTROL);
        keyMap.put("lalt", KeyEvent.VK_ALT);
        keyMap.put("ralt", KeyEvent.VK_ALT);
        keyMap.put("space", KeyEvent.VK_SPACE);
        keyMap.put("enter", KeyEvent.VK_ENTER);
        keyMap.put("backspace", KeyEvent.VK_BACK_SPACE);
        keyMap.put("tab", KeyEvent.VK_TAB);
        keyMap.put("esc", KeyEvent.VK_ESCAPE);
        keyMap.put("delete", KeyEvent.VK_DELETE);
        keyMap.put("insert", KeyEvent.VK_INSERT);
        keyMap.put("home", KeyEvent.VK_HOME);
        keyMap.put("end", KeyEvent.VK_END);
        keyMap.put("pageup", KeyEvent.VK_PAGE_UP);
        keyMap.put("pagedown", KeyEvent.VK_PAGE_DOWN);
        keyMap.put("up", KeyEvent.VK_UP);
        keyMap.put("down", KeyEvent.VK_DOWN);
        keyMap.put("left", KeyEvent.VK_LEFT);
        keyMap.put("right", KeyEvent.VK_RIGHT);
        keyMap.put("f1", KeyEvent.VK_F1);
        keyMap.put("f2", KeyEvent.VK_F2);
        keyMap.put("f3", KeyEvent.VK_F3);
        keyMap.put("f4", KeyEvent.VK_F4);
        keyMap.put("f5", KeyEvent.VK_F5);
        keyMap.put("f6", KeyEvent.VK_F6);
        keyMap.put("f7", KeyEvent.VK_F7);
        keyMap.put("f8", KeyEvent.VK_F8);
        keyMap.put("f9", KeyEvent.VK_F9);
        keyMap.put("f10", KeyEvent.VK_F10);
        keyMap.put("f11", KeyEvent.VK_F11);
        keyMap.put("f12", KeyEvent.VK_F12);
        keyMap.put("numlock", KeyEvent.VK_NUM_LOCK);
        keyMap.put("capslock", KeyEvent.VK_CAPS_LOCK);
        keyMap.put("scrolllock", KeyEvent.VK_SCROLL_LOCK);
        keyMap.put("pause", KeyEvent.VK_PAUSE);
        keyMap.put("printscreen", KeyEvent.VK_PRINTSCREEN);
        keyMap.put("windows", KeyEvent.VK_WINDOWS);
        keyMap.put("lmouse", InputEvent.BUTTON1_DOWN_MASK);
        keyMap.put("rmouse", InputEvent.BUTTON3_DOWN_MASK);
        
        // Alphabet keys
        for (char c = 'a'; c <= 'z'; c++) {
            keyMap.put(String.valueOf(c), KeyEvent.getExtendedKeyCodeForChar(c));
        }
    
        // Number keys
        for (int i = 0; i <= 9; i++) {
            keyMap.put(String.valueOf(i), KeyEvent.getExtendedKeyCodeForChar('0' + i));
        }
    
        // Numpad keys
        keyMap.put("numpad0", KeyEvent.VK_NUMPAD0);
        keyMap.put("numpad1", KeyEvent.VK_NUMPAD1);
        keyMap.put("numpad2", KeyEvent.VK_NUMPAD2);
        keyMap.put("numpad3", KeyEvent.VK_NUMPAD3);
        keyMap.put("numpad4", KeyEvent.VK_NUMPAD4);
        keyMap.put("numpad5", KeyEvent.VK_NUMPAD5);
        keyMap.put("numpad6", KeyEvent.VK_NUMPAD6);
        keyMap.put("numpad7", KeyEvent.VK_NUMPAD7);
        keyMap.put("numpad8", KeyEvent.VK_NUMPAD8);
        keyMap.put("numpad9", KeyEvent.VK_NUMPAD9);
        keyMap.put("numpadadd", KeyEvent.VK_ADD);
        keyMap.put("numpadsub", KeyEvent.VK_SUBTRACT);
        keyMap.put("numpadmul", KeyEvent.VK_MULTIPLY);
        keyMap.put("numpaddiv", KeyEvent.VK_DIVIDE);
        keyMap.put("numpaddecimal", KeyEvent.VK_DECIMAL);
        keyMap.put("numpadenter", KeyEvent.VK_ENTER);
    
        // Symbols
        keyMap.put("semicolon", KeyEvent.VK_SEMICOLON);
        keyMap.put("equals", KeyEvent.VK_EQUALS);
        keyMap.put("comma", KeyEvent.VK_COMMA);
        keyMap.put("minus", KeyEvent.VK_MINUS);
        keyMap.put("period", KeyEvent.VK_PERIOD);
        keyMap.put("slash", KeyEvent.VK_SLASH);
        keyMap.put("backslash", KeyEvent.VK_BACK_SLASH);
        keyMap.put("openbracket", KeyEvent.VK_OPEN_BRACKET);
        keyMap.put("closebracket", KeyEvent.VK_CLOSE_BRACKET);
        keyMap.put("quote", KeyEvent.VK_QUOTE);
    }
    
    private void initializeVariables() {
        variables = new HashMap<>();  // Initialize as an empty map, allowing dynamic variable creation
    }

    public void executeScript(String scriptPath) {
        try {
            List<String> lines = Files.readAllLines(Paths.get(scriptPath));

            // First pass: store labels and their line numbers
            for (int i = 0; i < lines.size(); i++) {
                String line = lines.get(i).trim();
                if (line.startsWith("#")) {
                    labels.put(line.substring(1).trim(), i);
                }
            }

            // Second pass: execute commands
            for (int i = 0; i < lines.size(); i++) {
                String line = lines.get(i).trim();
                if (line.isEmpty() || line.startsWith("#") || line.startsWith(";")) continue;  // Skip empty lines, labels, and comments

                if (!executeNextLine) {
                    executeNextLine = true; // Reset flag to execute subsequent lines
                    continue;
                }

                String command;
                String[] args = new String[0];
                int colonIndex = line.indexOf(':');
                if (colonIndex == -1) {
                    command = line;
                } else {
                    command = line.substring(0, colonIndex).trim();
                    args = line.substring(colonIndex + 1).trim().split(",");
                }

                int newLine = executeCommand(lines, command, args, i + 1);
                if (newLine >= 0) {
                    i = newLine - 1; // -1 because the for loop will increment i
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
            System.exit(0);
        }
    }

    private int executeCommand(List<String> lines, String command, String[] args, int lineNumber) throws InterruptedException {
        switch (command) {
            case "println":
                if (args.length != 1) {
                    System.out.println("Error on line " + lineNumber + ": println command requires exactly 1 argument");
                    return -1;
                }
                System.out.println(args[0].trim());
                break;
            case "print":
                if (args.length != 1) {
                    System.out.println("Error on line " + lineNumber + ": print command requires exactly 1 argument");
                    return -1;
                }
                System.out.print(args[0].trim());
                break;
            case "printnl":
                if (args.length != 0) {
                    System.out.println("Error on line " + lineNumber + ": printnl command requires no arguments");
                    return -1;
                }
                System.out.print("\n");
                break;
            case "move":
                if (args.length != 2) {
                    System.out.println("Error on line " + lineNumber + ": move command requires exactly 2 arguments");
                    return -1;
                }
                int x = getValue(args[0].trim());
                int y = getValue(args[1].trim());
                robot.mouseMove(x, y);
                break;
            case "autopress":
                if (args.length < 1) {
                    System.out.println("Error on line " + lineNumber + ": autopress command requires at least 1 argument");
                    return -1;
                }
                Thread.sleep(80);
                for (String arg : args) {
                    int key = keyMap.getOrDefault(arg.trim(), -1);
                    if (key != -1) {
                        if (key == InputEvent.BUTTON1_DOWN_MASK || key == InputEvent.BUTTON3_DOWN_MASK) {
                            robot.mousePress(key);
                        } else {
                            robot.keyPress(key);
                        }
                        pressedKeys.add(arg);
                    } else {
                        System.out.println("Error on line " + lineNumber + ": Invalid key: " + arg.trim());
                        return -1;
                    }
                    Thread.sleep(80);
                }
                Thread.sleep(80);
                for (String arg : args) {
                    int key = keyMap.getOrDefault(arg.trim(), -1);
                    if (key != -1) {
                        if (key == InputEvent.BUTTON1_DOWN_MASK || key == InputEvent.BUTTON3_DOWN_MASK) {
                            robot.mouseRelease(key);
                        } else {
                            robot.keyRelease(key);
                        }
                        pressedKeys.remove(arg);
                    } else {
                        System.out.println("Error on line " + lineNumber + ": Invalid key: " + arg.trim());
                        return -1;
                    }
                    Thread.sleep(80);
                }
                Thread.sleep(80);
                break;
            case "press":
                if (args.length < 1) {
                    System.out.println("Error on line " + lineNumber + ": press command requires at least 1 argument");
                    return -1;
                }
                Thread.sleep(40);
                for (String arg : args) {
                    int key = keyMap.getOrDefault(arg.trim(), -1);
                    if (key != -1) {
                        if (key == InputEvent.BUTTON1_DOWN_MASK || key == InputEvent.BUTTON3_DOWN_MASK) {
                            robot.mousePress(key);
                        } else {
                            robot.keyPress(key);
                        }
                        pressedKeys.add(arg);
                    } else {
                        System.out.println("Error on line " + lineNumber + ": Invalid key: " + arg.trim());
                        return -1;
                    }
                    Thread.sleep(40);
                }
                break;
            case "release":
                if (args.length < 1) {
                    System.out.println("Error on line " + lineNumber + ": release command requires at least 1 argument");
                    return -1;
                }
                Thread.sleep(40);
                for (String arg : args) {
                    int key = keyMap.getOrDefault(arg.trim(), -1);
                    if (key != -1) {
                        if (key == InputEvent.BUTTON1_DOWN_MASK || key == InputEvent.BUTTON3_DOWN_MASK) {
                            robot.mouseRelease(key);
                        } else {
                            robot.keyRelease(key);
                        }
                        pressedKeys.remove(arg);
                    } else {
                        System.out.println("Error on line " + lineNumber + ": Invalid key: " + arg.trim());
                        return -1;
                    }
                    Thread.sleep(40);
                }
                break;
            case "ifpressed": {
                if (args.length != 1) {
                    System.out.println("Error on line " + lineNumber + ": ifpressed command requires exactly 1 argument");
                    return -1;
                }
                String keystr = args[0].trim();
                int keyPressed = keyMap.getOrDefault(keystr, -1);
                if (keyPressed == -1) {
                    System.out.println("Error on line " + lineNumber + ": Invalid key: " + keystr);
                    return -1;
                }
                if (!pressedKeys.contains(keystr)) {
                    executeNextLine = false; // Skip the next line if the key is not pressed
                }
                break;
            }
            
            case "ifnotpressed": {
                if (args.length != 1) {
                    System.out.println("Error on line " + lineNumber + ": ifnotpressed command requires exactly 1 argument");
                    return -1;
                }
                String keystr = args[0].trim();
                int keyPressed = keyMap.getOrDefault(keystr, -1);
                if (keyPressed == -1) {
                    System.out.println("Error on line " + lineNumber + ": Invalid key: " + keystr);
                    return -1;
                }
                if (pressedKeys.contains(keystr)) {
                    executeNextLine = false; // Skip the next line if the key is pressed
                }
                break;
            }
            case "wait":
                if (args.length != 1) {
                    System.out.println("Error on line " + lineNumber + ": wait command requires exactly 1 argument");
                    return -1;
                }
                int duration = getValue(args[0].trim());
                Thread.sleep(duration);
                break;
            case "savecolor":
                if (args.length != 0) {
                    System.out.println("Error on line " + lineNumber + ": savecolor command requires no arguments");
                    return -1;
                }
                Point mousePosition = MouseInfo.getPointerInfo().getLocation();
                savedColor = robot.getPixelColor(mousePosition.x, mousePosition.y);
                break;
            case "printcolorrgb":
                if (args.length != 0) {
                    System.out.println("Error on line " + lineNumber + ": printcolor command requires no arguments");
                    return -1;
                }
                if (savedColor == null) {
                    System.out.println("Error on line " + lineNumber + ": No color saved, use savecolor command first");
                    return -1;
                }
                System.out.print("Saved Color: RGB(" + savedColor.getRed() + ", " + savedColor.getGreen() + ", " + savedColor.getBlue() +")");
                break;
            case "printcolorhex":
                if (args.length != 0) {
                    System.out.println("Error on line " + lineNumber + ": printcolor command requires no arguments");
                    return -1;
                }
                if (savedColor == null) {
                    System.out.println("Error on line " + lineNumber + ": No color saved, use savecolor command first");
                    return -1;
                }
                System.out.print("Hex: #" + Integer.toHexString(savedColor.getRGB()).substring(2).toUpperCase());
                break;
            case "ifcolor":
                if (args.length != 2) {
                    System.out.println("Error on line " + lineNumber + ": doifcolor command requires exactly 2 arguments");
                    return -1;
                }
                if (savedColor == null) {
                    System.out.println("Error on line " + lineNumber + ": No color saved, use savecolor command first");
                    return -1;
                }
                String colorHex = args[0].trim();
                int threshold = Integer.parseInt(args[1].trim(), 16);
                Color targetColor = Color.decode("#" + colorHex);
                if (!colorsMatch(savedColor, targetColor, threshold)) {
                    executeNextLine = false; // Skip the next line
                }
                break;
            case "printvar": {
                if (args.length != 1) {
                    System.out.println("Error on line " + lineNumber + ": printvar command requires exactly 1 argument");
                    return -1;
                }
                String varName = args[0].trim();
                if (!variables.containsKey(varName)) {
                    System.out.println("Error on line " + lineNumber + ": Variable not declared: " + varName);
                    return -1;
                }
                System.out.print(variables.get(varName));
                break;
            }
            case "goto":
                if (args.length != 1) {
                    System.out.println("Error on line " + lineNumber + ": goto command requires exactly 1 argument");
                    return -1;
                }
                String label = args[0].trim();
                if (!labels.containsKey(label)) {
                    System.out.println("Error on line " + lineNumber + ": Undefined label: " + label);
                    return -1;
                }
                return labels.get(label); // Jump to the label
            case "set":
                if (args.length != 2) {
                    System.out.println("Error on line " + lineNumber + ": set command requires exactly 2 arguments");
                    return -1;
                }
                String varName = args[0].trim();
                int value = getValue(args[1].trim());
                variables.put(varName, value);
                break;
            case "add":
                if (args.length != 2) {
                    System.out.println("Error on line " + lineNumber + ": add command requires exactly 2 arguments");
                    return -1;
                }
                varName = args[0].trim();
                if (!variables.containsKey(varName)) {
                    System.out.println("Error on line " + lineNumber + ": Variable not declared: " + varName);
                    return -1;
                }
                value = getValue(args[1].trim());
                variables.put(varName, variables.get(varName) + value);
                break;
            case "sub":
                if (args.length != 2) {
                    System.out.println("Error on line " + lineNumber + ": sub command requires exactly 2 arguments");
                    return -1;
                }
                varName = args[0].trim();
                if (!variables.containsKey(varName)) {
                    System.out.println("Error on line " + lineNumber + ": Variable not declared: " + varName);
                    return -1;
                }
                value = getValue(args[1].trim());
                variables.put(varName, variables.get(varName) - value);
                break;
            case "ifequal":
                if (args.length != 2) {
                    System.out.println("Error on line " + lineNumber + ": ifequal command requires exactly 2 arguments");
                    return -1;
                }
                varName = args[0].trim();
                if (!variables.containsKey(varName)) {
                    System.out.println("Error on line " + lineNumber + ": Variable not declared: " + varName);
                    return -1;
                }
                value = getValue(args[1].trim());
                if (variables.get(varName) != value) {
                    executeNextLine = false; // Skip the next line
                }
                break;
            case "ifgreater":
                if (args.length != 2) {
                    System.out.println("Error on line " + lineNumber + ": ifgreater command requires exactly 2 arguments");
                    return -1;
                }
                varName = args[0].trim();
                if (!variables.containsKey(varName)) {
                    System.out.println("Error on line " + lineNumber + ": Variable not declared: " + varName);
                    return -1;
                }
                value = getValue(args[1].trim());
                if (variables.get(varName) <= value) {
                    executeNextLine = false; // Skip the next line
                }
                break;
            case "ifless":
                if (args.length != 2) {
                    System.out.println("Error on line " + lineNumber + ": ifless command requires exactly 2 arguments");
                    return -1;
                }
                varName = args[0].trim();
                if (!variables.containsKey(varName)) {
                    System.out.println("Error on line " + lineNumber + ": Variable not declared: " + varName);
                    return -1;
                }
                value = getValue(args[1].trim());
                if (variables.get(varName) >= value) {
                    executeNextLine = false; // Skip the next line
                }
                break;
            default:
                System.out.println("Unknown command: " + command);
                return -1;
        }
        return -1; // No jump
    }

    private boolean colorsMatch(Color c1, Color c2, int threshold) {
        int rDiff = Math.abs(c1.getRed() - c2.getRed());
        int gDiff = Math.abs(c1.getGreen() - c2.getGreen());
        int bDiff = Math.abs(c1.getBlue() - c2.getBlue());
        return (rDiff <= threshold && gDiff <= threshold && bDiff <= threshold);
    }

    private int getValue(String arg) throws NumberFormatException {
        if (variables.containsKey(arg)) {
            return variables.get(arg);
        }
        return Integer.parseInt(arg);
    }
}

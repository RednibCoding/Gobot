import java.awt.AWTException;
import java.util.ArrayList;
import java.util.List;

public class Jbot {
    public static void main(String[] args) {
        List<String> argstmp = new ArrayList<>();
        argstmp.add("script.txt");
        args = argstmp.toArray(new String[0]);
        
        if (args.length != 1) {
            System.out.println("Usage: java jbot <script-file>");
            return;
        }
        try {
            ScriptProcessor processor = new ScriptProcessor();
            processor.executeScript(args[0]);
        } catch (AWTException e) {
            e.printStackTrace();
        }
    }
}

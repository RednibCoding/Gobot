import java.awt.AWTException;
import java.io.File;

public class Jbot {
    public static void main(String[] args) {

        // List<String> argstmp = new ArrayList<>();
        // argstmp.add("script.txt");
        // args = argstmp.toArray(new String[0]);
        
        if (args.length != 1) {
            System.out.println("Usage: java -jar jbot <script-file>");
            return;
        }

        File scriptFile = new File(args[0]);
        if (!scriptFile.exists() || !scriptFile.isFile()) {
            System.err.println("Error: The specified script file does not exist: " + args[0]);
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

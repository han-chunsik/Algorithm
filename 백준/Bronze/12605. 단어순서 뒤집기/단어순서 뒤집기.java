import java.io.*;
import java.util.*;

public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine());

        for (int i=1; i<n+1; i++){
            String l = br.readLine();

            Stack<String> stack = new Stack<>();
            Arrays.stream(l.split(" ")).forEach(stack::push);
            
            StringBuilder result = new StringBuilder();

            int size = stack.size();
            for (int j=0; j<size;j++){
                result.append(stack.pop());
                if (j < size -1){
                    result.append(" ");
                }
            }
            System.out.println(String.format("Case #%d: %s", i, result.toString()));
        }
    }
}
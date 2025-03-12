import java.io.*;
import java.util.*;

public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine());

        Stack<Integer> stack = new Stack<>();

        for (int i=1; i<n+1; i++){
            int h = Integer.parseInt(br.readLine());
            stack.push(h);
        }

        int max = 0;
        int result = 0;

        while (!stack.isEmpty()){
            if(stack.peek() > max){
                max = stack.peek();
                result++;
            }
            stack.pop();
        }
        System.out.println(result);
    }
}
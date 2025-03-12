import java.io.*;
import java.util.*;

public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringTokenizer stk = new StringTokenizer(br.readLine());
        int a = Integer.parseInt(stk.nextToken());
        int b = Integer.parseInt(stk.nextToken());
        int v = Integer.parseInt(stk.nextToken());

        int day = (v-b)/(a-b);

        if((v-a)%(a-b)!=0){
            day++;
        }

        System.out.println(day);
    }
}
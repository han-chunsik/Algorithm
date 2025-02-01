import java.io.*;
import java.util.*;
	/*
	*	정문에서 가장 짧은 거리에 있는 방
	*	w개의 방이 있는 h층 건물
	*	앨리베이터는 가장 왼쪽에 있으며, 정문도 거기에 가깝게 있음, 정문과 앨리베이터의 거리는 고려하지 않음
	*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int t = Integer.parseInt(br.readLine());

        List<Integer> resultList = new ArrayList<>(t);

        for(int i=0; i < t; i++) {
            StringTokenizer stk = new StringTokenizer(br.readLine());
            int h = Integer.parseInt(stk.nextToken());
            int w = Integer.parseInt(stk.nextToken());
            int n = Integer.parseInt(stk.nextToken());

            int f = n % h;
            int r = n/h + 1;
            if (f == 0) {
                r -= 1;
                f += h;
            }
            
            resultList.add(f*100 + r);
        }

        for (int result : resultList){
            System.out.println(result);
        }
        
    }
}
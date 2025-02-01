import java.io.*;
import java.util.*;
	/*
	* 1 2 3 4 5 6 7 8: ascending
	* 8 7 6 5 4 3 2 1: descending
	* 그 외: mixed
	*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringTokenizer stk = new StringTokenizer(br.readLine());

        List<Integer> nums = new ArrayList<>();

        while (stk.hasMoreTokens()) {
            int num = Integer.parseInt(stk.nextToken());
            nums.add(num);
        }

        List<Integer> ascending = new ArrayList<>(Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8));
        List<Integer> descending = new ArrayList<>(Arrays.asList(8, 7, 6, 5, 4, 3, 2, 1));

        if (ascending.equals(nums)){
            System.out.println("ascending");
        }else if (descending.equals(nums)){
            System.out.println("descending");
        }else{
            System.out.println("mixed");
        }
        
    }
}
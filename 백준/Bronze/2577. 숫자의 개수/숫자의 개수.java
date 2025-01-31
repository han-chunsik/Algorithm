import java.io.*;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
	/*
	*	A × B × C를 계산한 결과에 0부터 9까지 각각의 숫자가 몇 번씩 쓰였는지를 구하는 프로그램
	*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int a = Integer.parseInt(br.readLine());
        int b = Integer.parseInt(br.readLine());
        int c = Integer.parseInt(br.readLine());

        List<Integer> resultList = new ArrayList<>(Collections.nCopies(10, 0));
        
        for (int i=0; i<10; i++){
            for (char digit : String.valueOf(a*b*c).toCharArray()) {
                int num = Integer.valueOf(Character.toString(digit));
                if (num == i){
                    resultList.set(i, resultList.get(i)+1);
                }
            }
        }
        for (int result : resultList){
            System.out.println(result);
        }
    }
}
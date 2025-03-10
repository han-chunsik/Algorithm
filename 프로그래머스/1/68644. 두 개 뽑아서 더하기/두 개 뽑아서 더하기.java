// 입출력 예를 보면 결과 값에 중복값을 허용하지 않음
// - `5+2 = 7`, `0+7 = 7` 이지만 결과에는 하나만 있음
import java.util.HashSet;

class Solution {
    public int[] solution(int[] numbers) {
        HashSet<Integer> set = new HashSet<>();
        
        for (int i=0; i < numbers.length -1; i++) {
            for (int j=i+1; j < numbers.length; j++){
                set.add(numbers[i] + numbers[j]);
            } 
        }
        
        return set.stream().sorted().mapToInt(Integer::intValue).toArray();
    }
}
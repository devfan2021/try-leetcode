package duplicationInArray;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class DuplicationInArray {

	public int findRepeatNumber(int[] nums) {
		if (null == nums || nums.length == 0) {
			return -1;
		}

		Map<Integer, Integer> cacheMap = new HashMap<>();
		for (int i = 0; i < nums.length; i++) {
			if (cacheMap.containsKey(Integer.valueOf(nums[i]))) {
				return nums[i];
			}
			cacheMap.put(Integer.valueOf(nums[i]), 1);
		}
		return -1;
	}

	public int findRepeatNumber2(int[] nums) {
		if (null == nums || nums.length == 0) {
			return -1;
		}

		Set<Integer> cacheSet = new HashSet<>();
		for (int num : nums) {
			if (!cacheSet.add(num)) {
				return num;
			}
		}

		return -1;
	}

}

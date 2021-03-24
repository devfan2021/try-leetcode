package moreThanHalfNumber;

public class MoreThanHalfNumber {

	public int majorityElement(int[] nums) {
		if (null == nums || nums.length == 0) {
			return -1;
		}

		int votes = 0;
		int cmpValue = 0;
		for (int num : nums) {
			if (votes == 0) {
				cmpValue = num;
			}

			votes += (num == cmpValue ? 1 : -1);
		}

		// check the value
		int count = 0;
		for (int num : nums) {
			if (num == cmpValue) {
				count++;
			}
		}

		return count > (nums.length / 2) ? cmpValue : -1;
	}

}

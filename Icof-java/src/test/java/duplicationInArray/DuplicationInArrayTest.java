package duplicationInArray;

import org.junit.Assert;
import org.junit.Test;

public class DuplicationInArrayTest {

	@Test
	public void findRepeatNumber() {
		DuplicationInArray duplicationInArray = new DuplicationInArray();
		int[] nums = {1, 2, 3};
		Assert.assertEquals(-1, duplicationInArray.findRepeatNumber(nums));
		Assert.assertEquals(-1, duplicationInArray.findRepeatNumber2(nums));
		int[] nums2 = {1, 2, 3, 3};
		Assert.assertEquals(3, duplicationInArray.findRepeatNumber(nums2));
		Assert.assertEquals(3, duplicationInArray.findRepeatNumber2(nums2));
	}

}

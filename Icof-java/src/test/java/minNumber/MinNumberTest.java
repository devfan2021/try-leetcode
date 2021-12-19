package minNumber;

import org.junit.Assert;
import org.junit.Test;

public class MinNumberTest {

	@Test
	public void minNumber() {
		MinNumber minNumber = new MinNumber();
		int[] vals = {10, 4};
		Assert.assertEquals("104", minNumber.minNumber(vals));
		int[] vals2 = {10, 2};
		Assert.assertEquals("102", minNumber.minNumber(vals2));
	}

}

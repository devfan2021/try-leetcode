package minNumber;

public class MinNumber {

	public String minNumber(int[] nums) {
		String[] strs = new String[nums.length];
		for (int i = 0; i < nums.length; i++) {
			strs[i] = String.valueOf(nums[i]);
		}
		// Arrays.sort(strs, (x, y) -> (x + y).compareTo(y + x));
		quickSort(strs, 0, strs.length - 1);
		StringBuilder res = new StringBuilder();
		for (String val : strs) {
			res.append(val);
		}
		return res.toString();
	}

	private void quickSort(String[] arr, int low, int high) {
		if (low >= high) {
			return;
		}

		String pivot = arr[low];
		int left = low, right = high;
		while (left < right) {
			while (left < right && pivot.compareTo(arr[right]) <= 0) {
				right--;
			}
			if (pivot.compareTo(arr[right]) > 0) {
				arr[low] = arr[right];
				low++;
			}

			while (left < right && pivot.compareTo(arr[left]) > 0) {
				left++;
			}
			if (pivot.compareTo(arr[left]) < 0) {
				arr[right] = arr[low];
				right--;
			}
		}
		arr[left] = pivot;
		quickSort(arr, low, left - 1);
		quickSort(arr, left + 1, high);
	}

	private void quickSort(int[] arr, int left, int right) {
		if (left >= right) {
			return;
		}
		int low = left, high = right;
		int pivot = arr[low];
		// int partitionIndex = i;
		while (low < high) {
			while (pivot <= arr[high] && low < high) {
				high--;
			}
			if (arr[high] < pivot) {
				arr[low] = arr[high];
				low++;
			}

			while (pivot > arr[low] && low < high) {
				low++;
			}
			if (arr[low] > pivot) {
				arr[high] = arr[low];
				high--;
			}
		}
		arr[low] = pivot;

		quickSort(arr, left, low - 1);
		quickSort(arr, low + 1, right);
	}

	public static void main(String[] args) {
		MinNumber minNumber = new MinNumber();
		int[] vals = {8, 3, 10, 2, 7, 6, 9, 12};
		minNumber.quickSort(vals, 0, vals.length - 1);
		for (int i = 0; i < vals.length; i++) {
			System.out.print(vals[i] + ",");
		}

		System.out.println();

		int[] vals2 = {8, 3, 10, 2, 8, 6, 10, 12};
		minNumber.quickSort(vals2, 0, vals2.length - 1);
		for (int i = 0; i < vals2.length; i++) {
			System.out.print(vals2[i] + ",");
		}
	}

}

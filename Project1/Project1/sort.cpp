//#include <iostream>
//using namespace std;
//
////ð������
//template <class T>
//void bubbleSort(T* a, int m) {
//	for (int i = 0; i < m-1; i++) {
//		for (int j = 0; j < m - i - 1; j++) {
//			if (a[j] > a[j + 1]) 
//				swap(a[j], a[j + 1]);
//		}
//	}
//}
//
//void test01() {
//	int arr1[] = { 2,6,9,21,2,1,1,5 };
//	int len = sizeof(arr1) / sizeof(int);
//	bubbleSort(arr1, len);
//	for (int i = 0; i < len; i++) {
//		cout << arr1[i] << " ";
//	}
//}
//
//
//
////��������
//template <class T>
//void quickSort(T* a, int low, int height) {
//	if (low < height) {
//		int i = low, j = height;
//		int base = a[i];
//		//�ȴӻ�׼ֵ�ұ���
//		while (i < j && a[j] >= base) {
//			j--;
//		}
//		if (i<j) {
//			swap(a[j], a[i]);
//			i++;
//		}
//		//�ٴӻ�׼ֵ�����
//		while (i < j && a[i] <= base) {
//			i++;
//		}
//		if (i<j) {
//			swap(a[j], a[i]);
//			j--;
//		}
//		a[i] = base;
//
//		for (int i = 0; i < 6; i++) {
//			cout << a[i] << " ";
//		}
//		cout << endl;
//		quickSort(a, low, i - 1);
//		quickSort(a, i + 1, height);
//	}
//	
//}
//
//void test02() {
//	int arr2[] = { 10,6,4,3,5,1};
//	int len = sizeof(arr2) / sizeof(int);
//	quickSort(arr2, 0, len - 1);
//	/*for (int i = 0; i < len; i++) {
//		cout << arr2[i] << " ";
//	}*/
//}
//
//
//
////��������
//template <class T>
//void insertSort(T* a, int len) {
//	for (int i = 1; i < len; i++) {
//		int j = i - 1;
//		int tmp = a[i];
//		while (tmp < a[j] && j >= 0) {
//			a[j + 1] = a[j];
//			j--;
//		}
//		a[j + 1] = tmp;
//		for (int i = 0; i < len; i++) {
//			cout << a[i] << " ";
//		}
//		cout << endl;
//	}
//}
//
//void test03() {
//	int arr3[] = { 7,2,5,3,1,6,4 };
//	int len = sizeof(arr3) / sizeof(int);
//	insertSort(arr3, len);
//	/*for (int i = 0; i < len; i++) {
//		cout << arr3[i] << " ";
//	}*/
//
//}
//
//
//
////ѡ������
//template <class T>
//void SelectSort(T* a, int len) {
//	int begin = 0, end = len - 1;
//	while (begin < end) {
//		int maxa = begin;
//		int mina = end;
//		for (int i = begin; i <= end; i++) {
//			if (a[i] > a[maxa]) {
//				maxa = i;
//			}
//			if (a[i] < a[mina]) {
//				mina = i;
//			}
//		}
//		swap(a[begin], a[mina]);
//		if (begin == maxa) {
//			maxa = mina;
//		}
//		swap(a[end], a[maxa]);
//		for (int i = 0; i < len; i++) {
//			cout << a[i] << " ";
//		}
//		cout << endl;
//		begin++;
//		end--;
//	}
//}
//
//void test04() {
//	int arr4[] = { 3,4,7,1,2,6,8,5 };
//	int len = sizeof(arr4) / sizeof(arr4[0]);
//	SelectSort<int>(arr4, len);
//	/*for (int i = 0; i < len; i++) {
//		cout << arr4[i] << " ";
//	}*/
//}
//
//
//
////�鲢����
//template <class T>
//void merge(T* a, int low, int mid, int high, T* tmp) {//�ӿ�ϲ�
//	int i = low;
//	int j = mid + 1;
//	int k = low;
//	
//	//��ѡ��С��������tmp��ǰ�棬�������η�
//	while (i <= mid && j <= high) {
//		if (a[i] <= a[j]) {
//			tmp[k] = a[i];
//			k++;
//			i++;
//		}
//		else {
//			tmp[k] = a[j];
//			k++;
//			j++;
//		}
//	}
//
//	//�����ߣ����ź��򣩶�����Ԫ�أ�ֱ�����ź����tmp�����
//	while (i <= mid) {
//		tmp[k++] = a[i++];
//	}
//
//	//����ұߣ����ź��򣩶�����Ԫ�أ�ֱ�����ź����tmp�����
//	while (j <= high) {
//		tmp[k++] = a[j++];
//	}
//
//	//��tmp������ԭ����
//	for (int x = low; x <= high; x++) {
//		a[x] = tmp[x];
//		cout << a[x] << " ";
//	}
//	cout << endl;
//}
//
//template <class T>
//void mergeSort(T* a, int low, int high, T* tmp) {
//	int mid = (low + high) / 2;
//	if (low == high) {
//		return;
//	}
//	mergeSort(a, low, mid, tmp);
//	mergeSort(a, mid + 1, high, tmp);
//	merge(a, low, mid, high, tmp);
//	
//}
//
//void test05() {
//	int arr5[] = { 2,10,5,4,1,9,8,3,7,6 };
//	int len = sizeof(arr5) / sizeof(arr5[0]);
//	int tmp[10];
//	mergeSort(arr5, 0, len - 1, tmp);
//	/*for (int i = 0; i < len; i++) {
//		cout << arr5[i] << " ";
//	}*/
//}
//
//
//
//
//int main() {
//
//	//test01();
//	//test02();
//	//test03();
//	//test04();
//	test05();
//	/*int arr3[] = { 2,5,4,9,1};
//	quickSort<int>(arr3,0,4);
//	for (int i = 0; i < 5; i++) {
//		cout << arr3[i] << " ";
//	}*/
//	system("pause");
//	return 0;
//}
//
//

//#include <vector>
//#include <algorithm>
//#include <iostream>
//using namespace std;
//
//void MyPrint(int val)
//{
//	cout << val << endl;
//}
//
//void test01() {
//
//	vector<int> v;
//
//	v.push_back(10);
//	v.push_back(20);
//	v.push_back(30);
//	v.push_back(40);
//
//	
//	vector<int>::iterator pBegin = v.begin();//起始迭代器，指向容器中第一个数据
//	vector<int>::iterator pEnd = v.end(); //结束迭代器，指向容器元素的最后一个元素的下一个位置
//
//	/*while (pBegin != pEnd) {
//		cout << *pBegin << endl;
//		pBegin++;
//	}*/
//
//	/*for (vector<int>::iterator it = v.begin(); it != v.end(); it++) {
//		cout << *it << endl;
//	}
//	cout << endl;*/
//
//	for_each(v.begin(), v.end(), MyPrint);
//}
//
//
//int main() {
//	cout << __LINE__ << endl; //打印当前行号
//	cout << __FILE__ << endl; //打印当前源文件的文件名
//	test01();
//	system("pause");
//	return 0;
//}
//

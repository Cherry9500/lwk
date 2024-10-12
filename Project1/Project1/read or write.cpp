#include <iostream>
#include <string>
#include "Person.cpp"
using namespace std;

void test01()
{
	Person<string, int> p("Tom", 10);
	p.showPerson();
}

int main() {
	test01();
	system("pause");
	return 0;
}
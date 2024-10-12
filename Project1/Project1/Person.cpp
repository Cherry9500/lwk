//#include "person.h"
#include "Person.h" //解决方式1，包含cpp源文件
//解决方式2，将声明和实现写到一起，文件后缀名改为.hpp
//#include "person.hpp"

//构造函数 类外实现
template<class T1, class T2>
Person<T1, T2>::Person(T1 name, T2 age) {
	this->m_Name = name;
	this->m_Age = age;
}
//成员函数 类外实现
template<class T1, class T2>
void Person<T1, T2>::showPerson() {
	cout << "姓名: " << this->m_Name << " 年龄:" << this->m_Age << endl;
}


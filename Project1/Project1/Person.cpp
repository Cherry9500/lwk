//#include "person.h"
#include "Person.h" //�����ʽ1������cppԴ�ļ�
//�����ʽ2����������ʵ��д��һ���ļ���׺����Ϊ.hpp
//#include "person.hpp"

//���캯�� ����ʵ��
template<class T1, class T2>
Person<T1, T2>::Person(T1 name, T2 age) {
	this->m_Name = name;
	this->m_Age = age;
}
//��Ա���� ����ʵ��
template<class T1, class T2>
void Person<T1, T2>::showPerson() {
	cout << "����: " << this->m_Name << " ����:" << this->m_Age << endl;
}


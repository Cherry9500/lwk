#pragma once
#include <iostream>
using namespace std;

template <class T>
class MyArray {
public:
	//构造函数
	MyArray(int capacity)
	{
		cout << "MyArray的有参构造调用" << endl;
		this->m_Capacity = capacity;
		this->m_Size = 0;
		pAddress = new T[this->m_Capacity];
	}

	//拷贝构造
	MyArray(const MyArray& arr)
	{
		cout << "MyArray的拷贝构造调用" << endl;
		this->m_Capacity = arr.m_Capacity;
		this->m_Size = arr.m_Size;
		//this->pAddress = arr.pAddress;  会出现浅拷贝的问题，堆区的内存没有释放
		this->pAddress = new T[this->m_Capacity]; //深拷贝
		for (int i = 0; i < this->m_Size; i++)
		{
			//如果T为对象，而且还包含指针，必须需要重载 = 操作符，因为这个等号不是 构造 而是赋值，
			//普通类型可以直接= 但是指针类型需要深拷贝
			this->pAddress[i] = arr.pAddress[i];
		}
	}

	//operator= 防止浅拷贝，比如：a = b = c
	MyArray& operator=(const MyArray& myarray) {
		cout << "MyArray的operator= 调用" << endl;
		if (this->pAddress != NULL) {
			delete[] this->pAddress;
			this->pAddress = NULL;
			this->m_Capacity = 0;
			this->m_Size = 0;
		}

		this->m_Capacity = myarray.m_Capacity;
		this->m_Size = myarray.m_Size;
		this->pAddress = new T[this->m_Capacity]; //深拷贝
		for (int i = 0; i < this->m_Size; i++) {
			this->pAddress[i] = myarray.pAddress[i];
		}

		//返回自身
		return *this;
	}

	//重载[] 操作符 arr[0]
	T& operator [](int index)
	{
		return this->pAddress[index]; //不考虑越界，用户自己去处理
	}

	//尾插法
	void Push_back(const T& val)
	{
		if (this->m_Capacity == this->m_Size)
		{
			return;
		}
		this->pAddress[this->m_Size] = val;
		this->m_Size++;
	}

	//尾删法
	void Pop_back()
	{
		if (this->m_Size == 0)
		{
			return;
		}
		this->m_Size--;
	}

	//获取数组容量
	int getCapacity()
	{
		return this->m_Capacity;
	}
	//获取数组大小
	int getSize()
	{
		return this->m_Size;
	}
		

	//析构
	~MyArray()
	{
		cout << "MyArray的析构函数调用" << endl;
		if (this->pAddress != NULL)
		{
			delete[] this->pAddress;
			this->pAddress = NULL;
			this->m_Capacity = 0;
			this->m_Size = 0;
		}
	}


private:
	T* pAddress; //指向一个堆空间，这个空间存储真正的数据
	int m_Capacity; //容量
	int m_Size;
};


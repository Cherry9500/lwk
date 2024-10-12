#include <functional>
#include <iostream>

struct Foo {
    Foo(int num) : num_(num) {}
    void print_add(int i) const { std::cout << num_ + i << '\n'; }
    int num_;
};

void print_num(int i)
{
    std::cout << i << '\n';
}

struct PrintNum {
    void operator()(int i) const
    {
        std::cout << i << '\n';
    }
};

int main()
{
    // �洢���ɺ���
    std::function<void(int)> f_display = print_num;
    f_display(-9);

    // �洢 lambda
    std::function<void()> f_display_42 = []() { print_num(42); };
    f_display_42();

    // �洢�� std::bind ���õĽ��
    std::function<void()> f_display_31337 = std::bind(print_num, 31337);
    f_display_31337();

    // �洢����Ա�����ĵ���
    std::function<void(const Foo&, int)> f_add_display = &Foo::print_add;
    const Foo foo(314159);
    f_add_display(foo, 1);

    // �洢�����ݳ�Ա�������ĵ���
    std::function<int(Foo const&)> f_num = &Foo::num_;
    std::cout << "num_: " << f_num(foo) << '\n';

    // �洢����Ա����������ĵ���
    std::function<void(int)> f_add_display2 = std::bind(&Foo::print_add, foo, std::placeholders::_1);
    f_add_display2(2);

    // �洢����Ա�����Ͷ���ָ��ĵ���
    std::function<void(int)> f_add_display3 = std::bind(&Foo::print_add, &foo, std::placeholders::_1);
    f_add_display3(3);

    // �洢����������ĵ���
    std::function<void(int)> f_display_obj = PrintNum();
    f_display_obj(18);
}





//#include<iostream>
////#include<move>
//using namespace std;
//
//void Func(int& x) { cout << "��ֵ����" << endl; }
//void Func(const int& x) { cout << "const ��ֵ����" << endl; }
//
//void Func(int&& x) { cout << "��ֵ����" << endl; }
//void Func(const int&& x) { cout << "const ��ֵ����" << endl; }
//
//template<typename T>
//void PerfectForward(T&& t)
//{
//	//Func(t);//û��ʹ��forward��������ֵ�����ԣ��˻�Ϊ��ֵ
//	Func(forward<T>(t));
//}
//
//int main()
//{
//	PerfectForward(1);//��ֵ
//	int a = 10;
//	PerfectForward(a);
//	PerfectForward(move(a));
//
//	const int b = 20;
//	PerfectForward(b);
//	PerfectForward(move(b));
//	return 0;
//}

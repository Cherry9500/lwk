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
    // 存储自由函数
    std::function<void(int)> f_display = print_num;
    f_display(-9);

    // 存储 lambda
    std::function<void()> f_display_42 = []() { print_num(42); };
    f_display_42();

    // 存储到 std::bind 调用的结果
    std::function<void()> f_display_31337 = std::bind(print_num, 31337);
    f_display_31337();

    // 存储到成员函数的调用
    std::function<void(const Foo&, int)> f_add_display = &Foo::print_add;
    const Foo foo(314159);
    f_add_display(foo, 1);

    // 存储到数据成员访问器的调用
    std::function<int(Foo const&)> f_num = &Foo::num_;
    std::cout << "num_: " << f_num(foo) << '\n';

    // 存储到成员函数及对象的调用
    std::function<void(int)> f_add_display2 = std::bind(&Foo::print_add, foo, std::placeholders::_1);
    f_add_display2(2);

    // 存储到成员函数和对象指针的调用
    std::function<void(int)> f_add_display3 = std::bind(&Foo::print_add, &foo, std::placeholders::_1);
    f_add_display3(3);

    // 存储到函数对象的调用
    std::function<void(int)> f_display_obj = PrintNum();
    f_display_obj(18);
}





//#include<iostream>
////#include<move>
//using namespace std;
//
//void Func(int& x) { cout << "左值引用" << endl; }
//void Func(const int& x) { cout << "const 左值引用" << endl; }
//
//void Func(int&& x) { cout << "右值引用" << endl; }
//void Func(const int&& x) { cout << "const 右值引用" << endl; }
//
//template<typename T>
//void PerfectForward(T&& t)
//{
//	//Func(t);//没有使用forward保持其右值的属性，退化为左值
//	Func(forward<T>(t));
//}
//
//int main()
//{
//	PerfectForward(1);//右值
//	int a = 10;
//	PerfectForward(a);
//	PerfectForward(move(a));
//
//	const int b = 20;
//	PerfectForward(b);
//	PerfectForward(move(b));
//	return 0;
//}

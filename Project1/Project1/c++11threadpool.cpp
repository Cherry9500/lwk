//#include<iostream>
//#include<thread>
//#include<chrono>
//#include<mutex>
//#include<queue>
//#include<vector>
//#include<functional>
//#include<condition_variable>
//
//class Threadpool
//{
//public:
//	Threadpool(int numThreads) :stop(false) {
//		for (int i = 0; i < numThreads; i++) {
//			threads.emplace_back([this] { //emplace_back���push_back���ʡ��Դ
//				while (1) {
//					std::unique_lock<std::mutex> lock(mtx);
//					cv.wait(lock, [this] {return !tasks.empty() || stop; });
//
//					if (stop && tasks.empty())
//						return;
//
//					std::function<void()> task(std::move(tasks.front())); //move��ֵ����
//					tasks.pop();
//					//lock.unlock();
//					task();
//				}
//			}); 
//		}
//	}
//	~Threadpool() {
//		{
//			std::unique_lock<std::mutex> lock(mtx);
//			stop = true;
//		}
//		cv.notify_all();
//		for (auto& t : threads)
//			t.join();
//	};
//
//	//�ɱ����ģ��
//	template<class F, class...Args>
//	void enqueue(F&& f, Args&&...args) { 
//		//&&���������ã��ȿ���ֵҲ����ֵ
//		//forward������ת��
//		//blind��������
//		std::function<void()> task = std::bind(std::forward<F>(f), std::forward<Args>(args)...); 
//		{
//			std::unique_lock<std::mutex> lock(mtx);
//			tasks.emplace(std::move(task));
//		}
//		cv.notify_all();
//	}
//private: 
//	std::vector<std::thread> threads; 
//	std::queue<std::function<void()>> tasks; //c++11 function����ģ��
//
//	std::mutex mtx; 
//	std::condition_variable cv; 
//	
//	bool stop; 
//};
//
//int main() {
//	Threadpool pool(4); //��ʼ��4���߳�
//	//��10������
//	for (int i = 0; i < 10; i++) {
//		pool.enqueue([i] {
//			std::cout << "task : " << i << " is running " << std::endl;
//			std::this_thread::sleep_for(std::chrono::seconds(2));
//			std::cout << "task : " << i << " is done " << std::endl;
//			});
//	}
//	return 0;
//}
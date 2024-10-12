//#include<iostream>
//#include<thread>
//#include<chrono>
//#include<mutex>
//#include<condition_variable>
////using namespace std;
//
//std::mutex mtx;
//std::condition_variable cv;
//bool isReady = false;
//int data = 0;
//
//void producer() {
//	for (int i = 0; i < 5; i++) {
//		std::this_thread::sleep_for(std::chrono::seconds(2)); //ģ��ȴ�����
//		std::lock_guard<std::mutex> lock(mtx);
//		data = i * 10; //��������
//		isReady = true; //����״̬
//
//		std::cout << "Producer: " << data << std::endl;
//		cv.notify_all();
//	}
//}
//
//void consumer() {
//	while (true) {
//		std::unique_lock<std::mutex> lock(mtx);
//
//		cv.wait(lock, [] { return isReady; });
//
//		std::cout << "Consumer: " << data << std::endl;
//
//		isReady = false;
//
//		//lock.unlock();
//
//		//�˳�
//	}
//}
//
//int main(){
//	std::thread producer_thread(producer);
//	std::thread consumer_thread(consumer);
//
//	producer_thread.join();
//	consumer_thread.join();
//
//	system("pause");
//	return 0;
//}
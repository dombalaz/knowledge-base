# Simple Thread Safety

A data race in C++ occurs when two or more threads access the same memory location simultaneously, where at least one thread performs a write operation. Data race results in **undefined behaviour**. To prevent data races, atomic operations or mutexes has to be used, to prevent that.

This means that we cannot be reading from a `bool` that we are changing in another thread, like some stopping mechanism.

**Thread sanitizer** is your friend! Compile with `-fsanitize=thread`. Use it always during testing.

Bad:
```c++
#include <thread>

int main()
{
  bool stop{false};
  auto t = std::thread([&]() { stop = true; });
  while (!stop);
  t.join();
  return 0;
}
```
Compiling and running with `clang-19.1.0`:
```bash
ASM generation compiler returned: 0
Execution build compiler returned: 0
Program returned: 66
/opt/compiler-explorer/clang-19.1.0/bin/llvm-symbolizer: error: '[stack]': No such file or directory
==================
WARNING: ThreadSanitizer: data race (pid=1)
  Read of size 1 at 0x7fffffffe9fb by main thread:
    #0 main /app/example.cpp:7:11 (output.s+0xe624a)

  Previous write of size 1 at 0x7fffffffe9fb by thread T1:
    #0 main::$_0::operator()() const /app/example.cpp:6:37 (output.s+0xe6839)
    #1 void std::__invoke_impl<void, main::$_0>(std::__invoke_other, main::$_0&&) /opt/compiler-explorer/gcc-14.2.0/lib/gcc/x86_64-linux-gnu/14.2.0/../../../../include/c++/14.2.0/bits/invoke.h:61:14 (output.s+0xe67e5)
    #2 std::__invoke_result<main::$_0>::type std::__invoke<main::$_0>(main::$_0&&) /opt/compiler-explorer/gcc-14.2.0/lib/gcc/x86_64-linux-gnu/14.2.0/../../../../include/c++/14.2.0/bits/invoke.h:96:14 (output.s+0xe6755)
    #3 void std::thread::_Invoker<std::tuple<main::$_0>>::_M_invoke<0ul>(std::_Index_tuple<0ul>) /opt/compiler-explorer/gcc-14.2.0/lib/gcc/x86_64-linux-gnu/14.2.0/../../../../include/c++/14.2.0/bits/std_thread.h:301:13 (output.s+0xe670d)
    #4 std::thread::_Invoker<std::tuple<main::$_0>>::operator()() /opt/compiler-explorer/gcc-14.2.0/lib/gcc/x86_64-linux-gnu/14.2.0/../../../../include/c++/14.2.0/bits/std_thread.h:308:11 (output.s+0xe66b5)
    #5 std::thread::_State_impl<std::thread::_Invoker<std::tuple<main::$_0>>>::_M_run() /opt/compiler-explorer/gcc-14.2.0/lib/gcc/x86_64-linux-gnu/14.2.0/../../../../include/c++/14.2.0/bits/std_thread.h:253:13 (output.s+0xe6579)
    #6 <null> <null> (libstdc++.so.6+0xed0e3) (BuildId: 998334304023149e8c44e633d4a2c69800a2eb79)

  Location is stack of main thread.

  Location is global '??' at 0x7ffffffde000 ([stack]+0x209fb)

  Thread T1 (tid=3, finished) created by main thread at:
    #0 pthread_create /root/llvm-project/compiler-rt/lib/tsan/rtl/tsan_interceptors_posix.cpp:1023:3 (output.s+0x61fc1)
    #1 std::thread::_M_start_thread(std::unique_ptr<std::thread::_State, std::default_delete<std::thread::_State>>, void (*)()) <null> (libstdc++.so.6+0xed1b8) (BuildId: 998334304023149e8c44e633d4a2c69800a2eb79)
    #2 main /app/example.cpp:6:12 (output.s+0xe623d)

SUMMARY: ThreadSanitizer: data race /app/example.cpp:7:11 in main
==================
ThreadSanitizer: reported 1 warnings
```

Check example also on [godbolt](https://godbolt.org/z/d1E6x6GjP).

Fix with atomics:
```c++
#include <atomic>
#include <thread>

int main()
{
  std::atomic<bool> stop{false};
  auto t = std::thread([&]() { stop = true; });
  while (!stop);
  t.join();
  return 0;
}
```

跳表是经过改造后的链表结构，可以用来支持二分查找算法。

跳表支持高效的动态插入和删除

redis的有序集合是通过跳表实现的，还用到了散列表。

跳表使用空间换时间的设计思路，通过构建多级索引来提高查询的效率，实现了基于链表的“二分查找”。跳表是一种动态数据结构，支持快速的插入、删除、查找操作，时间复杂度都是O（logn）。

跳表的空间复杂度是O(n)。跳表是实现非常灵活，可以通过改变索引构建策略，有效平衡执行效率和内存消耗。
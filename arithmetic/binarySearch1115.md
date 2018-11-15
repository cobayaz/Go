
二分查找（binary Search）算法，又叫折半查找算法。<br>
二分查找针对的是一个有序的数据集合，查找思想有点类似分治思想。每次通过跟区间的中间元素相比，将待查找的区间缩小为之前的一半，直到找到要查找的元素，或者区间别缩小为0。<br>

二分查找算法的时间复杂度是O(logn)。每次查找后数据都会缩小为原来的一半，也就是会除以2.

logn是一个非常“恐怖”的数量级，即便n非常非常打，对应的logn也很小。在42亿个数据中用二分查找一个数据，最多也是比较32次。

二分查找依赖的是顺序表就结构，简单的说就是数组。

二分查找不能依赖链表等其他数据结构。主要是二分查找算法需要依据下标随机访问元素。数组按照下标随机访问数据的时间复杂度是O（1），而链表随机访问的时间复杂度是O(n)。如果数据使用链表存储，二分查找的时间复杂度就会变得很高。

硬说，二分查找只能用在数据是通过顺序表来存储的数据结构上。如果你的数据是通过其他数据结构存储的，则无法应用二分查找。

二分查找只能用在插入、删除操作不频繁，一次排序多次查找的场景中。针对动态变化的数据集合，二分查找将不再使用。动态数据集合需要用到二叉树相关的算法知识。

数据量太小也不适合二分查找。

二分查找的底层需要依赖数组这种数据结构，而数组为了支持随机访问的特性，要求内存空间连续，对内存的要求比较苛刻。

二分查找虽然性能比较优秀，但应用场景比较有限。底层必须依赖数组，并且还要求数据是有序的。对于较小规模的数据查找，我们直接使用顺序遍历就可以了，二分查找的优势并不明显。二分查找更加适合处理静态数据，也就是没有频繁的数据插入、删除操作。
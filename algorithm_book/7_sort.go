package algorithm_book


/*
1.算法分析
1.1时间复杂度
1.2空间复杂度


排序方法 平均情况 最好情况 最坏情况 辅助空间 稳定性
冒泡排序 n^2	     n	     n^2     1       稳定
选择排序 n^2      n^2     n^2     1       不稳定
插入排序 n^2      n       n^2     1       稳定
堆排序   nlogn    nlogn   nlogn   1       不稳定
归并排序 nlogn    nlogn   nlogn   n       稳定
快速排序 nlogn    nlogn   n^2     1       不稳定

快排最坏情况：选择的轴是最小的值或最大的值，造成划分的分区只有一个。退化成冒泡排序
*/

## npw介绍
npw：npartword，golang实现中文分词系统，主体分词逻辑有两个部分。   

- 1.前缀树查找字典，通过disctance或mmseg算法过滤分词    

- 2.维特比算法解隐马尔可夫模型，对词进行隐状态标注分词   

分词服务系统的实现，通过nmid的worker方式实现。       

- 1.分词系统服务端，需要实现nmid的worker，服务的实现十分简单，无需考虑通信问题，这些nmid解决。   

- 2.分词系统服务调用，只要通过nmid的client调用即可，任何nmid的client都可以随时跨服务器的使用分词系统。    


